// section: code
getDocumentShardsInfoOptions := service.NewGetDocumentShardsInfoOptions(
  "products",
  "small-appliances:1000042",
)

documentShardInfo, response, err := service.GetDocumentShardsInfo(getDocumentShardsInfoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentShardInfo, "", "  ")
fmt.Println(string(b))
