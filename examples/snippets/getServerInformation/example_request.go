// section: code
getServerInformationOptions := service.NewGetServerInformationOptions()

serverInformation, response, err := service.GetServerInformation(getServerInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(serverInformation, "", "  ")
fmt.Println(string(b))
