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
