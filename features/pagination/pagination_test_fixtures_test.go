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
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/gomega"
)

const (
	pageSize         = 10
	StatusBrokenJson = 600
	StatusBadIO      = 601
)

var errorCodes = []int{
	http.StatusBadRequest,
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusNotFound,
	http.StatusTooManyRequests,
	http.StatusInternalServerError,
	http.StatusBadGateway,
	http.StatusGatewayTimeout,
	StatusBrokenJson,
	StatusBadIO,
}

var testCases = []struct {
	descrition  string
	expectPages int
	expectItems int
}{
	{
		descrition:  "Confirms result is correct for an empty page",
		expectPages: 1, //  Need at least 1 empty page to know there are no more results
		expectItems: 0,
	},
	{
		descrition:  "Confirms result is correct for a partial page",
		expectPages: 1,
		expectItems: 1,
	},
	{
		descrition:  "Confirms result is correct for a page size minus one",
		expectPages: 1,
		expectItems: pageSize - 1,
	},
	{
		descrition:  "Confirms result is correct for a single page",
		expectPages: 1,
		expectItems: pageSize,
	},
	{
		descrition:  "Confirms result is correct for a page size plus one",
		expectPages: 2,
		expectItems: pageSize + 1,
	},
	{
		descrition:  "Confirms result is correct multiple pages exactly",
		expectPages: 3,
		expectItems: 3 * pageSize,
	},
	{
		descrition:  "Confirms result is correct multiple pages plus one",
		expectPages: 4,
		expectItems: 3*pageSize + 1,
	},
	{
		descrition:  "Confirms result is correct multiple pages minus one",
		expectPages: 4,
		expectItems: 4*pageSize - 1,
	},
}

type mockServiceKey string

var msKey = mockServiceKey("mock")

type mockDoc struct {
	N  int
	ID *string
}

