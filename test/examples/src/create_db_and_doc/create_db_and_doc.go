/**
 * © Copyright IBM Corporation 2020, 2022. All Rights Reserved.
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
	"encoding/json"
	"fmt"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func main() {
	// 1. Create a client with `CLOUDANT` default service name =============
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{},
	)
	if err != nil {
		panic(err)
	}

	// 2. Create a database ================================================
	exampleDbName := "orders"

	putDatabaseResult, putDatabaseResponse, err := client.PutDatabase(
		client.NewPutDatabaseOptions(exampleDbName),
	)
	if err != nil {
		if putDatabaseResponse.StatusCode == 412 {
			fmt.Printf("Cannot create \"%s\" database, it already exists.\n",
				exampleDbName)
		} else {
			panic(err)
		}
	}

	if putDatabaseResult != nil && *putDatabaseResult.Ok {
		fmt.Printf("\"%s\" database created.\n", exampleDbName)
	}

	// 3. Create a document ================================================
	// Create a document object with "example" id
	exampleDocID := "example"
	// Setting ID for the document is optional when "PostDocument" function
	// is used for CREATE.
	// When ID is not provided the server will generate one for your document.
	exampleDocument := cloudantv1.Document{
		ID: &exampleDocID,
	}

	// Add "name" and "joined" fields to the document
	exampleDocument.SetProperty("name", "Bob Smith")
	exampleDocument.SetProperty("joined", "2019-01-24T10:42:59.000Z")

	// Save the document in the database with "PostDocument" function
	createDocumentOptions := client.NewPostDocumentOptions(
		exampleDbName,
	).SetDocument(&exampleDocument)

	createDocumentResponse, _, err := client.PostDocument(createDocumentOptions)

	// =====================================================================
	// Note: saving the document can also be done with the "PutDocument"
	// function. In this case docID is required for a CREATE operation:
	/*
		createDocumentOptions := client.NewPutDocumentOptions(
			exampleDbName,
			exampleDocID,
		).SetDocument(&exampleDocument)

		createDocumentResponse, _, err := client.PutDocument(createDocumentOptions)
	*/
	// =====================================================================

	if err != nil {
		panic(err)
	}

	// Keeping track of the revision number of the document object
	// is necessary for further UPDATE/DELETE operations:
	// rev := createDocumentResponse.Rev

	// Print out the response body
	responseBody, _ := json.MarshalIndent(createDocumentResponse, "", "  ")
	fmt.Printf("You have created the document. Response body:\n%s\n", string(responseBody))
}
