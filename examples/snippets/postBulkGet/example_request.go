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
    Rev: core.StringPtr("3-917fa2381192822767f010b95b45325b"),
  },
  {
    ID: &docID,
    Rev: core.StringPtr("4-a5be949eeb7296747cc271766e9a498b"),
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
