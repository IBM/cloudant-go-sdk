// section: code
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
