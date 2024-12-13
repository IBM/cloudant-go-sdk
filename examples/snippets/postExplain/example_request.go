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

postExplainOptions := service.NewPostExplainOptions(
  "users",
  map[string]interface{}{
    "type": map[string]string{
      "$eq": "user",
    },
  },
)
postExplainOptions.SetExecutionStats(true)
postExplainOptions.SetLimit(10)

explainResult, response, err := service.PostExplain(postExplainOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(explainResult, "", "  ")
fmt.Println(string(b))
