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

getDocumentShardsInfoOptions := service.NewGetDocumentShardsInfoOptions(
  "products",
  "1000042",
)

documentShardInfo, response, err := service.GetDocumentShardsInfo(getDocumentShardsInfoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentShardInfo, "", "  ")
fmt.Println(string(b))
