# Code examples

<details open>
<summary>Table of Contents</summary>

<!-- toc -->
- [1. Create a database and add a document](#1-create-a-database-and-add-a-document)
- [2. Retrieve information from an existing database](#2-retrieve-information-from-an-existing-database)
- [3. Update your previously created document](#3-update-your-previously-created-document)
- [4. Delete your previously created document](#4-delete-your-previously-created-document)
- [Further code examples](#further-code-examples)
</details>

The following code examples
[authenticate with the environment variables](Authentication.md#authentication-with-environment-variables).

## 1. Create a database and add a document

**Note:** This example code assumes that `orders` database does not exist in your account.

This example code creates `orders` database and adds a new document "example"
into it. To connect, you must set your environment variables with
the *service url*, *authentication type* and *authentication credentials*
of your Cloudant service.

Cloudant environment variable naming starts with a *service name* prefix that identifies your service.
By default, this is `CLOUDANT`, see the settings in the
[authentication with environment variables section](Authentication.md#authentication-with-environment-variables).

If you would like to rename your Cloudant service from `CLOUDANT`,
you must use your defined service name as the prefix for all Cloudant related environment variables.

Once the environment variables are set, you can try out the code examples.

```go
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

	// Print out the response body
	responseBody, _ := json.MarshalIndent(createDocumentResponse, "", "  ")
	fmt.Printf("You have created the document. Response body:\n%s\n", string(responseBody))
}
```

When you run the code, you see a result similar to the following output.

```text
"orders" database created.
You have created the document. Response body:
{
  "id": "example",
  "rev": "1-1b403633540686aa32d013fda9041a5d",
  "ok": true
}
```

## 2. Retrieve information from an existing database

**Note**: This example code assumes that you have created both the `orders`
database and the `example` document by
[running the previous example code](#1-create-a-database-and-add-a-document)
successfully. Otherwise, the following error message occurs, "Cannot delete document because either 'orders'
database or 'example' document was not found."

<details open>
<summary>Gather database information example</summary>

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func main() {
	// 1. Create a client with `CLOUDANT` default service name ============
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{},
	)
	if err != nil {
		panic(err)
	}
	// 2. Get server information ===========================================
	serverInformationResult, _, err := client.GetServerInformation(
		client.NewGetServerInformationOptions(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server Version: %s\n", *serverInformationResult.Version)
	// 3. Get database information for "orders" ==========================
	dbName := "orders"
	databaseInformationResult, _, err := client.GetDatabaseInformation(
		client.NewGetDatabaseInformationOptions(
			dbName,
		),
	)
	if err != nil {
		panic(err)
	}
	// 4. Show document count in database ==================================
	fmt.Printf("Document count in \"%s\" database is %d.\n",
		*databaseInformationResult.DbName,
		*databaseInformationResult.DocCount)
	// 5. Get "example" document out of the database by document id ============
	documentExampleResult, _, err := client.GetDocument(
		client.NewGetDocumentOptions(
			dbName,
			"example",
		),
	)
	if err != nil {
		panic(err)
	}
	// 6. Print out the Document content ===================================
	exampleBuffer, _ := json.MarshalIndent(documentExampleResult, "", "  ")
	fmt.Println(string(exampleBuffer))
}
```

</details>
When you run the code, you see a result similar to the following output.

```text
Server Version: 3.2.1
Document count in "orders" database is 1.
{
  "_id": "example",
  "_rev": "1-1b403633540686aa32d013fda9041a5d",
  "joined": "2019-01-24T10:42:59.000Z",
  "name": "Bob Smith"
}
```

## 3. Update your previously created document

**Note**: This example code assumes that you have created both the `orders`
database and the `example` document by
[running the previous example code](#1-create-a-database-and-add-a-document)
successfully. Otherwise, the following error message occurs, "Cannot update document because either 'orders'
database or 'example' document was not found."

<details open>
<summary>Update code example</summary>

```go
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
```

</details>
When you run the code, you see a result similar to the following output.

```text
You have updated the document:
{
  "_id": "example",
  "_rev": "2-4e2178e85cffb32d38ba4e451f6ca376",
  "address": "19 Front Street, Darlington, DL5 1TY",
  "name": "Bob Smith"
}
```

## 4. Delete your previously created document

**Note**: This example code assumes that you have created both the `orders`
database and the `example` document by
[running the previous example code](#1-create-a-database-and-add-a-document)
successfully. Otherwise, the following error message occurs, "Cannot delete document because either 'orders'
database or 'example' document was not found."

<details open>
<summary>Delete code example</summary>

```go
package main

import (
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

	// 2. Delete the document ==============================================
	exampleDbName := "orders"
	exampleDocID := "example"

	// 2.1. Get the document if it previously existed in the database
	document, getDocumentResponse, err := client.GetDocument(
		client.NewGetDocumentOptions(
			exampleDbName,
			exampleDocID,
		),
	)
	if err != nil {
		if getDocumentResponse.StatusCode == 404 {
			fmt.Printf("Cannot delete document because "+
				"either \"%s\" database or \"%s\" document was not found.\n",
				exampleDbName,
				exampleDocID)
		} else {
			panic(err)
		}
	}
	// 2.2. Use its latest revision to delete
	if document != nil {
		deleteDocumentResult, _, err := client.DeleteDocument(
			client.NewDeleteDocumentOptions(
				exampleDbName,
				*document.ID, // docID is required for DELETE
			).SetRev(*document.Rev), // Rev is required for DELETE
		)
		if err != nil {
			panic(err)
		}

		if *deleteDocumentResult.Ok {
			fmt.Println("You have deleted the document.")
		}
	}
}
```

</details>
When you run the code, you see the following output.

```text
You have deleted the document.
```

## Further code examples

For a complete list of code examples, see the [examples directory](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.10/examples#examples-for-go).
