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

docID := "order00058"

postBulkGetOptions := service.NewPostBulkGetOptions(
  "orders",
  []cloudantv1.BulkGetQueryDocument{
    {
      ID: &docID,
      AttsSince: []string{"1-00000000000000000000000000000000"},
    },
  },
)

bulkGetResult, response, err := service.PostBulkGet(postBulkGetOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(bulkGetResult, "", "  ")
fmt.Println(string(b))
