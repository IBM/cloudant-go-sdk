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

postActivityTrackerEventsOptions := service.NewPostActivityTrackerEventsOptions(
  []string{
		cloudantv1.ActivityTrackerEventsTypesManagementConst,
		cloudantv1.ActivityTrackerEventsTypesDataConst,
	},
)

activityTrackerEvents, response, err := service.PostActivityTrackerEvents(postActivityTrackerEventsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(activityTrackerEvents, "", "  ")
fmt.Println(string(b))
