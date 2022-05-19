// section: code
postGeoCleanupOptions := service.NewPostGeoCleanupOptions(
  "stores",
)

ok, response, err := service.PostGeoCleanup(postGeoCleanupOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(ok, "", "  ")
fmt.Println(string(b))
