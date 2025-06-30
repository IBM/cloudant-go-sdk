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
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	minLimit = 1
	maxLimit = 200
)

var limitValidationRule = fmt.Sprintf("omitempty,min=%d,max=%d", minLimit, maxLimit)

var ErrNotImplemented = errors.New("not yet implemented")
var ErrNoMoreResults = errors.New("no more results available")

type pagerOptions interface {
	keyPagerOptions | bookmarkPagerOptions
}

type requestResult interface {
	keyRequestResult | bookmarkRequestResult
}

type paginatedRow interface {
	keyPaginatedRow | bookmarkPaginatedRow
}

// Pager is an interface for pagination of Cloudant query operations.
type Pager[T paginatedRow] interface {
	// HasNext returns false if there are no more pages.
	HasNext() bool

	// GetNext retrieves the next page of results.
	GetNext() ([]T, error)

	// GetNextWithContext retrieves the next page of results with user provided context.
	GetNextWithContext(context.Context) ([]T, error)

	// GetAll retrieves all elements from the pager.
	GetAll() ([]T, error)

	// GetAllWithContext retrieves all the elements from the pager with user provided context.
	GetAllWithContext(context.Context) ([]T, error)
}

// pagerImplementor is an internal interface of callbacks necessary for Pager interface implementation
type pagerImplementor[O pagerOptions, R requestResult, T paginatedRow] interface {
	nextRequestFunction(context.Context) (R, error)
	itemsGetter(R) ([]T, error)
	hasNext() bool
	getOptions() O
	setOptions(O)
	setNextPageOptions(R)
	getLimit() *int64
	setLimit(int64)
}

// basePager is a generic implementation of Pager interface.
// Fields:
//   - pager: An implementation of `paginated` iterface.
//   - options: An option structure of one of pagerOptions union types.
//   - err: Holds any errors encountered while fetching the next page of items.
//   - pageSize: The number of items per page, controlling the batch size of each query.
type basePager[O pagerOptions, R requestResult, T paginatedRow] struct {
	pager    pagerImplementor[O, R, T]
	options  O
	err      error
	pageSize int64
}

// newBasePager creates a new base pager for database operations.
func newBasePager[O pagerOptions, R requestResult, T paginatedRow](pd pagerImplementor[O, R, T]) Pager[T] {
	pageSize := getPageSizeFromOptionsLimit(pd)
	return &basePager[O, R, T]{
		pager:    pd,
		options:  pd.getOptions(),
		pageSize: pageSize,
	}
}

// HasNext returns false if there are no more pages.
func (p *basePager[O, R, T]) HasNext() bool {
	return p.pager.hasNext()
}

// GetNext retrieves the next page of results.
func (p *basePager[O, R, T]) GetNext() ([]T, error) {
	return p.GetNextWithContext(context.Background())
}

// GetNextWithContext retrieves the next page of results with user provided context.
func (p *basePager[O, R, T]) GetNextWithContext(ctx context.Context) ([]T, error) {
	if p.err != nil {
		return nil, p.err
	} else if !p.pager.hasNext() {
		return nil, ErrNoMoreResults
	}

	p.pager.setLimit(p.pageSize)
	result, err := p.pager.nextRequestFunction(ctx)
	if err != nil {
		return nil, err
	}
	err = validatePagerResponse(result)
	if err != nil {
		return nil, err
	}

	items, err := p.pager.itemsGetter(result)
	if err != nil {
		p.err = err
	}

	if p.pager.hasNext() {
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
	for p.HasNext() {
		items, err := p.GetNextWithContext(ctx)
		if err != nil {
			p.pager.setOptions(p.options)
			return nil, err
		}
		acc = append(acc, items...)
	}
	return acc, nil
}

// getPageSizeFromOptionsLimit infers pageSize from options limit or defaults to 200.
func getPageSizeFromOptionsLimit[O pagerOptions, R requestResult, T paginatedRow](pd pagerImplementor[O, R, T]) int64 {
	pageSize := int64(maxLimit)
	if pd.getLimit() != nil {
		pageSize = *pd.getLimit()
	}
	return pageSize
}

// validatePagerOptions validates the options struct using the provided rules.
func validatePagerOptions[O pagerOptions](rules map[string]string, options O) error {
	validate := validator.New()
	validate.RegisterStructValidationMapRules(rules, options)
	err := validate.Struct(options)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			switch e.Tag() {
			case "min":
				return fmt.Errorf("the provided limit %d is lower than the minimum page size value of %s", e.Value(), e.Param())
			case "max":
				return fmt.Errorf("the provided limit %d exceeds the maximum page size value of %s", e.Value(), e.Param())
			// This validates that the value is the default value and is almost the opposite of required.
			// i.e. it returns an error if the value is not the default.
			case "isdefault":
				return fmt.Errorf("the option %q is invalid when using pagination", e.Field())
			}
		}
	}
	return err
}

// validatePagerResponse validates the options struct using the provided rules.
func validatePagerResponse[R requestResult](response R) error {
	validate := validator.New()
	err := validate.Struct(response)
	return err
}
