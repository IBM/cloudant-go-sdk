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

allDocsQueries := []cloudantv1.AllDocsQuery{
  {
    Keys: []string{
      "1000042",
      "1000043",
    },
  },
  {
    Limit: core.Int64Ptr(3),
    Skip:  core.Int64Ptr(2),
  },
}
postAllDocsQueriesOptions := service.NewPostAllDocsQueriesOptions(
  "products",
  allDocsQueries,
)

allDocsQueriesResult, response, err := service.PostAllDocsQueries(postAllDocsQueriesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsQueriesResult, "", "  ")
fmt.Println(string(b))
