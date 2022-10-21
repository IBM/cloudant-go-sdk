# Examples for go

## getServerInformation

### get `/`

- [Example request](./getServerInformation/example_request.go)

## getActiveTasks

### get `/_active_tasks`

- [Example request](./getActiveTasks/example_request.go)

## getAllDbs

### get `/_all_dbs`

- [Example request](./getAllDbs/example_request.go)

## postApiKeys

### post `/_api/v2/api_keys`

- [Example request](./postApiKeys/example_request.go)

## putCloudantSecurity

### put `/_api/v2/db/{db}/_security`

- [Example request](./putCloudantSecurity/example_request.go)

## getActivityTrackerEvents

### get `/_api/v2/user/activity_tracker/events`

- [Example request](./getActivityTrackerEvents/example_request.go)

## postActivityTrackerEvents

### post `/_api/v2/user/activity_tracker/events`

- [Example request](./postActivityTrackerEvents/example_request.go)

## getCapacityThroughputInformation

### get `/_api/v2/user/capacity/throughput`

- [Example request](./getCapacityThroughputInformation/example_request.go)

## putCapacityThroughputConfiguration

### put `/_api/v2/user/capacity/throughput`

- [Example request](./putCapacityThroughputConfiguration/example_request.go)

## getCorsInformation

### get `/_api/v2/user/config/cors`

- [Example request](./getCorsInformation/example_request.go)

## putCorsConfiguration

### put `/_api/v2/user/config/cors`

- [Example request](./putCorsConfiguration/example_request.go)

## getCurrentThroughputInformation

### get `/_api/v2/user/current/throughput`

- [Example request](./getCurrentThroughputInformation/example_request.go)

## getDbUpdates

### get `/_db_updates`

- [Example request](./getDbUpdates/example_request.go)

## postDbsInfo

### post `/_dbs_info`

- [Example request](./postDbsInfo/example_request.go)

## getMembershipInformation

### get `/_membership`

- [Example request](./getMembershipInformation/example_request.go)

## deleteReplicationDocument

### delete `/_replicator/{doc_id}`

- [Example request](./deleteReplicationDocument/example_request.go)

## getReplicationDocument

### get `/_replicator/{doc_id}`

- [Example request](./getReplicationDocument/example_request.go)

## headReplicationDocument

### head `/_replicator/{doc_id}`

- [Example request](./headReplicationDocument/example_request.go)

## putReplicationDocument

### put `/_replicator/{doc_id}`

- [Example request](./putReplicationDocument/example_request.go)

## getSchedulerDocs

### get `/_scheduler/docs`

- [Example request](./getSchedulerDocs/example_request.go)

## getSchedulerDocument

### get `/_scheduler/docs/_replicator/{doc_id}`

- [Example request](./getSchedulerDocument/example_request.go)

## getSchedulerJobs

### get `/_scheduler/jobs`

- [Example request](./getSchedulerJobs/example_request.go)

## getSchedulerJob

### get `/_scheduler/jobs/{job_id}`

- [Example request](./getSchedulerJob/example_request.go)

## headSchedulerJob

### head `/_scheduler/jobs/{job_id}`

- [Example request](./headSchedulerJob/example_request.go)

## postSearchAnalyze

### post `/_search_analyze`

- [Example request](./postSearchAnalyze/example_request.go)

## getSessionInformation

### get `/_session`

- [Example request](./getSessionInformation/example_request.go)

## getUpInformation

### get `/_up`

- [Example request](./getUpInformation/example_request.go)

## getUuids

### get `/_uuids`

- [Example request](./getUuids/example_request.go)

## deleteDatabase

### delete `/{db}`

- [Example request](./deleteDatabase/example_request.go)

## getDatabaseInformation

### get `/{db}`

- [Example request](./getDatabaseInformation/example_request.go)

## headDatabase

### head `/{db}`

- [Example request](./headDatabase/example_request.go)

## postDocument

### post `/{db}`

- [Example request](./postDocument/example_request.go)

## putDatabase

### put `/{db}`

- [Example request](./putDatabase/example_request.go)

## postAllDocs

### post `/{db}/_all_docs`

- [Example request](./postAllDocs/example_request.go)
- [Example request as a stream](./postAllDocs/example_request_as_a_stream.go)

## postAllDocsQueries

### post `/{db}/_all_docs/queries`

- [Example request](./postAllDocsQueries/example_request.go)

## postBulkDocs

### post `/{db}/_bulk_docs`

- [Example request: create documents](./postBulkDocs/example_request_create_documents.go)
- [Example request: delete documents](./postBulkDocs/example_request_delete_documents.go)
- [Example request as a stream](./postBulkDocs/example_request_as_a_stream.go)

## postBulkGet

### post `/{db}/_bulk_get`

- [Example request](./postBulkGet/example_request.go)
- [Alternative example request for `open_revs=all`](./postBulkGet/alternative_example_request_for_open_revs_all.go)
- [Alternative example request for `atts_since`](./postBulkGet/alternative_example_request_for_atts_since.go)

## postChanges

### post `/{db}/_changes`

