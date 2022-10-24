// section: code
postExplainOptions := service.NewPostExplainOptions(
  "users",
  map[string]interface{}{
    "type": map[string]string{
      "$eq": "user",
    },
  },
)
postExplainOptions.SetExecutionStats(true)
postExplainOptions.SetLimit(10)

explainResult, response, err := service.PostExplain(postExplainOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(explainResult, "", "  ")
fmt.Println(string(b))
