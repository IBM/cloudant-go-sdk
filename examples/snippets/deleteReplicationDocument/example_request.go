// section: code
deleteReplicationDocumentOptions := service.NewDeleteReplicationDocumentOptions(
  "repldoc-example",
)
deleteReplicationDocumentOptions.SetRev("3-a0ccbdc6fe95b4184f9031d086034d85")

documentResult, response, err := service.DeleteReplicationDocument(deleteReplicationDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
