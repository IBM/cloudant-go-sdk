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

getAllDbsOptions := service.NewGetAllDbsOptions()

result, response, err := service.GetAllDbs(getAllDbsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(result, "", "  ")
fmt.Println(string(b))
