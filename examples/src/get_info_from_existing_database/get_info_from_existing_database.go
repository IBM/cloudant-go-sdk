// © Copyright IBM Corporation 2020. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

func main() {
	// 1. Create a Cloudant client with "EXAMPLES" service name ============
	client, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{
			ServiceName: "EXAMPLES",
		},
	)
	if err != nil {
		panic(err)
	}
	// 2. Get server information ===========================================
	serverInformationResult, _, err := client.GetServerInformation(
		client.NewGetServerInformationOptions(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server Version: %s\n", *serverInformationResult.Version)
	// 3. Get database information for "animaldb" ==========================
	dbName := "animaldb"
	databaseInformationResult, _, err := client.GetDatabaseInformation(
		client.NewGetDatabaseInformationOptions(
			dbName,
		),
	)
	if err != nil {
		panic(err)
	}
	// 4. Show document count in database ==================================
	fmt.Printf("Document count in \"%s\" database is %d.\n",
		*databaseInformationResult.DbName,
		*databaseInformationResult.DocCount)
	// 5. Get zebra document out of the database by document id ============
	documentAboutZebraResult, _, err := client.GetDocument(
		client.NewGetDocumentOptions(
			dbName,
			"zebra",
		),
	)
	if err != nil {
		panic(err)
	}
	// 6. Print out the Document content ===================================
	aboutZebraBuffer, _ := json.MarshalIndent(documentAboutZebraResult, "", "  ")
	fmt.Println(string(aboutZebraBuffer))
}
