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

localDocument := cloudantv1.Document{}
properties := map[string]interface{}{
  "type":            "order",
  "user":            "Bob Smith",
  "orderId":         "0007741142412418284",
  "userId":          "abc123",
  "total":           214.98,
  "deliveryAddress": "19 Front Street, Darlington, DL5 1TY",
  "delivered":       true,
  "courier":         "UPS",
  "courierId":       "15125425151261289",
  "date":            "2019-01-28T10:44:22.000Z",
}
for key, value := range properties {
  localDocument.SetProperty(key, value)
}

putLocalDocumentOptions := service.NewPutLocalDocumentOptions(
  "orders",
  "local-0007741142412418284",
)
putLocalDocumentOptions.SetDocument(&localDocument)

documentResult, response, err := service.PutLocalDocument(putLocalDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
