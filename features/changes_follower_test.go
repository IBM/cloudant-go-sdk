/**
 * © Copyright IBM Corporation 2022, 2023. All Rights Reserved.
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
	"errors"
	"fmt"
	"time"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

var expectedErrType *core.SDKProblem

var _ = Describe(`ChangesFollower initialization`, func() {
	var (
		service            *cloudantv1.CloudantV1
		postChangesOptions *cloudantv1.PostChangesOptions
	)

	BeforeEach(func() {
		var serviceErr error
		service, serviceErr = cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
			URL:           "http://localhost:5984",
			Authenticator: &core.NoAuthAuthenticator{},
		})

		Expect(serviceErr).ShouldNot(HaveOccurred())
		Expect(service).ToNot(BeNil())

		postChangesOptions = service.NewPostChangesOptions("db")
	})

	It(`Creates minimal ChangesFollower successfully`, func() {
		follower, followerErr := NewChangesFollower(service, postChangesOptions)

		Expect(followerErr).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())
	})

	It(`Creates minimal ChangesFollower with context successfully`, func() {
		ctx := context.Background()
		follower, followerErr := NewChangesFollowerWithContext(ctx, service, postChangesOptions)

		Expect(followerErr).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())
	})

	It(`Validates missing database name`, func() {
		postChangesOptions = service.NewPostChangesOptions("")
		follower, followerErr := NewChangesFollower(service, postChangesOptions)

		Expect(follower).To(BeNil())
		Expect(followerErr).Should(HaveOccurred())
		Expect(followerErr.Error()).To(MatchRegexp("Field validation for 'Db' failed"))
		Expect(errors.As(followerErr, &expectedErrType)).To(BeTrue())
	})

	Context("With valid PostChangesOptions", func() {
		var (
			follower    *ChangesFollower
			followerErr error
		)

		BeforeEach(func() {
			postChangesOptions := service.NewPostChangesOptions("db")
			follower, followerErr = NewChangesFollower(service, postChangesOptions)

			Expect(followerErr).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())
		})

		It(`Validates valid tolerance`, func() {
			err := follower.SetErrorTolerance(time.Second)

			Expect(err).ShouldNot(HaveOccurred())
		})

		It(`Validates negative tolerance`, func() {
			err := follower.SetErrorTolerance(-1 * time.Millisecond)

			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("error tolerance duration must not be negative"))
			Expect(errors.As(err, &expectedErrType)).To(BeTrue())
		})
	})

	It(`Creates ChangesFollower with valid client timeout`, func() {
		timeouts := []time.Duration{
			1 * time.Minute,
			2 * time.Minute,
			5 * time.Minute,
		}
		for _, timeout := range timeouts {
			client := core.DefaultHTTPClient()
			client.Timeout = timeout
			service.Service.SetHTTPClient(client)

			postChangesOptions := service.NewPostChangesOptions("db")
			follower, followerErr := NewChangesFollower(service, postChangesOptions)

			Expect(followerErr).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())
		}
	})

	It(`Fails to create ChangesFollower with invalid client timeout`, func() {
		timeouts := []time.Duration{
			15 * time.Second,
			30 * time.Second,
			LongpollTimeout,
		}
		for _, timeout := range timeouts {
			client := core.DefaultHTTPClient()
			client.Timeout = timeout
			service.Service.SetHTTPClient(client)

			postChangesOptions := service.NewPostChangesOptions("db")
			follower, followerErr := NewChangesFollower(service, postChangesOptions)

			Expect(follower).To(BeNil())
			Expect(followerErr).Should(HaveOccurred())
			Expect(followerErr.Error()).To(MatchRegexp("timeout must be at least 60000 ms"))
			Expect(errors.As(followerErr, &expectedErrType)).To(BeTrue())
		}
	})
})

var _ = Describe(`ChangesFollower options`, func() {
	var (
		service    *cloudantv1.CloudantV1
		serviceErr error
	)

	BeforeEach(func() {
		service, serviceErr = cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
			URL:           "http://localhost:5984",
			Authenticator: &core.NoAuthAuthenticator{},
		})

		Expect(serviceErr).ShouldNot(HaveOccurred())
		Expect(service).ToNot(BeNil())
	})

	Context("With valid PostChangesOptions", func() {
		var postChangesOptions *cloudantv1.PostChangesOptions

		BeforeEach(func() {
			postChangesOptions = service.NewPostChangesOptions("db")
			postChangesOptions.SetIncludeDocs(true)
			postChangesOptions.SetDocIds([]string{"foo", "bar", "baz"})
			postChangesOptions.SetAttEncodingInfo(true)
			postChangesOptions.SetAttachments(true)
			postChangesOptions.SetConflicts(true)
			postChangesOptions.SetFilter("_selector")
			postChangesOptions.SetSelector(map[string]interface{}{
				"selector": map[string]interface{}{"foo": "bar"},
			})

			Expect(postChangesOptions).ToNot(BeNil())
		})

		It(`Validate options valid cases`, func() {
			follower, followerErr := NewChangesFollower(service, postChangesOptions)

			Expect(followerErr).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())
		})

		It(`Set defaults`, func() {
			follower, followerErr := NewChangesFollower(service, postChangesOptions)

			Expect(followerErr).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())
			Expect(follower.mode).To(Equal(Finite))
			Expect(follower.options).ToNot(BeNil())

			o := follower.options
			Expect(*o.Feed).Should(Equal(cloudantv1.PostChangesOptionsFeedNormalConst))

			follower.mode = Listen
			follower.setOptionsDefaults()

			Expect(*o.Feed).Should(Equal(cloudantv1.PostChangesOptionsFeedLongpollConst))
			Expect(*o.Timeout).Should(Equal(LongpollTimeout.Milliseconds()))
		})

		It(`Set defaults with limit`, func() {
			follower, followerErr := NewChangesFollower(service, postChangesOptions)

			Expect(followerErr).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())
			Expect(follower.mode).To(Equal(Finite))
			Expect(follower.options).ToNot(BeNil())

			follower.setOptionsDefaults().withLimit(12)

			o := follower.options
			Expect(*o.Feed).Should(Equal(cloudantv1.PostChangesOptionsFeedNormalConst))

			follower.mode = Listen
			follower.setOptionsDefaults().withLimit(12)

			Expect(*o.Feed).Should(Equal(cloudantv1.PostChangesOptionsFeedLongpollConst))
			Expect(*o.Timeout).Should(Equal(LongpollTimeout.Milliseconds()))
			Expect(*o.Limit).Should(BeEquivalentTo(12))
		})

		It(`Set defaults with PostChangesOptions limit`, func() {
			postChangesOptions.SetLimit(24)
			follower, followerErr := NewChangesFollower(service, postChangesOptions)

			Expect(followerErr).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())
			Expect(follower.mode).To(Equal(Finite))
			Expect(follower.options).ToNot(BeNil())

			follower.setOptionsDefaults().withLimit(12)

			o := follower.options
			Expect(*o.Feed).Should(Equal(cloudantv1.PostChangesOptionsFeedNormalConst))

			follower.mode = Listen
			follower.setOptionsDefaults().withLimit(12)

			Expect(*o.Feed).Should(Equal(cloudantv1.PostChangesOptionsFeedLongpollConst))
			Expect(*o.Timeout).Should(Equal(LongpollTimeout.Milliseconds()))
			Expect(*o.Limit).Should(BeEquivalentTo(12))
		})
	})

	Context("With invalid PostChangesOptions", func() {
		var postChangesOptions *cloudantv1.PostChangesOptions
		var errFmt = "the option '%s' is invalid when using ChangesFollower"
		var errMsg string

		BeforeEach(func() {
			postChangesOptions = service.NewPostChangesOptions("db")

			Expect(postChangesOptions).ToNot(BeNil())
		})

		// AfterEach is an actual assertion step
		// the setup is happening in BeforeEach and each It sections
		// this is recommended ginkgo v1 approach to the table tests
		AfterEach(func() {
			follower, followerErr := NewChangesFollower(service, postChangesOptions)

			Expect(follower).To(BeNil())
			Expect(followerErr.Error()).To(MatchRegexp(errMsg))
			Expect(errors.As(followerErr, &expectedErrType)).To(BeTrue())
		})

		It(`Validate invalid option descending`, func() {
			postChangesOptions.SetDescending(true)
			errMsg = fmt.Sprintf(errFmt, "descending")
		})

		It(`Validate invalid option feed`, func() {
			postChangesOptions.SetFeed(cloudantv1.PostChangesOptionsFeedContinuousConst)
			errMsg = fmt.Sprintf(errFmt, "feed")
		})

		It(`Validate invalid option heartbeat`, func() {
			postChangesOptions.SetHeartbeat(150)
			errMsg = fmt.Sprintf(errFmt, "heartbeat")
		})

		It(`Validate invalid option lastEventId`, func() {
			postChangesOptions.SetLastEventID("9876-alotofcharactersthatarenotreallyrandom")
			errMsg = fmt.Sprintf(errFmt, "lastEventId")
		})

		It(`Validate invalid option timeout`, func() {
			postChangesOptions.SetTimeout(time.Hour.Milliseconds())
			errMsg = fmt.Sprintf(errFmt, "timeout")
		})

		It(`Validate invalid option filter`, func() {
			postChangesOptions.SetFilter("_view")
			errMsg = fmt.Sprintf(errFmt, "filter=_view")
		})

		It(`Validate options multiple invalid cases`, func() {
			postChangesOptions.SetDescending(true)
			postChangesOptions.SetFeed(cloudantv1.PostChangesOptionsFeedContinuousConst)
			postChangesOptions.SetHeartbeat(150)
			postChangesOptions.SetLastEventID("9876-alotofcharactersthatarenotreallyrandom")
			postChangesOptions.SetTimeout(time.Hour.Milliseconds())
			errMsg = "the options descending, feed, heartbeat, lastEventId, timeout are invalid when using ChangesFollower"
		})
	})
})

var _ = Describe(`ChangesFollower finite`, func() {
	var p = 100 * time.Millisecond

	It(`Checks that a FINITE mode completes successfully for a fixed number of batches.`, func() {
		batches := 6
		ms := NewMockServer(batches, noErrors)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		changesCh, err := follower.StartOneOff()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(changesCh).ToNot(BeNil())

		count := 0
		for ci := range changesCh {
			item, err := ci.Item()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).ToNot(Equal(cloudantv1.ChangesResultItem{}))
			count += 1
		}
		Expect(count).To(Equal(batches * BatchSize))
	})

	It(`Checks that a FINITE mode errors for all terminal errors.`, func() {
		for _, error := range terminalErrors {
			ms := NewMockErrorServer(error)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			changesCh, err := follower.StartOneOff()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(changesCh).ToNot(BeNil())

			ci := <-changesCh
			item, err := ci.Item()
			Expect(item).To(Equal(cloudantv1.ChangesResultItem{}))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(MatchRegexp(ErrorText(error)))
			Expect(errors.As(err, &expectedErrType)).To(BeTrue())

			ms.Stop()
		}
	})

	It(`Checks that a FINITE mode errors for all transient errors when not suppressing.`, func() {
		for _, error := range transientErrors {
			e := gmeasure.NewExperiment("measure runner duration")
			ms := NewMockErrorServer(error)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			err = follower.SetErrorTolerance(0)
			Expect(err).ShouldNot(HaveOccurred())

			changesCh, err := follower.StartOneOff()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(changesCh).ToNot(BeNil())

			e.MeasureDuration("stop after", func() {
				ci := <-changesCh
				item, err := ci.Item()
				Expect(item).To(Equal(cloudantv1.ChangesResultItem{}))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp(ErrorText(error)))
				Expect(errors.As(err, &expectedErrType)).To(BeTrue())
			}, gmeasure.Precision(p))
			runDuration := e.Get("stop after").Durations[0]
			Expect(runDuration).To(BeNumerically("<", 100*time.Millisecond))

			ms.Stop()
		}
	})

	It(`Checks that a FINITE mode repeatedly encountering transient errors will terminate with an exception after a duration.`, func() {
		for _, error := range transientErrors {
			e := gmeasure.NewExperiment("measure runner duration")
			ms := NewMockErrorServer(error)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			tolerance := 100 * time.Millisecond
			err = follower.SetErrorTolerance(tolerance)
			Expect(err).ShouldNot(HaveOccurred())

			changesCh, err := follower.StartOneOff()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(changesCh).ToNot(BeNil())

			e.MeasureDuration("stop after", func() {
				ci := <-changesCh
				item, err := ci.Item()
				Expect(item).To(Equal(cloudantv1.ChangesResultItem{}))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp(ErrorText(error)))
				Expect(errors.As(err, &expectedErrType)).To(BeTrue())
			}, gmeasure.Precision(p))
			runDuration := e.Get("stop after").Durations[0]
			Expect(runDuration).To(BeNumerically(">=", tolerance))

			ms.Stop()
		}
	})

	It(`Checks that a FINITE mode repeatedly encountering transient errors will complete successfully if not exceeding the duration.`, func() {
		batches := 5
		ms := NewMockServer(batches, transientErrors)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		err = follower.SetErrorTolerance(100 * time.Millisecond)
		Expect(err).ShouldNot(HaveOccurred())

		changesCh, err := follower.StartOneOff()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(changesCh).ToNot(BeNil())

		count := 0
		for ci := range changesCh {
			item, err := ci.Item()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).ToNot(Equal(cloudantv1.ChangesResultItem{}))
			count += 1
		}
		Expect(count).To(Equal(batches * BatchSize))
	})

	It(`Checks that a FINITE mode repeatedly encountering transient errors will keep trying indefinitely with max suppression.`, func() {
		for _, error := range transientErrors {
			ms := NewMockErrorServer(error)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			cfg := runnerConfig{
				mode:    Finite,
				timeout: 500 * time.Millisecond,
			}
			count, err := runner(follower, cfg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(Equal(0))

			ms.Stop()
		}
	})

	It(`Checks that a FINITE mode encountering transient errors will complete successfully with max suppression.`, func() {
		batches := 4
		ms := NewMockServer(batches, transientErrors)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		changesCh, err := follower.StartOneOff()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(changesCh).ToNot(BeNil())

		count := 0
		for ci := range changesCh {
			item, err := ci.Item()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(item).ToNot(Equal(cloudantv1.ChangesResultItem{}))
			count += 1
		}
		Expect(count).To(Equal(batches * BatchSize))
	})

	It(`Checks calling stop for the FINITE case.`, func() {
		e := gmeasure.NewExperiment("measure runner duration")
		ms := NewMockServer(maxBatches, noErrors)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		timeout := 5 * time.Second
		e.MeasureDuration("stop after", func() {
			cfg := runnerConfig{
				mode:      Finite,
				timeout:   timeout,
				stopAfter: 1000,
			}
			count, err := runner(follower, cfg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(BeNumerically(">", 1000))
		}, gmeasure.Precision(p))
		runDuration := e.Get("stop after").Durations[0]
		Expect(runDuration).To(BeNumerically("<", timeout))
	})

	It(`Checks that a FINITE follower returns error on terminal error at start.`, func() {
		for _, error := range terminalErrors {
			ms := NewMockServer(1, noErrors)
			ms.WithDbInfoError(error)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			postChangesOptions.SetIncludeDocs(true)
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			changesCh, err := follower.StartOneOff()
			Expect(changesCh).To(BeNil())
			Expect(err.Error()).To(MatchRegexp(ErrorText(error)))
			Expect(errors.As(err, &expectedErrType)).To(BeTrue())

			ms.Stop()
		}
	})

	It(`Checks that a FINITE follower can only be started once.`, func() {
		ms := NewMockServer(maxBatches, noErrors)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		cfg := runnerConfig{
			mode:      Finite,
			stopAfter: 1000,
		}
		count, err := runner(follower, cfg)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(count).To(Equal(0))

		errMsg := "cannot start a feed that has already started"
		changesCh, err := follower.StartOneOff()
		Expect(changesCh).To(BeNil())
		Expect(err.Error()).To(Equal(errMsg))
		Expect(errors.As(err, &expectedErrType)).To(BeTrue())

		changesCh, err = follower.Start()
		Expect(changesCh).To(BeNil())
		Expect(err.Error()).To(Equal(errMsg))
		Expect(errors.As(err, &expectedErrType)).To(BeTrue())
	})

	It(`Checks that setting a limit terminates the stream early for FINITE mode and limits smaller, the same and larger than the default batch size.`, func() {
		for _, limit := range limits {
			ms := NewMockServer(maxBatches, noErrors)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			postChangesOptions.SetLimit(int64(limit))
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			cfg := runnerConfig{
				mode:    Finite,
				timeout: time.Second,
			}
			count, err := runner(follower, cfg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(Equal(limit))

			ms.Stop()
		}
	})

	/*
	   For a time frame in 600ms an exponential backoff would make 3 retry
	   attempts (first immediately, for duration of 100ms, second after
	   that for duration of 200ms, and third after 100ms+200ms for duration
	   of 400ms).

	   In the same time frame a full jitter backoff would make more attempts
	   because of its random delay, realistically we can expect ~4-5.
	   We can safely triple this number, check for no more than 15 calls
	   and still be sure that we have delay working, because without it
	   we are looking at +1000 calls in the same time frame.
	*/
	It(`Checks that a FINITE follower delays between retries.`, func() {
		error := transientErrors[0]
		ms := NewMockErrorServer(error)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		cfg := runnerConfig{
			mode:    Finite,
			timeout: 600 * time.Millisecond,
		}
		count, err := runner(follower, cfg)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(count).To(Equal(0))
		Expect(ms.CallNumber()).To(BeNumerically("<=", 15))
	})

	/*
		Mocks a DB of 500000 docs of 523 bytes each to give an expected batch
		size of 5125

		523 bytes + 500 bytes of changes overhead = 1023 bytes
		5 MiB / 1023 bytes = 5125 docs per batch
	*/
	It(`Checks that setting includeDocs forces a calculation of batch size and asserts the size.`, func() {
		batches := 1
		ms := NewMockServer(batches, noErrors)
		ms.WithDbInfo(500000, 523)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		postChangesOptions.SetIncludeDocs(true)

		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		err = follower.SetErrorTolerance(0)
		Expect(err).ShouldNot(HaveOccurred())

		changesCh, err := follower.StartOneOff()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(changesCh).ToNot(BeNil())

		ci := <-changesCh
		item, err := ci.Item()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(item).ToNot(Equal(cloudantv1.ChangesResultItem{}))

		Expect(ms.Limit()).To(BeNumerically("==", 5125))
	})

	/*
	   Mocks a DB of 1 docs of less than 5 MiB size to give an expected batch
	   size of 0

	   Checks that the minimum batch_size of 1 is set.
	*/
	It(`Checks that setting includeDocs forces a calculation of batch size and asserts the size.`, func() {
		batches := 1
		ms := NewMockServer(batches, noErrors)
		ms.WithDbInfo(1, (5*1024*1024 - 1))
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		postChangesOptions.SetIncludeDocs(true)

		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		err = follower.SetErrorTolerance(0)
		Expect(err).ShouldNot(HaveOccurred())

		changesCh, err := follower.StartOneOff()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(changesCh).ToNot(BeNil())

		ci := <-changesCh
		item, err := ci.Item()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(item).ToNot(Equal(cloudantv1.ChangesResultItem{}))

		Expect(ms.Limit()).To(Equal(1))
	})

	It(`Checks that setting includeDocs and limit that below calculated batch sets batch size to limit.`, func() {
		batches := 1
		ms := NewMockServer(batches, noErrors)
		ms.WithDbInfo(500000, 523)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		postChangesOptions.SetLimit(1000)
		postChangesOptions.SetIncludeDocs(true)

		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		err = follower.SetErrorTolerance(0)
		Expect(err).ShouldNot(HaveOccurred())

		changesCh, err := follower.StartOneOff()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(changesCh).ToNot(BeNil())

		ci := <-changesCh
		item, err := ci.Item()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(item).ToNot(Equal(cloudantv1.ChangesResultItem{}))

		Expect(ms.Limit()).To(Equal(1000))
	})
})

