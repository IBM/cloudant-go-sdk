// section: code imports
import (
  "encoding/json"
  "fmt"

  "github.com/IBM/cloudant-go-sdk/cloudantv1"
)
// section: code
service, err := cloudantv1.NewCloudantV1(
  &cloudantv1.CloudantV1Options{},
)
if err != nil {
  panic(err)
}

deleteAttachmentOptions := service.NewDeleteAttachmentOptions(
  "products",
  "1000042",
  "product_details.txt",
)
deleteAttachmentOptions.SetRev("4-33333333333333333333333333333333")

documentResult, response, err := service.DeleteAttachment(deleteAttachmentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `product_details.txt` attachment in `1000042` document to exist. To create the attachment, see [Create or modify an attachment.](#putattachment)
