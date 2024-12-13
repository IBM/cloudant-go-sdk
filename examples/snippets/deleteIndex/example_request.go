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

deleteIndexOptions := service.NewDeleteIndexOptions(
  "users",
  "json-index",
  "json",
  "getUserByName",
)

ok, response, err := service.DeleteIndex(deleteIndexOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example will fail if `getUserByName` index doesn't exist. To create the index, see [Create a new index on a database.](#postindex)
