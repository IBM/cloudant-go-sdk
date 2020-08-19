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

	"github.com/IBM/go-sdk-core/v4/core"
	"github.ibm.com/cloudant/cloudant-go-sdk/cloudantv1"
)

func main() {
	// 1. Create a Cloudant client with "EXAMPLES" service name ============
	examplesClient, err := cloudantv1.NewCloudantV1UsingExternalConfig(
		&cloudantv1.CloudantV1Options{
			ServiceName: "EXAMPLES",
		},
	)
	if err != nil {
		panic(err)
	}
	// 2. Get server information ===========================================
	serverInformationResult, _, err := examplesClient.GetServerInformation(
		&cloudantv1.GetServerInformationOptions{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server Version: %s\n", *serverInformationResult.Version)
	// 3. Get database information for "animaldb" ==========================
	dbName := "animaldb"
	databaseInformationResult, _, err := examplesClient.GetDatabaseInformation(
		&cloudantv1.GetDatabaseInformationOptions{
			Db: &dbName,
		},
	)
	if err != nil {
		panic(err)
	}
	// 4. Show document count in database ==================================
	fmt.Printf("Document count in \"%s\" database is %d.\n",
		*databaseInformationResult.DbName,
		*databaseInformationResult.DocCount)
	// 5. Get zebra document out of the database by document id ============
	documentAboutZebraResult, _, err := examplesClient.GetDocument(
		&cloudantv1.GetDocumentOptions{
			Db:    &dbName,
			DocID: core.StringPtr("zebra"),
		},
	)
	if err != nil {
		panic(err)
	}
	// 6. Print out the Document content ===================================
	aboutZebraBuffer, _ := json.MarshalIndent(documentAboutZebraResult, "", "  ")
	fmt.Println(string(aboutZebraBuffer))
}
