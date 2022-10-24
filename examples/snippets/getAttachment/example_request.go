// section: code
getAttachmentOptions := service.NewGetAttachmentOptions(
  "products",
  "small-appliances:100001",
  "product_details.txt",
)

result, response, err := service.GetAttachment(getAttachmentOptions)
if err != nil {
  panic(err)
}

data, _ := ioutil.ReadAll(result)
fmt.Println("\n", string(data))
// section: markdown
// This example requires the `product_details.txt` attachment in `small-appliances:100001` document to exist. To create the attachment, see [Create or modify an attachment.](#putattachment)
