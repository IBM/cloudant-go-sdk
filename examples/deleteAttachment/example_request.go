// section: code
deleteAttachmentOptions := service.NewDeleteAttachmentOptions(
  "products",
  "small-appliances:100001",
  "product_details.txt",
)
deleteAttachmentOptions.SetRev("4-1a0d1cd6f40472509e9aac646183736a")

documentResult, response, err := service.DeleteAttachment(deleteAttachmentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `product_details.txt` attachment in `small-appliances:100001` document to exist. To create the attachment, see [Create or modify an attachment.](#putattachment)
