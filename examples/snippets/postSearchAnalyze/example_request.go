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

postSearchAnalyzeOptions := service.NewPostSearchAnalyzeOptions(
  "english",
  "running is fun",
)

searchAnalyzeResult, response, err := service.PostSearchAnalyze(postSearchAnalyzeOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchAnalyzeResult, "", "  ")
fmt.Println(string(b))
