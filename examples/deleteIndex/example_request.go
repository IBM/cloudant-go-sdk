// section: code
deleteIndexOptions := service.NewDeleteIndexOptions(
  "users",
  "json-index",
  "json",
  "getUserByName",
)

ok, response, err := service.DeleteIndex(deleteIndexOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example will fail if `getUserByName` index doesn't exist. To create the index, see [Create a new index on a database.](#postindex)
