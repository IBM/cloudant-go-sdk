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

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`BasePager tests`, func() {
	var (
		service *cloudantv1.CloudantV1
		opts    *cloudantv1.PostFindOptions
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
	})

	AfterEach(func() {
		service = nil
		opts = nil
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
		pd.makeItems(1)
		pager := newBasePager(pd)
		pager.GetNext()

		Expect(pager.HasNext()).To(BeTrue())
	})

	It(`Confirms BasePager HasNext is false for results less than limit`, func() {
		opts.SetLimit(1)
		pd := newTestPager(opts)
		pager := newBasePager(pd)
		pager.GetNext()

		Expect(pager.HasNext()).To(BeFalse())
	})

	It(`Confirms BasePager GetNext returns first page`, func() {
		opts.SetLimit(25)
		pd := newTestPager(opts)
		pd.makeItems(25)
		pager := newBasePager(pd)
		items, err := pager.GetNext()

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(pd.items))
	})

	It(`Confirms BasePager GetNext returns an error`, func() {
		opts.SetLimit(25)
		pd := newTestPager(opts)
		pd.makeItems(25)
		pd.setError(http.ErrServerClosed, 1)
		pager := newBasePager(pd)
		items, err := pager.GetNext()

		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(http.ErrServerClosed))
		Expect(items).Should(BeEmpty())
	})

	It(`Confirms BasePager GetNext returns correct pages consistently`, func() {
		opts.SetLimit(3)
		pd := newTestPager(opts)
		pd.makeItems(2 * 3)
		pager := newBasePager(pd)

		page, err := pager.GetNext()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(page).To(Equal(pd.items[:3]))

		page, err = pager.GetNext()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(page).To(Equal(pd.items[3:]))
	})

	It(`Confirms BasePager GetNext returns an error from base pager`, func() {
		opts.SetLimit(3)
		pd := newTestPager(opts)
		pd.makeItems(2 * 3)
		pager := newBasePager(pd)

		page, err := pager.GetNext()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(page).To(Equal(pd.items[:3]))

		pager.err = http.ErrServerClosed

		page, err = pager.GetNext()
		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(http.ErrServerClosed))
		Expect(page).Should(BeEmpty())
	})

	It(`Confirms BasePager GetNext is retriable`, func() {
		opts.SetLimit(3)
		pd := newTestPager(opts)
		pd.makeItems(2 * 3)
		pd.setError(http.ErrServerClosed, 4)
		pager := newBasePager(pd)

		page, err := pager.GetNext()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(page).To(Equal(pd.items[:3]))

		page, err = pager.GetNext()
		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(http.ErrServerClosed))
		Expect(page).Should(BeEmpty())

		pd.err = nil
		pd.errorItem = 0

		page, err = pager.GetNext()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(page).To(Equal(pd.items[3:]))
	})

	It(`Confirms BasePager GetNext cycles until empty on items fit to exact number of pages`, func() {
		opts.SetLimit(3)
		pd := newTestPager(opts)
		pd.makeItems(3 * 3)
		pager := newBasePager(pd)
		cycle := 0
		acc := make([]cloudantv1.Document, 0)
		for pager.HasNext() {
			cycle += 1
			items, err := pager.GetNext()

			Expect(err).ShouldNot(HaveOccurred())
			if cycle == 4 {
				Expect(items).To(BeEmpty())
			} else {
				Expect(items).Should(HaveLen(3))
			}
			acc = append(acc, items...)
		}
		Expect(cycle).To(Equal(4))
		Expect(acc).To(Equal(pd.items))
	})

	It(`Confirms BasePager GetNext cycles until empty on items exceeding exact number of pages`, func() {
		opts.SetLimit(3)
		pd := newTestPager(opts)
		pd.makeItems(3*3 + 1)
		pager := newBasePager(pd)
		cycle := 0
		acc := make([]cloudantv1.Document, 0)
		for pager.HasNext() {
			cycle += 1
			items, err := pager.GetNext()

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
		Expect(acc).To(Equal(pd.items))
	})

	It(`Confirms BasePager GetNext returns an error when pager is exhausted`, func() {
		opts.SetLimit(2)
		pd := newTestPager(opts)
		pd.makeItems(1)
		pager := newBasePager(pd)

		items, err := pager.GetNext()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(pd.items))

		Expect(pager.HasNext()).To(BeFalse())

		items, err = pager.GetNext()
		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(ErrNoMoreResults))
		Expect(items).To(BeEmpty())
	})

	It(`Confirms BasePager GetAll returns all items`, func() {
		opts.SetLimit(11)
		pd := newTestPager(opts)
		pd.makeItems(71)
		pager := newBasePager(pd)

		items, err := pager.GetAll()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(pd.items))
	})

	It(`Confirms BasePager GetAll returns an error`, func() {
		opts.SetLimit(11)
		pd := newTestPager(opts)
		pd.makeItems(71)
		pd.setError(http.ErrServerClosed, 12)
		pager := newBasePager(pd)

		items, err := pager.GetAll()
		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(http.ErrServerClosed))
		Expect(items).To(BeEmpty())
	})

	It(`Confirms BasePager GetAll is retriable`, func() {
		opts.SetLimit(11)
		pd := newTestPager(opts)
		pd.makeItems(71)
		pd.setError(http.ErrServerClosed, 12)
		pager := newBasePager(pd)

		items, err := pager.GetAll()
		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(http.ErrServerClosed))
		Expect(items).Should(BeEmpty())

		pd.err = nil
		pd.errorItem = 0

		items, err = pager.GetAll()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(pd.items))
	})

	It(`Confirms BasePager Pages works as iterator`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pd.makeItems(3*23 - 1)
		pager := newBasePager(pd)

		pageNum := 1
		pageSize := 23
		for page, err := range pager.Pages() {
			start := (pageNum - 1) * pageSize
			end := start + pageSize
			// last page
			if pageNum == 3 {
				end -= 1
				pageSize -= 1
			}
			Expect(err).ShouldNot(HaveOccurred())
			Expect(page).Should(HaveLen(pageSize))
			Expect(page).To(Equal(pd.items[start:end]))
			pageNum += 1
		}
	})

	It(`Confirms BasePager Pages supports break`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pd.makeItems(3*23 - 1)
		pager := newBasePager(pd)

		pageNum := 1
		pageSize := 23
		for page, err := range pager.Pages() {
			start := (pageNum - 1) * pageSize
			end := start + pageSize
			Expect(err).ShouldNot(HaveOccurred())
			Expect(page).Should(HaveLen(pageSize))
			Expect(page).To(Equal(pd.items[start:end]))
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
		pd.makeItems(3*23 - 1)
		pd.setError(http.ErrServerClosed, 24)
		pager := newBasePager(pd)

		pageNum := 1
		pageSize := 23
		for page, err := range pager.Pages() {
			if pageNum == 2 {
				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError(http.ErrServerClosed))
				continue
			}
			start := (pageNum - 1) * pageSize
			end := start + pageSize
			Expect(err).ShouldNot(HaveOccurred())
			Expect(page).To(Equal(pd.items[start:end]))
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
		pd.makeItems(3*23 - 1)
		pager := newBasePager(pd)

		i := 0
		for item, err := range pager.Rows() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).To(Equal(pd.items[i]))
			i += 1
		}
	})

	It(`Confirms BasePager Rows supports break`, func() {
		opts.SetLimit(23)
		pd := newTestPager(opts)
		pd.makeItems(3*23 - 1)
		pager := newBasePager(pd)

		i := 0
		for item, err := range pager.Rows() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).To(Equal(pd.items[i]))
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
		pd.makeItems(3*23 - 1)
		pd.setError(http.ErrServerClosed, 24)
		pager := newBasePager(pd)

		i := 0
		for item, err := range pager.Rows() {
			if i == 23 {
				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError(http.ErrServerClosed))
				continue
			}
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).To(Equal(pd.items[i]))
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
		pd.makeItems(5)
		pager := newBasePager(pd)
		cycle := 0
		for pager.HasNext() {
			cycle += 1
			_, err := pager.GetNext()

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
