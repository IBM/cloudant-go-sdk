// section: code
emailViewMapReduce, err := service.NewDesignDocumentViewsMapReduce(
  "function(doc) {" +
    "if(doc.email_verified  === true){ emit(doc.email, [doc.name, doc.email_verified, doc.joined])" +
  "}",
)
if err != nil {
  panic(err)
}

userIndexDefinition, err := service.NewSearchIndexDefinition(
  "function(doc) {" +
    "index(\"name\", doc.name); index(\"active\", doc.active);" +
  "}",
)
if err != nil {
  panic(err)
}

designDocument := &cloudantv1.DesignDocument{
  Views: map[string]cloudantv1.DesignDocumentViewsMapReduce{
    "getVerifiedEmails": *emailViewMapReduce,
  },
  Indexes: map[string]cloudantv1.SearchIndexDefinition{
    "activeUsers": *userIndexDefinition,
  },
}

putDesignDocumentOptions := service.NewPutDesignDocumentOptions(
  "users",
  "allusers",
  designDocument,
)

documentResult, response, err := service.PutDesignDocument(putDesignDocumentOptions)
if err != nil {
  panic(err)
}

b, _ := json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))

applianceProdIdViewMapReduce, err := service.NewDesignDocumentViewsMapReduce(
  "function(doc) {" +
    "emit(doc.productid, [doc.brand, doc.name, doc.description])" +
  "}",
)
if err != nil {
  panic(err)
}

priceIndexDefinition, err := service.NewSearchIndexDefinition(
  "function(doc) {" +
    "index(\"price\", doc.price);" +
  "}",
)
if err != nil {
  panic(err)
}

partitionedDesignDocument := &cloudantv1.DesignDocument{
  Views: map[string]cloudantv1.DesignDocumentViewsMapReduce{
    "byApplianceProdId": *applianceProdIdViewMapReduce,
  },
  Indexes: map[string]cloudantv1.SearchIndexDefinition{
    "findByPrice": *priceIndexDefinition,
  },
}

putPartitionedDesignDocumentOptions := service.NewPutDesignDocumentOptions(
  "products",
  "appliances",
  partitionedDesignDocument,
)

documentResult, response, err = service.PutDesignDocument(putPartitionedDesignDocumentOptions)
if err != nil {
  panic(err)
}

b, _ = json.MarshalIndent(documentResult, "", "  ")
fmt.Println(string(b))
// section: markdown
// This example creates `allusers` design document in the `users` database and `appliances` design document in the partitioned `products` database.
