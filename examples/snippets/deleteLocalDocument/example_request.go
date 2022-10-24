// section: code
deleteLocalDocumentOptions := service.NewDeleteLocalDocumentOptions(
  "orders",
  "local-0007741142412418284",
)

documentResult, response, err := service.DeleteLocalDocument(deleteLocalDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
