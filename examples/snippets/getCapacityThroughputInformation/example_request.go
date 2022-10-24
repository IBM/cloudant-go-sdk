// section: code
getCapacityThroughputInformationOptions := service.NewGetCapacityThroughputInformationOptions()

capacityThroughputInformation, response, err := service.GetCapacityThroughputInformation(getCapacityThroughputInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(capacityThroughputInformation, "", "  ")
fmt.Println(string(b))
