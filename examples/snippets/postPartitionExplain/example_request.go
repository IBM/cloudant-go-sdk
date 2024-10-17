// section: code
selector := map[string]interface{}{
  "userId": map[string]string{
    "$eq": "abc123",
  },
}

postPartitionExplainOptions := service.NewPostPartitionExplainOptions(
  "events",
  "ns1HJS13AMkK",
  selector,
)
postExplainOptions.SetExecutionStats(true)
postExplainOptions.SetLimit(10)


explainResult, response, err := service.PostPartitionExplain(postPartitionExplainOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(explainResult, "", "  ")
fmt.Println(string(b))
