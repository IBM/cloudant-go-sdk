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

getCorsInformationOptions := service.NewGetCorsInformationOptions()

corsConfiguration, response, err := service.GetCorsInformation(getCorsInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(corsConfiguration, "", "  ")
fmt.Println(string(b))
