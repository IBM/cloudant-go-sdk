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

getSchedulerJobOptions := service.NewGetSchedulerJobOptions(
  "7b94915cd8c4a0173c77c55cd0443939+continuous",
)

schedulerJob, response, err := service.GetSchedulerJob(getSchedulerJobOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerJob, "", "  ")
fmt.Println(string(b))
