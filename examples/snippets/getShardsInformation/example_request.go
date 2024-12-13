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

getShardsInformationOptions := service.NewGetShardsInformationOptions(
  "products",
)

shardsInformation, response, err := service.GetShardsInformation(getShardsInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(shardsInformation, "", "  ")
fmt.Println(string(b))
