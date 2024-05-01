# Examples for go

## getServerInformation

_GET `/`_

### [Example request](snippets/getServerInformation/example_request.go)

[embedmd]:# (snippets/getServerInformation/example_request.go)
```go
// section: code
getServerInformationOptions := service.NewGetServerInformationOptions()

serverInformation, response, err := service.GetServerInformation(getServerInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(serverInformation, "", "  ")
fmt.Println(string(b))
```

## getActiveTasks

_GET `/_active_tasks`_

### [Example request](snippets/getActiveTasks/example_request.go)

[embedmd]:# (snippets/getActiveTasks/example_request.go)
```go
// section: code
getActiveTasksOptions := service.NewGetActiveTasksOptions()

activeTask, response, err := service.GetActiveTasks(getActiveTasksOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(activeTask, "", "  ")
fmt.Println(string(b))
```

## getAllDbs

_GET `/_all_dbs`_

### [Example request](snippets/getAllDbs/example_request.go)

[embedmd]:# (snippets/getAllDbs/example_request.go)
```go
// section: code
getAllDbsOptions := service.NewGetAllDbsOptions()

result, response, err := service.GetAllDbs(getAllDbsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(result, "", "  ")
fmt.Println(string(b))
```

## postApiKeys

_POST `/_api/v2/api_keys`_

### [Example request](snippets/postApiKeys/example_request.go)

[embedmd]:# (snippets/postApiKeys/example_request.go)
```go
// section: code
postApiKeysOptions := service.NewPostApiKeysOptions()

apiKeysResult, response, err := service.PostApiKeys(postApiKeysOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(apiKeysResult, "", "  ")
fmt.Println(string(b))
```

## putCloudantSecurity

_PUT `/_api/v2/db/{db}/_security`_

### [Example request](snippets/putCloudantSecurity/example_request.go)

[embedmd]:# (snippets/putCloudantSecurity/example_request.go)
```go
// section: code
putCloudantSecurityConfigurationOptions := service.NewPutCloudantSecurityConfigurationOptions(
  "products",
  map[string][]string{
    "nobody": {cloudantv1.SecurityCloudantReaderConst},
  },
)

ok, response, err := service.PutCloudantSecurityConfiguration(putCloudantSecurityConfigurationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
// section: markdown
// The `nobody` username applies to all unauthenticated connection attempts. For example, if an application tries to read data from a database, but didn't identify itself, the task can continue only if the `nobody` user has the role `_reader`.
// section: markdown
// If instead of using Cloudant's security model for managing permissions you opt to use the Apache CouchDB `_users` database (that is using legacy credentials _and_ the `couchdb_auth_only:true` option) then be aware that the user must already exist in `_users` database before adding permissions. For information on the `_users` database, see <a href="https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-work-with-your-account#using-the-users-database-with-cloudant-nosql-db" target="_blank">Using the `_users` database with Cloudant</a>.
```

## getActivityTrackerEvents

_GET `/_api/v2/user/activity_tracker/events`_

### [Example request](snippets/getActivityTrackerEvents/example_request.go)

[embedmd]:# (snippets/getActivityTrackerEvents/example_request.go)
```go
// section: code
getActivityTrackerEventsOptions := service.NewGetActivityTrackerEventsOptions()

activityTrackerEvents, response, err := service.GetActivityTrackerEvents(getActivityTrackerEventsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(activityTrackerEvents, "", "  ")
fmt.Println(string(b))
```

## postActivityTrackerEvents

_POST `/_api/v2/user/activity_tracker/events`_

### [Example request](snippets/postActivityTrackerEvents/example_request.go)

[embedmd]:# (snippets/postActivityTrackerEvents/example_request.go)
```go
// section: code
postActivityTrackerEventsOptions := service.NewPostActivityTrackerEventsOptions(
  []string{"management"},
)

activityTrackerEvents, response, err := service.PostActivityTrackerEvents(postActivityTrackerEventsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(activityTrackerEvents, "", "  ")
fmt.Println(string(b))
```

## getCapacityThroughputInformation

_GET `/_api/v2/user/capacity/throughput`_

### [Example request](snippets/getCapacityThroughputInformation/example_request.go)

[embedmd]:# (snippets/getCapacityThroughputInformation/example_request.go)
```go
// section: code
getCapacityThroughputInformationOptions := service.NewGetCapacityThroughputInformationOptions()

capacityThroughputInformation, response, err := service.GetCapacityThroughputInformation(getCapacityThroughputInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(capacityThroughputInformation, "", "  ")
fmt.Println(string(b))
```

## putCapacityThroughputConfiguration

_PUT `/_api/v2/user/capacity/throughput`_

### [Example request](snippets/putCapacityThroughputConfiguration/example_request.go)

[embedmd]:# (snippets/putCapacityThroughputConfiguration/example_request.go)
```go
// section: code
putCapacityThroughputConfigurationOptions := service.NewPutCapacityThroughputConfigurationOptions(
  1,
)

capacityThroughputConfiguration, response, err := service.PutCapacityThroughputConfiguration(putCapacityThroughputConfigurationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(capacityThroughputConfiguration, "", "  ")
fmt.Println(string(b))
```

## getCorsInformation

_GET `/_api/v2/user/config/cors`_

### [Example request](snippets/getCorsInformation/example_request.go)

[embedmd]:# (snippets/getCorsInformation/example_request.go)
```go
// section: code
getCorsInformationOptions := service.NewGetCorsInformationOptions()

corsConfiguration, response, err := service.GetCorsInformation(getCorsInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(corsConfiguration, "", "  ")
fmt.Println(string(b))
```

## putCorsConfiguration

_PUT `/_api/v2/user/config/cors`_

### [Example request](snippets/putCorsConfiguration/example_request.go)

[embedmd]:# (snippets/putCorsConfiguration/example_request.go)
```go
// section: code
putCorsConfigurationOptions := service.NewPutCorsConfigurationOptions([]string{
  "https://example.com",
})
putCorsConfigurationOptions.SetEnableCors(true)

ok, response, err := service.PutCorsConfiguration(putCorsConfigurationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
```

## getCurrentThroughputInformation

_GET `/_api/v2/user/current/throughput`_

### [Example request](snippets/getCurrentThroughputInformation/example_request.go)

[embedmd]:# (snippets/getCurrentThroughputInformation/example_request.go)
```go
// section: code
getCurrentThroughputInformationOptions := service.NewGetCurrentThroughputInformationOptions()

currentThroughputInformation, response, err := service.GetCurrentThroughputInformation(getCurrentThroughputInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(currentThroughputInformation, "", "  ")
fmt.Println(string(b))
```

## getDbUpdates

_GET `/_db_updates`_

### [Example request](snippets/getDbUpdates/example_request.go)

