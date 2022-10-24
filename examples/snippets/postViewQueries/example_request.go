// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
postViewQueriesOptions := service.NewPostViewQueriesOptions(
  "users",
  "allusers",
  "getVerifiedEmails",
  []cloudantv1.ViewQuery{
    {
      IncludeDocs: core.BoolPtr(true),
      Limit:       core.Int64Ptr(5),
    },
    {
      Descending: core.BoolPtr(true),
      Skip:       core.Int64Ptr(1),
    },
  },
)

viewQueriesResult, response, err := service.PostViewQueries(postViewQueriesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(viewQueriesResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `getVerifiedEmails` view to exist. To create the design document with this view, see [Create or modify a design document.](#putdesigndocument)
