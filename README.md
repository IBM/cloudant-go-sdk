<!--
  The example codes and outputs below are generated using the `embedmd` go package.

      https://github.com/campoy/embedmd

  You should regenerate the example codes after making any changes to examples in the test/examples/ folder.

      embedmd -w README.md
  -->

[![Build Status](https://travis-ci.com/IBM/cloudant-go-sdk.svg?branch=master)](https://travis-ci.com/IBM/cloudant-go-sdk)
<!-- [![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release) -->

# IBM Cloudant Go SDK Version 0.0.26

Go client library to interact with the various [IBM Cloudant APIs](https://cloud.ibm.com/apidocs/cloudant?code=go).

Disclaimer: this SDK is being released initially as a **pre-release** version.
Changes might occur which impact applications that use this SDK.

<details>
<summary>Table of Contents</summary>

<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
    + [`go get` command](#go-get-command)
    + [Go modules](#go-modules)
    + [`dep` dependency manager](#dep-dependency-manager)
- [Getting started](#getting-started)
  * [Authentication](#authentication)
    + [IAM authentication](#iam-authentication)
    + [Basic authentication](#basic-authentication)
  * [Code examples](#code-examples)
    + [1. Retrieve information from an existing database](#1-retrieve-information-from-an-existing-database)
    + [2. Create your own database and add a document](#2-create-your-own-database-and-add-a-document)
    + [3. Update your previously created document](#3-update-your-previously-created-document)
    + [4. Delete your previously created document](#4-delete-your-previously-created-document)
- [Error handling](#error-handling)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Further resources](#further-resources)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

</details>

## Overview

The IBM Cloudant Go SDK allows developers to programmatically interact
with [Cloudant](https://cloud.ibm.com/apidocs/cloudant) with the help of
`cloudantv1` package.

## Features
The purpose of this Go SDK is to wrap most of the HTTP request APIs provided by Cloudant and
supply other functions to ease the usage of Cloudant.
This SDK should make life easier for programmers to do what’s really important for them: develop.

Reasons why you should consider using Cloudant SDK for Go in your project:

- Supported by IBM Cloudant.
- Includes all the most popular and latest supported endpoints for applications.
- Handles the authentication.
- HTTP2 support for higher performance connections to IBM Cloudant.
- Familiar user experience of IBM Cloud SDKs.
- Perform requests synchronously
- Safe for concurrent use by multiple goroutines

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.13 or above.

## Installation

The current version of this SDK: 0.0.26

There are a few different ways to download and install the Cloudant Go SDK project for use by your Go application:

#### `go get` command
Use this command to download and install the SDK to allow your Go application to
use it:

```terminal
go get -u github.com/IBM/cloudant-go-sdk/cloudantv1
```

#### Go modules
If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
  "github.com/IBM/cloudant-go-sdk/cloudantv1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

#### `dep` dependency manager
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```terminal
[[constraint]]
  name = "github.com/IBM/cloudant-go-sdk"
  version = "0.0.26"

```

then run `dep ensure`.

## Getting started

### Authentication

This library requires some of your [Cloudant service credentials](https://cloud.ibm.com/docs/Cloudant?topic=cloudant-creating-an-ibm-cloudant-instance-on-ibm-cloud#locating-your-service-credentials) to authenticate with your account.
1. `IAM`, `BASIC` or `NOAUTH` **authentication type**.
    1. [*IAM authentication*](#iam-authentication) is highly recommended when your back-end database server is [**Cloudant**](https://cloud.ibm.com/docs/Cloudant?topic=cloudant-ibm-cloud-identity-and-access-management-iam-). This authentication type requires a server-generated `apikey` instead of a user-given password.
    1. [*Basic* (or legacy) *authentication*](#basic-authentication) is a fallback for both [Cloudant](https://cloud.ibm.com/docs/services/Cloudant/api?topic=cloudant-authentication#basic-authentication) and [Apache CouchDB](https://docs.couchdb.org/en/stable/api/server/authn.html#basic-authentication) back-end database servers. This authentication type requires the good old `username` and `password` credentials.
    1. *Noauth* authentication does not need any credentials. Note that this authentication type will only work for queries against a database with read access for everyone.
1. The service `url`

You have to add these properties as your **environment variables**, because
some examples that follow assume that these variables are set.
To learn more about authentication configuration see the related documentation in the
[Cloudant API docs](https://cloud.ibm.com/apidocs/cloudant#authentication?code=go) or in the
[general SDK usage information](https://github.com/IBM/ibm-cloud-sdk-common#authentication).

#### IAM authentication

For Cloudant *IAM authentication* set the following environmental variables by replacing `<url>` and `<apikey>` with your proper [service credentials](https://cloud.ibm.com/docs/Cloudant?topic=cloudant-creating-an-ibm-cloudant-instance-on-ibm-cloud#locating-your-service-credentials). There is no need to set `CLOUDANT_AUTH_TYPE` to `IAM` because it is the default.
```vim
CLOUDANT_URL=<url>
CLOUDANT_APIKEY=<apikey>
```

#### Basic authentication

For *Basic authentication* set the following environmental variables by replacing `<url>`, `<username>` and `<password>` with your proper [service credentials](https://cloud.ibm.com/docs/Cloudant?topic=cloudant-creating-an-ibm-cloudant-instance-on-ibm-cloud#locating-your-service-credentials).

```vim
CLOUDANT_AUTH_TYPE=BASIC
CLOUDANT_URL=<url>
CLOUDANT_USERNAME=<username>
CLOUDANT_PASSWORD=<password>
```

**Note**: We recommend using [IAM](#iam-authentication) for Cloudant and
[Basic](#basic-authentication) for CouchDB authentication.

### Code examples

#### 1. Retrieve information from an existing database

This example code gathers some information about an existing database hosted on the https://examples.cloudant.com/ service `url`.
To do this, you need to extend your environment variables with the *service url*
and *authentication type* to use `NOAUTH` authentication while reaching the
`animaldb` database. This step is necessary for the SDK to distinguish the
`EXAMPLES` custom service name from the default service name which is
`CLOUDANT`.

```env
EXAMPLES_URL=https://examples.cloudant.com
EXAMPLES_AUTH_TYPE=NOAUTH
```

Once the environment variables are set, you can try out the code examples.

[embedmd]:# (examples/src/get_info_from_existing_database/get_info_from_existing_database.go /package main/ $)
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func main() {
	// 1. Create a Cloudant client with "EXAMPLES" service name ============
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{
			ServiceName: "EXAMPLES",
		},
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
	// 3. Get database information for "animaldb" ==========================
	dbName := "animaldb"
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
	// 5. Get zebra document out of the database by document id ============
	documentAboutZebraResult, _, err := client.GetDocument(
		client.NewGetDocumentOptions(
			dbName,
			"zebra",
		),
	)
	if err != nil {
		panic(err)
	}
	// 6. Print out the Document content ===================================
	aboutZebraBuffer, _ := json.MarshalIndent(documentAboutZebraResult, "", "  ")
	fmt.Println(string(aboutZebraBuffer))
}
```

The result of the code is similar to the following output.

[embedmd]:# (examples/output/get_info_from_existing_database.txt)
```txt
Server Version: 2.1.1
Document count in "animaldb" database is 11.
{
  "_id": "zebra",
  "_rev": "3-750dac460a6cc41e6999f8943b8e603e",
  "class": "mammal",
  "diet": "herbivore",
  "max_length": 2.5,
  "max_weight": 387,
  "min_length": 2,
  "min_weight": 175,
  "wiki_page": "http://en.wikipedia.org/wiki/Plains_zebra"
}
```

#### 2. Create your own database and add a document

Now comes the exciting part of creating your own `orders` database and adding a document
about *Bob Smith* with your own [IAM](#iam-authentication) or [Basic](#basic-authentication) service credentials.

[embedmd]:# (examples/src/create_db_and_doc/create_db_and_doc.go /package main/ $)
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
	// 3.1. Create a document object with "example" id
	exampleDocID    := "example"
	exampleDocument := cloudantv1.Document{
		ID: &exampleDocID,
	}

	// 3.2. Add "name" and "joined" fields to the document
	exampleDocument.SetProperty("name", "Bob Smith")
	exampleDocument.SetProperty("joined", "2019-01-24T10:42:99.000Z")

	// 3.3. Set the options to get the document out of the database if it exists
	getDocumentOptions := client.NewGetDocumentOptions(
		exampleDbName,
		exampleDocID,
	)

	// 3.4. Check the document existence in the database
	documentInfo, getDocumentResponse, err := client.GetDocument(
		getDocumentOptions,
	)
	if err != nil {
		if getDocumentResponse.StatusCode == 404 {
			// Document does not exist in database
		} else {
			panic(err)
		}
	}

	// 3.5. If it previously existed in the database, set revision of exampleDocument to the latest
	if documentInfo != nil {
		exampleDocument.Rev = documentInfo.Rev
		fmt.Printf("The document revision for \"%s\"  is set to \"%s\".\n",
			exampleDocID,
			*documentInfo.Rev)
	}

	// 3.6. Save the document in the database
	postDocumentOption := client.NewPostDocumentOptions(
		exampleDbName,
	).SetDocument(&exampleDocument)

	postDocumentResult, _, err := client.PostDocument(postDocumentOption)
	if err != nil {
		panic(err)
	}

	// 3.7. Keep track of the revision number from the `example` document object
	exampleDocument.Rev = postDocumentResult.Rev

	// 3.8. Print out the document content
	exampleDocumentContent, _ := json.MarshalIndent(exampleDocument, "", "  ")
	fmt.Printf("You have created the document:\n%s\n", string(exampleDocumentContent))
}
```

The result of the code is similar to the following output.

[embedmd]:# (examples/output/create_db_and_doc.txt)
```txt
"orders" database created.
You have created the document:
{
  "_id": "example",
  "_rev": "1-2c3b9502ed7c4a41d35c92bcf734869c",
  "joined": "2019-01-24T10:42:99.000Z",
  "name": "Bob Smith"
}
```

#### 3. Update your previously created document

**Note**: this example code assumes that you have created both the `orders`
database and the `example` document by [running this previous example
code](#2-create-your-own-database-and-add-a-document) successfully, otherwise
you get the `Cannot update document because either "orders" database or "example" document was not found.` message.

[embedmd]:# (examples/src/update_doc/update_doc.go /package main/ $)
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
		).SetDocument(document)

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
```

The result of the code is similar to the following output.

[embedmd]:# (examples/output/update_doc.txt)
```txt
You have updated the document:
{
  "_id": "example",
  "_rev": "2-6cc06e9484d776322f7e697c03fb23f7",
  "address": "19 Front Street, Darlington, DL5 1TY",
  "name": "Bob Smith"
}
```

#### 4. Delete your previously created document

**Note**: this example code assumes that you have created both the `orders`
database and the `example` document by [running this previous example
code](#2-create-your-own-database-and-add-a-document) successfully, otherwise
you get the `Cannot delete document because either "orders" database or "example" document was not found.` message.

[embedmd]:# (examples/src/delete_doc/delete_doc.go /package main/ $)
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
				*document.ID,
			).SetRev(*document.Rev),
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

The result of the code is the following output.

[embedmd]:# (examples/output/delete_doc.txt)
```txt
You have deleted the document.
```

## Error handling

For sample code on handling errors, please see [Cloudant API docs](https://cloud.ibm.com/apidocs/cloudant?code=go#error-handling)

## Using the SDK

For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md)

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/cloudant-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has
already reported the problem.

## Further resources
* [Cloudant API docs](https://cloud.ibm.com/apidocs/cloudant?code=go): API examples for Cloudant Go SDK.
* [Cloudant docs](https://cloud.ibm.com/docs/services/Cloudant?topic=cloudant-overview#overview): The official documentation page for Cloudant.
* [Cloudant Learning Center](https://developer.ibm.com/clouddataservices/docs/compose/cloudant/): The official learning center with several useful videos which help you to use Cloudant successfully.
* [Cloudant blog](https://blog.cloudant.com/): Many useful articles how to
  optimize Cloudant for common problems.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
