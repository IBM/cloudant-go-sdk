// section: code
postAllDocsOptions := service.NewPostAllDocsOptions(
  "orders",
)
postAllDocsOptions.SetIncludeDocs(true)
postAllDocsOptions.SetStartKey("abc")
postAllDocsOptions.SetLimit(10)

allDocsResult, response, err := service.PostAllDocsAsStream(postAllDocsOptions)
if err != nil {
    panic(err)
}

if allDocsResult != nil {
  defer allDocsResult.Close()
  outFile, err := os.Create("result.json")
  if err != nil {
    panic(err)
  }
  defer outFile.Close()
  if _, err = io.Copy(outFile, allDocsResult); err != nil {
    panic(err)
  }
}
