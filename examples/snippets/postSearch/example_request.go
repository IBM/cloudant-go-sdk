// section: code
postSearchOptions := service.NewPostSearchOptions(
  "users",
  "allusers",
  "activeUsers",
  "name:Jane* AND active:True",
)

searchResult, response, err := service.PostSearch(postSearchOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `activeUsers` Cloudant Search index to exist. To create the design document with this index, see [Create or modify a design document.](#putdesigndocument)
