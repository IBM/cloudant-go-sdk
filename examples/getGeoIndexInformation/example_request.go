// section: code
getGeoIndexInformationOptions := service.NewGetGeoIndexInformationOptions(
  "stores",
  "places",
  "pointsInEngland",
)

geoIndexInformation, response, err := service.GetGeoIndexInformation(getGeoIndexInformationOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(geoIndexInformation, "", "  ")
fmt.Println(string(b))
