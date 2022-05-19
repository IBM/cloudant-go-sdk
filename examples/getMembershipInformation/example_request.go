// section: code
getMembershipInformationOptions := service.NewGetMembershipInformationOptions()

membershipInformation, response, err := service.GetMembershipInformation(getMembershipInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(membershipInformation, "", "  ")
fmt.Println(string(b))
