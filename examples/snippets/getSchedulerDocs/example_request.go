// section: code
getSchedulerDocsOptions := service.NewGetSchedulerDocsOptions()
getSchedulerDocsOptions.SetLimit(100)
getSchedulerDocsOptions.SetStates([]string{"completed"})

schedulerDocsResult, response, err := service.GetSchedulerDocs(getSchedulerDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerDocsResult, "", "  ")
fmt.Println(string(b))
