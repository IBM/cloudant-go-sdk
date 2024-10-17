// section: code
selector := map[string]interface{}{
  "userId": map[string]string{
    "$eq": "abc123",
  },
}

postPartitionFindOptions := service.NewPostPartitionFindOptions(
  "events",
  "ns1HJS13AMkK",
  selector,
)
postPartitionFindOptions.SetFields([]string{
  "productId", "eventType", "date",
})

findResult, response, err := service.PostPartitionFind(postPartitionFindOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(findResult, "", "  ")
fmt.Println(string(b))
