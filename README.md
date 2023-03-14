[![Build Status](https://github.com/IBM/cloudant-go-sdk/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/IBM/cloudant-go-sdk/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/v/release/IBM/cloudant-go-sdk?include_prereleases&sort=semver)](https://github.com/IBM/cloudant-go-sdk/releases/latest)
[![Docs](https://img.shields.io/static/v1?label=Godoc&message=latest&color=blue)](https://pkg.go.dev/github.com/IBM/cloudant-go-sdk)

# IBM Cloudant Go SDK Version 0.3.3

IBM Cloudant Go SDK is a client library that interacts with the
[IBM Cloudant APIs](https://cloud.ibm.com/apidocs/cloudant?code=go).

Disclaimer: This library is still a 0.x release. We do consider this
library production-ready and capable, but there are still some
limitations we’re working to resolve, and refinements we want to
deliver. We are working really hard to minimise the disruption from
now until the 1.0 release, but there may still be some changes that
impact applications using this SDK. For now, be sure to pin versions
to avoid surprises.

<details>
<summary>Table of Contents</summary>

<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  * [`go get` command](#go-get-command)
  * [Go modules](#go-modules)
- [Authentication](#authentication)
  * [Authentication with environment variables](#authentication-with-environment-variables)
    + [IAM authentication](#iam-authentication)
    + [Session cookie authentication](#session-cookie-authentication)
    + [Basic authentication](#basic-authentication)
  * [Authentication with external configuration](#authentication-with-external-configuration)
  * [Programmatic authentication](#programmatic-authentication)
- [Using the SDK](#using-the-sdk)
  * [Request timeout configuration](#request-timeout-configuration)
  * [Code examples](#code-examples)
    + [1. Create a database and add a document](#1-create-a-database-and-add-a-document)
    + [2. Retrieve information from an existing database](#2-retrieve-information-from-an-existing-database)
    + [3. Update your previously created document](#3-update-your-previously-created-document)
    + [4. Delete your previously created document](#4-delete-your-previously-created-document)
    + [Further code examples](#further-code-examples)
  * [Error handling](#error-handling)
  * [Raw IO](#raw-io)
  * [Further resources](#further-resources)
- [Questions](#questions)
- [Issues](#issues)
- [Open source at IBM](#open-source-at-ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

</details>

## Overview

The IBM Cloudant Go SDK allows developers to programmatically
interact with [IBM Cloudant](https://cloud.ibm.com/apidocs/cloudant)
with the help of the `cloudantv1` package.

## Features

The purpose of this Go SDK is to wrap most of the HTTP request APIs
provided by Cloudant and supply other functions to ease the usage of Cloudant.
This SDK should make life easier for programmers to do what’s really important
to them: developing software.

Reasons why you should consider using Cloudant Go SDK in your
project:

- Supported by IBM Cloudant.
- Server compatibility with:
    - IBM Cloudant.
    - [Apache CouchDB 3.x](https://docs.couchdb.org/en/stable/) for data operations.
- Includes all the most popular and latest supported endpoints for
  applications.
- Handles the authentication.
- Familiar user experience with IBM Cloud SDKs.
- Flexibility to use either built-in models or byte-based requests and responses for documents.
- HTTP2 support for higher performance connections to IBM Cloudant.
- Perform requests synchronously.
- Safe for concurrent use by multiple goroutines.
- Transparently compresses request and response bodies.

## Prerequisites

- A
  [Cloudant](https://cloud.ibm.com/docs/Cloudant/getting-started.html#step-1-connect-to-your-cloudant-nosql-db-service-instance-on-ibm-cloud)
  service instance or a
  [CouchDB](https://docs.couchdb.org/en/latest/install/index.html)
  server.
- Go version 1.18 or 1.19.

## Installation

The current version of this SDK: 0.3.3

There are a few different ways to download and add the
Cloudant Go SDK project for use by your Go application:

### `go get` command

Use this command to download and add the SDK to allow your Go application to
use it:

```terminal
go get -u github.com/IBM/cloudant-go-sdk/cloudantv1@v0.3.3
```

### Go modules

If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
  "github.com/IBM/cloudant-go-sdk/cloudantv1"
)
```

then run `go mod tidy` to download and install the latest version if
the dependency is missing. This command will also update your Go
application's `go.mod` file.

## Authentication

[service-credentials]: https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-locating-your-service-credentials
[cloud-IAM-mgmt]: https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-managing-access-for-cloudant#introduction-iam-ai
[couch-cookie-auth]: https://docs.couchdb.org/en/stable/api/server/authn.html#cookie-authentication
[cloudant-cookie-auth]: https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-work-with-your-account#cookie-authentication
[couch-basic-auth]: https://docs.couchdb.org/en/stable/api/server/authn.html#basic-authentication
[cloudant-basic-auth]: https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-work-with-your-account#basic-authentication

This library requires some of your
[Cloudant service credentials][service-credentials] to authenticate with your
account.

1. `IAM`, `COUCHDB_SESSION`, `BASIC` or `NOAUTH` **authentication type**.
    1. [*IAM authentication*](#iam-authentication) is highly recommended when your
    back-end database server is [**Cloudant**][cloud-IAM-mgmt]. This
    authentication type requires a server-generated `apikey` instead of a
    user-given password. You can create one
    [here](https://cloud.ibm.com/iam/apikeys).
    1. [*Session cookie (`COUCHDB_SESSION`) authentication*](#session-cookie-authentication)
    is recommended for [Apache CouchDB][couch-cookie-auth] or for
    [Cloudant][cloudant-cookie-auth] when IAM is unavailable. It exchanges username
    and password credentials for an `AuthSession` cookie from the `/_session`
    endpoint.
    1. [*Basic* (or legacy) *authentication*](#basic-authentication) is a fallback
    for both [Cloudant][cloudant-basic-auth] and [Apache CouchDB][couch-basic-auth]
    back-end database servers. This authentication type requires the good old
    `username` and `password` credentials.
    1. *Noauth* authentication does not require credentials. Note that this
    authentication type only works with queries against a database with read
    access for everyone.
1. The service `url`.

There are several ways to **set** these properties:

1. As [environment variables](#authentication-with-environment-variables)
1. The [programmatic approach](#programmatic-authentication)
1. With an [external credentials file](#authentication-with-external-configuration)

### Authentication with environment variables

#### IAM authentication

For Cloudant *IAM authentication*, set the following environmental variables by
replacing the `<url>` and `<apikey>` with your proper
[service credentials][service-credentials]. There is no need to set
`CLOUDANT_AUTH_TYPE` to `IAM` because it is the default.

```bash
CLOUDANT_URL=<url>
CLOUDANT_APIKEY=<apikey>
```

#### Session cookie authentication

For `COUCHDB_SESSION` authentication, set the following environmental variables
by replacing the `<url>`, `<username>` and `<password>` with your proper
[service credentials][service-credentials].

```bash
CLOUDANT_AUTH_TYPE=COUCHDB_SESSION
CLOUDANT_URL=<url>
CLOUDANT_USERNAME=<username>
CLOUDANT_PASSWORD=<password>
```

#### Basic authentication

For *Basic authentication*, set the following environmental variables by
replacing the `<url>`, `<username>` and `<password>` with your proper
[service credentials][service-credentials].

```bash
CLOUDANT_AUTH_TYPE=BASIC
CLOUDANT_URL=<url>
CLOUDANT_USERNAME=<username>
CLOUDANT_PASSWORD=<password>
```

**Note**: There are also additional [*Bearer token*](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#bearer-token-authentication), [*Container*](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#container-authentication) and [*VPC Instance*](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#vpc-instance-authentication) authentication methods. For more details, please follow the provided links.
We recommend that you use [IAM](#iam-authentication) for Cloudant and
[Session](#session-cookie-authentication) for CouchDB authentication.

### Authentication with external configuration

To use an external configuration file, the
[Cloudant API docs](https://cloud.ibm.com/apidocs/cloudant?code=go#authentication-with-external-configuration),
or the
[general SDK usage information](https://github.com/IBM/ibm-cloud-sdk-common#using-external-configuration)
will guide you.

### Programmatic authentication

To learn more about how to use programmatic authentication, see the related
documentation in the
[Cloudant API docs](https://cloud.ibm.com/apidocs/cloudant?code=go#programmatic-authentication)
or in the
[Go SDK Core document](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md) about authentication.

## Using the SDK

For fundamental SDK usage information and config options, please see the common [IBM Cloud SDK](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md) documentation.

### Request timeout configuration

A 6m request timeout, which includes a 30s connect timeout, is set by default. Note that this also affects changes feed requests, regardless of a timeout set on `PostChangesOptions`. Be sure to set a request timeout appropriate to your application usage and environment.
The [request timeout](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md) section contains details on how to change the value.

**Note:** System settings may take precedence over configured timeout values.

### Code examples

The following code examples
[authenticate with the environment variables](#authenticate-with-environment-variables).

#### 1. Create a database and add a document

**Note:** This example code assumes that `orders` database does not exist in your account.

This example code creates `orders` database and adds a new document "example"
into it. To connect, you must set your environment variables with
the *service url*, *authentication type* and *authentication credentials*
of your Cloudant service.

Cloudant environment variable naming starts with a *service name* prefix that identifies your service.
By default this is `CLOUDANT`, see the settings in the
[authentication with environment variables section](#authentication-with-environment-variables).

If you would like to rename your Cloudant service from `CLOUDANT`,
you must use your defined service name as the prefix for all Cloudant related environment variables.

Once the environment variables are set, you can try out the code examples.

[embedmd]:# (test/examples/src/create_db_and_doc/create_db_and_doc.go /package main/ $)
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

	// Keeping track of the revision number of the document object
	// is necessary for further UPDATE/DELETE operations:
	exampleDocument.Rev = createDocumentResponse.Rev

	// Print out the document content
	exampleDocumentContent, _ := json.MarshalIndent(exampleDocument, "", "  ")
	fmt.Printf("You have created the document:\n%s\n", string(exampleDocumentContent))
}
```

When you run the code, you see a result similar to the following output.

[embedmd]:# (test/examples/output/create_db_and_doc.txt)
```txt
"orders" database created.
You have created the document:
{
  "_id": "example",
  "_rev": "1-1b403633540686aa32d013fda9041a5d",
  "joined": "2019-01-24T10:42:99.000Z",
  "name": "Bob Smith"
}
```

#### 2. Retrieve information from an existing database

**Note**: This example code assumes that you have created both the `orders`
database and the `example` document by
[running the previous example code](#1-create-a-database-and-add-a-document)
successfully. Otherwise, the following error message occurs, "Cannot delete document because either 'orders'
database or 'example' document was not found."

<details>
<summary>Gather database information example</summary>

[embedmd]:# (test/examples/src/get_info_from_existing_database/get_info_from_existing_database.go /package main/ $)
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

[embedmd]:# (test/examples/output/get_info_from_existing_database.txt)
```txt
Server Version: 2.1.1
Document count in "orders" database is 1.
{
  "_id": "example",
  "_rev": "1-1b403633540686aa32d013fda9041a5d",
  "name": "Bob Smith",
  "joined": "2019-01-24T10:42:99.000Z"
}
```

#### 3. Update your previously created document

**Note**: This example code assumes that you have created both the `orders`
database and the `example` document by
[running the previous example code](#1-create-a-database-and-add-a-document)
successfully. Otherwise, the following error message occurs, "Cannot update document because either 'orders'
database or 'example' document was not found."

<details>
<summary>Update code example</summary>

[embedmd]:# (test/examples/src/update_doc/update_doc.go /package main/ $)
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

[embedmd]:# (test/examples/output/update_doc.txt)
```txt
You have updated the document:
{
  "_id": "example",
  "_rev": "2-4e2178e85cffb32d38ba4e451f6ca376",
  "address": "19 Front Street, Darlington, DL5 1TY",
  "name": "Bob Smith"
}
```

#### 4. Delete your previously created document

**Note**: This example code assumes that you have created both the `orders`
database and the `example` document by
[running the previous example code](#1-create-a-database-and-add-a-document)
successfully. Otherwise, the following error message occurs, "Cannot delete document because either 'orders'
database or 'example' document was not found."

<details>
<summary>Delete code example</summary>

[embedmd]:# (test/examples/src/delete_doc/delete_doc.go /package main/ $)
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

[embedmd]:# (test/examples/output/delete_doc.txt)
```txt
You have deleted the document.
```

#### Further code examples

For a complete list of code examples, see the [examples directory](examples#examples-for-go).

### Error handling

For sample code on handling errors, see
[Cloudant API docs](https://cloud.ibm.com/apidocs/cloudant?code=go#error-handling).

### Raw IO

For endpoints that read or write document content it is possible to bypass
usage of the built-in struct with byte streams.

Depending on the specific SDK operation it may be possible to:
* accept a user-provided byte stream to send to the server as a request body
* return a byte stream of the server response body to the user

Request byte stream can be supplied for method options that have a `Body` property.
For these cases you can pass this byte stream directly to the HTTP request body.

Response byte stream is supported in functions with the suffix of `AsStream`.
The returned byte stream allows the response body to be consumed
without triggering JSON unmarshalling that is typically performed by the SDK.

The [update document](#3-update-your-previously-created-document) section
contains examples for both request and response byte stream cases.

The API reference contains further examples of using byte streams.
They are titled "Example request as stream" and are initially collapsed.
Expand them to see examples of:

- Byte requests:
  - [Bulk modify multiple documents in a database](https://cloud.ibm.com/apidocs/cloudant?code=go#postbulkdocs)

- Byte responses:
  - [Query a list of all documents in a database](https://cloud.ibm.com/apidocs/cloudant?code=go#postalldocs)
  - [Query the database document changes feed](https://cloud.ibm.com/apidocs/cloudant?code=go#postchanges)

### Further resources

- [Cloudant API docs](https://cloud.ibm.com/apidocs/cloudant?code=go):
  API reference including usage examples for Cloudant Go SDK API.
- [Godoc](https://pkg.go.dev/github.com/IBM/cloudant-go-sdk):
  Cloudant Go SDK API Documentation.
- [Cloudant docs](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-getting-started-with-cloudant):
  The official documentation page for Cloudant.
- [Cloudant blog](https://blog.cloudant.com/):
  Many useful articles about how to optimize Cloudant for common problems.

## Questions

If you are having difficulties using this SDK or have a question about the
IBM Cloud services, ask a question on
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues

If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/cloudant-go-sdk/issues).

Before you submit a bug report, search for
[similar issues](https://github.com/IBM/cloudant-go-sdk/issues?q=is%3Aissue) and review the
[KNOWN_ISSUES file](KNOWN_ISSUES.md) to verify that your issue hasn't been reported yet.

Please consult the [security policy](https://github.com/IBM/cloudant-go-sdk/security/policy) before opening security related issues.

## Open source at IBM

Find more open source projects on the [IBM GitHub](http://ibm.github.io/) page.

## Contributing

For more information, see [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK is released under the Apache 2.0 license. To read the full text of the license, see [LICENSE](LICENSE).
