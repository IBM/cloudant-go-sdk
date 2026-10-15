/**
 * © Copyright IBM Corporation 2026. All Rights Reserved.
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

type IndexInfoPagerOptions interface {
	*cloudantv1.GetIndexInfoOptions | *cloudantv1.PostIndexInfoOptions
}

func NewIndexInfoPagination[O IndexInfoPagerOptions](c *cloudantv1.CloudantV1, o O) Pagination[cloudantv1.IndexInfoResult] {
	return &paginationImplementor[O, cloudantv1.IndexInfoResult]{
		service:  c,
		options:  o,
		newPager: newIndexInfoPager[O],
	}
}

func newIndexInfoPager[O IndexInfoPagerOptions](c *cloudantv1.CloudantV1, o O) (Pager[cloudantv1.IndexInfoResult], error) {
	if err := validatePagerOptions(bookmarkPagerValidationRules, o); err != nil {
		return nil, err
	}

	switch opts := any(o).(type) {
	case *cloudantv1.GetIndexInfoOptions:
		pd := newGetIndexInfoBookmarkPager(c, opts)
		p := newBasePager(pd)

		return p, nil
	case *cloudantv1.PostIndexInfoOptions:
		pd := newPostIndexInfoBookmarkPager(c, opts)
		p := newBasePager(pd)

		return p, nil
	}

	return nil, ErrNotImplemented
}

func withBookmark(r *cloudantv1.IndexInfoResult) *cloudantv1.IndexInfoResult {
	if r != nil && r.Bookmark == nil {
		r.Bookmark = core.StringPtr("")
	}
	return r
}

func indexInfoHasNextPage(r *cloudantv1.IndexInfoResult, _ []cloudantv1.IndexInfoResult, _ *int64) bool {
	return r.Bookmark != nil && *r.Bookmark != ""
}

func newGetIndexInfoBookmarkPager(c *cloudantv1.CloudantV1, o *cloudantv1.GetIndexInfoOptions) *bookmarkPager[*cloudantv1.GetIndexInfoOptions, *cloudantv1.IndexInfoResult, cloudantv1.IndexInfoResult] {
	opts := *o
	return &bookmarkPager[*cloudantv1.GetIndexInfoOptions, *cloudantv1.IndexInfoResult, cloudantv1.IndexInfoResult]{
		service:     c,
		options:     &opts,
		hasNextPage: true,
		requestFunction: func(ctx context.Context, o *cloudantv1.GetIndexInfoOptions) (*cloudantv1.IndexInfoResult, *core.DetailedResponse, error) {
			result, response, err := c.GetIndexInfoWithContext(ctx, o)
			return withBookmark(result), response, err
		},
		resultItemsGetter: func(result *cloudantv1.IndexInfoResult) []cloudantv1.IndexInfoResult {
			return []cloudantv1.IndexInfoResult{*result}
		},
		bookmarkGetter: func(result *cloudantv1.IndexInfoResult) string { return *result.Bookmark },
		bookmarkSetter: opts.SetStartKey,
		optionsCloner: func(o *cloudantv1.GetIndexInfoOptions) *cloudantv1.GetIndexInfoOptions {
			opts := *o
			return &opts
		},
		limitGetter:     func() *int64 { return opts.Limit },
		limitSetter:     opts.SetLimit,
		skipSetter:      opts.SetSkip,
		hasNextFunction: indexInfoHasNextPage,
	}
}

func newPostIndexInfoBookmarkPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostIndexInfoOptions) *bookmarkPager[*cloudantv1.PostIndexInfoOptions, *cloudantv1.IndexInfoResult, cloudantv1.IndexInfoResult] {
	opts := *o
	return &bookmarkPager[*cloudantv1.PostIndexInfoOptions, *cloudantv1.IndexInfoResult, cloudantv1.IndexInfoResult]{
		service:     c,
		options:     &opts,
		hasNextPage: true,
		requestFunction: func(ctx context.Context, o *cloudantv1.PostIndexInfoOptions) (*cloudantv1.IndexInfoResult, *core.DetailedResponse, error) {
			result, response, err := c.PostIndexInfoWithContext(ctx, o)
			return withBookmark(result), response, err
		},
		resultItemsGetter: func(result *cloudantv1.IndexInfoResult) []cloudantv1.IndexInfoResult {
			return []cloudantv1.IndexInfoResult{*result}
		},

		bookmarkGetter: func(result *cloudantv1.IndexInfoResult) string { return *result.Bookmark },
		bookmarkSetter: opts.SetStartKey,
		optionsCloner: func(o *cloudantv1.PostIndexInfoOptions) *cloudantv1.PostIndexInfoOptions {
			opts := *o
			return &opts
		},
		limitGetter:     func() *int64 { return opts.Limit },
		limitSetter:     opts.SetLimit,
		skipSetter:      opts.SetSkip,
		hasNextFunction: indexInfoHasNextPage,
	}
}
