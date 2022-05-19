// section: code
getGeoOptions := service.NewGetGeoOptions(
  "stores",
  "places",
  "pointsInEngland",
)
getGeoOptions.SetBbox("-50.52,-4.46,54.59,1.45")
getGeoOptions.SetIncludeDocs(true)
getGeoOptions.SetNearest(true)

geoResult, response, err := service.GetGeo(getGeoOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(geoResult, "", "  ")
fmt.Println(string(b))
