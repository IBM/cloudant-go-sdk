// section: code
getUpInformationOptions := service.NewGetUpInformationOptions()

upInformation, response, err := service.GetUpInformation(getUpInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(upInformation, "", "  ")
fmt.Println(string(b))
