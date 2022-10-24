// section: code
getReplicationDocumentOptions := service.NewGetReplicationDocumentOptions(
  "repldoc-example",
)

replicationDocument, response, err := service.GetReplicationDocument(getReplicationDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(replicationDocument, "", "  ")
fmt.Println(string(b))
