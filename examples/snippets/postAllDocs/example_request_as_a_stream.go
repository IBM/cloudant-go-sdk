// section: code imports
import (
  "encoding/json"
  "fmt"
  "io"
  "os"

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

allDocsResult, response, err := service.PostAllDocsAsStream(postAllDocsOptions)
if err != nil {
    panic(err)
}

if allDocsResult != nil {
  defer allDocsResult.Close()
  outFile, err := os.Create("result.json")
  if err != nil {
    panic(err)
  }
  defer outFile.Close()
  if _, err = io.Copy(outFile, allDocsResult); err != nil {
    panic(err)
  }
}
