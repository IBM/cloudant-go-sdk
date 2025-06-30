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

package main

import (
	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/cloudant-go-sdk/features"
)

func main() {

	// Initialize service
	service, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{},
	)
	if err != nil {
		panic(err)
	}

	// Setup options
	opts := service.NewPostPartitionSearchOptions("events", "ns1HJS13AMkK", "checkout", "findByDate", "date:[2019-01-01T12:00:00.000Z TO 2019-01-31T12:00:00.000Z]")
	opts.SetLimit(50)

	// Create pagination
	// pagination can be reused without side-effects as a factory for iterators or pagers
	// options are fixed at pagination creation time
	pagination := features.NewSearchPagination(service, opts)

	// Option: iterate pages
	// Ideal for using a for/range loop with each page.
	// The iter.Seq2 returned from Pages() has an error as a second element.
	for page, err := range pagination.Pages() {
		// Break on err != nil
		// Do something with page
	}

	// Option: iterate rows
	// Ideal for using a for/range loop with each row.
	// The iter.Seq2 returned from Rows() has an error as a second element.
	for row, err := range pagination.Rows() {
		// Break on err != nil
		// Do something with row
	}

	// Option: use pager next page
	// For retrieving one page at a time with a method call.
	pager, err := pagination.Pager()
	if err != nil {
		panic(err)
	}

	for pager.HasNext() {
		page, err := pager.GetNext()
		// Break on err != nil
		// Do something with page
	}

	// Option: use pager all results
	// For retrieving all result rows in a single list
	allPager, err := pagination.Pager()
	if err != nil {
		panic(err)
	}

	// Note: all result rows may be very large!
	// Preferably use pagination.Rows() iterator for memory efficiency with large result sets.
	allRows, err := pager.GetAll()
	if err != nil {
		panic(err)
	}

	for _, row := range allRows {
		// Do something with row
	}
}
