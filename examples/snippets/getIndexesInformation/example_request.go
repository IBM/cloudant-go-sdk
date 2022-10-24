// section: code
getIndexesInformationOptions := service.NewGetIndexesInformationOptions(
  "users",
)

indexesInformation, response, err := service.GetIndexesInformation(getIndexesInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(indexesInformation, "", "  ")
fmt.Println(string(b))
