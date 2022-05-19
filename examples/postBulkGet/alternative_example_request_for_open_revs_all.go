// section: markdown
// This example requires an import for `github.com/IBM/go-sdk-core/v5/core`.
// section: code
postBulkGetOptions := service.NewPostBulkGetOptions(
  "orders",
  []cloudantv1.BulkGetQueryDocument{{ID: core.StringPtr("order00067")}},
)

bulkGetResult, response, err := service.PostBulkGet(postBulkGetOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(bulkGetResult, "", "  ")
fmt.Println(string(b))
