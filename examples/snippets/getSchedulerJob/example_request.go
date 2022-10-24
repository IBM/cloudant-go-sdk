// section: code
getSchedulerJobOptions := service.NewGetSchedulerJobOptions(
  "7b94915cd8c4a0173c77c55cd0443939+continuous",
)

schedulerJob, response, err := service.GetSchedulerJob(getSchedulerJobOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerJob, "", "  ")
fmt.Println(string(b))
