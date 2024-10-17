// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
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
