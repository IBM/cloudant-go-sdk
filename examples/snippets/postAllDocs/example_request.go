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

postAllDocsOptions := service.NewPostAllDocsOptions(
  "orders",
)
postAllDocsOptions.SetIncludeDocs(true)
postAllDocsOptions.SetStartKey("abc")
postAllDocsOptions.SetLimit(10)

allDocsResult, response, err := service.PostAllDocs(postAllDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsResult, "", "  ")
fmt.Println(string(b))
