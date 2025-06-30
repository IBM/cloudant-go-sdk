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
)

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

func newSearchPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostSearchOptions) *bookmarkPager[*cloudantv1.PostSearchOptions, *cloudantv1.SearchResult, cloudantv1.SearchResultRow] {
	return new(bookmarkPager[*cloudantv1.PostSearchOptions, *cloudantv1.SearchResult, cloudantv1.SearchResultRow])
}

func newSearchPartitionPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostPartitionSearchOptions) *bookmarkPager[*cloudantv1.PostPartitionSearchOptions, *cloudantv1.SearchResult, cloudantv1.SearchResultRow] {
	return new(bookmarkPager[*cloudantv1.PostPartitionSearchOptions, *cloudantv1.SearchResult, cloudantv1.SearchResultRow])
}
