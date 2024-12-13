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

postFindOptions := service.NewPostFindOptions(
  "users",
  map[string]interface{}{
    "address": map[string]string{
      "$regex": "Street",
    },
  },
)
postFindOptions.SetFields(
  []string{"_id", "type", "name", "email", "address"},
)
postFindOptions.SetLimit(3)

findResult, response, err := service.PostFind(postFindOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(findResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `getUserByAddress` Cloudant Query "text" index to exist. To create the index, see [Create a new index on a database.](#postindex)
