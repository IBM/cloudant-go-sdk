// section: code
deleteDocumentOptions := service.NewDeleteDocumentOptions(
  "events",
  "0007241142412418284",
)
deleteDocumentOptions.SetRev("2-9a0d1cd9f40472509e9aac6461837367")

documentResult, response, err := service.DeleteDocument(deleteDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
