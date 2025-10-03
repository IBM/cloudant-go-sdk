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

postRevsDiffOptions := service.NewPostRevsDiffOptions(
  "orders",
  map[string][]string{
    "order00077": {
      "1-00000000000000000000000000000000", // missing revision
      "2-11111111111111111111111111111111", // missing revision
      "3-22222222222222222222222222222222", // possible ancestor revision
    },
  },
)

mapStringRevsDiff, response, err := service.PostRevsDiff(postRevsDiffOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(mapStringRevsDiff, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the example revisions in the POST body to be replaced with valid revisions.
