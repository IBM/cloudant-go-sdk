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

// ViewPagerOptions defines options for paginating through views or partitioned views in a Cloudant database.
type ViewPagerOptions interface {
	*cloudantv1.PostViewOptions | *cloudantv1.PostPartitionViewOptions
}

// NewViewPagination creates a new pagination for views operations.
func NewViewPagination[O ViewPagerOptions](c *cloudantv1.CloudantV1, o O) Pagination[cloudantv1.ViewResultRow] {
	return &paginationImplementor[O, cloudantv1.ViewResultRow]{
		service:  c,
		options:  o,
		newPager: newViewPager[O],
	}
}

// NewViewPager creates a new pager for views operations.
func newViewPager[O ViewPagerOptions](c *cloudantv1.CloudantV1, o O) (Pager[cloudantv1.ViewResultRow], error) {
	if err := validatePagerOptions(keyPagerValidationRules, o); err != nil {
		return nil, err
	}

	switch opts := any(o).(type) {
	case *cloudantv1.PostViewOptions:
		pd := newViewKeyPager(c, opts)
		p := newBasePager(pd)

		return p, nil
	case *cloudantv1.PostPartitionViewOptions:
		pd := newViewPartitionKeyPager(c, opts)
		p := newBasePager(pd)

		return p, nil
	}

	return nil, ErrNotImplemented
}

func newViewKeyPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostViewOptions) *keyPager[*cloudantv1.PostViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow] {
	opts := *o
	return &keyPager[*cloudantv1.PostViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow]{
		service:            c,
		options:            &opts,
		hasNextPage:        true,
		requestFunction:    c.PostViewWithContext,
		resultItemsGetter:  func(result *cloudantv1.ViewResult) []cloudantv1.ViewResultRow { return result.Rows },
		startViewKeyGetter: func(item cloudantv1.ViewResultRow) any { return item.Key },
		startViewKeySetter: opts.SetStartKey,
		startKeyDocIDGetter: func(item cloudantv1.ViewResultRow) string {
			if item.ID != nil {
				return *item.ID
			}
			return ""
		},
		startKeyDocIDSetter: opts.SetStartKeyDocID,
		optionsCloner: func(o *cloudantv1.PostViewOptions) *cloudantv1.PostViewOptions {
			opts := *o
			return &opts
		},
		limitGetter: func() *int64 { return opts.Limit },
		limitSetter: opts.SetLimit,
	}
}

func newViewPartitionKeyPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostPartitionViewOptions) *keyPager[*cloudantv1.PostPartitionViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow] {
	opts := *o
	return &keyPager[*cloudantv1.PostPartitionViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow]{
		service:            c,
		options:            &opts,
		hasNextPage:        true,
		requestFunction:    c.PostPartitionViewWithContext,
		resultItemsGetter:  func(result *cloudantv1.ViewResult) []cloudantv1.ViewResultRow { return result.Rows },
		startViewKeyGetter: func(item cloudantv1.ViewResultRow) any { return item.Key },
		startViewKeySetter: opts.SetStartKey,
		startKeyDocIDGetter: func(item cloudantv1.ViewResultRow) string {
			if item.ID != nil {
				return *item.ID
			}
			return ""
		},
		startKeyDocIDSetter: opts.SetStartKeyDocID,
		optionsCloner: func(o *cloudantv1.PostPartitionViewOptions) *cloudantv1.PostPartitionViewOptions {
			opts := *o
			return &opts
		},
		limitGetter: func() *int64 { return opts.Limit },
		limitSetter: opts.SetLimit,
	}
}
