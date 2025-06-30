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
	requestFunction     func(context.Context, O) (R, *core.DetailedResponse, error)
	resultItemsGetter   func(R) []T
	startKeyGetter      func(T) string
	startKeySetter      func(string) O
	startViewKeyGetter  func(T) any
	startViewKeySetter  func(any) O
	startKeyDocIDGetter func(T) string
	startKeyDocIDSetter func(string) O
	limitGetter         func() *int64
	limitSetter         func(int64) O
}

func (p *keyPager[O, R, T]) nextRequestFunction(ctx context.Context) (R, error) {
	return nil, ErrNotImplemented
}

func (p *keyPager[O, R, T]) itemsGetter(result R) []T {
	return make([]T, 0)
}

func (p *keyPager[O, R, T]) setNextPageOptions(result R) {
}

func (p *keyPager[O, R, T]) getLimit() *int64 {
	return nil
}

func (p *keyPager[O, R, T]) setLimit(pageSize int64) {
}
