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

getSchedulerDocsOptions := service.NewGetSchedulerDocsOptions()
getSchedulerDocsOptions.SetLimit(100)
getSchedulerDocsOptions.SetStates([]string{"completed"})

schedulerDocsResult, response, err := service.GetSchedulerDocs(getSchedulerDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerDocsResult, "", "  ")
fmt.Println(string(b))
