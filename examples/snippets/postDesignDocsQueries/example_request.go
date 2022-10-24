// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
doc1 := cloudantv1.AllDocsQuery{
  Descending:  core.BoolPtr(true),
  IncludeDocs: core.BoolPtr(true),
  Limit:       core.Int64Ptr(10),
}

doc2 := cloudantv1.AllDocsQuery{
  InclusiveEnd: core.BoolPtr(true),
  Key:          core.StringPtr("_design/allusers"),
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
