// section: code imports
import (
  "encoding/json"
  "fmt"
  "io"
  "os"

  "github.com/IBM/cloudant-go-sdk/cloudantv1"
)
// section: code
service, err := cloudantv1.NewCloudantV1(
  &cloudantv1.CloudantV1Options{},
)
if err != nil {
  panic(err)
}

postChangesOptions := service.NewPostChangesOptions(
  "orders",
)

changesResult, response, err := service.PostChangesAsStream(postChangesOptions)
if err != nil {
  panic(err)
}

if changesResult != nil {
  defer changesResult.Close()
  outFile, err := os.Create("result.json")
  if err != nil {
    panic(err)
  }
  defer outFile.Close()
  if _, err = io.Copy(outFile, changesResult); err != nil {
    panic(err)
  }
}
