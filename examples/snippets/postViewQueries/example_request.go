// section: code imports
import (
  "encoding/json"
  "fmt"

  "github.com/IBM/cloudant-go-sdk/cloudantv1"
  "github.com/IBM/go-sdk-core/v5/core"
)
// section: code
service, err := cloudantv1.NewCloudantV1(
  &cloudantv1.CloudantV1Options{},
)
if err != nil {
  panic(err)
}

postViewQueriesOptions := service.NewPostViewQueriesOptions(
  "users",
  "allusers",
  "getVerifiedEmails",
  []cloudantv1.ViewQuery{
    {
      IncludeDocs: core.BoolPtr(true),
      Limit:       core.Int64Ptr(5),
    },
    {
      Descending: core.BoolPtr(true),
      Skip:       core.Int64Ptr(1),
    },
  },
)

viewQueriesResult, response, err := service.PostViewQueries(postViewQueriesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(viewQueriesResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `getVerifiedEmails` view to exist. To create the design document with this view, see [Create or modify a design document.](#putdesigndocument)
