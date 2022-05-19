// section: code
deleteDesignDocumentOptions := service.NewDeleteDesignDocumentOptions(
  "products",
  "appliances",
)
deleteDesignDocumentOptions.SetRev("1-98e6a25b3b45df62e7d47095ac15b16a")

documentResult, response, err := service.DeleteDesignDocument(deleteDesignDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This request requires the example revisions in the DELETE body to be replaced with valid revisions.
