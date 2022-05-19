// section: code
getDatabaseInformationOptions := service.NewGetDatabaseInformationOptions(
  "products",
)

databaseInformation, response, err := service.GetDatabaseInformation(getDatabaseInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(databaseInformation, "", "  ")
fmt.Println(string(b))
