// section: code
headDocumentOptions := service.NewHeadDocumentOptions(
  "events",
  "0007241142412418284",
)

response, err := service.HeadDocument(headDocumentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["Etag"])
