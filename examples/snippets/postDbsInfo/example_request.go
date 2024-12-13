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

postDbsInfoOptions := service.NewPostDbsInfoOptions([]string{
  "products",
  "users",
  "orders",
})

dbsInfoResult, response, err := service.PostDbsInfo(postDbsInfoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(dbsInfoResult, "", "  ")
fmt.Println(string(b))
