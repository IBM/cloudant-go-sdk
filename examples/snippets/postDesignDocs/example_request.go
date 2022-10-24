// section: code
postDesignDocsOptions := service.NewPostDesignDocsOptions(
  "users",
)
postDesignDocsOptions.SetAttachments(true)

allDocsResult, response, err := service.PostDesignDocs(postDesignDocsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(allDocsResult, "", "  ")
fmt.Println(string(b))
