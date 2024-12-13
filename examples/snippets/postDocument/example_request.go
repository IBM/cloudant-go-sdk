// section: code imports
import (
  "encoding/json"
  "fmt"

  "github.com/IBM/cloudant-go-sdk/cloudantv1"
  "github.com/IBM/go-sdk-core/v5/core"
)
// section: code
service, err := cloudantv1.NewCloudantV1(
  &cloudantv1.CloudantV1Options{},
)
if err != nil {
  panic(err)
}

productsDoc := cloudantv1.Document{
  ID: core.StringPtr("1000042"),
}
productsDoc.SetProperty("type", "product")
productsDoc.SetProperty("productId", "1000042")
productsDoc.SetProperty("brand", "Salter")
productsDoc.SetProperty("name", "Digital Kitchen Scales")
productsDoc.SetProperty("description", "Slim Colourful Design Electronic Cooking Appliance for Home / Kitchen, Weigh up to 5kg + Aquatronic for Liquids ml + fl. oz. 15Yr Guarantee - Green")
productsDoc.SetProperty("price", 14.99)
productsDoc.SetProperty("image", "assets/img/0gmsnghhew.jpg")

postDocumentOptions := service.NewPostDocumentOptions(
  "products",
)
postDocumentOptions.SetDocument(&productsDoc)

documentResult, response, err := service.PostDocument(postDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
