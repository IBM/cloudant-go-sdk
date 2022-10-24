// section: code
file, err := os.Open("upload.json")
if err != nil {
  panic(err)
}

postBulkDocsOptions := service.NewPostBulkDocsOptions(
  "events",
)

postBulkDocsOptions.SetBody(file)

documentResult, response, err := service.PostBulkDocs(postBulkDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// Content of upload.json
// section: code
{
  "docs": [
    {
      "_id": "0007241142412418284",
      "type": "event",
      "userid": "abc123",
      "eventType": "addedToBasket",
      "productId": "1000042",
      "date": "2019-01-28T10:44:22.000Z"
    },
    {
      "_id": "0007241142412418285",
      "type": "event",
      "userid": "abc123",
      "eventType": "addedToBasket",
      "productId": "1000050",
      "date": "2019-01-25T20:00:00.000Z"
    }
  ]
}
