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

deleteReplicationDocumentOptions := service.NewDeleteReplicationDocumentOptions(
  "repldoc-example",
)
deleteReplicationDocumentOptions.SetRev("1-00000000000000000000000000000000")

documentResult, response, err := service.DeleteReplicationDocument(deleteReplicationDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
