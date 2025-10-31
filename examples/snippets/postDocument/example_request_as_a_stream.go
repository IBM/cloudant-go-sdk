// section: code imports
import (
  "encoding/json"
  "fmt"
  "os"

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

productsDocReader, err := os.Open("products_doc.json")
if err != nil {
  panic(err)
}
defer productsDocReader.Close()

postDocumentOptions := service.NewPostDocumentOptions(
  "products",
)
postDocumentOptions.SetContentType("application/json")
postDocumentOptions.SetBody(productsDocReader)

documentResult, response, err := service.PostDocument(postDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
