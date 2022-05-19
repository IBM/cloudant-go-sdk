// section: code
putCorsConfigurationOptions := service.NewPutCorsConfigurationOptions([]string{
  "https://example.com",
})
putCorsConfigurationOptions.SetEnableCors(true)

ok, response, err := service.PutCorsConfiguration(putCorsConfigurationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
