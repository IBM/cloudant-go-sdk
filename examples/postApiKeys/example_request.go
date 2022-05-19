// section: code
postApiKeysOptions := service.NewPostApiKeysOptions()

apiKeysResult, response, err := service.PostApiKeys(postApiKeysOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(apiKeysResult, "", "  ")
fmt.Println(string(b))
