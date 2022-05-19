// section: code
postPartitionSearchOptions := service.NewPostPartitionSearchOptions(
  "products",
  "small-appliances",
  "appliances",
  "findByPrice",
  "price:[14 TO 20]",
)

searchResult, response, err := service.PostPartitionSearch(postPartitionSearchOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `findByPrice` Cloudant Search partitioned index to exist. To create the design document with this index, see [Create or modify a design document.](#putdesigndocument)
