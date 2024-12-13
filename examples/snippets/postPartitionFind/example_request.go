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

selector := map[string]interface{}{
  "userId": map[string]string{
    "$eq": "abc123",
  },
}

postPartitionFindOptions := service.NewPostPartitionFindOptions(
  "events",
  "ns1HJS13AMkK",
  selector,
)
postPartitionFindOptions.SetFields([]string{
  "productId", "eventType", "date",
})

findResult, response, err := service.PostPartitionFind(postPartitionFindOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(findResult, "", "  ")
fmt.Println(string(b))
