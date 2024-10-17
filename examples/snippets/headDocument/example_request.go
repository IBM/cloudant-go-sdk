// section: code
headDocumentOptions := service.NewHeadDocumentOptions(
  "orders",
  "order00058",
)

response, err := service.HeadDocument(headDocumentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["ETag"])
