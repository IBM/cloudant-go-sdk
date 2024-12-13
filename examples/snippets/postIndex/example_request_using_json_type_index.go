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

// Type "json" index fields require an object that maps the name of a field to a sort direction.
var indexField cloudantv1.IndexField
indexField.SetProperty("email", core.StringPtr("asc"))

postIndexOptions := service.NewPostIndexOptions(
  "users",
  &cloudantv1.IndexDefinition{
    Fields: []cloudantv1.IndexField{
      indexField,
    },
  },
)
postIndexOptions.SetDdoc("json-index")
postIndexOptions.SetName("getUserByEmail")
postIndexOptions.SetType("json")

indexResult, response, err := service.PostIndex(postIndexOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(indexResult, "", "  ")
fmt.Println(string(b))
