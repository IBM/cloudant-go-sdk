// section: code
putDatabaseOptions := service.NewPutDatabaseOptions(
  "products",
)
putDatabaseOptions.SetPartitioned(true)

ok, response, err := service.PutDatabase(putDatabaseOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
