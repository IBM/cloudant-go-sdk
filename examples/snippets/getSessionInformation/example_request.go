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

getSessionInformationOptions := service.NewGetSessionInformationOptions()

sessionInformation, response, err := service.GetSessionInformation(getSessionInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(sessionInformation, "", "  ")
fmt.Println(string(b))
// section: markdown
// For more details on Session Authentication, see [Authentication.](#authentication)
