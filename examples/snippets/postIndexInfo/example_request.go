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

indexInfoResult, response, err := service.PostIndexInfo(postIndexInfoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(indexInfoResult, "", "  ")
fmt.Println(string(b))