- [Example request](./postChanges/example_request.go)
- [Example request as a stream](./postChanges/example_request_as_a_stream.go)

## deleteDesignDocument

### delete `/{db}/_design/{ddoc}`

- [Example request](./deleteDesignDocument/example_request.go)

## getDesignDocument

### get `/{db}/_design/{ddoc}`

- [Example request](./getDesignDocument/example_request.go)

## headDesignDocument

### head `/{db}/_design/{ddoc}`

- [Example request](./headDesignDocument/example_request.go)

## putDesignDocument

### put `/{db}/_design/{ddoc}`

- [Example request](./putDesignDocument/example_request.go)

## getDesignDocumentInformation

### get `/{db}/_design/{ddoc}/_info`

- [Example request](./getDesignDocumentInformation/example_request.go)

## postSearch

### post `/{db}/_design/{ddoc}/_search/{index}`

- [Example request](./postSearch/example_request.go)

## getSearchInfo

### get `/{db}/_design/{ddoc}/_search_info/{index}`

- [Example request](./getSearchInfo/example_request.go)

## postView

### post `/{db}/_design/{ddoc}/_view/{view}`

- [Example request](./postView/example_request.go)

## postViewQueries

### post `/{db}/_design/{ddoc}/_view/{view}/queries`

- [Example request](./postViewQueries/example_request.go)

## postDesignDocs

### post `/{db}/_design_docs`

- [Example request](./postDesignDocs/example_request.go)

## postDesignDocsQueries

### post `/{db}/_design_docs/queries`

- [Example request](./postDesignDocsQueries/example_request.go)

## postExplain

### post `/{db}/_explain`

- [Example request](./postExplain/example_request.go)

## postFind

### post `/{db}/_find`

- [Example request for "json" index type](./postFind/example_request_for_json_index_type.go)
- [Example request for "text" index type](./postFind/example_request_for_text_index_type.go)

## getIndexesInformation

### get `/{db}/_index`

- [Example request](./getIndexesInformation/example_request.go)

## postIndex

### post `/{db}/_index`

- [Example request using "json" type index](./postIndex/example_request_using_json_type_index.go)
- [Example request using "text" type index](./postIndex/example_request_using_text_type_index.go)

## deleteIndex

### delete `/{db}/_index/_design/{ddoc}/{type}/{index}`

- [Example request](./deleteIndex/example_request.go)

## deleteLocalDocument

### delete `/{db}/_local/{doc_id}`

- [Example request](./deleteLocalDocument/example_request.go)

## getLocalDocument

### get `/{db}/_local/{doc_id}`

- [Example request](./getLocalDocument/example_request.go)

## putLocalDocument

### put `/{db}/_local/{doc_id}`

- [Example request](./putLocalDocument/example_request.go)

## postMissingRevs

### post `/{db}/_missing_revs`

- [Example request](./postMissingRevs/example_request.go)

## getPartitionInformation

### get `/{db}/_partition/{partition_key}`

- [Example request](./getPartitionInformation/example_request.go)

## postPartitionAllDocs

### post `/{db}/_partition/{partition_key}/_all_docs`

- [Example request](./postPartitionAllDocs/example_request.go)

## postPartitionSearch

### post `/{db}/_partition/{partition_key}/_design/{ddoc}/_search/{index}`

- [Example request](./postPartitionSearch/example_request.go)

## postPartitionView

### post `/{db}/_partition/{partition_key}/_design/{ddoc}/_view/{view}`

- [Example request](./postPartitionView/example_request.go)

## postPartitionFind

### post `/{db}/_partition/{partition_key}/_find`

- [Example request](./postPartitionFind/example_request.go)

## postRevsDiff

### post `/{db}/_revs_diff`

- [Example request](./postRevsDiff/example_request.go)

## getSecurity

### get `/{db}/_security`

- [Example request](./getSecurity/example_request.go)

## putSecurity

### put `/{db}/_security`

- [Example request](./putSecurity/example_request.go)

## getShardsInformation

### get `/{db}/_shards`

- [Example request](./getShardsInformation/example_request.go)

## getDocumentShardsInfo

### get `/{db}/_shards/{doc_id}`

- [Example request](./getDocumentShardsInfo/example_request.go)

## deleteDocument

### delete `/{db}/{doc_id}`

- [Example request](./deleteDocument/example_request.go)

## getDocument

### get `/{db}/{doc_id}`

- [Example request](./getDocument/example_request.go)

## headDocument

### head `/{db}/{doc_id}`

- [Example request](./headDocument/example_request.go)

## putDocument

### put `/{db}/{doc_id}`

- [Example request](./putDocument/example_request.go)

## deleteAttachment

### delete `/{db}/{doc_id}/{attachment_name}`

- [Example request](./deleteAttachment/example_request.go)

## getAttachment

### get `/{db}/{doc_id}/{attachment_name}`

- [Example request](./getAttachment/example_request.go)

## headAttachment

### head `/{db}/{doc_id}/{attachment_name}`

- [Example request](./headAttachment/example_request.go)

## putAttachment

### put `/{db}/{doc_id}/{attachment_name}`

- [Example request](./putAttachment/example_request.go)
