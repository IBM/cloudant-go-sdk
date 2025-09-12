[![Build Status](https://github.com/IBM/cloudant-go-sdk/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/IBM/cloudant-go-sdk/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/v/release/IBM/cloudant-go-sdk?include_prereleases&sort=semver)](https://github.com/IBM/cloudant-go-sdk/releases/latest)
[![Docs](https://img.shields.io/static/v1?label=Godoc&message=latest&color=blue)](https://pkg.go.dev/github.com/IBM/cloudant-go-sdk)

# IBM Cloudant Go SDK Version 0.10.6

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
  * [Automatic retries](#automatic-retries)
  * [Request timeout configuration](#request-timeout-configuration)
  * [Code examples](#code-examples)
  * [Error handling](#error-handling)
  * [Raw IO](#raw-io)
  * [Further resources](#further-resources)
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
- Built-in [Changes feed follower](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/docs/Changes_Follower.md)
- Built-in [Pagination](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/docs/Pagination.md) (beta)
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
- Go version 1.24 or 1.25.

## Installation

The current version of this SDK: 0.10.6

There are a few different ways to download and add the
Cloudant Go SDK project for use by your Go application:

### `go get` command

Use this command to download and add the SDK to allow your Go application to
use it:

```terminal
go get -u github.com/IBM/cloudant-go-sdk/cloudantv1@v0.10.6
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
[Cloudant service credentials][service-credentials] to authenticate with your
account.

There are several ways to **set** these authentication properties:

1. As [environment variables](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/docs/Authentication.md#authentication-with-environment-variables)
2. The [programmatic approach](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/docs/Authentication.md#programmatic-authentication)
3. With an [external credentials file](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/docs/Authentication.md#authentication-with-external-configuration)

The following section describes the different authentication types and provides environment variable examples.
Examples for other configuration methods are available by following the provided links.

### Authentication

Consult the [authentication document](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/docs/Authentication.md)
for comprehensive details of all the available authentication methods and how to configure them with environment settings
or programmatically.

Quick start for Cloudant with an IAM API key:
```sh
CLOUDANT_URL=https://~replaceWithYourUniqueHost~.cloudantnosqldb.appdomain.cloud # use your own Cloudant public or private URL # use your own Cloudant public or private URL
CLOUDANT_APIKEY=a1b2c3d4e5f6f1g4h7j3k6l9m2p5q8s1t4v7x0z3 # use your own IAM API key
```

Quick start for Apache CouchDB with a username/password session:
```sh
CLOUDANT_AUTH_TYPE=COUCHDB_SESSION
CLOUDANT_URL=http://~replaceWithYourUniqueHost~.example:5984 # use your CouchDB URL
CLOUDANT_USERNAME=username # replace with your username
CLOUDANT_PASSWORD=password # replace with your password
```

### Automatic retries

The SDK supports a generalized retry feature that can automatically retry on common errors.

The [automatic retries](https://github.com/IBM/ibm-cloud-sdk-common#automatic-retries) section has details on how to enable the retries with default values and customize the retries programmatically or with external configuration.

### Request timeout configuration

A 6m request timeout, which includes a 30s connect timeout, is set by default. Note that this also affects changes feed requests, regardless of a timeout set on `PostChangesOptions`. Be sure to set a request timeout appropriate to your application usage and environment.
The [request timeout](https://github.com/IBM/ibm-cloud-sdk-common#configuring-request-timeouts) section contains details on how to change the value.

**Note:** System settings may take precedence over configured timeout values.

### Code examples

Quick start example to list all databases (assumes environment variable [authentication](#authentication)):

```go
// section: code imports
import (
  "encoding/json"
  "fmt"

  "github.com/IBM/cloudant-go-sdk/cloudantv1"
)
// section: code
service, err := cloudantv1.NewCloudantV1(
  &cloudantv1.CloudantV1Options{},
)
if err != nil {
  panic(err)
}

getAllDbsOptions := service.NewGetAllDbsOptions()

result, response, err := service.GetAllDbs(getAllDbsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(result, "", "  ")
fmt.Println(string(b))
```

[More tutorial examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/docs/Examples.md) for creating a database
and document create, read, update and delete operations.

For a complete list of code examples, see the [examples directory](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/examples#examples-for-go).

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

The [update document](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/docs/Examples.md#3-update-your-previously-created-document) section
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

- [Cloudant Go SDK feature docs](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/docs)
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
[KNOWN_ISSUES file](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/KNOWN_ISSUES.md) to verify that your issue hasn't been reported yet.

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

For more information, see [CONTRIBUTING](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/CONTRIBUTING.md).

## License

This SDK is released under the Apache 2.0 license. To read the full text of the license, see [LICENSE](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.6/LICENSE).
