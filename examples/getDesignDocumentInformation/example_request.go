// section: code
getDesignDocumentInformationOptions := service.NewGetDesignDocumentInformationOptions(
  "products",
  "appliances",
)

designDocumentInformation, response, err := service.GetDesignDocumentInformation(getDesignDocumentInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(designDocumentInformation, "", "  ")
fmt.Println(string(b))
