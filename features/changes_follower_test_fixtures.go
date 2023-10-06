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
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	maxBatches       = math.MaxInt / BatchSize
	StatusBrokenJson = 600
	StatusBadIO      = 601
)

var noErrors []int = []int{}

var transientErrors []int = []int{
	http.StatusTooManyRequests,
	http.StatusInternalServerError,
	http.StatusBadGateway,
	http.StatusGatewayTimeout,
	StatusBrokenJson,
	StatusBadIO,
}

var terminalErrors []int = []int{
	http.StatusBadRequest,
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusNotFound,
}

var limits []int = []int{
	100,
	BatchSize,
	BatchSize + 123,
}

type MockGenerator interface {
	Next() (statusCode int, data []byte)
}

type MockChangesGenerator struct {
	batches     int
	errs        []int
	batchNum    int
	returnError bool
}

func NewMockChangesGenerator(batches int, errs []int) *MockChangesGenerator {
	return &MockChangesGenerator{
		batchNum:    1,
		returnError: false,
		batches:     batches,
		errs:        errs,
	}
}

func (mg *MockChangesGenerator) Next() (statusCode int, data []byte) {
	if mg.returnError {
		mg.returnError = false
		switch err := mg.errs[(mg.batchNum)%len(mg.errs)]; err {
		case StatusBrokenJson:
			return http.StatusOK, []byte(`{`)
		default:
			return err, []byte{}
		}
	}
	// this stands for "large" seq in empty result case
	final := mg.batches * BatchSize
	lastSeq := fmt.Sprintf("%d-abcdef", final)
	pending := 0
	results := make([]cloudantv1.ChangesResultItem, 0)
	if mg.batchNum <= mg.batches {
		// we start from doc idx 000001
		start := (mg.batchNum-1)*BatchSize + 1
		stop := start + BatchSize
		lastSeq = fmt.Sprintf("%d-abcdef", stop-1)
		pending = mg.batches*BatchSize - mg.batchNum*BatchSize
		for idx := start; idx < stop; idx++ {
			id := fmt.Sprintf("%06d", idx)
			seq := fmt.Sprintf("%d-abcdef", idx)
			results = append(results, cloudantv1.ChangesResultItem{
				ID:      core.StringPtr(id),
				Changes: make([]cloudantv1.Change, 0),
				Seq:     core.StringPtr(seq),
			})
		}
	}
	mg.batchNum += 1
	if len(mg.errs) > 0 {
		mg.returnError = true
	}
	data, err := json.Marshal(cloudantv1.ChangesResult{
		LastSeq: core.StringPtr(lastSeq),
		Pending: core.Int64Ptr(int64(pending)),
		Results: results,
	})
	Expect(err).ShouldNot(HaveOccurred())

	return http.StatusOK, data
}

type MockErrorGenerator struct {
	err  int
	data []byte
}

func NewMockErrorGenerator(err int) *MockErrorGenerator {
	switch err {
	case StatusBrokenJson:
		return &MockErrorGenerator{err: http.StatusOK, data: []byte(`{`)}
	default:
		return &MockErrorGenerator{err: err, data: []byte{}}
	}
}

func (mg *MockErrorGenerator) Next() (statusCode int, data []byte) {
	return mg.err, mg.data
}

type MockDbInfoGenerator struct {
	resp []byte
}

func NewDbInfoGenerator(docCount, docSize int) *MockDbInfoGenerator {
	data, err := json.Marshal(map[string]interface{}{
		"doc_count": docCount,
		"sizes": map[string]int{
			"external": docCount * docSize,
		},
	})
	Expect(err).ShouldNot(HaveOccurred())

	return &MockDbInfoGenerator{resp: data}
}

func (mg *MockDbInfoGenerator) Next() (statusCode int, data []byte) {
	return http.StatusOK, mg.resp
}

type MockServer struct {
	server        *httptest.Server
	mockGenerator MockGenerator
	dbInfo        MockGenerator
	callNumber    atomic.Int32
	limit         atomic.Int32
}