[embedmd]:# (snippets/getDbUpdates/example_request.go)
```go
// section: code
getDbUpdatesOptions := service.NewGetDbUpdatesOptions()
getDbUpdatesOptions.SetFeed("normal")
getDbUpdatesOptions.SetHeartbeat(10000)
getDbUpdatesOptions.SetSince("now")

dbUpdates, response, err := service.GetDbUpdates(getDbUpdatesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(dbUpdates, "", "  ")
fmt.Println(string(b))
// section: markdown
// This request requires `server_admin` access.
```

## postDbsInfo

_POST `/_dbs_info`_

### [Example request](snippets/postDbsInfo/example_request.go)

[embedmd]:# (snippets/postDbsInfo/example_request.go)
```go
// section: code
postDbsInfoOptions := service.NewPostDbsInfoOptions([]string{
  "products",
  "users",
  "orders",
})

dbsInfoResult, response, err := service.PostDbsInfo(postDbsInfoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(dbsInfoResult, "", "  ")
fmt.Println(string(b))
```

## getMembershipInformation

_GET `/_membership`_

### [Example request](snippets/getMembershipInformation/example_request.go)

[embedmd]:# (snippets/getMembershipInformation/example_request.go)
```go
// section: code
getMembershipInformationOptions := service.NewGetMembershipInformationOptions()

membershipInformation, response, err := service.GetMembershipInformation(getMembershipInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(membershipInformation, "", "  ")
fmt.Println(string(b))
```

## deleteReplicationDocument

_DELETE `/_replicator/{doc_id}`_

### [Example request](snippets/deleteReplicationDocument/example_request.go)

[embedmd]:# (snippets/deleteReplicationDocument/example_request.go)
```go
// section: code
deleteReplicationDocumentOptions := service.NewDeleteReplicationDocumentOptions(
  "repldoc-example",
)
deleteReplicationDocumentOptions.SetRev("3-a0ccbdc6fe95b4184f9031d086034d85")

documentResult, response, err := service.DeleteReplicationDocument(deleteReplicationDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
```

## getReplicationDocument

_GET `/_replicator/{doc_id}`_

### [Example request](snippets/getReplicationDocument/example_request.go)

[embedmd]:# (snippets/getReplicationDocument/example_request.go)
```go
// section: code
getReplicationDocumentOptions := service.NewGetReplicationDocumentOptions(
  "repldoc-example",
)

replicationDocument, response, err := service.GetReplicationDocument(getReplicationDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(replicationDocument, "", "  ")
fmt.Println(string(b))
```

## headReplicationDocument

_HEAD `/_replicator/{doc_id}`_

### [Example request](snippets/headReplicationDocument/example_request.go)

[embedmd]:# (snippets/headReplicationDocument/example_request.go)
```go
// section: code
headReplicationDocumentOptions := service.NewHeadReplicationDocumentOptions(
  "repldoc-example",
)

response, err := service.HeadReplicationDocument(headReplicationDocumentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["Etag"])
```

## putReplicationDocument

_PUT `/_replicator/{doc_id}`_

### [Example request](snippets/putReplicationDocument/example_request.go)

[embedmd]:# (snippets/putReplicationDocument/example_request.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
source, err := service.NewReplicationDatabase(
  "<your-source-service-url>/animaldb",
)
if err != nil {
  panic(err)
}

target, err := service.NewReplicationDatabase(
  "<your-target-service-url>" + "/" + "animaldb-target",
)
if err != nil {
  panic(err)
}

auth, err := service.NewReplicationDatabaseAuthIam(
  "<your-iam-api-key>",
)
if err != nil {
  panic(err)
}
target.Auth = &cloudantv1.ReplicationDatabaseAuth{Iam: auth}

replicationDoc, err := service.NewReplicationDocument(
  source,
  target,
)
if err != nil {
  panic(err)
}

replicationDoc.CreateTarget = core.BoolPtr(true)

putReplicationDocumentOptions := service.NewPutReplicationDocumentOptions(
  "repldoc-example",
  replicationDoc,
)

