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

getDocumentOptions := service.NewGetDocumentOptions(
  "products",
  "1000042",
)

document, response, err := service.GetDocument(getDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(document, "", "  ")
fmt.Println(string(b))
