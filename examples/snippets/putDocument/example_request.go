// section: code
eventDoc := cloudantv1.Document{}
eventDoc.SetProperty("type", "event")
eventDoc.SetProperty("userid", "abc123")
eventDoc.SetProperty("eventType", "addedToBasket")
eventDoc.SetProperty("productId", "1000042")
eventDoc.SetProperty("date", "2019-01-28T10:44:22.000Z")

putDocumentOptions := service.NewPutDocumentOptions(
  "events",
  "0007241142412418284",
)
putDocumentOptions.SetDocument(&eventDoc)

documentResult, response, err := service.PutDocument(putDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
