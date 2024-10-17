// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
eventDoc1 := cloudantv1.Document{
  ID: core.StringPtr("ns1HJS13AMkK:0007241142412418284"),
}
eventDoc1.SetProperty("type", "event")
eventDoc1.SetProperty("userId", "abc123")
eventDoc1.SetProperty("eventType", "addedToBasket")
eventDoc1.SetProperty("productId", "1000042")
eventDoc1.SetProperty("date", "2019-01-28T10:44:22.000Z")

eventDoc2 := cloudantv1.Document{
  ID: core.StringPtr("H8tDIwfadxp9:0007241142412418285"),
}
eventDoc2.SetProperty("type", "event")
eventDoc2.SetProperty("userId", "abc234")
eventDoc2.SetProperty("eventType", "addedToBasket")
eventDoc2.SetProperty("productId", "1000050")
eventDoc2.SetProperty("date", "2019-01-25T20:00:00.000Z")

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
