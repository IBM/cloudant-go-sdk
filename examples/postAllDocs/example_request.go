// section: code
postAllDocsOptions := service.NewPostAllDocsOptions(
  "orders",
)
postAllDocsOptions.SetIncludeDocs(true)
postAllDocsOptions.SetStartKey("abc")
postAllDocsOptions.SetLimit(10)

allDocsResult, response, err := service.PostAllDocs(postAllDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsResult, "", "  ")
fmt.Println(string(b))
