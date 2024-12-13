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

headAttachmentOptions := service.NewHeadAttachmentOptions(
  "products",
  "1000042",
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
// This example requires the `product_details.txt` attachment in `1000042` document to exist. To create the attachment, see [Create or modify an attachment.](#putattachment)
