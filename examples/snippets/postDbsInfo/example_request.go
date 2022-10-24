// section: code
postDbsInfoOptions := service.NewPostDbsInfoOptions([]string{
  "products",
  "users",
  "orders",
})

dbsInfoResult, response, err := service.PostDbsInfo(postDbsInfoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(dbsInfoResult, "", "  ")
fmt.Println(string(b))
