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

// AllDocsPagerOptions defines options for paginating through all documents or partitioned all documents in a Cloudant database.
type AllDocsPagerOptions interface {
	*cloudantv1.PostAllDocsOptions | *cloudantv1.PostPartitionAllDocsOptions
}

// NewAllDocsPagination creates a new pagination for all documents operations.
func NewAllDocsPagination[O AllDocsPagerOptions](c *cloudantv1.CloudantV1, o O) Pagination[cloudantv1.DocsResultRow] {
	return &paginationImplementor[O, cloudantv1.DocsResultRow]{
		service:  c,
		options:  o,
		newPager: newAllDocsPager[O],
	}
}

// newAllDocsPager creates a new pager for all documents operations.
func newAllDocsPager[O AllDocsPagerOptions](c *cloudantv1.CloudantV1, o O) (Pager[cloudantv1.DocsResultRow], error) {
	if err := validatePagerOptions(keyPagerValidationRules, o); err != nil {
		return nil, err
	}

	switch opts := any(o).(type) {
	case *cloudantv1.PostAllDocsOptions:
		pd := newAllDocsKeyPager(c, opts)
		p := newBasePager(pd)

		return p, nil
	case *cloudantv1.PostPartitionAllDocsOptions:
		pd := newAllDocsPartitionKeyPager(c, opts)
		p := newBasePager(pd)

		return p, nil
	}

	return nil, ErrNotImplemented
}

func newAllDocsKeyPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostAllDocsOptions) *keyPager[*cloudantv1.PostAllDocsOptions, *cloudantv1.AllDocsResult, cloudantv1.DocsResultRow] {
	opts := *o
	return &keyPager[*cloudantv1.PostAllDocsOptions, *cloudantv1.AllDocsResult, cloudantv1.DocsResultRow]{
		service:           c,
		options:           &opts,
		hasNextPage:       true,
		requestFunction:   c.PostAllDocsWithContext,
		resultItemsGetter: func(result *cloudantv1.AllDocsResult) []cloudantv1.DocsResultRow { return result.Rows },
		startKeyGetter: func(item cloudantv1.DocsResultRow) string {
			if item.Key != nil {
				return *item.Key
			}
			return ""
		},
		startKeySetter: opts.SetStartKey,
		optionsCloner: func(o *cloudantv1.PostAllDocsOptions) *cloudantv1.PostAllDocsOptions {
			opts := *o
			return &opts
		},
		limitGetter: func() *int64 { return opts.Limit },
		limitSetter: opts.SetLimit,
	}
}

func newAllDocsPartitionKeyPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostPartitionAllDocsOptions) *keyPager[*cloudantv1.PostPartitionAllDocsOptions, *cloudantv1.AllDocsResult, cloudantv1.DocsResultRow] {
	opts := *o
	return &keyPager[*cloudantv1.PostPartitionAllDocsOptions, *cloudantv1.AllDocsResult, cloudantv1.DocsResultRow]{
		service:           c,
		options:           &opts,
		hasNextPage:       true,
		requestFunction:   c.PostPartitionAllDocsWithContext,
		resultItemsGetter: func(result *cloudantv1.AllDocsResult) []cloudantv1.DocsResultRow { return result.Rows },
		startKeyGetter: func(item cloudantv1.DocsResultRow) string {
			if item.Key != nil {
				return *item.Key
			}
			return ""
		},
		startKeySetter: opts.SetStartKey,
		optionsCloner: func(o *cloudantv1.PostPartitionAllDocsOptions) *cloudantv1.PostPartitionAllDocsOptions {
			opts := *o
			return &opts
		},
		limitGetter: func() *int64 { return opts.Limit },
		limitSetter: opts.SetLimit,
	}
}
