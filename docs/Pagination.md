# Pagination

<details open>
<summary>Table of Contents</summary>

<!-- toc -->
- [Introduction](#introduction)
- [Limitations](#limitations)
- [Capacity considerations](#capacity-considerations)
- [Available operations](#available-operations)
- [Creating a pagination](#creating-a-pagination)
  * [Initialize the service](#initialize-the-service)
  * [Set the options](#set-the-options)
  * [Create the pagination](#create-the-pagination)
- [Using pagination](#using-pagination)
  * [Iterate pages](#iterate-pages)
  * [Iterate rows](#iterate-rows)
  * [Pager](#pager)
    + [Get each page from a pager](#get-each-page-from-a-pager)
    + [Get all results from a pager](#get-all-results-from-a-pager)
</details>

## Introduction

The pagination feature accepts options for a single operation and automatically
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

## Limitations

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

## Capacity considerations

Pagination can make many requests rapidly from a single program call.

For IBM Cloudant take care to ensure you have appropriate plan capacity
in place to avoid consuming all the permitted requests.
If there is no remaining plan allowance and retries are not enabled or insufficient
then a `429 Too Many Requests` error occurs.

## Available operations

Pagination is available for these operations:
* Query all documents [global](https://cloud.ibm.com/apidocs/cloudant?code=go#postalldocs)
  and [partitioned](https://cloud.ibm.com/apidocs/cloudant?code=go#postpartitionalldocs)
  * [Global all documents examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.10/test/examples/src/pagination/all_docs_pagination.go)
  * [Partitioned all documents examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.10/test/examples/src/pagination/partition_all_docs_pagination.go)
* Query all [design documents](https://cloud.ibm.com/apidocs/cloudant?code=go#postdesigndocs)
  * [Design documents examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.10/test/examples/src/pagination/design_docs_pagination.go)
* Query with selector syntax [global](https://cloud.ibm.com/apidocs/cloudant?code=go#postfind)
  and [partitioned](https://cloud.ibm.com/apidocs/cloudant?code=go#postpartitionfind)
  * [Global find selector query examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.10/test/examples/src/pagination/find_pagination.go)
  * [Partitioned find selector query examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.10/test/examples/src/pagination/partition_find_pagination.go)
* Query a search index [global](https://cloud.ibm.com/apidocs/cloudant?code=go#postsearch)
  and [partitioned](https://cloud.ibm.com/apidocs/cloudant?code=go#postpartitionsearch)
  * [Global search examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.10/test/examples/src/pagination/search_pagination.go)
  * [Partitioned search examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.10/test/examples/src/pagination/partition_search_pagination.go)
* Query a MapReduce view [global](https://cloud.ibm.com/apidocs/cloudant?code=go#postview)
  and [partitioned](https://cloud.ibm.com/apidocs/cloudant?code=go#postpartitionview)
  * [Global view examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.10/test/examples/src/pagination/view_pagination.go)
  * [Partitioned view examples](https://github.com/IBM/cloudant-go-sdk/tree/v0.10.10/test/examples/src/pagination/partition_view_pagination.go)

The examples presented in this `README` are for all documents in a partition.
The links in the list are to equivalent examples for each of the other available operations.

## Creating a pagination

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

### Initialize the service

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

### Set the options

<details open>
<summary>Go:</summary>

```go
// Setup options
opts := service.NewPostPartitionAllDocsOptions("events", "ns1HJS13AMkK")
opts.SetLimit(50)
```

</details>

### Create the pagination

<details open>
<summary>Go:</summary>

```go
// Create pagination
// pagination can be reused without side-effects as a factory for iterators or pagers
// options are fixed at pagination creation time
pagination := features.NewAllDocsPagination(service, opts)
```

</details>

## Using pagination

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

### Iterate pages

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

### Iterate rows

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

### Pager

The pager style is similar to other [IBM Cloud SDKs](https://github.com/IBM/ibm-cloud-sdk-common?tab=readme-ov-file#pagination).
Users familiar with that style of pagination may find using them preferable
to the native language style iterators.

In the Cloudant SDKs these pagers are single use and traverse the complete set of pages once and only once.
After exhaustion they cannot be re-used, simply create a new one from the pagination factory if needed.

Pagers are only valid for one of either page at a time or getting all results.
For example, calling for the next page then calling for all results causes an error.

#### Get each page from a pager

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

#### Get all results from a pager

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
