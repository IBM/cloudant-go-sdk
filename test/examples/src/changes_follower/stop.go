/**
 * © Copyright IBM Corporation 2023. All Rights Reserved.
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
