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
