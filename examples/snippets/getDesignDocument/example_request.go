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

getDesignDocumentOptions := service.NewGetDesignDocumentOptions(
  "products",
  "appliances",
)

designDocument, response, err := service.GetDesignDocument(getDesignDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(designDocument, "", "  ")
fmt.Println(string(b))
