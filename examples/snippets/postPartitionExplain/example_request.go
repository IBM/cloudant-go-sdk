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

postPartitionExplainOptions := service.NewPostPartitionExplainOptions(
  "events",
  "ns1HJS13AMkK",
  selector,
)
postExplainOptions.SetExecutionStats(true)
postExplainOptions.SetLimit(10)


explainResult, response, err := service.PostPartitionExplain(postPartitionExplainOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(explainResult, "", "  ")
fmt.Println(string(b))
