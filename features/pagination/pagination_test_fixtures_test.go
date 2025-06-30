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
	"strconv"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

type testPager struct {
	options   *cloudantv1.PostFindOptions
	items     []cloudantv1.Document
	cycle     int
	errorItem int
	err       error
}

func newTestPager(o *cloudantv1.PostFindOptions) *testPager {
	pager := &testPager{
		items: make([]cloudantv1.Document, 0),
	}
	pager.setOptions(o)
	return pager
}

// makeItems generated a given number of documents. It is not a part of pager interface
func (p *testPager) makeItems(itemsNum int) {
	for i := range itemsNum {
		id := fmt.Sprintf("%02d", i+1)
		p.items = append(p.items, cloudantv1.Document{ID: &id})
	}
}

// setError prepares an error to be returned from request function after given item
func (p *testPager) setError(err error, errorItem int) {
	p.err = err
	p.errorItem = errorItem
}

func (p *testPager) nextRequestFunction(ctx context.Context) (*cloudantv1.FindResult, error) {
	pageSize := int(*p.getLimit())
	cycle := 0
	if p.options.Bookmark != nil {
		if i, err := strconv.Atoi(*p.options.Bookmark); err == nil {
			cycle = i
		}
	}
	start := cycle * pageSize
	if start > len(p.items) {
		start = len(p.items)
	}
	end := start + pageSize
	if end > len(p.items) {
		end = len(p.items)
	}
	// error intercepter
	if p.err != nil && p.errorItem >= start && p.errorItem <= end {
		return nil, p.err
	}
	items := p.items[start:end]
	bookmark := fmt.Sprintf("%d", cycle+1)
	return &cloudantv1.FindResult{Docs: items, Bookmark: &bookmark}, nil
}

func (p *testPager) itemsGetter(result *cloudantv1.FindResult) []cloudantv1.Document {
	return result.Docs
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
