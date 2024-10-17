// section: code
getAttachmentOptions := service.NewGetAttachmentOptions(
  "products",
  "1000042",
  "product_details.txt",
)

result, response, err := service.GetAttachment(getAttachmentOptions)
if err != nil {
  panic(err)
}

data, _ := ioutil.ReadAll(result)
fmt.Println("\n", string(data))
// section: markdown
// This example requires the `product_details.txt` attachment in `1000042` document to exist. To create the attachment, see [Create or modify an attachment.](#putattachment)
