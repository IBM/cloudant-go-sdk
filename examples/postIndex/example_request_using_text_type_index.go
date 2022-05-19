// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
// Type "text" index fields require an object with a name and type properties for the field.
var indexField cloudantv1.IndexField
indexField.SetProperty("name", core.StringPtr("address"))
indexField.SetProperty("type", core.StringPtr("string"))

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
