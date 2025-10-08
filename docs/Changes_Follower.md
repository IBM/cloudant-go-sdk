# Changes feed follower

<details open>
<summary>Table of Contents</summary>

<!-- toc -->
- [Introduction](#introduction)
- [Modes of operation](#modes-of-operation)
- [Configuring the changes follower](#configuring-the-changes-follower)
- [Error suppression](#error-suppression)
- [Follower operation](#follower-operation)
- [Checkpoints](#checkpoints)
- [Code examples](#code-examples)
  * [Initializing a changes follower](#initializing-a-changes-follower)
  * [Starting the changes follower](#starting-the-changes-follower)
    + [Start mode for continuous listening](#start-mode-for-continuous-listening)
    + [Start mode for one-off fetching](#start-mode-for-one-off-fetching)
  * [Processing changes](#processing-changes)
    + [Process continuous changes](#process-continuous-changes)
    + [Process one-off changes](#process-one-off-changes)
  * [Stopping the changes follower](#stopping-the-changes-follower)
</details>

## Introduction

The SDK provides a changes feed follower utility.
This helper utility connects to the `_changes` endpoint and returns the individual change items.
It removes some complexity of using the `_changes` endpoint by setting some options automatically
and providing error suppression and retries.

*Tip: the changes feed often does not meet user expectations or assumptions.*

Consult the [Cloudant changes feed FAQ](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-faq-using-changes-feed)
to get a better understanding of the limitations and suitable use-cases before using the changes feed in your application.

## Modes of operation

There are two modes of operation:
* Start mode
  * Fetches the changes from the supplied `since` sequence (in this mode follower defaults to reading the feed from `now`).
  * Fetches all available changes and then continues listening for new changes indefinitely unless encountering an end condition.
  * An example use case for this mode is event driven workloads.
* Start one-off mode
  * Fetches the changes from the supplied `since` sequence (in this mode follower defaults to reading the feed from the beginning).
  * Fetches all available changes and then stops when either there are no further changes pending or encountering an end condition.
  * An example use case for this mode is ETL style workloads.

## Configuring the changes follower

The SDK's model of changes feed options is also used to configure the follower.
However, it is invalid to configure some options used internally by the follower implementation.
Supplying these options when instantiating the follower causes an error.
The invalid options are:
* `descending`
* `feed`
* `heartbeat`
* `lastEventId` - use `since` instead
* `timeout`
* Follower permits only the value of `_selector` for the `filter` option. This restriction is because selector
  based filters perform better than JavaScript backed filters. Configuring a non-selector based filter
  causes the follower to error.

Note that the `limit` parameter terminates the follower at the given number of changes in either
operating mode.

The changes follower requires the client to have HTTP timeouts of at least 1 minute and errors during
instantiation if it is insufficient. The default client configuration has sufficiently long timeouts.

For use-cases where these configuration limitations are too restrictive then write code to use the SDK's
[POST `_changes` API](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.7/examples#postchanges) instead of the follower.

## Error suppression

By default, the changes follower suppresses transient errors indefinitely and attempts to run to completion or listen forever as
dictated by the operating mode.
For applications where that is not desirable configure the optional error tolerance duration. This controls the time since
the last successful response that the follower suppresses transient errors. An example usage is an application grace period
before reporting an error and requiring intervention.

There are some additional points to consider for error suppression:
* Errors considered terminal, for example, the database not existing or invalid credentials are never suppressed and error immediately.
* The error suppression duration is not guaranteed to fire immediately after lapsing and is a minimum suppression time.
* The changes follower backs-off between retries and as such may remain paused for a short while after the transient errors have resolved.
* If the underlying SDK client used to initialize the follower also has retries configured then suppression of errors may last
  significantly longer than the follower's configured error tolerance duration depending on the specific options.

## Follower operation

For both modes:
* The end conditions are:
  * A terminal error (HTTP codes `400`, `401`, `403` `404`).
  * Transient errors occur for longer than the error tolerance duration. Transient errors are all other HTTP status codes and connection errors.
  * The number of changes received reaches the configured `limit`.
  * The application calls stop to terminate the feed early.

As is true for the `_changes` endpoint change items have *at least once* delivery and callers may receive
an individual item multiple times. When using the follower change items may repeat even within a limited
number of changes (that is using the `limit` option) this is a minor difference from using `limit` on the HTTP native API.

The follower is not optimized for some use cases and it is not recommended to use it in cases where:
* Setting `include_docs` and larger document sizes (for example > 10 kiB).
* The volume of changes is very high (if the rate of changes in the database exceeds the follower's rate of pulling them it can never catch-up).

In these use-cases use the SDK's [POST `_changes` API](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.7/examples#postchanges)
for  specific control over the number of change requests made and the content size of the responses.

## Checkpoints

The changes follower does not checkpoint since it has no information about whether the consuming application
has processed a change item after delivery. It is the application developer's responsibility
to store the sequence IDs to have appropriate checkpoints and to re-initialize the follower with the required
`since` value after, for example, the application restarts.

The frequency and conditions for checkpoints are application specific and some applications may be tolerant
of dropped changes. This section provides only general guidance on how to avoid missing changes.

To guarantee processing of all changes do not persist the sequence ID from a change item until *after*
the processing of the change item by the application has completed. As indicated previously change item
delivery is *at least once* so application code must be able to handle repeated changes already. It is
preferable to restart from an older `since` value and receive changes again than risk missing them.

The sequence IDs are available on each change item by default. However, the server omits sequence IDs from
some change items when using the `seq_interval` configuration option.
Infrequent sequence IDs may improve performance by reducing the amount of data transfer and server load,
but the tradeoff is repeating more changes if it is necessary to resume the changes follower.

Take extreme care persisting sequences if choosing to process change items in parallel as there
is a considerable risk of missing changes on a restart if the recorded sequence is out of order.

## Code examples

### Initializing a changes follower
```go
package main

import (
	"time"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/cloudant-go-sdk/features"
)

func main() {
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{},
	)
	if err != nil {
		panic(err)
	}

	// Required: the database name.
	postChangesOptions := client.NewPostChangesOptions("example")

	// Optional: return only 100 changes (including duplicates).
	postChangesOptions.SetLimit(100)

	// Optional: start from this sequence ID (e.g. with a value read from persistent storage).
	postChangesOptions.SetSince("3-g1AG3...")

	// Required: the Cloudant service client instance and an instance of PostChangesOptions
	follower, err := features.NewChangesFollower(client, postChangesOptions)
	if err != nil {
		panic(err)
	}

	// Optional: suppress transient errors for at least 10 seconds before terminating.
	err = follower.SetErrorTolerance(10 * time.Second)
	if err != nil {
		panic(err)
	}
}
```

### Starting the changes follower

#### Start mode for continuous listening
```go
package main

import (
	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/cloudant-go-sdk/features"
)

func main() {
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{},
	)
	if err != nil {
		panic(err)
	}

	postChangesOptions := client.NewPostChangesOptions("example")

	follower, err := features.NewChangesFollower(client, postChangesOptions)
	if err != nil {
		panic(err)
	}

	changesCh, err := follower.Start()
	if err != nil {
		panic(err)
	}

	// Note: changesCh channel will not do anything until it is read from.
	// Create a range loop to iterate over the flow of the changes items
}
```

#### Start mode for one-off fetching
```go
package main

import (
	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/cloudant-go-sdk/features"
)

func main() {
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{},
	)
	if err != nil {
		panic(err)
	}

	postChangesOptions := client.NewPostChangesOptions("example")

	follower, err := features.NewChangesFollower(client, postChangesOptions)
	if err != nil {
		panic(err)
	}

	changesCh, err = follower.StartOneOff()
	if err != nil {
		panic(err)
	}

	// Note: the returned channel will not do anything until it is read
	// Create a range loop to iterate over the flow of the change items
}
```

### Processing changes

#### Process continuous changes
```go
package main

import (
	"fmt"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/cloudant-go-sdk/features"
)

func main() {
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{},
	)
	if err != nil {
		panic(err)
	}

	postChangesOptions := client.NewPostChangesOptions("example")

	// Start from a previously persisted seq
	// Normally this would be read by the app from persistent storage
	// e.g. prevSeq = yourAppPersistenceReadFunc()
	prevSeq := "3-g1AG3..."
	postChangesOptions.SetSince(prevSeq)

	follower, err := features.NewChangesFollower(client, postChangesOptions)
	if err != nil {
		panic(err)
	}

	changesCh, err := follower.Start()
	if err != nil {
		panic(err)
	}

	for changesItem := range changesCh {
		// changes item returns an error on failed requests
		item, err := changesItem.Item()
		if err != nil {
			panic(err)
		}

		// do something with changes
		fmt.Printf("%s\n", *item.ID)
		for _, change := range item.Changes {
			fmt.Printf("%s\n", *change.Rev)
		}

		// when change item processing is complete app can store seq
		seq := *item.Seq
		// write seq to persistent storage for use as since
		// if required to resume later, e.g. yourAppPersistenceWriteFunc(seq)
		// keep processing changes until the application is terminated
		// or some other stop condition is reached
	}

	// Note: the returned channel continues to listen for changes
	// indefinitely. In other words iterating it is blocking,
	// code here will be unreachable until break from the range
	// or another stop condition is reached.
	// For long running followers careful consideration
	// should be made of where to call stop on the follower.
}
```

#### Process one-off changes
```go
package main

import (
	"fmt"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/cloudant-go-sdk/features"
)

func main() {
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{},
	)
	if err != nil {
		panic(err)
	}

	postChangesOptions := client.NewPostChangesOptions("example")

	// Start from a previously persisted seq
	// Normally this would be read by the app from persistent storage
	// e.g. prevSeq = yourAppPersistenceReadFunc()
	prevSeq := "3-g1AG3..."
	postChangesOptions.SetSince(prevSeq)

	follower, err := features.NewChangesFollower(client, postChangesOptions)
	if err != nil {
		panic(err)
	}

	changesCh, err := follower.StartOneOff()
	if err != nil {
		panic(err)
	}

	for changesItem := range changesCh {
		// changes item structure returns an error on failed requests
		item, err := changesItem.Item()
		if err != nil {
			panic(err)
		}

		// do something with changes
		fmt.Printf("%s\n", *item.ID)
		for _, change := range item.Changes {
			fmt.Printf("%s\n", *change.Rev)
		}

		// when change item processing is complete app can store seq
		seq := *item.Seq
		// write seq to persistent storage for use as since if required
		// to resume later, e.g. yourAppPersistenceWriteFunc(seq)
	}

	// Note: iterating the returned channel above is blocking, code here
	// will be unreachable until all changes are processed
	// or another stop condition is reached.
}
```

### Stopping the changes follower
```go
package main

import (
	"context"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/cloudant-go-sdk/features"
)

func main() {
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{},
	)
	if err != nil {
		panic(err)
	}

	postChangesOptions := client.NewPostChangesOptions("example")

	ctx, cancel := context.WithCancel(context.Background())
	// Option 1: call cancel function when ChangesFollower created
	// with a cancellable context.
	// Note that this is a good practice to always call cancel function,
	// e.g. in defer to prevent possible goroutine leak.
	defer cancel()
	follower, err := features.NewChangesFollowerWithContext(ctx, client, postChangesOptions)
	if err != nil {
		panic(err)
	}

	changesCh, err := follower.Start()
	if err != nil {
		panic(err)
	}

	for changesItem := range changesCh {
		// changes item structure returns an error on failed requests
		if changesItem.Error != nil {
			panic(changesItem.Error)
		}

		// Option 2: call stop after some condition
		// Note that at least one item must be returned from the channel
		// to reach this point. Additional changes may be processed
		// before the channel quits.
		follower.Stop()
	}

	// Option 3: call stop method when you want to end the continuous loop from
	// outside the channel. For example, you've put the changes follower in a
	// goroutine and need to call stop on the main goroutine.
	// Note: in this context the call must be made from a different
	// goroutine because code immediately following the range is unreachable
	// until the channel has quit.
	follower.Stop()
}
```
