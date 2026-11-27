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

postIndexInfoOptions := service.NewPostIndexInfoOptions(
  "recipes",
)
postIndexInfoOptions.SetKeys([]string{"_design/cookbook"})
postIndexInfoOptions.SetType([]string{"search"})

indexInfoResult, response, err := service.PostIndexInfo(postIndexInfoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(indexInfoResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `cookbook` design document to exist with a Cloudant Search index. To create the design document, see [Create or modify a design document.](#putdesigndocument)
