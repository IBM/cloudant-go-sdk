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
	"errors"
	"fmt"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

// IndexInfoPagerOptions defines options for paginating through index information in a Cloudant database.
//
// Beta: This feature is currently in beta and its interface and behavior are subject to change.
type IndexInfoPagerOptions interface {
	*cloudantv1.PostIndexInfoOptions
}

// NewIndexInfoPagination creates a new pagination for index information operations.
//
// Because index information contains multiple index types, each page contains a single
// IndexInfoResult with the SearchIndexes and ViewIndexes for that page.
//
// Beta: This feature is currently in beta and its interface and behavior are subject to change.
func NewIndexInfoPagination[O IndexInfoPagerOptions](c *cloudantv1.CloudantV1, o O) Pagination[cloudantv1.IndexInfoResult] {
	return &paginationImplementor[O, cloudantv1.IndexInfoResult]{
		service:  c,
		options:  o,
		newPager: newPostIndexInfoPager[O],
	}
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

func newPostIndexInfoPager[O IndexInfoPagerOptions](c *cloudantv1.CloudantV1, o O) (Pager[cloudantv1.IndexInfoResult], error) {
	if err := validatePagerOptions(indexInfoPagerValidationRules, o); err != nil {
		if errors.Is(err, ErrKeySet) {
			err = fmt.Errorf(`%w. No need to paginate as "Key" returns a single result for an ID`, err)
		}
		return nil, err
	}

	opts := *o
	pd := &bookmarkPager[*cloudantv1.PostIndexInfoOptions, *cloudantv1.IndexInfoResult, cloudantv1.IndexInfoResult]{
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
	return newBasePager(pd), nil
}
