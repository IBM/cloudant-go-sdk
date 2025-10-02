// section: code imports
import (
  "encoding/json"
  "fmt"

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

docID := "order00067"

bulkGetDocs := []cloudantv1.BulkGetQueryDocument{
  {
    ID: &docID,
    Rev: core.StringPtr("3-22222222222222222222222222222222"),
  },
  {
    ID: &docID,
    Rev: core.StringPtr("4-33333333333333333333333333333333"),
  },
}

postBulkGetOptions := service.NewPostBulkGetOptions(
  "orders",
  bulkGetDocs,
)
bulkGetResult, response, err := service.PostBulkGet(postBulkGetOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(bulkGetResult, "", "  ")
fmt.Println(string(b))
