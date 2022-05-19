// section: code
headDesignDocumentOptions := service.NewHeadDesignDocumentOptions(
  "products",
  "appliances",
)

response, err := service.HeadDesignDocument(headDesignDocumentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["Etag"])
