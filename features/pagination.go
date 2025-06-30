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
	"iter"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

type Pagination[T PaginatedRow] interface {
	// Pager returns new Pager.
	Pager() (Pager[T], error)

	// Pages returns an iterator for all pages from the pager.
	Pages() iter.Seq2[[]T, error]

	// PagesWithContext returns an iterator for all pages from the pager queried with user provided context.
	PagesWithContext(context.Context) iter.Seq2[[]T, error]

	// Rows returns an iterator for all elements from the pager.
	Rows() iter.Seq2[T, error]

	// RowsWithContext returns an iterator for all elements from the pager queried with user provided context.
	RowsWithContext(context.Context) iter.Seq2[T, error]
}

type paginationImplementor[O pagerOptions, T paginatedRow] struct {
	service  *cloudantv1.CloudantV1
	options  O
	newPager func(*cloudantv1.CloudantV1, O) (Pager[T], error)
}

func (pi *paginationImplementor[O, T]) Pager() (Pager[T], error) {
	return pi.newPager(pi.service, pi.options)
}

func (pi *paginationImplementor[O, T]) Pages() iter.Seq2[[]T, error] {
	return pi.PagesWithContext(context.Background())
}

func (pi *paginationImplementor[O, T]) PagesWithContext(ctx context.Context) iter.Seq2[[]T, error] {
	pager, err := pi.Pager()
	if err != nil {
		return func(yield func([]T, error) bool) {
			yield(nil, err)
		}
	}
	return pagesWithContext(ctx, pager)
}

func (pi *paginationImplementor[O, T]) Rows() iter.Seq2[T, error] {
	return pi.RowsWithContext(context.Background())
}

func (pi *paginationImplementor[O, T]) RowsWithContext(ctx context.Context) iter.Seq2[T, error] {
	pager, err := pi.Pager()
	if err != nil {
		return func(yield func(T, error) bool) {
			yield(*new(T), err)
		}
	}
	return rowsWithContext(ctx, pager)
}

// pagesWithContext returns an iterator for all pages from the pager queried with user provided context.
func pagesWithContext[T PaginatedRow](ctx context.Context, pd Pager[T]) iter.Seq2[[]T, error] {
	return func(yield func([]T, error) bool) {
		for pd.HasNext() {
			rows, err := pd.GetNextWithContext(ctx)
			if err != nil {
				yield(nil, err)
				return
			}
			if !yield(rows, nil) {
				return
			}
		}
	}
}

// rowsWithContext returns an iterator for all elements from the pager queried with user provided context.
func rowsWithContext[T PaginatedRow](ctx context.Context, pd Pager[T]) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for rows, err := range pagesWithContext(ctx, pd) {
			if err != nil {
				yield(*new(T), err)
				return
			}
			for _, row := range rows {
				if !yield(row, nil) {
					return
				}
			}
		}
	}
}
