// section: code
headDatabaseOptions := service.NewHeadDatabaseOptions(
  "products",
)

response, err := service.HeadDatabase(headDatabaseOptions)
if err != nil {
  panic(err)
}

fmt.Println(response.StatusCode)
