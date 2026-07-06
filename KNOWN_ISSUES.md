# Limitations, Restrictions, and Known Issues

## All Cloudant SDKs

### Cloudant (Gen 2) Compatibility

Cloudant Gen 2 is compatible with IAM authentication only.
The SDK authentication types `BASIC` and `COUCHDB_SESSION` do not work with Cloudant Gen 2 instances.

Instances of Cloudant Gen 2 do not provide some configuration and monitoring APIs that were available in Gen 1.

Migrating an application using this SDK from Cloudant Gen 1 to Cloudant Gen 2 requires either removing calls to these operations
or replacing them with alternatives.

This table summarizes the SDK operations that are incompatible with Cloudant Gen 2 instances and the recommended replacement operations.

| Gen 1 operation | Summary | Gen 1 Endpoint | API docs link | Replacement operation for Gen 2 |
|---|---|---|---|---|
| **Authentication and authorization** | | | | |
| `PostApiKeys` | Generates API keys for apps or persons to enable database access | `POST /_api/v2/api_keys` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#postapikeys) | Use IAM authentication |
| `PutCloudantSecurityConfiguration` | Modify only Cloudant related database permissions | `PUT /_api/v2/db/{db}/_security` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#putcloudantsecurity) | Use [database level IAM policies](https://cloud.ibm.com/docs/cloudant-gen2?topic=cloudant-gen2-managing-access-for-cloudant#database-level-iam-policies) |
| **Audit events configuration** | | | | |
`GetActivityTrackerEvents` | Retrieve activity tracking events information | `GET /_api/v2/user/activity_tracker/events` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#getactivitytrackerevents) | `management` events are always enabled, `data` events are not currently available in Gen 2 |
`PostActivityTrackerEvents` | Modify activity tracking events configuration | `POST /_api/v2/user/activity_tracker/events` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#postactivitytrackerevents) | `management` events are always enabled, `data` events are not currently available in Gen 2 |
| **CORS configuration** | | | | |
| `GetCorsInformation` | Retrieve CORS configuration information | `GET /_api/v2/user/config/cors` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#getcorsinformation) | Use the [Platform Services SDK](#using-the-platform-services-sdk)
| `PutCorsConfiguration` | Modify CORS configuration | `PUT /_api/v2/user/config/cors` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#putcorsconfiguration) | Use the [Platform Services SDK](#using-the-platform-services-sdk)
| **Capacity configuration** | | | | |
| `GetCapacityDatabasesInformation` | Retrieve maximum allowed database count | `GET /_api/v2/user/capacity/databases` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#getcapacitydatabasesinformation) | Use the [Platform Services SDK](#using-the-platform-services-sdk)
| `GetCapacityThroughputInformation` | Retrieve provisioned throughput capacity information | `GET /_api/v2/user/capacity/throughput` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#getcapacitythroughputinformation) | Use the [Platform Services SDK](#using-the-platform-services-sdk)
| `PutCapacityThroughputConfiguration` | Update the target provisioned throughput capacity | `PUT /_api/v2/user/capacity/throughput` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#putcapacitythroughputconfiguration) | Use the [Platform Services SDK](#using-the-platform-services-sdk)
| **Capacity monitoring** | | | | |
| `GetCurrentDatabasesInformation` | Retrieve current database count | `GET /_api/v2/user/current/databases` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#getcurrentdatabasesinformation) | Use the length of the list returned from the `GetAllDbs` operation.
| `GetCurrentThroughputInformation` | Retrieve the current provisioned throughput capacity consumption | `GET /_api/v2/user/current/throughput` | [API docs link](https://cloud.ibm.com/apidocs/cloudant/cloudant-gen1?code=go#getcurrentthroughputinformation) | Use a Prometheus instant query with IBM Cloud Monitoring APIs. For example to get the capacity consumption rate over the last minute use a query like `rate(ibm_cloudant_permitted_operations_total[1m])`. See the [IBM Cloud Monitoring docs for an example](https://cloud.ibm.com/docs/monitoring?topic=monitoring-metrics_api#metrics-api-sample-prom).

#### Using the Platform Services SDK

Use the [IBM Cloud Platform Services Go SDK](https://github.com/IBM/platform-services-go-sdk) Resource Controller APIs to programmatically configure your Cloudant Gen 2 instance.

Update the configuration values in the `parameters` mapping of key-value pairs.
Read the current values from the `extensions` mapping of key-value pairs.

| Key path in `parameters` or `extensions` | Value description |
| --- | --- |
`dataservices.cloudant.capacity_units` | The provisioned throughput capacity of the instance in [units](https://cloud.ibm.com/docs/cloudant-gen2?topic=cloudant-gen2-usage-and-charges#provisioned-throughput-capacity-units)
`dataservices.cloudant.configuration.cors` | The CORS configuration object of the instance with booleans for `enabled` and `allowCredentials` and a string array of `origins`
`dataservices.cloudant.configuration.db_count_limit` | Read only value of the maximum number of databases allowed on the instance

##### Viewing configuration

Use [Get a resource instance](https://cloud.ibm.com/apidocs/resource-controller/resource-controller?code=go#get-resource-instance) to retrieve the resource instance configuration.

##### Modifying configuration

For a new instance use [Create (provision) a new resource instance](https://cloud.ibm.com/apidocs/resource-controller/resource-controller?code=go#create-resource-instance).
For a pre-existing instance use [Update a resource instance](https://cloud.ibm.com/apidocs/resource-controller/resource-controller?code=go#update-resource-instance).


### Path elements containing the `+` character

Path elements containing the `+` character in the SDKs are not interoperable with:
* Cloudant 
* Apache CouchDB versions older than 3.2.0
* Apache CouchDB versions 3.2.0 or newer with the setting `decode_plus_to_space = true`

This is because standard URL encoding libraries following the [RFC3986 URI specification](https://tools.ietf.org/html/rfc3986#section-3.3) do not encode the `+` character in path elements.
* It is possible to workaround for document names with a `+` in the ID (e.g. `docidwith+char`) by using:
    * For reading: use the `post` all docs operation and the `key` or `keys` parameter with a value of the document ID including the `+`.
    * For writing: use the `post` document operation or `post` bulk docs operation with the value of the document ID including the `+`.
* There is no pre-encoding workaround because the result is a double encoding e.g. using `%2b` in the path element ends up being double encoded as `%252b`.

### Views

#### Objects as keys

Using JSON objects as keys (e.g. `start_key`, `end_key`, `key`, `keys`)
can cause inconsistent results because the ordering of the members of the JSON
object after serialization is not guaranteed.

### Documents

#### Attachments

The `atts_since` parameter is not supported when retrieving a document.
The workaround is to call `POST /{db}/_bulk_get` using the `atts_since` field under the `docs` request body. See the [alternative example request for `atts_since` using the `/_bulk_get` endpoint](https://cloud.ibm.com/apidocs/cloudant#postbulkget) in our API Docs.
Example JSON request body:
```json
{
  "docs": [{"id": "order00058", "atts_since": "1-99b02e08da151943c2dcb40090160bb8"}]
}
```

#### Open revisions

The `open_revs` parameter is not supported when retrieving a document.
If you want to retrieve documents with all leaf revisions (`open_revs=all`), the workaround is to call `POST /{db}/_bulk_get` using the `id` field within the `docs` array request body.
See the [alternative example request for `open_revs=all` using the `/_bulk_get` endpoint](https://cloud.ibm.com/apidocs/cloudant#postbulkget) in our API Docs.
Example JSON request body:
```json
{
  "docs": [{"id": "order00067"}]
}
```

If you want to retrieve documents of specified leaf revisions (e.g. `open_revs=["3-917fa2381192822767f010b95b45325b", "4-a5be949eeb7296747cc271766e9a498b"]`), the workaround is to call `POST /{db}/_bulk_get` using the same `id` value for each unique `rev` value within of the `docs` array request body.
See the [default example request using the `/_bulk_get` endpoint](https://cloud.ibm.com/apidocs/cloudant#postbulkget) in our API Docs.
Example JSON request body:
```json
{
  "docs": [
    {
      "id": "order00067",
      "rev": "3-917fa2381192822767f010b95b45325b"
    },
    {
      "id": "order00067",
      "rev": "4-a5be949eeb7296747cc271766e9a498b"
    }
  ]
}
```

### Compression

* Manually setting an `Accept-Encoding` header on requests will disable the transparent gzip decompression of response bodies from the server.
* Manually setting a `Content-Encoding` header on requests will disable the transparent gzip compression of request bodies to the server.

### Changes feed

#### Filter functions

The SDK does not support passing user-defined query or body parameters in `_changes` requests for dynamic filter functions in design documents. 
The workaround and recommended option is to use a `selector` type filter.
For example, if you are using a `_changes` request like `/{db}/_changes?filter=myDdoc/byName&name=Jane` with a filter function like:
```javascript
function(doc, req) {
    if (doc.name !== req.query.name) {
        return false;
    }
    return true; 
}
```
It can be replaced with a request using a selector filter:
```go
postChangesOptions := service.NewPostChangesOptions("example").
  SetFilter("_selector").
  SetSelector(map[string]interface{}{"name": "Jane"})

changesResult, response, err := service.PostChanges(postChangesOptions)
if err != nil {
  panic(err)
}
```


## Cloudant SDK for Go
<!-- KNOWN_ISSUES specific to Go -->
### Search
#### Analyzer definitions should be in object format

* In order to be able to deserialize a design document with a search
  index analyzer into the model object the analyzer must be stored in the design
  document described in object format, not a string format e.g.
  ```json
  {
    "analyzer": {
      "name": "keyword"
    }
  }
  ```
  not
  ```json
  {
    "analyzer": "keyword"
  }
  ```
  Note that analyzers created using the SDK's struct models use the object format. As such the issue will only manifest
  when trying to read a design document created from another source with this error:
  ```
  json: cannot unmarshal string into Go value of type map[string]json.RawMessage
  ```

#### Facet counting deserialization errors in no match cases

In Apache CouchDB versions `3.3.2` or older and Cloudant versions `8435` and older
when obtaining a `SearchResult` with `counts` if there are no matches a deserialization error is encountered.

The workaround is to use [Raw IO](/#raw-io) functions to custom deserialize the response.

### Request bodies containing the `headers` parameter

The `Headers` map is always reserved for the API request headers.
In order to avoid collision between the request headers and the `headers` parameter
that is contained by the request body, use the `HeadersVar` as
the request body `headers` parameter.

E.g. to set the `headers` parameter of your replication
target database, use the following approach:

```go
...
target, _ := cloudant.NewReplicationDatabase(
	"<your-service-url>" + "/" + "animaldb-target",
)

target.HeadersVar = make(map[string]string)
target.HeadersVar["Authorization"] = "Basic <your-base64-encoded-auth-key>"
...
```
The example above represents this JSON body:
```json
{
    ...
    "target": {
        "headers": {
            "Authorization": "Basic <your-base64-encoded-auth-key>"
        },
        "url": "<your-service-url>/animaldb-target"
    },
    ...
}
```

### Disabling request body compression

Some issues with older server versions can be worked around by disabling
compression of request bodies. This is an example of how to do that.

```go
import (
    "github.com/ibm/cloudant-go-sdk/cloudantv1"
)
client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
    &cloudantv1.CloudantV1Options{
        ServiceName: "{your-service-name}",
    },
)
client.SetEnableGzipCompression(false)
...
```

#### Request body write errors on 4xx responses from Apache CouchDB

This problem has been observed only when using request body compression (enabled by default), HTTP/1.1 protocol, and an Apache CouchDB server. It is not known to impact IBM Cloudant.

When a request is made that is rejected by the server (for example a `409 conflict` repsonse) then in very rare cases the client may return an error writing the request body (for example `write: connection reset by peer` or `write: broken pipe`) instead of the `4xx` error returned by the server. This occurs because the client is still writing the request body after the CouchDB server has already responded with an error and terminated the connection. In nearly all cases the server error is returned correctly. However, if you experience a scenario where this issue occurs a workaround is to disable request body compression as outlined in the example above.


### Replication Documents

* In order to be able to deserialize a replication document into the model
  object the source and target databases must be stored in an object format,
  not a string format e.g.
  ```json
  {
    "source": {
      "url": "https://<SERVICE_URL>/animaldb",
      "auth" : {
        "basic": {
          "username": "<SERVICE_USERNAME>",
          "password": "<SERVICE_PASSWORD>"
        }
      }
    },
    "target": {
      "url": "https://<TARGET_SERVICE_URL>/animaldb-target",
      "auth" : {
        "basic": {
          "username": "<TARGET_SERVICE_USERNAME>",
          "password": "<TARGET_SERVICE_PASSWORD>"
        }
      }
    }
  }
  ```
  not
  ```json
  {
    "source": "https://<SERVICE_USERNAME>:<SERVICE_PASSWORD>@<SERVICE_URL>",
    "target": "https://<TARGET_SERVICE_USERNAME>:<TARGET_SERVICE_PASSWORD>@<TARGET_SERVICE_URL>"
  }
  ```
  Note that replication documents created using the SDK's struct models use the object format. As such the issue will only manifest
  when trying to read a replication document created from another source with this error:
  ```sh
  json: cannot unmarshal string into Go value of type map[string]json.RawMessage
  ```
  If you need to read a replication document that uses string format source or target URLs the workaround is to use the `GetDocumentAsStream` [Raw IO](/#raw-io) function with the `_replicator` database name and your replication document id to custom deserialize the response.

### Queries

#### Explain response

With Apache CouchDB versions `3.3.2` or older and Cloudant versions `8442` and older
deserialization of the `_explain` response fails when the request does not
list any `fields` to project and hence the response includes `"fields": "all_fields"`.

If it is not possible to use a newer server version the workaround is to use
[Raw IO](/#raw-io) functions to custom deserialize the response.

### Design documents

#### Query language design documents

The type of the `map` function for `views` differs between design documents with `"language":"javascript"` (`map` function is a string) and `"language":"query"` (`map` function is an object).

The SDK model for a design document expects a `javascript` design document and the [`GetDesignDocument` operation](https://cloud.ibm.com/apidocs/cloudant?code=go#getdesigndocument) fails for a `query` design document with an error such as `error unmarshalling cloudantv1.DesignDocument: error unmarshalling property 'views' as map[string]cloudantv1.DesignDocumentViewsMapReduce: error unmarshalling property 'map': json: cannot unmarshal object into Go value of type string`.

In general manage `query` design documents using the `_index` endpoint operations: [`PostIndex`](https://cloud.ibm.com/apidocs/cloudant?code=go#postindex), [`DeleteIndex`](https://cloud.ibm.com/apidocs/cloudant?code=go#deleteindex), and [`GetIndexesInformation`](https://cloud.ibm.com/apidocs/cloudant?code=go#getindexesinformation). However, if you must retrieve the content of a `query` design document workaround the error by using the [`PostDesignDocs` operation](https://cloud.ibm.com/apidocs/cloudant?code=go#postdesigndocs) with the `IncludeDocs` parameter and the `Key` parameter to limit the response to a single specific design document. If a completely raw response is preferred then instead use the [`PostAllDocsAsStream operation`](https://cloud.ibm.com/apidocs/cloudant?code=go#postalldocs).
