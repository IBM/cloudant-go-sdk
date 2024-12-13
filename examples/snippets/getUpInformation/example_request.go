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

getUpInformationOptions := service.NewGetUpInformationOptions()

upInformation, response, err := service.GetUpInformation(getUpInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(upInformation, "", "  ")
fmt.Println(string(b))
