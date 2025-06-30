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

// FindPagerOptions defines options for paginating through queries or partitioned queries in a Cloudant database.
type FindPagerOptions interface {
	*cloudantv1.PostFindOptions | *cloudantv1.PostPartitionFindOptions
}

// NewFindPagination creates a new pagination for queries operations.
func NewFindPagination[O FindPagerOptions](c *cloudantv1.CloudantV1, o O) Pagination[cloudantv1.Document] {
	return &paginationImplementor[O, cloudantv1.Document]{
		service:  c,
		options:  o,
		newPager: NewFindPager[O],
	}
}

// NewFindPager creates a new pager for queries operations.
func NewFindPager[O FindPagerOptions](c *cloudantv1.CloudantV1, o O) (Pager[cloudantv1.Document], error) {
	if err := validatePagerOptions(keyPagerValidationRules, o); err != nil {
		return nil, err
	}

	switch opts := any(o).(type) {
	case *cloudantv1.PostFindOptions:
		pd := newFindPager(c, opts)
		p := newBasePager(pd)

		return p, nil
	case *cloudantv1.PostPartitionFindOptions:
		pd := newFindPartitionPager(c, opts)
		p := newBasePager(pd)

		return p, nil
	}

	return nil, ErrNotImplemented
}

func newFindPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostFindOptions) *bookmarkPager[*cloudantv1.PostFindOptions, *cloudantv1.FindResult, cloudantv1.Document] {
	opts := *o
	return &bookmarkPager[*cloudantv1.PostFindOptions, *cloudantv1.FindResult, cloudantv1.Document]{
		service:           c,
		options:           &opts,
		hasNextPage:       true,
		requestFunction:   c.PostFindWithContext,
		resultItemsGetter: func(result *cloudantv1.FindResult) []cloudantv1.Document { return result.Docs },
		bookmarkGetter:    func(result *cloudantv1.FindResult) string { return *result.Bookmark },
		bookmarkSetter:    opts.SetBookmark,
		optionsCloner: func(o *cloudantv1.PostFindOptions) *cloudantv1.PostFindOptions {
			opts := *o
			return &opts
		},
		limitGetter: func() *int64 { return opts.Limit },
		limitSetter: opts.SetLimit,
	}
}

func newFindPartitionPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostPartitionFindOptions) *bookmarkPager[*cloudantv1.PostPartitionFindOptions, *cloudantv1.FindResult, cloudantv1.Document] {
	opts := *o
	return &bookmarkPager[*cloudantv1.PostPartitionFindOptions, *cloudantv1.FindResult, cloudantv1.Document]{
		service:           c,
		options:           &opts,
		hasNextPage:       true,
		requestFunction:   c.PostPartitionFindWithContext,
		resultItemsGetter: func(result *cloudantv1.FindResult) []cloudantv1.Document { return result.Docs },
		bookmarkGetter:    func(result *cloudantv1.FindResult) string { return *result.Bookmark },
		bookmarkSetter:    opts.SetBookmark,
		optionsCloner: func(o *cloudantv1.PostPartitionFindOptions) *cloudantv1.PostPartitionFindOptions {
			opts := *o
			return &opts
		},
		limitGetter: func() *int64 { return opts.Limit },
		limitSetter: opts.SetLimit,
	}
}
