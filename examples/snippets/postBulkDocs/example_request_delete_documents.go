// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
eventDoc1 := cloudantv1.Document{
  ID: core.StringPtr("ns1HJS13AMkK:0007241142412418284"),
}
eventDoc1.Rev = core.StringPtr("1-5005d65514fe9e90f8eccf174af5dd64")
eventDoc1.Deleted = core.BoolPtr(true)

eventDoc2 := cloudantv1.Document{
  ID: core.StringPtr("H8tDIwfadxp9:0007241142412418285"),
}
eventDoc2.Rev = core.StringPtr("1-2d7810b054babeda4812b3924428d6d6")
eventDoc2.Deleted = core.BoolPtr(true)

postBulkDocsOptions := service.NewPostBulkDocsOptions(
  "events",
)
bulkDocs, err := service.NewBulkDocs(
  []cloudantv1.Document{
    eventDoc1,
    eventDoc2,
  },
)
if err != nil {
  panic(err)
}

postBulkDocsOptions.SetBulkDocs(bulkDocs)

documentResult, response, err := service.PostBulkDocs(postBulkDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
