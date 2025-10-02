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

deleteDesignDocumentOptions := service.NewDeleteDesignDocumentOptions(
  "products",
  "appliances",
)
deleteDesignDocumentOptions.SetRev("1-00000000000000000000000000000000")

documentResult, response, err := service.DeleteDesignDocument(deleteDesignDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This request requires the example revisions in the DELETE body to be replaced with valid revisions.
