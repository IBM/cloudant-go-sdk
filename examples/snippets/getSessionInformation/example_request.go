// section: code
getSessionInformationOptions := service.NewGetSessionInformationOptions()

sessionInformation, response, err := service.GetSessionInformation(getSessionInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(sessionInformation, "", "  ")
fmt.Println(string(b))
// section: markdown
// For more details on Session Authentication, see [Authentication.](#authentication)
