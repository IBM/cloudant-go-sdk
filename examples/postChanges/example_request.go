// section: code
postChangesOptions := service.NewPostChangesOptions(
  "orders",
)

changesResult, response, err := service.PostChanges(postChangesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(changesResult, "", "  ")
fmt.Println(string(b))
