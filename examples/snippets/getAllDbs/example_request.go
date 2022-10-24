// section: code
getAllDbsOptions := service.NewGetAllDbsOptions()

result, response, err := service.GetAllDbs(getAllDbsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(result, "", "  ")
fmt.Println(string(b))
