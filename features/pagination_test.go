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
	"net/http"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`Iterators tests`, func() {
	var (
		service *cloudantv1.CloudantV1
		opts    *cloudantv1.PostFindOptions
		ms      *mockService
		ctx     context.Context
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

		selector := make(map[string]any)
		opts = service.NewPostFindOptions("db", selector)
		Expect(opts).ToNot(BeNil())

		ms = newMockService()
		Expect(ms).ToNot(BeNil())

		ctx = toContext(ms)
	})

	AfterEach(func() {
		service = nil
		opts = nil
		ms = nil
		ctx = nil
	})

	It(`Confirms iterator Pages() accepts BasePager`, func() {
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		for _, err := range PagesWithContext(ctx, pager) {
			Expect(err).ShouldNot(HaveOccurred())
		}
	})

	It(`Confirms iterator Pages() works`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		docs := ms.documents()
		pageNum := 1
		pageSize := 23
		for page, err := range PagesWithContext(ctx, pager) {
			start := (pageNum - 1) * pageSize
			end := start + pageSize
			// last page
			if pageNum == 3 {
				end -= 1
				pageSize -= 1
			}
			Expect(err).ShouldNot(HaveOccurred())
			Expect(page).Should(HaveLen(pageSize))
			Expect(page).To(Equal(docs[start:end]))
			pageNum += 1
		}
	})

	It(`Confirms iterator Pages() supports break`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		docs := ms.documents()
		pageNum := 1
		pageSize := 23
		for page, err := range PagesWithContext(ctx, pager) {
			start := (pageNum - 1) * pageSize
			end := start + pageSize
			Expect(err).ShouldNot(HaveOccurred())
			Expect(page).Should(HaveLen(pageSize))
			Expect(page).To(Equal(docs[start:end]))
			if pageNum == 2 {
				break
			}
			pageNum += 1
		}
		Expect(pager.HasNext()).To(BeTrue())
		Expect(*pd.options.Bookmark).To(Equal("2"))
	})

	It(`Confirms iterator Pages() returns an error and stops cycle`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		ms.setError(http.ErrServerClosed, 24)
		docs := ms.documents()
		pageNum := 1
		pageSize := 23
		for page, err := range PagesWithContext(ctx, pager) {
			if pageNum == 2 {
				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError(http.ErrServerClosed))
				continue
			}
			start := (pageNum - 1) * pageSize
			end := start + pageSize
			Expect(err).ShouldNot(HaveOccurred())
			Expect(page).To(Equal(docs[start:end]))
			pageNum += 1
		}
		Expect(pager.HasNext()).To(BeTrue())
		// iterator exits on error even with continue
		Expect(pageNum).To(Equal(2))
		Expect(*pd.options.Bookmark).To(Equal("1"))
	})

	It(`Confirms iterator Rows() accepts BasePager`, func() {
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		for _, err := range RowsWithContext(ctx, pager) {
			Expect(err).ShouldNot(HaveOccurred())
		}
	})

	It(`Confirms iterator Rows() works`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		docs := ms.documents()
		i := 0
		for item, err := range RowsWithContext(ctx, pager) {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).To(Equal(docs[i]))
			i += 1
		}
	})

	It(`Confirms iterator Rows() supports break`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		docs := ms.documents()
		i := 0
		for item, err := range RowsWithContext(ctx, pager) {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).To(Equal(docs[i]))
			i += 1
			if i == 24 {
				break
			}
		}
		Expect(pager.HasNext()).To(BeTrue())
		Expect(*pd.options.Bookmark).To(Equal("2"))
	})

	It(`Confirms iterator Rows() returns an error and stops cycle`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		ms.setError(http.ErrServerClosed, 24)
		docs := ms.documents()
		i := 0
		for item, err := range RowsWithContext(ctx, pager) {
			if i == 23 {
				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError(http.ErrServerClosed))
				continue
			}
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).To(Equal(docs[i]))
			i += 1
		}
		Expect(pager.HasNext()).To(BeTrue())
		// iterator exits cycle on error
		Expect(i).To(Equal(23))
		Expect(*pd.options.Bookmark).To(Equal("1"))
	})

	It(`Confirms pagination Pager() returns a new pager`, func() {
		expectPager := newBasePager(newTestPager(opts))
		pagination := &paginationImplementor[*cloudantv1.PostFindOptions, cloudantv1.Document]{
			service: service,
			options: opts,
			newPager: func(c *cloudantv1.CloudantV1, opts *cloudantv1.PostFindOptions) (Pager[cloudantv1.Document], error) {
				return newBasePager(newTestPager(opts)), nil
			},
		}

		pager1, err1 := pagination.Pager()

		Expect(err1).ShouldNot(HaveOccurred())
		Expect(pager1).NotTo(BeNil())
		Expect(pager1).To(BeAssignableToTypeOf(expectPager))

		pager2, err2 := pagination.Pager()

		Expect(err2).ShouldNot(HaveOccurred())
		Expect(pager2).NotTo(BeNil())
		Expect(pager2).To(BeAssignableToTypeOf(expectPager))

		Expect(pager2).ToNot(BeIdenticalTo(pager1))
		Expect(pager2).To(Equal(pager1))
	})

	It(`Confirms pagination Pages() works`, func() {
		opts.SetLimit(23)
		pagination := &paginationImplementor[*cloudantv1.PostFindOptions, cloudantv1.Document]{
			service: service,
			options: opts,
			newPager: func(c *cloudantv1.CloudantV1, opts *cloudantv1.PostFindOptions) (Pager[cloudantv1.Document], error) {
				return newBasePager(newTestPager(opts)), nil
			},
		}

		ms.makeItems(3*23 - 1)
		docs := ms.documents()
		pageNum := 1
		pageSize := 23
		for page, err := range pagination.PagesWithContext(ctx) {
			start := (pageNum - 1) * pageSize
			end := start + pageSize
			// last page
			if pageNum == 3 {
				end -= 1
				pageSize -= 1
			}
			Expect(err).ShouldNot(HaveOccurred())
			Expect(page).Should(HaveLen(pageSize))
			Expect(page).To(Equal(docs[start:end]))
			pageNum += 1
		}
	})

	It(`Confirms pagination Rows() works`, func() {
		opts.SetLimit(23)
		pagination := &paginationImplementor[*cloudantv1.PostFindOptions, cloudantv1.Document]{
			service: service,
			options: opts,
			newPager: func(c *cloudantv1.CloudantV1, opts *cloudantv1.PostFindOptions) (Pager[cloudantv1.Document], error) {
				return newBasePager(newTestPager(opts)), nil
			},
		}

		ms.makeItems(3*23 - 1)
		docs := ms.documents()
		i := 0
		for item, err := range pagination.RowsWithContext(ctx) {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).To(Equal(docs[i]))
			i += 1
		}
	})
})
