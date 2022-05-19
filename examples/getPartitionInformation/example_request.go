// section: code
getPartitionInformationOptions := service.NewGetPartitionInformationOptions(
  "products",
  "small-appliances",
)

partitionInformation, response, err := service.GetPartitionInformation(getPartitionInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(partitionInformation, "", "  ")
fmt.Println(string(b))
