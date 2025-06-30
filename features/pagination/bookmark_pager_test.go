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

var _ = Describe(`BookmarkPager tests`, func() {
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

	It(`Creates BookmarkPager`, func() {
		pd := newTestBookmarkPager(opts)
		pager := newBasePager(pd)

		Expect(pager).NotTo(BeNil())
		Expect(pager.pager).To(Equal(pd))
	})

	It(`Confirms BookmarkPager has default pageSize`, func() {
		pd := newTestBookmarkPager(opts)
		pager := newBasePager(pd)

		Expect(pager.pageSize).To(BeEquivalentTo(200))
	})

	It(`Confirms BookmarkPager sets pageSize from a limit`, func() {
		opts.SetLimit(42)
		pd := newTestBookmarkPager(opts)
		pager := newBasePager(pd)

		Expect(pager.pageSize).To(BeEquivalentTo(42))
	})

	It(`Confirms BookmarkPager GetNext returns a page with a total items number less than a limit`, func() {
		opts.SetLimit(21)
		pd := newTestBookmarkPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(20)
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.documents()))
		Expect(pager.HasNext()).To(BeFalse())
	})

	It(`Confirms BookmarkPager GetNext returns a page with a total items number equal to a limit`, func() {
		opts.SetLimit(14)
		pd := newTestBookmarkPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(14)
		items, err := pager.GetNextWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.documents()))
		Expect(pager.HasNext()).To(BeTrue())
		Expect(*pd.options.Bookmark).To(Equal("1"))

		items, err = pager.GetNextWithContext(ctx)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(BeEmpty())
		Expect(pager.HasNext()).To(BeFalse())
	})

	It(`Confirms BookmarkPager GetNext returns a page with  a total items number greater than limit`, func() {
		opts.SetLimit(7)
		pd := newTestBookmarkPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(9)
		items, err := pager.GetNextWithContext(ctx)
		docs := ms.documents()

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(docs[:7]))
		Expect(pager.HasNext()).To(BeTrue())

		items, err = pager.GetNextWithContext(ctx)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(docs[7:]))
		Expect(pager.HasNext()).To(BeFalse())
	})

	It(`Confirms BookmarkPager GetAll returns all items`, func() {
		opts.SetLimit(3)
		pd := newTestBookmarkPager(opts)
		pager := newBasePager(pd)

		ms.makeItems(3 * 12)
		items, err := pager.GetAllWithContext(ctx)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(items).To(Equal(ms.documents()))
		Expect(pager.HasNext()).To(BeFalse())
	})
})
