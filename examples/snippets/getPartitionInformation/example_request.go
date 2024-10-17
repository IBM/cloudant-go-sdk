// section: code
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
