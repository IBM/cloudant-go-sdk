// section: code
localDocument := cloudantv1.Document{}
properties := map[string]interface{}{
  "type":            "order",
  "user":            "Bob Smith",
  "orderid":         "0007741142412418284",
  "userid":          "abc123",
  "total":           214.98,
  "deliveryAddress": "19 Front Street, Darlington, DL5 1TY",
  "delivered":       true,
  "courier":         "UPS",
  "courierid":       "15125425151261289",
  "date":            "2019-01-28T10:44:22.000Z",
}
for key, value := range properties {
  localDocument.SetProperty(key, value)
}

putLocalDocumentOptions := service.NewPutLocalDocumentOptions(
  "orders",
  "local-0007741142412418284",
)
putLocalDocumentOptions.SetDocument(&localDocument)

documentResult, response, err := service.PutLocalDocument(putLocalDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
