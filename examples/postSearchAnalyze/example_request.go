// section: code
postSearchAnalyzeOptions := service.NewPostSearchAnalyzeOptions(
  "english",
  "running is fun",
)

searchAnalyzeResult, response, err := service.PostSearchAnalyze(postSearchAnalyzeOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchAnalyzeResult, "", "  ")
fmt.Println(string(b))
