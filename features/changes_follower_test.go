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
	"fmt"
	"time"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

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

			o := follower.options
			Expect(*o.Feed).Should(Equal(cloudantv1.PostChangesOptionsFeedLongpollConst))
			Expect(*o.Timeout).Should(Equal(LongpollTimeout.Milliseconds()))
		})

		It(`Set defaults with limit`, func() {
			follower, followerErr := NewChangesFollower(service, postChangesOptions)

			Expect(followerErr).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			follower.setOptionsDefaults().withLimit(12)

			o := follower.options
			Expect(*o.Feed).Should(Equal(cloudantv1.PostChangesOptionsFeedLongpollConst))
			Expect(*o.Timeout).Should(Equal(LongpollTimeout.Milliseconds()))
			Expect(*o.Limit).Should(Equal(int64(12)))
		})

		It(`Set defaults with PostChangesOptions limit`, func() {
			postChangesOptions.SetLimit(24)
			follower, followerErr := NewChangesFollower(service, postChangesOptions)

			Expect(followerErr).ShouldNot(HaveOccurred())
			Expect(follower).ToNot(BeNil())

			follower.setOptionsDefaults().withLimit(12)

			o := follower.options
			Expect(*o.Feed).Should(Equal(cloudantv1.PostChangesOptionsFeedLongpollConst))
			Expect(*o.Timeout).Should(Equal(LongpollTimeout.Milliseconds()))
			Expect(*o.Limit).Should(Equal(int64(12)))
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

		changesCh, err = follower.Start()
		Expect(changesCh).To(BeNil())
		Expect(err.Error()).To(Equal(errMsg))
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

		changesCh, err = follower.Start()
		Expect(changesCh).To(BeNil())
		Expect(err.Error()).To(Equal(errMsg))
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
	})
})
