/**
 * © Copyright IBM Corporation 2020, 2021. All Rights Reserved.
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
	// 3.1. Create a document object with "example" id
	exampleDocID := "example"
	exampleDocument := cloudantv1.Document{
		ID: &exampleDocID,
	}

	// 3.2. Add "name" and "joined" fields to the document
	exampleDocument.SetProperty("name", "Bob Smith")
	exampleDocument.SetProperty("joined", "2019-01-24T10:42:99.000Z")

	// 3.3. Save the document in the database
	postDocumentOption := client.NewPostDocumentOptions(
		exampleDbName,
	).SetDocument(&exampleDocument)

	postDocumentResult, _, err := client.PostDocument(postDocumentOption)
	if err != nil {
		panic(err)
	}

	// 3.4. Keep track of the revision number from the `example` document object
	exampleDocument.Rev = postDocumentResult.Rev

	// 3.5. Print out the document content
	exampleDocumentContent, _ := json.MarshalIndent(exampleDocument, "", "  ")
	fmt.Printf("You have created the document:\n%s\n", string(exampleDocumentContent))
}
