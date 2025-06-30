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
	"github.com/IBM/go-sdk-core/v5/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`KeyPager tests`, func() {
	var (
		service *cloudantv1.CloudantV1
		opts    *cloudantv1.PostViewOptions
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

		opts = service.NewPostViewOptions("db", "ddoc", "view")
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

	It(`Creates KeyPager`, func() {
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)

		Expect(pager).NotTo(BeNil())

		p := pager.(*basePager[*cloudantv1.PostViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow])

		Expect(p.pager).To(Equal(pd))
	})

	It(`Confirms KeyPager has default pageSize`, func() {
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)
		p := pager.(*basePager[*cloudantv1.PostViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow])

		Expect(p.pageSize).To(BeEquivalentTo(200))
	})

	It(`Confirms KeyPager sets pageSize from a limit`, func() {
		opts.SetLimit(42)
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)
		p := pager.(*basePager[*cloudantv1.PostViewOptions, *cloudantv1.ViewResult, cloudantv1.ViewResultRow])

		Expect(p.pageSize).To(BeEquivalentTo(42))
	})

	It(`Confirms KeyPager GetNext returns a page with a total items number less than a limit`, func() {
		opts.SetLimit(21)
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(20)
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.viewRows()))
		Expect(pager.HasNext()).To(BeFalse())
	})

	It(`Confirms KeyPager GetNext returns a page with a total items number equal to a limit`, func() {
		opts.SetLimit(14)
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(14)
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.viewRows()))
		Expect(pager.HasNext()).To(BeFalse())
	})

	It(`Confirms KeyPager GetNext returns a page with a total items number equal to a limit plus one`, func() {
		opts.SetLimit(14)
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(15)
		items, err := pager.GetNextWithContext(ctx)
		rows := ms.viewRows()

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(rows[:14]))
		Expect(pager.HasNext()).To(BeTrue())

		items, err = pager.GetNextWithContext(ctx)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(rows[14:]))
		Expect(pager.HasNext()).To(BeFalse())
	})

	It(`Confirms KeyPager GetNext returns a page with  a total items number greater than limit`, func() {
		opts.SetLimit(7)
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(9)
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.viewRows()[:7]))
		Expect(pager.HasNext()).To(BeTrue())

		items, err = pager.GetNextWithContext(ctx)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.viewRows()[7:]))
		Expect(pager.HasNext()).To(BeFalse())
	})

	It(`Confirms KeyPager GetAll returns all items`, func() {
		opts.SetLimit(3)
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3 * 12)
		items, err := pager.GetAllWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.viewRows()))
		Expect(pager.HasNext()).To(BeFalse())
	})

	It(`Confirms KeyPager doesn't return an error on a bundary check when not advanced on a key-duplicated page`, func() {
		opts.SetLimit(1)
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(1)
		ms.duplicateItem()
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.viewRows()[:1]))
		Expect(pager.HasNext()).To(BeTrue())
	})

	It(`Confirms KeyPager does return an error on a bundary check when advanced on a key-duplicated page`, func() {
		opts.SetLimit(1)
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(1)
		ms.duplicateItem()
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.viewRows()[:1]))
		Expect(pager.HasNext()).To(BeTrue())

		items, err = pager.GetNextWithContext(ctx)

		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).To(HavePrefix("cannot paginate on a boundary containing identical keys"))
		Expect(items).Should(BeEmpty())
	})

	It(`Confirms KeyPager doesn't return an error on a bundary check when key-duplicated are on the same page`, func() {
		opts.SetLimit(3)
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(1)
		ms.duplicateItem()
		ms.makeItems(2)
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.viewRows()[:3]))
		Expect(pager.HasNext()).To(BeTrue())
	})

	It(`Confirms KeyPager doesn't return an error on a bundary check when there is no items left`, func() {
		opts.SetLimit(2)
		pd := newTestKeyPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(1)
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.viewRows()))
		Expect(pager.HasNext()).To(BeFalse())

		items, err = pager.GetNextWithContext(ctx)

		Expect(err).Should(HaveOccurred())
		Expect(err).Should(MatchError(ErrNoMoreResults))
		Expect(items).To(BeEmpty())
	})
})
