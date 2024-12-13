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

headReplicationDocumentOptions := service.NewHeadReplicationDocumentOptions(
  "repldoc-example",
)

response, err := service.HeadReplicationDocument(headReplicationDocumentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["ETag"])
