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

headDocumentOptions := service.NewHeadDocumentOptions(
  "orders",
  "order00058",
)

response, err := service.HeadDocument(headDocumentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["ETag"])
