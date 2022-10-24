// section: code
getSchedulerDocumentOptions := service.NewGetSchedulerDocumentOptions(
  "repldoc-example",
)

schedulerDocument, response, err := service.GetSchedulerDocument(getSchedulerDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(schedulerDocument, "", "  ")
fmt.Println(string(b))
