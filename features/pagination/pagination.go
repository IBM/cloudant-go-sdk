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

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

var ErrNotImplemented = errors.New("not yet implemented")

type paginatedRow interface {
	keyPaginatedRow | bookmarkPaginatedRow
}

// basePager is a generic structure designed to handle pagination logic.
// It works with any type T that satisfies the paginatedRow interface.
// Fields:
//   - service: A pointer to a CloudantV1 service client used for database interactions.
//   - hasNext: A boolean indicating if more pages may be available for fetching.
//   - pageSize: The number of items per page, controlling the batch size of each query.
//   - getNextWithContext: A function that retrieves the next set of data based on a given context,
//     returning a slice of type T and an error.
type basePager[T paginatedRow] struct {
	service            *cloudantv1.CloudantV1
	hasNext            bool
	pageSize           int64
	getNextWithContext func(context.Context) ([]T, error)
}

// HasNext returns false if there are no more pages.
func (p *basePager[T]) HasNext() bool {
	return p.hasNext
}

// GetNext retrieves the next page of results.
func (p *basePager[T]) GetNext() ([]T, error) {
	return p.GetNextWithContext(context.Background())
}

// GetNextWithContext retrieves the next page of results with user provided context.
func (p *basePager[T]) GetNextWithContext(ctx context.Context) ([]T, error) {
	return nil, ErrNotImplemented
}

// GetAll retrieves all elements from the pager.
func (p *basePager[T]) GetAll() ([]T, error) {
	return p.GetAllWithContext(context.Background())
}

// GetAllWithContext retrieves all the elements from the pager with user provided context.
func (p *basePager[T]) GetAllWithContext(ctx context.Context) ([]T, error) {
	return nil, ErrNotImplemented
}

// setGetNextWithContext sets the concrete implementation of getNextWithContext function for the pager.
func (p *basePager[T]) setGetNextWithContext(fn func(context.Context) ([]T, error)) {
	p.getNextWithContext = fn
}

// newBasePager creates a new base pager for database operations.
func newBasePager[T paginatedRow](c *cloudantv1.CloudantV1) *basePager[T] {
	return new(basePager[T])
}

type keyPaginatedRow interface {
	cloudantv1.DocsResultRow | cloudantv1.ViewResultRow
}

type keyPager[T keyPaginatedRow] struct {
	*basePager[T]
	lastRow T
}

func newKeyPager[T keyPaginatedRow](c *cloudantv1.CloudantV1) *keyPager[T] {
	return new(keyPager[T])
}

func (p *keyPager[T]) itemsGetter(rows []T) []T {
	return rows
}

// AllDocsPagerOptions defines options for paginating through all documents or partitioned documents in a Cloudant database.
type AllDocsPagerOptions interface {
	*cloudantv1.PostAllDocsOptions | *cloudantv1.PostPartitionAllDocsOptions
}

// AllDocsPager is an interface for pagination of database all-docs operations.
type AllDocsPager interface {
	// HasNext returns false if there are no more pages.
	HasNext() bool

	// GetNext retrieves the next page of results.
	GetNext() ([]cloudantv1.DocsResultRow, error)

	// GetNextWithContext retrieves the next page of results with user provided context.
	GetNextWithContext(context.Context) ([]cloudantv1.DocsResultRow, error)

	// GetAll retrieves all elements from the pager.
	GetAll() ([]cloudantv1.DocsResultRow, error)

	// GetAllWithContext retrieves all the elements from the pager with user provided context.
	GetAllWithContext(context.Context) ([]cloudantv1.DocsResultRow, error)
}

// NewAllDocsPager creates a new pager for all-docs operations.
func NewAllDocsPager[O AllDocsPagerOptions](c *cloudantv1.CloudantV1, o O) (AllDocsPager, error) {
	// generics initialiation example
	switch opts := any(o).(type) {
	case *cloudantv1.PostAllDocsOptions:
		optsClone := *opts
		p := newAllDocsBasePager(c, &optsClone)
		return p, ErrNotImplemented
	case *cloudantv1.PostPartitionAllDocsOptions:
		return nil, ErrNotImplemented
	}
	return nil, ErrNotImplemented
}

type allDocsBasePager[O AllDocsPagerOptions] struct {
	*keyPager[cloudantv1.DocsResultRow]
	options             O
	nextRequestFunction func(ctx context.Context, o O) (*cloudantv1.AllDocsResult, *core.DetailedResponse, error)
	nextKeySetter       func(string) O
}

func newAllDocsBasePager[O AllDocsPagerOptions](c *cloudantv1.CloudantV1, o O) *allDocsBasePager[O] {
	return new(allDocsBasePager[O])
}

func (p *allDocsBasePager[O]) getNextWithContext(ctx context.Context) ([]cloudantv1.DocsResultRow, error) {
	return nil, ErrNotImplemented
}

// ViewPagerOptions defines options for paginating through views or partitioned views in a Cloudant database.
type ViewPagerOptions interface {
	*cloudantv1.PostViewOptions | *cloudantv1.PostPartitionViewOptions
}

