// section: code
getShardsInformationOptions := service.NewGetShardsInformationOptions(
  "products",
)

shardsInformation, response, err := service.GetShardsInformation(getShardsInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(shardsInformation, "", "  ")
fmt.Println(string(b))
