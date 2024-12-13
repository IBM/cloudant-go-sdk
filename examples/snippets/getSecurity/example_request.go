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

getSecurityOptions := service.NewGetSecurityOptions(
  "products",
)

security, response, err := service.GetSecurity(getSecurityOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(security, "", "  ")
fmt.Println(string(b))
