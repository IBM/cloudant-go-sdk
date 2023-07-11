# Limitations, Restrictions, and Known Issues

## All Cloudant SDKs

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

## Cloudant SDK for Go
<!-- KNOWN_ISSUES specific to Go -->
### Search
#### Analyzer definitions should be in object format

* In order to be able to deserialize a design document with a search
  index analyzer into the model object the analyzer must be stored in the design
  document described in object format, not string format e.g.
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

When obtaining a `SearchResult` with `counts` if there are no matches a deserialization error is encountered.

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
