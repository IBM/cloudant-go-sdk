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

postSearchOptions := service.NewPostSearchOptions(
  "users",
  "allusers",
  "activeUsers",
  "name:Jane* AND active:True",
)

searchResult, response, err := service.PostSearch(postSearchOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(searchResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `activeUsers` Cloudant Search index to exist. To create the design document with this index, see [Create or modify a design document.](#putdesigndocument)
