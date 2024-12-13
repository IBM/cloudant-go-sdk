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

getUuidsOptions := service.NewGetUuidsOptions()
getUuidsOptions.SetCount(10)

uuidsResult, response, err := service.GetUuids(getUuidsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(uuidsResult, "", "  ")
fmt.Println(string(b))
