// section: code
getSecurityOptions := service.NewGetSecurityOptions(
  "products",
)

security, response, err := service.GetSecurity(getSecurityOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(security, "", "  ")
fmt.Println(string(b))
