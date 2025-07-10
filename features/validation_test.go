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
	"github.com/IBM/go-sdk-core/v5/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`Validation tests`, func() {
	var (
		service *cloudantv1.CloudantV1
	)

	BeforeEach(func() {
		var serviceErr error
		service, serviceErr = cloudantv1.NewCloudantV1(
			&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			},
		)
		Expect(serviceErr).ShouldNot(HaveOccurred())
		Expect(service).ToNot(BeNil())
	})

	AfterEach(func() {
		service = nil
	})

	Context("with key pager type of options", func() {
		var (
			allDocsOptions *cloudantv1.PostAllDocsOptions
			viewOptions    *cloudantv1.PostViewOptions
		)

		BeforeEach(func() {
			allDocsOptions = service.NewPostAllDocsOptions("db")
			viewOptions = service.NewPostViewOptions("db", "ddoc", "view")
		})

		AfterEach(func() {
			allDocsOptions = nil
			viewOptions = nil
		})

		Context("with valid options", func() {
			// AfterEach is an actual assertion step and the setup is happening in It sections
			// this is recommended ginkgo v1 approach to the table tests
			AfterEach(func() {
				allDocsErr := validatePagerOptions(keyPagerValidationRules, allDocsOptions)
				viewErr := validatePagerOptions(keyPagerValidationRules, viewOptions)

				Expect(allDocsErr).ShouldNot(HaveOccurred())
				Expect(viewErr).ShouldNot(HaveOccurred())
			})

			It(`Confirms no validation error when limit is not set`, func() {
				Expect(allDocsOptions.Limit).To(BeNil())
				Expect(viewOptions.Limit).To(BeNil())
			})

			It(`Confirms no validation error on limit equal to min`, func() {
				allDocsOptions.SetLimit(1)
				viewOptions.SetLimit(1)
			})

			It(`Confirms no validation error on limit less than max`, func() {
				allDocsOptions.SetLimit(199)
				viewOptions.SetLimit(199)
			})

			It(`Confirms no validation error on limit equal to max`, func() {
				allDocsOptions.SetLimit(200)
				viewOptions.SetLimit(200)
			})
		})

		Context("with invalid options", func() {
			var errMsg string

			BeforeEach(func() {
				errMsg = ""
			})

			AfterEach(func() {
				allDocsErr := validatePagerOptions(keyPagerValidationRules, allDocsOptions)
				viewErr := validatePagerOptions(keyPagerValidationRules, viewOptions)

				Expect(allDocsErr).Should(HaveOccurred())
				Expect(allDocsErr.Error()).To(MatchRegexp(errMsg))

				Expect(viewErr).Should(HaveOccurred())
				Expect(viewErr.Error()).To(MatchRegexp(errMsg))
			})

			It(`Confirms validation error on limit less than min`, func() {
				allDocsOptions.SetLimit(0)
				viewOptions.SetLimit(0)
				errMsg = "the provided limit 0 is lower than the minimum page size"
			})

			It(`Confirms validation error on limit greater than max`, func() {
				allDocsOptions.SetLimit(201)
				viewOptions.SetLimit(201)
				errMsg = "the provided limit 201 exceeds the maximum page size"
			})

			It(`Confirms validation error on presence of keys`, func() {
				allDocsOptions.SetKeys([]string{"key1", "key2"})
				viewOptions.SetKeys([]any{"key1", "key2"})
				errMsg = `the option "Keys" is invalid when using pagination`
			})
		})

		It(`Confirms all docs validation error on presence of key`, func() {
			errMsg := `the option "Key" is invalid when using pagination. No need to paginate as "Key" returns a single result for an ID`
			opts := service.NewPostAllDocsOptions("db")
			opts.SetKey("key1")
			pagination := NewAllDocsPagination(service, opts)
			_, err := pagination.Pager()

			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(MatchRegexp(errMsg))
		})

		It(`Confirms views validation error on presence of key`, func() {
			errMsg := `the option "Key" is invalid when using pagination. Use StartKey and EndKey instead`
			opts := service.NewPostViewOptions("db", "ddoc", "view")
			opts.SetKey("key1")
			pagination := NewViewPagination(service, opts)
			_, err := pagination.Pager()

			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(MatchRegexp(errMsg))
		})
	})

	Context("with bookmark pager type of options", func() {
		var (
			findOptions   *cloudantv1.PostFindOptions
			searchOptions *cloudantv1.PostSearchOptions
		)

		BeforeEach(func() {
			selector := make(map[string]any, 0)
			findOptions = service.NewPostFindOptions("db", selector)
			searchOptions = service.NewPostSearchOptions("db", "ddoc", "index", "*:*")
		})

		AfterEach(func() {
			findOptions = nil
			searchOptions = nil
		})

		Context("with valid options", func() {
			// AfterEach is an actual assertion step and the setup is happening in It sections
			// this is recommended ginkgo v1 approach to the table tests
			AfterEach(func() {
				findErr := validatePagerOptions(bookmarkPagerValidationRules, findOptions)
				searchErr := validatePagerOptions(searchPagerValidationRules, searchOptions)

				Expect(findErr).ShouldNot(HaveOccurred())
				Expect(searchErr).ShouldNot(HaveOccurred())
			})

			It(`Confirms no validation error when limit is not set`, func() {
				Expect(findOptions.Limit).To(BeNil())
				Expect(searchOptions.Limit).To(BeNil())
			})

			It(`Confirms no validation error on limit equal to min`, func() {
				findOptions.SetLimit(1)
				searchOptions.SetLimit(1)
			})

			It(`Confirms no validation error on limit less than max`, func() {
				findOptions.SetLimit(199)
				searchOptions.SetLimit(199)
			})

			It(`Confirms no validation error on limit equal to max`, func() {
				findOptions.SetLimit(200)
				searchOptions.SetLimit(200)
			})
		})

		Context("with invalid options", func() {
			var errMsg string

			BeforeEach(func() {
				errMsg = ""
			})

			AfterEach(func() {
				findErr := validatePagerOptions(bookmarkPagerValidationRules, findOptions)
				searchErr := validatePagerOptions(searchPagerValidationRules, searchOptions)

				Expect(findErr).Should(HaveOccurred())
				Expect(findErr.Error()).To(MatchRegexp(errMsg))

				Expect(searchErr).Should(HaveOccurred())
				Expect(searchErr.Error()).To(MatchRegexp(errMsg))
			})

			It(`Confirms validation error on limit less than min`, func() {
				findOptions.SetLimit(0)
				searchOptions.SetLimit(0)
				errMsg = "the provided limit 0 is lower than the minimum page size"
			})

			It(`Confirms validation error on limit greater than max`, func() {
				findOptions.SetLimit(201)
				searchOptions.SetLimit(201)
				errMsg = "the provided limit 201 exceeds the maximum page size"
			})
		})

		Context("with search invalid options", func() {
			var errMsg string

			BeforeEach(func() {
				errMsg = ""
				searchOptions = service.NewPostSearchOptions("db", "ddoc", "index", "*:*")
			})

			AfterEach(func() {
				searchErr := validatePagerOptions(searchPagerValidationRules, searchOptions)

				Expect(searchErr).Should(HaveOccurred())
				Expect(searchErr.Error()).To(MatchRegexp(errMsg))

				searchOptions = nil
			})

			It(`Confirms validation error on on presence of Counts`, func() {
				searchOptions.SetCounts([]string{"aTestFieldToCount"})
				errMsg = `the option "Counts" is invalid when using pagination`
			})

			It(`Confirms validation error on on presence of GroupField`, func() {
				searchOptions.SetGroupField("testField")
				errMsg = `the option "GroupField" is invalid when using pagination`
			})

			It(`Confirms validation error on on presence of GroupLimit`, func() {
				searchOptions.SetGroupLimit(6)
				errMsg = `the option "GroupLimit" is invalid when using pagination`
			})

			It(`Confirms validation error on on presence of GroupSort`, func() {
				searchOptions.SetGroupSort([]string{"aTestFieldToGroupSort"})
				errMsg = `the option "GroupSort" is invalid when using pagination`
			})

			It(`Confirms validation error on on presence of Ranges`, func() {
				searchOptions.SetRanges(map[string]map[string]string{
					"aTestFieldForRanges": {
						"low":  "[0 to 5}",
						"high": "[5 to 10]",
					},
				})
				errMsg = `the option "Ranges" is invalid when using pagination`
			})
		})
	})
})
