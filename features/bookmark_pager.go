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

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

var (
	bookmarkPagerValidationRules = map[string]string{
		"Limit": limitValidationRule,
	}
	searchPagerValidationRules = map[string]string{
		"Counts":     "isdefault",
		"GroupField": "isdefault",
		"GroupLimit": "isdefault",
		"GroupSort":  "isdefault",
		"Ranges":     "isdefault",
		"Limit":      limitValidationRule,
	}
)

type bookmarkPagerOptions interface {
	FindPagerOptions | SearchPagerOptions
}

type bookmarkRequestResult interface {
	*cloudantv1.FindResult | *cloudantv1.SearchResult
}

type bookmarkPaginatedRow interface {
	cloudantv1.Document | cloudantv1.SearchResultRow
}

type bookmarkPager[O bookmarkPagerOptions, R bookmarkRequestResult, T bookmarkPaginatedRow] struct {
	service           *cloudantv1.CloudantV1
	options           O
	hasNextPage       bool
	requestFunction   func(context.Context, O) (R, *core.DetailedResponse, error)
	resultItemsGetter func(R) []T
	bookmarkGetter    func(R) string
	bookmarkSetter    func(string) O
	optionsCloner     func(O) O
	limitGetter       func() *int64
	limitSetter       func(int64) O
	skipSetter        func(int64) O
}

func (p *bookmarkPager[O, R, T]) nextRequestFunction(ctx context.Context) (R, error) {
	result, _, err := p.requestFunction(ctx, p.options)
	return result, err
}

func (p *bookmarkPager[O, R, T]) itemsGetter(result R) ([]T, error) {
	items := p.resultItemsGetter(result)
	if p.limitGetter() != nil && len(items) < int(*p.limitGetter()) {
		p.hasNextPage = false
	}
	return items, nil
}

func (p *bookmarkPager[O, R, T]) hasNext() bool {
	return p.hasNextPage
}

func (p *bookmarkPager[O, R, T]) setNextPageOptions(result R) {
	if p.skipSetter != nil {
		p.skipSetter(0)
	}
	bookmark := p.bookmarkGetter(result)
	p.bookmarkSetter(bookmark)
}

func (p *bookmarkPager[O, R, T]) getOptions() O {
	return p.optionsCloner(p.options)
}

func (p *bookmarkPager[O, R, T]) setOptions(o O) {
	p.options = p.optionsCloner(o)
}

func (p *bookmarkPager[O, R, T]) getLimit() *int64 {
	return p.limitGetter()
}

func (p *bookmarkPager[O, R, T]) setLimit(pageSize int64) {
	p.limitSetter(pageSize)
}
