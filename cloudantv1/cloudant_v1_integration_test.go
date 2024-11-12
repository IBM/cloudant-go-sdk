//go:build integration

/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cloudantv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the cloudantv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`CloudantV1 Integration Tests`, func() {
	const externalConfigFile = "../cloudant_v1.env"

	var (
		err             error
		cloudantService *cloudantv1.CloudantV1
		serviceURL      string
		config          map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(cloudantv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			cloudantServiceOptions := &cloudantv1.CloudantV1Options{}

			cloudantService, err = cloudantv1.NewCloudantV1UsingExternalConfig(cloudantServiceOptions)
			Expect(err).To(BeNil())
			Expect(cloudantService).ToNot(BeNil())
			Expect(cloudantService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			cloudantService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetServerInformation - Retrieve server instance information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServerInformation(getServerInformationOptions *GetServerInformationOptions)`, func() {
			getServerInformationOptions := &cloudantv1.GetServerInformationOptions{}

			serverInformation, response, err := cloudantService.GetServerInformation(getServerInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serverInformation).ToNot(BeNil())
		})
	})

	Describe(`GetUuids - Retrieve one or more UUIDs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUuids(getUuidsOptions *GetUuidsOptions)`, func() {
			getUuidsOptions := &cloudantv1.GetUuidsOptions{
				Count: core.Int64Ptr(int64(1)),
			}

			uuidsResult, response, err := cloudantService.GetUuids(getUuidsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(uuidsResult).ToNot(BeNil())
		})
	})

	Describe(`GetCapacityThroughputInformation - Retrieve provisioned throughput capacity information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCapacityThroughputInformation(getCapacityThroughputInformationOptions *GetCapacityThroughputInformationOptions)`, func() {
			getCapacityThroughputInformationOptions := &cloudantv1.GetCapacityThroughputInformationOptions{}

			capacityThroughputInformation, response, err := cloudantService.GetCapacityThroughputInformation(getCapacityThroughputInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(capacityThroughputInformation).ToNot(BeNil())
		})
	})

	Describe(`PutCapacityThroughputConfiguration - Update the target provisioned throughput capacity`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutCapacityThroughputConfiguration(putCapacityThroughputConfigurationOptions *PutCapacityThroughputConfigurationOptions)`, func() {
			putCapacityThroughputConfigurationOptions := &cloudantv1.PutCapacityThroughputConfigurationOptions{
				Blocks: core.Int64Ptr(int64(10)),
			}

			capacityThroughputInformation, response, err := cloudantService.PutCapacityThroughputConfiguration(putCapacityThroughputConfigurationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(capacityThroughputInformation).ToNot(BeNil())
		})
	})

	Describe(`GetDbUpdates - Retrieve change events for all databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDbUpdates(getDbUpdatesOptions *GetDbUpdatesOptions)`, func() {
			getDbUpdatesOptions := &cloudantv1.GetDbUpdatesOptions{
				Descending: core.BoolPtr(false),
				Feed:       core.StringPtr("normal"),
				Heartbeat:  core.Int64Ptr(int64(0)),
				Limit:      core.Int64Ptr(int64(0)),
				Timeout:    core.Int64Ptr(int64(60000)),
				Since:      core.StringPtr("0"),
			}

			dbUpdates, response, err := cloudantService.GetDbUpdates(getDbUpdatesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dbUpdates).ToNot(BeNil())
		})
	})

	Describe(`PostChanges - Query the database document changes feed`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostChanges(postChangesOptions *PostChangesOptions)`, func() {
			postChangesOptions := &cloudantv1.PostChangesOptions{
				Db:              core.StringPtr("testString"),
				DocIds:          []string{"0007741142412418284"},
				Fields:          []string{"testString"},
				Selector:        map[string]interface{}{"anyKey": "anyValue"},
				LastEventID:     core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				Feed:            core.StringPtr("normal"),
				Filter:          core.StringPtr("testString"),
				Heartbeat:       core.Int64Ptr(int64(0)),
				IncludeDocs:     core.BoolPtr(false),
				Limit:           core.Int64Ptr(int64(0)),
				SeqInterval:     core.Int64Ptr(int64(1)),
				Since:           core.StringPtr("0"),
				Style:           core.StringPtr("main_only"),
				Timeout:         core.Int64Ptr(int64(60000)),
				View:            core.StringPtr("testString"),
			}

			changesResult, response, err := cloudantService.PostChanges(postChangesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(changesResult).ToNot(BeNil())
		})
	})

	Describe(`PostChangesAsStream - Query the database document changes feed as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostChangesAsStream(postChangesOptions *PostChangesOptions)`, func() {
			postChangesOptions := &cloudantv1.PostChangesOptions{
				Db:              core.StringPtr("testString"),
				DocIds:          []string{"0007741142412418284"},
				Fields:          []string{"testString"},
				Selector:        map[string]interface{}{"anyKey": "anyValue"},
				LastEventID:     core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				Feed:            core.StringPtr("normal"),
				Filter:          core.StringPtr("testString"),
				Heartbeat:       core.Int64Ptr(int64(0)),
				IncludeDocs:     core.BoolPtr(false),
				Limit:           core.Int64Ptr(int64(0)),
				SeqInterval:     core.Int64Ptr(int64(1)),
				Since:           core.StringPtr("0"),
				Style:           core.StringPtr("main_only"),
				Timeout:         core.Int64Ptr(int64(60000)),
				View:            core.StringPtr("testString"),
			}

			result, response, err := cloudantService.PostChangesAsStream(postChangesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`HeadDatabase - Retrieve the HTTP headers for a database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HeadDatabase(headDatabaseOptions *HeadDatabaseOptions)`, func() {
			headDatabaseOptions := &cloudantv1.HeadDatabaseOptions{
				Db: core.StringPtr("testString"),
			}

			response, err := cloudantService.HeadDatabase(headDatabaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetAllDbs - Query a list of all database names in the instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAllDbs(getAllDbsOptions *GetAllDbsOptions)`, func() {
			getAllDbsOptions := &cloudantv1.GetAllDbsOptions{
				Descending: core.BoolPtr(false),
				EndKey:     core.StringPtr("testString"),
				Limit:      core.Int64Ptr(int64(0)),
				Skip:       core.Int64Ptr(int64(0)),
				StartKey:   core.StringPtr("testString"),
			}

			result, response, err := cloudantService.GetAllDbs(getAllDbsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PostDbsInfo - Query information about multiple databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostDbsInfo(postDbsInfoOptions *PostDbsInfoOptions)`, func() {
			postDbsInfoOptions := &cloudantv1.PostDbsInfoOptions{
				Keys: []string{"products", "users", "orders"},
			}

			dbsInfoResult, response, err := cloudantService.PostDbsInfo(postDbsInfoOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dbsInfoResult).ToNot(BeNil())
		})
	})

	Describe(`GetDatabaseInformation - Retrieve information about a database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDatabaseInformation(getDatabaseInformationOptions *GetDatabaseInformationOptions)`, func() {
			getDatabaseInformationOptions := &cloudantv1.GetDatabaseInformationOptions{
				Db: core.StringPtr("testString"),
			}

			databaseInformation, response, err := cloudantService.GetDatabaseInformation(getDatabaseInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databaseInformation).ToNot(BeNil())
		})
	})

	Describe(`PutDatabase - Create a database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutDatabase(putDatabaseOptions *PutDatabaseOptions)`, func() {
			putDatabaseOptions := &cloudantv1.PutDatabaseOptions{
				Db:          core.StringPtr("testString"),
				Partitioned: core.BoolPtr(false),
				Q:           core.Int64Ptr(int64(26)),
			}

			ok, response, err := cloudantService.PutDatabase(putDatabaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(ok).ToNot(BeNil())
		})
	})

	Describe(`HeadDocument - Retrieve the HTTP headers for the document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HeadDocument(headDocumentOptions *HeadDocumentOptions)`, func() {
			headDocumentOptions := &cloudantv1.HeadDocumentOptions{
				Db:          core.StringPtr("testString"),
				DocID:       core.StringPtr("testString"),
				IfNoneMatch: core.StringPtr("testString"),
				Latest:      core.BoolPtr(false),
				Rev:         core.StringPtr("testString"),
			}

			response, err := cloudantService.HeadDocument(headDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`PostDocument - Create or modify a document in a database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostDocument(postDocumentOptions *PostDocumentOptions)`, func() {
			attachmentModel := &cloudantv1.Attachment{
				ContentType:   core.StringPtr("testString"),
				Data:          CreateMockByteArray("VGhpcyBpcyBhIG1vY2sgYnl0ZSBhcnJheSB2YWx1ZS4="),
				Digest:        core.StringPtr("testString"),
				EncodedLength: core.Int64Ptr(int64(0)),
				Encoding:      core.StringPtr("testString"),
				Follows:       core.BoolPtr(true),
				Length:        core.Int64Ptr(int64(0)),
				Revpos:        core.Int64Ptr(int64(1)),
				Stub:          core.BoolPtr(true),
			}

			revisionsModel := &cloudantv1.Revisions{
				Ids:   []string{"testString"},
				Start: core.Int64Ptr(int64(1)),
			}

			documentRevisionStatusModel := &cloudantv1.DocumentRevisionStatus{
				Rev:    core.StringPtr("testString"),
				Status: core.StringPtr("available"),
			}

			documentModel := &cloudantv1.Document{
				Attachments:      map[string]cloudantv1.Attachment{"key1": *attachmentModel},
				Conflicts:        []string{"testString"},
				Deleted:          core.BoolPtr(true),
				DeletedConflicts: []string{"testString"},
				ID:               core.StringPtr("exampleid"),
				LocalSeq:         core.StringPtr("testString"),
				Rev:              core.StringPtr("testString"),
				Revisions:        revisionsModel,
				RevsInfo:         []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel},
			}
			documentModel.Attachments["foo"] = *attachmentModel
			documentModel.SetProperty("brand", core.StringPtr("Foo"))
			documentModel.SetProperty("colours", core.StringPtr("[\"red\",\"green\",\"black\",\"blue\"]"))
			documentModel.SetProperty("description", core.StringPtr("Slim Colourful Design Electronic Cooking Appliance for ..."))
			documentModel.SetProperty("image", core.StringPtr("assets/img/0gmsnghhew.jpg"))
			documentModel.SetProperty("keywords", core.StringPtr("[\"Foo\",\"Scales\",\"Weight\",\"Digital\",\"Kitchen\"]"))
			documentModel.SetProperty("name", core.StringPtr("Digital Kitchen Scales"))
			documentModel.SetProperty("price", core.StringPtr("14.99"))
			documentModel.SetProperty("productid", core.StringPtr("1000042"))
			documentModel.SetProperty("taxonomy", core.StringPtr("[\"Home\",\"Kitchen\",\"Small Appliances\"]"))
			documentModel.SetProperty("type", core.StringPtr("product"))

			postDocumentOptions := &cloudantv1.PostDocumentOptions{
				Db:          core.StringPtr("testString"),
				Document:    documentModel,
				ContentType: core.StringPtr("application/json"),
				Batch:       core.StringPtr("ok"),
			}

			documentResult, response, err := cloudantService.PostDocument(postDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`PostAllDocs - Query a list of all documents in a database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostAllDocs(postAllDocsOptions *PostAllDocsOptions)`, func() {
			postAllDocsOptions := &cloudantv1.PostAllDocsOptions{
				Db:              core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(false),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(10)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          core.StringPtr("testString"),
				Key:             core.StringPtr("testString"),
				Keys:            []string{"testString"},
				StartKey:        core.StringPtr("0007741142412418284"),
			}

			allDocsResult, response, err := cloudantService.PostAllDocs(postAllDocsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allDocsResult).ToNot(BeNil())
		})
	})

	Describe(`PostAllDocsAsStream - Query a list of all documents in a database as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostAllDocsAsStream(postAllDocsOptions *PostAllDocsOptions)`, func() {
			postAllDocsOptions := &cloudantv1.PostAllDocsOptions{
				Db:              core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(false),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(10)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          core.StringPtr("testString"),
				Key:             core.StringPtr("testString"),
				Keys:            []string{"testString"},
				StartKey:        core.StringPtr("0007741142412418284"),
			}

			result, response, err := cloudantService.PostAllDocsAsStream(postAllDocsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PostAllDocsQueries - Multi-query the list of all documents in a database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostAllDocsQueries(postAllDocsQueriesOptions *PostAllDocsQueriesOptions)`, func() {
			allDocsQueryModel := &cloudantv1.AllDocsQuery{
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(false),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(0)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          core.StringPtr("testString"),
				Key:             core.StringPtr("testString"),
				Keys:            []string{"small-appliances:1000042", "small-appliances:1000043"},
				StartKey:        core.StringPtr("testString"),
			}

			postAllDocsQueriesOptions := &cloudantv1.PostAllDocsQueriesOptions{
				Db:      core.StringPtr("testString"),
				Queries: []cloudantv1.AllDocsQuery{*allDocsQueryModel},
			}

			allDocsQueriesResult, response, err := cloudantService.PostAllDocsQueries(postAllDocsQueriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allDocsQueriesResult).ToNot(BeNil())
		})
	})

	Describe(`PostAllDocsQueriesAsStream - Multi-query the list of all documents in a database as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostAllDocsQueriesAsStream(postAllDocsQueriesOptions *PostAllDocsQueriesOptions)`, func() {
			allDocsQueryModel := &cloudantv1.AllDocsQuery{
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(false),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(0)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          core.StringPtr("testString"),
				Key:             core.StringPtr("testString"),
				Keys:            []string{"small-appliances:1000042", "small-appliances:1000043"},
				StartKey:        core.StringPtr("testString"),
			}

			postAllDocsQueriesOptions := &cloudantv1.PostAllDocsQueriesOptions{
				Db:      core.StringPtr("testString"),
				Queries: []cloudantv1.AllDocsQuery{*allDocsQueryModel},
			}

			result, response, err := cloudantService.PostAllDocsQueriesAsStream(postAllDocsQueriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PostBulkDocs - Bulk modify multiple documents in a database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostBulkDocs(postBulkDocsOptions *PostBulkDocsOptions)`, func() {
			attachmentModel := &cloudantv1.Attachment{
				ContentType:   core.StringPtr("testString"),
				Data:          CreateMockByteArray("VGhpcyBpcyBhIG1vY2sgYnl0ZSBhcnJheSB2YWx1ZS4="),
				Digest:        core.StringPtr("testString"),
				EncodedLength: core.Int64Ptr(int64(0)),
				Encoding:      core.StringPtr("testString"),
				Follows:       core.BoolPtr(true),
				Length:        core.Int64Ptr(int64(0)),
				Revpos:        core.Int64Ptr(int64(1)),
				Stub:          core.BoolPtr(true),
			}

			revisionsModel := &cloudantv1.Revisions{
				Ids:   []string{"testString"},
				Start: core.Int64Ptr(int64(1)),
			}

			documentRevisionStatusModel := &cloudantv1.DocumentRevisionStatus{
				Rev:    core.StringPtr("testString"),
				Status: core.StringPtr("available"),
			}

			documentModel := &cloudantv1.Document{
				Attachments:      map[string]cloudantv1.Attachment{"key1": *attachmentModel},
				Conflicts:        []string{"testString"},
				Deleted:          core.BoolPtr(true),
				DeletedConflicts: []string{"testString"},
				ID:               core.StringPtr("0007241142412418284"),
				LocalSeq:         core.StringPtr("testString"),
				Rev:              core.StringPtr("testString"),
				Revisions:        revisionsModel,
				RevsInfo:         []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel},
			}
			documentModel.Attachments["foo"] = *attachmentModel
			documentModel.SetProperty("date", core.StringPtr("2019-01-28T10:44:22.000Z"))
			documentModel.SetProperty("eventType", core.StringPtr("addedToBasket"))
			documentModel.SetProperty("productId", core.StringPtr("1000042"))
			documentModel.SetProperty("type", core.StringPtr("event"))
			documentModel.SetProperty("userid", core.StringPtr("abc123"))

			bulkDocsModel := &cloudantv1.BulkDocs{
				Docs:     []cloudantv1.Document{*documentModel},
				NewEdits: core.BoolPtr(true),
			}

			postBulkDocsOptions := &cloudantv1.PostBulkDocsOptions{
				Db:       core.StringPtr("testString"),
				BulkDocs: bulkDocsModel,
			}

			documentResult, response, err := cloudantService.PostBulkDocs(postBulkDocsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`PostBulkGet - Bulk query revision information for multiple documents`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostBulkGet(postBulkGetOptions *PostBulkGetOptions)`, func() {
			bulkGetQueryDocumentModel := &cloudantv1.BulkGetQueryDocument{
				AttsSince: []string{"1-99b02e08da151943c2dcb40090160bb8"},
				ID:        core.StringPtr("order00067"),
				Rev:       core.StringPtr("3-917fa2381192822767f010b95b45325b"),
			}

			postBulkGetOptions := &cloudantv1.PostBulkGetOptions{
				Db:              core.StringPtr("testString"),
				Docs:            []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel},
				Attachments:     core.BoolPtr(false),
				AttEncodingInfo: core.BoolPtr(false),
				Latest:          core.BoolPtr(false),
				Revs:            core.BoolPtr(false),
			}

			bulkGetResult, response, err := cloudantService.PostBulkGet(postBulkGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bulkGetResult).ToNot(BeNil())
		})
	})

	Describe(`PostBulkGetAsMixed - Bulk query revision information for multiple documents as mixed`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostBulkGetAsMixed(postBulkGetOptions *PostBulkGetOptions)`, func() {
			bulkGetQueryDocumentModel := &cloudantv1.BulkGetQueryDocument{
				AttsSince: []string{"1-99b02e08da151943c2dcb40090160bb8"},
				ID:        core.StringPtr("order00067"),
				Rev:       core.StringPtr("3-917fa2381192822767f010b95b45325b"),
			}

			postBulkGetOptions := &cloudantv1.PostBulkGetOptions{
				Db:              core.StringPtr("testString"),
				Docs:            []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel},
				Attachments:     core.BoolPtr(false),
				AttEncodingInfo: core.BoolPtr(false),
				Latest:          core.BoolPtr(false),
				Revs:            core.BoolPtr(false),
			}

			result, response, err := cloudantService.PostBulkGetAsMixed(postBulkGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PostBulkGetAsRelated - Bulk query revision information for multiple documents as related`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostBulkGetAsRelated(postBulkGetOptions *PostBulkGetOptions)`, func() {
			bulkGetQueryDocumentModel := &cloudantv1.BulkGetQueryDocument{
				AttsSince: []string{"1-99b02e08da151943c2dcb40090160bb8"},
				ID:        core.StringPtr("order00067"),
				Rev:       core.StringPtr("3-917fa2381192822767f010b95b45325b"),
			}

			postBulkGetOptions := &cloudantv1.PostBulkGetOptions{
				Db:              core.StringPtr("testString"),
				Docs:            []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel},
				Attachments:     core.BoolPtr(false),
				AttEncodingInfo: core.BoolPtr(false),
				Latest:          core.BoolPtr(false),
				Revs:            core.BoolPtr(false),
			}

			result, response, err := cloudantService.PostBulkGetAsRelated(postBulkGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PostBulkGetAsStream - Bulk query revision information for multiple documents as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostBulkGetAsStream(postBulkGetOptions *PostBulkGetOptions)`, func() {
			bulkGetQueryDocumentModel := &cloudantv1.BulkGetQueryDocument{
				AttsSince: []string{"1-99b02e08da151943c2dcb40090160bb8"},
				ID:        core.StringPtr("order00067"),
				Rev:       core.StringPtr("3-917fa2381192822767f010b95b45325b"),
			}

			postBulkGetOptions := &cloudantv1.PostBulkGetOptions{
				Db:              core.StringPtr("testString"),
				Docs:            []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel},
				Attachments:     core.BoolPtr(false),
				AttEncodingInfo: core.BoolPtr(false),
				Latest:          core.BoolPtr(false),
				Revs:            core.BoolPtr(false),
			}

			result, response, err := cloudantService.PostBulkGetAsStream(postBulkGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetDocument - Retrieve a document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDocument(getDocumentOptions *GetDocumentOptions)`, func() {
			getDocumentOptions := &cloudantv1.GetDocumentOptions{
				Db:               core.StringPtr("testString"),
				DocID:            core.StringPtr("testString"),
				IfNoneMatch:      core.StringPtr("testString"),
				Attachments:      core.BoolPtr(false),
				AttEncodingInfo:  core.BoolPtr(false),
				Conflicts:        core.BoolPtr(false),
				DeletedConflicts: core.BoolPtr(false),
				Latest:           core.BoolPtr(false),
				LocalSeq:         core.BoolPtr(false),
				Meta:             core.BoolPtr(false),
				Rev:              core.StringPtr("testString"),
				Revs:             core.BoolPtr(false),
				RevsInfo:         core.BoolPtr(false),
			}

			document, response, err := cloudantService.GetDocument(getDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(document).ToNot(BeNil())
		})
	})

	Describe(`GetDocumentAsMixed - Retrieve a document as mixed`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDocumentAsMixed(getDocumentOptions *GetDocumentOptions)`, func() {
			getDocumentOptions := &cloudantv1.GetDocumentOptions{
				Db:               core.StringPtr("testString"),
				DocID:            core.StringPtr("testString"),
				IfNoneMatch:      core.StringPtr("testString"),
				Attachments:      core.BoolPtr(false),
				AttEncodingInfo:  core.BoolPtr(false),
				Conflicts:        core.BoolPtr(false),
				DeletedConflicts: core.BoolPtr(false),
				Latest:           core.BoolPtr(false),
				LocalSeq:         core.BoolPtr(false),
				Meta:             core.BoolPtr(false),
				Rev:              core.StringPtr("testString"),
				Revs:             core.BoolPtr(false),
				RevsInfo:         core.BoolPtr(false),
			}

			result, response, err := cloudantService.GetDocumentAsMixed(getDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetDocumentAsRelated - Retrieve a document as related`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDocumentAsRelated(getDocumentOptions *GetDocumentOptions)`, func() {
			getDocumentOptions := &cloudantv1.GetDocumentOptions{
				Db:               core.StringPtr("testString"),
				DocID:            core.StringPtr("testString"),
				IfNoneMatch:      core.StringPtr("testString"),
				Attachments:      core.BoolPtr(false),
				AttEncodingInfo:  core.BoolPtr(false),
				Conflicts:        core.BoolPtr(false),
				DeletedConflicts: core.BoolPtr(false),
				Latest:           core.BoolPtr(false),
				LocalSeq:         core.BoolPtr(false),
				Meta:             core.BoolPtr(false),
				Rev:              core.StringPtr("testString"),
				Revs:             core.BoolPtr(false),
				RevsInfo:         core.BoolPtr(false),
			}

			result, response, err := cloudantService.GetDocumentAsRelated(getDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetDocumentAsStream - Retrieve a document as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDocumentAsStream(getDocumentOptions *GetDocumentOptions)`, func() {
			getDocumentOptions := &cloudantv1.GetDocumentOptions{
				Db:               core.StringPtr("testString"),
				DocID:            core.StringPtr("testString"),
				IfNoneMatch:      core.StringPtr("testString"),
				Attachments:      core.BoolPtr(false),
				AttEncodingInfo:  core.BoolPtr(false),
				Conflicts:        core.BoolPtr(false),
				DeletedConflicts: core.BoolPtr(false),
				Latest:           core.BoolPtr(false),
				LocalSeq:         core.BoolPtr(false),
				Meta:             core.BoolPtr(false),
				Rev:              core.StringPtr("testString"),
				Revs:             core.BoolPtr(false),
				RevsInfo:         core.BoolPtr(false),
			}

			result, response, err := cloudantService.GetDocumentAsStream(getDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PutDocument - Create or modify a document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutDocument(putDocumentOptions *PutDocumentOptions)`, func() {
			attachmentModel := &cloudantv1.Attachment{
				ContentType:   core.StringPtr("testString"),
				Data:          CreateMockByteArray("VGhpcyBpcyBhIG1vY2sgYnl0ZSBhcnJheSB2YWx1ZS4="),
				Digest:        core.StringPtr("testString"),
				EncodedLength: core.Int64Ptr(int64(0)),
				Encoding:      core.StringPtr("testString"),
				Follows:       core.BoolPtr(true),
				Length:        core.Int64Ptr(int64(0)),
				Revpos:        core.Int64Ptr(int64(1)),
				Stub:          core.BoolPtr(true),
			}

			revisionsModel := &cloudantv1.Revisions{
				Ids:   []string{"testString"},
				Start: core.Int64Ptr(int64(1)),
			}

			documentRevisionStatusModel := &cloudantv1.DocumentRevisionStatus{
				Rev:    core.StringPtr("testString"),
				Status: core.StringPtr("available"),
			}

			documentModel := &cloudantv1.Document{
				Attachments:      map[string]cloudantv1.Attachment{"key1": *attachmentModel},
				Conflicts:        []string{"testString"},
				Deleted:          core.BoolPtr(true),
				DeletedConflicts: []string{"testString"},
				ID:               core.StringPtr("exampleid"),
				LocalSeq:         core.StringPtr("testString"),
				Rev:              core.StringPtr("testString"),
				Revisions:        revisionsModel,
				RevsInfo:         []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel},
			}
			documentModel.Attachments["foo"] = *attachmentModel
			documentModel.SetProperty("brand", core.StringPtr("Foo"))
			documentModel.SetProperty("colours", core.StringPtr("[\"red\",\"green\",\"black\",\"blue\"]"))
			documentModel.SetProperty("description", core.StringPtr("Slim Colourful Design Electronic Cooking Appliance for ..."))
			documentModel.SetProperty("image", core.StringPtr("assets/img/0gmsnghhew.jpg"))
			documentModel.SetProperty("keywords", core.StringPtr("[\"Foo\",\"Scales\",\"Weight\",\"Digital\",\"Kitchen\"]"))
			documentModel.SetProperty("name", core.StringPtr("Digital Kitchen Scales"))
			documentModel.SetProperty("price", core.StringPtr("14.99"))
			documentModel.SetProperty("productid", core.StringPtr("1000042"))
			documentModel.SetProperty("taxonomy", core.StringPtr("[\"Home\",\"Kitchen\",\"Small Appliances\"]"))
			documentModel.SetProperty("type", core.StringPtr("product"))

			putDocumentOptions := &cloudantv1.PutDocumentOptions{
				Db:          core.StringPtr("testString"),
				DocID:       core.StringPtr("testString"),
				Document:    documentModel,
				ContentType: core.StringPtr("application/json"),
				IfMatch:     core.StringPtr("testString"),
				Batch:       core.StringPtr("ok"),
				NewEdits:    core.BoolPtr(false),
				Rev:         core.StringPtr("testString"),
			}

			documentResult, response, err := cloudantService.PutDocument(putDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`HeadDesignDocument - Retrieve the HTTP headers for a design document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HeadDesignDocument(headDesignDocumentOptions *HeadDesignDocumentOptions)`, func() {
			headDesignDocumentOptions := &cloudantv1.HeadDesignDocumentOptions{
				Db:          core.StringPtr("testString"),
				Ddoc:        core.StringPtr("testString"),
				IfNoneMatch: core.StringPtr("testString"),
			}

			response, err := cloudantService.HeadDesignDocument(headDesignDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetDesignDocument - Retrieve a design document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDesignDocument(getDesignDocumentOptions *GetDesignDocumentOptions)`, func() {
			getDesignDocumentOptions := &cloudantv1.GetDesignDocumentOptions{
				Db:               core.StringPtr("testString"),
				Ddoc:             core.StringPtr("testString"),
				IfNoneMatch:      core.StringPtr("testString"),
				Attachments:      core.BoolPtr(false),
				AttEncodingInfo:  core.BoolPtr(false),
				Conflicts:        core.BoolPtr(false),
				DeletedConflicts: core.BoolPtr(false),
				Latest:           core.BoolPtr(false),
				LocalSeq:         core.BoolPtr(false),
				Meta:             core.BoolPtr(false),
				Rev:              core.StringPtr("testString"),
				Revs:             core.BoolPtr(false),
				RevsInfo:         core.BoolPtr(false),
			}

			designDocument, response, err := cloudantService.GetDesignDocument(getDesignDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(designDocument).ToNot(BeNil())
		})
	})

	Describe(`PutDesignDocument - Create or modify a design document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutDesignDocument(putDesignDocumentOptions *PutDesignDocumentOptions)`, func() {
			attachmentModel := &cloudantv1.Attachment{
				ContentType:   core.StringPtr("testString"),
				Data:          CreateMockByteArray("VGhpcyBpcyBhIG1vY2sgYnl0ZSBhcnJheSB2YWx1ZS4="),
				Digest:        core.StringPtr("testString"),
				EncodedLength: core.Int64Ptr(int64(0)),
				Encoding:      core.StringPtr("testString"),
				Follows:       core.BoolPtr(true),
				Length:        core.Int64Ptr(int64(0)),
				Revpos:        core.Int64Ptr(int64(1)),
				Stub:          core.BoolPtr(true),
			}

			revisionsModel := &cloudantv1.Revisions{
				Ids:   []string{"testString"},
				Start: core.Int64Ptr(int64(1)),
			}

			documentRevisionStatusModel := &cloudantv1.DocumentRevisionStatus{
				Rev:    core.StringPtr("testString"),
				Status: core.StringPtr("available"),
			}

			analyzerModel := &cloudantv1.Analyzer{
				Name:      core.StringPtr("classic"),
				Stopwords: []string{"testString"},
			}

			analyzerConfigurationModel := &cloudantv1.AnalyzerConfiguration{
				Name:      core.StringPtr("standard"),
				Stopwords: []string{"testString"},
				Fields:    map[string]cloudantv1.Analyzer{"key1": *analyzerModel},
			}
			analyzerConfigurationModel.Fields["foo"] = *analyzerModel

			searchIndexDefinitionModel := &cloudantv1.SearchIndexDefinition{
				Analyzer: analyzerConfigurationModel,
				Index:    core.StringPtr("function (doc) {\n  index(\"price\", doc.price);\n}"),
			}

			designDocumentOptionsModel := &cloudantv1.DesignDocumentOptions{
				Partitioned: core.BoolPtr(true),
			}

			designDocumentViewsMapReduceModel := &cloudantv1.DesignDocumentViewsMapReduce{
				Map:    core.StringPtr("function(doc) { \n  emit(doc.productid, [doc.brand, doc.name, doc.description]) \n}"),
				Reduce: core.StringPtr("testString"),
			}

			designDocumentModel := &cloudantv1.DesignDocument{
				Attachments:       map[string]cloudantv1.Attachment{"key1": *attachmentModel},
				Conflicts:         []string{"testString"},
				Deleted:           core.BoolPtr(true),
				DeletedConflicts:  []string{"testString"},
				ID:                core.StringPtr("_design/appliances"),
				LocalSeq:          core.StringPtr("testString"),
				Rev:               core.StringPtr("8-7e2537e5989294471061e0cfd7292725"),
				Revisions:         revisionsModel,
				RevsInfo:          []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel},
				Autoupdate:        core.BoolPtr(true),
				Filters:           map[string]string{"key1": "testString"},
				Indexes:           map[string]cloudantv1.SearchIndexDefinition{"key1": *searchIndexDefinitionModel},
				Language:          core.StringPtr("javascript"),
				Options:           designDocumentOptionsModel,
				ValidateDocUpdate: core.StringPtr("testString"),
				Views:             map[string]cloudantv1.DesignDocumentViewsMapReduce{"key1": *designDocumentViewsMapReduceModel},
			}
			designDocumentModel.Attachments["foo"] = *attachmentModel
			designDocumentModel.Indexes["foo"] = *searchIndexDefinitionModel
			designDocumentModel.Views["foo"] = *designDocumentViewsMapReduceModel
			designDocumentModel.SetProperty("foo", "testString")

			putDesignDocumentOptions := &cloudantv1.PutDesignDocumentOptions{
				Db:             core.StringPtr("testString"),
				Ddoc:           core.StringPtr("testString"),
				DesignDocument: designDocumentModel,
				IfMatch:        core.StringPtr("testString"),
				Batch:          core.StringPtr("ok"),
				NewEdits:       core.BoolPtr(false),
				Rev:            core.StringPtr("testString"),
			}

			documentResult, response, err := cloudantService.PutDesignDocument(putDesignDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`GetDesignDocumentInformation - Retrieve information about a design document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDesignDocumentInformation(getDesignDocumentInformationOptions *GetDesignDocumentInformationOptions)`, func() {
			getDesignDocumentInformationOptions := &cloudantv1.GetDesignDocumentInformationOptions{
				Db:   core.StringPtr("testString"),
				Ddoc: core.StringPtr("testString"),
			}

			designDocumentInformation, response, err := cloudantService.GetDesignDocumentInformation(getDesignDocumentInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(designDocumentInformation).ToNot(BeNil())
		})
	})

	Describe(`PostDesignDocs - Query a list of all design documents in a database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostDesignDocs(postDesignDocsOptions *PostDesignDocsOptions)`, func() {
			postDesignDocsOptions := &cloudantv1.PostDesignDocsOptions{
				Db:              core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(false),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(10)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          core.StringPtr("testString"),
				Key:             core.StringPtr("testString"),
				Keys:            []string{"testString"},
				StartKey:        core.StringPtr("0007741142412418284"),
			}

			allDocsResult, response, err := cloudantService.PostDesignDocs(postDesignDocsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allDocsResult).ToNot(BeNil())
		})
	})

	Describe(`PostDesignDocsQueries - Multi-query the list of all design documents`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostDesignDocsQueries(postDesignDocsQueriesOptions *PostDesignDocsQueriesOptions)`, func() {
			allDocsQueryModel := &cloudantv1.AllDocsQuery{
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(false),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(0)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          core.StringPtr("testString"),
				Key:             core.StringPtr("testString"),
				Keys:            []string{"small-appliances:1000042", "small-appliances:1000043"},
				StartKey:        core.StringPtr("testString"),
			}

			postDesignDocsQueriesOptions := &cloudantv1.PostDesignDocsQueriesOptions{
				Db:      core.StringPtr("testString"),
				Queries: []cloudantv1.AllDocsQuery{*allDocsQueryModel},
				Accept:  core.StringPtr("application/json"),
			}

			allDocsQueriesResult, response, err := cloudantService.PostDesignDocsQueries(postDesignDocsQueriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allDocsQueriesResult).ToNot(BeNil())
		})
	})

	Describe(`PostView - Query a MapReduce view`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostView(postViewOptions *PostViewOptions)`, func() {
			postViewOptions := &cloudantv1.PostViewOptions{
				Db:              core.StringPtr("testString"),
				Ddoc:            core.StringPtr("testString"),
				View:            core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(true),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(10)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          "testString",
				EndKeyDocID:     core.StringPtr("testString"),
				Group:           core.BoolPtr(false),
				GroupLevel:      core.Int64Ptr(int64(1)),
				Key:             "testString",
				Keys:            []interface{}{"examplekey"},
				Reduce:          core.BoolPtr(true),
				Stable:          core.BoolPtr(false),
				StartKey:        "testString",
				StartKeyDocID:   core.StringPtr("testString"),
				Update:          core.StringPtr("true"),
			}

			viewResult, response, err := cloudantService.PostView(postViewOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(viewResult).ToNot(BeNil())
		})
	})

	Describe(`PostViewAsStream - Query a MapReduce view as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostViewAsStream(postViewOptions *PostViewOptions)`, func() {
			postViewOptions := &cloudantv1.PostViewOptions{
				Db:              core.StringPtr("testString"),
				Ddoc:            core.StringPtr("testString"),
				View:            core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(true),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(10)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          "testString",
				EndKeyDocID:     core.StringPtr("testString"),
				Group:           core.BoolPtr(false),
				GroupLevel:      core.Int64Ptr(int64(1)),
				Key:             "testString",
				Keys:            []interface{}{"examplekey"},
				Reduce:          core.BoolPtr(true),
				Stable:          core.BoolPtr(false),
				StartKey:        "testString",
				StartKeyDocID:   core.StringPtr("testString"),
				Update:          core.StringPtr("true"),
			}

			result, response, err := cloudantService.PostViewAsStream(postViewOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PostViewQueries - Multi-query a MapReduce view`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostViewQueries(postViewQueriesOptions *PostViewQueriesOptions)`, func() {
			viewQueryModel := &cloudantv1.ViewQuery{
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(true),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(5)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          "testString",
				EndKeyDocID:     core.StringPtr("testString"),
				Group:           core.BoolPtr(false),
				GroupLevel:      core.Int64Ptr(int64(1)),
				Key:             "testString",
				Keys:            []interface{}{"testString"},
				Reduce:          core.BoolPtr(true),
				Stable:          core.BoolPtr(false),
				StartKey:        "testString",
				StartKeyDocID:   core.StringPtr("testString"),
				Update:          core.StringPtr("true"),
			}

			postViewQueriesOptions := &cloudantv1.PostViewQueriesOptions{
				Db:      core.StringPtr("testString"),
				Ddoc:    core.StringPtr("testString"),
				View:    core.StringPtr("testString"),
				Queries: []cloudantv1.ViewQuery{*viewQueryModel},
			}

			viewQueriesResult, response, err := cloudantService.PostViewQueries(postViewQueriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(viewQueriesResult).ToNot(BeNil())
		})
	})

	Describe(`PostViewQueriesAsStream - Multi-query a MapReduce view as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostViewQueriesAsStream(postViewQueriesOptions *PostViewQueriesOptions)`, func() {
			viewQueryModel := &cloudantv1.ViewQuery{
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(true),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(5)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          "testString",
				EndKeyDocID:     core.StringPtr("testString"),
				Group:           core.BoolPtr(false),
				GroupLevel:      core.Int64Ptr(int64(1)),
				Key:             "testString",
				Keys:            []interface{}{"testString"},
				Reduce:          core.BoolPtr(true),
				Stable:          core.BoolPtr(false),
				StartKey:        "testString",
				StartKeyDocID:   core.StringPtr("testString"),
				Update:          core.StringPtr("true"),
			}

			postViewQueriesOptions := &cloudantv1.PostViewQueriesOptions{
				Db:      core.StringPtr("testString"),
				Ddoc:    core.StringPtr("testString"),
				View:    core.StringPtr("testString"),
				Queries: []cloudantv1.ViewQuery{*viewQueryModel},
			}

			result, response, err := cloudantService.PostViewQueriesAsStream(postViewQueriesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetPartitionInformation - Retrieve information about a database partition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPartitionInformation(getPartitionInformationOptions *GetPartitionInformationOptions)`, func() {
			getPartitionInformationOptions := &cloudantv1.GetPartitionInformationOptions{
				Db:           core.StringPtr("testString"),
				PartitionKey: core.StringPtr("testString"),
			}

			partitionInformation, response, err := cloudantService.GetPartitionInformation(getPartitionInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(partitionInformation).ToNot(BeNil())
		})
	})

	Describe(`PostPartitionAllDocs - Query a list of all documents in a database partition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostPartitionAllDocs(postPartitionAllDocsOptions *PostPartitionAllDocsOptions)`, func() {
			postPartitionAllDocsOptions := &cloudantv1.PostPartitionAllDocsOptions{
				Db:              core.StringPtr("testString"),
				PartitionKey:    core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(false),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(10)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          core.StringPtr("testString"),
				Key:             core.StringPtr("testString"),
				Keys:            []string{"testString"},
				StartKey:        core.StringPtr("0007741142412418284"),
			}

			allDocsResult, response, err := cloudantService.PostPartitionAllDocs(postPartitionAllDocsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allDocsResult).ToNot(BeNil())
		})
	})

	Describe(`PostPartitionAllDocsAsStream - Query a list of all documents in a database partition as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostPartitionAllDocsAsStream(postPartitionAllDocsOptions *PostPartitionAllDocsOptions)`, func() {
			postPartitionAllDocsOptions := &cloudantv1.PostPartitionAllDocsOptions{
				Db:              core.StringPtr("testString"),
				PartitionKey:    core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(false),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(10)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          core.StringPtr("testString"),
				Key:             core.StringPtr("testString"),
				Keys:            []string{"testString"},
				StartKey:        core.StringPtr("0007741142412418284"),
			}

			result, response, err := cloudantService.PostPartitionAllDocsAsStream(postPartitionAllDocsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PostPartitionSearch - Query a database partition search index`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostPartitionSearch(postPartitionSearchOptions *PostPartitionSearchOptions)`, func() {
			postPartitionSearchOptions := &cloudantv1.PostPartitionSearchOptions{
				Db:               core.StringPtr("testString"),
				PartitionKey:     core.StringPtr("testString"),
				Ddoc:             core.StringPtr("testString"),
				Index:            core.StringPtr("testString"),
				Query:            core.StringPtr("name:Jane* AND active:True"),
				Bookmark:         core.StringPtr("testString"),
				HighlightFields:  []string{"testString"},
				HighlightNumber:  core.Int64Ptr(int64(1)),
				HighlightPostTag: core.StringPtr("</em>"),
				HighlightPreTag:  core.StringPtr("<em>"),
				HighlightSize:    core.Int64Ptr(int64(100)),
				IncludeDocs:      core.BoolPtr(false),
				IncludeFields:    []string{"testString"},
				Limit:            core.Int64Ptr(int64(3)),
				Sort:             []string{"testString"},
				Stale:            core.StringPtr("ok"),
			}

			searchResult, response, err := cloudantService.PostPartitionSearch(postPartitionSearchOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(searchResult).ToNot(BeNil())
		})
	})

	Describe(`PostPartitionSearchAsStream - Query a database partition search index as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostPartitionSearchAsStream(postPartitionSearchOptions *PostPartitionSearchOptions)`, func() {
			postPartitionSearchOptions := &cloudantv1.PostPartitionSearchOptions{
				Db:               core.StringPtr("testString"),
				PartitionKey:     core.StringPtr("testString"),
				Ddoc:             core.StringPtr("testString"),
				Index:            core.StringPtr("testString"),
				Query:            core.StringPtr("name:Jane* AND active:True"),
				Bookmark:         core.StringPtr("testString"),
				HighlightFields:  []string{"testString"},
				HighlightNumber:  core.Int64Ptr(int64(1)),
				HighlightPostTag: core.StringPtr("</em>"),
				HighlightPreTag:  core.StringPtr("<em>"),
				HighlightSize:    core.Int64Ptr(int64(100)),
				IncludeDocs:      core.BoolPtr(false),
				IncludeFields:    []string{"testString"},
				Limit:            core.Int64Ptr(int64(3)),
				Sort:             []string{"testString"},
				Stale:            core.StringPtr("ok"),
			}

			result, response, err := cloudantService.PostPartitionSearchAsStream(postPartitionSearchOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PostPartitionView - Query a database partition MapReduce view function`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostPartitionView(postPartitionViewOptions *PostPartitionViewOptions)`, func() {
			postPartitionViewOptions := &cloudantv1.PostPartitionViewOptions{
				Db:              core.StringPtr("testString"),
				PartitionKey:    core.StringPtr("testString"),
				Ddoc:            core.StringPtr("testString"),
				View:            core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(true),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(10)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          "testString",
				EndKeyDocID:     core.StringPtr("testString"),
				Group:           core.BoolPtr(false),
				GroupLevel:      core.Int64Ptr(int64(1)),
				Key:             "testString",
				Keys:            []interface{}{"examplekey"},
				Reduce:          core.BoolPtr(true),
				StartKey:        "testString",
				StartKeyDocID:   core.StringPtr("testString"),
				Update:          core.StringPtr("true"),
			}

			viewResult, response, err := cloudantService.PostPartitionView(postPartitionViewOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(viewResult).ToNot(BeNil())
		})
	})

	Describe(`PostPartitionViewAsStream - Query a database partition MapReduce view function as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostPartitionViewAsStream(postPartitionViewOptions *PostPartitionViewOptions)`, func() {
			postPartitionViewOptions := &cloudantv1.PostPartitionViewOptions{
				Db:              core.StringPtr("testString"),
				PartitionKey:    core.StringPtr("testString"),
				Ddoc:            core.StringPtr("testString"),
				View:            core.StringPtr("testString"),
				AttEncodingInfo: core.BoolPtr(false),
				Attachments:     core.BoolPtr(false),
				Conflicts:       core.BoolPtr(false),
				Descending:      core.BoolPtr(false),
				IncludeDocs:     core.BoolPtr(true),
				InclusiveEnd:    core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(10)),
				Skip:            core.Int64Ptr(int64(0)),
				UpdateSeq:       core.BoolPtr(false),
				EndKey:          "testString",
				EndKeyDocID:     core.StringPtr("testString"),
				Group:           core.BoolPtr(false),
				GroupLevel:      core.Int64Ptr(int64(1)),
				Key:             "testString",
				Keys:            []interface{}{"examplekey"},
				Reduce:          core.BoolPtr(true),
				StartKey:        "testString",
				StartKeyDocID:   core.StringPtr("testString"),
				Update:          core.StringPtr("true"),
			}

			result, response, err := cloudantService.PostPartitionViewAsStream(postPartitionViewOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PostPartitionExplain - Retrieve information about which partition index is used for a query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostPartitionExplain(postPartitionExplainOptions *PostPartitionExplainOptions)`, func() {
			postPartitionExplainOptions := &cloudantv1.PostPartitionExplainOptions{
				Db:             core.StringPtr("testString"),
				PartitionKey:   core.StringPtr("testString"),
				Selector:       map[string]interface{}{"anyKey": "anyValue"},
				AllowFallback:  core.BoolPtr(true),
				Bookmark:       core.StringPtr("testString"),
				Conflicts:      core.BoolPtr(true),
				ExecutionStats: core.BoolPtr(true),
				Fields:         []string{"productid", "name", "description"},
				Limit:          core.Int64Ptr(int64(25)),
				Skip:           core.Int64Ptr(int64(0)),
				Sort:           []map[string]string{map[string]string{"key1": "asc"}},
				Stable:         core.BoolPtr(true),
				Update:         core.StringPtr("true"),
				UseIndex:       []string{"testString"},
			}

			explainResult, response, err := cloudantService.PostPartitionExplain(postPartitionExplainOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(explainResult).ToNot(BeNil())
		})
	})

	Describe(`PostPartitionFind - Query a database partition index by using selector syntax`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostPartitionFind(postPartitionFindOptions *PostPartitionFindOptions)`, func() {
			postPartitionFindOptions := &cloudantv1.PostPartitionFindOptions{
				Db:             core.StringPtr("testString"),
				PartitionKey:   core.StringPtr("testString"),
				Selector:       map[string]interface{}{"anyKey": "anyValue"},
				AllowFallback:  core.BoolPtr(true),
				Bookmark:       core.StringPtr("testString"),
				Conflicts:      core.BoolPtr(true),
				ExecutionStats: core.BoolPtr(true),
				Fields:         []string{"productid", "name", "description"},
				Limit:          core.Int64Ptr(int64(25)),
				Skip:           core.Int64Ptr(int64(0)),
				Sort:           []map[string]string{map[string]string{"key1": "asc"}},
				Stable:         core.BoolPtr(true),
				Update:         core.StringPtr("true"),
				UseIndex:       []string{"testString"},
			}

			findResult, response, err := cloudantService.PostPartitionFind(postPartitionFindOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(findResult).ToNot(BeNil())
		})
	})

	Describe(`PostPartitionFindAsStream - Query a database partition index by using selector syntax as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostPartitionFindAsStream(postPartitionFindOptions *PostPartitionFindOptions)`, func() {
			postPartitionFindOptions := &cloudantv1.PostPartitionFindOptions{
				Db:             core.StringPtr("testString"),
				PartitionKey:   core.StringPtr("testString"),
				Selector:       map[string]interface{}{"anyKey": "anyValue"},
				AllowFallback:  core.BoolPtr(true),
				Bookmark:       core.StringPtr("testString"),
				Conflicts:      core.BoolPtr(true),
				ExecutionStats: core.BoolPtr(true),
				Fields:         []string{"productid", "name", "description"},
				Limit:          core.Int64Ptr(int64(25)),
				Skip:           core.Int64Ptr(int64(0)),
				Sort:           []map[string]string{map[string]string{"key1": "asc"}},
				Stable:         core.BoolPtr(true),
				Update:         core.StringPtr("true"),
				UseIndex:       []string{"testString"},
			}

			result, response, err := cloudantService.PostPartitionFindAsStream(postPartitionFindOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PostExplain - Retrieve information about which index is used for a query`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostExplain(postExplainOptions *PostExplainOptions)`, func() {
			postExplainOptions := &cloudantv1.PostExplainOptions{
				Db:             core.StringPtr("testString"),
				Selector:       map[string]interface{}{"anyKey": "anyValue"},
				AllowFallback:  core.BoolPtr(true),
				Bookmark:       core.StringPtr("testString"),
				Conflicts:      core.BoolPtr(true),
				ExecutionStats: core.BoolPtr(true),
				Fields:         []string{"_id", "type", "name", "email"},
				Limit:          core.Int64Ptr(int64(3)),
				Skip:           core.Int64Ptr(int64(0)),
				Sort:           []map[string]string{map[string]string{"key1": "asc"}},
				Stable:         core.BoolPtr(true),
				Update:         core.StringPtr("true"),
				UseIndex:       []string{"testString"},
				R:              core.Int64Ptr(int64(1)),
			}

			explainResult, response, err := cloudantService.PostExplain(postExplainOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(explainResult).ToNot(BeNil())
		})
	})

	Describe(`PostFind - Query an index by using selector syntax`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostFind(postFindOptions *PostFindOptions)`, func() {
			postFindOptions := &cloudantv1.PostFindOptions{
				Db:             core.StringPtr("testString"),
				Selector:       map[string]interface{}{"anyKey": "anyValue"},
				AllowFallback:  core.BoolPtr(true),
				Bookmark:       core.StringPtr("testString"),
				Conflicts:      core.BoolPtr(true),
				ExecutionStats: core.BoolPtr(true),
				Fields:         []string{"_id", "type", "name", "email"},
				Limit:          core.Int64Ptr(int64(3)),
				Skip:           core.Int64Ptr(int64(0)),
				Sort:           []map[string]string{map[string]string{"key1": "asc"}},
				Stable:         core.BoolPtr(true),
				Update:         core.StringPtr("true"),
				UseIndex:       []string{"testString"},
				R:              core.Int64Ptr(int64(1)),
			}

			findResult, response, err := cloudantService.PostFind(postFindOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(findResult).ToNot(BeNil())
		})
	})

	Describe(`PostFindAsStream - Query an index by using selector syntax as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostFindAsStream(postFindOptions *PostFindOptions)`, func() {
			postFindOptions := &cloudantv1.PostFindOptions{
				Db:             core.StringPtr("testString"),
				Selector:       map[string]interface{}{"anyKey": "anyValue"},
				AllowFallback:  core.BoolPtr(true),
				Bookmark:       core.StringPtr("testString"),
				Conflicts:      core.BoolPtr(true),
				ExecutionStats: core.BoolPtr(true),
				Fields:         []string{"_id", "type", "name", "email"},
				Limit:          core.Int64Ptr(int64(3)),
				Skip:           core.Int64Ptr(int64(0)),
				Sort:           []map[string]string{map[string]string{"key1": "asc"}},
				Stable:         core.BoolPtr(true),
				Update:         core.StringPtr("true"),
				UseIndex:       []string{"testString"},
				R:              core.Int64Ptr(int64(1)),
			}

			result, response, err := cloudantService.PostFindAsStream(postFindOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetIndexesInformation - Retrieve information about all indexes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetIndexesInformation(getIndexesInformationOptions *GetIndexesInformationOptions)`, func() {
			getIndexesInformationOptions := &cloudantv1.GetIndexesInformationOptions{
				Db: core.StringPtr("testString"),
			}

			indexesInformation, response, err := cloudantService.GetIndexesInformation(getIndexesInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(indexesInformation).ToNot(BeNil())
		})
	})

	Describe(`PostIndex - Create a new index on a database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostIndex(postIndexOptions *PostIndexOptions)`, func() {
			analyzerModel := &cloudantv1.Analyzer{
				Name:      core.StringPtr("classic"),
				Stopwords: []string{"testString"},
			}

			indexTextOperatorDefaultFieldModel := &cloudantv1.IndexTextOperatorDefaultField{
				Analyzer: analyzerModel,
				Enabled:  core.BoolPtr(true),
			}

			indexFieldModel := &cloudantv1.IndexField{
				Name: core.StringPtr("asc"),
				Type: core.StringPtr("boolean"),
			}
			indexFieldModel.SetProperty("foo", core.StringPtr("asc"))

			indexDefinitionModel := &cloudantv1.IndexDefinition{
				DefaultAnalyzer:       analyzerModel,
				DefaultField:          indexTextOperatorDefaultFieldModel,
				Fields:                []cloudantv1.IndexField{*indexFieldModel},
				IndexArrayLengths:     core.BoolPtr(true),
				PartialFilterSelector: map[string]interface{}{"anyKey": "anyValue"},
			}

			postIndexOptions := &cloudantv1.PostIndexOptions{
				Db:          core.StringPtr("testString"),
				Index:       indexDefinitionModel,
				Ddoc:        core.StringPtr("json-index"),
				Name:        core.StringPtr("getUserByName"),
				Partitioned: core.BoolPtr(true),
				Type:        core.StringPtr("json"),
			}

			indexResult, response, err := cloudantService.PostIndex(postIndexOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(indexResult).ToNot(BeNil())
		})
	})

	Describe(`PostSearchAnalyze - Query tokenization of sample text`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostSearchAnalyze(postSearchAnalyzeOptions *PostSearchAnalyzeOptions)`, func() {
			postSearchAnalyzeOptions := &cloudantv1.PostSearchAnalyzeOptions{
				Analyzer: core.StringPtr("english"),
				Text:     core.StringPtr("running is fun"),
			}

			searchAnalyzeResult, response, err := cloudantService.PostSearchAnalyze(postSearchAnalyzeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(searchAnalyzeResult).ToNot(BeNil())
		})
	})

	Describe(`PostSearch - Query a search index`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostSearch(postSearchOptions *PostSearchOptions)`, func() {
			postSearchOptions := &cloudantv1.PostSearchOptions{
				Db:               core.StringPtr("testString"),
				Ddoc:             core.StringPtr("testString"),
				Index:            core.StringPtr("testString"),
				Query:            core.StringPtr("name:Jane* AND active:True"),
				Bookmark:         core.StringPtr("testString"),
				HighlightFields:  []string{"testString"},
				HighlightNumber:  core.Int64Ptr(int64(1)),
				HighlightPostTag: core.StringPtr("</em>"),
				HighlightPreTag:  core.StringPtr("<em>"),
				HighlightSize:    core.Int64Ptr(int64(100)),
				IncludeDocs:      core.BoolPtr(false),
				IncludeFields:    []string{"testString"},
				Limit:            core.Int64Ptr(int64(3)),
				Sort:             []string{"testString"},
				Stale:            core.StringPtr("ok"),
				Counts:           []string{"testString"},
				Drilldown:        [][]string{[]string{"testString"}},
				GroupField:       core.StringPtr("testString"),
				GroupLimit:       core.Int64Ptr(int64(1)),
				GroupSort:        []string{"testString"},
				Ranges:           map[string]map[string]map[string]string{"key1": map[string]map[string]string{"key1": map[string]string{"key1": "testString"}}},
			}

			searchResult, response, err := cloudantService.PostSearch(postSearchOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(searchResult).ToNot(BeNil())
		})
	})

	Describe(`PostSearchAsStream - Query a search index as stream`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostSearchAsStream(postSearchOptions *PostSearchOptions)`, func() {
			postSearchOptions := &cloudantv1.PostSearchOptions{
				Db:               core.StringPtr("testString"),
				Ddoc:             core.StringPtr("testString"),
				Index:            core.StringPtr("testString"),
				Query:            core.StringPtr("name:Jane* AND active:True"),
				Bookmark:         core.StringPtr("testString"),
				HighlightFields:  []string{"testString"},
				HighlightNumber:  core.Int64Ptr(int64(1)),
				HighlightPostTag: core.StringPtr("</em>"),
				HighlightPreTag:  core.StringPtr("<em>"),
				HighlightSize:    core.Int64Ptr(int64(100)),
				IncludeDocs:      core.BoolPtr(false),
				IncludeFields:    []string{"testString"},
				Limit:            core.Int64Ptr(int64(3)),
				Sort:             []string{"testString"},
				Stale:            core.StringPtr("ok"),
				Counts:           []string{"testString"},
				Drilldown:        [][]string{[]string{"testString"}},
				GroupField:       core.StringPtr("testString"),
				GroupLimit:       core.Int64Ptr(int64(1)),
				GroupSort:        []string{"testString"},
				Ranges:           map[string]map[string]map[string]string{"key1": map[string]map[string]string{"key1": map[string]string{"key1": "testString"}}},
			}

			result, response, err := cloudantService.PostSearchAsStream(postSearchOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetSearchInfo - Retrieve information about a search index`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSearchInfo(getSearchInfoOptions *GetSearchInfoOptions)`, func() {
			getSearchInfoOptions := &cloudantv1.GetSearchInfoOptions{
				Db:    core.StringPtr("testString"),
				Ddoc:  core.StringPtr("testString"),
				Index: core.StringPtr("testString"),
			}

			searchInfoResult, response, err := cloudantService.GetSearchInfo(getSearchInfoOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(searchInfoResult).ToNot(BeNil())
		})
	})

	Describe(`HeadReplicationDocument - Retrieve the HTTP headers for a persistent replication`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HeadReplicationDocument(headReplicationDocumentOptions *HeadReplicationDocumentOptions)`, func() {
			headReplicationDocumentOptions := &cloudantv1.HeadReplicationDocumentOptions{
				DocID:       core.StringPtr("testString"),
				IfNoneMatch: core.StringPtr("testString"),
			}

			response, err := cloudantService.HeadReplicationDocument(headReplicationDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`HeadSchedulerDocument - Retrieve HTTP headers for a replication scheduler document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HeadSchedulerDocument(headSchedulerDocumentOptions *HeadSchedulerDocumentOptions)`, func() {
			headSchedulerDocumentOptions := &cloudantv1.HeadSchedulerDocumentOptions{
				DocID: core.StringPtr("testString"),
			}

			response, err := cloudantService.HeadSchedulerDocument(headSchedulerDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`HeadSchedulerJob - Retrieve the HTTP headers for a replication scheduler job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HeadSchedulerJob(headSchedulerJobOptions *HeadSchedulerJobOptions)`, func() {
			headSchedulerJobOptions := &cloudantv1.HeadSchedulerJobOptions{
				JobID: core.StringPtr("testString"),
			}

			response, err := cloudantService.HeadSchedulerJob(headSchedulerJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`PostReplicator - Create a persistent replication with a generated ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostReplicator(postReplicatorOptions *PostReplicatorOptions)`, func() {
			attachmentModel := &cloudantv1.Attachment{
				ContentType:   core.StringPtr("testString"),
				Data:          CreateMockByteArray("VGhpcyBpcyBhIG1vY2sgYnl0ZSBhcnJheSB2YWx1ZS4="),
				Digest:        core.StringPtr("testString"),
				EncodedLength: core.Int64Ptr(int64(0)),
				Encoding:      core.StringPtr("testString"),
				Follows:       core.BoolPtr(true),
				Length:        core.Int64Ptr(int64(0)),
				Revpos:        core.Int64Ptr(int64(1)),
				Stub:          core.BoolPtr(true),
			}

			revisionsModel := &cloudantv1.Revisions{
				Ids:   []string{"testString"},
				Start: core.Int64Ptr(int64(1)),
			}

			documentRevisionStatusModel := &cloudantv1.DocumentRevisionStatus{
				Rev:    core.StringPtr("testString"),
				Status: core.StringPtr("available"),
			}

			replicationCreateTargetParametersModel := &cloudantv1.ReplicationCreateTargetParameters{
				N:           core.Int64Ptr(int64(3)),
				Partitioned: core.BoolPtr(false),
				Q:           core.Int64Ptr(int64(1)),
			}

			replicationDatabaseAuthBasicModel := &cloudantv1.ReplicationDatabaseAuthBasic{
				Password: core.StringPtr("testString"),
				Username: core.StringPtr("testString"),
			}

			replicationDatabaseAuthIamModel := &cloudantv1.ReplicationDatabaseAuthIam{
				ApiKey: core.StringPtr("testString"),
			}

			replicationDatabaseAuthModel := &cloudantv1.ReplicationDatabaseAuth{
				Basic: replicationDatabaseAuthBasicModel,
				Iam:   replicationDatabaseAuthIamModel,
			}

			replicationDatabaseModel := &cloudantv1.ReplicationDatabase{
				Auth:       replicationDatabaseAuthModel,
				HeadersVar: map[string]string{"key1": "testString"},
				URL:        core.StringPtr("https://my-source-instance.cloudantnosqldb.appdomain.cloud.example/animaldb"),
			}

			userContextModel := &cloudantv1.UserContext{
				Db:    core.StringPtr("testString"),
				Name:  core.StringPtr("john"),
				Roles: []string{"_replicator"},
			}

			replicationDocumentModel := &cloudantv1.ReplicationDocument{
				Attachments:        map[string]cloudantv1.Attachment{"key1": *attachmentModel},
				Conflicts:          []string{"testString"},
				Deleted:            core.BoolPtr(true),
				DeletedConflicts:   []string{"testString"},
				ID:                 core.StringPtr("testString"),
				LocalSeq:           core.StringPtr("testString"),
				Rev:                core.StringPtr("testString"),
				Revisions:          revisionsModel,
				RevsInfo:           []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel},
				Cancel:             core.BoolPtr(false),
				CheckpointInterval: core.Int64Ptr(int64(4500)),
				ConnectionTimeout:  core.Int64Ptr(int64(15000)),
				Continuous:         core.BoolPtr(true),
				CreateTarget:       core.BoolPtr(true),
				CreateTargetParams: replicationCreateTargetParametersModel,
				DocIds:             []string{"badger", "lemur", "llama"},
				Filter:             core.StringPtr("ddoc/my_filter"),
				HTTPConnections:    core.Int64Ptr(int64(10)),
				Owner:              core.StringPtr("testString"),
				QueryParams:        map[string]string{"key1": "testString"},
				RetriesPerRequest:  core.Int64Ptr(int64(3)),
				Selector:           map[string]interface{}{"anyKey": "anyValue"},
				SinceSeq:           core.StringPtr("34-g1AAAAGjeJzLYWBgYMlgTmGQT0lKzi9KdU"),
				SocketOptions:      core.StringPtr("[{keepalive, true}, {nodelay, false}]"),
				Source:             replicationDatabaseModel,
				SourceProxy:        core.StringPtr("testString"),
				Target:             replicationDatabaseModel,
				TargetProxy:        core.StringPtr("testString"),
				UseBulkGet:         core.BoolPtr(true),
				UseCheckpoints:     core.BoolPtr(false),
				UserCtx:            userContextModel,
				WinningRevsOnly:    core.BoolPtr(false),
				WorkerBatchSize:    core.Int64Ptr(int64(400)),
				WorkerProcesses:    core.Int64Ptr(int64(3)),
			}
			replicationDocumentModel.Attachments["foo"] = *attachmentModel
			replicationDocumentModel.SetProperty("foo", "testString")

			postReplicatorOptions := &cloudantv1.PostReplicatorOptions{
				ReplicationDocument: replicationDocumentModel,
				Batch:               core.StringPtr("ok"),
			}

			documentResult, response, err := cloudantService.PostReplicator(postReplicatorOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`GetReplicationDocument - Retrieve the configuration for a persistent replication`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReplicationDocument(getReplicationDocumentOptions *GetReplicationDocumentOptions)`, func() {
			getReplicationDocumentOptions := &cloudantv1.GetReplicationDocumentOptions{
				DocID:            core.StringPtr("testString"),
				IfNoneMatch:      core.StringPtr("testString"),
				Attachments:      core.BoolPtr(false),
				AttEncodingInfo:  core.BoolPtr(false),
				Conflicts:        core.BoolPtr(false),
				DeletedConflicts: core.BoolPtr(false),
				Latest:           core.BoolPtr(false),
				LocalSeq:         core.BoolPtr(false),
				Meta:             core.BoolPtr(false),
				Rev:              core.StringPtr("testString"),
				Revs:             core.BoolPtr(false),
				RevsInfo:         core.BoolPtr(false),
			}

			replicationDocument, response, err := cloudantService.GetReplicationDocument(getReplicationDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(replicationDocument).ToNot(BeNil())
		})
	})

	Describe(`PutReplicationDocument - Create or modify a persistent replication`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutReplicationDocument(putReplicationDocumentOptions *PutReplicationDocumentOptions)`, func() {
			attachmentModel := &cloudantv1.Attachment{
				ContentType:   core.StringPtr("testString"),
				Data:          CreateMockByteArray("VGhpcyBpcyBhIG1vY2sgYnl0ZSBhcnJheSB2YWx1ZS4="),
				Digest:        core.StringPtr("testString"),
				EncodedLength: core.Int64Ptr(int64(0)),
				Encoding:      core.StringPtr("testString"),
				Follows:       core.BoolPtr(true),
				Length:        core.Int64Ptr(int64(0)),
				Revpos:        core.Int64Ptr(int64(1)),
				Stub:          core.BoolPtr(true),
			}

			revisionsModel := &cloudantv1.Revisions{
				Ids:   []string{"testString"},
				Start: core.Int64Ptr(int64(1)),
			}

			documentRevisionStatusModel := &cloudantv1.DocumentRevisionStatus{
				Rev:    core.StringPtr("testString"),
				Status: core.StringPtr("available"),
			}

			replicationCreateTargetParametersModel := &cloudantv1.ReplicationCreateTargetParameters{
				N:           core.Int64Ptr(int64(3)),
				Partitioned: core.BoolPtr(false),
				Q:           core.Int64Ptr(int64(1)),
			}

			replicationDatabaseAuthBasicModel := &cloudantv1.ReplicationDatabaseAuthBasic{
				Password: core.StringPtr("testString"),
				Username: core.StringPtr("testString"),
			}

			replicationDatabaseAuthIamModel := &cloudantv1.ReplicationDatabaseAuthIam{
				ApiKey: core.StringPtr("testString"),
			}

			replicationDatabaseAuthModel := &cloudantv1.ReplicationDatabaseAuth{
				Basic: replicationDatabaseAuthBasicModel,
				Iam:   replicationDatabaseAuthIamModel,
			}

			replicationDatabaseModel := &cloudantv1.ReplicationDatabase{
				Auth:       replicationDatabaseAuthModel,
				HeadersVar: map[string]string{"key1": "testString"},
				URL:        core.StringPtr("https://my-source-instance.cloudantnosqldb.appdomain.cloud.example/animaldb"),
			}

			userContextModel := &cloudantv1.UserContext{
				Db:    core.StringPtr("testString"),
				Name:  core.StringPtr("john"),
				Roles: []string{"_replicator"},
			}

			replicationDocumentModel := &cloudantv1.ReplicationDocument{
				Attachments:        map[string]cloudantv1.Attachment{"key1": *attachmentModel},
				Conflicts:          []string{"testString"},
				Deleted:            core.BoolPtr(true),
				DeletedConflicts:   []string{"testString"},
				ID:                 core.StringPtr("testString"),
				LocalSeq:           core.StringPtr("testString"),
				Rev:                core.StringPtr("testString"),
				Revisions:          revisionsModel,
				RevsInfo:           []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel},
				Cancel:             core.BoolPtr(false),
				CheckpointInterval: core.Int64Ptr(int64(4500)),
				ConnectionTimeout:  core.Int64Ptr(int64(15000)),
				Continuous:         core.BoolPtr(true),
				CreateTarget:       core.BoolPtr(true),
				CreateTargetParams: replicationCreateTargetParametersModel,
				DocIds:             []string{"badger", "lemur", "llama"},
				Filter:             core.StringPtr("ddoc/my_filter"),
				HTTPConnections:    core.Int64Ptr(int64(10)),
				Owner:              core.StringPtr("testString"),
				QueryParams:        map[string]string{"key1": "testString"},
				RetriesPerRequest:  core.Int64Ptr(int64(3)),
				Selector:           map[string]interface{}{"anyKey": "anyValue"},
				SinceSeq:           core.StringPtr("34-g1AAAAGjeJzLYWBgYMlgTmGQT0lKzi9KdU"),
				SocketOptions:      core.StringPtr("[{keepalive, true}, {nodelay, false}]"),
				Source:             replicationDatabaseModel,
				SourceProxy:        core.StringPtr("testString"),
				Target:             replicationDatabaseModel,
				TargetProxy:        core.StringPtr("testString"),
				UseBulkGet:         core.BoolPtr(true),
				UseCheckpoints:     core.BoolPtr(false),
				UserCtx:            userContextModel,
				WinningRevsOnly:    core.BoolPtr(false),
				WorkerBatchSize:    core.Int64Ptr(int64(400)),
				WorkerProcesses:    core.Int64Ptr(int64(3)),
			}
			replicationDocumentModel.Attachments["foo"] = *attachmentModel
			replicationDocumentModel.SetProperty("foo", "testString")

			putReplicationDocumentOptions := &cloudantv1.PutReplicationDocumentOptions{
				DocID:               core.StringPtr("testString"),
				ReplicationDocument: replicationDocumentModel,
				IfMatch:             core.StringPtr("testString"),
				Batch:               core.StringPtr("ok"),
				NewEdits:            core.BoolPtr(false),
				Rev:                 core.StringPtr("testString"),
			}

			documentResult, response, err := cloudantService.PutReplicationDocument(putReplicationDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`GetSchedulerDocs - Retrieve replication scheduler documents`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSchedulerDocs(getSchedulerDocsOptions *GetSchedulerDocsOptions)`, func() {
			getSchedulerDocsOptions := &cloudantv1.GetSchedulerDocsOptions{
				Limit:  core.Int64Ptr(int64(0)),
				Skip:   core.Int64Ptr(int64(0)),
				States: []string{"initializing"},
			}

			schedulerDocsResult, response, err := cloudantService.GetSchedulerDocs(getSchedulerDocsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schedulerDocsResult).ToNot(BeNil())
		})
	})

	Describe(`GetSchedulerDocument - Retrieve a replication scheduler document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSchedulerDocument(getSchedulerDocumentOptions *GetSchedulerDocumentOptions)`, func() {
			getSchedulerDocumentOptions := &cloudantv1.GetSchedulerDocumentOptions{
				DocID: core.StringPtr("testString"),
			}

			schedulerDocument, response, err := cloudantService.GetSchedulerDocument(getSchedulerDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schedulerDocument).ToNot(BeNil())
		})
	})

	Describe(`GetSchedulerJobs - Retrieve replication scheduler jobs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSchedulerJobs(getSchedulerJobsOptions *GetSchedulerJobsOptions)`, func() {
			getSchedulerJobsOptions := &cloudantv1.GetSchedulerJobsOptions{
				Limit: core.Int64Ptr(int64(25)),
				Skip:  core.Int64Ptr(int64(0)),
			}

			schedulerJobsResult, response, err := cloudantService.GetSchedulerJobs(getSchedulerJobsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schedulerJobsResult).ToNot(BeNil())
		})
	})

	Describe(`GetSchedulerJob - Retrieve a replication scheduler job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSchedulerJob(getSchedulerJobOptions *GetSchedulerJobOptions)`, func() {
			getSchedulerJobOptions := &cloudantv1.GetSchedulerJobOptions{
				JobID: core.StringPtr("testString"),
			}

			schedulerJob, response, err := cloudantService.GetSchedulerJob(getSchedulerJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schedulerJob).ToNot(BeNil())
		})
	})

	Describe(`GetSessionInformation - Retrieve current session cookie information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSessionInformation(getSessionInformationOptions *GetSessionInformationOptions)`, func() {
			getSessionInformationOptions := &cloudantv1.GetSessionInformationOptions{}

			sessionInformation, response, err := cloudantService.GetSessionInformation(getSessionInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sessionInformation).ToNot(BeNil())
		})
	})

	Describe(`GetSecurity - Retrieve database permissions information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSecurity(getSecurityOptions *GetSecurityOptions)`, func() {
			getSecurityOptions := &cloudantv1.GetSecurityOptions{
				Db: core.StringPtr("testString"),
			}

			security, response, err := cloudantService.GetSecurity(getSecurityOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(security).ToNot(BeNil())
		})
	})

	Describe(`PutSecurity - Modify database permissions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutSecurity(putSecurityOptions *PutSecurityOptions)`, func() {
			securityObjectModel := &cloudantv1.SecurityObject{
				Names: []string{"superuser"},
				Roles: []string{"admins"},
			}

			putSecurityOptions := &cloudantv1.PutSecurityOptions{
				Db:              core.StringPtr("testString"),
				Admins:          securityObjectModel,
				Members:         securityObjectModel,
				Cloudant:        map[string][]string{"key1": []string{"_reader"}},
				CouchdbAuthOnly: core.BoolPtr(true),
			}

			ok, response, err := cloudantService.PutSecurity(putSecurityOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ok).ToNot(BeNil())
		})
	})

	Describe(`PostApiKeys - Generates API keys for apps or persons to enable database access`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostApiKeys(postApiKeysOptions *PostApiKeysOptions)`, func() {
			postApiKeysOptions := &cloudantv1.PostApiKeysOptions{}

			apiKeysResult, response, err := cloudantService.PostApiKeys(postApiKeysOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKeysResult).ToNot(BeNil())
		})
	})

	Describe(`PutCloudantSecurityConfiguration - Modify only Cloudant related database permissions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutCloudantSecurityConfiguration(putCloudantSecurityConfigurationOptions *PutCloudantSecurityConfigurationOptions)`, func() {
			securityObjectModel := &cloudantv1.SecurityObject{
				Names: []string{"testString"},
				Roles: []string{"testString"},
			}

			putCloudantSecurityConfigurationOptions := &cloudantv1.PutCloudantSecurityConfigurationOptions{
				Db:              core.StringPtr("testString"),
				Cloudant:        map[string][]string{"key1": []string{"_reader"}},
				Admins:          securityObjectModel,
				Members:         securityObjectModel,
				CouchdbAuthOnly: core.BoolPtr(true),
			}

			ok, response, err := cloudantService.PutCloudantSecurityConfiguration(putCloudantSecurityConfigurationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ok).ToNot(BeNil())
		})
	})

	Describe(`GetCorsInformation - Retrieve CORS configuration information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCorsInformation(getCorsInformationOptions *GetCorsInformationOptions)`, func() {
			getCorsInformationOptions := &cloudantv1.GetCorsInformationOptions{}

			corsInformation, response, err := cloudantService.GetCorsInformation(getCorsInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(corsInformation).ToNot(BeNil())
		})
	})

	Describe(`PutCorsConfiguration - Modify CORS configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutCorsConfiguration(putCorsConfigurationOptions *PutCorsConfigurationOptions)`, func() {
			putCorsConfigurationOptions := &cloudantv1.PutCorsConfigurationOptions{
				Origins:          []string{"https://example.com", "https://www.example.com"},
				AllowCredentials: core.BoolPtr(true),
				EnableCors:       core.BoolPtr(true),
			}

			ok, response, err := cloudantService.PutCorsConfiguration(putCorsConfigurationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ok).ToNot(BeNil())
		})
	})

	Describe(`HeadAttachment - Retrieve the HTTP headers for an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HeadAttachment(headAttachmentOptions *HeadAttachmentOptions)`, func() {
			headAttachmentOptions := &cloudantv1.HeadAttachmentOptions{
				Db:             core.StringPtr("testString"),
				DocID:          core.StringPtr("testString"),
				AttachmentName: core.StringPtr("testString"),
				IfMatch:        core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
				Rev:            core.StringPtr("testString"),
			}

			response, err := cloudantService.HeadAttachment(headAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetAttachment - Retrieve an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAttachment(getAttachmentOptions *GetAttachmentOptions)`, func() {
			getAttachmentOptions := &cloudantv1.GetAttachmentOptions{
				Db:             core.StringPtr("testString"),
				DocID:          core.StringPtr("testString"),
				AttachmentName: core.StringPtr("testString"),
				Accept:         core.StringPtr("testString"),
				IfMatch:        core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
				Range:          core.StringPtr("testString"),
				Rev:            core.StringPtr("testString"),
			}

			result, response, err := cloudantService.GetAttachment(getAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PutAttachment - Create or modify an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutAttachment(putAttachmentOptions *PutAttachmentOptions)`, func() {
			putAttachmentOptions := &cloudantv1.PutAttachmentOptions{
				Db:             core.StringPtr("testString"),
				DocID:          core.StringPtr("testString"),
				AttachmentName: core.StringPtr("testString"),
				Attachment:     CreateMockReader("This is a mock file."),
				ContentType:    core.StringPtr("application/octet-stream"),
				IfMatch:        core.StringPtr("testString"),
				Rev:            core.StringPtr("testString"),
			}

			documentResult, response, err := cloudantService.PutAttachment(putAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`HeadLocalDocument - Retrieve HTTP headers for a local document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HeadLocalDocument(headLocalDocumentOptions *HeadLocalDocumentOptions)`, func() {
			headLocalDocumentOptions := &cloudantv1.HeadLocalDocumentOptions{
				Db:          core.StringPtr("testString"),
				DocID:       core.StringPtr("testString"),
				IfNoneMatch: core.StringPtr("testString"),
			}

			response, err := cloudantService.HeadLocalDocument(headLocalDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetLocalDocument - Retrieve a local document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLocalDocument(getLocalDocumentOptions *GetLocalDocumentOptions)`, func() {
			getLocalDocumentOptions := &cloudantv1.GetLocalDocumentOptions{
				Db:              core.StringPtr("testString"),
				DocID:           core.StringPtr("testString"),
				Accept:          core.StringPtr("application/json"),
				IfNoneMatch:     core.StringPtr("testString"),
				Attachments:     core.BoolPtr(false),
				AttEncodingInfo: core.BoolPtr(false),
				LocalSeq:        core.BoolPtr(false),
			}

			document, response, err := cloudantService.GetLocalDocument(getLocalDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(document).ToNot(BeNil())
		})
	})

	Describe(`PutLocalDocument - Create or modify a local document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutLocalDocument(putLocalDocumentOptions *PutLocalDocumentOptions)`, func() {
			attachmentModel := &cloudantv1.Attachment{
				ContentType:   core.StringPtr("testString"),
				Data:          CreateMockByteArray("VGhpcyBpcyBhIG1vY2sgYnl0ZSBhcnJheSB2YWx1ZS4="),
				Digest:        core.StringPtr("testString"),
				EncodedLength: core.Int64Ptr(int64(0)),
				Encoding:      core.StringPtr("testString"),
				Follows:       core.BoolPtr(true),
				Length:        core.Int64Ptr(int64(0)),
				Revpos:        core.Int64Ptr(int64(1)),
				Stub:          core.BoolPtr(true),
			}

			revisionsModel := &cloudantv1.Revisions{
				Ids:   []string{"testString"},
				Start: core.Int64Ptr(int64(1)),
			}

			documentRevisionStatusModel := &cloudantv1.DocumentRevisionStatus{
				Rev:    core.StringPtr("testString"),
				Status: core.StringPtr("available"),
			}

			documentModel := &cloudantv1.Document{
				Attachments:      map[string]cloudantv1.Attachment{"key1": *attachmentModel},
				Conflicts:        []string{"testString"},
				Deleted:          core.BoolPtr(true),
				DeletedConflicts: []string{"testString"},
				ID:               core.StringPtr("exampleid"),
				LocalSeq:         core.StringPtr("testString"),
				Rev:              core.StringPtr("testString"),
				Revisions:        revisionsModel,
				RevsInfo:         []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel},
			}
			documentModel.Attachments["foo"] = *attachmentModel
			documentModel.SetProperty("brand", core.StringPtr("Foo"))
			documentModel.SetProperty("colours", core.StringPtr("[\"red\",\"green\",\"black\",\"blue\"]"))
			documentModel.SetProperty("description", core.StringPtr("Slim Colourful Design Electronic Cooking Appliance for ..."))
			documentModel.SetProperty("image", core.StringPtr("assets/img/0gmsnghhew.jpg"))
			documentModel.SetProperty("keywords", core.StringPtr("[\"Foo\",\"Scales\",\"Weight\",\"Digital\",\"Kitchen\"]"))
			documentModel.SetProperty("name", core.StringPtr("Digital Kitchen Scales"))
			documentModel.SetProperty("price", core.StringPtr("14.99"))
			documentModel.SetProperty("productid", core.StringPtr("1000042"))
			documentModel.SetProperty("taxonomy", core.StringPtr("[\"Home\",\"Kitchen\",\"Small Appliances\"]"))
			documentModel.SetProperty("type", core.StringPtr("product"))

			putLocalDocumentOptions := &cloudantv1.PutLocalDocumentOptions{
				Db:          core.StringPtr("testString"),
				DocID:       core.StringPtr("testString"),
				Document:    documentModel,
				ContentType: core.StringPtr("application/json"),
				Batch:       core.StringPtr("ok"),
			}

			documentResult, response, err := cloudantService.PutLocalDocument(putLocalDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`PostRevsDiff - Query the document revisions and possible ancestors missing from the database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostRevsDiff(postRevsDiffOptions *PostRevsDiffOptions)`, func() {
			postRevsDiffOptions := &cloudantv1.PostRevsDiffOptions{
				Db:                core.StringPtr("testString"),
				DocumentRevisions: map[string][]string{"key1": []string{"testString"}},
			}

			mapStringRevsDiff, response, err := cloudantService.PostRevsDiff(postRevsDiffOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringRevsDiff).ToNot(BeNil())
		})
	})

	Describe(`GetShardsInformation - Retrieve shard information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetShardsInformation(getShardsInformationOptions *GetShardsInformationOptions)`, func() {
			getShardsInformationOptions := &cloudantv1.GetShardsInformationOptions{
				Db: core.StringPtr("testString"),
			}

			shardsInformation, response, err := cloudantService.GetShardsInformation(getShardsInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shardsInformation).ToNot(BeNil())
		})
	})

	Describe(`GetDocumentShardsInfo - Retrieve shard information for a specific document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDocumentShardsInfo(getDocumentShardsInfoOptions *GetDocumentShardsInfoOptions)`, func() {
			getDocumentShardsInfoOptions := &cloudantv1.GetDocumentShardsInfoOptions{
				Db:    core.StringPtr("testString"),
				DocID: core.StringPtr("testString"),
			}

			documentShardInfo, response, err := cloudantService.GetDocumentShardsInfo(getDocumentShardsInfoOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(documentShardInfo).ToNot(BeNil())
		})
	})

	Describe(`HeadUpInformation - Retrieve HTTP headers about whether the server is up`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`HeadUpInformation(headUpInformationOptions *HeadUpInformationOptions)`, func() {
			headUpInformationOptions := &cloudantv1.HeadUpInformationOptions{}

			response, err := cloudantService.HeadUpInformation(headUpInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetActiveTasks - Retrieve list of running tasks`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetActiveTasks(getActiveTasksOptions *GetActiveTasksOptions)`, func() {
			getActiveTasksOptions := &cloudantv1.GetActiveTasksOptions{}

			activeTask, response, err := cloudantService.GetActiveTasks(getActiveTasksOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(activeTask).ToNot(BeNil())
		})
	})

	Describe(`GetMembershipInformation - Retrieve cluster membership information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMembershipInformation(getMembershipInformationOptions *GetMembershipInformationOptions)`, func() {
			getMembershipInformationOptions := &cloudantv1.GetMembershipInformationOptions{}

			membershipInformation, response, err := cloudantService.GetMembershipInformation(getMembershipInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(membershipInformation).ToNot(BeNil())
		})
	})

	Describe(`GetUpInformation - Retrieve information about whether the server is up`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUpInformation(getUpInformationOptions *GetUpInformationOptions)`, func() {
			getUpInformationOptions := &cloudantv1.GetUpInformationOptions{}

			upInformation, response, err := cloudantService.GetUpInformation(getUpInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(upInformation).ToNot(BeNil())
		})
	})

	Describe(`GetActivityTrackerEvents - Retrieve Activity Tracker events information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetActivityTrackerEvents(getActivityTrackerEventsOptions *GetActivityTrackerEventsOptions)`, func() {
			getActivityTrackerEventsOptions := &cloudantv1.GetActivityTrackerEventsOptions{}

			activityTrackerEvents, response, err := cloudantService.GetActivityTrackerEvents(getActivityTrackerEventsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(activityTrackerEvents).ToNot(BeNil())
		})
	})

	Describe(`PostActivityTrackerEvents - Modify Activity Tracker events configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostActivityTrackerEvents(postActivityTrackerEventsOptions *PostActivityTrackerEventsOptions)`, func() {
			postActivityTrackerEventsOptions := &cloudantv1.PostActivityTrackerEventsOptions{
				Types: []string{"management", "data"},
			}

			ok, response, err := cloudantService.PostActivityTrackerEvents(postActivityTrackerEventsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ok).ToNot(BeNil())
		})
	})

	Describe(`GetCurrentThroughputInformation - Retrieve the current provisioned throughput capacity consumption`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCurrentThroughputInformation(getCurrentThroughputInformationOptions *GetCurrentThroughputInformationOptions)`, func() {
			getCurrentThroughputInformationOptions := &cloudantv1.GetCurrentThroughputInformationOptions{}

			currentThroughputInformation, response, err := cloudantService.GetCurrentThroughputInformation(getCurrentThroughputInformationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(currentThroughputInformation).ToNot(BeNil())
		})
	})

	Describe(`DeleteDatabase - Delete a database`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDatabase(deleteDatabaseOptions *DeleteDatabaseOptions)`, func() {
			deleteDatabaseOptions := &cloudantv1.DeleteDatabaseOptions{
				Db: core.StringPtr("testString"),
			}

			ok, response, err := cloudantService.DeleteDatabase(deleteDatabaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ok).ToNot(BeNil())
		})
	})

	Describe(`DeleteDocument - Delete a document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions)`, func() {
			deleteDocumentOptions := &cloudantv1.DeleteDocumentOptions{
				Db:      core.StringPtr("testString"),
				DocID:   core.StringPtr("testString"),
				IfMatch: core.StringPtr("testString"),
				Batch:   core.StringPtr("ok"),
				Rev:     core.StringPtr("testString"),
			}

			documentResult, response, err := cloudantService.DeleteDocument(deleteDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`DeleteDesignDocument - Delete a design document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDesignDocument(deleteDesignDocumentOptions *DeleteDesignDocumentOptions)`, func() {
			deleteDesignDocumentOptions := &cloudantv1.DeleteDesignDocumentOptions{
				Db:      core.StringPtr("testString"),
				Ddoc:    core.StringPtr("testString"),
				IfMatch: core.StringPtr("testString"),
				Batch:   core.StringPtr("ok"),
				Rev:     core.StringPtr("testString"),
			}

			documentResult, response, err := cloudantService.DeleteDesignDocument(deleteDesignDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`DeleteIndex - Delete an index`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteIndex(deleteIndexOptions *DeleteIndexOptions)`, func() {
			deleteIndexOptions := &cloudantv1.DeleteIndexOptions{
				Db:    core.StringPtr("testString"),
				Ddoc:  core.StringPtr("testString"),
				Type:  core.StringPtr("json"),
				Index: core.StringPtr("testString"),
			}

			ok, response, err := cloudantService.DeleteIndex(deleteIndexOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ok).ToNot(BeNil())
		})
	})

	Describe(`DeleteReplicationDocument - Cancel a persistent replication`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteReplicationDocument(deleteReplicationDocumentOptions *DeleteReplicationDocumentOptions)`, func() {
			deleteReplicationDocumentOptions := &cloudantv1.DeleteReplicationDocumentOptions{
				DocID:   core.StringPtr("testString"),
				IfMatch: core.StringPtr("testString"),
				Batch:   core.StringPtr("ok"),
				Rev:     core.StringPtr("testString"),
			}

			documentResult, response, err := cloudantService.DeleteReplicationDocument(deleteReplicationDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`DeleteAttachment - Delete an attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAttachment(deleteAttachmentOptions *DeleteAttachmentOptions)`, func() {
			deleteAttachmentOptions := &cloudantv1.DeleteAttachmentOptions{
				Db:             core.StringPtr("testString"),
				DocID:          core.StringPtr("testString"),
				AttachmentName: core.StringPtr("testString"),
				IfMatch:        core.StringPtr("testString"),
				Rev:            core.StringPtr("testString"),
				Batch:          core.StringPtr("ok"),
			}

			documentResult, response, err := cloudantService.DeleteAttachment(deleteAttachmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(documentResult).ToNot(BeNil())
		})
	})

	Describe(`DeleteLocalDocument - Delete a local document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLocalDocument(deleteLocalDocumentOptions *DeleteLocalDocumentOptions)`, func() {
			deleteLocalDocumentOptions := &cloudantv1.DeleteLocalDocumentOptions{
				Db:    core.StringPtr("testString"),
				DocID: core.StringPtr("testString"),
				Batch: core.StringPtr("ok"),
			}

			documentResult, response, err := cloudantService.DeleteLocalDocument(deleteLocalDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(documentResult).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
