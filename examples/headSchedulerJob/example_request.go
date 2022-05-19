// section: code
headSchedulerJobOptions := service.NewHeadSchedulerJobOptions(
  "7b94915cd8c4a0173c77c55cd0443939+continuous",
)

response, err := service.HeadSchedulerJob(headSchedulerJobOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
