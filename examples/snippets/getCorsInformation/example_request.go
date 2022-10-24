// section: code
getCorsInformationOptions := service.NewGetCorsInformationOptions()

corsConfiguration, response, err := service.GetCorsInformation(getCorsInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(corsConfiguration, "", "  ")
fmt.Println(string(b))
