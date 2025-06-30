/**
 * © Copyright IBM Corporation 2025. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package features

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

type mockServiceKey string

var msKey = mockServiceKey("mock")

type mockDoc struct {
	ID *string
}

// mockService is a struct holding mocked results
type mockService struct {
	items     []mockDoc
	errorItem int
	err       error
}

// newMockService creates a new mock service for testing
func newMockService() *mockService {
	return &mockService{}
}

// toContext builds a context that keeps a given mock service
func toContext(s *mockService) context.Context {
	return context.WithValue(context.Background(), msKey, s)
}

// fromContext restores mock service from given context
func fromContext(ctx context.Context) (*mockService, error) {
	if ms := ctx.Value(msKey); ms != nil {
		if s, ok := ms.(*mockService); ok {
			return s, nil
		}
		return nil, fmt.Errorf("mock service on the context is not an instance of *mockService")
	}
	return nil, fmt.Errorf("can't find mock service on the context")
}

// makeItems generated a given number of mock items
func (s *mockService) makeItems(itemsNum int) {
	for i := range itemsNum {
		id := fmt.Sprintf("%02d", i+1)
		s.items = append(s.items, mockDoc{ID: &id})
	}
}

// getItems mocks the actual service call, returning either requested items or an error
func (s *mockService) getItems(start, limit int) ([]mockDoc, error) {
	acc := make([]mockDoc, 0)
	for i, d := range s.items {
		if s.err != nil && s.errorItem == i+1 {
			return nil, s.err
		}
		if i+1 >= start {
			acc = append(acc, d)
		}
		if len(acc) == limit {
			return acc, nil
		}
	}
	return acc, nil
}

// getDocuments is a converter from slice of mock documents to a slice of cloudantv1.Document
func (s *mockService) getDocuments(items []mockDoc) []cloudantv1.Document {
	docs := make([]cloudantv1.Document, len(items))
	for i, d := range items {
		docs[i] = cloudantv1.Document{ID: d.ID}
	}
	return docs
}

// documents returns all mock service's documents as a slice of cloudantv1.Document
func (s *mockService) documents() []cloudantv1.Document {
	return s.getDocuments(s.items)
}

// getViewRows is a converter from slice of mock documents to a slice of cloudantv1.ViewResultRow
func (s *mockService) getViewRows(items []mockDoc) []cloudantv1.ViewResultRow {
	rows := make([]cloudantv1.ViewResultRow, len(items))
	for i, d := range items {
		rows[i] = cloudantv1.ViewResultRow{ID: d.ID, Key: *d.ID, Value: *d.ID}
	}
	return rows
}

// viewRows returns all mock service's documents as a slice of cloudantv1.ViewResultRow
func (s *mockService) viewRows() []cloudantv1.ViewResultRow {
	return s.getViewRows(s.items)
}

// setError prepares an error to be returned from request function after a given item
func (s *mockService) setError(err error, errorItem int) {
	s.err = err
	s.errorItem = errorItem
}

// duplicateItem duplicates a last item in the items collection
func (s *mockService) duplicateItem() {
	lastItem := slices.Clone(s.items[len(s.items)-1:])
	s.items = append(s.items, lastItem...)
}

// testPager is pager implementation to test basePager functionality
type testPager struct {
	options     *cloudantv1.PostFindOptions
	items       []cloudantv1.Document
	hasNextPage bool
}

func newTestPager(o *cloudantv1.PostFindOptions) *testPager {
	pager := &testPager{
		items:       make([]cloudantv1.Document, 0),
		hasNextPage: true,
	}
	pager.setOptions(o)
	return pager
}

func (p *testPager) nextRequestFunction(ctx context.Context) (*cloudantv1.FindResult, error) {
	ms, err := fromContext(ctx)
	if err != nil {
		return nil, err
	}

	limit := int(*p.getLimit())
	page := 0
	if p.options.Bookmark != nil {
		if i, err := strconv.Atoi(*p.options.Bookmark); err == nil {
			page = i
		}
	}

	items, err := ms.getItems(page*limit+1, limit)
	if err != nil {
		return nil, ms.err
	}

	if len(items) < limit {
		p.hasNextPage = false
	}

	docs := ms.getDocuments(items)
	bookmark := fmt.Sprintf("%d", page+1)
	return &cloudantv1.FindResult{Docs: docs, Bookmark: &bookmark}, nil
}

func (p *testPager) itemsGetter(result *cloudantv1.FindResult) ([]cloudantv1.Document, error) {
	return result.Docs, nil
}

func (p *testPager) hasNext() bool {
	return p.hasNextPage
}

func (p *testPager) getOptions() *cloudantv1.PostFindOptions {
	opts := *p.options
	return &opts
}

func (p *testPager) setOptions(o *cloudantv1.PostFindOptions) {
	opts := *o
	p.options = &opts
}

func (p *testPager) setNextPageOptions(result *cloudantv1.FindResult) {
	p.options.SetBookmark(*result.Bookmark)
}

func (p *testPager) getLimit() *int64 {
	return p.options.Limit
}

func (p *testPager) setLimit(pageSize int64) {
	p.options.SetLimit(pageSize)
}

// newTestKeyPager returns a keyPager configured for tests
func newTestKeyPager(o *cloudantv1.PostViewOptions) *keyPager[*cloudantv1.PostViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow] {
	opts := *o
	return &keyPager[*cloudantv1.PostViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow]{
		options:     &opts,
		hasNextPage: true,
		requestFunction: func(ctx context.Context, o *cloudantv1.PostViewOptions) (*cloudantv1.ViewResult, *core.DetailedResponse, error) {
			ms, err := fromContext(ctx)
			if err != nil {
				return nil, nil, err
			}

			limit := int(*opts.Limit)
			startKey := 1
			if opts.StartKeyDocID != nil {
				if i, err := strconv.Atoi(*opts.StartKeyDocID); err == nil {
					startKey = i
				}
			}

			items, err := ms.getItems(startKey, limit)
			if err != nil {
				return nil, nil, ms.err
			}

			rows := ms.getViewRows(items)
			return &cloudantv1.ViewResult{Rows: rows}, nil, nil
		},
		resultItemsGetter:   func(result *cloudantv1.ViewResult) []cloudantv1.ViewResultRow { return result.Rows },
		startViewKeyGetter:  func(item cloudantv1.ViewResultRow) any { return item.Key },
		startViewKeySetter:  opts.SetStartKey,
		startKeyDocIDGetter: func(item cloudantv1.ViewResultRow) string { return *item.ID },
		startKeyDocIDSetter: opts.SetStartKeyDocID,
		optionsCloner: func(o *cloudantv1.PostViewOptions) *cloudantv1.PostViewOptions {
			opts := *o
			return &opts
		},
		limitGetter: func() *int64 { return opts.Limit },
		limitSetter: opts.SetLimit,
	}
}
