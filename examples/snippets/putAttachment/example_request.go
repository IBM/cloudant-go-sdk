// section: code
putAttachmentOptions := service.NewPutAttachmentOptions(
  "products",
  "1000042",
  "product_details.txt",
  ioutil.NopCloser(
    bytes.NewReader([]byte("This appliance includes...")),
  ),
  "text/plain",
)

documentResult, response, err := service.PutAttachment(putAttachmentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
