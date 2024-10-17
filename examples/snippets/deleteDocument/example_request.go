// section: code
deleteDocumentOptions := service.NewDeleteDocumentOptions(
  "orders",
  "order00058",
)
deleteDocumentOptions.SetRev("1-99b02e08da151943c2dcb40090160bb8")

documentResult, response, err := service.DeleteDocument(deleteDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
