// section: code
headAttachmentOptions := service.NewHeadAttachmentOptions(
  "products",
  "small-appliances:100001",
  "product_details.txt",
)

response, err := service.HeadAttachment(headAttachmentOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
fmt.Println(response.Headers["Content-Length"])
fmt.Println(response.Headers["Content-Type"])
// section: markdown
// This example requires the `product_details.txt` attachment in `small-appliances:100001` document to exist. To create the attachment, see [Create or modify an attachment.](#putattachment)
