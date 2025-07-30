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

// Type "text" index fields require an object with a name and type properties for the field.
indexField := cloudantv1.IndexField{
  Name: core.StringPtr("address"),
  Type: core.StringPtr(cloudantv1.IndexFieldTypeStringConst),
}

postIndexOptions := service.NewPostIndexOptions(
  "users",
  &cloudantv1.IndexDefinition{
    Fields: []cloudantv1.IndexField{
      indexField,
    },
  },
)
postIndexOptions.SetDdoc("text-index")
postIndexOptions.SetName("getUserByAddress")
postIndexOptions.SetType("text")

indexResult, response, err := service.PostIndex(postIndexOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(indexResult, "", "  ")
fmt.Println(string(b))
