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

postViewOptions := service.NewPostViewOptions(
  "users",
  "allusers",
  "getVerifiedEmails",
)

viewResult, response, err := service.PostView(postViewOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(viewResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example requires the `getVerifiedEmails` view to exist. To create the design document with this view, see [Create or modify a design document.](#putdesigndocument)
