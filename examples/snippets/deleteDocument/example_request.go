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

deleteDocumentOptions := service.NewDeleteDocumentOptions(
  "orders",
  "order00058",
)
deleteDocumentOptions.SetRev("1-99b02e08da151943c2dcb40090160bb8")

documentResult, response, err := service.DeleteDocument(deleteDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
