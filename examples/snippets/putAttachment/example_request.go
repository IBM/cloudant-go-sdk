// section: code imports
import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"

  "github.com/IBM/cloudant-go-sdk/cloudantv1"
)
// section: code
service, err := cloudantv1.NewCloudantV1(
  &cloudantv1.CloudantV1Options{},
)
if err != nil {
  panic(err)
}

putAttachmentOptions := service.NewPutAttachmentOptions(
  "products",
  "1000042",
  "product_details.txt",
  io.NopCloser(
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
