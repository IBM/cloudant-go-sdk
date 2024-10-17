// section: code
headDesignDocumentOptions := service.NewHeadDesignDocumentOptions(
  "events",
  "checkout",
)

response, err := service.HeadDesignDocument(headDesignDocumentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["ETag"])