var _ = Describe(`ChangesFollower listen`, func() {
	It(`Checks that a LISTEN mode completes successfully (after stopping) with some batches.`, func() {
		ms := NewMockServer(maxBatches, noErrors)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		cfg := runnerConfig{
			mode:    Listen,
			timeout: 5 * time.Second,
		}
		count, err := runner(follower, cfg)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(count).To(BeNumerically(">", 2*BatchSize+1))
	})

	It(`Checks that a LISTEN mode errors for all terminal errors.`, func() {
		for _, error := range terminalErrors {
			ms := NewMockErrorServer(error)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			cfg := runnerConfig{
				mode:    Listen,
				timeout: time.Second,
			}
			_, err = runner(follower, cfg)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(MatchRegexp(ErrorText(error)))
			Expect(errors.As(err, &expectedErrType)).To(BeTrue())

			ms.Stop()
		}
	})

	It(`Checks that a LISTEN mode errors for all transient errors when not suppressing.`, func() {
		for _, error := range transientErrors {
			ms := NewMockErrorServer(error)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			err = follower.SetErrorTolerance(0)
			Expect(err).ShouldNot(HaveOccurred())

			cfg := runnerConfig{
				mode:    Listen,
				timeout: time.Second,
			}
			_, err = runner(follower, cfg)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(MatchRegexp(ErrorText(error)))
			Expect(errors.As(err, &expectedErrType)).To(BeTrue())

			ms.Stop()
		}
	})

	It(`Checks that a LISTEN mode errors for all transient errors
        when exceeding the suppression duration.`, func() {
		for _, error := range transientErrors {
			ms := NewMockErrorServer(error)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			err = follower.SetErrorTolerance(100 * time.Millisecond)
			Expect(err).ShouldNot(HaveOccurred())

			cfg := runnerConfig{
				mode:    Listen,
				timeout: time.Second,
			}
			_, err = runner(follower, cfg)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(MatchRegexp(ErrorText(error)))
			Expect(errors.As(err, &expectedErrType)).To(BeTrue())

			ms.Stop()
		}
	})

	It(`Checks that a LISTEN mode gets changes and can be stopped cleanly with transient errors when not exceeding the suppression duration.`, func() {
		batches := 2
		ms := NewMockServer(batches, transientErrors)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		err = follower.SetErrorTolerance(300 * time.Second)
		Expect(err).ShouldNot(HaveOccurred())

		cfg := runnerConfig{
			mode:    Listen,
			timeout: time.Second,
		}
		count, err := runner(follower, cfg)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(count).To(Equal(batches * BatchSize))
	})

	It(`Checks that a LISTEN mode keeps running with transient errors (until stopped cleanly) with max suppression.`, func() {
		for _, error := range transientErrors {
			ms := NewMockErrorServer(error)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			cfg := runnerConfig{
				mode:    Listen,
				timeout: time.Second,
			}
			count, err := runner(follower, cfg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(Equal(0))
			Expect(ms.CallNumber()).To(BeNumerically(">", 1))

			ms.Stop()
		}
	})

	It(`Checks that a LISTEN mode runs through transient errors with max suppression to receive changes until stopped.`, func() {
		batches := 3
		ms := NewMockServer(batches, transientErrors)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		cfg := runnerConfig{
			mode:    Listen,
			timeout: time.Second,
		}
		count, err := runner(follower, cfg)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(count).To(Equal(batches * BatchSize))
	})

	It(`Checks calling stop for the LISTEN case.`, func() {
		e := gmeasure.NewExperiment("measure runner duration")
		ms := NewMockServer(maxBatches, noErrors)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		timeout := 5 * time.Second
		e.MeasureDuration("stop after", func() {
			cfg := runnerConfig{
				mode:      Listen,
				timeout:   timeout,
				stopAfter: 1000,
			}
			count, err := runner(follower, cfg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(BeNumerically(">", 1000))
		}, gmeasure.Precision(100*time.Millisecond))
		runDuration := e.Get("stop after").Durations[0]
		Expect(runDuration).To(BeNumerically("<", timeout))
	})

	It(`Checks that a LISTEN follower returns error on terminal error at start.`, func() {
		for _, error := range terminalErrors {
			ms := NewMockServer(1, noErrors)
			ms.WithDbInfoError(error)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			postChangesOptions.SetIncludeDocs(true)
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			changesCh, err := follower.Start()
			Expect(changesCh).To(BeNil())
			Expect(err.Error()).To(MatchRegexp(ErrorText(error)))
			Expect(errors.As(err, &expectedErrType)).To(BeTrue())

			ms.Stop()
		}
	})

	It(`Checks that a LISTEN follower can only be started once.`, func() {
		ms := NewMockServer(maxBatches, noErrors)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		cfg := runnerConfig{
			mode:      Listen,
			stopAfter: 1000,
		}
		count, err := runner(follower, cfg)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(count).To(Equal(0))

		errMsg := "cannot start a feed that has already started"
		changesCh, err := follower.StartOneOff()
		Expect(changesCh).To(BeNil())
		Expect(err.Error()).To(Equal(errMsg))
		Expect(errors.As(err, &expectedErrType)).To(BeTrue())

		changesCh, err = follower.Start()
		Expect(changesCh).To(BeNil())
		Expect(err.Error()).To(Equal(errMsg))
		Expect(errors.As(err, &expectedErrType)).To(BeTrue())
	})

	It(`Checks that setting a limit terminates the stream early for LISTEN mode and limits smaller, the same and larger than the default batch size.`, func() {
		for _, limit := range limits {
			ms := NewMockServer(maxBatches, noErrors)
			service := ms.Start()

			postChangesOptions := service.NewPostChangesOptions("db")
			postChangesOptions.SetLimit(int64(limit))
			follower, err := NewChangesFollower(service, postChangesOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			cfg := runnerConfig{
				mode:    Listen,
				timeout: time.Second,
			}
			count, err := runner(follower, cfg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(Equal(limit))

			ms.Stop()
		}
	})

	// See the FINITE version of the test for additional comments.
	It(`Checks that a LISTEN follower delays between retries.`, func() {
		error := transientErrors[0]
		ms := NewMockErrorServer(error)
		service := ms.Start()
		defer ms.Stop()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		cfg := runnerConfig{
			mode:    Listen,
			timeout: 600 * time.Millisecond,
		}
		count, err := runner(follower, cfg)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(count).To(Equal(0))
		Expect(ms.CallNumber()).To(BeNumerically("<=", 15))
	})
})

var _ = Describe(`ChangesFollower with context`, func() {
	var (
		runnerTimeout = time.Second
		p             = 200 * time.Millisecond
	)

	It(`Checks passing context with timeout.`, func() {
		e := gmeasure.NewExperiment("measure runner duration")
		ms := NewMockServer(maxBatches, noErrors)
		service := ms.Start()
		defer ms.Stop()

		timeout := 500 * time.Millisecond
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollowerWithContext(ctx, service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		e.MeasureDuration("stop after", func() {
			cfg := runnerConfig{
				mode:    Listen,
				timeout: runnerTimeout,
			}
			count, err := runner(follower, cfg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(BeNumerically(">", 0))
		}, gmeasure.Precision(p))
		runDuration := e.Get("stop after").Durations[0]
		Expect(runDuration).To(BeNumerically("~", timeout, p))
		Expect(runDuration).To(BeNumerically("<", runnerTimeout))
		Expect(ctx.Err()).To(Equal(context.DeadlineExceeded))
	})

	It(`Checks passing context with deadline.`, func() {
		e := gmeasure.NewExperiment("measure runner duration")
		ms := NewMockServer(maxBatches, noErrors)
		service := ms.Start()
		defer ms.Stop()

		duration := 500 * time.Millisecond
		deadline := time.Now().Add(duration)
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		defer cancel()

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollowerWithContext(ctx, service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		e.MeasureDuration("stop after", func() {
			cfg := runnerConfig{
				mode:    Listen,
				timeout: runnerTimeout,
			}
			count, err := runner(follower, cfg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(BeNumerically(">", 0))
		}, gmeasure.Precision(p))
		runDuration := e.Get("stop after").Durations[0]
		Expect(runDuration).To(BeNumerically("~", duration, p))
		Expect(runDuration).To(BeNumerically("<", runnerTimeout))
		Expect(ctx.Err()).To(Equal(context.DeadlineExceeded))
	})

	It(`Checks passing context with cancel.`, func() {
		e := gmeasure.NewExperiment("measure runner duration")
		ms := NewMockServer(maxBatches, noErrors)
		service := ms.Start()
		defer ms.Stop()

		timeout := 500 * time.Millisecond
		ctx, cancel := context.WithCancel(context.Background())
		time.AfterFunc(timeout, cancel)

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollowerWithContext(ctx, service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		e.MeasureDuration("stop after", func() {
			cfg := runnerConfig{
				mode:    Listen,
				timeout: runnerTimeout,
			}
			count, err := runner(follower, cfg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(BeNumerically(">", 0))
		}, gmeasure.Precision(p))
		runDuration := e.Get("stop after").Durations[0]
		Expect(runDuration).To(BeNumerically("~", timeout, p))
		Expect(runDuration).To(BeNumerically("<", runnerTimeout))
		Expect(ctx.Err()).To(Equal(context.Canceled))
	})

	It(`Checks passing context with timeout on empty changes feed.`, func() {
		ms := NewMockServer(0, noErrors)
		service := ms.Start()
		defer ms.Stop()

		timeout := 100 * time.Millisecond
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		time.AfterFunc(timeout, cancel)

		postChangesOptions := service.NewPostChangesOptions("db")
		follower, err := NewChangesFollowerWithContext(ctx, service, postChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(follower).ToNot(BeNil())

		changesCh, err := follower.Start()
		Expect(err).ShouldNot(HaveOccurred())
		time.Sleep(2 * timeout)
		ci := <-changesCh
		Expect(changesCh).Should(BeClosed())
		item, err := ci.Item()
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).To(Equal("can't read from a closed channel"))
		Expect(errors.As(err, &expectedErrType)).To(BeTrue())
		Expect(item).To(Equal(cloudantv1.ChangesResultItem{}))
	})
})

// ---------------------------------------------------------------------------
// seqMarkers helpers
// ---------------------------------------------------------------------------

// seqStr builds a seq string from an integer, e.g. seqStr(11) -> "11-aa".
func seqStr(n int) string {
	return fmt.Sprintf("%d-aa", n)
}

// makeTestRow builds a ChangesResultItem with the given seq (nil for null rows).
func makeTestRow(s *string) cloudantv1.ChangesResultItem {
	return cloudantv1.ChangesResultItem{
		ID:      core.StringPtr("doc"),
		Changes: []cloudantv1.Change{},
		Seq:     s,
	}
}

// testPageData holds the raw page fields used to populate seqMarkers.
type testPageData struct {
	results []cloudantv1.ChangesResultItem
	lastSeq string
}

// testPageType is the factory for the 9 page types.
//
//	Type 1: rows=[b, b+1],     lastSeq=b+1  (last row == last_seq, no nulls)
//	Type 2: rows=[b, b+1],     lastSeq=b+2  (last row != last_seq, no nulls)
//	Type 3: rows=[null, b+1],  lastSeq=b+1  (leading null, last row == last_seq)
//	Type 4: rows=[null, b+1],  lastSeq=b+2  (leading null, last row != last_seq)
//	Type 5: rows=[b, null],    lastSeq=b+1  (trailing null last row)
//	Type 6: rows=[b, null],    lastSeq=b+2  (trailing null last row, last_seq beyond)
//	Type 7: rows=[null, null], lastSeq=b+1  (all nulls)
//	Type 8: rows=[null, null], lastSeq=b+2  (all nulls, last_seq beyond)
//	Type 9: rows=[],           lastSeq=b    (empty page)
//
// What gets stored in seqMarkers per type:
//
//	Types 1,3: ROW('(b+1)-aa'), PAGE('(b+1)-aa')
//	Types 2,4: ROW('(b+1)-aa'), PAGE('(b+2)-aa')
//	Types 5,7: ROW(nil),        PAGE('(b+1)-aa')
//	Types 6,8: ROW(nil),        PAGE('(b+2)-aa')
//	Type  9:   PAGE('b-aa')  (no ROW)
func testPageType(t, base int) testPageData {
	s := func(n int) *string { v := seqStr(n); return &v }
	switch t {
	case 1:
		return testPageData{results: []cloudantv1.ChangesResultItem{makeTestRow(s(base)), makeTestRow(s(base + 1))}, lastSeq: seqStr(base + 1)}
	case 2:
		return testPageData{results: []cloudantv1.ChangesResultItem{makeTestRow(s(base)), makeTestRow(s(base + 1))}, lastSeq: seqStr(base + 2)}
	case 3:
		return testPageData{results: []cloudantv1.ChangesResultItem{makeTestRow(nil), makeTestRow(s(base + 1))}, lastSeq: seqStr(base + 1)}
	case 4:
		return testPageData{results: []cloudantv1.ChangesResultItem{makeTestRow(nil), makeTestRow(s(base + 1))}, lastSeq: seqStr(base + 2)}
	case 5:
		return testPageData{results: []cloudantv1.ChangesResultItem{makeTestRow(s(base)), makeTestRow(nil)}, lastSeq: seqStr(base + 1)}
	case 6:
		return testPageData{results: []cloudantv1.ChangesResultItem{makeTestRow(s(base)), makeTestRow(nil)}, lastSeq: seqStr(base + 2)}
	case 7:
		return testPageData{results: []cloudantv1.ChangesResultItem{makeTestRow(nil), makeTestRow(nil)}, lastSeq: seqStr(base + 1)}
	case 8:
		return testPageData{results: []cloudantv1.ChangesResultItem{makeTestRow(nil), makeTestRow(nil)}, lastSeq: seqStr(base + 2)}
	case 9:
		return testPageData{results: []cloudantv1.ChangesResultItem{}, lastSeq: seqStr(base)}
	default:
		panic(fmt.Sprintf("unknown page type: %d", t))
	}
}

// populateTestFollower builds a ChangesFollower and fills its seqMarkers
// by calling the real updateSeqMarkers method for each page.
func populateTestFollower(pages []testPageData) *ChangesFollower {
	cf := &ChangesFollower{}
	for _, p := range pages {
		ls := p.lastSeq
		cf.updateSeqMarkers(p.results, &ls)
	}
	return cf
}

// lastSeqSinceHelper populates a follower with pages and calls lastSeqSince.
func lastSeqSinceHelper(pages []testPageData, querySeq string) string {
	return populateTestFollower(pages).lastSeqSince(querySeq)
}

// ---------------------------------------------------------------------------
// seqMarkers / lastSeqSince tests
// ---------------------------------------------------------------------------

var _ = Describe(`seqMarkers / lastSeqSince`, func() {

	// -----------------------------------------------------------------------
	// Not-found / empty edge cases
	// -----------------------------------------------------------------------

	It(`testLastSeqSinceNotFound`, func() {
		result := lastSeqSinceHelper([]testPageData{testPageType(1, 10)}, "999-ff")
		Expect(result).To(Equal("999-ff"))
	})

	It(`testLastSeqSinceEmptySeqMarkers`, func() {
		result := lastSeqSinceHelper([]testPageData{}, "1-aa")
		Expect(result).To(Equal("1-aa"))
	})

	// -----------------------------------------------------------------------
	// Per-page-type: single page
	// -----------------------------------------------------------------------

	DescribeTable(`testLastSeqSinceAlone`,
		func(pageT, base int, querySeq, expected string) {
			result := lastSeqSinceHelper([]testPageData{testPageType(pageT, base)}, querySeq)
			Expect(result).To(Equal(expected))
		},
		Entry("Type 1: last row seq (== last_seq)", 1, 10, seqStr(11), seqStr(11)),
		Entry("Type 3: last row seq (== last_seq)", 3, 10, seqStr(11), seqStr(11)),
		Entry("Type 2: last row seq -> last_seq", 2, 10, seqStr(11), seqStr(12)),
		Entry("Type 2: last_seq key -> itself", 2, 10, seqStr(12), seqStr(12)),
		Entry("Type 4: last row seq -> last_seq", 4, 10, seqStr(11), seqStr(12)),
		Entry("Type 4: last_seq key -> itself", 4, 10, seqStr(12), seqStr(12)),
		Entry("Type 5: non-stored row seq unchanged", 5, 10, seqStr(10), seqStr(10)),
		Entry("Type 5: last_seq key -> itself", 5, 10, seqStr(11), seqStr(11)),
		Entry("Type 6: non-stored row seq unchanged", 6, 10, seqStr(10), seqStr(10)),
		Entry("Type 6: last_seq key -> itself", 6, 10, seqStr(12), seqStr(12)),
		Entry("Type 7: last_seq key -> itself", 7, 10, seqStr(11), seqStr(11)),
		Entry("Type 8: last_seq key -> itself", 8, 10, seqStr(12), seqStr(12)),
		Entry("Type 9: last_seq key -> itself", 9, 10, seqStr(10), seqStr(10)),
	)

	// -----------------------------------------------------------------------
	// Per-page-type: followed by a non-empty page (type 1 at base 20)
	// Page 2 inserts ROW('21-aa') which blocks advancement.
	// -----------------------------------------------------------------------

	DescribeTable(`testLastSeqSinceFollowedByNonEmpty`,
		func(pageT, base int, querySeq, expected string) {
			result := lastSeqSinceHelper([]testPageData{testPageType(pageT, base), testPageType(1, 20)}, querySeq)
			Expect(result).To(Equal(expected))
		},
		Entry("Type 1 + non-empty: blocked by p2 ROW", 1, 10, seqStr(11), seqStr(11)),
		Entry("Type 2 + non-empty: last row seq -> p1 last_seq", 2, 10, seqStr(11), seqStr(12)),
		Entry("Type 2 + non-empty: last_seq key -> p1 last_seq", 2, 10, seqStr(12), seqStr(12)),
		Entry("Type 3 + non-empty: blocked by p2 ROW", 3, 10, seqStr(11), seqStr(11)),
		Entry("Type 4 + non-empty: last row seq -> p1 last_seq", 4, 10, seqStr(11), seqStr(12)),
		Entry("Type 4 + non-empty: last_seq key -> p1 last_seq", 4, 10, seqStr(12), seqStr(12)),
		Entry("Type 5 + non-empty: blocked by p2 ROW", 5, 10, seqStr(11), seqStr(11)),
		Entry("Type 6 + non-empty: blocked by p2 ROW", 6, 10, seqStr(12), seqStr(12)),
		Entry("Type 7 + non-empty: blocked by p2 ROW", 7, 10, seqStr(11), seqStr(11)),
		Entry("Type 8 + non-empty: blocked by p2 ROW", 8, 10, seqStr(12), seqStr(12)),
		Entry("Type 9 + non-empty: blocked by p2 ROW", 9, 10, seqStr(10), seqStr(10)),
	)

	// -----------------------------------------------------------------------
	// Per-page-type: followed by an empty page (type 9 at base 20)
	// Page 2 inserts only PAGE('20-aa') — no ROW to block, advances to '20-aa'.
	// -----------------------------------------------------------------------

	DescribeTable(`testLastSeqSinceFollowedByEmpty`,
		func(pageT, base int, querySeq, expected string) {
			result := lastSeqSinceHelper([]testPageData{testPageType(pageT, base), testPageType(9, 20)}, querySeq)
			Expect(result).To(Equal(expected))
		},
		Entry("Type 1 + empty: advances to p2 last_seq", 1, 10, seqStr(11), seqStr(20)),
		Entry("Type 2 + empty: last row seq advances to p2", 2, 10, seqStr(11), seqStr(20)),
		Entry("Type 2 + empty: last_seq key advances to p2", 2, 10, seqStr(12), seqStr(20)),
		Entry("Type 3 + empty: advances to p2 last_seq", 3, 10, seqStr(11), seqStr(20)),
		Entry("Type 4 + empty: last row seq advances to p2", 4, 10, seqStr(11), seqStr(20)),
		Entry("Type 4 + empty: last_seq key advances to p2", 4, 10, seqStr(12), seqStr(20)),
		Entry("Type 5 + empty: last_seq advances to p2", 5, 10, seqStr(11), seqStr(20)),
		Entry("Type 6 + empty: last_seq advances to p2", 6, 10, seqStr(12), seqStr(20)),
		Entry("Type 7 + empty: last_seq advances to p2", 7, 10, seqStr(11), seqStr(20)),
		Entry("Type 8 + empty: last_seq advances to p2", 8, 10, seqStr(12), seqStr(20)),
		Entry("Type 9 + empty: advances to p2 last_seq", 9, 10, seqStr(10), seqStr(20)),
	)

	// -----------------------------------------------------------------------
	// All 8 three-page sequences of empty (E=type 9) and non-empty (N=type 1).
	// Query from page 1's last_seq key. E adds only PAGE; N adds ROW+PAGE.
	// -----------------------------------------------------------------------

	DescribeTable(`testLastSeqSince3PageSequence`,
		func(types []int, bases []int, querySeq, expected string) {
			pages := make([]testPageData, len(types))
			for i, t := range types {
				pages[i] = testPageType(t, bases[i])
			}
			result := lastSeqSinceHelper(pages, querySeq)
			Expect(result).To(Equal(expected))
		},
		Entry("NNN: blocked by p2 ROW -> p1 last_seq",
			[]int{1, 1, 1}, []int{10, 20, 30}, seqStr(11), seqStr(11)),
		Entry("NNE: blocked by p2 ROW -> p1 last_seq",
			[]int{1, 1, 9}, []int{10, 20, 30}, seqStr(11), seqStr(11)),
		Entry("NEE: advances through both empty pages",
			[]int{1, 9, 9}, []int{10, 20, 30}, seqStr(11), seqStr(30)),
		Entry("NEN: advances through p2 empty, stops at p3 ROW",
			[]int{1, 9, 1}, []int{10, 20, 30}, seqStr(11), seqStr(20)),
		Entry("ENN: blocked by p2 ROW -> p1 last_seq",
			[]int{9, 1, 1}, []int{10, 20, 30}, seqStr(10), seqStr(10)),
		Entry("ENE: blocked by p2 ROW -> p1 last_seq",
			[]int{9, 1, 9}, []int{10, 20, 30}, seqStr(10), seqStr(10)),
		Entry("EEN: advances through p2, stops at p3 ROW",
			[]int{9, 9, 1}, []int{10, 20, 30}, seqStr(10), seqStr(20)),
		Entry("EEE: advances through all three empty pages",
			[]int{9, 9, 9}, []int{10, 20, 30}, seqStr(10), seqStr(30)),
	)

	// -----------------------------------------------------------------------
	// Eviction
	//
	// Each non-empty page (type 2) adds 2 entries (ROW + PAGE). With
	// CAPACITY=200 and EVICTION_COUNT=20, adding 101 pages triggers one
	// eviction of the oldest 20 entries (first 10 pages). Entries for
	// page 0 (base=0) should be gone; most recent should remain.
	// -----------------------------------------------------------------------

	It(`testLastSeqSinceEviction`, func() {
		pages := make([]testPageData, 101)
		for i := range pages {
			pages[i] = testPageType(2, i*10)
		}
		cf := populateTestFollower(pages)

		// Page 0 (base=0): row=seqStr(1), page=seqStr(2) — evicted, returns input unchanged
		Expect(cf.lastSeqSince(seqStr(1))).To(Equal(seqStr(1)))
		Expect(cf.lastSeqSince(seqStr(2))).To(Equal(seqStr(2)))

		// Most recent page (base=1000): row=seqStr(1001), page=seqStr(1002) — still present
		Expect(cf.lastSeqSince(seqStr(1001))).To(Equal(seqStr(1002)))
		Expect(cf.lastSeqSince(seqStr(1002))).To(Equal(seqStr(1002)))
	})

	// -----------------------------------------------------------------------
	// Nil seq row (seq_interval scenario) — would panic without nil guard
	// -----------------------------------------------------------------------

	It(`testLastSeqSinceNilRowSeqDoesNotPanic`, func() {
		// Types 5/6/7/8 store ROW(nil). Querying a seq not in the markers
		// must return input unchanged without panicking.
		pages := []testPageData{testPageType(5, 10)}
		Expect(func() {
			result := lastSeqSinceHelper(pages, seqStr(10))
			Expect(result).To(Equal(seqStr(10)))
		}).NotTo(Panic())
	})
})

// ---------------------------------------------------------------------------
// GetLastSeqNewerThan tests
// ---------------------------------------------------------------------------

var _ = Describe(`GetLastSeqNewerThan`, func() {
	var (
		glsnService            *cloudantv1.CloudantV1
		glsnPostChangesOptions *cloudantv1.PostChangesOptions
	)

	BeforeEach(func() {
		var serviceErr error
		glsnService, serviceErr = cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
			URL:           "http://localhost:5984",
			Authenticator: &core.NoAuthAuthenticator{},
		})
		Expect(serviceErr).ShouldNot(HaveOccurred())
		glsnPostChangesOptions = glsnService.NewPostChangesOptions("db")
	})

	It(`testGetLastSeqNewerThanWithEmptyString`, func() {
		follower, err := NewChangesFollower(glsnService, glsnPostChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())

		_, err = follower.GetLastSeqNewerThan("")
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("the provided sequence ID cannot be null or empty"))
		Expect(errors.As(err, &expectedErrType)).To(BeTrue())
	})

	It(`testGetLastSeqNewerThanBeforeFeedStarts`, func() {
		follower, err := NewChangesFollower(glsnService, glsnPostChangesOptions)
		Expect(err).ShouldNot(HaveOccurred())

		result, err := follower.GetLastSeqNewerThan("seq-a")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(result).To(Equal("seq-a"))
	})

	It(`testGetLastSeqNewerThanUnknownSeq`, func() {
		ms := NewMockServer(1, noErrors)
		svc := ms.Start()
		defer ms.Stop()

		opts := svc.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(svc, opts)
		Expect(err).ShouldNot(HaveOccurred())

		ch, err := follower.StartOneOff()
		Expect(err).ShouldNot(HaveOccurred())
		for range ch {
		}

		result, err := follower.GetLastSeqNewerThan("seq-unknown")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(result).To(Equal("seq-unknown"))
	})

	It(`testGetLastSeqNewerThanMiddleOfBatch`, func() {
		ms := NewMockServer(1, noErrors)
		svc := ms.Start()
		defer ms.Stop()

		opts := svc.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(svc, opts)
		Expect(err).ShouldNot(HaveOccurred())

		ch, err := follower.StartOneOff()
		Expect(err).ShouldNot(HaveOccurred())

		var items []cloudantv1.ChangesResultItem
		for ci := range ch {
			item, itemErr := ci.Item()
			Expect(itemErr).ShouldNot(HaveOccurred())
			items = append(items, item)
		}
		Expect(len(items)).To(BeNumerically(">", 2))

		// Only the last item's seq is stored — middle items are not in seqMarkers.
		seqA := *items[0].Seq
		seqB := *items[1].Seq

		resultA, err := follower.GetLastSeqNewerThan(seqA)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resultA).To(Equal(seqA))

		resultB, err := follower.GetLastSeqNewerThan(seqB)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resultB).To(Equal(seqB))
	})

	It(`testGetLastSeqNewerThanEndToEnd`, func() {
		ms := NewMockServer(1, noErrors)
		svc := ms.Start()
		defer ms.Stop()

		opts := svc.NewPostChangesOptions("db")
		follower, err := NewChangesFollower(svc, opts)
		Expect(err).ShouldNot(HaveOccurred())

		ch, err := follower.StartOneOff()
		Expect(err).ShouldNot(HaveOccurred())

		var lastItem cloudantv1.ChangesResultItem
		for ci := range ch {
			item, itemErr := ci.Item()
			Expect(itemErr).ShouldNot(HaveOccurred())
			lastItem = item
		}

		// The last item's seq is the stored row entry. MockChangesGenerator
		// produces pages where last row seq == last_seq, so result equals input.
		result, err := follower.GetLastSeqNewerThan(*lastItem.Seq)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(result).To(Equal(*lastItem.Seq))
	})
})
