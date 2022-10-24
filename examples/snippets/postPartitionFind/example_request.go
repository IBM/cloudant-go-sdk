// section: code
selector := map[string]interface{}{
  "type": map[string]string{
    "$eq": "product",
  },
}

postPartitionFindOptions := service.NewPostPartitionFindOptions(
  "products",
  "small-appliances",
  selector,
)
postPartitionFindOptions.SetFields([]string{
  "productid", "name", "description",
})

findResult, response, err := service.PostPartitionFind(postPartitionFindOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(findResult, "", "  ")
fmt.Println(string(b))
