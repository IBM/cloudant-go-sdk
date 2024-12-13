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

getCurrentThroughputInformationOptions := service.NewGetCurrentThroughputInformationOptions()

currentThroughputInformation, response, err := service.GetCurrentThroughputInformation(getCurrentThroughputInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(currentThroughputInformation, "", "  ")
fmt.Println(string(b))