// mockService is a struct holding mocked results
type mockService struct {
	items      []mockDoc
	errorItem  int
	err        error
	statusCode int
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
		item := mockDoc{N: i + 1}
		id := fmt.Sprintf("%02d", item.N)
		item.ID = &id
		s.items = append(s.items, item)
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

// setHTTPError prepares an http error to be returned from request function after a given item
func (s *mockService) setHTTPError(statusCode int, errorItem int) {
	s.statusCode = statusCode
	s.err = errors.New(statusText(statusCode))
	s.errorItem = errorItem
}

// statusText converts an error code into expecrted error string
func statusText(code int) string {
	switch code {
	case StatusBrokenJson:
		return "An error occurred while processing the HTTP response"
	case StatusBadIO:
		return "An error occurred while reading the response body"
	default:
		return http.StatusText(code)
	}
}

// duplicateItem duplicates a last item in the items collection
func (s *mockService) duplicateItem() {
	lastItem := slices.Clone(s.items[len(s.items)-1:])
	s.items = append(s.items, lastItem...)
}

// mockServerCallback is a callback function for httptest server in pager concrete tests
func mockServerCallback(w http.ResponseWriter, r *http.Request, ms *mockService) {
	Expect(r.Method).To(Equal("POST"))

	body, err := io.ReadAll(r.Body)
	Expect(err).ShouldNot(HaveOccurred())

	q := struct {
		Limit    int    `json:"limit"`
		StartKey string `json:"start_key"`
		Bookmark string `json:"bookmark"`
	}{}
	err = json.Unmarshal(body, &q)
	Expect(err).ShouldNot(HaveOccurred())

	startKey := 0
	if q.StartKey != "" {
		startKey, err = strconv.Atoi(q.StartKey)
		Expect(err).ShouldNot(HaveOccurred())
	} else if q.Bookmark != "" {
		startKey, err = strconv.Atoi(q.Bookmark)
		Expect(err).ShouldNot(HaveOccurred())
	}

	var statusCode int
	var data []byte

	items, err := ms.getItems(startKey, q.Limit)
	if err != nil {
		statusCode = ms.statusCode
		data = []byte(err.Error())

		if statusCode == StatusBadIO {
			w.WriteHeader(http.StatusOK)
			hj, ok := w.(http.Hijacker)
			if !ok {
				http.Error(w, "can't create hijack rw", http.StatusInternalServerError)
				return
			}
			conn, _, err := hj.Hijack()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = conn.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		if statusCode == StatusBrokenJson {
			statusCode = http.StatusOK
			data = []byte(`{`)
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(data)
		return
	} else {
		currentPath := r.URL.EscapedPath()
		statusCode = http.StatusOK
		if strings.Contains(currentPath, "_all_docs") || strings.Contains(currentPath, "_view") {
			rows := make([]cloudantv1.DocsResultRow, len(items))
			for i, doc := range ms.getDocuments(items) {
				rows[i] = cloudantv1.DocsResultRow{
					ID:  doc.ID,
					Key: doc.ID,
					Doc: &doc,
				}
			}
			total := int64(len(ms.items))
			resp := cloudantv1.AllDocsResult{
				Rows:      rows,
				TotalRows: &total,
			}
			jsonResp, err := json.Marshal(&resp)
			Expect(err).ShouldNot(HaveOccurred())
			data = jsonResp
		} else if strings.Contains(currentPath, "_find") {
			rows := make([]cloudantv1.Document, len(items))
			bookmark := ""
			for i, doc := range ms.getDocuments(items) {
				rows[i] = cloudantv1.Document{
					ID: doc.ID,
				}
				bookmark = strconv.Itoa(items[i].N + 1)
			}
			resp := cloudantv1.FindResult{
				Docs:     rows,
				Bookmark: &bookmark,
			}
			jsonResp, err := json.Marshal(&resp)
			Expect(err).ShouldNot(HaveOccurred())
			data = jsonResp
		} else if strings.Contains(currentPath, "_search") {
			rows := make([]cloudantv1.SearchResultRow, len(items))
			bookmark := ""
			for i, doc := range ms.getDocuments(items) {
				rows[i] = cloudantv1.SearchResultRow{
					ID:  doc.ID,
					Doc: &doc,
				}
				bookmark = strconv.Itoa(items[i].N + 1)
			}
			total := int64(len(ms.items))
			resp := cloudantv1.SearchResult{
				Rows:      rows,
				TotalRows: &total,
				Bookmark:  &bookmark,
			}
			jsonResp, err := json.Marshal(&resp)
			Expect(err).ShouldNot(HaveOccurred())
			data = jsonResp
		} else {
			statusCode = http.StatusNotFound
		}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	//nolint:errcheck
	w.Write(data)
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

// newTestBookmarkPager returns a bookmarkPager configured for tests
func newTestBookmarkPager(o *cloudantv1.PostFindOptions) *bookmarkPager[*cloudantv1.PostFindOptions, *cloudantv1.FindResult, cloudantv1.Document] {
	opts := *o
	return &bookmarkPager[*cloudantv1.PostFindOptions, *cloudantv1.FindResult, cloudantv1.Document]{
		options:     &opts,
		hasNextPage: true,
		requestFunction: func(ctx context.Context, o *cloudantv1.PostFindOptions) (*cloudantv1.FindResult, *core.DetailedResponse, error) {
			ms, err := fromContext(ctx)
			if err != nil {
				return nil, nil, err
			}

			limit := int(*opts.Limit)
			page := 0
			if opts.Bookmark != nil {
				if i, err := strconv.Atoi(*opts.Bookmark); err == nil {
					page = i
				}
			}
			items, err := ms.getItems(page*limit+1, limit)
			if err != nil {
				return nil, nil, ms.err
			}

			docs := ms.getDocuments(items)
			bookmark := fmt.Sprintf("%d", page+1)
			return &cloudantv1.FindResult{Docs: docs, Bookmark: &bookmark}, nil, nil
		},
		resultItemsGetter: func(result *cloudantv1.FindResult) []cloudantv1.Document { return result.Docs },
		bookmarkGetter:    func(result *cloudantv1.FindResult) string { return *result.Bookmark },
		bookmarkSetter:    opts.SetBookmark,
		optionsCloner: func(o *cloudantv1.PostFindOptions) *cloudantv1.PostFindOptions {
			opts := *o
			return &opts
		},
		limitGetter: func() *int64 { return opts.Limit },
		limitSetter: opts.SetLimit,
	}
}

func runGetNextAssertion[T paginatedRow](pager Pager[T], expectPages, expectItems int) {
	pageCount := 0
	items := 0
	uniqueItems := make(map[string]bool, 0)
	for pager.HasNext() {
		rows, err := pager.GetNext()
		Expect(err).ShouldNot(HaveOccurred())
		pageCount += 1
		items += len(rows)
		for _, doc := range rows {
			data, err := json.Marshal(doc)
			Expect(err).ShouldNot(HaveOccurred())
			uniqueItems[fmt.Sprintf("%x", md5.Sum(data))] = true
		}
	}

	Expect(pageCount).To(Equal(expectPages))
	Expect(items).To(Equal(expectItems))
	Expect(uniqueItems).Should(HaveLen(expectItems))
}

func runGetNextWithErrorAssertion[T paginatedRow](pager Pager[T], expectedError string, expectItems int) {
	// assertion for an error on a second page
	if expectItems > pageSize {
		items, err := pager.GetNext()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).Should(HaveLen(pageSize))
	}

	page, err := pager.GetNext()

	Expect(err).Should(HaveOccurred())
	Expect(err).Should(MatchError(ContainSubstring(expectedError)))
	Expect(page).Should(BeEmpty())
}
