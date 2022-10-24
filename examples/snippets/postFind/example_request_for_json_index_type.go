// section: code
postFindOptions := service.NewPostFindOptions(
  "users",
  map[string]interface{}{
    "email_verified": map[string]bool{
      "$eq": true,
    },
  },
)
postFindOptions.SetFields(
  []string{"_id", "type", "name", "email"},
)
postFindOptions.SetSort([]map[string]string{{"email": "desc"}})
postFindOptions.SetLimit(3)

findResult, response, err := service.PostFind(postFindOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(findResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `getUserByEmail` Cloudant Query "json" index to exist. To create the index, see [Create a new index on a database.](#postindex)
