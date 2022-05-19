// section: code
postActivityTrackerEventsOptions := service.NewPostActivityTrackerEventsOptions(
  []string{"management"},
)

activityTrackerEvents, response, err := service.PostActivityTrackerEvents(postActivityTrackerEventsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(activityTrackerEvents, "", "  ")
fmt.Println(string(b))
