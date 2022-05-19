// section: code
putCapacityThroughputConfigurationOptions := service.NewPutCapacityThroughputConfigurationOptions(
  1,
)

capacityThroughputConfiguration, response, err := service.PutCapacityThroughputConfiguration(putCapacityThroughputConfigurationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(capacityThroughputConfiguration, "", "  ")
fmt.Println(string(b))
