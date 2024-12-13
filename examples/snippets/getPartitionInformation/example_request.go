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

getPartitionInformationOptions := service.NewGetPartitionInformationOptions(
  "events",
  "ns1HJS13AMkK",
)

partitionInformation, response, err := service.GetPartitionInformation(getPartitionInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(partitionInformation, "", "  ")
fmt.Println(string(b))
