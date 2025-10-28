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

doc1 := cloudantv1.AllDocsQuery{
  Descending:  core.BoolPtr(true),
  IncludeDocs: core.BoolPtr(true),
  Limit:       core.Int64Ptr(10),
}

doc2 := cloudantv1.AllDocsQuery{
  InclusiveEnd: core.BoolPtr(true),
  StartKey:     core.StringPtr("_design/allusers"),
  Skip:         core.Int64Ptr(1),
}

postDesignDocsQueriesOptions := service.NewPostDesignDocsQueriesOptions(
  "users",
  []cloudantv1.AllDocsQuery{
    doc1,
    doc2,
  },
)

allDocsQueriesResult, response, err := service.PostDesignDocsQueries(postDesignDocsQueriesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsQueriesResult, "", "  ")
fmt.Println(string(b))
