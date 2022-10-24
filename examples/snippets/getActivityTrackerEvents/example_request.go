// section: code
getActivityTrackerEventsOptions := service.NewGetActivityTrackerEventsOptions()

activityTrackerEvents, response, err := service.GetActivityTrackerEvents(getActivityTrackerEventsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(activityTrackerEvents, "", "  ")
fmt.Println(string(b))
