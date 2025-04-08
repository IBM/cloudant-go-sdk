/**
 * © Copyright IBM Corporation 2022, 2024. All Rights Reserved.
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
	"math"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/cloudant-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
)

// Max timedelta is limited at max int64. A compiler will stop attempts to
// exceed it.
const forever time.Duration = math.MaxInt64

// Minimal client timeout set to 1 minute.
const minClientTimeout time.Duration = time.Minute

// LongpollTimeout is set to give the changes request a chance to be answered
// before the client timeout it is set to 3 seconds less.
const LongpollTimeout time.Duration = minClientTimeout - 3*time.Second

// BatchSize is the default number of document to pull in a single request.
const BatchSize int = 10000

// Base delay between unsuccessful attempts to pull changes feed
// in presence of transient errors
const baseDelay time.Duration = 100 * time.Millisecond

// Once we reach this number of retries we'll be capping the backoff
var expRetryGate int = int(math.Log(float64(LongpollTimeout/baseDelay)) / math.Log(2))

// Mode are enums for changes follower's operation mode.
type Mode int

const (
	// Finite mode is an alias for StartOneOff()
	Finite Mode = iota
	// Listen mode is an alias for Start()
	Listen
)

// TransientErrorSuppression are enums for changes follower's
// transient errors suppression mode.
type TransientErrorSuppression int

const (
	// Always suppress transient errors
	Always TransientErrorSuppression = iota
	// Never suppress transient errors
	Never
	// Timer specifies a time duration for suppressing transient errors
	Timer
)

// ChangesFollower is a helper for using the changes feed.
//
// There are two modes of operation:
//
//	StartOneOff() to fetch the changes from the supplied since sequence
//	until there are no further pending changes.
//	Start() to fetch the changes from the supplied since sequence
//	and then continuing to listen indefinitely for further new changes.
//
// The starting sequence ID can be changed for either mode by using "Since".
// By default when using:
//
//	StartOneOff() the feed will start from the beginning.
//	Start() the feed will start from "now".
//
// In either mode the response stream can be terminated early by calling Stop().
// By default ChangesFollower will suppress transient errors indefinitely
// and endeavour to run to completion or listen forever. For applications
// where that behaviour is not desirable an alternate option is
// available where a duration may be specified to limit the time since the
// last successful response that transient errors will be suppressed.
// It should be noted that some errors considered terminal, for example, the
// database not existing or invalid credentials are never suppressed and will
// return an Error immediately.
// The named attributes for "PostChangesOptions" are used to configure the behaviour of the ChangesFollower.
// However, a subset of the attributes are invalid as they are configured
// internally by the implementation and will cause an Error to be returned
// if supplied.
// These invalid options are:
//   - Descending
//   - Feed
//   - Heartbeat
//   - LastEventID
//   - Timeout
//
// Only the value of "_selector" is permitted for
// the PostChangesOptions's "Filter" option.
// Selector based filters perform better than JS based filters and using one
// of the alternative JS based filter types will cause ChangesFollower
// to return an Error.
// It should also be noted that the "Limit" parameter will truncate the
// response at the given number of changes in either operating mode.
// The ChangesFollower requires the Cloudant client to have HTTP timeout
// of at least 1 minute.
// The default client configuration has a sufficiently long timeout.
type ChangesFollower struct {
	client           *cloudantv1.CloudantV1
	options          *cloudantv1.PostChangesOptions
	mode             Mode
	since            string
	limit            int
	retry            int
	errorTolerance   time.Duration
	suppression      TransientErrorSuppression
	successTimestamp time.Time
	ctx              context.Context
	cancel           context.CancelFunc
	running          bool
	runLock          sync.Mutex
	logger           core.Logger
}

// ChangesItem is a wrapper structure around cloudantv1.ChangesResultItem
// with addtitional attribute Error for errors received during the run.
type ChangesItem struct {
	item  cloudantv1.ChangesResultItem
	error error
}

// Item is a ChangesItem's getter for cloudantv1.ChangesResultItem
// that either returns an acquired item or an error received
// during its fetch.
func (ci ChangesItem) Item() (cloudantv1.ChangesResultItem, error) {
	if ci.error == nil && ci.item.ID == nil {
		err := core.SDKErrorf(nil, "can't read from a closed channel", "changes-follower-closed-channel", common.GetComponentInfo())
		return ci.item, err
	}
	return ci.item, ci.error
}

// NewChangesFollower returns a new ChangesFollower or an error if provided
// configuration is invalid.
func NewChangesFollower(c *cloudantv1.CloudantV1, o *cloudantv1.PostChangesOptions) (*ChangesFollower, error) {
	ctx := context.Background()
	return NewChangesFollowerWithContext(ctx, c, o)
}

// NewChangesFollowerWithContext returns a new ChangesFollower initiated
// with a given context or an error if provided configuration is invalid.
func NewChangesFollowerWithContext(ctx context.Context, c *cloudantv1.CloudantV1, o *cloudantv1.PostChangesOptions) (*ChangesFollower, error) {
	err := validateOptions(o)
	if err != nil {
		return nil, err
	}

	client := c.Service.GetHTTPClient()
	if client.Timeout > 0 && client.Timeout < minClientTimeout {
		err := fmt.Errorf("to use ChangesFollower the client timeout must be at least %d ms. The client timeout is %d ms", minClientTimeout/time.Millisecond, client.Timeout/time.Millisecond)
		return nil, core.SDKErrorf(err, "", "changes-follower-invalid-timeout", common.GetComponentInfo())
	}

	cf := &ChangesFollower{
		client:  c,
		options: o,
		limit:   -1,
		// this is default, since we are setting error tolerance separately
		suppression:    Always,
		errorTolerance: forever,
		logger:         core.GetLogger(),
	}

	if o.Limit != nil {
		cf.limit = int(*o.Limit)
	}
	cf.setOptionsDefaults()

	cf.ctx, cf.cancel = context.WithCancel(ctx)
	return cf, nil
}

// SetErrorTolerance sets the duration to suppress errors, measured
// from the previous successful request.
func (cf *ChangesFollower) SetErrorTolerance(d time.Duration) error {
	if d < 0 {
		return core.SDKErrorf(nil, "error tolerance duration must not be negative", "changes-follower-invalid-tolerance", common.GetComponentInfo())
	} else if d == 0 {
		cf.suppression = Never
	} else if d < forever {
		cf.suppression = Timer
	}
	cf.errorTolerance = d
	return nil
}

// Start returns a channel that will stream all available changes
// and keep listening for new changes until reaching an end condition.
//
// The end conditions are:
//   - a terminal error (e.g. unauthorized client).
//   - transient errors occur for longer than the error
//     suppression duration.
//   - the number of changes received reaches the limit specified
//     in the "PostChangesOptions" args used to instantiate
//     this ChangesFollower.
//   - ChangesFollower's Stop() is called.
//
// The same change may be received more than once.
//
// Returns a channel of ChangesItem structs or an error
// if ChangesFollower's Start() or StartOneOff() was already called
// or terminal error is recevied from the service during setup.
func (cf *ChangesFollower) Start() (<-chan ChangesItem, error) {
	return cf.run(Listen)
}

// StartOneOff returns a channel that will stream all available changes
// until there are no further changes pending or reaching an end condition.
//
// The end conditions are:
//   - a terminal error (e.g. unauthorized client).
//   - transient errors occur for longer than the error
//     suppression duration.
//   - the number of changes received reaches the limit specified
//     in the "PostChangesOptions" used to instantiate this ChangesFollower.
//   - ChangesFollower's Stop() is called.
//
// The same change may be received more than once.
//
// Returns a channel of ChangesItem structs or an error
// if ChangesFollower's Start() or StartOneOff() was already called
// or terminal error is recevied from the service during setup.
func (cf *ChangesFollower) StartOneOff() (<-chan ChangesItem, error) {
	return cf.run(Finite)
}

// Stop this ChangesFollower.
func (cf *ChangesFollower) Stop() {
	cf.cancel()
}

func validateOptions(o *cloudantv1.PostChangesOptions) error {
	// this validates that database was properly set
	err := core.ValidateStruct(o, "postChangesOptions")
	if err != nil {
		return err
	}

	errAttrs := make([]string, 0)

	if o.Descending != nil {
		errAttrs = append(errAttrs, "descending")
	}
	if o.Feed != nil {
		errAttrs = append(errAttrs, "feed")
	}
	if o.Heartbeat != nil {
		errAttrs = append(errAttrs, "heartbeat")
	}
	if o.LastEventID != nil {
		errAttrs = append(errAttrs, "lastEventId")
	}
	if o.Timeout != nil {
		errAttrs = append(errAttrs, "timeout")
	}
	if o.Filter != nil && *o.Filter != "_selector" {
		errAttrs = append(errAttrs, fmt.Sprintf("filter=%s", *o.Filter))
	}
	if len(errAttrs) == 1 {
		err := fmt.Errorf("the option '%s' is invalid when using ChangesFollower", errAttrs[0])
		return core.SDKErrorf(err, "", "changes-follower-validation-failed", common.GetComponentInfo())
	}
	if len(errAttrs) > 0 {
		err := fmt.Errorf("the options %s are invalid when using ChangesFollower", strings.Join(errAttrs, ", "))
		return core.SDKErrorf(err, "", "changes-follower-validation-failed", common.GetComponentInfo())
	}

	return nil
}

func (cf *ChangesFollower) setOptionsDefaults() *ChangesFollower {
	switch cf.mode {
	case Finite:
		cf.options.SetFeed(cloudantv1.PostChangesOptionsFeedNormalConst)
	case Listen:
		cf.options.SetFeed(cloudantv1.PostChangesOptionsFeedLongpollConst)
		cf.options.SetTimeout(LongpollTimeout.Milliseconds())
	}
	return cf
}

func (cf *ChangesFollower) withLimit(limit int) *ChangesFollower {
	cf.logger.Debug("Applying changes limit %d", limit)
	cf.options.SetLimit(int64(limit))
	return cf
}

func (cf *ChangesFollower) run(m Mode) (<-chan ChangesItem, error) {
	defer cf.runLock.Unlock()
	cf.runLock.Lock()

	if cf.running {
		return nil, core.SDKErrorf(nil, "cannot start a feed that has already started", "changes-follower-feed-started", common.GetComponentInfo())
	}
	cf.running = true

	switch m {
	case Finite:
		cf.mode = Finite
		cf.since = "0"
	case Listen:
		cf.mode = Listen
		cf.since = "now"
	}

	if cf.options.Since != nil {
		cf.since = *cf.options.Since
	}

	batchSize := BatchSize
	if cf.options.IncludeDocs != nil && *cf.options.IncludeDocs {
		o := cf.client.NewGetDatabaseInformationOptions(*cf.options.Db)

		result, _, err := cf.client.GetDatabaseInformation(o)
		if err != nil {
			return nil, err
		}

		docs := 0
		externalSize := 0
		if result.DocCount != nil {
			docs = int(*result.DocCount)
		}
		if result.Sizes != nil && (*result.Sizes).External != nil {
			externalSize = int(*result.Sizes.External)
		}
		if externalSize > 0 && docs > 0 {
			batchSize = 5 * 1024 * 1024 / (externalSize/docs + 500)
			if batchSize < 1 {
				batchSize = 1
			}
		}
	}

	if cf.limit > 0 && cf.limit < batchSize {
		batchSize = cf.limit
	}

	cf.setOptionsDefaults().withLimit(batchSize)
	cf.successTimestamp = time.Now()

	changes := make(chan ChangesItem)
	go func() {
		defer close(changes)
		for batch := range cf.getChangesBatch() {
			if errors.Is(cf.ctx.Err(), context.Canceled) || errors.Is(cf.ctx.Err(), context.DeadlineExceeded) {
				return
			} else if batch.error != nil {
				select {
				case changes <- ChangesItem{error: batch.error}:
					return
				case <-cf.ctx.Done():
					return
				}
			}
			for _, item := range batch.items {
				if cf.limit == 0 {
					cf.Stop()
					return
				}
				select {
				case changes <- ChangesItem{item: item}:
				case <-cf.ctx.Done():
					return
				}
				if cf.limit > 0 {
					cf.limit--
				}
			}
		}
	}()
	return changes, nil
}

type changesItems struct {
	items []cloudantv1.ChangesResultItem
	error error
}

func (cf *ChangesFollower) getChangesBatch() chan changesItems {
	changes := make(chan changesItems, 1)
	go func() {
		defer close(changes)
		for {
			select {
			case <-cf.ctx.Done():
				return
			default:
			}
			cf.options.SetSince(cf.since)
			result, resp, err := cf.client.PostChangesWithContext(cf.ctx, cf.options)
			if err != nil {
				cf.logger.Debug("Error getting changes: %s", err)
				if resp == nil || isTerminalError(resp.GetStatusCode()) {
					cf.logger.Debug("Terminal error.")
					changes <- changesItems{error: err}
					return
				}
				if cf.suppression == Never || (cf.suppression == Timer && cf.successTimestamp.Add(cf.errorTolerance).Before(time.Now())) {
					cf.logger.Debug("Error tolerance deadline exceeded.")
					changes <- changesItems{error: err}
					return
				}
				cf.retryDelay()
				continue
			}
			cf.since = *result.LastSeq
			cf.retry = 0
			if cf.suppression == Timer {
				cf.successTimestamp = time.Now()
			}
			changes <- changesItems{items: result.Results}
			if cf.mode == Finite && *result.Pending == 0 {
				return
			}
		}
	}()
	return changes
}

func isTerminalError(code int) bool {
	switch code {
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound:
		return true
	default:
		return false
	}
}

// retryDelay implements full jitter delay algorithm.
//
// This is an exponential capped backoff with added jitter to spread
// retry calls in case of multiple followers started simultaneously
// for different feeds on the same account.
//
// The base delay is set to 100 ms and cap is set to the changes
// feed pull's timeout.
//
// Algorithm reference: https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
func (cf *ChangesFollower) retryDelay() {
	expDelay := int(LongpollTimeout)
	if cf.retry < expRetryGate {
		expDelay = int(math.Pow(2, float64(cf.retry)) * float64(baseDelay))
	}
	jitterDelay := rand.Intn(expDelay)
	time.Sleep(time.Duration(jitterDelay))
	cf.retry++
}
