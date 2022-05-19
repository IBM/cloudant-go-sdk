// section: code
getDocumentOptions := service.NewGetDocumentOptions(
  "products",
  "small-appliances:1000042",
)

document, response, err := service.GetDocument(getDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(document, "", "  ")
fmt.Println(string(b))
