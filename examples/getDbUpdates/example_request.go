// section: code
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
