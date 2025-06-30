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
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`Search pager tests`, func() {
	var (
		ms               *mockService
		server           *httptest.Server
		service          *cloudantv1.CloudantV1
		expectPages      int
		expectItems      int
		expectStatusCode int
		expectedError    string
	)

	BeforeEach(func() {
		ms = newMockService()

		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer GinkgoRecover()
			mockServerCallback(w, r, ms)
		}))

		var serviceErr error
		service, serviceErr = cloudantv1.NewCloudantV1(
			&cloudantv1.CloudantV1Options{
				URL:           server.URL,
				Authenticator: &core.NoAuthAuthenticator{},
			},
		)
		Expect(serviceErr).ShouldNot(HaveOccurred())
		Expect(service).ToNot(BeNil())
		service.SetEnableGzipCompression(false)
	})

	AfterEach(func() {
		server.Close()
		ms = nil
		service = nil
	})

	Context("with PostSearch options:", func() {
		Context("Successful cases", func() {
			AfterEach(func() {
				opts := service.NewPostSearchOptions("db", "ddoc", "index", "*:*")
				opts.SetLimit(int64(defaultTestPageSize))
				pagination := NewSearchPagination(service, opts)

				Expect(pagination).ToNot(BeNil())

				ms.makeItems(expectItems)

				runPagesAssertion(pagination, expectPages)
				runRowsAssertion(pagination, expectItems)
				runPagerAssertion(pagination, expectPages, expectItems)
			})

			for _, test := range concretePagerTestCases {
				It(test.descrition, func() {
					expectPages = test.expectPages
					// we expect one last empty page on none-partial pages from bookmark pagers
					if test.expectItems > 0 && test.expectItems%defaultTestPageSize == 0 {
						expectPages += 1
					}
					expectItems = test.expectItems
				})
			}
		})

		Context("Error cases", func() {
			AfterEach(func() {
				opts := service.NewPostSearchOptions("db", "ddoc", "index", "*:*")
				opts.SetLimit(int64(defaultTestPageSize))
				pagination := NewSearchPagination(service, opts)

				Expect(pagination).ToNot(BeNil())

				ms.makeItems(expectItems)
				ms.setHTTPError(expectStatusCode, expectItems-2)

				runPagesWithErrorAssertion(pagination, expectedError, expectPages)
				runRowsWithErrorAssertion(pagination, expectedError, expectPages*defaultTestPageSize)
				runPagerWithErrorAssertion(pagination, expectedError, expectItems)
			})

			for pageNum := range 2 {
				for _, code := range append(terminalErrors, transientErrors...) {
					It(fmt.Sprintf("Confirms error is returned on page %d for code %d", (pageNum+1), code), func() {
						expectPages = pageNum
						expectItems = defaultTestPageSize * (pageNum + 1)
						expectStatusCode = code
						expectedError = statusText(expectStatusCode)
					})
				}
			}
		})
	})

	Context("with PostPartitionSearch options:", func() {
		Context("Successful cases", func() {
			AfterEach(func() {
				opts := service.NewPostPartitionSearchOptions("db", "partition", "ddoc", "index", "*:*")
				opts.SetLimit(int64(defaultTestPageSize))
				pagination := NewSearchPagination(service, opts)

				Expect(pagination).ToNot(BeNil())

				ms.makeItems(expectItems)

				runPagesAssertion(pagination, expectPages)
				runRowsAssertion(pagination, expectItems)
				runPagerAssertion(pagination, expectPages, expectItems)
			})

			for _, test := range concretePagerTestCases {
				It(test.descrition, func() {
					expectPages = test.expectPages
					if test.expectItems > 0 && test.expectItems%defaultTestPageSize == 0 {
						expectPages += 1
					}
					expectItems = test.expectItems
				})
			}
		})

		Context("Error cases", func() {
			AfterEach(func() {
				opts := service.NewPostPartitionSearchOptions("db", "partition", "ddoc", "index", "*:*")
				opts.SetLimit(int64(defaultTestPageSize))
				pagination := NewSearchPagination(service, opts)

				Expect(pagination).ToNot(BeNil())

				ms.makeItems(expectItems)
				ms.setHTTPError(expectStatusCode, expectItems-2)

				runPagesWithErrorAssertion(pagination, expectedError, expectPages)
				runRowsWithErrorAssertion(pagination, expectedError, expectPages*defaultTestPageSize)
				runPagerWithErrorAssertion(pagination, expectedError, expectItems)
			})

			for pageNum := range 2 {
				for _, code := range append(terminalErrors, transientErrors...) {
					It(fmt.Sprintf("Confirms error is returned on page %d for code %d", (pageNum+1), code), func() {
						expectPages = pageNum
						expectItems = defaultTestPageSize * (pageNum + 1)
						expectStatusCode = code
						expectedError = statusText(expectStatusCode)
					})
				}
			}
		})
	})
})
