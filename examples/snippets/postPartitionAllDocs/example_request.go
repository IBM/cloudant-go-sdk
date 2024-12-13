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

postPartitionAllDocsOptions := service.NewPostPartitionAllDocsOptions(
  "events",
  "ns1HJS13AMkK",
)
postPartitionAllDocsOptions.SetIncludeDocs(true)

allDocsResult, response, err := service.PostPartitionAllDocs(postPartitionAllDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsResult, "", "  ")
fmt.Println(string(b))
