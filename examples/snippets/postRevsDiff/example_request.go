// section: code
postRevsDiffOptions := service.NewPostRevsDiffOptions(
  "orders",
  map[string][]string{
    "order00077": {
      "<1-missing-revision>",
      "<2-missing-revision>",
      "<3-possible-ancestor-revision>",
    },
  },
)

mapStringRevsDiff, response, err := service.PostRevsDiff(postRevsDiffOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(mapStringRevsDiff, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the example revisions in the POST body to be replaced with valid revisions.
