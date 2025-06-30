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
	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

// SearchPagerOptions defines options for paginating through searches or partitioned searches in a Cloudant database.
type SearchPagerOptions interface {
	*cloudantv1.PostSearchOptions | *cloudantv1.PostPartitionSearchOptions
}

// NewSearchPagination creates a new pagination for searches operations.
func NewSearchPagination[O SearchPagerOptions](c *cloudantv1.CloudantV1, o O) Pagination[cloudantv1.SearchResultRow] {
	return &paginationImplementor[O, cloudantv1.SearchResultRow]{
		service:  c,
		options:  o,
		newPager: NewSearchPager[O],
	}
}

// NewSearchPager creates a new pager for searches operations.
func NewSearchPager[O SearchPagerOptions](c *cloudantv1.CloudantV1, o O) (Pager[cloudantv1.SearchResultRow], error) {
	switch opts := any(o).(type) {
	case *cloudantv1.PostSearchOptions:
		if err := validatePagerOptions(searchPagerValidationRules, opts); err != nil {
			return nil, err
		}

		pd := newSearchPager(c, opts)
		p := newBasePager(pd)

		return p, nil
	case *cloudantv1.PostPartitionSearchOptions:
		if err := validatePagerOptions(bookmarkPagerValidationRules, opts); err != nil {
			return nil, err
		}

		pd := newSearchPartitionPager(c, opts)
		p := newBasePager(pd)

		return p, nil
	}
	return nil, ErrNotImplemented
}

func newSearchPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostSearchOptions) *bookmarkPager[*cloudantv1.PostSearchOptions, *cloudantv1.SearchResult, cloudantv1.SearchResultRow] {
	opts := *o
	return &bookmarkPager[*cloudantv1.PostSearchOptions, *cloudantv1.SearchResult, cloudantv1.SearchResultRow]{
		service:           c,
		options:           &opts,
		hasNextPage:       true,
		requestFunction:   c.PostSearchWithContext,
		resultItemsGetter: func(result *cloudantv1.SearchResult) []cloudantv1.SearchResultRow { return result.Rows },
		bookmarkGetter:    func(result *cloudantv1.SearchResult) string { return *result.Bookmark },
		bookmarkSetter:    opts.SetBookmark,
		optionsCloner: func(o *cloudantv1.PostSearchOptions) *cloudantv1.PostSearchOptions {
			opts := *o
			return &opts
		},
		limitGetter: func() *int64 { return opts.Limit },
		limitSetter: opts.SetLimit,
	}
}

func newSearchPartitionPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostPartitionSearchOptions) *bookmarkPager[*cloudantv1.PostPartitionSearchOptions, *cloudantv1.SearchResult, cloudantv1.SearchResultRow] {
	opts := *o
	return &bookmarkPager[*cloudantv1.PostPartitionSearchOptions, *cloudantv1.SearchResult, cloudantv1.SearchResultRow]{
		service:           c,
		options:           &opts,
		hasNextPage:       true,
		requestFunction:   c.PostPartitionSearchWithContext,
		resultItemsGetter: func(result *cloudantv1.SearchResult) []cloudantv1.SearchResultRow { return result.Rows },
		bookmarkGetter:    func(result *cloudantv1.SearchResult) string { return *result.Bookmark },
		bookmarkSetter:    opts.SetBookmark,
		optionsCloner: func(o *cloudantv1.PostPartitionSearchOptions) *cloudantv1.PostPartitionSearchOptions {
			opts := *o
			return &opts
		},
		limitGetter: func() *int64 { return opts.Limit },
		limitSetter: opts.SetLimit,
	}
}
