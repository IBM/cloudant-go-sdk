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
	"fmt"
	"net/http"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`BasePager tests`, func() {
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

	It(`Creates BasePager`, func() {
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		Expect(pager).NotTo(BeNil())
		Expect(pager.pager).To(Equal(pd))
	})

	It(`Confirms BasePager has default pageSize`, func() {
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		Expect(pager.pageSize).To(BeEquivalentTo(20))
	})

	It(`Confirms BasePager sets pageSize from limit`, func() {
		opts.SetLimit(42)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		Expect(pager.pageSize).To(BeEquivalentTo(42))
	})

	It(`Confirms BasePager default HasNext is true`, func() {
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		Expect(pager.HasNext()).To(BeTrue())
	})

	It(`Confirms BasePager HasNext is true for results equal to limit`, func() {
		opts.SetLimit(1)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(1)
		pager.GetNextWithContext(ctx)

		Expect(pager.HasNext()).To(BeTrue())
	})

	It(`Confirms BasePager HasNext is false for results less than limit`, func() {
		opts.SetLimit(1)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		pager.GetNextWithContext(ctx)

		Expect(pager.HasNext()).To(BeFalse())
	})

	It(`Confirms BasePager GetNext returns first page`, func() {
		opts.SetLimit(25)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(25)
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.documents()))
	})

	It(`Confirms BasePager GetNext returns an error`, func() {
		opts.SetLimit(25)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(25)
		ms.setError(http.ErrServerClosed, 1)
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(http.ErrServerClosed))
		Expect(items).Should(BeEmpty())
	})

	It(`Confirms BasePager GetNext returns correct pages consistently`, func() {
		opts.SetLimit(3)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(2 * 3)
		page, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(page).To(Equal(ms.documents()[:3]))

		page, err = pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(page).To(Equal(ms.documents()[3:]))
	})

	It(`Confirms BasePager GetNext returns an error from base pager`, func() {
		opts.SetLimit(3)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(2 * 3)
		page, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(page).To(Equal(ms.documents()[:3]))

		pager.err = http.ErrServerClosed
		page, err = pager.GetNextWithContext(ctx)

		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(http.ErrServerClosed))
		Expect(page).Should(BeEmpty())
	})

	It(`Confirms BasePager GetNext is retriable`, func() {
		opts.SetLimit(3)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(2 * 3)
		ms.setError(http.ErrServerClosed, 4)
		page, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(page).To(Equal(ms.documents()[:3]))

		page, err = pager.GetNextWithContext(ctx)

		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(http.ErrServerClosed))
		Expect(page).Should(BeEmpty())

		ms.err = nil
		ms.errorItem = 0
		page, err = pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(page).To(Equal(ms.documents()[3:]))
	})

	It(`Confirms BasePager GetNext cycles until empty on items fit to exact number of pages`, func() {
		opts.SetLimit(3)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3 * 3)
		cycle := 0
		acc := make([]cloudantv1.Document, 0)
		for pager.HasNext() {
			cycle += 1
			items, err := pager.GetNextWithContext(ctx)

			Expect(err).ShouldNot(HaveOccurred())
			if cycle == 4 {
				Expect(items).To(BeEmpty())
			} else {
				Expect(items).Should(HaveLen(3))
			}
			acc = append(acc, items...)
		}
		Expect(cycle).To(Equal(4))
		Expect(acc).To(Equal(ms.documents()))
	})

	It(`Confirms BasePager GetNext cycles until empty on items exceeding exact number of pages`, func() {
		opts.SetLimit(3)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*3 + 1)
		cycle := 0
		acc := make([]cloudantv1.Document, 0)
		for pager.HasNext() {
			cycle += 1
			items, err := pager.GetNextWithContext(ctx)

			Expect(err).ShouldNot(HaveOccurred())
			switch cycle {
			case 5:
				Expect(items).To(BeEmpty())
			case 4:
				Expect(items).Should(HaveLen(1))
			default:
				Expect(items).Should(HaveLen(3))
			}
			acc = append(acc, items...)
		}
		Expect(cycle).To(Equal(4))
		Expect(acc).To(Equal(ms.documents()))
	})

	It(`Confirms BasePager GetNext returns an error when pager is exhausted`, func() {
		opts.SetLimit(2)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(1)
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.documents()))
		Expect(pager.HasNext()).To(BeFalse())

		items, err = pager.GetNextWithContext(ctx)

		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(ErrNoMoreResults))
		Expect(items).To(BeEmpty())
	})

	It(`Confirms BasePager GetAll returns all items`, func() {
		opts.SetLimit(11)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(71)
		items, err := pager.GetAllWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.documents()))
	})

	It(`Confirms BasePager GetAll returns an error`, func() {
		opts.SetLimit(11)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(71)
		ms.setError(http.ErrServerClosed, 12)
		items, err := pager.GetAllWithContext(ctx)

		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(http.ErrServerClosed))
		Expect(items).To(BeEmpty())
	})

	It(`Confirms BasePager GetAll is retriable`, func() {
		opts.SetLimit(11)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(71)
		ms.setError(http.ErrServerClosed, 12)
		items, err := pager.GetAllWithContext(ctx)

		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(http.ErrServerClosed))
		Expect(items).Should(BeEmpty())

		ms.err = nil
		ms.errorItem = 0
		items, err = pager.GetAllWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.documents()))
	})

	It(`Confirms BasePager Pages works as iterator`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		docs := ms.documents()
		pageNum := 1
		pageSize := 23
		for page, err := range pager.PagesWithContext(ctx) {
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

	It(`Confirms BasePager Pages supports break`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		docs := ms.documents()
		pageNum := 1
		pageSize := 23
		for page, err := range pager.PagesWithContext(ctx) {
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

	It(`Confirms BasePager Pages returns an error and stops cycle`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		ms.setError(http.ErrServerClosed, 24)
		docs := ms.documents()
		pageNum := 1
		pageSize := 23
		for page, err := range pager.PagesWithContext(ctx) {
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

	It(`Confirms BasePager Rows works as iterator`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		docs := ms.documents()
		i := 0
		for item, err := range pager.RowsWithContext(ctx) {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).To(Equal(docs[i]))
			i += 1
		}
	})

	It(`Confirms BasePager Rows supports break`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		docs := ms.documents()
		i := 0
		for item, err := range pager.RowsWithContext(ctx) {
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

	It(`Confirms BasePager Rows returns an error and stops cycle`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3*23 - 1)
		ms.setError(http.ErrServerClosed, 24)
		docs := ms.documents()
		i := 0
		for item, err := range pager.RowsWithContext(ctx) {
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

	It(`Confirms BasePager sets next page options`, func() {
		opts.SetLimit(1)
		pd := newTestPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(5)
		cycle := 0
		for pager.HasNext() {
			cycle += 1
			_, err := pager.GetNextWithContext(ctx)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(opts.Bookmark).To(BeNil())
			if pager.HasNext() {
				expect := fmt.Sprintf("%d", cycle)
				Expect(*pd.options.Bookmark).To(Equal(expect))
			} else {
				expect := fmt.Sprintf("%d", cycle-1)
				Expect(*pd.options.Bookmark).To(Equal(expect))
			}
		}
		Expect(cycle).To(Equal(6))
	})
})
