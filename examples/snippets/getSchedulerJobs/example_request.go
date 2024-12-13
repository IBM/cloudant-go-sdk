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

getSchedulerJobsOptions := service.NewGetSchedulerJobsOptions()
getSchedulerJobsOptions.SetLimit(100)

schedulerJobsResult, response, err := service.GetSchedulerJobs(getSchedulerJobsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerJobsResult, "", "  ")
fmt.Println(string(b))
