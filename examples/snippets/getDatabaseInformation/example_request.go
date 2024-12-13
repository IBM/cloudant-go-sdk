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

getDatabaseInformationOptions := service.NewGetDatabaseInformationOptions(
  "products",
)

databaseInformation, response, err := service.GetDatabaseInformation(getDatabaseInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(databaseInformation, "", "  ")
fmt.Println(string(b))
