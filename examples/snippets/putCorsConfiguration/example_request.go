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

putCorsConfigurationOptions := service.NewPutCorsConfigurationOptions([]string{
  "https://example.com",
})
putCorsConfigurationOptions.SetEnableCors(true)

ok, response, err := service.PutCorsConfiguration(putCorsConfigurationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