documentResult, response, err := service.PutReplicationDocument(putReplicationDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
```

## getSchedulerDocs

_GET `/_scheduler/docs`_

### [Example request](snippets/getSchedulerDocs/example_request.go)

[embedmd]:# (snippets/getSchedulerDocs/example_request.go)
```go
// section: code
getSchedulerDocsOptions := service.NewGetSchedulerDocsOptions()
getSchedulerDocsOptions.SetLimit(100)
getSchedulerDocsOptions.SetStates([]string{"completed"})

schedulerDocsResult, response, err := service.GetSchedulerDocs(getSchedulerDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerDocsResult, "", "  ")
fmt.Println(string(b))
```

## getSchedulerDocument

_GET `/_scheduler/docs/_replicator/{doc_id}`_

### [Example request](snippets/getSchedulerDocument/example_request.go)

[embedmd]:# (snippets/getSchedulerDocument/example_request.go)
```go
// section: code
getSchedulerDocumentOptions := service.NewGetSchedulerDocumentOptions(
  "repldoc-example",
)

schedulerDocument, response, err := service.GetSchedulerDocument(getSchedulerDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerDocument, "", "  ")
fmt.Println(string(b))
```

## getSchedulerJobs

_GET `/_scheduler/jobs`_

### [Example request](snippets/getSchedulerJobs/example_request.go)

[embedmd]:# (snippets/getSchedulerJobs/example_request.go)
```go
// section: code
getSchedulerJobsOptions := service.NewGetSchedulerJobsOptions()
getSchedulerJobsOptions.SetLimit(100)

schedulerJobsResult, response, err := service.GetSchedulerJobs(getSchedulerJobsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerJobsResult, "", "  ")
fmt.Println(string(b))
```

## getSchedulerJob

_GET `/_scheduler/jobs/{job_id}`_

### [Example request](snippets/getSchedulerJob/example_request.go)

[embedmd]:# (snippets/getSchedulerJob/example_request.go)
```go
// section: code
getSchedulerJobOptions := service.NewGetSchedulerJobOptions(
  "7b94915cd8c4a0173c77c55cd0443939+continuous",
)

schedulerJob, response, err := service.GetSchedulerJob(getSchedulerJobOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerJob, "", "  ")
fmt.Println(string(b))
```

## headSchedulerJob

_HEAD `/_scheduler/jobs/{job_id}`_

### [Example request](snippets/headSchedulerJob/example_request.go)

[embedmd]:# (snippets/headSchedulerJob/example_request.go)
```go
// section: code
headSchedulerJobOptions := service.NewHeadSchedulerJobOptions(
  "7b94915cd8c4a0173c77c55cd0443939+continuous",
)

response, err := service.HeadSchedulerJob(headSchedulerJobOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
```

## postSearchAnalyze

_POST `/_search_analyze`_

### [Example request](snippets/postSearchAnalyze/example_request.go)

[embedmd]:# (snippets/postSearchAnalyze/example_request.go)
```go
// section: code
postSearchAnalyzeOptions := service.NewPostSearchAnalyzeOptions(
  "english",
  "running is fun",
)

searchAnalyzeResult, response, err := service.PostSearchAnalyze(postSearchAnalyzeOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchAnalyzeResult, "", "  ")
fmt.Println(string(b))
```

## getSessionInformation

_GET `/_session`_

### [Example request](snippets/getSessionInformation/example_request.go)

[embedmd]:# (snippets/getSessionInformation/example_request.go)
```go
// section: code
getSessionInformationOptions := service.NewGetSessionInformationOptions()

sessionInformation, response, err := service.GetSessionInformation(getSessionInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(sessionInformation, "", "  ")
fmt.Println(string(b))
// section: markdown
// For more details on Session Authentication, see [Authentication.](#authentication)
```

## getUpInformation

_GET `/_up`_

### [Example request](snippets/getUpInformation/example_request.go)

[embedmd]:# (snippets/getUpInformation/example_request.go)
```go
// section: code
getUpInformationOptions := service.NewGetUpInformationOptions()

upInformation, response, err := service.GetUpInformation(getUpInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(upInformation, "", "  ")
fmt.Println(string(b))
```

## getUuids

_GET `/_uuids`_

### [Example request](snippets/getUuids/example_request.go)

[embedmd]:# (snippets/getUuids/example_request.go)
```go
// section: code
getUuidsOptions := service.NewGetUuidsOptions()
getUuidsOptions.SetCount(10)

uuidsResult, response, err := service.GetUuids(getUuidsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(uuidsResult, "", "  ")
fmt.Println(string(b))
```

## deleteDatabase

_DELETE `/{db}`_

### [Example request](snippets/deleteDatabase/example_request.go)

[embedmd]:# (snippets/deleteDatabase/example_request.go)
```go
// section: code
deleteDatabaseOptions := service.NewDeleteDatabaseOptions(
  "<db-name>",
)

ok, response, err := service.DeleteDatabase(deleteDatabaseOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
```

## getDatabaseInformation

_GET `/{db}`_

### [Example request](snippets/getDatabaseInformation/example_request.go)

[embedmd]:# (snippets/getDatabaseInformation/example_request.go)
```go
// section: code
getDatabaseInformationOptions := service.NewGetDatabaseInformationOptions(
  "products",
)

databaseInformation, response, err := service.GetDatabaseInformation(getDatabaseInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(databaseInformation, "", "  ")
fmt.Println(string(b))
```

## headDatabase

_HEAD `/{db}`_

### [Example request](snippets/headDatabase/example_request.go)

[embedmd]:# (snippets/headDatabase/example_request.go)
```go
// section: code
headDatabaseOptions := service.NewHeadDatabaseOptions(
  "products",
)

response, err := service.HeadDatabase(headDatabaseOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
```

## postDocument

_POST `/{db}`_

### [Example request](snippets/postDocument/example_request.go)

[embedmd]:# (snippets/postDocument/example_request.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
productsDoc := cloudantv1.Document{
  ID: core.StringPtr("small-appliances:1000042"),
}
productsDoc.SetProperty("type", "product")
productsDoc.SetProperty("productid", "1000042")
productsDoc.SetProperty("brand", "Salter")
productsDoc.SetProperty("name", "Digital Kitchen Scales")
productsDoc.SetProperty("description", "Slim Colourful Design Electronic Cooking Appliance for Home / Kitchen, Weigh up to 5kg + Aquatronic for Liquids ml + fl. oz. 15Yr Guarantee - Green")
productsDoc.SetProperty("price", 14.99)
productsDoc.SetProperty("image", "assets/img/0gmsnghhew.jpg")

postDocumentOptions := service.NewPostDocumentOptions(
  "products",
)
postDocumentOptions.SetDocument(&productsDoc)

documentResult, response, err := service.PostDocument(postDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
```

## putDatabase

_PUT `/{db}`_

### [Example request](snippets/putDatabase/example_request.go)

[embedmd]:# (snippets/putDatabase/example_request.go)
```go
// section: code
putDatabaseOptions := service.NewPutDatabaseOptions(
  "products",
)
putDatabaseOptions.SetPartitioned(true)

ok, response, err := service.PutDatabase(putDatabaseOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
```

## postAllDocs

_POST `/{db}/_all_docs`_

### [Example request](snippets/postAllDocs/example_request.go)

[embedmd]:# (snippets/postAllDocs/example_request.go)
```go
// section: code
postAllDocsOptions := service.NewPostAllDocsOptions(
  "orders",
)
postAllDocsOptions.SetIncludeDocs(true)
postAllDocsOptions.SetStartKey("abc")
postAllDocsOptions.SetLimit(10)

allDocsResult, response, err := service.PostAllDocs(postAllDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsResult, "", "  ")
fmt.Println(string(b))
```

### [Example request as a stream](snippets/postAllDocs/example_request_as_a_stream.go)

[embedmd]:# (snippets/postAllDocs/example_request_as_a_stream.go)
```go
// section: code
postAllDocsOptions := service.NewPostAllDocsOptions(
  "orders",
)
postAllDocsOptions.SetIncludeDocs(true)
postAllDocsOptions.SetStartKey("abc")
postAllDocsOptions.SetLimit(10)

allDocsResult, response, err := service.PostAllDocsAsStream(postAllDocsOptions)
if err != nil {
    panic(err)
}

if allDocsResult != nil {
  defer allDocsResult.Close()
  outFile, err := os.Create("result.json")
  if err != nil {
    panic(err)
  }
  defer outFile.Close()
  if _, err = io.Copy(outFile, allDocsResult); err != nil {
    panic(err)
  }
}
```

## postAllDocsQueries

_POST `/{db}/_all_docs/queries`_

### [Example request](snippets/postAllDocsQueries/example_request.go)

[embedmd]:# (snippets/postAllDocsQueries/example_request.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
allDocsQueries := []cloudantv1.AllDocsQuery{
  {
    Keys: []string{
      "small-appliances:1000042",
      "small-appliances:1000043",
    },
  },
  {
    Limit: core.Int64Ptr(3),
    Skip:  core.Int64Ptr(2),
  },
}
postAllDocsQueriesOptions := service.NewPostAllDocsQueriesOptions(
  "products",
  allDocsQueries,
)

allDocsQueriesResult, response, err := service.PostAllDocsQueries(postAllDocsQueriesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsQueriesResult, "", "  ")
fmt.Println(string(b))
```

## postBulkDocs

_POST `/{db}/_bulk_docs`_

### [Example request: create documents](snippets/postBulkDocs/example_request_create_documents.go)

[embedmd]:# (snippets/postBulkDocs/example_request_create_documents.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
eventDoc1 := cloudantv1.Document{
  ID: core.StringPtr("0007241142412418284"),
}
eventDoc1.SetProperty("type", "event")
eventDoc1.SetProperty("userid", "abc123")
eventDoc1.SetProperty("eventType", "addedToBasket")
eventDoc1.SetProperty("productId", "1000042")
eventDoc1.SetProperty("date", "2019-01-28T10:44:22.000Z")

eventDoc2 := cloudantv1.Document{
  ID: core.StringPtr("0007241142412418285"),
}
eventDoc2.SetProperty("type", "event")
eventDoc2.SetProperty("userid", "abc234")
eventDoc2.SetProperty("eventType", "addedToBasket")
eventDoc2.SetProperty("productId", "1000050")
eventDoc2.SetProperty("date", "2019-01-25T20:00:00.000Z")

postBulkDocsOptions := service.NewPostBulkDocsOptions(
  "events",
)
bulkDocs, err := service.NewBulkDocs(
  []cloudantv1.Document{
    eventDoc1,
    eventDoc2,
  },
)
if err != nil {
  panic(err)
}

postBulkDocsOptions.SetBulkDocs(bulkDocs)

documentResult, response, err := service.PostBulkDocs(postBulkDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
```

### [Example request: delete documents](snippets/postBulkDocs/example_request_delete_documents.go)

[embedmd]:# (snippets/postBulkDocs/example_request_delete_documents.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
eventDoc1 := cloudantv1.Document{
  ID: core.StringPtr("0007241142412418284"),
}
eventDoc1.Rev = core.StringPtr("1-5005d65514fe9e90f8eccf174af5dd64")
eventDoc1.Deleted = core.BoolPtr(true)

eventDoc2 := cloudantv1.Document{
  ID: core.StringPtr("0007241142412418285"),
}
eventDoc2.Rev = core.StringPtr("1-2d7810b054babeda4812b3924428d6d6")
eventDoc2.Deleted = core.BoolPtr(true)

postBulkDocsOptions := service.NewPostBulkDocsOptions(
  "events",
)
bulkDocs, err := service.NewBulkDocs(
  []cloudantv1.Document{
    eventDoc1,
    eventDoc2,
  },
)
if err != nil {
  panic(err)
}

postBulkDocsOptions.SetBulkDocs(bulkDocs)

documentResult, response, err := service.PostBulkDocs(postBulkDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
```

### [Example request as a stream](snippets/postBulkDocs/example_request_as_a_stream.go)

[embedmd]:# (snippets/postBulkDocs/example_request_as_a_stream.go)
```go
// section: code
file, err := os.Open("upload.json")
if err != nil {
  panic(err)
}

postBulkDocsOptions := service.NewPostBulkDocsOptions(
  "events",
)

postBulkDocsOptions.SetBody(file)

documentResult, response, err := service.PostBulkDocs(postBulkDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// Content of upload.json
// section: code
{
  "docs": [
    {
      "_id": "0007241142412418284",
      "type": "event",
      "userid": "abc123",
      "eventType": "addedToBasket",
      "productId": "1000042",
      "date": "2019-01-28T10:44:22.000Z"
    },
    {
      "_id": "0007241142412418285",
      "type": "event",
      "userid": "abc123",
      "eventType": "addedToBasket",
      "productId": "1000050",
      "date": "2019-01-25T20:00:00.000Z"
    }
  ]
}
```

## postBulkGet

_POST `/{db}/_bulk_get`_

### [Example request](snippets/postBulkGet/example_request.go)

[embedmd]:# (snippets/postBulkGet/example_request.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
docID := "order00067"

bulkGetDocs := []cloudantv1.BulkGetQueryDocument{
  {
    ID: &docID,
    Rev: core.StringPtr("3-917fa2381192822767f010b95b45325b"),
  },
  {
    ID: &docID,
    Rev: core.StringPtr("4-a5be949eeb7296747cc271766e9a498b"),
  },
}

postBulkGetOptions := service.NewPostBulkGetOptions(
  "orders",
  bulkGetDocs,
)
bulkGetResult, response, err := service.PostBulkGet(postBulkGetOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(bulkGetResult, "", "  ")
fmt.Println(string(b))
```

### [Alternative example request for `open_revs=all`](snippets/postBulkGet/alternative_example_request_for_open_revs_all.go)

[embedmd]:# (snippets/postBulkGet/alternative_example_request_for_open_revs_all.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
postBulkGetOptions := service.NewPostBulkGetOptions(
  "orders",
  []cloudantv1.BulkGetQueryDocument{{ID: core.StringPtr("order00067")}},
)

bulkGetResult, response, err := service.PostBulkGet(postBulkGetOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(bulkGetResult, "", "  ")
fmt.Println(string(b))
```

### [Alternative example request for `atts_since`](snippets/postBulkGet/alternative_example_request_for_atts_since.go)

[embedmd]:# (snippets/postBulkGet/alternative_example_request_for_atts_since.go)
```go
// section: code
docID := "order00058"

postBulkGetOptions := service.NewPostBulkGetOptions(
  "orders",
  []cloudantv1.BulkGetQueryDocument{
    {
      ID: &docID,
      AttsSince: []string{"1-99b02e08da151943c2dcb40090160bb8"},
    },
  },
)

bulkGetResult, response, err := service.PostBulkGet(postBulkGetOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(bulkGetResult, "", "  ")
fmt.Println(string(b))
```

## postChanges

_POST `/{db}/_changes`_

### [Example request](snippets/postChanges/example_request.go)

[embedmd]:# (snippets/postChanges/example_request.go)
```go
// section: code
postChangesOptions := service.NewPostChangesOptions(
  "orders",
)

changesResult, response, err := service.PostChanges(postChangesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(changesResult, "", "  ")
fmt.Println(string(b))
```

### [Example request as a stream](snippets/postChanges/example_request_as_a_stream.go)

[embedmd]:# (snippets/postChanges/example_request_as_a_stream.go)
```go
// section: code
postChangesOptions := service.NewPostChangesOptions(
  "orders",
)

changesResult, response, err := service.PostChangesAsStream(postChangesOptions)
if err != nil {
  panic(err)
}

if changesResult != nil {
  defer changesResult.Close()
  outFile, err := os.Create("result.json")
  if err != nil {
    panic(err)
  }
  defer outFile.Close()
  if _, err = io.Copy(outFile, changesResult); err != nil {
    panic(err)
  }
}
```

## deleteDesignDocument

_DELETE `/{db}/_design/{ddoc}`_

### [Example request](snippets/deleteDesignDocument/example_request.go)

[embedmd]:# (snippets/deleteDesignDocument/example_request.go)
```go
// section: code
deleteDesignDocumentOptions := service.NewDeleteDesignDocumentOptions(
  "products",
  "appliances",
)
deleteDesignDocumentOptions.SetRev("1-98e6a25b3b45df62e7d47095ac15b16a")

documentResult, response, err := service.DeleteDesignDocument(deleteDesignDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This request requires the example revisions in the DELETE body to be replaced with valid revisions.
```

## getDesignDocument

_GET `/{db}/_design/{ddoc}`_

### [Example request](snippets/getDesignDocument/example_request.go)

[embedmd]:# (snippets/getDesignDocument/example_request.go)
```go
// section: code
getDesignDocumentOptions := service.NewGetDesignDocumentOptions(
  "products",
  "appliances",
)
getDesignDocumentOptions.SetLatest(true)

designDocument, response, err := service.GetDesignDocument(getDesignDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(designDocument, "", "  ")
fmt.Println(string(b))
```

## headDesignDocument

_HEAD `/{db}/_design/{ddoc}`_

### [Example request](snippets/headDesignDocument/example_request.go)

[embedmd]:# (snippets/headDesignDocument/example_request.go)
```go
// section: code
headDesignDocumentOptions := service.NewHeadDesignDocumentOptions(
  "products",
  "appliances",
)

response, err := service.HeadDesignDocument(headDesignDocumentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["Etag"])
```

## putDesignDocument

_PUT `/{db}/_design/{ddoc}`_

### [Example request](snippets/putDesignDocument/example_request.go)

[embedmd]:# (snippets/putDesignDocument/example_request.go)
```go
// section: code
emailViewMapReduce, err := service.NewDesignDocumentViewsMapReduce(
  "function(doc) {" +
    "if(doc.email_verified  === true){ emit(doc.email, [doc.name, doc.email_verified, doc.joined])" +
  "}",
)
if err != nil {
  panic(err)
}

userIndexDefinition, err := service.NewSearchIndexDefinition(
  "function(doc) {" +
    "index(\"name\", doc.name); index(\"active\", doc.active);" +
  "}",
)
if err != nil {
  panic(err)
}

designDocument := &cloudantv1.DesignDocument{
  Views: map[string]cloudantv1.DesignDocumentViewsMapReduce{
    "getVerifiedEmails": *emailViewMapReduce,
  },
  Indexes: map[string]cloudantv1.SearchIndexDefinition{
    "activeUsers": *userIndexDefinition,
  },
}

putDesignDocumentOptions := service.NewPutDesignDocumentOptions(
  "users",
  "allusers",
  designDocument,
)

documentResult, response, err := service.PutDesignDocument(putDesignDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))

applianceProdIdViewMapReduce, err := service.NewDesignDocumentViewsMapReduce(
  "function(doc) {" +
    "emit(doc.productid, [doc.brand, doc.name, doc.description])" +
  "}",
)
if err != nil {
  panic(err)
}

priceIndexDefinition, err := service.NewSearchIndexDefinition(
  "function(doc) {" +
    "index(\"price\", doc.price);" +
  "}",
)
if err != nil {
  panic(err)
}

partitionedDesignDocument := &cloudantv1.DesignDocument{
  Views: map[string]cloudantv1.DesignDocumentViewsMapReduce{
    "byApplianceProdId": *applianceProdIdViewMapReduce,
  },
  Indexes: map[string]cloudantv1.SearchIndexDefinition{
    "findByPrice": *priceIndexDefinition,
  },
}

putPartitionedDesignDocumentOptions := service.NewPutDesignDocumentOptions(
  "products",
  "appliances",
  partitionedDesignDocument,
)

documentResult, response, err = service.PutDesignDocument(putPartitionedDesignDocumentOptions)
if err != nil {
  panic(err)
}

b, _ = json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example creates `allusers` design document in the `users` database and `appliances` design document in the partitioned `products` database.
```

## getDesignDocumentInformation

_GET `/{db}/_design/{ddoc}/_info`_

### [Example request](snippets/getDesignDocumentInformation/example_request.go)

[embedmd]:# (snippets/getDesignDocumentInformation/example_request.go)
```go
// section: code
getDesignDocumentInformationOptions := service.NewGetDesignDocumentInformationOptions(
  "products",
  "appliances",
)

designDocumentInformation, response, err := service.GetDesignDocumentInformation(getDesignDocumentInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(designDocumentInformation, "", "  ")
fmt.Println(string(b))
```

## postSearch

_POST `/{db}/_design/{ddoc}/_search/{index}`_

### [Example request](snippets/postSearch/example_request.go)

[embedmd]:# (snippets/postSearch/example_request.go)
```go
// section: code
postSearchOptions := service.NewPostSearchOptions(
  "users",
  "allusers",
  "activeUsers",
  "name:Jane* AND active:True",
)

searchResult, response, err := service.PostSearch(postSearchOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `activeUsers` Cloudant Search index to exist. To create the design document with this index, see [Create or modify a design document.](#putdesigndocument)
```

## getSearchInfo

_GET `/{db}/_design/{ddoc}/_search_info/{index}`_

### [Example request](snippets/getSearchInfo/example_request.go)

[embedmd]:# (snippets/getSearchInfo/example_request.go)
```go
// section: code
getSearchInfoOptions := service.NewGetSearchInfoOptions(
  "products",
  "appliances",
  "findByPrice",
)

searchInfoResult, response, err := service.GetSearchInfo(getSearchInfoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchInfoResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `findByPrice` Cloudant Search partitioned index to exist. To create the design document with this index, see [Create or modify a design document.](#putdesigndocument)
```

## postView

_POST `/{db}/_design/{ddoc}/_view/{view}`_

### [Example request](snippets/postView/example_request.go)

[embedmd]:# (snippets/postView/example_request.go)
```go
// section: code
postViewOptions := service.NewPostViewOptions(
  "users",
  "allusers",
  "getVerifiedEmails",
)

viewResult, response, err := service.PostView(postViewOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(viewResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `getVerifiedEmails` view to exist. To create the design document with this view, see [Create or modify a design document.](#putdesigndocument)
```

## postViewQueries

_POST `/{db}/_design/{ddoc}/_view/{view}/queries`_

### [Example request](snippets/postViewQueries/example_request.go)

[embedmd]:# (snippets/postViewQueries/example_request.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
postViewQueriesOptions := service.NewPostViewQueriesOptions(
  "users",
  "allusers",
  "getVerifiedEmails",
  []cloudantv1.ViewQuery{
    {
      IncludeDocs: core.BoolPtr(true),
      Limit:       core.Int64Ptr(5),
    },
    {
      Descending: core.BoolPtr(true),
      Skip:       core.Int64Ptr(1),
    },
  },
)

viewQueriesResult, response, err := service.PostViewQueries(postViewQueriesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(viewQueriesResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `getVerifiedEmails` view to exist. To create the design document with this view, see [Create or modify a design document.](#putdesigndocument)
```

## postDesignDocs

_POST `/{db}/_design_docs`_

### [Example request](snippets/postDesignDocs/example_request.go)

[embedmd]:# (snippets/postDesignDocs/example_request.go)
```go
// section: code
postDesignDocsOptions := service.NewPostDesignDocsOptions(
  "users",
)
postDesignDocsOptions.SetAttachments(true)

allDocsResult, response, err := service.PostDesignDocs(postDesignDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsResult, "", "  ")
fmt.Println(string(b))
```

## postDesignDocsQueries

_POST `/{db}/_design_docs/queries`_

### [Example request](snippets/postDesignDocsQueries/example_request.go)

[embedmd]:# (snippets/postDesignDocsQueries/example_request.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
doc1 := cloudantv1.AllDocsQuery{
  Descending:  core.BoolPtr(true),
  IncludeDocs: core.BoolPtr(true),
  Limit:       core.Int64Ptr(10),
}

doc2 := cloudantv1.AllDocsQuery{
  InclusiveEnd: core.BoolPtr(true),
  Key:          core.StringPtr("_design/allusers"),
  Skip:         core.Int64Ptr(1),
}

postDesignDocsQueriesOptions := service.NewPostDesignDocsQueriesOptions(
  "users",
  []cloudantv1.AllDocsQuery{
    doc1,
    doc2,
  },
)

allDocsQueriesResult, response, err := service.PostDesignDocsQueries(postDesignDocsQueriesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsQueriesResult, "", "  ")
fmt.Println(string(b))
```

## postExplain

_POST `/{db}/_explain`_

### [Example request](snippets/postExplain/example_request.go)

[embedmd]:# (snippets/postExplain/example_request.go)
```go
// section: code
postExplainOptions := service.NewPostExplainOptions(
  "users",
  map[string]interface{}{
    "type": map[string]string{
      "$eq": "user",
    },
  },
)
postExplainOptions.SetExecutionStats(true)
postExplainOptions.SetLimit(10)

explainResult, response, err := service.PostExplain(postExplainOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(explainResult, "", "  ")
fmt.Println(string(b))
```

## postFind

_POST `/{db}/_find`_

### [Example request for "json" index type](snippets/postFind/example_request_for_json_index_type.go)

[embedmd]:# (snippets/postFind/example_request_for_json_index_type.go)
```go
// section: code
postFindOptions := service.NewPostFindOptions(
  "users",
  map[string]interface{}{
    "email_verified": map[string]bool{
      "$eq": true,
    },
  },
)
postFindOptions.SetFields(
  []string{"_id", "type", "name", "email"},
)
postFindOptions.SetSort([]map[string]string{{"email": "desc"}})
postFindOptions.SetLimit(3)

findResult, response, err := service.PostFind(postFindOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(findResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `getUserByEmail` Cloudant Query "json" index to exist. To create the index, see [Create a new index on a database.](#postindex)
```

### [Example request for "text" index type](snippets/postFind/example_request_for_text_index_type.go)

[embedmd]:# (snippets/postFind/example_request_for_text_index_type.go)
```go
// section: code
postFindOptions := service.NewPostFindOptions(
  "users",
  map[string]interface{}{
    "address": map[string]string{
      "$regex": "Street",
    },
  },
)
postFindOptions.SetFields(
  []string{"_id", "type", "name", "email", "address"},
)
postFindOptions.SetLimit(3)

findResult, response, err := service.PostFind(postFindOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(findResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `getUserByAddress` Cloudant Query "text" index to exist. To create the index, see [Create a new index on a database.](#postindex)
```

## getIndexesInformation

_GET `/{db}/_index`_

### [Example request](snippets/getIndexesInformation/example_request.go)

[embedmd]:# (snippets/getIndexesInformation/example_request.go)
```go
// section: code
getIndexesInformationOptions := service.NewGetIndexesInformationOptions(
  "users",
)

indexesInformation, response, err := service.GetIndexesInformation(getIndexesInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(indexesInformation, "", "  ")
fmt.Println(string(b))
```

## postIndex

_POST `/{db}/_index`_

### [Example request using "json" type index](snippets/postIndex/example_request_using_json_type_index.go)

[embedmd]:# (snippets/postIndex/example_request_using_json_type_index.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
// Type "json" index fields require an object that maps the name of a field to a sort direction.
var indexField cloudantv1.IndexField
indexField.SetProperty("email", core.StringPtr("asc"))

postIndexOptions := service.NewPostIndexOptions(
  "users",
  &cloudantv1.IndexDefinition{
    Fields: []cloudantv1.IndexField{
      indexField,
    },
  },
)
postIndexOptions.SetDdoc("json-index")
postIndexOptions.SetName("getUserByEmail")
postIndexOptions.SetType("json")

indexResult, response, err := service.PostIndex(postIndexOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(indexResult, "", "  ")
fmt.Println(string(b))
```

### [Example request using "text" type index](snippets/postIndex/example_request_using_text_type_index.go)

[embedmd]:# (snippets/postIndex/example_request_using_text_type_index.go)
```go
// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
// Type "text" index fields require an object with a name and type properties for the field.
var indexField cloudantv1.IndexField
indexField.SetProperty("name", core.StringPtr("address"))
indexField.SetProperty("type", core.StringPtr("string"))

postIndexOptions := service.NewPostIndexOptions(
  "users",
  &cloudantv1.IndexDefinition{
    Fields: []cloudantv1.IndexField{
      indexField,
    },
  },
)
postIndexOptions.SetDdoc("text-index")
postIndexOptions.SetName("getUserByAddress")
postIndexOptions.SetType("text")

indexResult, response, err := service.PostIndex(postIndexOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(indexResult, "", "  ")
fmt.Println(string(b))
```

## deleteIndex

_DELETE `/{db}/_index/_design/{ddoc}/{type}/{index}`_

### [Example request](snippets/deleteIndex/example_request.go)

[embedmd]:# (snippets/deleteIndex/example_request.go)
```go
// section: code
deleteIndexOptions := service.NewDeleteIndexOptions(
  "users",
  "json-index",
  "json",
  "getUserByName",
)

ok, response, err := service.DeleteIndex(deleteIndexOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example will fail if `getUserByName` index doesn't exist. To create the index, see [Create a new index on a database.](#postindex)
```

## deleteLocalDocument

_DELETE `/{db}/_local/{doc_id}`_

### [Example request](snippets/deleteLocalDocument/example_request.go)

[embedmd]:# (snippets/deleteLocalDocument/example_request.go)
```go
// section: code
deleteLocalDocumentOptions := service.NewDeleteLocalDocumentOptions(
  "orders",
  "local-0007741142412418284",
)

documentResult, response, err := service.DeleteLocalDocument(deleteLocalDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
```

## getLocalDocument

_GET `/{db}/_local/{doc_id}`_

### [Example request](snippets/getLocalDocument/example_request.go)

[embedmd]:# (snippets/getLocalDocument/example_request.go)
```go
// section: code
getLocalDocumentOptions := service.NewGetLocalDocumentOptions(
  "orders",
  "local-0007741142412418284",
)

document, response, err := service.GetLocalDocument(getLocalDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(document, "", "  ")
fmt.Println(string(b))
```

## putLocalDocument

_PUT `/{db}/_local/{doc_id}`_

### [Example request](snippets/putLocalDocument/example_request.go)

[embedmd]:# (snippets/putLocalDocument/example_request.go)
```go
// section: code
localDocument := cloudantv1.Document{}
properties := map[string]interface{}{
  "type":            "order",
  "user":            "Bob Smith",
  "orderid":         "0007741142412418284",
  "userid":          "abc123",
  "total":           214.98,
  "deliveryAddress": "19 Front Street, Darlington, DL5 1TY",
  "delivered":       true,
  "courier":         "UPS",
  "courierid":       "15125425151261289",
  "date":            "2019-01-28T10:44:22.000Z",
}
for key, value := range properties {
  localDocument.SetProperty(key, value)
}

putLocalDocumentOptions := service.NewPutLocalDocumentOptions(
  "orders",
  "local-0007741142412418284",
)
putLocalDocumentOptions.SetDocument(&localDocument)

documentResult, response, err := service.PutLocalDocument(putLocalDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
```

## postMissingRevs

_POST `/{db}/_missing_revs`_

### [Example request](snippets/postMissingRevs/example_request.go)

[embedmd]:# (snippets/postMissingRevs/example_request.go)
```go
// section: code
postMissingRevsOptions := service.NewPostMissingRevsOptions(
  "orders",
  map[string][]string{
    "order00077": {"<order00077-existing-revision>", "<2-missing-revision>"},
  },
)

missingRevsResult, response, err := service.PostMissingRevs(postMissingRevsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(missingRevsResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the example revisions in the POST body to be replaced with valid revisions.
```

## getPartitionInformation

_GET `/{db}/_partition/{partition_key}`_

### [Example request](snippets/getPartitionInformation/example_request.go)

[embedmd]:# (snippets/getPartitionInformation/example_request.go)
```go
// section: code
getPartitionInformationOptions := service.NewGetPartitionInformationOptions(
  "products",
  "small-appliances",
)

partitionInformation, response, err := service.GetPartitionInformation(getPartitionInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(partitionInformation, "", "  ")
fmt.Println(string(b))
```

## postPartitionAllDocs

_POST `/{db}/_partition/{partition_key}/_all_docs`_

### [Example request](snippets/postPartitionAllDocs/example_request.go)

[embedmd]:# (snippets/postPartitionAllDocs/example_request.go)
```go
// section: code
postPartitionAllDocsOptions := service.NewPostPartitionAllDocsOptions(
  "products",
  "small-appliances",
)
postPartitionAllDocsOptions.SetIncludeDocs(true)

allDocsResult, response, err := service.PostPartitionAllDocs(postPartitionAllDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsResult, "", "  ")
fmt.Println(string(b))
```

## postPartitionSearch

_POST `/{db}/_partition/{partition_key}/_design/{ddoc}/_search/{index}`_

### [Example request](snippets/postPartitionSearch/example_request.go)

[embedmd]:# (snippets/postPartitionSearch/example_request.go)
```go
// section: code
postPartitionSearchOptions := service.NewPostPartitionSearchOptions(
  "products",
  "small-appliances",
  "appliances",
  "findByPrice",
  "price:[14 TO 20]",
)

searchResult, response, err := service.PostPartitionSearch(postPartitionSearchOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `findByPrice` Cloudant Search partitioned index to exist. To create the design document with this index, see [Create or modify a design document.](#putdesigndocument)
```

## postPartitionView

_POST `/{db}/_partition/{partition_key}/_design/{ddoc}/_view/{view}`_

### [Example request](snippets/postPartitionView/example_request.go)

[embedmd]:# (snippets/postPartitionView/example_request.go)
```go
// section: code
postPartitionViewOptions := service.NewPostPartitionViewOptions(
  "products",
  "small-appliances",
  "appliances",
  "byApplianceProdId",
)
postPartitionViewOptions.SetIncludeDocs(true)
postPartitionViewOptions.SetLimit(10)

viewResult, response, err := service.PostPartitionView(postPartitionViewOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(viewResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `byApplianceProdId` partitioned view to exist. To create the view, see [Create or modify a design document.](#putdesigndocument)
```

## postPartitionFind

_POST `/{db}/_partition/{partition_key}/_find`_

### [Example request](snippets/postPartitionFind/example_request.go)

[embedmd]:# (snippets/postPartitionFind/example_request.go)
```go
// section: code
selector := map[string]interface{}{
  "type": map[string]string{
    "$eq": "product",
  },
}

postPartitionFindOptions := service.NewPostPartitionFindOptions(
  "products",
  "small-appliances",
  selector,
)
postPartitionFindOptions.SetFields([]string{
  "productid", "name", "description",
})

findResult, response, err := service.PostPartitionFind(postPartitionFindOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(findResult, "", "  ")
fmt.Println(string(b))
```

## postRevsDiff

_POST `/{db}/_revs_diff`_

### [Example request](snippets/postRevsDiff/example_request.go)

[embedmd]:# (snippets/postRevsDiff/example_request.go)
```go
// section: code
postRevsDiffOptions := service.NewPostRevsDiffOptions(
  "orders",
  map[string][]string{
    "order00077": {
      "<1-missing-revision>",
      "<2-missing-revision>",
      "<3-possible-ancestor-revision>",
    },
  },
)

mapStringRevsDiff, response, err := service.PostRevsDiff(postRevsDiffOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(mapStringRevsDiff, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the example revisions in the POST body to be replaced with valid revisions.
```

## getSecurity

_GET `/{db}/_security`_

### [Example request](snippets/getSecurity/example_request.go)

[embedmd]:# (snippets/getSecurity/example_request.go)
```go
// section: code
getSecurityOptions := service.NewGetSecurityOptions(
  "products",
)

security, response, err := service.GetSecurity(getSecurityOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(security, "", "  ")
fmt.Println(string(b))
```

## putSecurity

_PUT `/{db}/_security`_

### [Example request](snippets/putSecurity/example_request.go)

[embedmd]:# (snippets/putSecurity/example_request.go)
```go
// section: code
putSecurityOptions := service.NewPutSecurityOptions(
  "products",
)
putSecurityOptions.SetMembers(&cloudantv1.SecurityObject{
  Names: []string{"user1", "user2"},
  Roles: []string{"developers"},
})

ok, response, err := service.PutSecurity(putSecurityOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
// section: markdown
// The `nobody` username applies to all unauthenticated connection attempts. For example, if an application tries to read data from a database, but didn't identify itself, the task can continue only if the `nobody` user has the role `_reader`.
// section: markdown
// If instead of using Cloudant's security model for managing permissions you opt to use the Apache CouchDB `_users` database (that is using legacy credentials _and_ the `couchdb_auth_only:true` option) then be aware that the user must already exist in `_users` database before adding permissions. For information on the `_users` database, see <a href="https://cloud.ibm.com/docs/Cloudant?topic=Cloudant-work-with-your-account#using-the-users-database-with-cloudant-nosql-db" target="_blank">Using the `_users` database with Cloudant</a>.
```

## getShardsInformation

_GET `/{db}/_shards`_

### [Example request](snippets/getShardsInformation/example_request.go)

[embedmd]:# (snippets/getShardsInformation/example_request.go)
```go
// section: code
getShardsInformationOptions := service.NewGetShardsInformationOptions(
  "products",
)

shardsInformation, response, err := service.GetShardsInformation(getShardsInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(shardsInformation, "", "  ")
fmt.Println(string(b))
```

## getDocumentShardsInfo

_GET `/{db}/_shards/{doc_id}`_

### [Example request](snippets/getDocumentShardsInfo/example_request.go)

[embedmd]:# (snippets/getDocumentShardsInfo/example_request.go)
```go
// section: code
getDocumentShardsInfoOptions := service.NewGetDocumentShardsInfoOptions(
  "products",
  "small-appliances:1000042",
)

documentShardInfo, response, err := service.GetDocumentShardsInfo(getDocumentShardsInfoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentShardInfo, "", "  ")
fmt.Println(string(b))
```

## deleteDocument

_DELETE `/{db}/{doc_id}`_

### [Example request](snippets/deleteDocument/example_request.go)

[embedmd]:# (snippets/deleteDocument/example_request.go)
```go
// section: code
deleteDocumentOptions := service.NewDeleteDocumentOptions(
  "events",
  "0007241142412418284",
)
deleteDocumentOptions.SetRev("2-9a0d1cd9f40472509e9aac6461837367")

documentResult, response, err := service.DeleteDocument(deleteDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
```

## getDocument

_GET `/{db}/{doc_id}`_

### [Example request](snippets/getDocument/example_request.go)

[embedmd]:# (snippets/getDocument/example_request.go)
```go
// section: code
getDocumentOptions := service.NewGetDocumentOptions(
  "products",
  "small-appliances:1000042",
)

document, response, err := service.GetDocument(getDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(document, "", "  ")
fmt.Println(string(b))
```

## headDocument

_HEAD `/{db}/{doc_id}`_

### [Example request](snippets/headDocument/example_request.go)

[embedmd]:# (snippets/headDocument/example_request.go)
```go
// section: code
headDocumentOptions := service.NewHeadDocumentOptions(
  "events",
  "0007241142412418284",
)

response, err := service.HeadDocument(headDocumentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["Etag"])
```

## putDocument

_PUT `/{db}/{doc_id}`_

### [Example request](snippets/putDocument/example_request.go)

[embedmd]:# (snippets/putDocument/example_request.go)
```go
// section: code
eventDoc := cloudantv1.Document{}
eventDoc.SetProperty("type", "event")
eventDoc.SetProperty("userid", "abc123")
eventDoc.SetProperty("eventType", "addedToBasket")
eventDoc.SetProperty("productId", "1000042")
eventDoc.SetProperty("date", "2019-01-28T10:44:22.000Z")

putDocumentOptions := service.NewPutDocumentOptions(
  "events",
  "0007241142412418284",
)
putDocumentOptions.SetDocument(&eventDoc)

documentResult, response, err := service.PutDocument(putDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
```

## deleteAttachment

_DELETE `/{db}/{doc_id}/{attachment_name}`_

### [Example request](snippets/deleteAttachment/example_request.go)

[embedmd]:# (snippets/deleteAttachment/example_request.go)
```go
// section: code
deleteAttachmentOptions := service.NewDeleteAttachmentOptions(
  "products",
  "small-appliances:100001",
  "product_details.txt",
)
deleteAttachmentOptions.SetRev("4-1a0d1cd6f40472509e9aac646183736a")

documentResult, response, err := service.DeleteAttachment(deleteAttachmentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `product_details.txt` attachment in `small-appliances:100001` document to exist. To create the attachment, see [Create or modify an attachment.](#putattachment)
```

## getAttachment

_GET `/{db}/{doc_id}/{attachment_name}`_

### [Example request](snippets/getAttachment/example_request.go)

[embedmd]:# (snippets/getAttachment/example_request.go)
```go
// section: code
getAttachmentOptions := service.NewGetAttachmentOptions(
  "products",
  "small-appliances:100001",
  "product_details.txt",
)

result, response, err := service.GetAttachment(getAttachmentOptions)
if err != nil {
  panic(err)
}

data, _ := ioutil.ReadAll(result)
fmt.Println("\n", string(data))
// section: markdown
// This example requires the `product_details.txt` attachment in `small-appliances:100001` document to exist. To create the attachment, see [Create or modify an attachment.](#putattachment)
```

## headAttachment

_HEAD `/{db}/{doc_id}/{attachment_name}`_

### [Example request](snippets/headAttachment/example_request.go)

[embedmd]:# (snippets/headAttachment/example_request.go)
```go
// section: code
headAttachmentOptions := service.NewHeadAttachmentOptions(
  "products",
  "small-appliances:100001",
  "product_details.txt",
)

response, err := service.HeadAttachment(headAttachmentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["Content-Length"])
fmt.Println(response.Headers["Content-Type"])
// section: markdown
// This example requires the `product_details.txt` attachment in `small-appliances:100001` document to exist. To create the attachment, see [Create or modify an attachment.](#putattachment)
```

## putAttachment

_PUT `/{db}/{doc_id}/{attachment_name}`_

### [Example request](snippets/putAttachment/example_request.go)

[embedmd]:# (snippets/putAttachment/example_request.go)
```go
// section: code
putAttachmentOptions := service.NewPutAttachmentOptions(
  "products",
  "small-appliances:100001",
  "product_details.txt",
  ioutil.NopCloser(
    bytes.NewReader([]byte("This appliance includes...")),
  ),
  "text/plain",
)

documentResult, response, err := service.PutAttachment(putAttachmentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
```
