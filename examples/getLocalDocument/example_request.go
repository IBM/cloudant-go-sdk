// section: code
getLocalDocumentOptions := service.NewGetLocalDocumentOptions(
  "orders",
  "local-0007741142412418284",
)

document, response, err := service.GetLocalDocument(getLocalDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(document, "", "  ")
fmt.Println(string(b))
