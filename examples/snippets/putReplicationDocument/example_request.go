// section: code imports
import (
  "encoding/json"
  "fmt"

  "github.com/IBM/cloudant-go-sdk/cloudantv1"
  "github.com/IBM/go-sdk-core/v5/core"
)
// section: code
service, err := cloudantv1.NewCloudantV1(
  &cloudantv1.CloudantV1Options{},
)
if err != nil {
  panic(err)
}

source, err := service.NewReplicationDatabase(
  "https://~replace-with-source-host~.cloudantnosqldb.appdomain.cloud/animaldb",
)
if err != nil {
  panic(err)
}

target, err := service.NewReplicationDatabase(
  "https://~replace-with-target-host~.cloudantnosqldb.appdomain.cloud/animaldb-target",
)
if err != nil {
  panic(err)
}

auth, err := service.NewReplicationDatabaseAuthIam(
  "a1b2c3d4e5f6f1g4h7j3k6l9m2p5q8s1t4v7x0z3", //use your own IAM API key
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
