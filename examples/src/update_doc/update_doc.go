// © Copyright IBM Corporation 2020. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"

	"github.ibm.com/cloudant/cloudant-go-sdk/cloudantv1"
)

func main() {
	// 1. Create a client with `IBM_CLOUDANT` default service name =========
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{},
	)
	if err != nil {
		panic(err)
	}

	// 2. Update the document ==============================================
	exampleDbName := "orders"
	exampleDocID  := "example"

	// 2.1. Get the document if it previously existed in the database
	document, getDocumentResponse, err := client.GetDocument(
		client.NewGetDocumentOptions(
			exampleDbName,
			exampleDocID,
		),
	)
	if err != nil {
		if getDocumentResponse.StatusCode == 404 {
			fmt.Printf("Cannot update document because "+
				"either \"%s\"  database or \"%s\" document was not found.\n",
				exampleDbName,
				exampleDocID)
		} else {
			panic(err)
		}
	}

	if document != nil {
		// 2.2. Make some modifiction in the document content
		// 2.2.1. Add Bob Smith's address to the document
		document.SetProperty("address", "19 Front Street, Darlington, DL5 1TY")
		// 2.2.2. Remove the joined property from document object
		delete(document.GetProperties(), "joined")

		// 2.3. Update the document in the database
		postDocumentOption := client.NewPostDocumentOptions(
			exampleDbName,
		).SetDocument(document).SetContentType("application/json")

		postDocumentResult, _, err := client.PostDocument(
			postDocumentOption,
		)
		if err != nil {
			panic(err)
		}

		// 2.4. Keep track the revision number of the document object
		document.Rev = postDocumentResult.Rev

		// 2.5. Print out the new document content
		documentContent, _ := json.MarshalIndent(document, "", "  ")
		fmt.Printf("You have updated the document:\n%s\n", string(documentContent))
	}
}
