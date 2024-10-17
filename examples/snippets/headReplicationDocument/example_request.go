// section: code
headReplicationDocumentOptions := service.NewHeadReplicationDocumentOptions(
  "repldoc-example",
)

response, err := service.HeadReplicationDocument(headReplicationDocumentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["ETag"])
