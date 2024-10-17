// section: code
postPartitionViewOptions := service.NewPostPartitionViewOptions(
  "events",
  "ns1HJS13AMkK",
  "checkout",
  "byProductId",
)
postPartitionViewOptions.SetIncludeDocs(true)
postPartitionViewOptions.SetLimit(10)

viewResult, response, err := service.PostPartitionView(postPartitionViewOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(viewResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `byProductId` partitioned view to exist. To create the view, see [Create or modify a design document.](#putdesigndocument)
