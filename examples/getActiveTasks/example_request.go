// section: code
getActiveTasksOptions := service.NewGetActiveTasksOptions()

activeTask, response, err := service.GetActiveTasks(getActiveTasksOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(activeTask, "", "  ")
fmt.Println(string(b))
