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

getIndexesInformationOptions := service.NewGetIndexesInformationOptions(
  "users",
)

indexesInformation, response, err := service.GetIndexesInformation(getIndexesInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(indexesInformation, "", "  ")
fmt.Println(string(b))
