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

func newViewPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostViewOptions) *keyPager[*cloudantv1.PostViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow] {
	return new(keyPager[*cloudantv1.PostViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow])
}

func newViewPartitionPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostPartitionViewOptions) *keyPager[*cloudantv1.PostPartitionViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow] {
	return new(keyPager[*cloudantv1.PostPartitionViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow])
}
