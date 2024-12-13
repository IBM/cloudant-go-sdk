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

eventDoc := cloudantv1.Document{}
eventDoc.SetProperty("type", "event")
eventDoc.SetProperty("userId", "abc123")
eventDoc.SetProperty("eventType", "addedToBasket")
eventDoc.SetProperty("productId", "1000042")
eventDoc.SetProperty("date", "2019-01-28T10:44:22.000Z")

putDocumentOptions := service.NewPutDocumentOptions(
  "events",
  "ns1HJS13AMkK:0007241142412418284",
)
putDocumentOptions.SetDocument(&eventDoc)

documentResult, response, err := service.PutDocument(putDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
