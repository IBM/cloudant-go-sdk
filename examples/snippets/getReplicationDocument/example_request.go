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

getReplicationDocumentOptions := service.NewGetReplicationDocumentOptions(
  "repldoc-example",
)

replicationDocument, response, err := service.GetReplicationDocument(getReplicationDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(replicationDocument, "", "  ")
fmt.Println(string(b))
