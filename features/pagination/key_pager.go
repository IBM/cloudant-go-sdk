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

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

type keyPagerOptions interface {
	AllDocsPagerOptions | ViewPagerOptions
}

type keyRequestResult interface {
	*cloudantv1.AllDocsResult | *cloudantv1.ViewResult
}

type keyPaginatedRow interface {
	cloudantv1.DocsResultRow | cloudantv1.ViewResultRow
}

type keyPager[O keyPagerOptions, R keyRequestResult, T keyPaginatedRow] struct {
	service             *cloudantv1.CloudantV1
	options             O
	hasNextPage         bool
	requestFunction     func(context.Context, O) (R, *core.DetailedResponse, error)
	resultItemsGetter   func(R) []T
	startKeyGetter      func(T) string
	startKeySetter      func(string) O
	startViewKeyGetter  func(T) any
	startViewKeySetter  func(any) O
	startKeyDocIDGetter func(T) string
	startKeyDocIDSetter func(string) O
	optionsCloner       func(O) O
	limitGetter         func() *int64
	limitSetter         func(int64) O
}

func (p *keyPager[O, R, T]) nextRequestFunction(ctx context.Context) (R, error) {
	result, _, err := p.requestFunction(ctx, p.options)
	return result, err
}

func (p *keyPager[O, R, T]) itemsGetter(result R) ([]T, error) {
	items := p.resultItemsGetter(result)
	if p.limitGetter() != nil && len(items) < int(*p.limitGetter()) {
		p.hasNextPage = false
		return items, nil
	}

	var err error
	itemsNum := len(items) - 1
	if p.startViewKeyGetter != nil && p.startKeyDocIDGetter != nil {
		lastItem := items[itemsNum]
		penultimateItem := items[itemsNum-1]
		lID := p.startKeyDocIDGetter(lastItem)
		pID := p.startKeyDocIDGetter(penultimateItem)
		lKey := p.startViewKeyGetter(lastItem)
		pKey := p.startViewKeyGetter(penultimateItem)
		if lID == pID && pKey == lKey {
			err = fmt.Errorf("cannot paginate on a boundary containing identical keys %q and document IDs %q", lKey, lID)
		}
	}

	return items[:itemsNum], err
}

func (p *keyPager[O, R, T]) hasNext() bool {
	return p.hasNextPage
}

func (p *keyPager[O, R, T]) setNextPageOptions(result R) {
	items := p.resultItemsGetter(result)
	if len(items) == 0 {
		return
	}
	itemsNum := len(items) - 1
	lastItem := items[itemsNum]
	if p.startKeySetter != nil {
		startKey := p.startKeyGetter(lastItem)
		p.startKeySetter(startKey)
	}
	if p.startViewKeySetter != nil {
		startViewKey := p.startViewKeyGetter(lastItem)
		p.startViewKeySetter(startViewKey)
	}
	if p.startKeyDocIDSetter != nil {
		startKeyDocID := p.startKeyDocIDGetter(lastItem)
		p.startKeyDocIDSetter(startKeyDocID)
	}
}

func (p *keyPager[O, R, T]) getOptions() O {
	return p.optionsCloner(p.options)
}

func (p *keyPager[O, R, T]) setOptions(o O) {
	p.options = p.optionsCloner(o)
}

func (p *keyPager[O, R, T]) getLimit() *int64 {
	return p.limitGetter()
}

func (p *keyPager[O, R, T]) setLimit(pageSize int64) {
	p.limitSetter(pageSize + 1)
}
