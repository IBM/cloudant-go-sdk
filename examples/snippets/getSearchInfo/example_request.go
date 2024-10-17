// section: code
getSearchInfoOptions := service.NewGetSearchInfoOptions(
  "events",
  "checkout",
  "findByDate",
)

searchInfoResult, response, err := service.GetSearchInfo(getSearchInfoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchInfoResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `findByDate` Cloudant Search partitioned index to exist. To create the design document with this index, see [Create or modify a design document.](#putdesigndocument)
