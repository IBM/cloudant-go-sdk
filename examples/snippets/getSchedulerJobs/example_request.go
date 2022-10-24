// section: code
getSchedulerJobsOptions := service.NewGetSchedulerJobsOptions()
getSchedulerJobsOptions.SetLimit(100)

schedulerJobsResult, response, err := service.GetSchedulerJobs(getSchedulerJobsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerJobsResult, "", "  ")
fmt.Println(string(b))
