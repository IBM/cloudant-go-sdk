// section: code
postPartitionAllDocsOptions := service.NewPostPartitionAllDocsOptions(
  "events",
  "ns1HJS13AMkK",
)
postPartitionAllDocsOptions.SetIncludeDocs(true)

allDocsResult, response, err := service.PostPartitionAllDocs(postPartitionAllDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsResult, "", "  ")
fmt.Println(string(b))
