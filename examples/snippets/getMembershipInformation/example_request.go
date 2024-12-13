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

getMembershipInformationOptions := service.NewGetMembershipInformationOptions()

membershipInformation, response, err := service.GetMembershipInformation(getMembershipInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(membershipInformation, "", "  ")
fmt.Println(string(b))
