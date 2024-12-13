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

getDesignDocumentInformationOptions := service.NewGetDesignDocumentInformationOptions(
  "products",
  "appliances",
)

designDocumentInformation, response, err := service.GetDesignDocumentInformation(getDesignDocumentInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(designDocumentInformation, "", "  ")
fmt.Println(string(b))
