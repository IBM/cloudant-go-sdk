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