// ViewPager is an interface for pagination of database views operations.
type ViewPager interface {
	// HasNext returns false if there are no more pages.
	HasNext() bool

	// GetNext retrieves the next page of results.
	GetNext() ([]cloudantv1.ViewResultRow, error)

	// GetNextWithContext retrieves the next page of results with user provided context.
	GetNextWithContext(context.Context) ([]cloudantv1.ViewResultRow, error)

	// GetAll retrieves all elements from the pager.
	GetAll() ([]cloudantv1.ViewResultRow, error)

	// GetAllWithContext retrieves all the elements from the pager with user provided context.
	GetAllWithContext(context.Context) ([]cloudantv1.ViewResultRow, error)
}

// NewViewPager creates a new pager for view operations.
func NewViewPager[O ViewPagerOptions](c *cloudantv1.CloudantV1, o O) (ViewPager, error) {
	return nil, ErrNotImplemented
}

type viewBasePager[O ViewPagerOptions] struct {
	*keyPager[cloudantv1.ViewResultRow]
	options             O
	nextRequestFunction func(ctx context.Context, o O) (*cloudantv1.ViewResult, *core.DetailedResponse, error)
	nextKeySetter       func(any) O
	nextKeyIDSetter     func(string) O
}

func newViewBasePager[O ViewPagerOptions](c *cloudantv1.CloudantV1, o O) *viewBasePager[O] {
	return new(viewBasePager[O])
}

func (p *viewBasePager[O]) getNextWithContext(ctx context.Context) ([]cloudantv1.ViewResultRow, error) {
	return nil, ErrNotImplemented
}

type bookmarkPaginatedRow interface {
	cloudantv1.Document | cloudantv1.SearchResultRow
}

type bookmarkPager[T paginatedRow] struct {
	*basePager[T]
}

func newBookmarkPager[T bookmarkPaginatedRow](c *cloudantv1.CloudantV1) *bookmarkPager[T] {
	return new(bookmarkPager[T])
}

func (p *bookmarkPager[T]) itemsGetter(docs []T) []T {
	return docs
}

// FindPagerOptions defines options for paginating through queries or partitioned queries in a Cloudant database.
type FindPagerOptions interface {
	*cloudantv1.PostFindOptions | *cloudantv1.PostPartitionFindOptions
}

// FindPager is an interface for pagination of database query operations.
type FindPager interface {
	// HasNext returns false if there are no more pages.
	HasNext() bool

	// GetNext retrieves the next page of results.
	GetNext() ([]cloudantv1.Document, error)

	// GetNextWithContext retrieves the next page of results with user provided context.
	GetNextWithContext(context.Context) ([]cloudantv1.Document, error)

	// GetAll retrieves all elements from the pager.
	GetAll() ([]cloudantv1.Document, error)

	// GetAllWithContext retrieves all the elements from the pager with user provided context.
	GetAllWithContext(context.Context) ([]cloudantv1.Document, error)
}

// NewFindPager creates a new pager for query operations.
func NewFindPager[O FindPagerOptions](c *cloudantv1.CloudantV1, o O) (FindPager, error) {
	return nil, ErrNotImplemented
}

type findBasePager[O FindPagerOptions] struct {
	*bookmarkPager[cloudantv1.Document]
	options             O
	nextRequestFunction func(ctx context.Context, o O) (*cloudantv1.FindResult, *core.DetailedResponse, error)
	bookmarkSetter      func(string) O
}

func newFindBasePager[O FindPagerOptions](c *cloudantv1.CloudantV1, o O) *findBasePager[O] {
	return new(findBasePager[O])
}

func (p *findBasePager[O]) getNextWithContext(ctx context.Context) ([]cloudantv1.Document, error) {
	return nil, ErrNotImplemented
}

// SearchPagerOptions defines options for paginating through searches or partitioned searches in a Cloudant database.
type SearchPagerOptions interface {
	*cloudantv1.PostSearchOptions | *cloudantv1.PostPartitionSearchOptions
}

// SearchPager is an interface for pagination of database search operations.
type SearchPager interface {
	// HasNext returns false if there are no more pages.
	HasNext() bool

	// GetNext retrieves the next page of results.
	GetNext() ([]cloudantv1.SearchResultRow, error)

	// GetNextWithContext retrieves the next page of results with user provided context.
	GetNextWithContext(context.Context) ([]cloudantv1.SearchResultRow, error)

	// GetAll retrieves all elements from the pager.
	GetAll() ([]cloudantv1.SearchResultRow, error)

	// GetAllWithContext retrieves all the elements from the pager with user provided context.
	GetAllWithContext(context.Context) ([]cloudantv1.SearchResultRow, error)
}

// NewSearchPager creates a new pager for search operations.
func NewSearchPager[O SearchPagerOptions](c *cloudantv1.CloudantV1, o O) (SearchPager, error) {
	return nil, ErrNotImplemented
}

type searchBasePager[O SearchPagerOptions] struct {
	*bookmarkPager[cloudantv1.Document]
	options             O
	nextRequestFunction func(ctx context.Context, o O) (*cloudantv1.SearchResult, *core.DetailedResponse, error)
	bookmarkSetter      func(string) O
}

func newSearchBasePager[O SearchPagerOptions](c *cloudantv1.CloudantV1, o O) *searchBasePager[O] {
	return new(searchBasePager[O])
}

func (p *searchBasePager[O]) getNextWithContext(ctx context.Context) ([]cloudantv1.SearchResultRow, error) {
	return nil, ErrNotImplemented
}
