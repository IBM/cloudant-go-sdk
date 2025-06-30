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

func newFindPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostFindOptions) *bookmarkPager[*cloudantv1.PostFindOptions, *cloudantv1.FindResult, cloudantv1.Document] {
	return new(bookmarkPager[*cloudantv1.PostFindOptions, *cloudantv1.FindResult, cloudantv1.Document])
}

func newFindPartitionPager(c *cloudantv1.CloudantV1, o *cloudantv1.PostPartitionFindOptions) *bookmarkPager[*cloudantv1.PostPartitionFindOptions, *cloudantv1.FindResult, cloudantv1.Document] {
	return new(bookmarkPager[*cloudantv1.PostPartitionFindOptions, *cloudantv1.FindResult, cloudantv1.Document])
}
