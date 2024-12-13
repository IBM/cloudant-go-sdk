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

postChangesOptions := service.NewPostChangesOptions(
  "orders",
)

changesResult, response, err := service.PostChanges(postChangesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(changesResult, "", "  ")
fmt.Println(string(b))
