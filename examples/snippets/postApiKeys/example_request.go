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

postApiKeysOptions := service.NewPostApiKeysOptions()

apiKeysResult, response, err := service.PostApiKeys(postApiKeysOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(apiKeysResult, "", "  ")
fmt.Println(string(b))
