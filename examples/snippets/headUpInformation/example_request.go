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

headUpInformationOptions := service.NewHeadUpInformationOptions()

response, err := service.HeadUpInformation(headUpInformationOptions)
if err != nil {
  if response.GetStatusCode() == 503 {	
    fmt.Println("Service is unavailable: Status code:", response.GetStatusCode())
  } else {
    fmt.Println("Issue performing health check: Status code:", response.GetStatusCode())
  }
  return
}

fmt.Println("Service is up and healthy")
