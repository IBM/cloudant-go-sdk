// section: code
docID := "order00058"

postBulkGetOptions := service.NewPostBulkGetOptions(
  "orders",
  []cloudantv1.BulkGetQueryDocument{
    {
      ID: &docID,
      AttsSince: []string{"1-99b02e08da151943c2dcb40090160bb8"},
    },
  },
)

bulkGetResult, response, err := service.PostBulkGet(postBulkGetOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(bulkGetResult, "", "  ")
fmt.Println(string(b))
