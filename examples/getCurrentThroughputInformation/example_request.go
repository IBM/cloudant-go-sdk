// section: code
getCurrentThroughputInformationOptions := service.NewGetCurrentThroughputInformationOptions()

currentThroughputInformation, response, err := service.GetCurrentThroughputInformation(getCurrentThroughputInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(currentThroughputInformation, "", "  ")
fmt.Println(string(b))
