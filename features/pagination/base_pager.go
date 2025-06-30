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
	"errors"
	"iter"
)

var ErrNotImplemented = errors.New("not yet implemented")
var ErrNoMoreResults = errors.New("no more results available")

type basePagerOptions interface {
	keyPagerOptions | bookmarkPagerOptions
}

type requestResult interface {
	keyRequestResult | bookmarkRequestResult
}

type paginatedRow interface {
	keyPaginatedRow | bookmarkPaginatedRow
}

// pager is an interface implementing callbacks necessary for Pager interface
type pager[O basePagerOptions, R requestResult, T paginatedRow] interface {
	nextRequestFunction(context.Context) (R, error)
	itemsGetter(R) []T
	getOptions() O
	setOptions(O)
	setNextPageOptions(R)
	getLimit() *int64
	setLimit(int64)
}

// basePager is a generic structure designed to handle pagination logic.
// It works with any type T that satisfies the paginatedRow interface.
// Fields:
//   - pager: An implementation of `pager` iterface.
//   - hasNext: A boolean indicating if there are no more pages available for fetching.
//   - pageSize: The number of items per page, controlling the batch size of each query.
type basePager[O basePagerOptions, R requestResult, T paginatedRow] struct {
	pager    pager[O, R, T]
	options  O
	err      error
	hasNext  bool
	pageSize int64
}

// newBasePager creates a new base pager for database operations.
func newBasePager[O basePagerOptions, R requestResult, T paginatedRow](pd pager[O, R, T]) *basePager[O, R, T] {
	pageSize := getPageSizeFromOptionsLimit(pd)
	return &basePager[O, R, T]{
		pager:    pd,
		options:  pd.getOptions(),
		hasNext:  true,
		pageSize: pageSize,
	}
}

// HasNext returns false if there are no more pages.
func (p *basePager[O, R, T]) HasNext() bool {
	return p.hasNext
}

// GetNext retrieves the next page of results.
func (p *basePager[O, R, T]) GetNext() ([]T, error) {
	return p.GetNextWithContext(context.Background())
}

// GetNextWithContext retrieves the next page of results with user provided context.
func (p *basePager[O, R, T]) GetNextWithContext(ctx context.Context) ([]T, error) {
	if p.err != nil {
		return nil, p.err
	} else if !p.hasNext {
		return nil, ErrNoMoreResults
	}

	p.pager.setLimit(p.pageSize)
	result, err := p.pager.nextRequestFunction(ctx)
	if err != nil {
		return nil, err
	}

	items := p.pager.itemsGetter(result)

	if len(items) < int(p.pageSize) {
		p.hasNext = false
	} else {
		p.pager.setNextPageOptions(result)
	}
	return items, nil
}

// GetAll retrieves all elements from the pager.
func (p *basePager[O, R, T]) GetAll() ([]T, error) {
	return p.GetAllWithContext(context.Background())
}

// GetAllWithContext retrieves all the elements from the pager with user provided context.
func (p *basePager[O, R, T]) GetAllWithContext(ctx context.Context) ([]T, error) {
	acc := make([]T, 0)
	for item, err := range p.RowsWithContext(ctx) {
		if err != nil {
			p.pager.setOptions(p.options)
			return nil, err
		}
		acc = append(acc, item)
	}
	return acc, nil
}

// Pages returns an iterator for all pages from the pager.
func (p *basePager[O, R, T]) Pages() iter.Seq2[[]T, error] {
	return p.PagesWithContext(context.Background())
}

// PagesWithContext returns an iterator for all pages from the pager queried with user provided context.
func (p *basePager[O, R, T]) PagesWithContext(ctx context.Context) iter.Seq2[[]T, error] {
	return func(yield func([]T, error) bool) {
		for p.HasNext() {
			rows, err := p.GetNextWithContext(ctx)
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

// Rows returns an iterator for all elements from the pager.
func (p *basePager[O, R, T]) Rows() iter.Seq2[T, error] {
	return p.RowsWithContext(context.Background())
}

// RowsWithContext returns an iterator for all elements from the pager queried with user provided context.
func (p *basePager[O, R, T]) RowsWithContext(ctx context.Context) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for rows, err := range p.PagesWithContext(ctx) {
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

// getPageSizeFromOptionsLimit infers pageSize from options limit or defaults to 20.
func getPageSizeFromOptionsLimit[O basePagerOptions, R requestResult, T paginatedRow](pd pager[O, R, T]) int64 {
	pageSize := int64(20)
	if pd.getLimit() != nil {
		pageSize = *pd.getLimit()
	}
	return pageSize
}
