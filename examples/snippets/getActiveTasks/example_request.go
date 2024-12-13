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

getActiveTasksOptions := service.NewGetActiveTasksOptions()

activeTask, response, err := service.GetActiveTasks(getActiveTasksOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(activeTask, "", "  ")
fmt.Println(string(b))
