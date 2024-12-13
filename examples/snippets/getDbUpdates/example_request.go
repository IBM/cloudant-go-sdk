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

getDbUpdatesOptions := service.NewGetDbUpdatesOptions()
getDbUpdatesOptions.SetFeed("normal")
getDbUpdatesOptions.SetHeartbeat(10000)
getDbUpdatesOptions.SetSince("now")

dbUpdates, response, err := service.GetDbUpdates(getDbUpdatesOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(dbUpdates, "", "  ")
fmt.Println(string(b))
// section: markdown
// This request requires `server_admin` access.
