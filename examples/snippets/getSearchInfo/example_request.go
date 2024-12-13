// section: code imports
import (
  "encoding/json"
  "fmt"

  "github.com/IBM/cloudant-go-sdk/cloudantv1"
)
// section: code
service, err := cloudantv1.NewCloudantV1(
  &cloudantv1.CloudantV1Options{},
)
if err != nil {
  panic(err)
}

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
