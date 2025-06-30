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

// DesignDocsPagerOptions defines options for paginating through design documents in a Cloudant database.
type DesignDocsPagerOptions interface {
	*cloudantv1.PostDesignDocsOptions
}

// NewDesignDocsPagination creates a new pagination for design documents operations.
func NewDesignDocsPagination[O DesignDocsPagerOptions](c *cloudantv1.CloudantV1, o O) Pagination[cloudantv1.DocsResultRow] {
	return &paginationImplementor[O, cloudantv1.DocsResultRow]{
		service:  c,
		options:  o,
		newPager: NewDesignDocsPager[O],
	}
}

// NewDesignDocsPager creates a new pager for design documents operations.
func NewDesignDocsPager[O DesignDocsPagerOptions](c *cloudantv1.CloudantV1, o O) (Pager[cloudantv1.DocsResultRow], error) {
	return nil, ErrNotImplemented
}

func newDesignDocsPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostDesignDocsOptions) *keyPager[*cloudantv1.PostDesignDocsOptions, *cloudantv1.AllDocsResult, cloudantv1.DocsResultRow] {
	return new(keyPager[*cloudantv1.PostDesignDocsOptions, *cloudantv1.AllDocsResult, cloudantv1.DocsResultRow])
}
