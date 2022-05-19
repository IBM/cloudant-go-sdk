// section: code
getUuidsOptions := service.NewGetUuidsOptions()
getUuidsOptions.SetCount(10)

uuidsResult, response, err := service.GetUuids(getUuidsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(uuidsResult, "", "  ")
fmt.Println(string(b))
