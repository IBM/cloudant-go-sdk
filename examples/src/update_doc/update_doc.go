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

  // 2. Update the document ==============================================
  exampleDbName := "orders"
  exampleDocID := "example"

  // Get the document if it previously existed in the database
  document, getDocumentResponse, err := client.GetDocument(
    client.NewGetDocumentOptions(
      exampleDbName,
      exampleDocID,
    ),
  )

  // =====================================================================
  // Note: for response byte stream use:
  /*
  	documentAsByteStream, getDocumentResponse, err := client.GetDocumentAsStream(
  		client.NewGetDocumentOptions(
  			exampleDbName,
  			exampleDocID,
  		),
  	)
  */
  // =====================================================================

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
    // Make some modification in the document content
    // Add Bob Smith's address to the document
    document.SetProperty("address", "19 Front Street, Darlington, DL5 1TY")
    // Remove the joined property from document object
    delete(document.GetProperties(), "joined")

    // Update the document in the database
    updateDocumentOptions := client.NewPostDocumentOptions(
      exampleDbName,
    ).SetDocument(document)

    // =================================================================
    // Note: for request byte stream use:
    /*
    	postDocumentOption := client.NewPostDocumentOptions(
    		exampleDbName,
    	).SetBody(documentAsByteStream)
    */
    // =================================================================

    updateDocumentResponse, _, err := client.PostDocument(
      updateDocumentOptions,
    )

    // =================================================================
    // Note: updating the document can also be done with the "PutDocument"
    // function. DocID and Rev are required for an UPDATE operation
    // but Rev can be provided in the document object too:
    /*
    	updateDocumentOptions := client.NewPutDocumentOptions(
    		exampleDbName,
    		core.StringNilMapper(document.ID), // docID is a required parameter
    	).SetDocument(document) // Rev in the document object CAN replace below SetRev

    	updateDocumentOptions.SetRev(core.StringNilMapper(document.Rev))

    	updateDocumentResponse, _, err := client.PutDocument(
    		updateDocumentOptions,
    	)
    */
    // =================================================================

    if err != nil {
      panic(err)
    }

    // Keeping track of the latest revision number of the document object
    // is necessary for further UPDATE/DELETE operations:
    document.Rev = updateDocumentResponse.Rev

    // Print out the new document content
    documentContent, _ := json.MarshalIndent(document, "", "  ")
    fmt.Printf("You have updated the document:\n%s\n", string(documentContent))
  }
}
