// section: code
postMissingRevsOptions := service.NewPostMissingRevsOptions(
  "orders",
  map[string][]string{
    "order00077": {"<order00077-existing-revision>", "<2-missing-revision>"},
  },
)

missingRevsResult, response, err := service.PostMissingRevs(postMissingRevsOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(missingRevsResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the example revisions in the POST body to be replaced with valid revisions.
