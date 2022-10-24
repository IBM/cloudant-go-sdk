// section: code
deleteDatabaseOptions := service.NewDeleteDatabaseOptions(
  "<db-name>",
)

ok, response, err := service.DeleteDatabase(deleteDatabaseOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
