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

package cloudantv1_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

const externalConfigFile = "../cloudant.env"

var (
	err          error
	service      *cloudantv1.CloudantV1
	serviceURL   string
	config       map[string]string
	configLoaded = false
	dbName       string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`ExampleServiceV1 Integration Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties("SERVER")
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			configLoaded = true
			fmt.Printf("Service URL: %s\n", serviceURL)

			var found bool
			dbName, found = os.LookupEnv("DATABASE_NAME")
			if !found {
				dbName = "stores"
			}
		})
	})

	Describe(`Service-level tests`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			service, err = cloudantv1.NewCloudantV1UsingExternalConfig(
				&cloudantv1.CloudantV1Options{ServiceName: "SERVER"})
			Expect(err).To(BeNil())
			Expect(service).ToNot(BeNil())
			Expect(service.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`validate`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`server information`, func() {
			result, _, err := service.GetServerInformation(&cloudantv1.GetServerInformationOptions{})
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())

			Expect(*result.Couchdb).ToNot(BeNil())
			Expect(*result.Version).ToNot(BeNil())
		})

		It(`db exists`, func() {
			response, err := service.HeadDatabase(&cloudantv1.HeadDatabaseOptions{Db: &dbName})
			Expect(err).To(BeNil())
			Expect(response).ToNot(BeNil())

			Expect(len(response.Headers)).ToNot(BeZero())
		})

		It(`all docs`, func() {
			result, _, err := service.PostAllDocs(&cloudantv1.PostAllDocsOptions{Db: &dbName})
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())

			Expect(len(result.Rows)).ToNot(BeZero())
		})
	})
})
