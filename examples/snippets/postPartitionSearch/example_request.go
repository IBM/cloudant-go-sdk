// section: code
postPartitionSearchOptions := service.NewPostPartitionSearchOptions(
  "events",
  "ns1HJS13AMkK",
  "checkout",
  "findByDate",
  "date:[2019-01-01T12:00:00.000Z TO 2019-01-31T12:00:00.000Z]",
)

searchResult, response, err := service.PostPartitionSearch(postPartitionSearchOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `findByDate` Cloudant Search partitioned index to exist. To create the design document with this index, see [Create or modify a design document.](#putdesigndocument)