func NewMockServer(batches int, errs []int) *MockServer {
	mg := NewMockChangesGenerator(batches, errs)
	ms := MockServer{mockGenerator: mg}
	ms.callNumber.Store(0)
	return &ms
}

func NewMockErrorServer(err int) *MockServer {
	mg := NewMockErrorGenerator(err)
	ms := MockServer{mockGenerator: mg}
	ms.callNumber.Store(0)
	ms.limit.Store(0)
	return &ms
}

func (ms *MockServer) WithDbInfo(docCount, docSize int) *MockServer {
	ms.dbInfo = NewDbInfoGenerator(docCount, docSize)
	return ms
}

func (ms *MockServer) WithDbInfoError(err int) *MockServer {
	ms.dbInfo = NewMockErrorGenerator(err)
	return ms
}

func (ms *MockServer) Start() *cloudantv1.CloudantV1 {
	ms.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer GinkgoRecover()

		var statusCode int
		var data []byte
		switch path := r.URL.EscapedPath(); path {
		case "/db":
			Expect(r.Method).To(Equal("GET"))
			statusCode, data = ms.dbInfo.Next()
		default:
			postChangesPath := "/db/_changes"
			Expect(path).To(Equal(postChangesPath))
			Expect(r.Method).To(Equal("POST"))
			statusCode, data = ms.mockGenerator.Next()
			if statusCode == StatusBadIO {
				// return header and shut down the connection to get
				// "An error occurred while reading the response body: EOF"
				ms.callNumber.Add(1)
				w.WriteHeader(statusCode)
				hj, ok := w.(http.Hijacker)
				if !ok {
					http.Error(w, "can't create hijack rw", http.StatusInternalServerError)
					return
				}
				conn, _, err := hj.Hijack()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				err = conn.Close()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}

			l, err := strconv.Atoi(r.URL.Query().Get("limit"))
			Expect(err).ShouldNot(HaveOccurred())
			ms.limit.Store(int32(l))
		}
		// Set mock response
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(statusCode)
		//nolint:errcheck
		w.Write(data)

		ms.callNumber.Add(1)
	}))

	service, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
		URL:           ms.server.URL,
		Authenticator: &core.NoAuthAuthenticator{},
	})
	Expect(serviceErr).ShouldNot(HaveOccurred())
	Expect(service).ToNot(BeNil())

	return service
}

func (ms *MockServer) Limit() int {
	return int(ms.limit.Load())
}

func (ms *MockServer) CallNumber() int {
	return int(ms.callNumber.Load())
}

func (ms *MockServer) Stop() {
	ms.server.Close()
}

func ErrorText(err int) string {
	switch err {
	case StatusBrokenJson:
		return "An error occurred while processing the HTTP response"
	default:
		return http.StatusText(err)
	}
}

type runnerConfig struct {
	mode      Mode
	timeout   time.Duration
	stopAfter int
}

func runner(cf *ChangesFollower, c runnerConfig) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	countCh := make(chan int)
	errors := make(chan error)
	go func() {
		defer GinkgoRecover()
		defer close(countCh)
		defer close(errors)
		count := 0
		var changesCh <-chan ChangesItem
		var err error
		switch c.mode {
		case Finite:
			changesCh, err = cf.StartOneOff()
		case Listen:
			changesCh, err = cf.Start()
		}
		Expect(err).ShouldNot(HaveOccurred())
		Expect(changesCh).ToNot(BeNil())
		for {
			select {
			case <-ctx.Done():
				cf.Stop()
				// check that changesCh was closed
				select {
				case <-changesCh:
				default:
				}
				// return last count
				countCh <- count
				return
			case ci, ok := <-changesCh:
				// Follower quit and closed the changes channel
				if !ok || (c.stopAfter > 0 && count > c.stopAfter) {
					countCh <- count
					return
				}
				item, err := ci.Item()
				if err != nil {
					errors <- err
					return
				}
				Expect(item).ToNot(Equal(cloudantv1.ChangesResultItem{}))
				count += 1
			}
		}
	}()
	select {
	case error := <-errors:
		return 0, error
	case count := <-countCh:
		return count, nil
	}
}
