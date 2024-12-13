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

headDatabaseOptions := service.NewHeadDatabaseOptions(
  "products",
)

response, err := service.HeadDatabase(headDatabaseOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
