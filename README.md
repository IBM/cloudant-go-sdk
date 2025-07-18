[![Build Status](https://github.com/IBM/cloudant-go-sdk/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/IBM/cloudant-go-sdk/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/v/release/IBM/cloudant-go-sdk?include_prereleases&sort=semver)](https://github.com/IBM/cloudant-go-sdk/releases/latest)
[![Docs](https://img.shields.io/static/v1?label=Godoc&message=latest&color=blue)](https://pkg.go.dev/github.com/IBM/cloudant-go-sdk)

# IBM Cloudant Go SDK Version 0.10.4

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

<!-- toc -->
- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  * [`go get` command](#go-get-command)
  * [Go modules](#go-modules)
- [Using the SDK](#using-the-sdk)
  * [Authentication](#authentication)
    + [Authentication with environment variables](#authentication-with-environment-variables)
      - [IAM API key authentication](#iam-api-key-authentication)
      - [IAM Trusted profile (container) authentication](#iam-trusted-profile-container-authentication)
      - [IAM Trusted profile (VPC) authentication](#iam-trusted-profile-vpc-authentication)
      - [IAM Trusted profile (assume identity) authentication](#iam-trusted-profile-assume-identity-authentication)
      - [Session cookie authentication](#session-cookie-authentication)
    + [Authentication with external configuration](#authentication-with-external-configuration)
    + [Programmatic authentication](#programmatic-authentication)
  * [Automatic retries](#automatic-retries)
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
  * [Changes feed follower](#changes-feed-follower)
    + [Introduction](#introduction)
    + [Modes of operation](#modes-of-operation)
    + [Configuring the changes follower](#configuring-the-changes-follower)
    + [Error suppression](#error-suppression)
    + [Follower operation](#follower-operation)
    + [Checkpoints](#checkpoints)
    + [Code examples](#code-examples-1)
      - [Initializing a changes follower](#initializing-a-changes-follower)
      - [Starting the changes follower](#starting-the-changes-follower)
        * [Start mode for continuous listening](#start-mode-for-continuous-listening)
        * [Start mode for one-off fetching](#start-mode-for-one-off-fetching)
      - [Processing changes](#processing-changes)
        * [Process continuous changes](#process-continuous-changes)
        * [Process one-off changes](#process-one-off-changes)
      - [Stopping the changes follower](#stopping-the-changes-follower)
  * [Pagination (beta)](#pagination-beta)
    + [Introduction](#introduction-1)
    + [Limitations](#limitations)
    + [Capacity considerations](#capacity-considerations)
    + [Available operations](#available-operations)
    + [Creating a pagination](#creating-a-pagination)
      - [Initialize the service](#initialize-the-service)
      - [Set the options](#set-the-options)
      - [Create the pagination](#create-the-pagination)
    + [Using pagination](#using-pagination)
      - [Iterate pages](#iterate-pages)
      - [Iterate rows](#iterate-rows)
      - [Pager](#pager)
        * [Get each page from a pager](#get-each-page-from-a-pager)
        * [Get all results from a pager](#get-all-results-from-a-pager)
- [Questions](#questions)
- [Issues](#issues)
- [Versioning and LTS support](#versioning-and-lts-support)
- [Open source at IBM](#open-source-at-ibm)
- [Contributing](#contributing)
- [License](#license)

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
- Built-in [Changes feed follower](#changes-feed-follower)
- Built-in [Pagination](#pagination-beta) (beta)
- HTTP2 support for higher performance connections to IBM Cloudant.
- Perform requests synchronously.
- Safe for concurrent use by multiple goroutines.
- Transparently compresses request and response bodies.

## Prerequisites

- A
  [Cloudant](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-getting-started-with-cloudant)
  service instance or a
  [CouchDB](https://docs.couchdb.org/en/latest/install/index.html)
  server.
- Go version 1.23 or 1.24.

## Installation

The current version of this SDK: 0.10.4

There are a few different ways to download and add the
Cloudant Go SDK project for use by your Go application:

### `go get` command

Use this command to download and add the SDK to allow your Go application to
use it:

```terminal
go get -u github.com/IBM/cloudant-go-sdk/cloudantv1@v0.10.4
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

## Using the SDK

For fundamental SDK usage information and config options, please see the common [IBM Cloud SDK](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md) documentation.

This library requires configuration with a service URL and
[Cloudant service credentials](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-locating-your-service-credentials) to authenticate with your
account.

There are several ways to **set** these authentication properties:

1. As [environment variables](#authentication-with-environment-variables)
2. The [programmatic approach](#programmatic-authentication)
3. With an [external credentials file](#authentication-with-external-configuration)

The following section describes the different authentication types and provides environment variable examples.
Examples for other configuration methods are available by following the provided links.

### Authentication


This library requires credentials to authenticate with IBM Cloudant. These credentials may be:
* IBM Cloud IAM credentials (with authentication types `CONTAINER`, `VPC`, `IAMASSUME` and `IAM`)
  * [IBM Cloud account](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-managing-access-for-cloudant#introduction-iam-ai) user, service ID or trusted profile credentials
    that have access granted to the IBM Cloud Cloudant resource instance.
  * [IBM Cloudant service credentials](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-locating-your-service-credentials) generated by the IBM Cloud Cloudant resource instance.
* Username and password credentials (with authentication types `COUCHDB_SESSION` and `BASIC`)
  * [IBM Cloudant service credentials](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-locating-your-service-credentials) generated for an IBM Cloud Cloudant resource instance not configured as `IAM only`.
  * IBM Cloudant [legacy credentials](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-work-with-your-account#basic-authentication) (username and password) for instances not in IBM Cloud.
  * IBM Cloudant [legacy API keys](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-work-with-your-account#api-keys).

| Authentication type | Recommended for | `AUTH_TYPE` | Description |
| --- | --- | --- | --- |
| IAM Trusted Profiles (compute resource [container](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#container-authentication)) | Cloudant<BR>(SDK running in IBM Cloud IKS) | `CONTAINER` | Obtains a compute resource (CR) token from the container.<BR>Exchanges the CR token for an IAM `access_token`.<BR>Adds an `Authorization: Bearer <access_token>` header to each HTTP request.<BR>Automatically renews the access token when needed. |
| IAM Trusted Profiles (compute resource [VPC](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#vpc-instance-authentication)) | Cloudant<BR>(SDK running in IBM Cloud VPC) | `VPC` | Obtains an identity token from the VPC instance metadata.<BR>Exchanges the identity token for an IAM `access_token`.<BR>Adds an `Authorization: Bearer <access_token>` header to each HTTP request.<BR>Automatically renews the access token when needed. |
| IAM Trusted Profiles ([assume identity](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#identity-and-access-management-iam-authentication-grant-type-assume)) | Cloudant | `IAMASSUME` | Exchanges an IAM API key for an IAM `access_token` (same as `IAM` auth type).<BR>Uses that initial token to obtain a second `access_token` from IAM with the assumed identity information.<BR>Adds an `Authorization: Bearer <access_token>` header to each HTTP request.<BR>Automatically renews the access token when needed. |
| [IAM API key](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#identity-and-access-management-iam-authentication-grant-type-apikey) | Cloudant | `IAM` | Exchanges an IAM API key for an IAM `access_token`.<BR>Adds an `Authorization: Bearer <access_token>` header to each HTTP request.<BR>Automatically renews the access token when needed. |
| [Session cookie](#session-cookie-authentication) | [Cloudant](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-work-with-your-account#cookie-authentication)<BR>(legacy credentials & instances without IAM)<BR><BR>[Apache CouchDB](https://docs.couchdb.org/en/stable/api/server/authn.html#cookie-authentication) | `COUCHDB_SESSION` | Exchanges credentials with `/_session` endpoint to retrieve a cookie.<BR>Adds `Cookie` header and content to each HTTP request.<BR>Automatically renews session when needed. |
| [Bearer token](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#bearer-token-authentication) | [Apache CouchDB](https://docs.couchdb.org/en/stable/api/server/authn.html#jwt-authentication)<BR>(using JWT authentication) | `BEARERTOKEN` | Adds an `Authorization: Bearer <token>` to each HTTP request.<BR>No token management or renewal.<BR>Also compatible with IAM access tokens managed independently of the SDK. |
| [Basic](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#basic-authentication) | [Apache CouchDB](https://docs.couchdb.org/en/stable/api/server/authn.html#basic-authentication)<BR>(if cookies are not enabled) | `BASIC` | Adds an `Authorization: Basic <encoded username and password>` header to each HTTP request. |
| [None](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md#no-auth-authentication) | - | `NOAUTH` | Note that this authentication type only works for operations against a database allowing access for unauthenticated users. |

The default authentication type for the SDK is `CONTAINER` unless supplying `APIKEY` configuration, which changes the default authentication type to `IAM`.

#### Authentication with environment variables

The default service name is `CLOUDANT` so these examples use `CLOUDANT_` prefixed names.

Any custom service name prefix is valid, provided it matches the name used to instantiate the SDK client
and applied to all configuration options.

##### IAM API key authentication

For Cloudant *IAM API key authentication*, set the following environmental variables by
replacing the `<url>` and `<apikey>` with your proper
[service credentials](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-locating-your-service-credentials). There is no need to set
`CLOUDANT_AUTH_TYPE` to `IAM` because it is the default when supplying an `APIKEY`.

```bash
CLOUDANT_URL=<url>
CLOUDANT_APIKEY=<apikey>
```

##### IAM Trusted profile (container) authentication

For Cloudant *IAM Trusted profile compute resource container authentication*, set the following environmental variables by
replacing the `<url>` and `<id>` with your values. There is no need to set
`CLOUDANT_AUTH_TYPE` to `CONTAINER` because it is the default.

```bash
CLOUDANT_URL=<url>
CLOUDANT_IAM_PROFILE_ID=<id>
```

Alternatives to `CLOUDANT_IAM_PROFILE_ID`:
* `CLOUDANT_IAM_PROFILE_NAME`

##### IAM Trusted profile (VPC) authentication

For Cloudant *IAM Trusted profile compute resource vpc authentication*, set the following environmental variables by
replacing the `<url>` and `<id>` with your values.

```bash
CLOUDANT_AUTH_TYPE=VPC
CLOUDANT_URL=<url>
CLOUDANT_IAM_PROFILE_ID=<id>
```

Alternatives to `CLOUDANT_IAM_PROFILE_ID`:
* `CLOUDANT_IAM_PROFILE_CRN`
* No profile information (uses the default trusted profile linked to the compute resource)

##### IAM Trusted profile (assume identity) authentication

For Cloudant *IAM Trusted profile assume authentication*, set the following environmental variables by
replacing the `<url>` and `<id>` with your values.

```bash
CLOUDANT_AUTH_TYPE=IAMASSUME
CLOUDANT_URL=<url>
CLOUDANT_IAM_PROFILE_ID=<id>
```

Alternatives to `CLOUDANT_IAM_PROFILE_ID`:
* `CLOUDANT_IAM_PROFILE_CRN`
* `CLOUDANT_IAM_PROFILE_NAME` *and* `CLOUDANT_IAM_ACCOUNT_ID` (ID of the account that contains the named trusted profile)

##### Session cookie authentication

For `COUCHDB_SESSION` authentication, set the following environmental variables
by replacing the `<url>`, `<username>` and `<password>` with your proper
[service credentials](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-locating-your-service-credentials).

```bash
CLOUDANT_AUTH_TYPE=COUCHDB_SESSION
CLOUDANT_URL=<url>
CLOUDANT_USERNAME=<username>
CLOUDANT_PASSWORD=<password>
```

#### Authentication with external configuration

For more information about using an external configuration file, see the related documentation in
[Cloudant API docs](https://cloud.ibm.com/apidocs/cloudant?code=go#authentication-with-external-configuration),
or the
[general SDK usage information](https://github.com/IBM/ibm-cloud-sdk-common#using-external-configuration).

#### Programmatic authentication

To learn more about how to use programmatic authentication, see the related
documentation in the
[Cloudant API docs](https://cloud.ibm.com/apidocs/cloudant?code=go#programmatic-authentication)
or in the
[Go SDK Core document](https://github.com/IBM/go-sdk-core/blob/main/Authentication.md) about authentication.

### Automatic retries

The SDK supports a generalized retry feature that can automatically retry on common errors.

The [automatic retries](https://github.com/IBM/ibm-cloud-sdk-common#automatic-retries) section has details on how to enable the retries with default values and customize the retries programmatically or with external configuration.

### Request timeout configuration

A 6m request timeout, which includes a 30s connect timeout, is set by default. Note that this also affects changes feed requests, regardless of a timeout set on `PostChangesOptions`. Be sure to set a request timeout appropriate to your application usage and environment.
The [request timeout](https://github.com/IBM/ibm-cloud-sdk-common#configuring-request-timeouts) section contains details on how to change the value.

**Note:** System settings may take precedence over configured timeout values.

### Code examples

The following code examples
[authenticate with the environment variables](#authentication-with-environment-variables).

#### 1. Create a database and add a document

**Note:** This example code assumes that `orders` database does not exist in your account.

This example code creates `orders` database and adds a new document "example"
into it. To connect, you must set your environment variables with
the *service url*, *authentication type* and *authentication credentials*
of your Cloudant service.

Cloudant environment variable naming starts with a *service name* prefix that identifies your service.
By default, this is `CLOUDANT`, see the settings in the
[authentication with environment variables section](#authentication-with-environment-variables).

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

	// Keeping track of the revision number of the document object
	// is necessary for further UPDATE/DELETE operations:
	exampleDocument.Rev = createDocumentResponse.Rev

	// Print out the document content
	exampleDocumentContent, _ := json.MarshalIndent(exampleDocument, "", "  ")
	fmt.Printf("You have created the document:\n%s\n", string(exampleDocumentContent))
}
```

When you run the code, you see a result similar to the following output.

```text
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

#### 4. Delete your previously created document

**Note**: This example code assumes that you have created both the `orders`
database and the `example` document by
[running the previous example code](#1-create-a-database-and-add-a-document)
successfully. Otherwise, the following error message occurs, "Cannot delete document because either 'orders'
database or 'example' document was not found."

<details>
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

#### Further code examples

For a complete list of code examples, see the [examples directory](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/examples#examples-for-go).

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

### Changes feed follower

#### Introduction

The SDK provides a changes feed follower utility.
This helper utility connects to the `_changes` endpoint and returns the individual change items.
It removes some complexity of using the `_changes` endpoint by setting some options automatically
and providing error suppression and retries.

*Tip: the changes feed often does not meet user expectations or assumptions.*

Consult the [Cloudant changes feed FAQ](https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-faq-using-changes-feed)
to get a better understanding of the limitations and suitable use-cases before using the changes feed in your application.

#### Modes of operation

There are two modes of operation:
* Start mode
  * Fetches the changes from the supplied `since` sequence (in this mode follower defaults to reading the feed from `now`).
  * Fetches all available changes and then continues listening for new changes indefinitely unless encountering an end condition.
  * An example use case for this mode is event driven workloads.
* Start one-off mode
  * Fetches the changes from the supplied `since` sequence (in this mode follower defaults to reading the feed from the beginning).
  * Fetches all available changes and then stops when either there are no further changes pending or encountering an end condition.
  * An example use case for this mode is ETL style workloads.

#### Configuring the changes follower

The SDK's model of changes feed options is also used to configure the follower.
However, a subset of the options used internally by the follower implementation are invalid.
Supplying these options when instantiating the follower causes an error.
The invalid options are:
* `descending`
* `feed`
* `heartbeat`
* `lastEventId` - use `since` instead
* `timeout`
* Follower permits only the value `_selector` for the `filter` option. This restriction is because selector
  based filters perform better than JavaScript backed filters. Configuring a non-selector based filter
  causes the follower to error.

Note that the `limit` parameter terminates the follower at the given number of changes in either
operating mode.

The changes follower requires the client to have HTTP timeouts of at least 1 minute and errors during
instantiation if it is insufficient. The default client configuration has sufficiently long timeouts.

For use-cases where these configuration limitations are too restrictive then write code to use the SDK's
[POST `_changes` API](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/examples#postchanges) instead of the follower.

#### Error suppression

By default, the changes follower suppresses transient errors indefinitely and attempts to run to completion or listen forever as
dictated by the operating mode.
For applications where that is not desirable configure the optional error tolerance duration. This controls the time since
the last successful response that the follower suppresses transient errors. An example usage is an application grace period
before reporting an error and requiring intervention.

There are some additional points to consider for error suppression:
* Errors considered terminal, for example, the database not existing or invalid credentials are never suppressed and error immediately.
* The error suppression duration is not guaranteed to fire immediately after lapsing and is a minimum suppression time.
* The changes follower backs-off between retries and as such may remain paused for a short while after the transient errors have resolved.
* If the underlying SDK client used to initialize the follower also has retries configured then suppression of errors may last
  significantly longer than the follower's configured error tolerance duration depending on the specific options.

#### Follower operation

For both modes:
* The end conditions are:
  * A terminal error (HTTP codes `400`, `401`, `403` `404`).
  * Transient errors occur for longer than the error tolerance duration. Transient errors are all other HTTP status codes and connection errors.
  * The number of changes received reaches the configured `limit`.
  * The application calls stop to terminate the feed early.

As is true for the `_changes` endpoint change items have *at least once* delivery and callers may receive
an individual item multiple times. When using the follower change items may repeat even within a limited
number of changes (that is using the `limit` option) this is a minor difference from using `limit` on the HTTP native API.

The follower is not optimized for some use cases and it is not recommended to use it in cases where:
* Setting `include_docs` and larger document sizes (for example > 10 kiB).
* The volume of changes is very high (if the rate of changes in the database exceeds the follower's rate of pulling them it can never catch-up).

In these use-cases use the SDK's [POST `_changes` API](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/examples#postchanges)
for  specific control over the number of change requests made and the content size of the responses.

#### Checkpoints

The changes follower does not checkpoint since it has no information about whether the consuming application
has processed a change item after delivery. It is the application developer's responsibility
to store the sequence IDs to have appropriate checkpoints and to re-initialize the follower with the required
`since` value after, for example, the application restarts.

The frequency and conditions for checkpoints are application specific and some applications may be tolerant
of dropped changes. This section provides only general guidance on how to avoid missing changes.

To guarantee processing of all changes do not persist the sequence ID from a change item until *after*
the processing of the change item by the application has completed. As indicated previously change item
delivery is *at least once* so application code must be able to handle repeated changes already. It is
preferable to restart from an older `since` value and receive changes again than risk missing them.

The sequence IDs are available on each change item by default. However, the server omits sequence IDs from
some change items when using the `seq_interval` configuration option.
Infrequent sequence IDs may improve performance by reducing the amount of data transfer and server load,
but the tradeoff is repeating more changes if it is necessary to resume the changes follower.

Take extreme care persisting sequences if choosing to process change items in parallel as there
is a considerable risk of missing changes on a restart if the recorded sequence is out of order.

#### Code examples

##### Initializing a changes follower
```go
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
```

##### Starting the changes follower

###### Start mode for continuous listening
```go
package main

import (
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

	follower, err := features.NewChangesFollower(client, postChangesOptions)
	if err != nil {
		panic(err)
	}

	changesCh, err := follower.Start()
	if err != nil {
		panic(err)
	}

	// Note: changesCh channel will not do anything until it is read from.
	// Create a range loop to iterate over the flow of the changes items
}
```

###### Start mode for one-off fetching
```go
package main

import (
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

	follower, err := features.NewChangesFollower(client, postChangesOptions)
	if err != nil {
		panic(err)
	}

	changesCh, err = follower.StartOneOff()
	if err != nil {
		panic(err)
	}

	// Note: the returned channel will not do anything until it is read
	// Create a range loop to iterate over the flow of the change items
}
```

##### Processing changes

###### Process continuous changes
```go
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

	changesCh, err := follower.Start()
	if err != nil {
		panic(err)
	}

	for changesItem := range changesCh {
		// changes item returns an error on failed requests
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
		// write seq to persistent storage for use as since
		// if required to resume later, e.g. yourAppPersistenceWriteFunc(seq)
		// keep processing changes until the application is terminated
		// or some other stop condition is reached
	}

	// Note: the returned channel continues to listen for changes
	// indefinitely. In other words iterating it is blocking,
	// code here will be unreachable until break from the range
	// or another stop condition is reached.
	// For long running followers careful consideration
	// should be made of where to call stop on the follower.
}
```

###### Process one-off changes
```go
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
```

##### Stopping the changes follower
```go
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
```

### Pagination (beta)

#### Introduction

The pagination feature (currently beta) accepts options for a single operation and automatically
creates the multiple requests to the server necessary to page through the results a fixed number at a time.

Pagination is a best-practice to break apart large queries into multiple server requests.
This has a number of advantages:
* Keeping requests within server imposed limits, for example
  * `200` max results for text search
  * `2000` max results for partitioned queries
* Fetching only the necessary data, for example
  * User finds required result on first page, no need to continue fetching results
* Reducing the duration of any individual query
  * Reduce risk of query timing out on the server
  * Reduce risk of network request timeouts

#### Limitations

Limitations of pagination:
* Forward only, no backwards paging
* Limitations on `_all_docs` and `_design_docs` operations
  * No pagination for `key` option.
    There is no need to paginate as IDs are unique and this returns only a single row.
    This is better achieved with a single document get request.
  * No pagination for `keys` option.
* Limitations on `_view` operations
  * No pagination for `key` option. Pass the same `key` as a start and end key instead.
  * No pagination for `keys` option.
  * Views that emit multiple identical keys (with the same or different values)
    from the same document cannot paginate if those key rows with the same ID
    span a page boundary.
    The pagination feature detects this condition and an error occurs.
    It may be possible to workaround using a different page size.
* Limitations on `_search` operations
  * No pagination of grouped results.
  * No pagination of faceted `counts` or `ranges` results.

#### Capacity considerations

Pagination can make many requests rapidly from a single program call.

For IBM Cloudant take care to ensure you have appropriate plan capacity
in place to avoid consuming all the permitted requests.
If there is no remaining plan allowance and retries are not enabled or insufficient
then a `429 Too Many Requests` error occurs.

#### Available operations

Pagination is available for these operations:
* Query all documents [global](https://cloud.ibm.com/apidocs/cloudant?code=go#postalldocs)
  and [partitioned](https://cloud.ibm.com/apidocs/cloudant?code=go#postpartitionalldocs)
  * [Global all documents examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/test/examples/src/pagination/all_docs_pagination.go)
  * [Partitioned all documents examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/test/examples/src/pagination/partition_all_docs_pagination.go)
* Query all [design documents](https://cloud.ibm.com/apidocs/cloudant?code=go#postdesigndocs)
  * [Design documents examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/test/examples/src/pagination/design_docs_pagination.go)
* Query with selector syntax [global](https://cloud.ibm.com/apidocs/cloudant?code=go#postfind)
  and [partitioned](https://cloud.ibm.com/apidocs/cloudant?code=go#postpartitionfind)
  * [Global find selector query examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/test/examples/src/pagination/find_pagination.go)
  * [Partitioned find selector query examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/test/examples/src/pagination/partition_find_pagination.go)
* Query a search index [global](https://cloud.ibm.com/apidocs/cloudant?code=go#postsearch)
  and [partitioned](https://cloud.ibm.com/apidocs/cloudant?code=go#postpartitionsearch)
  * [Global search examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/test/examples/src/pagination/search_pagination.go)
  * [Partitioned search examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/test/examples/src/pagination/partition_search_pagination.go)
* Query a MapReduce view [global](https://cloud.ibm.com/apidocs/cloudant?code=go#postview)
  and [partitioned](https://cloud.ibm.com/apidocs/cloudant?code=go#postpartitionview)
  * [Global view examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/test/examples/src/pagination/view_pagination.go)
  * [Partitioned view examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/test/examples/src/pagination/partition_view_pagination.go)

The examples presented in this `README` are for all documents in a partition.
The links in the list are to equivalent examples for each of the other available operations.

#### Creating a pagination

Make a new pagination from a client
and the options for the chosen operation.
Use the `limit` option to configure the page size (default and maximum `200`).

Imports required for these examples:

<details open>
<summary>Go:</summary>

```go
import (
	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/cloudant-go-sdk/features"
)
```

</details>

##### Initialize the service

<details open>
<summary>Go:</summary>

```go
// Initialize service
service, err := cloudantv1.NewCloudantV1UsingExternalConfig(
	&cloudantv1.CloudantV1Options{},
)
if err != nil {
	panic(err)
}
```

</details>

##### Set the options

<details open>
<summary>Go:</summary>

```go
// Setup options
opts := service.NewPostPartitionAllDocsOptions("events", "ns1HJS13AMkK")
opts.SetLimit(50)
```

</details>

##### Create the pagination

<details open>
<summary>Go:</summary>

```go
// Create pagination
// pagination can be reused without side-effects as a factory for iterators or pagers
// options are fixed at pagination creation time
pagination := features.NewAllDocsPagination(service, opts)
```

</details>

#### Using pagination

Once you have a pagination factory there are multiple options available.

* Iterate pages
* Iterate rows
* Get each page from a pager
* Get all results from a pager

All the paging styles produce equivalent results and make identical page requests.
The style of paging to choose depends on the use case requirements
in particular whether to process a page at a time or a row at a time.

The pagination factory is reusable and can repeatedly produce new instances
of the same or different pagination styles for the same operation options.

Here are examples for each paging style.

##### Iterate pages

Iterating pages is ideal for using a range iterator to process a page at a time.

<details open>
<summary>Go:</summary>

```go
// Option: iterate pages
// Ideal for using a for/range loop with each page.
// The iter.Seq2 returned from Pages() has an error as a second element.
for page, err := range pagination.Pages() {
	// Break on err != nil
	// Do something with page
}
```

</details>

##### Iterate rows

Iterating rows is ideal for using a range iterator to process a result row at a time.

<details open>
<summary>Go:</summary>

```go
// Option: iterate rows
// Ideal for using a for/range loop with each row.
// The iter.Seq2 returned from Rows() has an error as a second element.
for row, err := range pagination.Rows() {
	// Break on err != nil
	// Do something with row
}
```

</details>

##### Pager

The pager style is similar to other [IBM Cloud SDKs](https://github.com/IBM/ibm-cloud-sdk-common?tab=readme-ov-file#pagination).
Users familiar with that style of pagination may find using them preferable
to the native language style iterators.

In the Cloudant SDKs these pagers are single use and traverse the complete set of pages once and only once.
After exhaustion they cannot be re-used, simply create a new one from the pagination factory if needed.

Pagers are only valid for one of either page at a time or getting all results.
For example, calling for the next page then calling for all results causes an error.

###### Get each page from a pager

This is useful for calling to retrieve one page at a time, for example,
in a user interface with a "next page" interaction.

If calling for the next page errors, it is valid to call for the next page again
to continue paging.

<details open>
<summary>Go:</summary>

```go
// Option: use pager next page
// For retrieving one page at a time with a method call.
pager, err := pagination.Pager()
if err != nil {
	panic(err)
}
```

</details>

###### Get all results from a pager

This is useful to retrieve all results in a single call.
However, this approach requires sufficient memory for the entire collection of results.
So although it may be convenient for small result sets generally prefer iterating pages
or rows with the other paging styles, especially for large result sets.

If calling for all the results errors, then calling for all the results again restarts the pagination.

<details open>
<summary>Go:</summary>

```go
// Option: use pager all results
// For retrieving all result rows in a single list
allPager, err := pagination.Pager()
if err != nil {
	panic(err)
}
```

</details>

## Questions

If you are having difficulties using this SDK or have a question about the
IBM Cloud services, ask a question on
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues

If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/cloudant-go-sdk/issues).

Before you submit a bug report, search for
[similar issues](https://github.com/IBM/cloudant-go-sdk/issues?q=is%3Aissue) and review the
[KNOWN_ISSUES file](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/KNOWN_ISSUES.md) to verify that your issue hasn't been reported yet.

Please consult the [security policy](https://github.com/IBM/cloudant-go-sdk/security/policy) before opening security related issues.

## Versioning and LTS support

This SDK follows semantic versioning with respect to the definition of user facing APIs.
This means under some circumstances breaking changes may occur within a major or minor version
of the SDK related to changes in supported language platforms.

The SDK is supported on the available LTS releases of the language platform.
The LTS language versions are listed in the prerequisites:
* [LTS versions currently supported by the SDK](https://github.com/IBM/cloudant-go-sdk/#prerequisites)
* [LTS versions for this release of the SDK](#prerequisites)

Incompatible changes from new language versions are not added to the SDK
until they are available in the minimum supported language version.

When language LTS versions move out of support the following will happen:
* Existing SDK releases will continue to run on obsolete language versions, but will no longer be supported.
* The minimum language version supported by the SDK will be updated to the next available LTS.
* New language features may be added in subsequent SDK releases that will cause breaking changes
  if the new releases of the SDK are used with older, now unsupported, language levels.

## Open source at IBM

Find more open source projects on the [IBM GitHub](http://ibm.github.io/) page.

## Contributing

For more information, see [CONTRIBUTING](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/CONTRIBUTING.md).

## License

This SDK is released under the Apache 2.0 license. To read the full text of the license, see [LICENSE](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.4/LICENSE).
