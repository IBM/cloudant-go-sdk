// section: code
getDesignDocumentOptions := service.NewGetDesignDocumentOptions(
  "products",
  "appliances",
)
getDesignDocumentOptions.SetLatest(true)

designDocument, response, err := service.GetDesignDocument(getDesignDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(designDocument, "", "  ")
fmt.Println(string(b))
