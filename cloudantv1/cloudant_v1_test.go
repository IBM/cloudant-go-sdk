/**
 * (C) Copyright IBM Corp. 2020.
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
	"bytes"
	"fmt"
	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`CloudantV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetServerInformation(getServerInformationOptions *GetServerInformationOptions) - Operation response error`, func() {
		getServerInformationPath := "/"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServerInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetServerInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetServerInformationOptions model
				getServerInformationOptionsModel := new(cloudantv1.GetServerInformationOptions)
				getServerInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetServerInformation(getServerInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetServerInformation(getServerInformationOptions *GetServerInformationOptions)`, func() {
		getServerInformationPath := "/"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServerInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"couchdb": "Couchdb", "features": ["Features"], "vendor": {"name": "Name", "variant": "Variant", "version": "Version"}, "version": "Version"}`)
				}))
			})
			It(`Invoke GetServerInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetServerInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServerInformationOptions model
				getServerInformationOptionsModel := new(cloudantv1.GetServerInformationOptions)
				getServerInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetServerInformation(getServerInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetServerInformation with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetServerInformationOptions model
				getServerInformationOptionsModel := new(cloudantv1.GetServerInformationOptions)
				getServerInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetServerInformation(getServerInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMembershipInformation(getMembershipInformationOptions *GetMembershipInformationOptions) - Operation response error`, func() {
		getMembershipInformationPath := "/_membership"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMembershipInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMembershipInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetMembershipInformationOptions model
				getMembershipInformationOptionsModel := new(cloudantv1.GetMembershipInformationOptions)
				getMembershipInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetMembershipInformation(getMembershipInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMembershipInformation(getMembershipInformationOptions *GetMembershipInformationOptions)`, func() {
		getMembershipInformationPath := "/_membership"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMembershipInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"all_nodes": ["AllNodes"], "cluster_nodes": ["ClusterNodes"]}`)
				}))
			})
			It(`Invoke GetMembershipInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetMembershipInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMembershipInformationOptions model
				getMembershipInformationOptionsModel := new(cloudantv1.GetMembershipInformationOptions)
				getMembershipInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetMembershipInformation(getMembershipInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetMembershipInformation with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetMembershipInformationOptions model
				getMembershipInformationOptionsModel := new(cloudantv1.GetMembershipInformationOptions)
				getMembershipInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetMembershipInformation(getMembershipInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUuids(getUuidsOptions *GetUuidsOptions) - Operation response error`, func() {
		getUuidsPath := "/_uuids"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUuidsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for count query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUuids with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetUuidsOptions model
				getUuidsOptionsModel := new(cloudantv1.GetUuidsOptions)
				getUuidsOptionsModel.Count = core.Int64Ptr(int64(1))
				getUuidsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetUuids(getUuidsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUuids(getUuidsOptions *GetUuidsOptions)`, func() {
		getUuidsPath := "/_uuids"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUuidsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for count query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"uuids": ["Uuids"]}`)
				}))
			})
			It(`Invoke GetUuids successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetUuids(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUuidsOptions model
				getUuidsOptionsModel := new(cloudantv1.GetUuidsOptions)
				getUuidsOptionsModel.Count = core.Int64Ptr(int64(1))
				getUuidsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetUuids(getUuidsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetUuids with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetUuidsOptions model
				getUuidsOptionsModel := new(cloudantv1.GetUuidsOptions)
				getUuidsOptionsModel.Count = core.Int64Ptr(int64(1))
				getUuidsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetUuids(getUuidsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`HeadDatabase(headDatabaseOptions *HeadDatabaseOptions)`, func() {
		headDatabasePath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(headDatabasePath))
					Expect(req.Method).To(Equal("HEAD"))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke HeadDatabase successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cloudantService.HeadDatabase(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HeadDatabaseOptions model
				headDatabaseOptionsModel := new(cloudantv1.HeadDatabaseOptions)
				headDatabaseOptionsModel.Db = core.StringPtr("testString")
				headDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cloudantService.HeadDatabase(headDatabaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HeadDatabase with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the HeadDatabaseOptions model
				headDatabaseOptionsModel := new(cloudantv1.HeadDatabaseOptions)
				headDatabaseOptionsModel.Db = core.StringPtr("testString")
				headDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cloudantService.HeadDatabase(headDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HeadDatabaseOptions model with no property values
				headDatabaseOptionsModelNew := new(cloudantv1.HeadDatabaseOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cloudantService.HeadDatabase(headDatabaseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAllDbs(getAllDbsOptions *GetAllDbsOptions)`, func() {
		getAllDbsPath := "/_all_dbs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllDbsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for descending query parameter

					Expect(req.URL.Query()["endkey"]).To(Equal([]string{"testString"}))


					// TODO: Add check for limit query parameter


					// TODO: Add check for skip query parameter

					Expect(req.URL.Query()["startkey"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `["OperationResponse"]`)
				}))
			})
			It(`Invoke GetAllDbs successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetAllDbs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAllDbsOptions model
				getAllDbsOptionsModel := new(cloudantv1.GetAllDbsOptions)
				getAllDbsOptionsModel.Descending = core.BoolPtr(true)
				getAllDbsOptionsModel.Endkey = core.StringPtr("testString")
				getAllDbsOptionsModel.Limit = core.Int64Ptr(int64(0))
				getAllDbsOptionsModel.Skip = core.Int64Ptr(int64(0))
				getAllDbsOptionsModel.Startkey = core.StringPtr("testString")
				getAllDbsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetAllDbs(getAllDbsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetAllDbs with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetAllDbsOptions model
				getAllDbsOptionsModel := new(cloudantv1.GetAllDbsOptions)
				getAllDbsOptionsModel.Descending = core.BoolPtr(true)
				getAllDbsOptionsModel.Endkey = core.StringPtr("testString")
				getAllDbsOptionsModel.Limit = core.Int64Ptr(int64(0))
				getAllDbsOptionsModel.Skip = core.Int64Ptr(int64(0))
				getAllDbsOptionsModel.Startkey = core.StringPtr("testString")
				getAllDbsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetAllDbs(getAllDbsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDbsInfo(postDbsInfoOptions *PostDbsInfoOptions) - Operation response error`, func() {
		postDbsInfoPath := "/_dbs_info"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDbsInfoPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDbsInfo with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostDbsInfoOptions model
				postDbsInfoOptionsModel := new(cloudantv1.PostDbsInfoOptions)
				postDbsInfoOptionsModel.Keys = []string{"testString"}
				postDbsInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostDbsInfo(postDbsInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostDbsInfo(postDbsInfoOptions *PostDbsInfoOptions)`, func() {
		postDbsInfoPath := "/_dbs_info"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDbsInfoPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"info": {"cluster": {"n": 1, "q": 1, "r": 1, "w": 1}, "committed_update_seq": "CommittedUpdateSeq", "compact_running": true, "compacted_seq": "CompactedSeq", "db_name": "DbName", "disk_format_version": 17, "doc_count": 0, "doc_del_count": 0, "engine": "Engine", "props": {"partitioned": false}, "sizes": {"active": 6, "external": 8, "file": 4}, "update_seq": "UpdateSeq", "uuid": "UUID"}, "key": "Key"}]`)
				}))
			})
			It(`Invoke PostDbsInfo successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostDbsInfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostDbsInfoOptions model
				postDbsInfoOptionsModel := new(cloudantv1.PostDbsInfoOptions)
				postDbsInfoOptionsModel.Keys = []string{"testString"}
				postDbsInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostDbsInfo(postDbsInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostDbsInfo with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostDbsInfoOptions model
				postDbsInfoOptionsModel := new(cloudantv1.PostDbsInfoOptions)
				postDbsInfoOptionsModel.Keys = []string{"testString"}
				postDbsInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostDbsInfo(postDbsInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDatabase(deleteDatabaseOptions *DeleteDatabaseOptions) - Operation response error`, func() {
		deleteDatabasePath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabasePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDatabase with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteDatabaseOptions model
				deleteDatabaseOptionsModel := new(cloudantv1.DeleteDatabaseOptions)
				deleteDatabaseOptionsModel.Db = core.StringPtr("testString")
				deleteDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.DeleteDatabase(deleteDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteDatabase(deleteDatabaseOptions *DeleteDatabaseOptions)`, func() {
		deleteDatabasePath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabasePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ok": true}`)
				}))
			})
			It(`Invoke DeleteDatabase successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.DeleteDatabase(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDatabaseOptions model
				deleteDatabaseOptionsModel := new(cloudantv1.DeleteDatabaseOptions)
				deleteDatabaseOptionsModel.Db = core.StringPtr("testString")
				deleteDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.DeleteDatabase(deleteDatabaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteDatabase with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteDatabaseOptions model
				deleteDatabaseOptionsModel := new(cloudantv1.DeleteDatabaseOptions)
				deleteDatabaseOptionsModel.Db = core.StringPtr("testString")
				deleteDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.DeleteDatabase(deleteDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDatabaseOptions model with no property values
				deleteDatabaseOptionsModelNew := new(cloudantv1.DeleteDatabaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.DeleteDatabase(deleteDatabaseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDatabaseInformation(getDatabaseInformationOptions *GetDatabaseInformationOptions) - Operation response error`, func() {
		getDatabaseInformationPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDatabaseInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDatabaseInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDatabaseInformationOptions model
				getDatabaseInformationOptionsModel := new(cloudantv1.GetDatabaseInformationOptions)
				getDatabaseInformationOptionsModel.Db = core.StringPtr("testString")
				getDatabaseInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetDatabaseInformation(getDatabaseInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDatabaseInformation(getDatabaseInformationOptions *GetDatabaseInformationOptions)`, func() {
		getDatabaseInformationPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDatabaseInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cluster": {"n": 1, "q": 1, "r": 1, "w": 1}, "committed_update_seq": "CommittedUpdateSeq", "compact_running": true, "compacted_seq": "CompactedSeq", "db_name": "DbName", "disk_format_version": 17, "doc_count": 0, "doc_del_count": 0, "engine": "Engine", "props": {"partitioned": false}, "sizes": {"active": 6, "external": 8, "file": 4}, "update_seq": "UpdateSeq", "uuid": "UUID"}`)
				}))
			})
			It(`Invoke GetDatabaseInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetDatabaseInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDatabaseInformationOptions model
				getDatabaseInformationOptionsModel := new(cloudantv1.GetDatabaseInformationOptions)
				getDatabaseInformationOptionsModel.Db = core.StringPtr("testString")
				getDatabaseInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetDatabaseInformation(getDatabaseInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDatabaseInformation with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDatabaseInformationOptions model
				getDatabaseInformationOptionsModel := new(cloudantv1.GetDatabaseInformationOptions)
				getDatabaseInformationOptionsModel.Db = core.StringPtr("testString")
				getDatabaseInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetDatabaseInformation(getDatabaseInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDatabaseInformationOptions model with no property values
				getDatabaseInformationOptionsModelNew := new(cloudantv1.GetDatabaseInformationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetDatabaseInformation(getDatabaseInformationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutDatabase(putDatabaseOptions *PutDatabaseOptions) - Operation response error`, func() {
		putDatabasePath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDatabasePath))
					Expect(req.Method).To(Equal("PUT"))

					// TODO: Add check for partitioned query parameter


					// TODO: Add check for q query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutDatabase with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PutDatabaseOptions model
				putDatabaseOptionsModel := new(cloudantv1.PutDatabaseOptions)
				putDatabaseOptionsModel.Db = core.StringPtr("testString")
				putDatabaseOptionsModel.Partitioned = core.BoolPtr(true)
				putDatabaseOptionsModel.Q = core.Int64Ptr(int64(1))
				putDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PutDatabase(putDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutDatabase(putDatabaseOptions *PutDatabaseOptions)`, func() {
		putDatabasePath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDatabasePath))
					Expect(req.Method).To(Equal("PUT"))

					// TODO: Add check for partitioned query parameter


					// TODO: Add check for q query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"ok": true}`)
				}))
			})
			It(`Invoke PutDatabase successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PutDatabase(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PutDatabaseOptions model
				putDatabaseOptionsModel := new(cloudantv1.PutDatabaseOptions)
				putDatabaseOptionsModel.Db = core.StringPtr("testString")
				putDatabaseOptionsModel.Partitioned = core.BoolPtr(true)
				putDatabaseOptionsModel.Q = core.Int64Ptr(int64(1))
				putDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PutDatabase(putDatabaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PutDatabase with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PutDatabaseOptions model
				putDatabaseOptionsModel := new(cloudantv1.PutDatabaseOptions)
				putDatabaseOptionsModel.Db = core.StringPtr("testString")
				putDatabaseOptionsModel.Partitioned = core.BoolPtr(true)
				putDatabaseOptionsModel.Q = core.Int64Ptr(int64(1))
				putDatabaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PutDatabase(putDatabaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutDatabaseOptions model with no property values
				putDatabaseOptionsModelNew := new(cloudantv1.PutDatabaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PutDatabase(putDatabaseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostChanges(postChangesOptions *PostChangesOptions) - Operation response error`, func() {
		postChangesPath := "/testString/_changes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postChangesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Last-Event-Id"]).ToNot(BeNil())
					Expect(req.Header["Last-Event-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for attachments query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for descending query parameter

					Expect(req.URL.Query()["feed"]).To(Equal([]string{"continuous"}))

					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))


					// TODO: Add check for heartbeat query parameter


					// TODO: Add check for include_docs query parameter


					// TODO: Add check for limit query parameter


					// TODO: Add check for seq_interval query parameter

					Expect(req.URL.Query()["since"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["style"]).To(Equal([]string{"testString"}))


					// TODO: Add check for timeout query parameter

					Expect(req.URL.Query()["view"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostChanges with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostChangesOptions model
				postChangesOptionsModel := new(cloudantv1.PostChangesOptions)
				postChangesOptionsModel.Db = core.StringPtr("testString")
				postChangesOptionsModel.DocIds = []string{"testString"}
				postChangesOptionsModel.Fields = []string{"testString"}
				postChangesOptionsModel.Selector = make(map[string]interface{})
				postChangesOptionsModel.LastEventID = core.StringPtr("testString")
				postChangesOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postChangesOptionsModel.Attachments = core.BoolPtr(true)
				postChangesOptionsModel.Conflicts = core.BoolPtr(true)
				postChangesOptionsModel.Descending = core.BoolPtr(true)
				postChangesOptionsModel.Feed = core.StringPtr("continuous")
				postChangesOptionsModel.Filter = core.StringPtr("testString")
				postChangesOptionsModel.Heartbeat = core.Int64Ptr(int64(0))
				postChangesOptionsModel.IncludeDocs = core.BoolPtr(true)
				postChangesOptionsModel.Limit = core.Int64Ptr(int64(0))
				postChangesOptionsModel.SeqInterval = core.Int64Ptr(int64(1))
				postChangesOptionsModel.Since = core.StringPtr("testString")
				postChangesOptionsModel.Style = core.StringPtr("testString")
				postChangesOptionsModel.Timeout = core.Int64Ptr(int64(0))
				postChangesOptionsModel.View = core.StringPtr("testString")
				postChangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostChanges(postChangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostChanges(postChangesOptions *PostChangesOptions)`, func() {
		postChangesPath := "/testString/_changes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postChangesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Last-Event-Id"]).ToNot(BeNil())
					Expect(req.Header["Last-Event-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for attachments query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for descending query parameter

					Expect(req.URL.Query()["feed"]).To(Equal([]string{"continuous"}))

					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))


					// TODO: Add check for heartbeat query parameter


					// TODO: Add check for include_docs query parameter


					// TODO: Add check for limit query parameter


					// TODO: Add check for seq_interval query parameter

					Expect(req.URL.Query()["since"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["style"]).To(Equal([]string{"testString"}))


					// TODO: Add check for timeout query parameter

					Expect(req.URL.Query()["view"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"last_seq": "LastSeq", "pending": 7, "results": [{"changes": [{"rev": "Rev"}], "deleted": false, "id": "ID", "seq": "Seq"}]}`)
				}))
			})
			It(`Invoke PostChanges successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostChanges(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostChangesOptions model
				postChangesOptionsModel := new(cloudantv1.PostChangesOptions)
				postChangesOptionsModel.Db = core.StringPtr("testString")
				postChangesOptionsModel.DocIds = []string{"testString"}
				postChangesOptionsModel.Fields = []string{"testString"}
				postChangesOptionsModel.Selector = make(map[string]interface{})
				postChangesOptionsModel.LastEventID = core.StringPtr("testString")
				postChangesOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postChangesOptionsModel.Attachments = core.BoolPtr(true)
				postChangesOptionsModel.Conflicts = core.BoolPtr(true)
				postChangesOptionsModel.Descending = core.BoolPtr(true)
				postChangesOptionsModel.Feed = core.StringPtr("continuous")
				postChangesOptionsModel.Filter = core.StringPtr("testString")
				postChangesOptionsModel.Heartbeat = core.Int64Ptr(int64(0))
				postChangesOptionsModel.IncludeDocs = core.BoolPtr(true)
				postChangesOptionsModel.Limit = core.Int64Ptr(int64(0))
				postChangesOptionsModel.SeqInterval = core.Int64Ptr(int64(1))
				postChangesOptionsModel.Since = core.StringPtr("testString")
				postChangesOptionsModel.Style = core.StringPtr("testString")
				postChangesOptionsModel.Timeout = core.Int64Ptr(int64(0))
				postChangesOptionsModel.View = core.StringPtr("testString")
				postChangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostChanges(postChangesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostChanges with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostChangesOptions model
				postChangesOptionsModel := new(cloudantv1.PostChangesOptions)
				postChangesOptionsModel.Db = core.StringPtr("testString")
				postChangesOptionsModel.DocIds = []string{"testString"}
				postChangesOptionsModel.Fields = []string{"testString"}
				postChangesOptionsModel.Selector = make(map[string]interface{})
				postChangesOptionsModel.LastEventID = core.StringPtr("testString")
				postChangesOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postChangesOptionsModel.Attachments = core.BoolPtr(true)
				postChangesOptionsModel.Conflicts = core.BoolPtr(true)
				postChangesOptionsModel.Descending = core.BoolPtr(true)
				postChangesOptionsModel.Feed = core.StringPtr("continuous")
				postChangesOptionsModel.Filter = core.StringPtr("testString")
				postChangesOptionsModel.Heartbeat = core.Int64Ptr(int64(0))
				postChangesOptionsModel.IncludeDocs = core.BoolPtr(true)
				postChangesOptionsModel.Limit = core.Int64Ptr(int64(0))
				postChangesOptionsModel.SeqInterval = core.Int64Ptr(int64(1))
				postChangesOptionsModel.Since = core.StringPtr("testString")
				postChangesOptionsModel.Style = core.StringPtr("testString")
				postChangesOptionsModel.Timeout = core.Int64Ptr(int64(0))
				postChangesOptionsModel.View = core.StringPtr("testString")
				postChangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostChanges(postChangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostChangesOptions model with no property values
				postChangesOptionsModelNew := new(cloudantv1.PostChangesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostChanges(postChangesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostChangesAsStream(postChangesOptions *PostChangesOptions)`, func() {
		postChangesAsStreamPath := "/testString/_changes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postChangesAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Last-Event-Id"]).ToNot(BeNil())
					Expect(req.Header["Last-Event-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for attachments query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for descending query parameter

					Expect(req.URL.Query()["feed"]).To(Equal([]string{"continuous"}))

					Expect(req.URL.Query()["filter"]).To(Equal([]string{"testString"}))


					// TODO: Add check for heartbeat query parameter


					// TODO: Add check for include_docs query parameter


					// TODO: Add check for limit query parameter


					// TODO: Add check for seq_interval query parameter

					Expect(req.URL.Query()["since"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["style"]).To(Equal([]string{"testString"}))


					// TODO: Add check for timeout query parameter

					Expect(req.URL.Query()["view"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostChangesAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostChangesAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostChangesOptions model
				postChangesOptionsModel := new(cloudantv1.PostChangesOptions)
				postChangesOptionsModel.Db = core.StringPtr("testString")
				postChangesOptionsModel.DocIds = []string{"testString"}
				postChangesOptionsModel.Fields = []string{"testString"}
				postChangesOptionsModel.Selector = make(map[string]interface{})
				postChangesOptionsModel.LastEventID = core.StringPtr("testString")
				postChangesOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postChangesOptionsModel.Attachments = core.BoolPtr(true)
				postChangesOptionsModel.Conflicts = core.BoolPtr(true)
				postChangesOptionsModel.Descending = core.BoolPtr(true)
				postChangesOptionsModel.Feed = core.StringPtr("continuous")
				postChangesOptionsModel.Filter = core.StringPtr("testString")
				postChangesOptionsModel.Heartbeat = core.Int64Ptr(int64(0))
				postChangesOptionsModel.IncludeDocs = core.BoolPtr(true)
				postChangesOptionsModel.Limit = core.Int64Ptr(int64(0))
				postChangesOptionsModel.SeqInterval = core.Int64Ptr(int64(1))
				postChangesOptionsModel.Since = core.StringPtr("testString")
				postChangesOptionsModel.Style = core.StringPtr("testString")
				postChangesOptionsModel.Timeout = core.Int64Ptr(int64(0))
				postChangesOptionsModel.View = core.StringPtr("testString")
				postChangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostChangesAsStream(postChangesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostChangesAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostChangesOptions model
				postChangesOptionsModel := new(cloudantv1.PostChangesOptions)
				postChangesOptionsModel.Db = core.StringPtr("testString")
				postChangesOptionsModel.DocIds = []string{"testString"}
				postChangesOptionsModel.Fields = []string{"testString"}
				postChangesOptionsModel.Selector = make(map[string]interface{})
				postChangesOptionsModel.LastEventID = core.StringPtr("testString")
				postChangesOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postChangesOptionsModel.Attachments = core.BoolPtr(true)
				postChangesOptionsModel.Conflicts = core.BoolPtr(true)
				postChangesOptionsModel.Descending = core.BoolPtr(true)
				postChangesOptionsModel.Feed = core.StringPtr("continuous")
				postChangesOptionsModel.Filter = core.StringPtr("testString")
				postChangesOptionsModel.Heartbeat = core.Int64Ptr(int64(0))
				postChangesOptionsModel.IncludeDocs = core.BoolPtr(true)
				postChangesOptionsModel.Limit = core.Int64Ptr(int64(0))
				postChangesOptionsModel.SeqInterval = core.Int64Ptr(int64(1))
				postChangesOptionsModel.Since = core.StringPtr("testString")
				postChangesOptionsModel.Style = core.StringPtr("testString")
				postChangesOptionsModel.Timeout = core.Int64Ptr(int64(0))
				postChangesOptionsModel.View = core.StringPtr("testString")
				postChangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostChangesAsStream(postChangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostChangesOptions model with no property values
				postChangesOptionsModelNew := new(cloudantv1.PostChangesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostChangesAsStream(postChangesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`HeadDocument(headDocumentOptions *HeadDocumentOptions)`, func() {
		headDocumentPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(headDocumentPath))
					Expect(req.Method).To(Equal("HEAD"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for latest query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke HeadDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cloudantService.HeadDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HeadDocumentOptions model
				headDocumentOptionsModel := new(cloudantv1.HeadDocumentOptions)
				headDocumentOptionsModel.Db = core.StringPtr("testString")
				headDocumentOptionsModel.DocID = core.StringPtr("testString")
				headDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				headDocumentOptionsModel.Latest = core.BoolPtr(true)
				headDocumentOptionsModel.Rev = core.StringPtr("testString")
				headDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cloudantService.HeadDocument(headDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HeadDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the HeadDocumentOptions model
				headDocumentOptionsModel := new(cloudantv1.HeadDocumentOptions)
				headDocumentOptionsModel.Db = core.StringPtr("testString")
				headDocumentOptionsModel.DocID = core.StringPtr("testString")
				headDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				headDocumentOptionsModel.Latest = core.BoolPtr(true)
				headDocumentOptionsModel.Rev = core.StringPtr("testString")
				headDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cloudantService.HeadDocument(headDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HeadDocumentOptions model with no property values
				headDocumentOptionsModelNew := new(cloudantv1.HeadDocumentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cloudantService.HeadDocument(headDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDocument(postDocumentOptions *PostDocumentOptions) - Operation response error`, func() {
		postDocumentPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PostDocumentOptions model
				postDocumentOptionsModel := new(cloudantv1.PostDocumentOptions)
				postDocumentOptionsModel.Db = core.StringPtr("testString")
				postDocumentOptionsModel.Document = documentModel
				postDocumentOptionsModel.ContentType = core.StringPtr("application/json")
				postDocumentOptionsModel.Batch = core.StringPtr("ok")
				postDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostDocument(postDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostDocument(postDocumentOptions *PostDocumentOptions)`, func() {
		postDocumentPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke PostDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PostDocumentOptions model
				postDocumentOptionsModel := new(cloudantv1.PostDocumentOptions)
				postDocumentOptionsModel.Db = core.StringPtr("testString")
				postDocumentOptionsModel.Document = documentModel
				postDocumentOptionsModel.ContentType = core.StringPtr("application/json")
				postDocumentOptionsModel.Batch = core.StringPtr("ok")
				postDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostDocument(postDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PostDocumentOptions model
				postDocumentOptionsModel := new(cloudantv1.PostDocumentOptions)
				postDocumentOptionsModel.Db = core.StringPtr("testString")
				postDocumentOptionsModel.Document = documentModel
				postDocumentOptionsModel.ContentType = core.StringPtr("application/json")
				postDocumentOptionsModel.Batch = core.StringPtr("ok")
				postDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostDocument(postDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostDocumentOptions model with no property values
				postDocumentOptionsModelNew := new(cloudantv1.PostDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostDocument(postDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostAllDocs(postAllDocsOptions *PostAllDocsOptions) - Operation response error`, func() {
		postAllDocsPath := "/testString/_all_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postAllDocsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostAllDocs with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostAllDocsOptions model
				postAllDocsOptionsModel := new(cloudantv1.PostAllDocsOptions)
				postAllDocsOptionsModel.Db = core.StringPtr("testString")
				postAllDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postAllDocsOptionsModel.Attachments = core.BoolPtr(true)
				postAllDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postAllDocsOptionsModel.Descending = core.BoolPtr(true)
				postAllDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postAllDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postAllDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postAllDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postAllDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postAllDocsOptionsModel.Endkey = core.StringPtr("testString")
				postAllDocsOptionsModel.Key = core.StringPtr("testString")
				postAllDocsOptionsModel.Keys = []string{"testString"}
				postAllDocsOptionsModel.Startkey = core.StringPtr("testString")
				postAllDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostAllDocs(postAllDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostAllDocs(postAllDocsOptions *PostAllDocsOptions)`, func() {
		postAllDocsPath := "/testString/_all_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postAllDocsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "rows": [{"caused_by": "CausedBy", "error": "Error", "reason": "Reason", "doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "id": "ID", "key": "Key", "value": {"rev": "Rev"}}], "update_seq": "UpdateSeq"}`)
				}))
			})
			It(`Invoke PostAllDocs successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostAllDocs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostAllDocsOptions model
				postAllDocsOptionsModel := new(cloudantv1.PostAllDocsOptions)
				postAllDocsOptionsModel.Db = core.StringPtr("testString")
				postAllDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postAllDocsOptionsModel.Attachments = core.BoolPtr(true)
				postAllDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postAllDocsOptionsModel.Descending = core.BoolPtr(true)
				postAllDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postAllDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postAllDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postAllDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postAllDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postAllDocsOptionsModel.Endkey = core.StringPtr("testString")
				postAllDocsOptionsModel.Key = core.StringPtr("testString")
				postAllDocsOptionsModel.Keys = []string{"testString"}
				postAllDocsOptionsModel.Startkey = core.StringPtr("testString")
				postAllDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostAllDocs(postAllDocsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostAllDocs with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostAllDocsOptions model
				postAllDocsOptionsModel := new(cloudantv1.PostAllDocsOptions)
				postAllDocsOptionsModel.Db = core.StringPtr("testString")
				postAllDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postAllDocsOptionsModel.Attachments = core.BoolPtr(true)
				postAllDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postAllDocsOptionsModel.Descending = core.BoolPtr(true)
				postAllDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postAllDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postAllDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postAllDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postAllDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postAllDocsOptionsModel.Endkey = core.StringPtr("testString")
				postAllDocsOptionsModel.Key = core.StringPtr("testString")
				postAllDocsOptionsModel.Keys = []string{"testString"}
				postAllDocsOptionsModel.Startkey = core.StringPtr("testString")
				postAllDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostAllDocs(postAllDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostAllDocsOptions model with no property values
				postAllDocsOptionsModelNew := new(cloudantv1.PostAllDocsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostAllDocs(postAllDocsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostAllDocsAsStream(postAllDocsOptions *PostAllDocsOptions)`, func() {
		postAllDocsAsStreamPath := "/testString/_all_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postAllDocsAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostAllDocsAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostAllDocsAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostAllDocsOptions model
				postAllDocsOptionsModel := new(cloudantv1.PostAllDocsOptions)
				postAllDocsOptionsModel.Db = core.StringPtr("testString")
				postAllDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postAllDocsOptionsModel.Attachments = core.BoolPtr(true)
				postAllDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postAllDocsOptionsModel.Descending = core.BoolPtr(true)
				postAllDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postAllDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postAllDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postAllDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postAllDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postAllDocsOptionsModel.Endkey = core.StringPtr("testString")
				postAllDocsOptionsModel.Key = core.StringPtr("testString")
				postAllDocsOptionsModel.Keys = []string{"testString"}
				postAllDocsOptionsModel.Startkey = core.StringPtr("testString")
				postAllDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostAllDocsAsStream(postAllDocsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostAllDocsAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostAllDocsOptions model
				postAllDocsOptionsModel := new(cloudantv1.PostAllDocsOptions)
				postAllDocsOptionsModel.Db = core.StringPtr("testString")
				postAllDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postAllDocsOptionsModel.Attachments = core.BoolPtr(true)
				postAllDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postAllDocsOptionsModel.Descending = core.BoolPtr(true)
				postAllDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postAllDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postAllDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postAllDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postAllDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postAllDocsOptionsModel.Endkey = core.StringPtr("testString")
				postAllDocsOptionsModel.Key = core.StringPtr("testString")
				postAllDocsOptionsModel.Keys = []string{"testString"}
				postAllDocsOptionsModel.Startkey = core.StringPtr("testString")
				postAllDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostAllDocsAsStream(postAllDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostAllDocsOptions model with no property values
				postAllDocsOptionsModelNew := new(cloudantv1.PostAllDocsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostAllDocsAsStream(postAllDocsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostAllDocsQueries(postAllDocsQueriesOptions *PostAllDocsQueriesOptions) - Operation response error`, func() {
		postAllDocsQueriesPath := "/testString/_all_docs/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postAllDocsQueriesPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostAllDocsQueries with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostAllDocsQueriesOptions model
				postAllDocsQueriesOptionsModel := new(cloudantv1.PostAllDocsQueriesOptions)
				postAllDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postAllDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postAllDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostAllDocsQueries(postAllDocsQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostAllDocsQueries(postAllDocsQueriesOptions *PostAllDocsQueriesOptions)`, func() {
		postAllDocsQueriesPath := "/testString/_all_docs/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postAllDocsQueriesPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"total_rows": 0, "rows": [{"caused_by": "CausedBy", "error": "Error", "reason": "Reason", "doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "id": "ID", "key": "Key", "value": {"rev": "Rev"}}], "update_seq": "UpdateSeq"}]}`)
				}))
			})
			It(`Invoke PostAllDocsQueries successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostAllDocsQueries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostAllDocsQueriesOptions model
				postAllDocsQueriesOptionsModel := new(cloudantv1.PostAllDocsQueriesOptions)
				postAllDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postAllDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postAllDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostAllDocsQueries(postAllDocsQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostAllDocsQueries with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostAllDocsQueriesOptions model
				postAllDocsQueriesOptionsModel := new(cloudantv1.PostAllDocsQueriesOptions)
				postAllDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postAllDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postAllDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostAllDocsQueries(postAllDocsQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostAllDocsQueriesOptions model with no property values
				postAllDocsQueriesOptionsModelNew := new(cloudantv1.PostAllDocsQueriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostAllDocsQueries(postAllDocsQueriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostAllDocsQueriesAsStream(postAllDocsQueriesOptions *PostAllDocsQueriesOptions)`, func() {
		postAllDocsQueriesAsStreamPath := "/testString/_all_docs/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postAllDocsQueriesAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostAllDocsQueriesAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostAllDocsQueriesAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostAllDocsQueriesOptions model
				postAllDocsQueriesOptionsModel := new(cloudantv1.PostAllDocsQueriesOptions)
				postAllDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postAllDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postAllDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostAllDocsQueriesAsStream(postAllDocsQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostAllDocsQueriesAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostAllDocsQueriesOptions model
				postAllDocsQueriesOptionsModel := new(cloudantv1.PostAllDocsQueriesOptions)
				postAllDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postAllDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postAllDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostAllDocsQueriesAsStream(postAllDocsQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostAllDocsQueriesOptions model with no property values
				postAllDocsQueriesOptionsModelNew := new(cloudantv1.PostAllDocsQueriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostAllDocsQueriesAsStream(postAllDocsQueriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostBulkDocs(postBulkDocsOptions *PostBulkDocsOptions) - Operation response error`, func() {
		postBulkDocsPath := "/testString/_bulk_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postBulkDocsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostBulkDocs with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the BulkDocs model
				bulkDocsModel := new(cloudantv1.BulkDocs)
				bulkDocsModel.Docs = []cloudantv1.Document{*documentModel}
				bulkDocsModel.NewEdits = core.BoolPtr(true)

				// Construct an instance of the PostBulkDocsOptions model
				postBulkDocsOptionsModel := new(cloudantv1.PostBulkDocsOptions)
				postBulkDocsOptionsModel.Db = core.StringPtr("testString")
				postBulkDocsOptionsModel.BulkDocs = bulkDocsModel
				postBulkDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostBulkDocs(postBulkDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostBulkDocs(postBulkDocsOptions *PostBulkDocsOptions)`, func() {
		postBulkDocsPath := "/testString/_bulk_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postBulkDocsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `[{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}]`)
				}))
			})
			It(`Invoke PostBulkDocs successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostBulkDocs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the BulkDocs model
				bulkDocsModel := new(cloudantv1.BulkDocs)
				bulkDocsModel.Docs = []cloudantv1.Document{*documentModel}
				bulkDocsModel.NewEdits = core.BoolPtr(true)

				// Construct an instance of the PostBulkDocsOptions model
				postBulkDocsOptionsModel := new(cloudantv1.PostBulkDocsOptions)
				postBulkDocsOptionsModel.Db = core.StringPtr("testString")
				postBulkDocsOptionsModel.BulkDocs = bulkDocsModel
				postBulkDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostBulkDocs(postBulkDocsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostBulkDocs with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the BulkDocs model
				bulkDocsModel := new(cloudantv1.BulkDocs)
				bulkDocsModel.Docs = []cloudantv1.Document{*documentModel}
				bulkDocsModel.NewEdits = core.BoolPtr(true)

				// Construct an instance of the PostBulkDocsOptions model
				postBulkDocsOptionsModel := new(cloudantv1.PostBulkDocsOptions)
				postBulkDocsOptionsModel.Db = core.StringPtr("testString")
				postBulkDocsOptionsModel.BulkDocs = bulkDocsModel
				postBulkDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostBulkDocs(postBulkDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostBulkDocsOptions model with no property values
				postBulkDocsOptionsModelNew := new(cloudantv1.PostBulkDocsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostBulkDocs(postBulkDocsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostBulkGet(postBulkGetOptions *PostBulkGetOptions) - Operation response error`, func() {
		postBulkGetPath := "/testString/_bulk_get"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postBulkGetPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for revs query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostBulkGet with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the BulkGetQueryDocument model
				bulkGetQueryDocumentModel := new(cloudantv1.BulkGetQueryDocument)
				bulkGetQueryDocumentModel.AttsSince = []string{"testString"}
				bulkGetQueryDocumentModel.ID = core.StringPtr("foo")
				bulkGetQueryDocumentModel.OpenRevs = []string{"testString"}
				bulkGetQueryDocumentModel.Rev = core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")

				// Construct an instance of the PostBulkGetOptions model
				postBulkGetOptionsModel := new(cloudantv1.PostBulkGetOptions)
				postBulkGetOptionsModel.Db = core.StringPtr("testString")
				postBulkGetOptionsModel.Docs = []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel}
				postBulkGetOptionsModel.Attachments = core.BoolPtr(true)
				postBulkGetOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postBulkGetOptionsModel.Latest = core.BoolPtr(true)
				postBulkGetOptionsModel.Revs = core.BoolPtr(true)
				postBulkGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostBulkGet(postBulkGetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostBulkGet(postBulkGetOptions *PostBulkGetOptions)`, func() {
		postBulkGetPath := "/testString/_bulk_get"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postBulkGetPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for revs query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"docs": [{"error": {"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}, "ok": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}}], "id": "ID"}]}`)
				}))
			})
			It(`Invoke PostBulkGet successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostBulkGet(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BulkGetQueryDocument model
				bulkGetQueryDocumentModel := new(cloudantv1.BulkGetQueryDocument)
				bulkGetQueryDocumentModel.AttsSince = []string{"testString"}
				bulkGetQueryDocumentModel.ID = core.StringPtr("foo")
				bulkGetQueryDocumentModel.OpenRevs = []string{"testString"}
				bulkGetQueryDocumentModel.Rev = core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")

				// Construct an instance of the PostBulkGetOptions model
				postBulkGetOptionsModel := new(cloudantv1.PostBulkGetOptions)
				postBulkGetOptionsModel.Db = core.StringPtr("testString")
				postBulkGetOptionsModel.Docs = []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel}
				postBulkGetOptionsModel.Attachments = core.BoolPtr(true)
				postBulkGetOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postBulkGetOptionsModel.Latest = core.BoolPtr(true)
				postBulkGetOptionsModel.Revs = core.BoolPtr(true)
				postBulkGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostBulkGet(postBulkGetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostBulkGet with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the BulkGetQueryDocument model
				bulkGetQueryDocumentModel := new(cloudantv1.BulkGetQueryDocument)
				bulkGetQueryDocumentModel.AttsSince = []string{"testString"}
				bulkGetQueryDocumentModel.ID = core.StringPtr("foo")
				bulkGetQueryDocumentModel.OpenRevs = []string{"testString"}
				bulkGetQueryDocumentModel.Rev = core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")

				// Construct an instance of the PostBulkGetOptions model
				postBulkGetOptionsModel := new(cloudantv1.PostBulkGetOptions)
				postBulkGetOptionsModel.Db = core.StringPtr("testString")
				postBulkGetOptionsModel.Docs = []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel}
				postBulkGetOptionsModel.Attachments = core.BoolPtr(true)
				postBulkGetOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postBulkGetOptionsModel.Latest = core.BoolPtr(true)
				postBulkGetOptionsModel.Revs = core.BoolPtr(true)
				postBulkGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostBulkGet(postBulkGetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostBulkGetOptions model with no property values
				postBulkGetOptionsModelNew := new(cloudantv1.PostBulkGetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostBulkGet(postBulkGetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostBulkGetAsMixed(postBulkGetOptions *PostBulkGetOptions)`, func() {
		postBulkGetAsMixedPath := "/testString/_bulk_get"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postBulkGetAsMixedPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for revs query parameter

					res.Header().Set("Content-type", "multipart/mixed")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke PostBulkGetAsMixed successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostBulkGetAsMixed(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BulkGetQueryDocument model
				bulkGetQueryDocumentModel := new(cloudantv1.BulkGetQueryDocument)
				bulkGetQueryDocumentModel.AttsSince = []string{"testString"}
				bulkGetQueryDocumentModel.ID = core.StringPtr("foo")
				bulkGetQueryDocumentModel.OpenRevs = []string{"testString"}
				bulkGetQueryDocumentModel.Rev = core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")

				// Construct an instance of the PostBulkGetOptions model
				postBulkGetOptionsModel := new(cloudantv1.PostBulkGetOptions)
				postBulkGetOptionsModel.Db = core.StringPtr("testString")
				postBulkGetOptionsModel.Docs = []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel}
				postBulkGetOptionsModel.Attachments = core.BoolPtr(true)
				postBulkGetOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postBulkGetOptionsModel.Latest = core.BoolPtr(true)
				postBulkGetOptionsModel.Revs = core.BoolPtr(true)
				postBulkGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostBulkGetAsMixed(postBulkGetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostBulkGetAsMixed with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the BulkGetQueryDocument model
				bulkGetQueryDocumentModel := new(cloudantv1.BulkGetQueryDocument)
				bulkGetQueryDocumentModel.AttsSince = []string{"testString"}
				bulkGetQueryDocumentModel.ID = core.StringPtr("foo")
				bulkGetQueryDocumentModel.OpenRevs = []string{"testString"}
				bulkGetQueryDocumentModel.Rev = core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")

				// Construct an instance of the PostBulkGetOptions model
				postBulkGetOptionsModel := new(cloudantv1.PostBulkGetOptions)
				postBulkGetOptionsModel.Db = core.StringPtr("testString")
				postBulkGetOptionsModel.Docs = []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel}
				postBulkGetOptionsModel.Attachments = core.BoolPtr(true)
				postBulkGetOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postBulkGetOptionsModel.Latest = core.BoolPtr(true)
				postBulkGetOptionsModel.Revs = core.BoolPtr(true)
				postBulkGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostBulkGetAsMixed(postBulkGetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostBulkGetOptions model with no property values
				postBulkGetOptionsModelNew := new(cloudantv1.PostBulkGetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostBulkGetAsMixed(postBulkGetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostBulkGetAsRelated(postBulkGetOptions *PostBulkGetOptions)`, func() {
		postBulkGetAsRelatedPath := "/testString/_bulk_get"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postBulkGetAsRelatedPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for revs query parameter

					res.Header().Set("Content-type", "multipart/related")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke PostBulkGetAsRelated successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostBulkGetAsRelated(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BulkGetQueryDocument model
				bulkGetQueryDocumentModel := new(cloudantv1.BulkGetQueryDocument)
				bulkGetQueryDocumentModel.AttsSince = []string{"testString"}
				bulkGetQueryDocumentModel.ID = core.StringPtr("foo")
				bulkGetQueryDocumentModel.OpenRevs = []string{"testString"}
				bulkGetQueryDocumentModel.Rev = core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")

				// Construct an instance of the PostBulkGetOptions model
				postBulkGetOptionsModel := new(cloudantv1.PostBulkGetOptions)
				postBulkGetOptionsModel.Db = core.StringPtr("testString")
				postBulkGetOptionsModel.Docs = []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel}
				postBulkGetOptionsModel.Attachments = core.BoolPtr(true)
				postBulkGetOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postBulkGetOptionsModel.Latest = core.BoolPtr(true)
				postBulkGetOptionsModel.Revs = core.BoolPtr(true)
				postBulkGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostBulkGetAsRelated(postBulkGetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostBulkGetAsRelated with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the BulkGetQueryDocument model
				bulkGetQueryDocumentModel := new(cloudantv1.BulkGetQueryDocument)
				bulkGetQueryDocumentModel.AttsSince = []string{"testString"}
				bulkGetQueryDocumentModel.ID = core.StringPtr("foo")
				bulkGetQueryDocumentModel.OpenRevs = []string{"testString"}
				bulkGetQueryDocumentModel.Rev = core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")

				// Construct an instance of the PostBulkGetOptions model
				postBulkGetOptionsModel := new(cloudantv1.PostBulkGetOptions)
				postBulkGetOptionsModel.Db = core.StringPtr("testString")
				postBulkGetOptionsModel.Docs = []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel}
				postBulkGetOptionsModel.Attachments = core.BoolPtr(true)
				postBulkGetOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postBulkGetOptionsModel.Latest = core.BoolPtr(true)
				postBulkGetOptionsModel.Revs = core.BoolPtr(true)
				postBulkGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostBulkGetAsRelated(postBulkGetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostBulkGetOptions model with no property values
				postBulkGetOptionsModelNew := new(cloudantv1.PostBulkGetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostBulkGetAsRelated(postBulkGetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostBulkGetAsStream(postBulkGetOptions *PostBulkGetOptions)`, func() {
		postBulkGetAsStreamPath := "/testString/_bulk_get"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postBulkGetAsStreamPath))
					Expect(req.Method).To(Equal("POST"))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for revs query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostBulkGetAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostBulkGetAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BulkGetQueryDocument model
				bulkGetQueryDocumentModel := new(cloudantv1.BulkGetQueryDocument)
				bulkGetQueryDocumentModel.AttsSince = []string{"testString"}
				bulkGetQueryDocumentModel.ID = core.StringPtr("foo")
				bulkGetQueryDocumentModel.OpenRevs = []string{"testString"}
				bulkGetQueryDocumentModel.Rev = core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")

				// Construct an instance of the PostBulkGetOptions model
				postBulkGetOptionsModel := new(cloudantv1.PostBulkGetOptions)
				postBulkGetOptionsModel.Db = core.StringPtr("testString")
				postBulkGetOptionsModel.Docs = []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel}
				postBulkGetOptionsModel.Attachments = core.BoolPtr(true)
				postBulkGetOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postBulkGetOptionsModel.Latest = core.BoolPtr(true)
				postBulkGetOptionsModel.Revs = core.BoolPtr(true)
				postBulkGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostBulkGetAsStream(postBulkGetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostBulkGetAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the BulkGetQueryDocument model
				bulkGetQueryDocumentModel := new(cloudantv1.BulkGetQueryDocument)
				bulkGetQueryDocumentModel.AttsSince = []string{"testString"}
				bulkGetQueryDocumentModel.ID = core.StringPtr("foo")
				bulkGetQueryDocumentModel.OpenRevs = []string{"testString"}
				bulkGetQueryDocumentModel.Rev = core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")

				// Construct an instance of the PostBulkGetOptions model
				postBulkGetOptionsModel := new(cloudantv1.PostBulkGetOptions)
				postBulkGetOptionsModel.Db = core.StringPtr("testString")
				postBulkGetOptionsModel.Docs = []cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel}
				postBulkGetOptionsModel.Attachments = core.BoolPtr(true)
				postBulkGetOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postBulkGetOptionsModel.Latest = core.BoolPtr(true)
				postBulkGetOptionsModel.Revs = core.BoolPtr(true)
				postBulkGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostBulkGetAsStream(postBulkGetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostBulkGetOptions model with no property values
				postBulkGetOptionsModelNew := new(cloudantv1.PostBulkGetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostBulkGetAsStream(postBulkGetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions) - Operation response error`, func() {
		deleteDocumentPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(cloudantv1.DeleteDocumentOptions)
				deleteDocumentOptionsModel.Db = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocID = core.StringPtr("testString")
				deleteDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteDocumentOptionsModel.Rev = core.StringPtr("testString")
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteDocument(deleteDocumentOptions *DeleteDocumentOptions)`, func() {
		deleteDocumentPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke DeleteDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.DeleteDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(cloudantv1.DeleteDocumentOptions)
				deleteDocumentOptionsModel.Db = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocID = core.StringPtr("testString")
				deleteDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteDocumentOptionsModel.Rev = core.StringPtr("testString")
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteDocumentOptions model
				deleteDocumentOptionsModel := new(cloudantv1.DeleteDocumentOptions)
				deleteDocumentOptionsModel.Db = core.StringPtr("testString")
				deleteDocumentOptionsModel.DocID = core.StringPtr("testString")
				deleteDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteDocumentOptionsModel.Rev = core.StringPtr("testString")
				deleteDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.DeleteDocument(deleteDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDocumentOptions model with no property values
				deleteDocumentOptionsModelNew := new(cloudantv1.DeleteDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.DeleteDocument(deleteDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDocument(getDocumentOptions *GetDocumentOptions) - Operation response error`, func() {
		getDocumentPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for deleted_conflicts query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for local_seq query parameter


					// TODO: Add check for meta query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))


					// TODO: Add check for revs query parameter


					// TODO: Add check for revs_info query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDocumentOptions model
				getDocumentOptionsModel := new(cloudantv1.GetDocumentOptions)
				getDocumentOptionsModel.Db = core.StringPtr("testString")
				getDocumentOptionsModel.DocID = core.StringPtr("testString")
				getDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDocumentOptionsModel.AttsSince = []string{"testString"}
				getDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetDocument(getDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocument(getDocumentOptions *GetDocumentOptions)`, func() {
		getDocumentPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for deleted_conflicts query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for local_seq query parameter


					// TODO: Add check for meta query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))


					// TODO: Add check for revs query parameter


					// TODO: Add check for revs_info query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}`)
				}))
			})
			It(`Invoke GetDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDocumentOptions model
				getDocumentOptionsModel := new(cloudantv1.GetDocumentOptions)
				getDocumentOptionsModel.Db = core.StringPtr("testString")
				getDocumentOptionsModel.DocID = core.StringPtr("testString")
				getDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDocumentOptionsModel.AttsSince = []string{"testString"}
				getDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetDocument(getDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDocumentOptions model
				getDocumentOptionsModel := new(cloudantv1.GetDocumentOptions)
				getDocumentOptionsModel.Db = core.StringPtr("testString")
				getDocumentOptionsModel.DocID = core.StringPtr("testString")
				getDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDocumentOptionsModel.AttsSince = []string{"testString"}
				getDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetDocument(getDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentOptions model with no property values
				getDocumentOptionsModelNew := new(cloudantv1.GetDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetDocument(getDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocumentAsMixed(getDocumentOptions *GetDocumentOptions)`, func() {
		getDocumentAsMixedPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentAsMixedPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for deleted_conflicts query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for local_seq query parameter


					// TODO: Add check for meta query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))


					// TODO: Add check for revs query parameter


					// TODO: Add check for revs_info query parameter

					res.Header().Set("Content-type", "multipart/mixed")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetDocumentAsMixed successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetDocumentAsMixed(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDocumentOptions model
				getDocumentOptionsModel := new(cloudantv1.GetDocumentOptions)
				getDocumentOptionsModel.Db = core.StringPtr("testString")
				getDocumentOptionsModel.DocID = core.StringPtr("testString")
				getDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDocumentOptionsModel.AttsSince = []string{"testString"}
				getDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetDocumentAsMixed(getDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDocumentAsMixed with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDocumentOptions model
				getDocumentOptionsModel := new(cloudantv1.GetDocumentOptions)
				getDocumentOptionsModel.Db = core.StringPtr("testString")
				getDocumentOptionsModel.DocID = core.StringPtr("testString")
				getDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDocumentOptionsModel.AttsSince = []string{"testString"}
				getDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetDocumentAsMixed(getDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentOptions model with no property values
				getDocumentOptionsModelNew := new(cloudantv1.GetDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetDocumentAsMixed(getDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocumentAsRelated(getDocumentOptions *GetDocumentOptions)`, func() {
		getDocumentAsRelatedPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentAsRelatedPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for deleted_conflicts query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for local_seq query parameter


					// TODO: Add check for meta query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))


					// TODO: Add check for revs query parameter


					// TODO: Add check for revs_info query parameter

					res.Header().Set("Content-type", "multipart/related")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetDocumentAsRelated successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetDocumentAsRelated(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDocumentOptions model
				getDocumentOptionsModel := new(cloudantv1.GetDocumentOptions)
				getDocumentOptionsModel.Db = core.StringPtr("testString")
				getDocumentOptionsModel.DocID = core.StringPtr("testString")
				getDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDocumentOptionsModel.AttsSince = []string{"testString"}
				getDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetDocumentAsRelated(getDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDocumentAsRelated with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDocumentOptions model
				getDocumentOptionsModel := new(cloudantv1.GetDocumentOptions)
				getDocumentOptionsModel.Db = core.StringPtr("testString")
				getDocumentOptionsModel.DocID = core.StringPtr("testString")
				getDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDocumentOptionsModel.AttsSince = []string{"testString"}
				getDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetDocumentAsRelated(getDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentOptions model with no property values
				getDocumentOptionsModelNew := new(cloudantv1.GetDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetDocumentAsRelated(getDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocumentAsStream(getDocumentOptions *GetDocumentOptions)`, func() {
		getDocumentAsStreamPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentAsStreamPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for deleted_conflicts query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for local_seq query parameter


					// TODO: Add check for meta query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))


					// TODO: Add check for revs query parameter


					// TODO: Add check for revs_info query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke GetDocumentAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetDocumentAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDocumentOptions model
				getDocumentOptionsModel := new(cloudantv1.GetDocumentOptions)
				getDocumentOptionsModel.Db = core.StringPtr("testString")
				getDocumentOptionsModel.DocID = core.StringPtr("testString")
				getDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDocumentOptionsModel.AttsSince = []string{"testString"}
				getDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetDocumentAsStream(getDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke GetDocumentAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDocumentOptions model
				getDocumentOptionsModel := new(cloudantv1.GetDocumentOptions)
				getDocumentOptionsModel.Db = core.StringPtr("testString")
				getDocumentOptionsModel.DocID = core.StringPtr("testString")
				getDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDocumentOptionsModel.AttsSince = []string{"testString"}
				getDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetDocumentAsStream(getDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentOptions model with no property values
				getDocumentOptionsModelNew := new(cloudantv1.GetDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetDocumentAsStream(getDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutDocument(putDocumentOptions *PutDocumentOptions) - Operation response error`, func() {
		putDocumentPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDocumentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))


					// TODO: Add check for new_edits query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PutDocumentOptions model
				putDocumentOptionsModel := new(cloudantv1.PutDocumentOptions)
				putDocumentOptionsModel.Db = core.StringPtr("testString")
				putDocumentOptionsModel.DocID = core.StringPtr("testString")
				putDocumentOptionsModel.Document = documentModel
				putDocumentOptionsModel.ContentType = core.StringPtr("application/json")
				putDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				putDocumentOptionsModel.Batch = core.StringPtr("ok")
				putDocumentOptionsModel.NewEdits = core.BoolPtr(true)
				putDocumentOptionsModel.Rev = core.StringPtr("testString")
				putDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PutDocument(putDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutDocument(putDocumentOptions *PutDocumentOptions)`, func() {
		putDocumentPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDocumentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))


					// TODO: Add check for new_edits query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke PutDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PutDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PutDocumentOptions model
				putDocumentOptionsModel := new(cloudantv1.PutDocumentOptions)
				putDocumentOptionsModel.Db = core.StringPtr("testString")
				putDocumentOptionsModel.DocID = core.StringPtr("testString")
				putDocumentOptionsModel.Document = documentModel
				putDocumentOptionsModel.ContentType = core.StringPtr("application/json")
				putDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				putDocumentOptionsModel.Batch = core.StringPtr("ok")
				putDocumentOptionsModel.NewEdits = core.BoolPtr(true)
				putDocumentOptionsModel.Rev = core.StringPtr("testString")
				putDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PutDocument(putDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PutDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PutDocumentOptions model
				putDocumentOptionsModel := new(cloudantv1.PutDocumentOptions)
				putDocumentOptionsModel.Db = core.StringPtr("testString")
				putDocumentOptionsModel.DocID = core.StringPtr("testString")
				putDocumentOptionsModel.Document = documentModel
				putDocumentOptionsModel.ContentType = core.StringPtr("application/json")
				putDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				putDocumentOptionsModel.Batch = core.StringPtr("ok")
				putDocumentOptionsModel.NewEdits = core.BoolPtr(true)
				putDocumentOptionsModel.Rev = core.StringPtr("testString")
				putDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PutDocument(putDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutDocumentOptions model with no property values
				putDocumentOptionsModelNew := new(cloudantv1.PutDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PutDocument(putDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`HeadDesignDocument(headDesignDocumentOptions *HeadDesignDocumentOptions)`, func() {
		headDesignDocumentPath := "/testString/_design/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(headDesignDocumentPath))
					Expect(req.Method).To(Equal("HEAD"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke HeadDesignDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cloudantService.HeadDesignDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HeadDesignDocumentOptions model
				headDesignDocumentOptionsModel := new(cloudantv1.HeadDesignDocumentOptions)
				headDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				headDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				headDesignDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				headDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cloudantService.HeadDesignDocument(headDesignDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HeadDesignDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the HeadDesignDocumentOptions model
				headDesignDocumentOptionsModel := new(cloudantv1.HeadDesignDocumentOptions)
				headDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				headDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				headDesignDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				headDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cloudantService.HeadDesignDocument(headDesignDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HeadDesignDocumentOptions model with no property values
				headDesignDocumentOptionsModelNew := new(cloudantv1.HeadDesignDocumentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cloudantService.HeadDesignDocument(headDesignDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDesignDocument(deleteDesignDocumentOptions *DeleteDesignDocumentOptions) - Operation response error`, func() {
		deleteDesignDocumentPath := "/testString/_design/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDesignDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDesignDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteDesignDocumentOptions model
				deleteDesignDocumentOptionsModel := new(cloudantv1.DeleteDesignDocumentOptions)
				deleteDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteDesignDocumentOptionsModel.Rev = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.DeleteDesignDocument(deleteDesignDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteDesignDocument(deleteDesignDocumentOptions *DeleteDesignDocumentOptions)`, func() {
		deleteDesignDocumentPath := "/testString/_design/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDesignDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke DeleteDesignDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.DeleteDesignDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDesignDocumentOptions model
				deleteDesignDocumentOptionsModel := new(cloudantv1.DeleteDesignDocumentOptions)
				deleteDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteDesignDocumentOptionsModel.Rev = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.DeleteDesignDocument(deleteDesignDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteDesignDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteDesignDocumentOptions model
				deleteDesignDocumentOptionsModel := new(cloudantv1.DeleteDesignDocumentOptions)
				deleteDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteDesignDocumentOptionsModel.Rev = core.StringPtr("testString")
				deleteDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.DeleteDesignDocument(deleteDesignDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDesignDocumentOptions model with no property values
				deleteDesignDocumentOptionsModelNew := new(cloudantv1.DeleteDesignDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.DeleteDesignDocument(deleteDesignDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDesignDocument(getDesignDocumentOptions *GetDesignDocumentOptions) - Operation response error`, func() {
		getDesignDocumentPath := "/testString/_design/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDesignDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for deleted_conflicts query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for local_seq query parameter


					// TODO: Add check for meta query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))


					// TODO: Add check for revs query parameter


					// TODO: Add check for revs_info query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDesignDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDesignDocumentOptions model
				getDesignDocumentOptionsModel := new(cloudantv1.GetDesignDocumentOptions)
				getDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				getDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				getDesignDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDesignDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDesignDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDesignDocumentOptionsModel.AttsSince = []string{"testString"}
				getDesignDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDesignDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDesignDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDesignDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDesignDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDesignDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDesignDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDesignDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDesignDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetDesignDocument(getDesignDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDesignDocument(getDesignDocumentOptions *GetDesignDocumentOptions)`, func() {
		getDesignDocumentPath := "/testString/_design/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDesignDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for deleted_conflicts query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for local_seq query parameter


					// TODO: Add check for meta query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))


					// TODO: Add check for revs query parameter


					// TODO: Add check for revs_info query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}], "autoupdate": true, "filters": {"mapKey": "Inner"}, "indexes": {"mapKey": {"analyzer": {"name": "classic", "stopwords": ["Stopwords"], "fields": {"mapKey": {"name": "classic", "stopwords": ["Stopwords"]}}}, "index": "Index"}}, "language": "Language", "options": {"partitioned": false}, "updates": {"mapKey": "Inner"}, "validate_doc_update": {"mapKey": "Inner"}, "views": {"mapKey": {"map": "Map", "reduce": "Reduce"}}, "st_indexes": {"mapKey": {"index": "Index"}}}`)
				}))
			})
			It(`Invoke GetDesignDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetDesignDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDesignDocumentOptions model
				getDesignDocumentOptionsModel := new(cloudantv1.GetDesignDocumentOptions)
				getDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				getDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				getDesignDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDesignDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDesignDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDesignDocumentOptionsModel.AttsSince = []string{"testString"}
				getDesignDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDesignDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDesignDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDesignDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDesignDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDesignDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDesignDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDesignDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDesignDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetDesignDocument(getDesignDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDesignDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDesignDocumentOptions model
				getDesignDocumentOptionsModel := new(cloudantv1.GetDesignDocumentOptions)
				getDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				getDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				getDesignDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getDesignDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getDesignDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getDesignDocumentOptionsModel.AttsSince = []string{"testString"}
				getDesignDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getDesignDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getDesignDocumentOptionsModel.Latest = core.BoolPtr(true)
				getDesignDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getDesignDocumentOptionsModel.Meta = core.BoolPtr(true)
				getDesignDocumentOptionsModel.OpenRevs = []string{"testString"}
				getDesignDocumentOptionsModel.Rev = core.StringPtr("testString")
				getDesignDocumentOptionsModel.Revs = core.BoolPtr(true)
				getDesignDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetDesignDocument(getDesignDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDesignDocumentOptions model with no property values
				getDesignDocumentOptionsModelNew := new(cloudantv1.GetDesignDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetDesignDocument(getDesignDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutDesignDocument(putDesignDocumentOptions *PutDesignDocumentOptions) - Operation response error`, func() {
		putDesignDocumentPath := "/testString/_design/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDesignDocumentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))


					// TODO: Add check for new_edits query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutDesignDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Analyzer model
				analyzerModel := new(cloudantv1.Analyzer)
				analyzerModel.Name = core.StringPtr("classic")
				analyzerModel.Stopwords = []string{"testString"}

				// Construct an instance of the AnalyzerConfiguration model
				analyzerConfigurationModel := new(cloudantv1.AnalyzerConfiguration)
				analyzerConfigurationModel.Name = core.StringPtr("classic")
				analyzerConfigurationModel.Stopwords = []string{"testString"}
				analyzerConfigurationModel.Fields = make(map[string]cloudantv1.Analyzer)
				analyzerConfigurationModel.Fields["foo"] = *analyzerModel

				// Construct an instance of the SearchIndexDefinition model
				searchIndexDefinitionModel := new(cloudantv1.SearchIndexDefinition)
				searchIndexDefinitionModel.Analyzer = analyzerConfigurationModel
				searchIndexDefinitionModel.Index = core.StringPtr("testString")

				// Construct an instance of the DesignDocumentOptions model
				designDocumentOptionsModel := new(cloudantv1.DesignDocumentOptions)
				designDocumentOptionsModel.Partitioned = core.BoolPtr(true)

				// Construct an instance of the DesignDocumentViewsMapReduce model
				designDocumentViewsMapReduceModel := new(cloudantv1.DesignDocumentViewsMapReduce)
				designDocumentViewsMapReduceModel.Map = core.StringPtr("testString")
				designDocumentViewsMapReduceModel.Reduce = core.StringPtr("testString")

				// Construct an instance of the GeoIndexDefinition model
				geoIndexDefinitionModel := new(cloudantv1.GeoIndexDefinition)
				geoIndexDefinitionModel.Index = core.StringPtr("testString")

				// Construct an instance of the DesignDocument model
				designDocumentModel := new(cloudantv1.DesignDocument)
				designDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				designDocumentModel.Conflicts = []string{"testString"}
				designDocumentModel.Deleted = core.BoolPtr(true)
				designDocumentModel.DeletedConflicts = []string{"testString"}
				designDocumentModel.ID = core.StringPtr("testString")
				designDocumentModel.LocalSeq = core.StringPtr("testString")
				designDocumentModel.Rev = core.StringPtr("testString")
				designDocumentModel.Revisions = revisionsModel
				designDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				designDocumentModel.Autoupdate = core.BoolPtr(true)
				designDocumentModel.Filters = make(map[string]string)
				designDocumentModel.Indexes = make(map[string]cloudantv1.SearchIndexDefinition)
				designDocumentModel.Language = core.StringPtr("testString")
				designDocumentModel.Options = designDocumentOptionsModel
				designDocumentModel.Updates = make(map[string]string)
				designDocumentModel.ValidateDocUpdate = make(map[string]string)
				designDocumentModel.Views = make(map[string]cloudantv1.DesignDocumentViewsMapReduce)
				designDocumentModel.StIndexes = make(map[string]cloudantv1.GeoIndexDefinition)
				designDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				designDocumentModel.Attachments["foo"] = *attachmentModel
				designDocumentModel.Indexes["foo"] = *searchIndexDefinitionModel
				designDocumentModel.Views["foo"] = *designDocumentViewsMapReduceModel
				designDocumentModel.StIndexes["foo"] = *geoIndexDefinitionModel

				// Construct an instance of the PutDesignDocumentOptions model
				putDesignDocumentOptionsModel := new(cloudantv1.PutDesignDocumentOptions)
				putDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				putDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				putDesignDocumentOptionsModel.DesignDocument = designDocumentModel
				putDesignDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				putDesignDocumentOptionsModel.Batch = core.StringPtr("ok")
				putDesignDocumentOptionsModel.NewEdits = core.BoolPtr(true)
				putDesignDocumentOptionsModel.Rev = core.StringPtr("testString")
				putDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PutDesignDocument(putDesignDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutDesignDocument(putDesignDocumentOptions *PutDesignDocumentOptions)`, func() {
		putDesignDocumentPath := "/testString/_design/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDesignDocumentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))


					// TODO: Add check for new_edits query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke PutDesignDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PutDesignDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Analyzer model
				analyzerModel := new(cloudantv1.Analyzer)
				analyzerModel.Name = core.StringPtr("classic")
				analyzerModel.Stopwords = []string{"testString"}

				// Construct an instance of the AnalyzerConfiguration model
				analyzerConfigurationModel := new(cloudantv1.AnalyzerConfiguration)
				analyzerConfigurationModel.Name = core.StringPtr("classic")
				analyzerConfigurationModel.Stopwords = []string{"testString"}
				analyzerConfigurationModel.Fields = make(map[string]cloudantv1.Analyzer)
				analyzerConfigurationModel.Fields["foo"] = *analyzerModel

				// Construct an instance of the SearchIndexDefinition model
				searchIndexDefinitionModel := new(cloudantv1.SearchIndexDefinition)
				searchIndexDefinitionModel.Analyzer = analyzerConfigurationModel
				searchIndexDefinitionModel.Index = core.StringPtr("testString")

				// Construct an instance of the DesignDocumentOptions model
				designDocumentOptionsModel := new(cloudantv1.DesignDocumentOptions)
				designDocumentOptionsModel.Partitioned = core.BoolPtr(true)

				// Construct an instance of the DesignDocumentViewsMapReduce model
				designDocumentViewsMapReduceModel := new(cloudantv1.DesignDocumentViewsMapReduce)
				designDocumentViewsMapReduceModel.Map = core.StringPtr("testString")
				designDocumentViewsMapReduceModel.Reduce = core.StringPtr("testString")

				// Construct an instance of the GeoIndexDefinition model
				geoIndexDefinitionModel := new(cloudantv1.GeoIndexDefinition)
				geoIndexDefinitionModel.Index = core.StringPtr("testString")

				// Construct an instance of the DesignDocument model
				designDocumentModel := new(cloudantv1.DesignDocument)
				designDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				designDocumentModel.Conflicts = []string{"testString"}
				designDocumentModel.Deleted = core.BoolPtr(true)
				designDocumentModel.DeletedConflicts = []string{"testString"}
				designDocumentModel.ID = core.StringPtr("testString")
				designDocumentModel.LocalSeq = core.StringPtr("testString")
				designDocumentModel.Rev = core.StringPtr("testString")
				designDocumentModel.Revisions = revisionsModel
				designDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				designDocumentModel.Autoupdate = core.BoolPtr(true)
				designDocumentModel.Filters = make(map[string]string)
				designDocumentModel.Indexes = make(map[string]cloudantv1.SearchIndexDefinition)
				designDocumentModel.Language = core.StringPtr("testString")
				designDocumentModel.Options = designDocumentOptionsModel
				designDocumentModel.Updates = make(map[string]string)
				designDocumentModel.ValidateDocUpdate = make(map[string]string)
				designDocumentModel.Views = make(map[string]cloudantv1.DesignDocumentViewsMapReduce)
				designDocumentModel.StIndexes = make(map[string]cloudantv1.GeoIndexDefinition)
				designDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				designDocumentModel.Attachments["foo"] = *attachmentModel
				designDocumentModel.Indexes["foo"] = *searchIndexDefinitionModel
				designDocumentModel.Views["foo"] = *designDocumentViewsMapReduceModel
				designDocumentModel.StIndexes["foo"] = *geoIndexDefinitionModel

				// Construct an instance of the PutDesignDocumentOptions model
				putDesignDocumentOptionsModel := new(cloudantv1.PutDesignDocumentOptions)
				putDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				putDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				putDesignDocumentOptionsModel.DesignDocument = designDocumentModel
				putDesignDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				putDesignDocumentOptionsModel.Batch = core.StringPtr("ok")
				putDesignDocumentOptionsModel.NewEdits = core.BoolPtr(true)
				putDesignDocumentOptionsModel.Rev = core.StringPtr("testString")
				putDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PutDesignDocument(putDesignDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PutDesignDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Analyzer model
				analyzerModel := new(cloudantv1.Analyzer)
				analyzerModel.Name = core.StringPtr("classic")
				analyzerModel.Stopwords = []string{"testString"}

				// Construct an instance of the AnalyzerConfiguration model
				analyzerConfigurationModel := new(cloudantv1.AnalyzerConfiguration)
				analyzerConfigurationModel.Name = core.StringPtr("classic")
				analyzerConfigurationModel.Stopwords = []string{"testString"}
				analyzerConfigurationModel.Fields = make(map[string]cloudantv1.Analyzer)
				analyzerConfigurationModel.Fields["foo"] = *analyzerModel

				// Construct an instance of the SearchIndexDefinition model
				searchIndexDefinitionModel := new(cloudantv1.SearchIndexDefinition)
				searchIndexDefinitionModel.Analyzer = analyzerConfigurationModel
				searchIndexDefinitionModel.Index = core.StringPtr("testString")

				// Construct an instance of the DesignDocumentOptions model
				designDocumentOptionsModel := new(cloudantv1.DesignDocumentOptions)
				designDocumentOptionsModel.Partitioned = core.BoolPtr(true)

				// Construct an instance of the DesignDocumentViewsMapReduce model
				designDocumentViewsMapReduceModel := new(cloudantv1.DesignDocumentViewsMapReduce)
				designDocumentViewsMapReduceModel.Map = core.StringPtr("testString")
				designDocumentViewsMapReduceModel.Reduce = core.StringPtr("testString")

				// Construct an instance of the GeoIndexDefinition model
				geoIndexDefinitionModel := new(cloudantv1.GeoIndexDefinition)
				geoIndexDefinitionModel.Index = core.StringPtr("testString")

				// Construct an instance of the DesignDocument model
				designDocumentModel := new(cloudantv1.DesignDocument)
				designDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				designDocumentModel.Conflicts = []string{"testString"}
				designDocumentModel.Deleted = core.BoolPtr(true)
				designDocumentModel.DeletedConflicts = []string{"testString"}
				designDocumentModel.ID = core.StringPtr("testString")
				designDocumentModel.LocalSeq = core.StringPtr("testString")
				designDocumentModel.Rev = core.StringPtr("testString")
				designDocumentModel.Revisions = revisionsModel
				designDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				designDocumentModel.Autoupdate = core.BoolPtr(true)
				designDocumentModel.Filters = make(map[string]string)
				designDocumentModel.Indexes = make(map[string]cloudantv1.SearchIndexDefinition)
				designDocumentModel.Language = core.StringPtr("testString")
				designDocumentModel.Options = designDocumentOptionsModel
				designDocumentModel.Updates = make(map[string]string)
				designDocumentModel.ValidateDocUpdate = make(map[string]string)
				designDocumentModel.Views = make(map[string]cloudantv1.DesignDocumentViewsMapReduce)
				designDocumentModel.StIndexes = make(map[string]cloudantv1.GeoIndexDefinition)
				designDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				designDocumentModel.Attachments["foo"] = *attachmentModel
				designDocumentModel.Indexes["foo"] = *searchIndexDefinitionModel
				designDocumentModel.Views["foo"] = *designDocumentViewsMapReduceModel
				designDocumentModel.StIndexes["foo"] = *geoIndexDefinitionModel

				// Construct an instance of the PutDesignDocumentOptions model
				putDesignDocumentOptionsModel := new(cloudantv1.PutDesignDocumentOptions)
				putDesignDocumentOptionsModel.Db = core.StringPtr("testString")
				putDesignDocumentOptionsModel.Ddoc = core.StringPtr("testString")
				putDesignDocumentOptionsModel.DesignDocument = designDocumentModel
				putDesignDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				putDesignDocumentOptionsModel.Batch = core.StringPtr("ok")
				putDesignDocumentOptionsModel.NewEdits = core.BoolPtr(true)
				putDesignDocumentOptionsModel.Rev = core.StringPtr("testString")
				putDesignDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PutDesignDocument(putDesignDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutDesignDocumentOptions model with no property values
				putDesignDocumentOptionsModelNew := new(cloudantv1.PutDesignDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PutDesignDocument(putDesignDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDesignDocumentInformation(getDesignDocumentInformationOptions *GetDesignDocumentInformationOptions) - Operation response error`, func() {
		getDesignDocumentInformationPath := "/testString/_design/testString/_info"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDesignDocumentInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDesignDocumentInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDesignDocumentInformationOptions model
				getDesignDocumentInformationOptionsModel := new(cloudantv1.GetDesignDocumentInformationOptions)
				getDesignDocumentInformationOptionsModel.Db = core.StringPtr("testString")
				getDesignDocumentInformationOptionsModel.Ddoc = core.StringPtr("testString")
				getDesignDocumentInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetDesignDocumentInformation(getDesignDocumentInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDesignDocumentInformation(getDesignDocumentInformationOptions *GetDesignDocumentInformationOptions)`, func() {
		getDesignDocumentInformationPath := "/testString/_design/testString/_info"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDesignDocumentInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "view_index": {"compact_running": true, "language": "Language", "signature": "Signature", "sizes": {"active": 6, "external": 8, "file": 4}, "update_seq": "UpdateSeq", "updater_running": true, "waiting_clients": 0, "waiting_commit": false}}`)
				}))
			})
			It(`Invoke GetDesignDocumentInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetDesignDocumentInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDesignDocumentInformationOptions model
				getDesignDocumentInformationOptionsModel := new(cloudantv1.GetDesignDocumentInformationOptions)
				getDesignDocumentInformationOptionsModel.Db = core.StringPtr("testString")
				getDesignDocumentInformationOptionsModel.Ddoc = core.StringPtr("testString")
				getDesignDocumentInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetDesignDocumentInformation(getDesignDocumentInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDesignDocumentInformation with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDesignDocumentInformationOptions model
				getDesignDocumentInformationOptionsModel := new(cloudantv1.GetDesignDocumentInformationOptions)
				getDesignDocumentInformationOptionsModel.Db = core.StringPtr("testString")
				getDesignDocumentInformationOptionsModel.Ddoc = core.StringPtr("testString")
				getDesignDocumentInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetDesignDocumentInformation(getDesignDocumentInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDesignDocumentInformationOptions model with no property values
				getDesignDocumentInformationOptionsModelNew := new(cloudantv1.GetDesignDocumentInformationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetDesignDocumentInformation(getDesignDocumentInformationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDesignDocs(postDesignDocsOptions *PostDesignDocsOptions) - Operation response error`, func() {
		postDesignDocsPath := "/testString/_design_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDesignDocsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDesignDocs with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostDesignDocsOptions model
				postDesignDocsOptionsModel := new(cloudantv1.PostDesignDocsOptions)
				postDesignDocsOptionsModel.Db = core.StringPtr("testString")
				postDesignDocsOptionsModel.Accept = core.StringPtr("application/json")
				postDesignDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postDesignDocsOptionsModel.Attachments = core.BoolPtr(true)
				postDesignDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postDesignDocsOptionsModel.Descending = core.BoolPtr(true)
				postDesignDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postDesignDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postDesignDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postDesignDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postDesignDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postDesignDocsOptionsModel.Endkey = core.StringPtr("testString")
				postDesignDocsOptionsModel.Key = core.StringPtr("testString")
				postDesignDocsOptionsModel.Keys = []string{"testString"}
				postDesignDocsOptionsModel.Startkey = core.StringPtr("testString")
				postDesignDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostDesignDocs(postDesignDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostDesignDocs(postDesignDocsOptions *PostDesignDocsOptions)`, func() {
		postDesignDocsPath := "/testString/_design_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDesignDocsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "rows": [{"caused_by": "CausedBy", "error": "Error", "reason": "Reason", "doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "id": "ID", "key": "Key", "value": {"rev": "Rev"}}], "update_seq": "UpdateSeq"}`)
				}))
			})
			It(`Invoke PostDesignDocs successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostDesignDocs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostDesignDocsOptions model
				postDesignDocsOptionsModel := new(cloudantv1.PostDesignDocsOptions)
				postDesignDocsOptionsModel.Db = core.StringPtr("testString")
				postDesignDocsOptionsModel.Accept = core.StringPtr("application/json")
				postDesignDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postDesignDocsOptionsModel.Attachments = core.BoolPtr(true)
				postDesignDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postDesignDocsOptionsModel.Descending = core.BoolPtr(true)
				postDesignDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postDesignDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postDesignDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postDesignDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postDesignDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postDesignDocsOptionsModel.Endkey = core.StringPtr("testString")
				postDesignDocsOptionsModel.Key = core.StringPtr("testString")
				postDesignDocsOptionsModel.Keys = []string{"testString"}
				postDesignDocsOptionsModel.Startkey = core.StringPtr("testString")
				postDesignDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostDesignDocs(postDesignDocsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostDesignDocs with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostDesignDocsOptions model
				postDesignDocsOptionsModel := new(cloudantv1.PostDesignDocsOptions)
				postDesignDocsOptionsModel.Db = core.StringPtr("testString")
				postDesignDocsOptionsModel.Accept = core.StringPtr("application/json")
				postDesignDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postDesignDocsOptionsModel.Attachments = core.BoolPtr(true)
				postDesignDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postDesignDocsOptionsModel.Descending = core.BoolPtr(true)
				postDesignDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postDesignDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postDesignDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postDesignDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postDesignDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postDesignDocsOptionsModel.Endkey = core.StringPtr("testString")
				postDesignDocsOptionsModel.Key = core.StringPtr("testString")
				postDesignDocsOptionsModel.Keys = []string{"testString"}
				postDesignDocsOptionsModel.Startkey = core.StringPtr("testString")
				postDesignDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostDesignDocs(postDesignDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostDesignDocsOptions model with no property values
				postDesignDocsOptionsModelNew := new(cloudantv1.PostDesignDocsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostDesignDocs(postDesignDocsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDesignDocsQueries(postDesignDocsQueriesOptions *PostDesignDocsQueriesOptions) - Operation response error`, func() {
		postDesignDocsQueriesPath := "/testString/_design_docs/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDesignDocsQueriesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDesignDocsQueries with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostDesignDocsQueriesOptions model
				postDesignDocsQueriesOptionsModel := new(cloudantv1.PostDesignDocsQueriesOptions)
				postDesignDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postDesignDocsQueriesOptionsModel.Accept = core.StringPtr("application/json")
				postDesignDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postDesignDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostDesignDocsQueries(postDesignDocsQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostDesignDocsQueries(postDesignDocsQueriesOptions *PostDesignDocsQueriesOptions)`, func() {
		postDesignDocsQueriesPath := "/testString/_design_docs/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDesignDocsQueriesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"total_rows": 0, "rows": [{"caused_by": "CausedBy", "error": "Error", "reason": "Reason", "doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "id": "ID", "key": "Key", "value": {"rev": "Rev"}}], "update_seq": "UpdateSeq"}]}`)
				}))
			})
			It(`Invoke PostDesignDocsQueries successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostDesignDocsQueries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostDesignDocsQueriesOptions model
				postDesignDocsQueriesOptionsModel := new(cloudantv1.PostDesignDocsQueriesOptions)
				postDesignDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postDesignDocsQueriesOptionsModel.Accept = core.StringPtr("application/json")
				postDesignDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postDesignDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostDesignDocsQueries(postDesignDocsQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostDesignDocsQueries with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostDesignDocsQueriesOptions model
				postDesignDocsQueriesOptionsModel := new(cloudantv1.PostDesignDocsQueriesOptions)
				postDesignDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postDesignDocsQueriesOptionsModel.Accept = core.StringPtr("application/json")
				postDesignDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postDesignDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostDesignDocsQueries(postDesignDocsQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostDesignDocsQueriesOptions model with no property values
				postDesignDocsQueriesOptionsModelNew := new(cloudantv1.PostDesignDocsQueriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostDesignDocsQueries(postDesignDocsQueriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`PostView(postViewOptions *PostViewOptions) - Operation response error`, func() {
		postViewPath := "/testString/_design/testString/_view/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postViewPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostView with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostViewOptions model
				postViewOptionsModel := new(cloudantv1.PostViewOptions)
				postViewOptionsModel.Db = core.StringPtr("testString")
				postViewOptionsModel.Ddoc = core.StringPtr("testString")
				postViewOptionsModel.View = core.StringPtr("testString")
				postViewOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postViewOptionsModel.Attachments = core.BoolPtr(true)
				postViewOptionsModel.Conflicts = core.BoolPtr(true)
				postViewOptionsModel.Descending = core.BoolPtr(true)
				postViewOptionsModel.IncludeDocs = core.BoolPtr(true)
				postViewOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postViewOptionsModel.Limit = core.Int64Ptr(int64(0))
				postViewOptionsModel.Skip = core.Int64Ptr(int64(0))
				postViewOptionsModel.UpdateSeq = core.BoolPtr(true)
				postViewOptionsModel.Endkey = core.StringPtr("testString")
				postViewOptionsModel.EndkeyDocid = core.StringPtr("testString")
				postViewOptionsModel.Group = core.BoolPtr(true)
				postViewOptionsModel.GroupLevel = core.Int64Ptr(int64(1))
				postViewOptionsModel.Key = core.StringPtr("testString")
				postViewOptionsModel.Keys = []interface{}{"testString"}
				postViewOptionsModel.Reduce = core.BoolPtr(true)
				postViewOptionsModel.Stable = core.BoolPtr(true)
				postViewOptionsModel.Startkey = core.StringPtr("testString")
				postViewOptionsModel.StartkeyDocid = core.StringPtr("testString")
				postViewOptionsModel.Update = core.StringPtr("true")
				postViewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostView(postViewOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostView(postViewOptions *PostViewOptions)`, func() {
		postViewPath := "/testString/_design/testString/_view/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postViewPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "update_seq": "UpdateSeq", "rows": [{"caused_by": "CausedBy", "error": "Error", "reason": "Reason", "doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "id": "ID", "key": "anyValue", "value": "anyValue"}]}`)
				}))
			})
			It(`Invoke PostView successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostView(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostViewOptions model
				postViewOptionsModel := new(cloudantv1.PostViewOptions)
				postViewOptionsModel.Db = core.StringPtr("testString")
				postViewOptionsModel.Ddoc = core.StringPtr("testString")
				postViewOptionsModel.View = core.StringPtr("testString")
				postViewOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postViewOptionsModel.Attachments = core.BoolPtr(true)
				postViewOptionsModel.Conflicts = core.BoolPtr(true)
				postViewOptionsModel.Descending = core.BoolPtr(true)
				postViewOptionsModel.IncludeDocs = core.BoolPtr(true)
				postViewOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postViewOptionsModel.Limit = core.Int64Ptr(int64(0))
				postViewOptionsModel.Skip = core.Int64Ptr(int64(0))
				postViewOptionsModel.UpdateSeq = core.BoolPtr(true)
				postViewOptionsModel.Endkey = core.StringPtr("testString")
				postViewOptionsModel.EndkeyDocid = core.StringPtr("testString")
				postViewOptionsModel.Group = core.BoolPtr(true)
				postViewOptionsModel.GroupLevel = core.Int64Ptr(int64(1))
				postViewOptionsModel.Key = core.StringPtr("testString")
				postViewOptionsModel.Keys = []interface{}{"testString"}
				postViewOptionsModel.Reduce = core.BoolPtr(true)
				postViewOptionsModel.Stable = core.BoolPtr(true)
				postViewOptionsModel.Startkey = core.StringPtr("testString")
				postViewOptionsModel.StartkeyDocid = core.StringPtr("testString")
				postViewOptionsModel.Update = core.StringPtr("true")
				postViewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostView(postViewOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostView with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostViewOptions model
				postViewOptionsModel := new(cloudantv1.PostViewOptions)
				postViewOptionsModel.Db = core.StringPtr("testString")
				postViewOptionsModel.Ddoc = core.StringPtr("testString")
				postViewOptionsModel.View = core.StringPtr("testString")
				postViewOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postViewOptionsModel.Attachments = core.BoolPtr(true)
				postViewOptionsModel.Conflicts = core.BoolPtr(true)
				postViewOptionsModel.Descending = core.BoolPtr(true)
				postViewOptionsModel.IncludeDocs = core.BoolPtr(true)
				postViewOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postViewOptionsModel.Limit = core.Int64Ptr(int64(0))
				postViewOptionsModel.Skip = core.Int64Ptr(int64(0))
				postViewOptionsModel.UpdateSeq = core.BoolPtr(true)
				postViewOptionsModel.Endkey = core.StringPtr("testString")
				postViewOptionsModel.EndkeyDocid = core.StringPtr("testString")
				postViewOptionsModel.Group = core.BoolPtr(true)
				postViewOptionsModel.GroupLevel = core.Int64Ptr(int64(1))
				postViewOptionsModel.Key = core.StringPtr("testString")
				postViewOptionsModel.Keys = []interface{}{"testString"}
				postViewOptionsModel.Reduce = core.BoolPtr(true)
				postViewOptionsModel.Stable = core.BoolPtr(true)
				postViewOptionsModel.Startkey = core.StringPtr("testString")
				postViewOptionsModel.StartkeyDocid = core.StringPtr("testString")
				postViewOptionsModel.Update = core.StringPtr("true")
				postViewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostView(postViewOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostViewOptions model with no property values
				postViewOptionsModelNew := new(cloudantv1.PostViewOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostView(postViewOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostViewAsStream(postViewOptions *PostViewOptions)`, func() {
		postViewAsStreamPath := "/testString/_design/testString/_view/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postViewAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostViewAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostViewAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostViewOptions model
				postViewOptionsModel := new(cloudantv1.PostViewOptions)
				postViewOptionsModel.Db = core.StringPtr("testString")
				postViewOptionsModel.Ddoc = core.StringPtr("testString")
				postViewOptionsModel.View = core.StringPtr("testString")
				postViewOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postViewOptionsModel.Attachments = core.BoolPtr(true)
				postViewOptionsModel.Conflicts = core.BoolPtr(true)
				postViewOptionsModel.Descending = core.BoolPtr(true)
				postViewOptionsModel.IncludeDocs = core.BoolPtr(true)
				postViewOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postViewOptionsModel.Limit = core.Int64Ptr(int64(0))
				postViewOptionsModel.Skip = core.Int64Ptr(int64(0))
				postViewOptionsModel.UpdateSeq = core.BoolPtr(true)
				postViewOptionsModel.Endkey = core.StringPtr("testString")
				postViewOptionsModel.EndkeyDocid = core.StringPtr("testString")
				postViewOptionsModel.Group = core.BoolPtr(true)
				postViewOptionsModel.GroupLevel = core.Int64Ptr(int64(1))
				postViewOptionsModel.Key = core.StringPtr("testString")
				postViewOptionsModel.Keys = []interface{}{"testString"}
				postViewOptionsModel.Reduce = core.BoolPtr(true)
				postViewOptionsModel.Stable = core.BoolPtr(true)
				postViewOptionsModel.Startkey = core.StringPtr("testString")
				postViewOptionsModel.StartkeyDocid = core.StringPtr("testString")
				postViewOptionsModel.Update = core.StringPtr("true")
				postViewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostViewAsStream(postViewOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostViewAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostViewOptions model
				postViewOptionsModel := new(cloudantv1.PostViewOptions)
				postViewOptionsModel.Db = core.StringPtr("testString")
				postViewOptionsModel.Ddoc = core.StringPtr("testString")
				postViewOptionsModel.View = core.StringPtr("testString")
				postViewOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postViewOptionsModel.Attachments = core.BoolPtr(true)
				postViewOptionsModel.Conflicts = core.BoolPtr(true)
				postViewOptionsModel.Descending = core.BoolPtr(true)
				postViewOptionsModel.IncludeDocs = core.BoolPtr(true)
				postViewOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postViewOptionsModel.Limit = core.Int64Ptr(int64(0))
				postViewOptionsModel.Skip = core.Int64Ptr(int64(0))
				postViewOptionsModel.UpdateSeq = core.BoolPtr(true)
				postViewOptionsModel.Endkey = core.StringPtr("testString")
				postViewOptionsModel.EndkeyDocid = core.StringPtr("testString")
				postViewOptionsModel.Group = core.BoolPtr(true)
				postViewOptionsModel.GroupLevel = core.Int64Ptr(int64(1))
				postViewOptionsModel.Key = core.StringPtr("testString")
				postViewOptionsModel.Keys = []interface{}{"testString"}
				postViewOptionsModel.Reduce = core.BoolPtr(true)
				postViewOptionsModel.Stable = core.BoolPtr(true)
				postViewOptionsModel.Startkey = core.StringPtr("testString")
				postViewOptionsModel.StartkeyDocid = core.StringPtr("testString")
				postViewOptionsModel.Update = core.StringPtr("true")
				postViewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostViewAsStream(postViewOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostViewOptions model with no property values
				postViewOptionsModelNew := new(cloudantv1.PostViewOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostViewAsStream(postViewOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostViewQueries(postViewQueriesOptions *PostViewQueriesOptions) - Operation response error`, func() {
		postViewQueriesPath := "/testString/_design/testString/_view/testString/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postViewQueriesPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostViewQueries with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the ViewQuery model
				viewQueryModel := new(cloudantv1.ViewQuery)
				viewQueryModel.AttEncodingInfo = core.BoolPtr(true)
				viewQueryModel.Attachments = core.BoolPtr(true)
				viewQueryModel.Conflicts = core.BoolPtr(true)
				viewQueryModel.Descending = core.BoolPtr(true)
				viewQueryModel.IncludeDocs = core.BoolPtr(true)
				viewQueryModel.InclusiveEnd = core.BoolPtr(true)
				viewQueryModel.Limit = core.Int64Ptr(int64(0))
				viewQueryModel.Skip = core.Int64Ptr(int64(0))
				viewQueryModel.UpdateSeq = core.BoolPtr(true)
				viewQueryModel.Endkey = core.StringPtr("testString")
				viewQueryModel.EndkeyDocid = core.StringPtr("testString")
				viewQueryModel.Group = core.BoolPtr(true)
				viewQueryModel.GroupLevel = core.Int64Ptr(int64(1))
				viewQueryModel.Key = core.StringPtr("testString")
				viewQueryModel.Keys = []interface{}{"testString"}
				viewQueryModel.Reduce = core.BoolPtr(true)
				viewQueryModel.Stable = core.BoolPtr(true)
				viewQueryModel.Startkey = core.StringPtr("testString")
				viewQueryModel.StartkeyDocid = core.StringPtr("testString")
				viewQueryModel.Update = core.StringPtr("true")

				// Construct an instance of the PostViewQueriesOptions model
				postViewQueriesOptionsModel := new(cloudantv1.PostViewQueriesOptions)
				postViewQueriesOptionsModel.Db = core.StringPtr("testString")
				postViewQueriesOptionsModel.Ddoc = core.StringPtr("testString")
				postViewQueriesOptionsModel.View = core.StringPtr("testString")
				postViewQueriesOptionsModel.Queries = []cloudantv1.ViewQuery{*viewQueryModel}
				postViewQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostViewQueries(postViewQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostViewQueries(postViewQueriesOptions *PostViewQueriesOptions)`, func() {
		postViewQueriesPath := "/testString/_design/testString/_view/testString/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postViewQueriesPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"total_rows": 0, "update_seq": "UpdateSeq", "rows": [{"caused_by": "CausedBy", "error": "Error", "reason": "Reason", "doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "id": "ID", "key": "anyValue", "value": "anyValue"}]}]}`)
				}))
			})
			It(`Invoke PostViewQueries successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostViewQueries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ViewQuery model
				viewQueryModel := new(cloudantv1.ViewQuery)
				viewQueryModel.AttEncodingInfo = core.BoolPtr(true)
				viewQueryModel.Attachments = core.BoolPtr(true)
				viewQueryModel.Conflicts = core.BoolPtr(true)
				viewQueryModel.Descending = core.BoolPtr(true)
				viewQueryModel.IncludeDocs = core.BoolPtr(true)
				viewQueryModel.InclusiveEnd = core.BoolPtr(true)
				viewQueryModel.Limit = core.Int64Ptr(int64(0))
				viewQueryModel.Skip = core.Int64Ptr(int64(0))
				viewQueryModel.UpdateSeq = core.BoolPtr(true)
				viewQueryModel.Endkey = core.StringPtr("testString")
				viewQueryModel.EndkeyDocid = core.StringPtr("testString")
				viewQueryModel.Group = core.BoolPtr(true)
				viewQueryModel.GroupLevel = core.Int64Ptr(int64(1))
				viewQueryModel.Key = core.StringPtr("testString")
				viewQueryModel.Keys = []interface{}{"testString"}
				viewQueryModel.Reduce = core.BoolPtr(true)
				viewQueryModel.Stable = core.BoolPtr(true)
				viewQueryModel.Startkey = core.StringPtr("testString")
				viewQueryModel.StartkeyDocid = core.StringPtr("testString")
				viewQueryModel.Update = core.StringPtr("true")

				// Construct an instance of the PostViewQueriesOptions model
				postViewQueriesOptionsModel := new(cloudantv1.PostViewQueriesOptions)
				postViewQueriesOptionsModel.Db = core.StringPtr("testString")
				postViewQueriesOptionsModel.Ddoc = core.StringPtr("testString")
				postViewQueriesOptionsModel.View = core.StringPtr("testString")
				postViewQueriesOptionsModel.Queries = []cloudantv1.ViewQuery{*viewQueryModel}
				postViewQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostViewQueries(postViewQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostViewQueries with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the ViewQuery model
				viewQueryModel := new(cloudantv1.ViewQuery)
				viewQueryModel.AttEncodingInfo = core.BoolPtr(true)
				viewQueryModel.Attachments = core.BoolPtr(true)
				viewQueryModel.Conflicts = core.BoolPtr(true)
				viewQueryModel.Descending = core.BoolPtr(true)
				viewQueryModel.IncludeDocs = core.BoolPtr(true)
				viewQueryModel.InclusiveEnd = core.BoolPtr(true)
				viewQueryModel.Limit = core.Int64Ptr(int64(0))
				viewQueryModel.Skip = core.Int64Ptr(int64(0))
				viewQueryModel.UpdateSeq = core.BoolPtr(true)
				viewQueryModel.Endkey = core.StringPtr("testString")
				viewQueryModel.EndkeyDocid = core.StringPtr("testString")
				viewQueryModel.Group = core.BoolPtr(true)
				viewQueryModel.GroupLevel = core.Int64Ptr(int64(1))
				viewQueryModel.Key = core.StringPtr("testString")
				viewQueryModel.Keys = []interface{}{"testString"}
				viewQueryModel.Reduce = core.BoolPtr(true)
				viewQueryModel.Stable = core.BoolPtr(true)
				viewQueryModel.Startkey = core.StringPtr("testString")
				viewQueryModel.StartkeyDocid = core.StringPtr("testString")
				viewQueryModel.Update = core.StringPtr("true")

				// Construct an instance of the PostViewQueriesOptions model
				postViewQueriesOptionsModel := new(cloudantv1.PostViewQueriesOptions)
				postViewQueriesOptionsModel.Db = core.StringPtr("testString")
				postViewQueriesOptionsModel.Ddoc = core.StringPtr("testString")
				postViewQueriesOptionsModel.View = core.StringPtr("testString")
				postViewQueriesOptionsModel.Queries = []cloudantv1.ViewQuery{*viewQueryModel}
				postViewQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostViewQueries(postViewQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostViewQueriesOptions model with no property values
				postViewQueriesOptionsModelNew := new(cloudantv1.PostViewQueriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostViewQueries(postViewQueriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostViewQueriesAsStream(postViewQueriesOptions *PostViewQueriesOptions)`, func() {
		postViewQueriesAsStreamPath := "/testString/_design/testString/_view/testString/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postViewQueriesAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostViewQueriesAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostViewQueriesAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ViewQuery model
				viewQueryModel := new(cloudantv1.ViewQuery)
				viewQueryModel.AttEncodingInfo = core.BoolPtr(true)
				viewQueryModel.Attachments = core.BoolPtr(true)
				viewQueryModel.Conflicts = core.BoolPtr(true)
				viewQueryModel.Descending = core.BoolPtr(true)
				viewQueryModel.IncludeDocs = core.BoolPtr(true)
				viewQueryModel.InclusiveEnd = core.BoolPtr(true)
				viewQueryModel.Limit = core.Int64Ptr(int64(0))
				viewQueryModel.Skip = core.Int64Ptr(int64(0))
				viewQueryModel.UpdateSeq = core.BoolPtr(true)
				viewQueryModel.Endkey = core.StringPtr("testString")
				viewQueryModel.EndkeyDocid = core.StringPtr("testString")
				viewQueryModel.Group = core.BoolPtr(true)
				viewQueryModel.GroupLevel = core.Int64Ptr(int64(1))
				viewQueryModel.Key = core.StringPtr("testString")
				viewQueryModel.Keys = []interface{}{"testString"}
				viewQueryModel.Reduce = core.BoolPtr(true)
				viewQueryModel.Stable = core.BoolPtr(true)
				viewQueryModel.Startkey = core.StringPtr("testString")
				viewQueryModel.StartkeyDocid = core.StringPtr("testString")
				viewQueryModel.Update = core.StringPtr("true")

				// Construct an instance of the PostViewQueriesOptions model
				postViewQueriesOptionsModel := new(cloudantv1.PostViewQueriesOptions)
				postViewQueriesOptionsModel.Db = core.StringPtr("testString")
				postViewQueriesOptionsModel.Ddoc = core.StringPtr("testString")
				postViewQueriesOptionsModel.View = core.StringPtr("testString")
				postViewQueriesOptionsModel.Queries = []cloudantv1.ViewQuery{*viewQueryModel}
				postViewQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostViewQueriesAsStream(postViewQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostViewQueriesAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the ViewQuery model
				viewQueryModel := new(cloudantv1.ViewQuery)
				viewQueryModel.AttEncodingInfo = core.BoolPtr(true)
				viewQueryModel.Attachments = core.BoolPtr(true)
				viewQueryModel.Conflicts = core.BoolPtr(true)
				viewQueryModel.Descending = core.BoolPtr(true)
				viewQueryModel.IncludeDocs = core.BoolPtr(true)
				viewQueryModel.InclusiveEnd = core.BoolPtr(true)
				viewQueryModel.Limit = core.Int64Ptr(int64(0))
				viewQueryModel.Skip = core.Int64Ptr(int64(0))
				viewQueryModel.UpdateSeq = core.BoolPtr(true)
				viewQueryModel.Endkey = core.StringPtr("testString")
				viewQueryModel.EndkeyDocid = core.StringPtr("testString")
				viewQueryModel.Group = core.BoolPtr(true)
				viewQueryModel.GroupLevel = core.Int64Ptr(int64(1))
				viewQueryModel.Key = core.StringPtr("testString")
				viewQueryModel.Keys = []interface{}{"testString"}
				viewQueryModel.Reduce = core.BoolPtr(true)
				viewQueryModel.Stable = core.BoolPtr(true)
				viewQueryModel.Startkey = core.StringPtr("testString")
				viewQueryModel.StartkeyDocid = core.StringPtr("testString")
				viewQueryModel.Update = core.StringPtr("true")

				// Construct an instance of the PostViewQueriesOptions model
				postViewQueriesOptionsModel := new(cloudantv1.PostViewQueriesOptions)
				postViewQueriesOptionsModel.Db = core.StringPtr("testString")
				postViewQueriesOptionsModel.Ddoc = core.StringPtr("testString")
				postViewQueriesOptionsModel.View = core.StringPtr("testString")
				postViewQueriesOptionsModel.Queries = []cloudantv1.ViewQuery{*viewQueryModel}
				postViewQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostViewQueriesAsStream(postViewQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostViewQueriesOptions model with no property values
				postViewQueriesOptionsModelNew := new(cloudantv1.PostViewQueriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostViewQueriesAsStream(postViewQueriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetPartitionInformation(getPartitionInformationOptions *GetPartitionInformationOptions) - Operation response error`, func() {
		getPartitionInformationPath := "/testString/_partition/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPartitionInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPartitionInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetPartitionInformationOptions model
				getPartitionInformationOptionsModel := new(cloudantv1.GetPartitionInformationOptions)
				getPartitionInformationOptionsModel.Db = core.StringPtr("testString")
				getPartitionInformationOptionsModel.PartitionKey = core.StringPtr("testString")
				getPartitionInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetPartitionInformation(getPartitionInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPartitionInformation(getPartitionInformationOptions *GetPartitionInformationOptions)`, func() {
		getPartitionInformationPath := "/testString/_partition/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPartitionInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"db_name": "DbName", "doc_count": 0, "doc_del_count": 0, "partition": "Partition", "partitioned_indexes": {"count": 0, "indexes": {"search": 0, "view": 0}, "limit": 0}, "sizes": {"active": 0, "external": 0}}`)
				}))
			})
			It(`Invoke GetPartitionInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetPartitionInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPartitionInformationOptions model
				getPartitionInformationOptionsModel := new(cloudantv1.GetPartitionInformationOptions)
				getPartitionInformationOptionsModel.Db = core.StringPtr("testString")
				getPartitionInformationOptionsModel.PartitionKey = core.StringPtr("testString")
				getPartitionInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetPartitionInformation(getPartitionInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetPartitionInformation with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetPartitionInformationOptions model
				getPartitionInformationOptionsModel := new(cloudantv1.GetPartitionInformationOptions)
				getPartitionInformationOptionsModel.Db = core.StringPtr("testString")
				getPartitionInformationOptionsModel.PartitionKey = core.StringPtr("testString")
				getPartitionInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetPartitionInformation(getPartitionInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPartitionInformationOptions model with no property values
				getPartitionInformationOptionsModelNew := new(cloudantv1.GetPartitionInformationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetPartitionInformation(getPartitionInformationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostPartitionAllDocs(postPartitionAllDocsOptions *PostPartitionAllDocsOptions) - Operation response error`, func() {
		postPartitionAllDocsPath := "/testString/_partition/testString/_all_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionAllDocsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostPartitionAllDocs with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionAllDocsOptions model
				postPartitionAllDocsOptionsModel := new(cloudantv1.PostPartitionAllDocsOptions)
				postPartitionAllDocsOptionsModel.Db = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Attachments = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Descending = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionAllDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionAllDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Endkey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Key = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Keys = []string{"testString"}
				postPartitionAllDocsOptionsModel.Startkey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostPartitionAllDocs(postPartitionAllDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostPartitionAllDocs(postPartitionAllDocsOptions *PostPartitionAllDocsOptions)`, func() {
		postPartitionAllDocsPath := "/testString/_partition/testString/_all_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionAllDocsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "rows": [{"caused_by": "CausedBy", "error": "Error", "reason": "Reason", "doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "id": "ID", "key": "Key", "value": {"rev": "Rev"}}], "update_seq": "UpdateSeq"}`)
				}))
			})
			It(`Invoke PostPartitionAllDocs successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostPartitionAllDocs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostPartitionAllDocsOptions model
				postPartitionAllDocsOptionsModel := new(cloudantv1.PostPartitionAllDocsOptions)
				postPartitionAllDocsOptionsModel.Db = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Attachments = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Descending = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionAllDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionAllDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Endkey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Key = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Keys = []string{"testString"}
				postPartitionAllDocsOptionsModel.Startkey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostPartitionAllDocs(postPartitionAllDocsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostPartitionAllDocs with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionAllDocsOptions model
				postPartitionAllDocsOptionsModel := new(cloudantv1.PostPartitionAllDocsOptions)
				postPartitionAllDocsOptionsModel.Db = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Attachments = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Descending = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionAllDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionAllDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Endkey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Key = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Keys = []string{"testString"}
				postPartitionAllDocsOptionsModel.Startkey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostPartitionAllDocs(postPartitionAllDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostPartitionAllDocsOptions model with no property values
				postPartitionAllDocsOptionsModelNew := new(cloudantv1.PostPartitionAllDocsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostPartitionAllDocs(postPartitionAllDocsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostPartitionAllDocsAsStream(postPartitionAllDocsOptions *PostPartitionAllDocsOptions)`, func() {
		postPartitionAllDocsAsStreamPath := "/testString/_partition/testString/_all_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionAllDocsAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostPartitionAllDocsAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostPartitionAllDocsAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostPartitionAllDocsOptions model
				postPartitionAllDocsOptionsModel := new(cloudantv1.PostPartitionAllDocsOptions)
				postPartitionAllDocsOptionsModel.Db = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Attachments = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Descending = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionAllDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionAllDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Endkey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Key = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Keys = []string{"testString"}
				postPartitionAllDocsOptionsModel.Startkey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostPartitionAllDocsAsStream(postPartitionAllDocsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostPartitionAllDocsAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionAllDocsOptions model
				postPartitionAllDocsOptionsModel := new(cloudantv1.PostPartitionAllDocsOptions)
				postPartitionAllDocsOptionsModel.Db = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Attachments = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Descending = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionAllDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionAllDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postPartitionAllDocsOptionsModel.Endkey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Key = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Keys = []string{"testString"}
				postPartitionAllDocsOptionsModel.Startkey = core.StringPtr("testString")
				postPartitionAllDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostPartitionAllDocsAsStream(postPartitionAllDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostPartitionAllDocsOptions model with no property values
				postPartitionAllDocsOptionsModelNew := new(cloudantv1.PostPartitionAllDocsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostPartitionAllDocsAsStream(postPartitionAllDocsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostPartitionSearch(postPartitionSearchOptions *PostPartitionSearchOptions) - Operation response error`, func() {
		postPartitionSearchPath := "/testString/_partition/testString/_design/testString/_search/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionSearchPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostPartitionSearch with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionSearchOptions model
				postPartitionSearchOptionsModel := new(cloudantv1.PostPartitionSearchOptions)
				postPartitionSearchOptionsModel.Db = core.StringPtr("testString")
				postPartitionSearchOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Ddoc = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Index = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Bookmark = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightFields = []string{"testString"}
				postPartitionSearchOptionsModel.HighlightNumber = core.Int64Ptr(int64(1))
				postPartitionSearchOptionsModel.HighlightPostTag = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightPreTag = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightSize = core.Int64Ptr(int64(1))
				postPartitionSearchOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionSearchOptionsModel.IncludeFields = []string{"testString"}
				postPartitionSearchOptionsModel.Limit = core.Int64Ptr(int64(3))
				postPartitionSearchOptionsModel.Query = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Sort = []string{"testString"}
				postPartitionSearchOptionsModel.Stale = core.StringPtr("ok")
				postPartitionSearchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostPartitionSearch(postPartitionSearchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostPartitionSearch(postPartitionSearchOptions *PostPartitionSearchOptions)`, func() {
		postPartitionSearchPath := "/testString/_partition/testString/_design/testString/_search/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionSearchPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "bookmark": "Bookmark", "by": "By", "counts": {"mapKey": {"mapKey": 0}}, "ranges": {"mapKey": {"mapKey": 0}}, "rows": [{"doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "fields": {"mapKey": "anyValue"}, "highlights": {"mapKey": ["Inner"]}, "id": "ID"}], "groups": [{"total_rows": 0, "bookmark": "Bookmark", "by": "By", "counts": {"mapKey": {"mapKey": 0}}, "ranges": {"mapKey": {"mapKey": 0}}, "rows": [{"doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "fields": {"mapKey": "anyValue"}, "highlights": {"mapKey": ["Inner"]}, "id": "ID"}]}]}`)
				}))
			})
			It(`Invoke PostPartitionSearch successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostPartitionSearch(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostPartitionSearchOptions model
				postPartitionSearchOptionsModel := new(cloudantv1.PostPartitionSearchOptions)
				postPartitionSearchOptionsModel.Db = core.StringPtr("testString")
				postPartitionSearchOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Ddoc = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Index = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Bookmark = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightFields = []string{"testString"}
				postPartitionSearchOptionsModel.HighlightNumber = core.Int64Ptr(int64(1))
				postPartitionSearchOptionsModel.HighlightPostTag = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightPreTag = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightSize = core.Int64Ptr(int64(1))
				postPartitionSearchOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionSearchOptionsModel.IncludeFields = []string{"testString"}
				postPartitionSearchOptionsModel.Limit = core.Int64Ptr(int64(3))
				postPartitionSearchOptionsModel.Query = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Sort = []string{"testString"}
				postPartitionSearchOptionsModel.Stale = core.StringPtr("ok")
				postPartitionSearchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostPartitionSearch(postPartitionSearchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostPartitionSearch with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionSearchOptions model
				postPartitionSearchOptionsModel := new(cloudantv1.PostPartitionSearchOptions)
				postPartitionSearchOptionsModel.Db = core.StringPtr("testString")
				postPartitionSearchOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Ddoc = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Index = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Bookmark = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightFields = []string{"testString"}
				postPartitionSearchOptionsModel.HighlightNumber = core.Int64Ptr(int64(1))
				postPartitionSearchOptionsModel.HighlightPostTag = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightPreTag = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightSize = core.Int64Ptr(int64(1))
				postPartitionSearchOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionSearchOptionsModel.IncludeFields = []string{"testString"}
				postPartitionSearchOptionsModel.Limit = core.Int64Ptr(int64(3))
				postPartitionSearchOptionsModel.Query = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Sort = []string{"testString"}
				postPartitionSearchOptionsModel.Stale = core.StringPtr("ok")
				postPartitionSearchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostPartitionSearch(postPartitionSearchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostPartitionSearchOptions model with no property values
				postPartitionSearchOptionsModelNew := new(cloudantv1.PostPartitionSearchOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostPartitionSearch(postPartitionSearchOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostPartitionSearchAsStream(postPartitionSearchOptions *PostPartitionSearchOptions)`, func() {
		postPartitionSearchAsStreamPath := "/testString/_partition/testString/_design/testString/_search/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionSearchAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostPartitionSearchAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostPartitionSearchAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostPartitionSearchOptions model
				postPartitionSearchOptionsModel := new(cloudantv1.PostPartitionSearchOptions)
				postPartitionSearchOptionsModel.Db = core.StringPtr("testString")
				postPartitionSearchOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Ddoc = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Index = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Bookmark = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightFields = []string{"testString"}
				postPartitionSearchOptionsModel.HighlightNumber = core.Int64Ptr(int64(1))
				postPartitionSearchOptionsModel.HighlightPostTag = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightPreTag = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightSize = core.Int64Ptr(int64(1))
				postPartitionSearchOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionSearchOptionsModel.IncludeFields = []string{"testString"}
				postPartitionSearchOptionsModel.Limit = core.Int64Ptr(int64(3))
				postPartitionSearchOptionsModel.Query = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Sort = []string{"testString"}
				postPartitionSearchOptionsModel.Stale = core.StringPtr("ok")
				postPartitionSearchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostPartitionSearchAsStream(postPartitionSearchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostPartitionSearchAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionSearchOptions model
				postPartitionSearchOptionsModel := new(cloudantv1.PostPartitionSearchOptions)
				postPartitionSearchOptionsModel.Db = core.StringPtr("testString")
				postPartitionSearchOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Ddoc = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Index = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Bookmark = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightFields = []string{"testString"}
				postPartitionSearchOptionsModel.HighlightNumber = core.Int64Ptr(int64(1))
				postPartitionSearchOptionsModel.HighlightPostTag = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightPreTag = core.StringPtr("testString")
				postPartitionSearchOptionsModel.HighlightSize = core.Int64Ptr(int64(1))
				postPartitionSearchOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionSearchOptionsModel.IncludeFields = []string{"testString"}
				postPartitionSearchOptionsModel.Limit = core.Int64Ptr(int64(3))
				postPartitionSearchOptionsModel.Query = core.StringPtr("testString")
				postPartitionSearchOptionsModel.Sort = []string{"testString"}
				postPartitionSearchOptionsModel.Stale = core.StringPtr("ok")
				postPartitionSearchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostPartitionSearchAsStream(postPartitionSearchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostPartitionSearchOptions model with no property values
				postPartitionSearchOptionsModelNew := new(cloudantv1.PostPartitionSearchOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostPartitionSearchAsStream(postPartitionSearchOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostPartitionView(postPartitionViewOptions *PostPartitionViewOptions) - Operation response error`, func() {
		postPartitionViewPath := "/testString/_partition/testString/_design/testString/_view/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionViewPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostPartitionView with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionViewOptions model
				postPartitionViewOptionsModel := new(cloudantv1.PostPartitionViewOptions)
				postPartitionViewOptionsModel.Db = core.StringPtr("testString")
				postPartitionViewOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionViewOptionsModel.Ddoc = core.StringPtr("testString")
				postPartitionViewOptionsModel.View = core.StringPtr("testString")
				postPartitionViewOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postPartitionViewOptionsModel.Attachments = core.BoolPtr(true)
				postPartitionViewOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionViewOptionsModel.Descending = core.BoolPtr(true)
				postPartitionViewOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionViewOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postPartitionViewOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionViewOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionViewOptionsModel.UpdateSeq = core.BoolPtr(true)
				postPartitionViewOptionsModel.Endkey = core.StringPtr("testString")
				postPartitionViewOptionsModel.EndkeyDocid = core.StringPtr("testString")
				postPartitionViewOptionsModel.Group = core.BoolPtr(true)
				postPartitionViewOptionsModel.GroupLevel = core.Int64Ptr(int64(1))
				postPartitionViewOptionsModel.Key = core.StringPtr("testString")
				postPartitionViewOptionsModel.Keys = []interface{}{"testString"}
				postPartitionViewOptionsModel.Reduce = core.BoolPtr(true)
				postPartitionViewOptionsModel.Stable = core.BoolPtr(true)
				postPartitionViewOptionsModel.Startkey = core.StringPtr("testString")
				postPartitionViewOptionsModel.StartkeyDocid = core.StringPtr("testString")
				postPartitionViewOptionsModel.Update = core.StringPtr("true")
				postPartitionViewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostPartitionView(postPartitionViewOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostPartitionView(postPartitionViewOptions *PostPartitionViewOptions)`, func() {
		postPartitionViewPath := "/testString/_partition/testString/_design/testString/_view/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionViewPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "update_seq": "UpdateSeq", "rows": [{"caused_by": "CausedBy", "error": "Error", "reason": "Reason", "doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "id": "ID", "key": "anyValue", "value": "anyValue"}]}`)
				}))
			})
			It(`Invoke PostPartitionView successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostPartitionView(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostPartitionViewOptions model
				postPartitionViewOptionsModel := new(cloudantv1.PostPartitionViewOptions)
				postPartitionViewOptionsModel.Db = core.StringPtr("testString")
				postPartitionViewOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionViewOptionsModel.Ddoc = core.StringPtr("testString")
				postPartitionViewOptionsModel.View = core.StringPtr("testString")
				postPartitionViewOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postPartitionViewOptionsModel.Attachments = core.BoolPtr(true)
				postPartitionViewOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionViewOptionsModel.Descending = core.BoolPtr(true)
				postPartitionViewOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionViewOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postPartitionViewOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionViewOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionViewOptionsModel.UpdateSeq = core.BoolPtr(true)
				postPartitionViewOptionsModel.Endkey = core.StringPtr("testString")
				postPartitionViewOptionsModel.EndkeyDocid = core.StringPtr("testString")
				postPartitionViewOptionsModel.Group = core.BoolPtr(true)
				postPartitionViewOptionsModel.GroupLevel = core.Int64Ptr(int64(1))
				postPartitionViewOptionsModel.Key = core.StringPtr("testString")
				postPartitionViewOptionsModel.Keys = []interface{}{"testString"}
				postPartitionViewOptionsModel.Reduce = core.BoolPtr(true)
				postPartitionViewOptionsModel.Stable = core.BoolPtr(true)
				postPartitionViewOptionsModel.Startkey = core.StringPtr("testString")
				postPartitionViewOptionsModel.StartkeyDocid = core.StringPtr("testString")
				postPartitionViewOptionsModel.Update = core.StringPtr("true")
				postPartitionViewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostPartitionView(postPartitionViewOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostPartitionView with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionViewOptions model
				postPartitionViewOptionsModel := new(cloudantv1.PostPartitionViewOptions)
				postPartitionViewOptionsModel.Db = core.StringPtr("testString")
				postPartitionViewOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionViewOptionsModel.Ddoc = core.StringPtr("testString")
				postPartitionViewOptionsModel.View = core.StringPtr("testString")
				postPartitionViewOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postPartitionViewOptionsModel.Attachments = core.BoolPtr(true)
				postPartitionViewOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionViewOptionsModel.Descending = core.BoolPtr(true)
				postPartitionViewOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionViewOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postPartitionViewOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionViewOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionViewOptionsModel.UpdateSeq = core.BoolPtr(true)
				postPartitionViewOptionsModel.Endkey = core.StringPtr("testString")
				postPartitionViewOptionsModel.EndkeyDocid = core.StringPtr("testString")
				postPartitionViewOptionsModel.Group = core.BoolPtr(true)
				postPartitionViewOptionsModel.GroupLevel = core.Int64Ptr(int64(1))
				postPartitionViewOptionsModel.Key = core.StringPtr("testString")
				postPartitionViewOptionsModel.Keys = []interface{}{"testString"}
				postPartitionViewOptionsModel.Reduce = core.BoolPtr(true)
				postPartitionViewOptionsModel.Stable = core.BoolPtr(true)
				postPartitionViewOptionsModel.Startkey = core.StringPtr("testString")
				postPartitionViewOptionsModel.StartkeyDocid = core.StringPtr("testString")
				postPartitionViewOptionsModel.Update = core.StringPtr("true")
				postPartitionViewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostPartitionView(postPartitionViewOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostPartitionViewOptions model with no property values
				postPartitionViewOptionsModelNew := new(cloudantv1.PostPartitionViewOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostPartitionView(postPartitionViewOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostPartitionViewAsStream(postPartitionViewOptions *PostPartitionViewOptions)`, func() {
		postPartitionViewAsStreamPath := "/testString/_partition/testString/_design/testString/_view/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionViewAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostPartitionViewAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostPartitionViewAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostPartitionViewOptions model
				postPartitionViewOptionsModel := new(cloudantv1.PostPartitionViewOptions)
				postPartitionViewOptionsModel.Db = core.StringPtr("testString")
				postPartitionViewOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionViewOptionsModel.Ddoc = core.StringPtr("testString")
				postPartitionViewOptionsModel.View = core.StringPtr("testString")
				postPartitionViewOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postPartitionViewOptionsModel.Attachments = core.BoolPtr(true)
				postPartitionViewOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionViewOptionsModel.Descending = core.BoolPtr(true)
				postPartitionViewOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionViewOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postPartitionViewOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionViewOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionViewOptionsModel.UpdateSeq = core.BoolPtr(true)
				postPartitionViewOptionsModel.Endkey = core.StringPtr("testString")
				postPartitionViewOptionsModel.EndkeyDocid = core.StringPtr("testString")
				postPartitionViewOptionsModel.Group = core.BoolPtr(true)
				postPartitionViewOptionsModel.GroupLevel = core.Int64Ptr(int64(1))
				postPartitionViewOptionsModel.Key = core.StringPtr("testString")
				postPartitionViewOptionsModel.Keys = []interface{}{"testString"}
				postPartitionViewOptionsModel.Reduce = core.BoolPtr(true)
				postPartitionViewOptionsModel.Stable = core.BoolPtr(true)
				postPartitionViewOptionsModel.Startkey = core.StringPtr("testString")
				postPartitionViewOptionsModel.StartkeyDocid = core.StringPtr("testString")
				postPartitionViewOptionsModel.Update = core.StringPtr("true")
				postPartitionViewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostPartitionViewAsStream(postPartitionViewOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostPartitionViewAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionViewOptions model
				postPartitionViewOptionsModel := new(cloudantv1.PostPartitionViewOptions)
				postPartitionViewOptionsModel.Db = core.StringPtr("testString")
				postPartitionViewOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionViewOptionsModel.Ddoc = core.StringPtr("testString")
				postPartitionViewOptionsModel.View = core.StringPtr("testString")
				postPartitionViewOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postPartitionViewOptionsModel.Attachments = core.BoolPtr(true)
				postPartitionViewOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionViewOptionsModel.Descending = core.BoolPtr(true)
				postPartitionViewOptionsModel.IncludeDocs = core.BoolPtr(true)
				postPartitionViewOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postPartitionViewOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionViewOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionViewOptionsModel.UpdateSeq = core.BoolPtr(true)
				postPartitionViewOptionsModel.Endkey = core.StringPtr("testString")
				postPartitionViewOptionsModel.EndkeyDocid = core.StringPtr("testString")
				postPartitionViewOptionsModel.Group = core.BoolPtr(true)
				postPartitionViewOptionsModel.GroupLevel = core.Int64Ptr(int64(1))
				postPartitionViewOptionsModel.Key = core.StringPtr("testString")
				postPartitionViewOptionsModel.Keys = []interface{}{"testString"}
				postPartitionViewOptionsModel.Reduce = core.BoolPtr(true)
				postPartitionViewOptionsModel.Stable = core.BoolPtr(true)
				postPartitionViewOptionsModel.Startkey = core.StringPtr("testString")
				postPartitionViewOptionsModel.StartkeyDocid = core.StringPtr("testString")
				postPartitionViewOptionsModel.Update = core.StringPtr("true")
				postPartitionViewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostPartitionViewAsStream(postPartitionViewOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostPartitionViewOptions model with no property values
				postPartitionViewOptionsModelNew := new(cloudantv1.PostPartitionViewOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostPartitionViewAsStream(postPartitionViewOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostPartitionFind(postPartitionFindOptions *PostPartitionFindOptions) - Operation response error`, func() {
		postPartitionFindPath := "/testString/_partition/testString/_find"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionFindPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostPartitionFind with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionFindOptions model
				postPartitionFindOptionsModel := new(cloudantv1.PostPartitionFindOptions)
				postPartitionFindOptionsModel.Db = core.StringPtr("testString")
				postPartitionFindOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionFindOptionsModel.Bookmark = core.StringPtr("testString")
				postPartitionFindOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionFindOptionsModel.ExecutionStats = core.BoolPtr(true)
				postPartitionFindOptionsModel.Fields = []string{"testString"}
				postPartitionFindOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionFindOptionsModel.Selector = make(map[string]interface{})
				postPartitionFindOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionFindOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postPartitionFindOptionsModel.Stable = core.BoolPtr(true)
				postPartitionFindOptionsModel.Update = core.StringPtr("false")
				postPartitionFindOptionsModel.UseIndex = []string{"testString"}
				postPartitionFindOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostPartitionFind(postPartitionFindOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostPartitionFind(postPartitionFindOptions *PostPartitionFindOptions)`, func() {
		postPartitionFindPath := "/testString/_partition/testString/_find"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionFindPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bookmark": "Bookmark", "docs": [{"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}], "execution_stats": {"execution_time_ms": 15, "results_returned": 0, "total_docs_examined": 0, "total_keys_examined": 0, "total_quorum_docs_examined": 0}, "warning": "Warning"}`)
				}))
			})
			It(`Invoke PostPartitionFind successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostPartitionFind(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostPartitionFindOptions model
				postPartitionFindOptionsModel := new(cloudantv1.PostPartitionFindOptions)
				postPartitionFindOptionsModel.Db = core.StringPtr("testString")
				postPartitionFindOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionFindOptionsModel.Bookmark = core.StringPtr("testString")
				postPartitionFindOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionFindOptionsModel.ExecutionStats = core.BoolPtr(true)
				postPartitionFindOptionsModel.Fields = []string{"testString"}
				postPartitionFindOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionFindOptionsModel.Selector = make(map[string]interface{})
				postPartitionFindOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionFindOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postPartitionFindOptionsModel.Stable = core.BoolPtr(true)
				postPartitionFindOptionsModel.Update = core.StringPtr("false")
				postPartitionFindOptionsModel.UseIndex = []string{"testString"}
				postPartitionFindOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostPartitionFind(postPartitionFindOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostPartitionFind with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionFindOptions model
				postPartitionFindOptionsModel := new(cloudantv1.PostPartitionFindOptions)
				postPartitionFindOptionsModel.Db = core.StringPtr("testString")
				postPartitionFindOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionFindOptionsModel.Bookmark = core.StringPtr("testString")
				postPartitionFindOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionFindOptionsModel.ExecutionStats = core.BoolPtr(true)
				postPartitionFindOptionsModel.Fields = []string{"testString"}
				postPartitionFindOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionFindOptionsModel.Selector = make(map[string]interface{})
				postPartitionFindOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionFindOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postPartitionFindOptionsModel.Stable = core.BoolPtr(true)
				postPartitionFindOptionsModel.Update = core.StringPtr("false")
				postPartitionFindOptionsModel.UseIndex = []string{"testString"}
				postPartitionFindOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostPartitionFind(postPartitionFindOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostPartitionFindOptions model with no property values
				postPartitionFindOptionsModelNew := new(cloudantv1.PostPartitionFindOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostPartitionFind(postPartitionFindOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostPartitionFindAsStream(postPartitionFindOptions *PostPartitionFindOptions)`, func() {
		postPartitionFindAsStreamPath := "/testString/_partition/testString/_find"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postPartitionFindAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostPartitionFindAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostPartitionFindAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostPartitionFindOptions model
				postPartitionFindOptionsModel := new(cloudantv1.PostPartitionFindOptions)
				postPartitionFindOptionsModel.Db = core.StringPtr("testString")
				postPartitionFindOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionFindOptionsModel.Bookmark = core.StringPtr("testString")
				postPartitionFindOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionFindOptionsModel.ExecutionStats = core.BoolPtr(true)
				postPartitionFindOptionsModel.Fields = []string{"testString"}
				postPartitionFindOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionFindOptionsModel.Selector = make(map[string]interface{})
				postPartitionFindOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionFindOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postPartitionFindOptionsModel.Stable = core.BoolPtr(true)
				postPartitionFindOptionsModel.Update = core.StringPtr("false")
				postPartitionFindOptionsModel.UseIndex = []string{"testString"}
				postPartitionFindOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostPartitionFindAsStream(postPartitionFindOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostPartitionFindAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostPartitionFindOptions model
				postPartitionFindOptionsModel := new(cloudantv1.PostPartitionFindOptions)
				postPartitionFindOptionsModel.Db = core.StringPtr("testString")
				postPartitionFindOptionsModel.PartitionKey = core.StringPtr("testString")
				postPartitionFindOptionsModel.Bookmark = core.StringPtr("testString")
				postPartitionFindOptionsModel.Conflicts = core.BoolPtr(true)
				postPartitionFindOptionsModel.ExecutionStats = core.BoolPtr(true)
				postPartitionFindOptionsModel.Fields = []string{"testString"}
				postPartitionFindOptionsModel.Limit = core.Int64Ptr(int64(0))
				postPartitionFindOptionsModel.Selector = make(map[string]interface{})
				postPartitionFindOptionsModel.Skip = core.Int64Ptr(int64(0))
				postPartitionFindOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postPartitionFindOptionsModel.Stable = core.BoolPtr(true)
				postPartitionFindOptionsModel.Update = core.StringPtr("false")
				postPartitionFindOptionsModel.UseIndex = []string{"testString"}
				postPartitionFindOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostPartitionFindAsStream(postPartitionFindOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostPartitionFindOptions model with no property values
				postPartitionFindOptionsModelNew := new(cloudantv1.PostPartitionFindOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostPartitionFindAsStream(postPartitionFindOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`PostExplain(postExplainOptions *PostExplainOptions) - Operation response error`, func() {
		postExplainPath := "/testString/_explain"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postExplainPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostExplain with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostExplainOptions model
				postExplainOptionsModel := new(cloudantv1.PostExplainOptions)
				postExplainOptionsModel.Db = core.StringPtr("testString")
				postExplainOptionsModel.Bookmark = core.StringPtr("testString")
				postExplainOptionsModel.Conflicts = core.BoolPtr(true)
				postExplainOptionsModel.ExecutionStats = core.BoolPtr(true)
				postExplainOptionsModel.Fields = []string{"testString"}
				postExplainOptionsModel.Limit = core.Int64Ptr(int64(0))
				postExplainOptionsModel.Selector = make(map[string]interface{})
				postExplainOptionsModel.Skip = core.Int64Ptr(int64(0))
				postExplainOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postExplainOptionsModel.Stable = core.BoolPtr(true)
				postExplainOptionsModel.Update = core.StringPtr("false")
				postExplainOptionsModel.UseIndex = []string{"testString"}
				postExplainOptionsModel.R = core.Int64Ptr(int64(1))
				postExplainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostExplain(postExplainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostExplain(postExplainOptions *PostExplainOptions)`, func() {
		postExplainPath := "/testString/_explain"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postExplainPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dbname": "Dbname", "fields": ["Fields"], "index": {"ddoc": "Ddoc", "def": {"default_analyzer": {"name": "classic", "stopwords": ["Stopwords"]}, "default_field": {"analyzer": {"name": "classic", "stopwords": ["Stopwords"]}, "enabled": false}, "fields": [{"name": "Name", "type": "boolean"}], "index_array_lengths": false}, "name": "Name", "type": "json"}, "limit": 0, "opts": {"mapKey": "anyValue"}, "range": {"end_key": ["anyValue"], "start_key": ["anyValue"]}, "selector": {"mapKey": "anyValue"}, "skip": 0}`)
				}))
			})
			It(`Invoke PostExplain successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostExplain(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostExplainOptions model
				postExplainOptionsModel := new(cloudantv1.PostExplainOptions)
				postExplainOptionsModel.Db = core.StringPtr("testString")
				postExplainOptionsModel.Bookmark = core.StringPtr("testString")
				postExplainOptionsModel.Conflicts = core.BoolPtr(true)
				postExplainOptionsModel.ExecutionStats = core.BoolPtr(true)
				postExplainOptionsModel.Fields = []string{"testString"}
				postExplainOptionsModel.Limit = core.Int64Ptr(int64(0))
				postExplainOptionsModel.Selector = make(map[string]interface{})
				postExplainOptionsModel.Skip = core.Int64Ptr(int64(0))
				postExplainOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postExplainOptionsModel.Stable = core.BoolPtr(true)
				postExplainOptionsModel.Update = core.StringPtr("false")
				postExplainOptionsModel.UseIndex = []string{"testString"}
				postExplainOptionsModel.R = core.Int64Ptr(int64(1))
				postExplainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostExplain(postExplainOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostExplain with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostExplainOptions model
				postExplainOptionsModel := new(cloudantv1.PostExplainOptions)
				postExplainOptionsModel.Db = core.StringPtr("testString")
				postExplainOptionsModel.Bookmark = core.StringPtr("testString")
				postExplainOptionsModel.Conflicts = core.BoolPtr(true)
				postExplainOptionsModel.ExecutionStats = core.BoolPtr(true)
				postExplainOptionsModel.Fields = []string{"testString"}
				postExplainOptionsModel.Limit = core.Int64Ptr(int64(0))
				postExplainOptionsModel.Selector = make(map[string]interface{})
				postExplainOptionsModel.Skip = core.Int64Ptr(int64(0))
				postExplainOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postExplainOptionsModel.Stable = core.BoolPtr(true)
				postExplainOptionsModel.Update = core.StringPtr("false")
				postExplainOptionsModel.UseIndex = []string{"testString"}
				postExplainOptionsModel.R = core.Int64Ptr(int64(1))
				postExplainOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostExplain(postExplainOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostExplainOptions model with no property values
				postExplainOptionsModelNew := new(cloudantv1.PostExplainOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostExplain(postExplainOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostFind(postFindOptions *PostFindOptions) - Operation response error`, func() {
		postFindPath := "/testString/_find"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postFindPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostFind with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostFindOptions model
				postFindOptionsModel := new(cloudantv1.PostFindOptions)
				postFindOptionsModel.Db = core.StringPtr("testString")
				postFindOptionsModel.Bookmark = core.StringPtr("testString")
				postFindOptionsModel.Conflicts = core.BoolPtr(true)
				postFindOptionsModel.ExecutionStats = core.BoolPtr(true)
				postFindOptionsModel.Fields = []string{"testString"}
				postFindOptionsModel.Limit = core.Int64Ptr(int64(0))
				postFindOptionsModel.Selector = make(map[string]interface{})
				postFindOptionsModel.Skip = core.Int64Ptr(int64(0))
				postFindOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postFindOptionsModel.Stable = core.BoolPtr(true)
				postFindOptionsModel.Update = core.StringPtr("false")
				postFindOptionsModel.UseIndex = []string{"testString"}
				postFindOptionsModel.R = core.Int64Ptr(int64(1))
				postFindOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostFind(postFindOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostFind(postFindOptions *PostFindOptions)`, func() {
		postFindPath := "/testString/_find"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postFindPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bookmark": "Bookmark", "docs": [{"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}], "execution_stats": {"execution_time_ms": 15, "results_returned": 0, "total_docs_examined": 0, "total_keys_examined": 0, "total_quorum_docs_examined": 0}, "warning": "Warning"}`)
				}))
			})
			It(`Invoke PostFind successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostFind(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostFindOptions model
				postFindOptionsModel := new(cloudantv1.PostFindOptions)
				postFindOptionsModel.Db = core.StringPtr("testString")
				postFindOptionsModel.Bookmark = core.StringPtr("testString")
				postFindOptionsModel.Conflicts = core.BoolPtr(true)
				postFindOptionsModel.ExecutionStats = core.BoolPtr(true)
				postFindOptionsModel.Fields = []string{"testString"}
				postFindOptionsModel.Limit = core.Int64Ptr(int64(0))
				postFindOptionsModel.Selector = make(map[string]interface{})
				postFindOptionsModel.Skip = core.Int64Ptr(int64(0))
				postFindOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postFindOptionsModel.Stable = core.BoolPtr(true)
				postFindOptionsModel.Update = core.StringPtr("false")
				postFindOptionsModel.UseIndex = []string{"testString"}
				postFindOptionsModel.R = core.Int64Ptr(int64(1))
				postFindOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostFind(postFindOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostFind with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostFindOptions model
				postFindOptionsModel := new(cloudantv1.PostFindOptions)
				postFindOptionsModel.Db = core.StringPtr("testString")
				postFindOptionsModel.Bookmark = core.StringPtr("testString")
				postFindOptionsModel.Conflicts = core.BoolPtr(true)
				postFindOptionsModel.ExecutionStats = core.BoolPtr(true)
				postFindOptionsModel.Fields = []string{"testString"}
				postFindOptionsModel.Limit = core.Int64Ptr(int64(0))
				postFindOptionsModel.Selector = make(map[string]interface{})
				postFindOptionsModel.Skip = core.Int64Ptr(int64(0))
				postFindOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postFindOptionsModel.Stable = core.BoolPtr(true)
				postFindOptionsModel.Update = core.StringPtr("false")
				postFindOptionsModel.UseIndex = []string{"testString"}
				postFindOptionsModel.R = core.Int64Ptr(int64(1))
				postFindOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostFind(postFindOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostFindOptions model with no property values
				postFindOptionsModelNew := new(cloudantv1.PostFindOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostFind(postFindOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostFindAsStream(postFindOptions *PostFindOptions)`, func() {
		postFindAsStreamPath := "/testString/_find"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postFindAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostFindAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostFindAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostFindOptions model
				postFindOptionsModel := new(cloudantv1.PostFindOptions)
				postFindOptionsModel.Db = core.StringPtr("testString")
				postFindOptionsModel.Bookmark = core.StringPtr("testString")
				postFindOptionsModel.Conflicts = core.BoolPtr(true)
				postFindOptionsModel.ExecutionStats = core.BoolPtr(true)
				postFindOptionsModel.Fields = []string{"testString"}
				postFindOptionsModel.Limit = core.Int64Ptr(int64(0))
				postFindOptionsModel.Selector = make(map[string]interface{})
				postFindOptionsModel.Skip = core.Int64Ptr(int64(0))
				postFindOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postFindOptionsModel.Stable = core.BoolPtr(true)
				postFindOptionsModel.Update = core.StringPtr("false")
				postFindOptionsModel.UseIndex = []string{"testString"}
				postFindOptionsModel.R = core.Int64Ptr(int64(1))
				postFindOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostFindAsStream(postFindOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostFindAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostFindOptions model
				postFindOptionsModel := new(cloudantv1.PostFindOptions)
				postFindOptionsModel.Db = core.StringPtr("testString")
				postFindOptionsModel.Bookmark = core.StringPtr("testString")
				postFindOptionsModel.Conflicts = core.BoolPtr(true)
				postFindOptionsModel.ExecutionStats = core.BoolPtr(true)
				postFindOptionsModel.Fields = []string{"testString"}
				postFindOptionsModel.Limit = core.Int64Ptr(int64(0))
				postFindOptionsModel.Selector = make(map[string]interface{})
				postFindOptionsModel.Skip = core.Int64Ptr(int64(0))
				postFindOptionsModel.Sort = []map[string]string{make(map[string]string)}
				postFindOptionsModel.Stable = core.BoolPtr(true)
				postFindOptionsModel.Update = core.StringPtr("false")
				postFindOptionsModel.UseIndex = []string{"testString"}
				postFindOptionsModel.R = core.Int64Ptr(int64(1))
				postFindOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostFindAsStream(postFindOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostFindOptions model with no property values
				postFindOptionsModelNew := new(cloudantv1.PostFindOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostFindAsStream(postFindOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetIndexesInformation(getIndexesInformationOptions *GetIndexesInformationOptions) - Operation response error`, func() {
		getIndexesInformationPath := "/testString/_index"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIndexesInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetIndexesInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetIndexesInformationOptions model
				getIndexesInformationOptionsModel := new(cloudantv1.GetIndexesInformationOptions)
				getIndexesInformationOptionsModel.Db = core.StringPtr("testString")
				getIndexesInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetIndexesInformation(getIndexesInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetIndexesInformation(getIndexesInformationOptions *GetIndexesInformationOptions)`, func() {
		getIndexesInformationPath := "/testString/_index"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIndexesInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "indexes": [{"ddoc": "Ddoc", "def": {"default_analyzer": {"name": "classic", "stopwords": ["Stopwords"]}, "default_field": {"analyzer": {"name": "classic", "stopwords": ["Stopwords"]}, "enabled": false}, "fields": [{"name": "Name", "type": "boolean"}], "index_array_lengths": false}, "name": "Name", "type": "json"}]}`)
				}))
			})
			It(`Invoke GetIndexesInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetIndexesInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetIndexesInformationOptions model
				getIndexesInformationOptionsModel := new(cloudantv1.GetIndexesInformationOptions)
				getIndexesInformationOptionsModel.Db = core.StringPtr("testString")
				getIndexesInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetIndexesInformation(getIndexesInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetIndexesInformation with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetIndexesInformationOptions model
				getIndexesInformationOptionsModel := new(cloudantv1.GetIndexesInformationOptions)
				getIndexesInformationOptionsModel.Db = core.StringPtr("testString")
				getIndexesInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetIndexesInformation(getIndexesInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetIndexesInformationOptions model with no property values
				getIndexesInformationOptionsModelNew := new(cloudantv1.GetIndexesInformationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetIndexesInformation(getIndexesInformationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostIndex(postIndexOptions *PostIndexOptions) - Operation response error`, func() {
		postIndexPath := "/testString/_index"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postIndexPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostIndex with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Analyzer model
				analyzerModel := new(cloudantv1.Analyzer)
				analyzerModel.Name = core.StringPtr("classic")
				analyzerModel.Stopwords = []string{"testString"}

				// Construct an instance of the IndexTextOperatorDefaultField model
				indexTextOperatorDefaultFieldModel := new(cloudantv1.IndexTextOperatorDefaultField)
				indexTextOperatorDefaultFieldModel.Analyzer = analyzerModel
				indexTextOperatorDefaultFieldModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the IndexField model
				indexFieldModel := new(cloudantv1.IndexField)
				indexFieldModel.Name = core.StringPtr("testString")
				indexFieldModel.Type = core.StringPtr("boolean")
				indexFieldModel.SetProperty("foo", core.StringPtr("asc"))

				// Construct an instance of the IndexDefinition model
				indexDefinitionModel := new(cloudantv1.IndexDefinition)
				indexDefinitionModel.DefaultAnalyzer = analyzerModel
				indexDefinitionModel.DefaultField = indexTextOperatorDefaultFieldModel
				indexDefinitionModel.Fields = []cloudantv1.IndexField{*indexFieldModel}
				indexDefinitionModel.IndexArrayLengths = core.BoolPtr(true)

				// Construct an instance of the PostIndexOptions model
				postIndexOptionsModel := new(cloudantv1.PostIndexOptions)
				postIndexOptionsModel.Db = core.StringPtr("testString")
				postIndexOptionsModel.Ddoc = core.StringPtr("testString")
				postIndexOptionsModel.Def = indexDefinitionModel
				postIndexOptionsModel.Index = indexDefinitionModel
				postIndexOptionsModel.Name = core.StringPtr("testString")
				postIndexOptionsModel.PartialFilterSelector = make(map[string]interface{})
				postIndexOptionsModel.Partitioned = core.BoolPtr(true)
				postIndexOptionsModel.Type = core.StringPtr("json")
				postIndexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostIndex(postIndexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostIndex(postIndexOptions *PostIndexOptions)`, func() {
		postIndexPath := "/testString/_index"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postIndexPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "result": "created"}`)
				}))
			})
			It(`Invoke PostIndex successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostIndex(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Analyzer model
				analyzerModel := new(cloudantv1.Analyzer)
				analyzerModel.Name = core.StringPtr("classic")
				analyzerModel.Stopwords = []string{"testString"}

				// Construct an instance of the IndexTextOperatorDefaultField model
				indexTextOperatorDefaultFieldModel := new(cloudantv1.IndexTextOperatorDefaultField)
				indexTextOperatorDefaultFieldModel.Analyzer = analyzerModel
				indexTextOperatorDefaultFieldModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the IndexField model
				indexFieldModel := new(cloudantv1.IndexField)
				indexFieldModel.Name = core.StringPtr("testString")
				indexFieldModel.Type = core.StringPtr("boolean")
				indexFieldModel.SetProperty("foo", core.StringPtr("asc"))

				// Construct an instance of the IndexDefinition model
				indexDefinitionModel := new(cloudantv1.IndexDefinition)
				indexDefinitionModel.DefaultAnalyzer = analyzerModel
				indexDefinitionModel.DefaultField = indexTextOperatorDefaultFieldModel
				indexDefinitionModel.Fields = []cloudantv1.IndexField{*indexFieldModel}
				indexDefinitionModel.IndexArrayLengths = core.BoolPtr(true)

				// Construct an instance of the PostIndexOptions model
				postIndexOptionsModel := new(cloudantv1.PostIndexOptions)
				postIndexOptionsModel.Db = core.StringPtr("testString")
				postIndexOptionsModel.Ddoc = core.StringPtr("testString")
				postIndexOptionsModel.Def = indexDefinitionModel
				postIndexOptionsModel.Index = indexDefinitionModel
				postIndexOptionsModel.Name = core.StringPtr("testString")
				postIndexOptionsModel.PartialFilterSelector = make(map[string]interface{})
				postIndexOptionsModel.Partitioned = core.BoolPtr(true)
				postIndexOptionsModel.Type = core.StringPtr("json")
				postIndexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostIndex(postIndexOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostIndex with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Analyzer model
				analyzerModel := new(cloudantv1.Analyzer)
				analyzerModel.Name = core.StringPtr("classic")
				analyzerModel.Stopwords = []string{"testString"}

				// Construct an instance of the IndexTextOperatorDefaultField model
				indexTextOperatorDefaultFieldModel := new(cloudantv1.IndexTextOperatorDefaultField)
				indexTextOperatorDefaultFieldModel.Analyzer = analyzerModel
				indexTextOperatorDefaultFieldModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the IndexField model
				indexFieldModel := new(cloudantv1.IndexField)
				indexFieldModel.Name = core.StringPtr("testString")
				indexFieldModel.Type = core.StringPtr("boolean")
				indexFieldModel.SetProperty("foo", core.StringPtr("asc"))

				// Construct an instance of the IndexDefinition model
				indexDefinitionModel := new(cloudantv1.IndexDefinition)
				indexDefinitionModel.DefaultAnalyzer = analyzerModel
				indexDefinitionModel.DefaultField = indexTextOperatorDefaultFieldModel
				indexDefinitionModel.Fields = []cloudantv1.IndexField{*indexFieldModel}
				indexDefinitionModel.IndexArrayLengths = core.BoolPtr(true)

				// Construct an instance of the PostIndexOptions model
				postIndexOptionsModel := new(cloudantv1.PostIndexOptions)
				postIndexOptionsModel.Db = core.StringPtr("testString")
				postIndexOptionsModel.Ddoc = core.StringPtr("testString")
				postIndexOptionsModel.Def = indexDefinitionModel
				postIndexOptionsModel.Index = indexDefinitionModel
				postIndexOptionsModel.Name = core.StringPtr("testString")
				postIndexOptionsModel.PartialFilterSelector = make(map[string]interface{})
				postIndexOptionsModel.Partitioned = core.BoolPtr(true)
				postIndexOptionsModel.Type = core.StringPtr("json")
				postIndexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostIndex(postIndexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostIndexOptions model with no property values
				postIndexOptionsModelNew := new(cloudantv1.PostIndexOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostIndex(postIndexOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteIndex(deleteIndexOptions *DeleteIndexOptions) - Operation response error`, func() {
		deleteIndexPath := "/testString/_index/_design/testString/json/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteIndexPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteIndex with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteIndexOptions model
				deleteIndexOptionsModel := new(cloudantv1.DeleteIndexOptions)
				deleteIndexOptionsModel.Db = core.StringPtr("testString")
				deleteIndexOptionsModel.Ddoc = core.StringPtr("testString")
				deleteIndexOptionsModel.Type = core.StringPtr("json")
				deleteIndexOptionsModel.Index = core.StringPtr("testString")
				deleteIndexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.DeleteIndex(deleteIndexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteIndex(deleteIndexOptions *DeleteIndexOptions)`, func() {
		deleteIndexPath := "/testString/_index/_design/testString/json/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteIndexPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ok": true}`)
				}))
			})
			It(`Invoke DeleteIndex successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.DeleteIndex(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteIndexOptions model
				deleteIndexOptionsModel := new(cloudantv1.DeleteIndexOptions)
				deleteIndexOptionsModel.Db = core.StringPtr("testString")
				deleteIndexOptionsModel.Ddoc = core.StringPtr("testString")
				deleteIndexOptionsModel.Type = core.StringPtr("json")
				deleteIndexOptionsModel.Index = core.StringPtr("testString")
				deleteIndexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.DeleteIndex(deleteIndexOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteIndex with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteIndexOptions model
				deleteIndexOptionsModel := new(cloudantv1.DeleteIndexOptions)
				deleteIndexOptionsModel.Db = core.StringPtr("testString")
				deleteIndexOptionsModel.Ddoc = core.StringPtr("testString")
				deleteIndexOptionsModel.Type = core.StringPtr("json")
				deleteIndexOptionsModel.Index = core.StringPtr("testString")
				deleteIndexOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.DeleteIndex(deleteIndexOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteIndexOptions model with no property values
				deleteIndexOptionsModelNew := new(cloudantv1.DeleteIndexOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.DeleteIndex(deleteIndexOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`PostSearchAnalyze(postSearchAnalyzeOptions *PostSearchAnalyzeOptions) - Operation response error`, func() {
		postSearchAnalyzePath := "/_search_analyze"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postSearchAnalyzePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostSearchAnalyze with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostSearchAnalyzeOptions model
				postSearchAnalyzeOptionsModel := new(cloudantv1.PostSearchAnalyzeOptions)
				postSearchAnalyzeOptionsModel.Analyzer = core.StringPtr("arabic")
				postSearchAnalyzeOptionsModel.Text = core.StringPtr("testString")
				postSearchAnalyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostSearchAnalyze(postSearchAnalyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostSearchAnalyze(postSearchAnalyzeOptions *PostSearchAnalyzeOptions)`, func() {
		postSearchAnalyzePath := "/_search_analyze"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postSearchAnalyzePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tokens": ["Tokens"]}`)
				}))
			})
			It(`Invoke PostSearchAnalyze successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostSearchAnalyze(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostSearchAnalyzeOptions model
				postSearchAnalyzeOptionsModel := new(cloudantv1.PostSearchAnalyzeOptions)
				postSearchAnalyzeOptionsModel.Analyzer = core.StringPtr("arabic")
				postSearchAnalyzeOptionsModel.Text = core.StringPtr("testString")
				postSearchAnalyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostSearchAnalyze(postSearchAnalyzeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostSearchAnalyze with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostSearchAnalyzeOptions model
				postSearchAnalyzeOptionsModel := new(cloudantv1.PostSearchAnalyzeOptions)
				postSearchAnalyzeOptionsModel.Analyzer = core.StringPtr("arabic")
				postSearchAnalyzeOptionsModel.Text = core.StringPtr("testString")
				postSearchAnalyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostSearchAnalyze(postSearchAnalyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostSearch(postSearchOptions *PostSearchOptions) - Operation response error`, func() {
		postSearchPath := "/testString/_design/testString/_search/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postSearchPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostSearch with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostSearchOptions model
				postSearchOptionsModel := new(cloudantv1.PostSearchOptions)
				postSearchOptionsModel.Db = core.StringPtr("testString")
				postSearchOptionsModel.Ddoc = core.StringPtr("testString")
				postSearchOptionsModel.Index = core.StringPtr("testString")
				postSearchOptionsModel.Bookmark = core.StringPtr("testString")
				postSearchOptionsModel.HighlightFields = []string{"testString"}
				postSearchOptionsModel.HighlightNumber = core.Int64Ptr(int64(1))
				postSearchOptionsModel.HighlightPostTag = core.StringPtr("testString")
				postSearchOptionsModel.HighlightPreTag = core.StringPtr("testString")
				postSearchOptionsModel.HighlightSize = core.Int64Ptr(int64(1))
				postSearchOptionsModel.IncludeDocs = core.BoolPtr(true)
				postSearchOptionsModel.IncludeFields = []string{"testString"}
				postSearchOptionsModel.Limit = core.Int64Ptr(int64(3))
				postSearchOptionsModel.Query = core.StringPtr("testString")
				postSearchOptionsModel.Sort = []string{"testString"}
				postSearchOptionsModel.Stale = core.StringPtr("ok")
				postSearchOptionsModel.Counts = []string{"testString"}
				postSearchOptionsModel.Drilldown = [][]string{[]string{"testString"}}
				postSearchOptionsModel.GroupField = core.StringPtr("testString")
				postSearchOptionsModel.GroupLimit = core.Int64Ptr(int64(1))
				postSearchOptionsModel.GroupSort = []string{"testString"}
				postSearchOptionsModel.Ranges = make(map[string]map[string]map[string]string)
				postSearchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostSearch(postSearchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostSearch(postSearchOptions *PostSearchOptions)`, func() {
		postSearchPath := "/testString/_design/testString/_search/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postSearchPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "bookmark": "Bookmark", "by": "By", "counts": {"mapKey": {"mapKey": 0}}, "ranges": {"mapKey": {"mapKey": 0}}, "rows": [{"doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "fields": {"mapKey": "anyValue"}, "highlights": {"mapKey": ["Inner"]}, "id": "ID"}], "groups": [{"total_rows": 0, "bookmark": "Bookmark", "by": "By", "counts": {"mapKey": {"mapKey": 0}}, "ranges": {"mapKey": {"mapKey": 0}}, "rows": [{"doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "fields": {"mapKey": "anyValue"}, "highlights": {"mapKey": ["Inner"]}, "id": "ID"}]}]}`)
				}))
			})
			It(`Invoke PostSearch successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostSearch(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostSearchOptions model
				postSearchOptionsModel := new(cloudantv1.PostSearchOptions)
				postSearchOptionsModel.Db = core.StringPtr("testString")
				postSearchOptionsModel.Ddoc = core.StringPtr("testString")
				postSearchOptionsModel.Index = core.StringPtr("testString")
				postSearchOptionsModel.Bookmark = core.StringPtr("testString")
				postSearchOptionsModel.HighlightFields = []string{"testString"}
				postSearchOptionsModel.HighlightNumber = core.Int64Ptr(int64(1))
				postSearchOptionsModel.HighlightPostTag = core.StringPtr("testString")
				postSearchOptionsModel.HighlightPreTag = core.StringPtr("testString")
				postSearchOptionsModel.HighlightSize = core.Int64Ptr(int64(1))
				postSearchOptionsModel.IncludeDocs = core.BoolPtr(true)
				postSearchOptionsModel.IncludeFields = []string{"testString"}
				postSearchOptionsModel.Limit = core.Int64Ptr(int64(3))
				postSearchOptionsModel.Query = core.StringPtr("testString")
				postSearchOptionsModel.Sort = []string{"testString"}
				postSearchOptionsModel.Stale = core.StringPtr("ok")
				postSearchOptionsModel.Counts = []string{"testString"}
				postSearchOptionsModel.Drilldown = [][]string{[]string{"testString"}}
				postSearchOptionsModel.GroupField = core.StringPtr("testString")
				postSearchOptionsModel.GroupLimit = core.Int64Ptr(int64(1))
				postSearchOptionsModel.GroupSort = []string{"testString"}
				postSearchOptionsModel.Ranges = make(map[string]map[string]map[string]string)
				postSearchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostSearch(postSearchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostSearch with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostSearchOptions model
				postSearchOptionsModel := new(cloudantv1.PostSearchOptions)
				postSearchOptionsModel.Db = core.StringPtr("testString")
				postSearchOptionsModel.Ddoc = core.StringPtr("testString")
				postSearchOptionsModel.Index = core.StringPtr("testString")
				postSearchOptionsModel.Bookmark = core.StringPtr("testString")
				postSearchOptionsModel.HighlightFields = []string{"testString"}
				postSearchOptionsModel.HighlightNumber = core.Int64Ptr(int64(1))
				postSearchOptionsModel.HighlightPostTag = core.StringPtr("testString")
				postSearchOptionsModel.HighlightPreTag = core.StringPtr("testString")
				postSearchOptionsModel.HighlightSize = core.Int64Ptr(int64(1))
				postSearchOptionsModel.IncludeDocs = core.BoolPtr(true)
				postSearchOptionsModel.IncludeFields = []string{"testString"}
				postSearchOptionsModel.Limit = core.Int64Ptr(int64(3))
				postSearchOptionsModel.Query = core.StringPtr("testString")
				postSearchOptionsModel.Sort = []string{"testString"}
				postSearchOptionsModel.Stale = core.StringPtr("ok")
				postSearchOptionsModel.Counts = []string{"testString"}
				postSearchOptionsModel.Drilldown = [][]string{[]string{"testString"}}
				postSearchOptionsModel.GroupField = core.StringPtr("testString")
				postSearchOptionsModel.GroupLimit = core.Int64Ptr(int64(1))
				postSearchOptionsModel.GroupSort = []string{"testString"}
				postSearchOptionsModel.Ranges = make(map[string]map[string]map[string]string)
				postSearchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostSearch(postSearchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostSearchOptions model with no property values
				postSearchOptionsModelNew := new(cloudantv1.PostSearchOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostSearch(postSearchOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostSearchAsStream(postSearchOptions *PostSearchOptions)`, func() {
		postSearchAsStreamPath := "/testString/_design/testString/_search/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postSearchAsStreamPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke PostSearchAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostSearchAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostSearchOptions model
				postSearchOptionsModel := new(cloudantv1.PostSearchOptions)
				postSearchOptionsModel.Db = core.StringPtr("testString")
				postSearchOptionsModel.Ddoc = core.StringPtr("testString")
				postSearchOptionsModel.Index = core.StringPtr("testString")
				postSearchOptionsModel.Bookmark = core.StringPtr("testString")
				postSearchOptionsModel.HighlightFields = []string{"testString"}
				postSearchOptionsModel.HighlightNumber = core.Int64Ptr(int64(1))
				postSearchOptionsModel.HighlightPostTag = core.StringPtr("testString")
				postSearchOptionsModel.HighlightPreTag = core.StringPtr("testString")
				postSearchOptionsModel.HighlightSize = core.Int64Ptr(int64(1))
				postSearchOptionsModel.IncludeDocs = core.BoolPtr(true)
				postSearchOptionsModel.IncludeFields = []string{"testString"}
				postSearchOptionsModel.Limit = core.Int64Ptr(int64(3))
				postSearchOptionsModel.Query = core.StringPtr("testString")
				postSearchOptionsModel.Sort = []string{"testString"}
				postSearchOptionsModel.Stale = core.StringPtr("ok")
				postSearchOptionsModel.Counts = []string{"testString"}
				postSearchOptionsModel.Drilldown = [][]string{[]string{"testString"}}
				postSearchOptionsModel.GroupField = core.StringPtr("testString")
				postSearchOptionsModel.GroupLimit = core.Int64Ptr(int64(1))
				postSearchOptionsModel.GroupSort = []string{"testString"}
				postSearchOptionsModel.Ranges = make(map[string]map[string]map[string]string)
				postSearchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostSearchAsStream(postSearchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke PostSearchAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostSearchOptions model
				postSearchOptionsModel := new(cloudantv1.PostSearchOptions)
				postSearchOptionsModel.Db = core.StringPtr("testString")
				postSearchOptionsModel.Ddoc = core.StringPtr("testString")
				postSearchOptionsModel.Index = core.StringPtr("testString")
				postSearchOptionsModel.Bookmark = core.StringPtr("testString")
				postSearchOptionsModel.HighlightFields = []string{"testString"}
				postSearchOptionsModel.HighlightNumber = core.Int64Ptr(int64(1))
				postSearchOptionsModel.HighlightPostTag = core.StringPtr("testString")
				postSearchOptionsModel.HighlightPreTag = core.StringPtr("testString")
				postSearchOptionsModel.HighlightSize = core.Int64Ptr(int64(1))
				postSearchOptionsModel.IncludeDocs = core.BoolPtr(true)
				postSearchOptionsModel.IncludeFields = []string{"testString"}
				postSearchOptionsModel.Limit = core.Int64Ptr(int64(3))
				postSearchOptionsModel.Query = core.StringPtr("testString")
				postSearchOptionsModel.Sort = []string{"testString"}
				postSearchOptionsModel.Stale = core.StringPtr("ok")
				postSearchOptionsModel.Counts = []string{"testString"}
				postSearchOptionsModel.Drilldown = [][]string{[]string{"testString"}}
				postSearchOptionsModel.GroupField = core.StringPtr("testString")
				postSearchOptionsModel.GroupLimit = core.Int64Ptr(int64(1))
				postSearchOptionsModel.GroupSort = []string{"testString"}
				postSearchOptionsModel.Ranges = make(map[string]map[string]map[string]string)
				postSearchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostSearchAsStream(postSearchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostSearchOptions model with no property values
				postSearchOptionsModelNew := new(cloudantv1.PostSearchOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostSearchAsStream(postSearchOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSearchInfo(getSearchInfoOptions *GetSearchInfoOptions) - Operation response error`, func() {
		getSearchInfoPath := "/testString/_design/testString/_search_info/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSearchInfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSearchInfo with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSearchInfoOptions model
				getSearchInfoOptionsModel := new(cloudantv1.GetSearchInfoOptions)
				getSearchInfoOptionsModel.Db = core.StringPtr("testString")
				getSearchInfoOptionsModel.Ddoc = core.StringPtr("testString")
				getSearchInfoOptionsModel.Index = core.StringPtr("testString")
				getSearchInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetSearchInfo(getSearchInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSearchInfo(getSearchInfoOptions *GetSearchInfoOptions)`, func() {
		getSearchInfoPath := "/testString/_design/testString/_search_info/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSearchInfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "search_index": {"committed_seq": 12, "disk_size": 0, "doc_count": 0, "doc_del_count": 0, "pending_seq": 10}}`)
				}))
			})
			It(`Invoke GetSearchInfo successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetSearchInfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSearchInfoOptions model
				getSearchInfoOptionsModel := new(cloudantv1.GetSearchInfoOptions)
				getSearchInfoOptionsModel.Db = core.StringPtr("testString")
				getSearchInfoOptionsModel.Ddoc = core.StringPtr("testString")
				getSearchInfoOptionsModel.Index = core.StringPtr("testString")
				getSearchInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetSearchInfo(getSearchInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSearchInfo with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSearchInfoOptions model
				getSearchInfoOptionsModel := new(cloudantv1.GetSearchInfoOptions)
				getSearchInfoOptionsModel.Db = core.StringPtr("testString")
				getSearchInfoOptionsModel.Ddoc = core.StringPtr("testString")
				getSearchInfoOptionsModel.Index = core.StringPtr("testString")
				getSearchInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetSearchInfo(getSearchInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSearchInfoOptions model with no property values
				getSearchInfoOptionsModelNew := new(cloudantv1.GetSearchInfoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetSearchInfo(getSearchInfoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetGeo(getGeoOptions *GetGeoOptions) - Operation response error`, func() {
		getGeoPath := "/testString/_design/testString/_geo/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGeoPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["bbox"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["bookmark"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["format"]).To(Equal([]string{"legacy"}))

					Expect(req.URL.Query()["g"]).To(Equal([]string{"testString"}))


					// TODO: Add check for include_docs query parameter


					// TODO: Add check for lat query parameter


					// TODO: Add check for limit query parameter


					// TODO: Add check for lon query parameter


					// TODO: Add check for nearest query parameter


					// TODO: Add check for radius query parameter


					// TODO: Add check for rangex query parameter


					// TODO: Add check for rangey query parameter

					Expect(req.URL.Query()["relation"]).To(Equal([]string{"contains"}))


					// TODO: Add check for skip query parameter

					Expect(req.URL.Query()["stale"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGeo with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetGeoOptions model
				getGeoOptionsModel := new(cloudantv1.GetGeoOptions)
				getGeoOptionsModel.Db = core.StringPtr("testString")
				getGeoOptionsModel.Ddoc = core.StringPtr("testString")
				getGeoOptionsModel.Index = core.StringPtr("testString")
				getGeoOptionsModel.Bbox = core.StringPtr("testString")
				getGeoOptionsModel.Bookmark = core.StringPtr("testString")
				getGeoOptionsModel.Format = core.StringPtr("legacy")
				getGeoOptionsModel.G = core.StringPtr("testString")
				getGeoOptionsModel.IncludeDocs = core.BoolPtr(true)
				getGeoOptionsModel.Lat = core.Float64Ptr(float64(-90))
				getGeoOptionsModel.Limit = core.Int64Ptr(int64(0))
				getGeoOptionsModel.Lon = core.Float64Ptr(float64(-180))
				getGeoOptionsModel.Nearest = core.BoolPtr(true)
				getGeoOptionsModel.Radius = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Rangex = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Rangey = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Relation = core.StringPtr("contains")
				getGeoOptionsModel.Skip = core.Int64Ptr(int64(0))
				getGeoOptionsModel.Stale = core.StringPtr("ok")
				getGeoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetGeo(getGeoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetGeo(getGeoOptions *GetGeoOptions)`, func() {
		getGeoPath := "/testString/_design/testString/_geo/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGeoPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["bbox"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["bookmark"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["format"]).To(Equal([]string{"legacy"}))

					Expect(req.URL.Query()["g"]).To(Equal([]string{"testString"}))


					// TODO: Add check for include_docs query parameter


					// TODO: Add check for lat query parameter


					// TODO: Add check for limit query parameter


					// TODO: Add check for lon query parameter


					// TODO: Add check for nearest query parameter


					// TODO: Add check for radius query parameter


					// TODO: Add check for rangex query parameter


					// TODO: Add check for rangey query parameter

					Expect(req.URL.Query()["relation"]).To(Equal([]string{"contains"}))


					// TODO: Add check for skip query parameter

					Expect(req.URL.Query()["stale"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"bookmark": "Bookmark", "features": [{"_id": "ID", "_rev": "Rev", "bbox": [4], "geometry": {"type": "Point", "coordinates": ["anyValue"]}, "properties": {"mapKey": "anyValue"}, "type": "Feature"}], "rows": [{"doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "geometry": {"type": "Point", "coordinates": ["anyValue"]}, "id": "ID", "rev": "Rev"}], "type": "FeatureCollection"}`)
				}))
			})
			It(`Invoke GetGeo successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetGeo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGeoOptions model
				getGeoOptionsModel := new(cloudantv1.GetGeoOptions)
				getGeoOptionsModel.Db = core.StringPtr("testString")
				getGeoOptionsModel.Ddoc = core.StringPtr("testString")
				getGeoOptionsModel.Index = core.StringPtr("testString")
				getGeoOptionsModel.Bbox = core.StringPtr("testString")
				getGeoOptionsModel.Bookmark = core.StringPtr("testString")
				getGeoOptionsModel.Format = core.StringPtr("legacy")
				getGeoOptionsModel.G = core.StringPtr("testString")
				getGeoOptionsModel.IncludeDocs = core.BoolPtr(true)
				getGeoOptionsModel.Lat = core.Float64Ptr(float64(-90))
				getGeoOptionsModel.Limit = core.Int64Ptr(int64(0))
				getGeoOptionsModel.Lon = core.Float64Ptr(float64(-180))
				getGeoOptionsModel.Nearest = core.BoolPtr(true)
				getGeoOptionsModel.Radius = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Rangex = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Rangey = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Relation = core.StringPtr("contains")
				getGeoOptionsModel.Skip = core.Int64Ptr(int64(0))
				getGeoOptionsModel.Stale = core.StringPtr("ok")
				getGeoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetGeo(getGeoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetGeo with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetGeoOptions model
				getGeoOptionsModel := new(cloudantv1.GetGeoOptions)
				getGeoOptionsModel.Db = core.StringPtr("testString")
				getGeoOptionsModel.Ddoc = core.StringPtr("testString")
				getGeoOptionsModel.Index = core.StringPtr("testString")
				getGeoOptionsModel.Bbox = core.StringPtr("testString")
				getGeoOptionsModel.Bookmark = core.StringPtr("testString")
				getGeoOptionsModel.Format = core.StringPtr("legacy")
				getGeoOptionsModel.G = core.StringPtr("testString")
				getGeoOptionsModel.IncludeDocs = core.BoolPtr(true)
				getGeoOptionsModel.Lat = core.Float64Ptr(float64(-90))
				getGeoOptionsModel.Limit = core.Int64Ptr(int64(0))
				getGeoOptionsModel.Lon = core.Float64Ptr(float64(-180))
				getGeoOptionsModel.Nearest = core.BoolPtr(true)
				getGeoOptionsModel.Radius = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Rangex = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Rangey = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Relation = core.StringPtr("contains")
				getGeoOptionsModel.Skip = core.Int64Ptr(int64(0))
				getGeoOptionsModel.Stale = core.StringPtr("ok")
				getGeoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetGeo(getGeoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGeoOptions model with no property values
				getGeoOptionsModelNew := new(cloudantv1.GetGeoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetGeo(getGeoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetGeoAsStream(getGeoOptions *GetGeoOptions)`, func() {
		getGeoAsStreamPath := "/testString/_design/testString/_geo/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGeoAsStreamPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["bbox"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["bookmark"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["format"]).To(Equal([]string{"legacy"}))

					Expect(req.URL.Query()["g"]).To(Equal([]string{"testString"}))


					// TODO: Add check for include_docs query parameter


					// TODO: Add check for lat query parameter


					// TODO: Add check for limit query parameter


					// TODO: Add check for lon query parameter


					// TODO: Add check for nearest query parameter


					// TODO: Add check for radius query parameter


					// TODO: Add check for rangex query parameter


					// TODO: Add check for rangey query parameter

					Expect(req.URL.Query()["relation"]).To(Equal([]string{"contains"}))


					// TODO: Add check for skip query parameter

					Expect(req.URL.Query()["stale"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"foo": "this is a mock response for JSON streaming"}`)
				}))
			})
			It(`Invoke GetGeoAsStream successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetGeoAsStream(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGeoOptions model
				getGeoOptionsModel := new(cloudantv1.GetGeoOptions)
				getGeoOptionsModel.Db = core.StringPtr("testString")
				getGeoOptionsModel.Ddoc = core.StringPtr("testString")
				getGeoOptionsModel.Index = core.StringPtr("testString")
				getGeoOptionsModel.Bbox = core.StringPtr("testString")
				getGeoOptionsModel.Bookmark = core.StringPtr("testString")
				getGeoOptionsModel.Format = core.StringPtr("legacy")
				getGeoOptionsModel.G = core.StringPtr("testString")
				getGeoOptionsModel.IncludeDocs = core.BoolPtr(true)
				getGeoOptionsModel.Lat = core.Float64Ptr(float64(-90))
				getGeoOptionsModel.Limit = core.Int64Ptr(int64(0))
				getGeoOptionsModel.Lon = core.Float64Ptr(float64(-180))
				getGeoOptionsModel.Nearest = core.BoolPtr(true)
				getGeoOptionsModel.Radius = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Rangex = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Rangey = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Relation = core.StringPtr("contains")
				getGeoOptionsModel.Skip = core.Int64Ptr(int64(0))
				getGeoOptionsModel.Stale = core.StringPtr("ok")
				getGeoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetGeoAsStream(getGeoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Verify streamed JSON response.
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(string(buffer)).To(Equal(`{"foo": "this is a mock response for JSON streaming"}`))
			})
			It(`Invoke GetGeoAsStream with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetGeoOptions model
				getGeoOptionsModel := new(cloudantv1.GetGeoOptions)
				getGeoOptionsModel.Db = core.StringPtr("testString")
				getGeoOptionsModel.Ddoc = core.StringPtr("testString")
				getGeoOptionsModel.Index = core.StringPtr("testString")
				getGeoOptionsModel.Bbox = core.StringPtr("testString")
				getGeoOptionsModel.Bookmark = core.StringPtr("testString")
				getGeoOptionsModel.Format = core.StringPtr("legacy")
				getGeoOptionsModel.G = core.StringPtr("testString")
				getGeoOptionsModel.IncludeDocs = core.BoolPtr(true)
				getGeoOptionsModel.Lat = core.Float64Ptr(float64(-90))
				getGeoOptionsModel.Limit = core.Int64Ptr(int64(0))
				getGeoOptionsModel.Lon = core.Float64Ptr(float64(-180))
				getGeoOptionsModel.Nearest = core.BoolPtr(true)
				getGeoOptionsModel.Radius = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Rangex = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Rangey = core.Float64Ptr(float64(0))
				getGeoOptionsModel.Relation = core.StringPtr("contains")
				getGeoOptionsModel.Skip = core.Int64Ptr(int64(0))
				getGeoOptionsModel.Stale = core.StringPtr("ok")
				getGeoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetGeoAsStream(getGeoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGeoOptions model with no property values
				getGeoOptionsModelNew := new(cloudantv1.GetGeoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetGeoAsStream(getGeoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostGeoCleanup(postGeoCleanupOptions *PostGeoCleanupOptions) - Operation response error`, func() {
		postGeoCleanupPath := "/testString/_geo_cleanup"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postGeoCleanupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostGeoCleanup with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostGeoCleanupOptions model
				postGeoCleanupOptionsModel := new(cloudantv1.PostGeoCleanupOptions)
				postGeoCleanupOptionsModel.Db = core.StringPtr("testString")
				postGeoCleanupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostGeoCleanup(postGeoCleanupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostGeoCleanup(postGeoCleanupOptions *PostGeoCleanupOptions)`, func() {
		postGeoCleanupPath := "/testString/_geo_cleanup"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postGeoCleanupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"ok": true}`)
				}))
			})
			It(`Invoke PostGeoCleanup successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostGeoCleanup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostGeoCleanupOptions model
				postGeoCleanupOptionsModel := new(cloudantv1.PostGeoCleanupOptions)
				postGeoCleanupOptionsModel.Db = core.StringPtr("testString")
				postGeoCleanupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostGeoCleanup(postGeoCleanupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostGeoCleanup with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostGeoCleanupOptions model
				postGeoCleanupOptionsModel := new(cloudantv1.PostGeoCleanupOptions)
				postGeoCleanupOptionsModel.Db = core.StringPtr("testString")
				postGeoCleanupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostGeoCleanup(postGeoCleanupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostGeoCleanupOptions model with no property values
				postGeoCleanupOptionsModelNew := new(cloudantv1.PostGeoCleanupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostGeoCleanup(postGeoCleanupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGeoIndexInformation(getGeoIndexInformationOptions *GetGeoIndexInformationOptions) - Operation response error`, func() {
		getGeoIndexInformationPath := "/testString/_design/testString/_geo_info/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGeoIndexInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGeoIndexInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetGeoIndexInformationOptions model
				getGeoIndexInformationOptionsModel := new(cloudantv1.GetGeoIndexInformationOptions)
				getGeoIndexInformationOptionsModel.Db = core.StringPtr("testString")
				getGeoIndexInformationOptionsModel.Ddoc = core.StringPtr("testString")
				getGeoIndexInformationOptionsModel.Index = core.StringPtr("testString")
				getGeoIndexInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetGeoIndexInformation(getGeoIndexInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetGeoIndexInformation(getGeoIndexInformationOptions *GetGeoIndexInformationOptions)`, func() {
		getGeoIndexInformationPath := "/testString/_design/testString/_geo_info/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGeoIndexInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"geo_index": {"data_size": 0, "disk_size": 0, "doc_count": 0}}`)
				}))
			})
			It(`Invoke GetGeoIndexInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetGeoIndexInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGeoIndexInformationOptions model
				getGeoIndexInformationOptionsModel := new(cloudantv1.GetGeoIndexInformationOptions)
				getGeoIndexInformationOptionsModel.Db = core.StringPtr("testString")
				getGeoIndexInformationOptionsModel.Ddoc = core.StringPtr("testString")
				getGeoIndexInformationOptionsModel.Index = core.StringPtr("testString")
				getGeoIndexInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetGeoIndexInformation(getGeoIndexInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetGeoIndexInformation with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetGeoIndexInformationOptions model
				getGeoIndexInformationOptionsModel := new(cloudantv1.GetGeoIndexInformationOptions)
				getGeoIndexInformationOptionsModel.Db = core.StringPtr("testString")
				getGeoIndexInformationOptionsModel.Ddoc = core.StringPtr("testString")
				getGeoIndexInformationOptionsModel.Index = core.StringPtr("testString")
				getGeoIndexInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetGeoIndexInformation(getGeoIndexInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGeoIndexInformationOptions model with no property values
				getGeoIndexInformationOptionsModelNew := new(cloudantv1.GetGeoIndexInformationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetGeoIndexInformation(getGeoIndexInformationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetDbUpdates(getDbUpdatesOptions *GetDbUpdatesOptions) - Operation response error`, func() {
		getDbUpdatesPath := "/_db_updates"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDbUpdatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["feed"]).To(Equal([]string{"continuous"}))


					// TODO: Add check for heartbeat query parameter


					// TODO: Add check for timeout query parameter

					Expect(req.URL.Query()["since"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDbUpdates with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDbUpdatesOptions model
				getDbUpdatesOptionsModel := new(cloudantv1.GetDbUpdatesOptions)
				getDbUpdatesOptionsModel.Feed = core.StringPtr("continuous")
				getDbUpdatesOptionsModel.Heartbeat = core.Int64Ptr(int64(0))
				getDbUpdatesOptionsModel.Timeout = core.Int64Ptr(int64(0))
				getDbUpdatesOptionsModel.Since = core.StringPtr("testString")
				getDbUpdatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetDbUpdates(getDbUpdatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDbUpdates(getDbUpdatesOptions *GetDbUpdatesOptions)`, func() {
		getDbUpdatesPath := "/_db_updates"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDbUpdatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["feed"]).To(Equal([]string{"continuous"}))


					// TODO: Add check for heartbeat query parameter


					// TODO: Add check for timeout query parameter

					Expect(req.URL.Query()["since"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"last_seq": "LastSeq", "results": [{"account": "Account", "dbname": "Dbname", "seq": "Seq", "type": "created"}]}`)
				}))
			})
			It(`Invoke GetDbUpdates successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetDbUpdates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDbUpdatesOptions model
				getDbUpdatesOptionsModel := new(cloudantv1.GetDbUpdatesOptions)
				getDbUpdatesOptionsModel.Feed = core.StringPtr("continuous")
				getDbUpdatesOptionsModel.Heartbeat = core.Int64Ptr(int64(0))
				getDbUpdatesOptionsModel.Timeout = core.Int64Ptr(int64(0))
				getDbUpdatesOptionsModel.Since = core.StringPtr("testString")
				getDbUpdatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetDbUpdates(getDbUpdatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDbUpdates with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDbUpdatesOptions model
				getDbUpdatesOptionsModel := new(cloudantv1.GetDbUpdatesOptions)
				getDbUpdatesOptionsModel.Feed = core.StringPtr("continuous")
				getDbUpdatesOptionsModel.Heartbeat = core.Int64Ptr(int64(0))
				getDbUpdatesOptionsModel.Timeout = core.Int64Ptr(int64(0))
				getDbUpdatesOptionsModel.Since = core.StringPtr("testString")
				getDbUpdatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetDbUpdates(getDbUpdatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`HeadReplicationDocument(headReplicationDocumentOptions *HeadReplicationDocumentOptions)`, func() {
		headReplicationDocumentPath := "/_replicator/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(headReplicationDocumentPath))
					Expect(req.Method).To(Equal("HEAD"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke HeadReplicationDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cloudantService.HeadReplicationDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HeadReplicationDocumentOptions model
				headReplicationDocumentOptionsModel := new(cloudantv1.HeadReplicationDocumentOptions)
				headReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				headReplicationDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				headReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cloudantService.HeadReplicationDocument(headReplicationDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HeadReplicationDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the HeadReplicationDocumentOptions model
				headReplicationDocumentOptionsModel := new(cloudantv1.HeadReplicationDocumentOptions)
				headReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				headReplicationDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				headReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cloudantService.HeadReplicationDocument(headReplicationDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HeadReplicationDocumentOptions model with no property values
				headReplicationDocumentOptionsModelNew := new(cloudantv1.HeadReplicationDocumentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cloudantService.HeadReplicationDocument(headReplicationDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`HeadSchedulerJob(headSchedulerJobOptions *HeadSchedulerJobOptions)`, func() {
		headSchedulerJobPath := "/_scheduler/jobs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(headSchedulerJobPath))
					Expect(req.Method).To(Equal("HEAD"))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke HeadSchedulerJob successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cloudantService.HeadSchedulerJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HeadSchedulerJobOptions model
				headSchedulerJobOptionsModel := new(cloudantv1.HeadSchedulerJobOptions)
				headSchedulerJobOptionsModel.JobID = core.StringPtr("testString")
				headSchedulerJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cloudantService.HeadSchedulerJob(headSchedulerJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HeadSchedulerJob with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the HeadSchedulerJobOptions model
				headSchedulerJobOptionsModel := new(cloudantv1.HeadSchedulerJobOptions)
				headSchedulerJobOptionsModel.JobID = core.StringPtr("testString")
				headSchedulerJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cloudantService.HeadSchedulerJob(headSchedulerJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HeadSchedulerJobOptions model with no property values
				headSchedulerJobOptionsModelNew := new(cloudantv1.HeadSchedulerJobOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cloudantService.HeadSchedulerJob(headSchedulerJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostReplicate(postReplicateOptions *PostReplicateOptions) - Operation response error`, func() {
		postReplicatePath := "/_replicate"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postReplicatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostReplicate with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the ReplicationCreateTargetParameters model
				replicationCreateTargetParametersModel := new(cloudantv1.ReplicationCreateTargetParameters)
				replicationCreateTargetParametersModel.N = core.Int64Ptr(int64(1))
				replicationCreateTargetParametersModel.Partitioned = core.BoolPtr(true)
				replicationCreateTargetParametersModel.Q = core.Int64Ptr(int64(1))

				// Construct an instance of the ReplicationDatabaseAuthIam model
				replicationDatabaseAuthIamModel := new(cloudantv1.ReplicationDatabaseAuthIam)
				replicationDatabaseAuthIamModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ReplicationDatabaseAuth model
				replicationDatabaseAuthModel := new(cloudantv1.ReplicationDatabaseAuth)
				replicationDatabaseAuthModel.Iam = replicationDatabaseAuthIamModel

				// Construct an instance of the ReplicationDatabase model
				replicationDatabaseModel := new(cloudantv1.ReplicationDatabase)
				replicationDatabaseModel.Auth = replicationDatabaseAuthModel
				replicationDatabaseModel.HeadersVar = make(map[string]string)
				replicationDatabaseModel.URL = core.StringPtr("testString")

				// Construct an instance of the UserContext model
				userContextModel := new(cloudantv1.UserContext)
				userContextModel.Db = core.StringPtr("testString")
				userContextModel.Name = core.StringPtr("testString")
				userContextModel.Roles = []string{"_reader"}

				// Construct an instance of the ReplicationDocument model
				replicationDocumentModel := new(cloudantv1.ReplicationDocument)
				replicationDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				replicationDocumentModel.Conflicts = []string{"testString"}
				replicationDocumentModel.Deleted = core.BoolPtr(true)
				replicationDocumentModel.DeletedConflicts = []string{"testString"}
				replicationDocumentModel.ID = core.StringPtr("testString")
				replicationDocumentModel.LocalSeq = core.StringPtr("testString")
				replicationDocumentModel.Rev = core.StringPtr("testString")
				replicationDocumentModel.Revisions = revisionsModel
				replicationDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				replicationDocumentModel.Cancel = core.BoolPtr(true)
				replicationDocumentModel.CheckpointInterval = core.Int64Ptr(int64(0))
				replicationDocumentModel.ConnectionTimeout = core.Int64Ptr(int64(0))
				replicationDocumentModel.Continuous = core.BoolPtr(true)
				replicationDocumentModel.CreateTarget = core.BoolPtr(true)
				replicationDocumentModel.CreateTargetParams = replicationCreateTargetParametersModel
				replicationDocumentModel.DocIds = []string{"testString"}
				replicationDocumentModel.Filter = core.StringPtr("testString")
				replicationDocumentModel.HTTPConnections = core.Int64Ptr(int64(1))
				replicationDocumentModel.QueryParams = make(map[string]string)
				replicationDocumentModel.RetriesPerRequest = core.Int64Ptr(int64(0))
				replicationDocumentModel.Selector = make(map[string]interface{})
				replicationDocumentModel.SinceSeq = core.StringPtr("testString")
				replicationDocumentModel.SocketOptions = core.StringPtr("testString")
				replicationDocumentModel.Source = replicationDatabaseModel
				replicationDocumentModel.SourceProxy = core.StringPtr("testString")
				replicationDocumentModel.Target = replicationDatabaseModel
				replicationDocumentModel.TargetProxy = core.StringPtr("testString")
				replicationDocumentModel.UseCheckpoints = core.BoolPtr(true)
				replicationDocumentModel.UserCtx = userContextModel
				replicationDocumentModel.WorkerBatchSize = core.Int64Ptr(int64(1))
				replicationDocumentModel.WorkerProcesses = core.Int64Ptr(int64(1))
				replicationDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				replicationDocumentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PostReplicateOptions model
				postReplicateOptionsModel := new(cloudantv1.PostReplicateOptions)
				postReplicateOptionsModel.ReplicationDocument = replicationDocumentModel
				postReplicateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostReplicate(postReplicateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostReplicate(postReplicateOptions *PostReplicateOptions)`, func() {
		postReplicatePath := "/_replicate"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postReplicatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"doc_write_failures": 0, "docs_read": 0, "docs_written": 0, "end_last_seq": "EndLastSeq", "end_time": "EndTime", "missing_checked": 0, "missing_found": 0, "recorded_seq": "RecordedSeq", "session_id": "SessionID", "start_last_seq": "StartLastSeq", "start_time": "StartTime"}], "ok": true, "replication_id_version": 0, "session_id": "SessionID", "source_last_seq": "SourceLastSeq"}`)
				}))
			})
			It(`Invoke PostReplicate successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostReplicate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the ReplicationCreateTargetParameters model
				replicationCreateTargetParametersModel := new(cloudantv1.ReplicationCreateTargetParameters)
				replicationCreateTargetParametersModel.N = core.Int64Ptr(int64(1))
				replicationCreateTargetParametersModel.Partitioned = core.BoolPtr(true)
				replicationCreateTargetParametersModel.Q = core.Int64Ptr(int64(1))

				// Construct an instance of the ReplicationDatabaseAuthIam model
				replicationDatabaseAuthIamModel := new(cloudantv1.ReplicationDatabaseAuthIam)
				replicationDatabaseAuthIamModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ReplicationDatabaseAuth model
				replicationDatabaseAuthModel := new(cloudantv1.ReplicationDatabaseAuth)
				replicationDatabaseAuthModel.Iam = replicationDatabaseAuthIamModel

				// Construct an instance of the ReplicationDatabase model
				replicationDatabaseModel := new(cloudantv1.ReplicationDatabase)
				replicationDatabaseModel.Auth = replicationDatabaseAuthModel
				replicationDatabaseModel.HeadersVar = make(map[string]string)
				replicationDatabaseModel.URL = core.StringPtr("testString")

				// Construct an instance of the UserContext model
				userContextModel := new(cloudantv1.UserContext)
				userContextModel.Db = core.StringPtr("testString")
				userContextModel.Name = core.StringPtr("testString")
				userContextModel.Roles = []string{"_reader"}

				// Construct an instance of the ReplicationDocument model
				replicationDocumentModel := new(cloudantv1.ReplicationDocument)
				replicationDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				replicationDocumentModel.Conflicts = []string{"testString"}
				replicationDocumentModel.Deleted = core.BoolPtr(true)
				replicationDocumentModel.DeletedConflicts = []string{"testString"}
				replicationDocumentModel.ID = core.StringPtr("testString")
				replicationDocumentModel.LocalSeq = core.StringPtr("testString")
				replicationDocumentModel.Rev = core.StringPtr("testString")
				replicationDocumentModel.Revisions = revisionsModel
				replicationDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				replicationDocumentModel.Cancel = core.BoolPtr(true)
				replicationDocumentModel.CheckpointInterval = core.Int64Ptr(int64(0))
				replicationDocumentModel.ConnectionTimeout = core.Int64Ptr(int64(0))
				replicationDocumentModel.Continuous = core.BoolPtr(true)
				replicationDocumentModel.CreateTarget = core.BoolPtr(true)
				replicationDocumentModel.CreateTargetParams = replicationCreateTargetParametersModel
				replicationDocumentModel.DocIds = []string{"testString"}
				replicationDocumentModel.Filter = core.StringPtr("testString")
				replicationDocumentModel.HTTPConnections = core.Int64Ptr(int64(1))
				replicationDocumentModel.QueryParams = make(map[string]string)
				replicationDocumentModel.RetriesPerRequest = core.Int64Ptr(int64(0))
				replicationDocumentModel.Selector = make(map[string]interface{})
				replicationDocumentModel.SinceSeq = core.StringPtr("testString")
				replicationDocumentModel.SocketOptions = core.StringPtr("testString")
				replicationDocumentModel.Source = replicationDatabaseModel
				replicationDocumentModel.SourceProxy = core.StringPtr("testString")
				replicationDocumentModel.Target = replicationDatabaseModel
				replicationDocumentModel.TargetProxy = core.StringPtr("testString")
				replicationDocumentModel.UseCheckpoints = core.BoolPtr(true)
				replicationDocumentModel.UserCtx = userContextModel
				replicationDocumentModel.WorkerBatchSize = core.Int64Ptr(int64(1))
				replicationDocumentModel.WorkerProcesses = core.Int64Ptr(int64(1))
				replicationDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				replicationDocumentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PostReplicateOptions model
				postReplicateOptionsModel := new(cloudantv1.PostReplicateOptions)
				postReplicateOptionsModel.ReplicationDocument = replicationDocumentModel
				postReplicateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostReplicate(postReplicateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostReplicate with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the ReplicationCreateTargetParameters model
				replicationCreateTargetParametersModel := new(cloudantv1.ReplicationCreateTargetParameters)
				replicationCreateTargetParametersModel.N = core.Int64Ptr(int64(1))
				replicationCreateTargetParametersModel.Partitioned = core.BoolPtr(true)
				replicationCreateTargetParametersModel.Q = core.Int64Ptr(int64(1))

				// Construct an instance of the ReplicationDatabaseAuthIam model
				replicationDatabaseAuthIamModel := new(cloudantv1.ReplicationDatabaseAuthIam)
				replicationDatabaseAuthIamModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ReplicationDatabaseAuth model
				replicationDatabaseAuthModel := new(cloudantv1.ReplicationDatabaseAuth)
				replicationDatabaseAuthModel.Iam = replicationDatabaseAuthIamModel

				// Construct an instance of the ReplicationDatabase model
				replicationDatabaseModel := new(cloudantv1.ReplicationDatabase)
				replicationDatabaseModel.Auth = replicationDatabaseAuthModel
				replicationDatabaseModel.HeadersVar = make(map[string]string)
				replicationDatabaseModel.URL = core.StringPtr("testString")

				// Construct an instance of the UserContext model
				userContextModel := new(cloudantv1.UserContext)
				userContextModel.Db = core.StringPtr("testString")
				userContextModel.Name = core.StringPtr("testString")
				userContextModel.Roles = []string{"_reader"}

				// Construct an instance of the ReplicationDocument model
				replicationDocumentModel := new(cloudantv1.ReplicationDocument)
				replicationDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				replicationDocumentModel.Conflicts = []string{"testString"}
				replicationDocumentModel.Deleted = core.BoolPtr(true)
				replicationDocumentModel.DeletedConflicts = []string{"testString"}
				replicationDocumentModel.ID = core.StringPtr("testString")
				replicationDocumentModel.LocalSeq = core.StringPtr("testString")
				replicationDocumentModel.Rev = core.StringPtr("testString")
				replicationDocumentModel.Revisions = revisionsModel
				replicationDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				replicationDocumentModel.Cancel = core.BoolPtr(true)
				replicationDocumentModel.CheckpointInterval = core.Int64Ptr(int64(0))
				replicationDocumentModel.ConnectionTimeout = core.Int64Ptr(int64(0))
				replicationDocumentModel.Continuous = core.BoolPtr(true)
				replicationDocumentModel.CreateTarget = core.BoolPtr(true)
				replicationDocumentModel.CreateTargetParams = replicationCreateTargetParametersModel
				replicationDocumentModel.DocIds = []string{"testString"}
				replicationDocumentModel.Filter = core.StringPtr("testString")
				replicationDocumentModel.HTTPConnections = core.Int64Ptr(int64(1))
				replicationDocumentModel.QueryParams = make(map[string]string)
				replicationDocumentModel.RetriesPerRequest = core.Int64Ptr(int64(0))
				replicationDocumentModel.Selector = make(map[string]interface{})
				replicationDocumentModel.SinceSeq = core.StringPtr("testString")
				replicationDocumentModel.SocketOptions = core.StringPtr("testString")
				replicationDocumentModel.Source = replicationDatabaseModel
				replicationDocumentModel.SourceProxy = core.StringPtr("testString")
				replicationDocumentModel.Target = replicationDatabaseModel
				replicationDocumentModel.TargetProxy = core.StringPtr("testString")
				replicationDocumentModel.UseCheckpoints = core.BoolPtr(true)
				replicationDocumentModel.UserCtx = userContextModel
				replicationDocumentModel.WorkerBatchSize = core.Int64Ptr(int64(1))
				replicationDocumentModel.WorkerProcesses = core.Int64Ptr(int64(1))
				replicationDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				replicationDocumentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PostReplicateOptions model
				postReplicateOptionsModel := new(cloudantv1.PostReplicateOptions)
				postReplicateOptionsModel.ReplicationDocument = replicationDocumentModel
				postReplicateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostReplicate(postReplicateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteReplicationDocument(deleteReplicationDocumentOptions *DeleteReplicationDocumentOptions) - Operation response error`, func() {
		deleteReplicationDocumentPath := "/_replicator/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteReplicationDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteReplicationDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteReplicationDocumentOptions model
				deleteReplicationDocumentOptionsModel := new(cloudantv1.DeleteReplicationDocumentOptions)
				deleteReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				deleteReplicationDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteReplicationDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteReplicationDocumentOptionsModel.Rev = core.StringPtr("testString")
				deleteReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.DeleteReplicationDocument(deleteReplicationDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteReplicationDocument(deleteReplicationDocumentOptions *DeleteReplicationDocumentOptions)`, func() {
		deleteReplicationDocumentPath := "/_replicator/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteReplicationDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke DeleteReplicationDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.DeleteReplicationDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteReplicationDocumentOptions model
				deleteReplicationDocumentOptionsModel := new(cloudantv1.DeleteReplicationDocumentOptions)
				deleteReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				deleteReplicationDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteReplicationDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteReplicationDocumentOptionsModel.Rev = core.StringPtr("testString")
				deleteReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.DeleteReplicationDocument(deleteReplicationDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteReplicationDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteReplicationDocumentOptions model
				deleteReplicationDocumentOptionsModel := new(cloudantv1.DeleteReplicationDocumentOptions)
				deleteReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				deleteReplicationDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteReplicationDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteReplicationDocumentOptionsModel.Rev = core.StringPtr("testString")
				deleteReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.DeleteReplicationDocument(deleteReplicationDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteReplicationDocumentOptions model with no property values
				deleteReplicationDocumentOptionsModelNew := new(cloudantv1.DeleteReplicationDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.DeleteReplicationDocument(deleteReplicationDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReplicationDocument(getReplicationDocumentOptions *GetReplicationDocumentOptions) - Operation response error`, func() {
		getReplicationDocumentPath := "/_replicator/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReplicationDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for deleted_conflicts query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for local_seq query parameter


					// TODO: Add check for meta query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))


					// TODO: Add check for revs query parameter


					// TODO: Add check for revs_info query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReplicationDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetReplicationDocumentOptions model
				getReplicationDocumentOptionsModel := new(cloudantv1.GetReplicationDocumentOptions)
				getReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				getReplicationDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getReplicationDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.AttsSince = []string{"testString"}
				getReplicationDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.Latest = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.Meta = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.OpenRevs = []string{"testString"}
				getReplicationDocumentOptionsModel.Rev = core.StringPtr("testString")
				getReplicationDocumentOptionsModel.Revs = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetReplicationDocument(getReplicationDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetReplicationDocument(getReplicationDocumentOptions *GetReplicationDocumentOptions)`, func() {
		getReplicationDocumentPath := "/_replicator/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReplicationDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for conflicts query parameter


					// TODO: Add check for deleted_conflicts query parameter


					// TODO: Add check for latest query parameter


					// TODO: Add check for local_seq query parameter


					// TODO: Add check for meta query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))


					// TODO: Add check for revs query parameter


					// TODO: Add check for revs_info query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}], "cancel": true, "checkpoint_interval": 0, "connection_timeout": 0, "continuous": true, "create_target": true, "create_target_params": {"n": 1, "partitioned": false, "q": 1}, "doc_ids": ["DocIds"], "filter": "Filter", "http_connections": 1, "query_params": {"mapKey": "Inner"}, "retries_per_request": 0, "selector": {"mapKey": {"anyKey": "anyValue"}}, "since_seq": "SinceSeq", "socket_options": "SocketOptions", "source": {"auth": {"iam": {"api_key": "ApiKey"}}, "headers": {"mapKey": "Inner"}, "url": "URL"}, "source_proxy": "SourceProxy", "target": {"auth": {"iam": {"api_key": "ApiKey"}}, "headers": {"mapKey": "Inner"}, "url": "URL"}, "target_proxy": "TargetProxy", "use_checkpoints": true, "user_ctx": {"db": "Db", "name": "Name", "roles": ["_reader"]}, "worker_batch_size": 1, "worker_processes": 1}`)
				}))
			})
			It(`Invoke GetReplicationDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetReplicationDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReplicationDocumentOptions model
				getReplicationDocumentOptionsModel := new(cloudantv1.GetReplicationDocumentOptions)
				getReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				getReplicationDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getReplicationDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.AttsSince = []string{"testString"}
				getReplicationDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.Latest = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.Meta = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.OpenRevs = []string{"testString"}
				getReplicationDocumentOptionsModel.Rev = core.StringPtr("testString")
				getReplicationDocumentOptionsModel.Revs = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetReplicationDocument(getReplicationDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetReplicationDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetReplicationDocumentOptions model
				getReplicationDocumentOptionsModel := new(cloudantv1.GetReplicationDocumentOptions)
				getReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				getReplicationDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getReplicationDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.AttsSince = []string{"testString"}
				getReplicationDocumentOptionsModel.Conflicts = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.DeletedConflicts = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.Latest = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.Meta = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.OpenRevs = []string{"testString"}
				getReplicationDocumentOptionsModel.Rev = core.StringPtr("testString")
				getReplicationDocumentOptionsModel.Revs = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.RevsInfo = core.BoolPtr(true)
				getReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetReplicationDocument(getReplicationDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReplicationDocumentOptions model with no property values
				getReplicationDocumentOptionsModelNew := new(cloudantv1.GetReplicationDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetReplicationDocument(getReplicationDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutReplicationDocument(putReplicationDocumentOptions *PutReplicationDocumentOptions) - Operation response error`, func() {
		putReplicationDocumentPath := "/_replicator/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putReplicationDocumentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))


					// TODO: Add check for new_edits query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutReplicationDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the ReplicationCreateTargetParameters model
				replicationCreateTargetParametersModel := new(cloudantv1.ReplicationCreateTargetParameters)
				replicationCreateTargetParametersModel.N = core.Int64Ptr(int64(1))
				replicationCreateTargetParametersModel.Partitioned = core.BoolPtr(true)
				replicationCreateTargetParametersModel.Q = core.Int64Ptr(int64(1))

				// Construct an instance of the ReplicationDatabaseAuthIam model
				replicationDatabaseAuthIamModel := new(cloudantv1.ReplicationDatabaseAuthIam)
				replicationDatabaseAuthIamModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ReplicationDatabaseAuth model
				replicationDatabaseAuthModel := new(cloudantv1.ReplicationDatabaseAuth)
				replicationDatabaseAuthModel.Iam = replicationDatabaseAuthIamModel

				// Construct an instance of the ReplicationDatabase model
				replicationDatabaseModel := new(cloudantv1.ReplicationDatabase)
				replicationDatabaseModel.Auth = replicationDatabaseAuthModel
				replicationDatabaseModel.HeadersVar = make(map[string]string)
				replicationDatabaseModel.URL = core.StringPtr("http://myserver.example:5984/foo-db")

				// Construct an instance of the UserContext model
				userContextModel := new(cloudantv1.UserContext)
				userContextModel.Db = core.StringPtr("testString")
				userContextModel.Name = core.StringPtr("john")
				userContextModel.Roles = []string{"_reader"}

				// Construct an instance of the ReplicationDocument model
				replicationDocumentModel := new(cloudantv1.ReplicationDocument)
				replicationDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				replicationDocumentModel.Conflicts = []string{"testString"}
				replicationDocumentModel.Deleted = core.BoolPtr(true)
				replicationDocumentModel.DeletedConflicts = []string{"testString"}
				replicationDocumentModel.ID = core.StringPtr("testString")
				replicationDocumentModel.LocalSeq = core.StringPtr("testString")
				replicationDocumentModel.Rev = core.StringPtr("testString")
				replicationDocumentModel.Revisions = revisionsModel
				replicationDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				replicationDocumentModel.Cancel = core.BoolPtr(false)
				replicationDocumentModel.CheckpointInterval = core.Int64Ptr(int64(4500))
				replicationDocumentModel.ConnectionTimeout = core.Int64Ptr(int64(15000))
				replicationDocumentModel.Continuous = core.BoolPtr(true)
				replicationDocumentModel.CreateTarget = core.BoolPtr(true)
				replicationDocumentModel.CreateTargetParams = replicationCreateTargetParametersModel
				replicationDocumentModel.DocIds = []string{"testString"}
				replicationDocumentModel.Filter = core.StringPtr("ddoc/my_filter")
				replicationDocumentModel.HTTPConnections = core.Int64Ptr(int64(10))
				replicationDocumentModel.QueryParams = make(map[string]string)
				replicationDocumentModel.RetriesPerRequest = core.Int64Ptr(int64(3))
				replicationDocumentModel.Selector = make(map[string]interface{})
				replicationDocumentModel.SinceSeq = core.StringPtr("34-g1AAAAGjeJzLYWBgYMlgTmGQT0lKzi9KdU")
				replicationDocumentModel.SocketOptions = core.StringPtr("[{keepalive, true}, {nodelay, false}]")
				replicationDocumentModel.Source = replicationDatabaseModel
				replicationDocumentModel.SourceProxy = core.StringPtr("http://my-source-proxy.example:8888")
				replicationDocumentModel.Target = replicationDatabaseModel
				replicationDocumentModel.TargetProxy = core.StringPtr("http://my-target-proxy.example:8888")
				replicationDocumentModel.UseCheckpoints = core.BoolPtr(false)
				replicationDocumentModel.UserCtx = userContextModel
				replicationDocumentModel.WorkerBatchSize = core.Int64Ptr(int64(400))
				replicationDocumentModel.WorkerProcesses = core.Int64Ptr(int64(3))
				replicationDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				replicationDocumentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PutReplicationDocumentOptions model
				putReplicationDocumentOptionsModel := new(cloudantv1.PutReplicationDocumentOptions)
				putReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				putReplicationDocumentOptionsModel.ReplicationDocument = replicationDocumentModel
				putReplicationDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				putReplicationDocumentOptionsModel.Batch = core.StringPtr("ok")
				putReplicationDocumentOptionsModel.NewEdits = core.BoolPtr(true)
				putReplicationDocumentOptionsModel.Rev = core.StringPtr("testString")
				putReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PutReplicationDocument(putReplicationDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutReplicationDocument(putReplicationDocumentOptions *PutReplicationDocumentOptions)`, func() {
		putReplicationDocumentPath := "/_replicator/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putReplicationDocumentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))


					// TODO: Add check for new_edits query parameter

					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke PutReplicationDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PutReplicationDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the ReplicationCreateTargetParameters model
				replicationCreateTargetParametersModel := new(cloudantv1.ReplicationCreateTargetParameters)
				replicationCreateTargetParametersModel.N = core.Int64Ptr(int64(1))
				replicationCreateTargetParametersModel.Partitioned = core.BoolPtr(true)
				replicationCreateTargetParametersModel.Q = core.Int64Ptr(int64(1))

				// Construct an instance of the ReplicationDatabaseAuthIam model
				replicationDatabaseAuthIamModel := new(cloudantv1.ReplicationDatabaseAuthIam)
				replicationDatabaseAuthIamModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ReplicationDatabaseAuth model
				replicationDatabaseAuthModel := new(cloudantv1.ReplicationDatabaseAuth)
				replicationDatabaseAuthModel.Iam = replicationDatabaseAuthIamModel

				// Construct an instance of the ReplicationDatabase model
				replicationDatabaseModel := new(cloudantv1.ReplicationDatabase)
				replicationDatabaseModel.Auth = replicationDatabaseAuthModel
				replicationDatabaseModel.HeadersVar = make(map[string]string)
				replicationDatabaseModel.URL = core.StringPtr("http://myserver.example:5984/foo-db")

				// Construct an instance of the UserContext model
				userContextModel := new(cloudantv1.UserContext)
				userContextModel.Db = core.StringPtr("testString")
				userContextModel.Name = core.StringPtr("john")
				userContextModel.Roles = []string{"_reader"}

				// Construct an instance of the ReplicationDocument model
				replicationDocumentModel := new(cloudantv1.ReplicationDocument)
				replicationDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				replicationDocumentModel.Conflicts = []string{"testString"}
				replicationDocumentModel.Deleted = core.BoolPtr(true)
				replicationDocumentModel.DeletedConflicts = []string{"testString"}
				replicationDocumentModel.ID = core.StringPtr("testString")
				replicationDocumentModel.LocalSeq = core.StringPtr("testString")
				replicationDocumentModel.Rev = core.StringPtr("testString")
				replicationDocumentModel.Revisions = revisionsModel
				replicationDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				replicationDocumentModel.Cancel = core.BoolPtr(false)
				replicationDocumentModel.CheckpointInterval = core.Int64Ptr(int64(4500))
				replicationDocumentModel.ConnectionTimeout = core.Int64Ptr(int64(15000))
				replicationDocumentModel.Continuous = core.BoolPtr(true)
				replicationDocumentModel.CreateTarget = core.BoolPtr(true)
				replicationDocumentModel.CreateTargetParams = replicationCreateTargetParametersModel
				replicationDocumentModel.DocIds = []string{"testString"}
				replicationDocumentModel.Filter = core.StringPtr("ddoc/my_filter")
				replicationDocumentModel.HTTPConnections = core.Int64Ptr(int64(10))
				replicationDocumentModel.QueryParams = make(map[string]string)
				replicationDocumentModel.RetriesPerRequest = core.Int64Ptr(int64(3))
				replicationDocumentModel.Selector = make(map[string]interface{})
				replicationDocumentModel.SinceSeq = core.StringPtr("34-g1AAAAGjeJzLYWBgYMlgTmGQT0lKzi9KdU")
				replicationDocumentModel.SocketOptions = core.StringPtr("[{keepalive, true}, {nodelay, false}]")
				replicationDocumentModel.Source = replicationDatabaseModel
				replicationDocumentModel.SourceProxy = core.StringPtr("http://my-source-proxy.example:8888")
				replicationDocumentModel.Target = replicationDatabaseModel
				replicationDocumentModel.TargetProxy = core.StringPtr("http://my-target-proxy.example:8888")
				replicationDocumentModel.UseCheckpoints = core.BoolPtr(false)
				replicationDocumentModel.UserCtx = userContextModel
				replicationDocumentModel.WorkerBatchSize = core.Int64Ptr(int64(400))
				replicationDocumentModel.WorkerProcesses = core.Int64Ptr(int64(3))
				replicationDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				replicationDocumentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PutReplicationDocumentOptions model
				putReplicationDocumentOptionsModel := new(cloudantv1.PutReplicationDocumentOptions)
				putReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				putReplicationDocumentOptionsModel.ReplicationDocument = replicationDocumentModel
				putReplicationDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				putReplicationDocumentOptionsModel.Batch = core.StringPtr("ok")
				putReplicationDocumentOptionsModel.NewEdits = core.BoolPtr(true)
				putReplicationDocumentOptionsModel.Rev = core.StringPtr("testString")
				putReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PutReplicationDocument(putReplicationDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PutReplicationDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the ReplicationCreateTargetParameters model
				replicationCreateTargetParametersModel := new(cloudantv1.ReplicationCreateTargetParameters)
				replicationCreateTargetParametersModel.N = core.Int64Ptr(int64(1))
				replicationCreateTargetParametersModel.Partitioned = core.BoolPtr(true)
				replicationCreateTargetParametersModel.Q = core.Int64Ptr(int64(1))

				// Construct an instance of the ReplicationDatabaseAuthIam model
				replicationDatabaseAuthIamModel := new(cloudantv1.ReplicationDatabaseAuthIam)
				replicationDatabaseAuthIamModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ReplicationDatabaseAuth model
				replicationDatabaseAuthModel := new(cloudantv1.ReplicationDatabaseAuth)
				replicationDatabaseAuthModel.Iam = replicationDatabaseAuthIamModel

				// Construct an instance of the ReplicationDatabase model
				replicationDatabaseModel := new(cloudantv1.ReplicationDatabase)
				replicationDatabaseModel.Auth = replicationDatabaseAuthModel
				replicationDatabaseModel.HeadersVar = make(map[string]string)
				replicationDatabaseModel.URL = core.StringPtr("http://myserver.example:5984/foo-db")

				// Construct an instance of the UserContext model
				userContextModel := new(cloudantv1.UserContext)
				userContextModel.Db = core.StringPtr("testString")
				userContextModel.Name = core.StringPtr("john")
				userContextModel.Roles = []string{"_reader"}

				// Construct an instance of the ReplicationDocument model
				replicationDocumentModel := new(cloudantv1.ReplicationDocument)
				replicationDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				replicationDocumentModel.Conflicts = []string{"testString"}
				replicationDocumentModel.Deleted = core.BoolPtr(true)
				replicationDocumentModel.DeletedConflicts = []string{"testString"}
				replicationDocumentModel.ID = core.StringPtr("testString")
				replicationDocumentModel.LocalSeq = core.StringPtr("testString")
				replicationDocumentModel.Rev = core.StringPtr("testString")
				replicationDocumentModel.Revisions = revisionsModel
				replicationDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				replicationDocumentModel.Cancel = core.BoolPtr(false)
				replicationDocumentModel.CheckpointInterval = core.Int64Ptr(int64(4500))
				replicationDocumentModel.ConnectionTimeout = core.Int64Ptr(int64(15000))
				replicationDocumentModel.Continuous = core.BoolPtr(true)
				replicationDocumentModel.CreateTarget = core.BoolPtr(true)
				replicationDocumentModel.CreateTargetParams = replicationCreateTargetParametersModel
				replicationDocumentModel.DocIds = []string{"testString"}
				replicationDocumentModel.Filter = core.StringPtr("ddoc/my_filter")
				replicationDocumentModel.HTTPConnections = core.Int64Ptr(int64(10))
				replicationDocumentModel.QueryParams = make(map[string]string)
				replicationDocumentModel.RetriesPerRequest = core.Int64Ptr(int64(3))
				replicationDocumentModel.Selector = make(map[string]interface{})
				replicationDocumentModel.SinceSeq = core.StringPtr("34-g1AAAAGjeJzLYWBgYMlgTmGQT0lKzi9KdU")
				replicationDocumentModel.SocketOptions = core.StringPtr("[{keepalive, true}, {nodelay, false}]")
				replicationDocumentModel.Source = replicationDatabaseModel
				replicationDocumentModel.SourceProxy = core.StringPtr("http://my-source-proxy.example:8888")
				replicationDocumentModel.Target = replicationDatabaseModel
				replicationDocumentModel.TargetProxy = core.StringPtr("http://my-target-proxy.example:8888")
				replicationDocumentModel.UseCheckpoints = core.BoolPtr(false)
				replicationDocumentModel.UserCtx = userContextModel
				replicationDocumentModel.WorkerBatchSize = core.Int64Ptr(int64(400))
				replicationDocumentModel.WorkerProcesses = core.Int64Ptr(int64(3))
				replicationDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				replicationDocumentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PutReplicationDocumentOptions model
				putReplicationDocumentOptionsModel := new(cloudantv1.PutReplicationDocumentOptions)
				putReplicationDocumentOptionsModel.DocID = core.StringPtr("testString")
				putReplicationDocumentOptionsModel.ReplicationDocument = replicationDocumentModel
				putReplicationDocumentOptionsModel.IfMatch = core.StringPtr("testString")
				putReplicationDocumentOptionsModel.Batch = core.StringPtr("ok")
				putReplicationDocumentOptionsModel.NewEdits = core.BoolPtr(true)
				putReplicationDocumentOptionsModel.Rev = core.StringPtr("testString")
				putReplicationDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PutReplicationDocument(putReplicationDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutReplicationDocumentOptions model with no property values
				putReplicationDocumentOptionsModelNew := new(cloudantv1.PutReplicationDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PutReplicationDocument(putReplicationDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchedulerDocs(getSchedulerDocsOptions *GetSchedulerDocsOptions) - Operation response error`, func() {
		getSchedulerDocsPath := "/_scheduler/docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchedulerDocsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter


					// TODO: Add check for skip query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSchedulerDocs with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSchedulerDocsOptions model
				getSchedulerDocsOptionsModel := new(cloudantv1.GetSchedulerDocsOptions)
				getSchedulerDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				getSchedulerDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				getSchedulerDocsOptionsModel.States = []string{"initializing"}
				getSchedulerDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetSchedulerDocs(getSchedulerDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSchedulerDocs(getSchedulerDocsOptions *GetSchedulerDocsOptions)`, func() {
		getSchedulerDocsPath := "/_scheduler/docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchedulerDocsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter


					// TODO: Add check for skip query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "docs": [{"database": "Database", "doc_id": "DocID", "error_count": 0, "id": "ID", "info": {"changes_pending": 0, "checkpointed_source_seq": "CheckpointedSourceSeq", "doc_write_failures": 0, "docs_read": 0, "docs_written": 0, "error": "Error", "missing_revisions_found": 0, "revisions_checked": 0, "source_seq": "SourceSeq", "through_seq": "ThroughSeq"}, "last_updated": "2019-01-01T12:00:00", "node": "Node", "source": "Source", "source_proxy": "SourceProxy", "start_time": "2019-01-01T12:00:00", "state": "initializing", "target": "Target", "target_proxy": "TargetProxy"}]}`)
				}))
			})
			It(`Invoke GetSchedulerDocs successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetSchedulerDocs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchedulerDocsOptions model
				getSchedulerDocsOptionsModel := new(cloudantv1.GetSchedulerDocsOptions)
				getSchedulerDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				getSchedulerDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				getSchedulerDocsOptionsModel.States = []string{"initializing"}
				getSchedulerDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetSchedulerDocs(getSchedulerDocsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSchedulerDocs with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSchedulerDocsOptions model
				getSchedulerDocsOptionsModel := new(cloudantv1.GetSchedulerDocsOptions)
				getSchedulerDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				getSchedulerDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				getSchedulerDocsOptionsModel.States = []string{"initializing"}
				getSchedulerDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetSchedulerDocs(getSchedulerDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchedulerDocument(getSchedulerDocumentOptions *GetSchedulerDocumentOptions) - Operation response error`, func() {
		getSchedulerDocumentPath := "/_scheduler/docs/_replicator/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchedulerDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSchedulerDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSchedulerDocumentOptions model
				getSchedulerDocumentOptionsModel := new(cloudantv1.GetSchedulerDocumentOptions)
				getSchedulerDocumentOptionsModel.DocID = core.StringPtr("testString")
				getSchedulerDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetSchedulerDocument(getSchedulerDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSchedulerDocument(getSchedulerDocumentOptions *GetSchedulerDocumentOptions)`, func() {
		getSchedulerDocumentPath := "/_scheduler/docs/_replicator/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchedulerDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"database": "Database", "doc_id": "DocID", "error_count": 0, "id": "ID", "info": {"changes_pending": 0, "checkpointed_source_seq": "CheckpointedSourceSeq", "doc_write_failures": 0, "docs_read": 0, "docs_written": 0, "error": "Error", "missing_revisions_found": 0, "revisions_checked": 0, "source_seq": "SourceSeq", "through_seq": "ThroughSeq"}, "last_updated": "2019-01-01T12:00:00", "node": "Node", "source": "Source", "source_proxy": "SourceProxy", "start_time": "2019-01-01T12:00:00", "state": "initializing", "target": "Target", "target_proxy": "TargetProxy"}`)
				}))
			})
			It(`Invoke GetSchedulerDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetSchedulerDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchedulerDocumentOptions model
				getSchedulerDocumentOptionsModel := new(cloudantv1.GetSchedulerDocumentOptions)
				getSchedulerDocumentOptionsModel.DocID = core.StringPtr("testString")
				getSchedulerDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetSchedulerDocument(getSchedulerDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSchedulerDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSchedulerDocumentOptions model
				getSchedulerDocumentOptionsModel := new(cloudantv1.GetSchedulerDocumentOptions)
				getSchedulerDocumentOptionsModel.DocID = core.StringPtr("testString")
				getSchedulerDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetSchedulerDocument(getSchedulerDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSchedulerDocumentOptions model with no property values
				getSchedulerDocumentOptionsModelNew := new(cloudantv1.GetSchedulerDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetSchedulerDocument(getSchedulerDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchedulerJobs(getSchedulerJobsOptions *GetSchedulerJobsOptions) - Operation response error`, func() {
		getSchedulerJobsPath := "/_scheduler/jobs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchedulerJobsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter


					// TODO: Add check for skip query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSchedulerJobs with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSchedulerJobsOptions model
				getSchedulerJobsOptionsModel := new(cloudantv1.GetSchedulerJobsOptions)
				getSchedulerJobsOptionsModel.Limit = core.Int64Ptr(int64(0))
				getSchedulerJobsOptionsModel.Skip = core.Int64Ptr(int64(0))
				getSchedulerJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetSchedulerJobs(getSchedulerJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSchedulerJobs(getSchedulerJobsOptions *GetSchedulerJobsOptions)`, func() {
		getSchedulerJobsPath := "/_scheduler/jobs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchedulerJobsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter


					// TODO: Add check for skip query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "jobs": [{"database": "Database", "doc_id": "DocID", "history": [{"timestamp": "2019-01-01T12:00:00", "type": "Type"}], "id": "ID", "info": {"changes_pending": 0, "checkpointed_source_seq": "CheckpointedSourceSeq", "doc_write_failures": 0, "docs_read": 0, "docs_written": 0, "error": "Error", "missing_revisions_found": 0, "revisions_checked": 0, "source_seq": "SourceSeq", "through_seq": "ThroughSeq"}, "node": "Node", "pid": "Pid", "source": "Source", "start_time": "2019-01-01T12:00:00", "target": "Target", "user": "User"}]}`)
				}))
			})
			It(`Invoke GetSchedulerJobs successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetSchedulerJobs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchedulerJobsOptions model
				getSchedulerJobsOptionsModel := new(cloudantv1.GetSchedulerJobsOptions)
				getSchedulerJobsOptionsModel.Limit = core.Int64Ptr(int64(0))
				getSchedulerJobsOptionsModel.Skip = core.Int64Ptr(int64(0))
				getSchedulerJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetSchedulerJobs(getSchedulerJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSchedulerJobs with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSchedulerJobsOptions model
				getSchedulerJobsOptionsModel := new(cloudantv1.GetSchedulerJobsOptions)
				getSchedulerJobsOptionsModel.Limit = core.Int64Ptr(int64(0))
				getSchedulerJobsOptionsModel.Skip = core.Int64Ptr(int64(0))
				getSchedulerJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetSchedulerJobs(getSchedulerJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchedulerJob(getSchedulerJobOptions *GetSchedulerJobOptions) - Operation response error`, func() {
		getSchedulerJobPath := "/_scheduler/jobs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchedulerJobPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSchedulerJob with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSchedulerJobOptions model
				getSchedulerJobOptionsModel := new(cloudantv1.GetSchedulerJobOptions)
				getSchedulerJobOptionsModel.JobID = core.StringPtr("testString")
				getSchedulerJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetSchedulerJob(getSchedulerJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSchedulerJob(getSchedulerJobOptions *GetSchedulerJobOptions)`, func() {
		getSchedulerJobPath := "/_scheduler/jobs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchedulerJobPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"database": "Database", "doc_id": "DocID", "history": [{"timestamp": "2019-01-01T12:00:00", "type": "Type"}], "id": "ID", "info": {"changes_pending": 0, "checkpointed_source_seq": "CheckpointedSourceSeq", "doc_write_failures": 0, "docs_read": 0, "docs_written": 0, "error": "Error", "missing_revisions_found": 0, "revisions_checked": 0, "source_seq": "SourceSeq", "through_seq": "ThroughSeq"}, "node": "Node", "pid": "Pid", "source": "Source", "start_time": "2019-01-01T12:00:00", "target": "Target", "user": "User"}`)
				}))
			})
			It(`Invoke GetSchedulerJob successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetSchedulerJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchedulerJobOptions model
				getSchedulerJobOptionsModel := new(cloudantv1.GetSchedulerJobOptions)
				getSchedulerJobOptionsModel.JobID = core.StringPtr("testString")
				getSchedulerJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetSchedulerJob(getSchedulerJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSchedulerJob with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSchedulerJobOptions model
				getSchedulerJobOptionsModel := new(cloudantv1.GetSchedulerJobOptions)
				getSchedulerJobOptionsModel.JobID = core.StringPtr("testString")
				getSchedulerJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetSchedulerJob(getSchedulerJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSchedulerJobOptions model with no property values
				getSchedulerJobOptionsModelNew := new(cloudantv1.GetSchedulerJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetSchedulerJob(getSchedulerJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetSessionInformation(getSessionInformationOptions *GetSessionInformationOptions) - Operation response error`, func() {
		getSessionInformationPath := "/_session"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSessionInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSessionInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSessionInformationOptions model
				getSessionInformationOptionsModel := new(cloudantv1.GetSessionInformationOptions)
				getSessionInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetSessionInformation(getSessionInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSessionInformation(getSessionInformationOptions *GetSessionInformationOptions)`, func() {
		getSessionInformationPath := "/_session"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSessionInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ok": true, "info": {"authenticated": "Authenticated", "authentication_db": "AuthenticationDb", "authentication_handlers": ["AuthenticationHandlers"]}, "userCtx": {"db": "Db", "name": "Name", "roles": ["_reader"]}}`)
				}))
			})
			It(`Invoke GetSessionInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetSessionInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSessionInformationOptions model
				getSessionInformationOptionsModel := new(cloudantv1.GetSessionInformationOptions)
				getSessionInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetSessionInformation(getSessionInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSessionInformation with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSessionInformationOptions model
				getSessionInformationOptionsModel := new(cloudantv1.GetSessionInformationOptions)
				getSessionInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetSessionInformation(getSessionInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteIamSession(deleteIamSessionOptions *DeleteIamSessionOptions) - Operation response error`, func() {
		deleteIamSessionPath := "/_iam_session"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteIamSessionPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteIamSession with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteIamSessionOptions model
				deleteIamSessionOptionsModel := new(cloudantv1.DeleteIamSessionOptions)
				deleteIamSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.DeleteIamSession(deleteIamSessionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteIamSession(deleteIamSessionOptions *DeleteIamSessionOptions)`, func() {
		deleteIamSessionPath := "/_iam_session"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteIamSessionPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ok": true}`)
				}))
			})
			It(`Invoke DeleteIamSession successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.DeleteIamSession(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteIamSessionOptions model
				deleteIamSessionOptionsModel := new(cloudantv1.DeleteIamSessionOptions)
				deleteIamSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.DeleteIamSession(deleteIamSessionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteIamSession with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteIamSessionOptions model
				deleteIamSessionOptionsModel := new(cloudantv1.DeleteIamSessionOptions)
				deleteIamSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.DeleteIamSession(deleteIamSessionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetIamSessionInformation(getIamSessionInformationOptions *GetIamSessionInformationOptions) - Operation response error`, func() {
		getIamSessionInformationPath := "/_iam_session"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIamSessionInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetIamSessionInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetIamSessionInformationOptions model
				getIamSessionInformationOptionsModel := new(cloudantv1.GetIamSessionInformationOptions)
				getIamSessionInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetIamSessionInformation(getIamSessionInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetIamSessionInformation(getIamSessionInformationOptions *GetIamSessionInformationOptions)`, func() {
		getIamSessionInformationPath := "/_iam_session"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIamSessionInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "ok": true, "scope": "Scope", "type": "Type"}`)
				}))
			})
			It(`Invoke GetIamSessionInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetIamSessionInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetIamSessionInformationOptions model
				getIamSessionInformationOptionsModel := new(cloudantv1.GetIamSessionInformationOptions)
				getIamSessionInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetIamSessionInformation(getIamSessionInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetIamSessionInformation with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetIamSessionInformationOptions model
				getIamSessionInformationOptionsModel := new(cloudantv1.GetIamSessionInformationOptions)
				getIamSessionInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetIamSessionInformation(getIamSessionInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostIamSession(postIamSessionOptions *PostIamSessionOptions) - Operation response error`, func() {
		postIamSessionPath := "/_iam_session"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postIamSessionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostIamSession with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostIamSessionOptions model
				postIamSessionOptionsModel := new(cloudantv1.PostIamSessionOptions)
				postIamSessionOptionsModel.AccessToken = core.StringPtr("testString")
				postIamSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostIamSession(postIamSessionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostIamSession(postIamSessionOptions *PostIamSessionOptions)`, func() {
		postIamSessionPath := "/_iam_session"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postIamSessionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ok": true}`)
				}))
			})
			It(`Invoke PostIamSession successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostIamSession(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostIamSessionOptions model
				postIamSessionOptionsModel := new(cloudantv1.PostIamSessionOptions)
				postIamSessionOptionsModel.AccessToken = core.StringPtr("testString")
				postIamSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostIamSession(postIamSessionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostIamSession with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostIamSessionOptions model
				postIamSessionOptionsModel := new(cloudantv1.PostIamSessionOptions)
				postIamSessionOptionsModel.AccessToken = core.StringPtr("testString")
				postIamSessionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostIamSession(postIamSessionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetSecurity(getSecurityOptions *GetSecurityOptions) - Operation response error`, func() {
		getSecurityPath := "/testString/_security"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSecurityPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSecurity with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSecurityOptions model
				getSecurityOptionsModel := new(cloudantv1.GetSecurityOptions)
				getSecurityOptionsModel.Db = core.StringPtr("testString")
				getSecurityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetSecurity(getSecurityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSecurity(getSecurityOptions *GetSecurityOptions)`, func() {
		getSecurityPath := "/testString/_security"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSecurityPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"admins": {"names": ["Names"], "roles": ["Roles"]}, "members": {"names": ["Names"], "roles": ["Roles"]}, "cloudant": {"mapKey": ["_reader"]}, "couchdb_auth_only": false}`)
				}))
			})
			It(`Invoke GetSecurity successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetSecurity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSecurityOptions model
				getSecurityOptionsModel := new(cloudantv1.GetSecurityOptions)
				getSecurityOptionsModel.Db = core.StringPtr("testString")
				getSecurityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetSecurity(getSecurityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSecurity with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetSecurityOptions model
				getSecurityOptionsModel := new(cloudantv1.GetSecurityOptions)
				getSecurityOptionsModel.Db = core.StringPtr("testString")
				getSecurityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetSecurity(getSecurityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSecurityOptions model with no property values
				getSecurityOptionsModelNew := new(cloudantv1.GetSecurityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetSecurity(getSecurityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutSecurity(putSecurityOptions *PutSecurityOptions) - Operation response error`, func() {
		putSecurityPath := "/testString/_security"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putSecurityPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutSecurity with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the SecurityObject model
				securityObjectModel := new(cloudantv1.SecurityObject)
				securityObjectModel.Names = []string{"testString"}
				securityObjectModel.Roles = []string{"testString"}

				// Construct an instance of the PutSecurityOptions model
				putSecurityOptionsModel := new(cloudantv1.PutSecurityOptions)
				putSecurityOptionsModel.Db = core.StringPtr("testString")
				putSecurityOptionsModel.Admins = securityObjectModel
				putSecurityOptionsModel.Members = securityObjectModel
				putSecurityOptionsModel.Cloudant = make(map[string][]string)
				putSecurityOptionsModel.CouchdbAuthOnly = core.BoolPtr(true)
				putSecurityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PutSecurity(putSecurityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutSecurity(putSecurityOptions *PutSecurityOptions)`, func() {
		putSecurityPath := "/testString/_security"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putSecurityPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ok": true}`)
				}))
			})
			It(`Invoke PutSecurity successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PutSecurity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SecurityObject model
				securityObjectModel := new(cloudantv1.SecurityObject)
				securityObjectModel.Names = []string{"testString"}
				securityObjectModel.Roles = []string{"testString"}

				// Construct an instance of the PutSecurityOptions model
				putSecurityOptionsModel := new(cloudantv1.PutSecurityOptions)
				putSecurityOptionsModel.Db = core.StringPtr("testString")
				putSecurityOptionsModel.Admins = securityObjectModel
				putSecurityOptionsModel.Members = securityObjectModel
				putSecurityOptionsModel.Cloudant = make(map[string][]string)
				putSecurityOptionsModel.CouchdbAuthOnly = core.BoolPtr(true)
				putSecurityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PutSecurity(putSecurityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PutSecurity with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the SecurityObject model
				securityObjectModel := new(cloudantv1.SecurityObject)
				securityObjectModel.Names = []string{"testString"}
				securityObjectModel.Roles = []string{"testString"}

				// Construct an instance of the PutSecurityOptions model
				putSecurityOptionsModel := new(cloudantv1.PutSecurityOptions)
				putSecurityOptionsModel.Db = core.StringPtr("testString")
				putSecurityOptionsModel.Admins = securityObjectModel
				putSecurityOptionsModel.Members = securityObjectModel
				putSecurityOptionsModel.Cloudant = make(map[string][]string)
				putSecurityOptionsModel.CouchdbAuthOnly = core.BoolPtr(true)
				putSecurityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PutSecurity(putSecurityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutSecurityOptions model with no property values
				putSecurityOptionsModelNew := new(cloudantv1.PutSecurityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PutSecurity(putSecurityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostApiKeys(postApiKeysOptions *PostApiKeysOptions) - Operation response error`, func() {
		postApiKeysPath := "/_api/v2/api_keys"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postApiKeysPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostApiKeys with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostApiKeysOptions model
				postApiKeysOptionsModel := new(cloudantv1.PostApiKeysOptions)
				postApiKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostApiKeys(postApiKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostApiKeys(postApiKeysOptions *PostApiKeysOptions)`, func() {
		postApiKeysPath := "/_api/v2/api_keys"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postApiKeysPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"ok": true, "key": "Key", "password": "Password"}`)
				}))
			})
			It(`Invoke PostApiKeys successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostApiKeys(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostApiKeysOptions model
				postApiKeysOptionsModel := new(cloudantv1.PostApiKeysOptions)
				postApiKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostApiKeys(postApiKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostApiKeys with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostApiKeysOptions model
				postApiKeysOptionsModel := new(cloudantv1.PostApiKeysOptions)
				postApiKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostApiKeys(postApiKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutCloudantSecurityConfiguration(putCloudantSecurityConfigurationOptions *PutCloudantSecurityConfigurationOptions) - Operation response error`, func() {
		putCloudantSecurityConfigurationPath := "/_api/v2/db/testString/_security"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putCloudantSecurityConfigurationPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutCloudantSecurityConfiguration with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the SecurityObject model
				securityObjectModel := new(cloudantv1.SecurityObject)
				securityObjectModel.Names = []string{"testString"}
				securityObjectModel.Roles = []string{"testString"}

				// Construct an instance of the PutCloudantSecurityConfigurationOptions model
				putCloudantSecurityConfigurationOptionsModel := new(cloudantv1.PutCloudantSecurityConfigurationOptions)
				putCloudantSecurityConfigurationOptionsModel.Db = core.StringPtr("testString")
				putCloudantSecurityConfigurationOptionsModel.Admins = securityObjectModel
				putCloudantSecurityConfigurationOptionsModel.Members = securityObjectModel
				putCloudantSecurityConfigurationOptionsModel.Cloudant = make(map[string][]string)
				putCloudantSecurityConfigurationOptionsModel.CouchdbAuthOnly = core.BoolPtr(true)
				putCloudantSecurityConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PutCloudantSecurityConfiguration(putCloudantSecurityConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutCloudantSecurityConfiguration(putCloudantSecurityConfigurationOptions *PutCloudantSecurityConfigurationOptions)`, func() {
		putCloudantSecurityConfigurationPath := "/_api/v2/db/testString/_security"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putCloudantSecurityConfigurationPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ok": true}`)
				}))
			})
			It(`Invoke PutCloudantSecurityConfiguration successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PutCloudantSecurityConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SecurityObject model
				securityObjectModel := new(cloudantv1.SecurityObject)
				securityObjectModel.Names = []string{"testString"}
				securityObjectModel.Roles = []string{"testString"}

				// Construct an instance of the PutCloudantSecurityConfigurationOptions model
				putCloudantSecurityConfigurationOptionsModel := new(cloudantv1.PutCloudantSecurityConfigurationOptions)
				putCloudantSecurityConfigurationOptionsModel.Db = core.StringPtr("testString")
				putCloudantSecurityConfigurationOptionsModel.Admins = securityObjectModel
				putCloudantSecurityConfigurationOptionsModel.Members = securityObjectModel
				putCloudantSecurityConfigurationOptionsModel.Cloudant = make(map[string][]string)
				putCloudantSecurityConfigurationOptionsModel.CouchdbAuthOnly = core.BoolPtr(true)
				putCloudantSecurityConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PutCloudantSecurityConfiguration(putCloudantSecurityConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PutCloudantSecurityConfiguration with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the SecurityObject model
				securityObjectModel := new(cloudantv1.SecurityObject)
				securityObjectModel.Names = []string{"testString"}
				securityObjectModel.Roles = []string{"testString"}

				// Construct an instance of the PutCloudantSecurityConfigurationOptions model
				putCloudantSecurityConfigurationOptionsModel := new(cloudantv1.PutCloudantSecurityConfigurationOptions)
				putCloudantSecurityConfigurationOptionsModel.Db = core.StringPtr("testString")
				putCloudantSecurityConfigurationOptionsModel.Admins = securityObjectModel
				putCloudantSecurityConfigurationOptionsModel.Members = securityObjectModel
				putCloudantSecurityConfigurationOptionsModel.Cloudant = make(map[string][]string)
				putCloudantSecurityConfigurationOptionsModel.CouchdbAuthOnly = core.BoolPtr(true)
				putCloudantSecurityConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PutCloudantSecurityConfiguration(putCloudantSecurityConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutCloudantSecurityConfigurationOptions model with no property values
				putCloudantSecurityConfigurationOptionsModelNew := new(cloudantv1.PutCloudantSecurityConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PutCloudantSecurityConfiguration(putCloudantSecurityConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetCorsInformation(getCorsInformationOptions *GetCorsInformationOptions) - Operation response error`, func() {
		getCorsInformationPath := "/_api/v2/user/config/cors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCorsInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCorsInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetCorsInformationOptions model
				getCorsInformationOptionsModel := new(cloudantv1.GetCorsInformationOptions)
				getCorsInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetCorsInformation(getCorsInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCorsInformation(getCorsInformationOptions *GetCorsInformationOptions)`, func() {
		getCorsInformationPath := "/_api/v2/user/config/cors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCorsInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"allow_credentials": true, "enable_cors": true, "origins": ["Origins"]}`)
				}))
			})
			It(`Invoke GetCorsInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetCorsInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCorsInformationOptions model
				getCorsInformationOptionsModel := new(cloudantv1.GetCorsInformationOptions)
				getCorsInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetCorsInformation(getCorsInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetCorsInformation with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetCorsInformationOptions model
				getCorsInformationOptionsModel := new(cloudantv1.GetCorsInformationOptions)
				getCorsInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetCorsInformation(getCorsInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutCorsConfiguration(putCorsConfigurationOptions *PutCorsConfigurationOptions) - Operation response error`, func() {
		putCorsConfigurationPath := "/_api/v2/user/config/cors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putCorsConfigurationPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutCorsConfiguration with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PutCorsConfigurationOptions model
				putCorsConfigurationOptionsModel := new(cloudantv1.PutCorsConfigurationOptions)
				putCorsConfigurationOptionsModel.AllowCredentials = core.BoolPtr(true)
				putCorsConfigurationOptionsModel.EnableCors = core.BoolPtr(true)
				putCorsConfigurationOptionsModel.Origins = []string{"testString"}
				putCorsConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PutCorsConfiguration(putCorsConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutCorsConfiguration(putCorsConfigurationOptions *PutCorsConfigurationOptions)`, func() {
		putCorsConfigurationPath := "/_api/v2/user/config/cors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putCorsConfigurationPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ok": true}`)
				}))
			})
			It(`Invoke PutCorsConfiguration successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PutCorsConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PutCorsConfigurationOptions model
				putCorsConfigurationOptionsModel := new(cloudantv1.PutCorsConfigurationOptions)
				putCorsConfigurationOptionsModel.AllowCredentials = core.BoolPtr(true)
				putCorsConfigurationOptionsModel.EnableCors = core.BoolPtr(true)
				putCorsConfigurationOptionsModel.Origins = []string{"testString"}
				putCorsConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PutCorsConfiguration(putCorsConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PutCorsConfiguration with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PutCorsConfigurationOptions model
				putCorsConfigurationOptionsModel := new(cloudantv1.PutCorsConfigurationOptions)
				putCorsConfigurationOptionsModel.AllowCredentials = core.BoolPtr(true)
				putCorsConfigurationOptionsModel.EnableCors = core.BoolPtr(true)
				putCorsConfigurationOptionsModel.Origins = []string{"testString"}
				putCorsConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PutCorsConfiguration(putCorsConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`HeadAttachment(headAttachmentOptions *HeadAttachmentOptions)`, func() {
		headAttachmentPath := "/testString/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(headAttachmentPath))
					Expect(req.Method).To(Equal("HEAD"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke HeadAttachment successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := cloudantService.HeadAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HeadAttachmentOptions model
				headAttachmentOptionsModel := new(cloudantv1.HeadAttachmentOptions)
				headAttachmentOptionsModel.Db = core.StringPtr("testString")
				headAttachmentOptionsModel.DocID = core.StringPtr("testString")
				headAttachmentOptionsModel.AttachmentName = core.StringPtr("testString")
				headAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				headAttachmentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				headAttachmentOptionsModel.Rev = core.StringPtr("testString")
				headAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = cloudantService.HeadAttachment(headAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HeadAttachment with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the HeadAttachmentOptions model
				headAttachmentOptionsModel := new(cloudantv1.HeadAttachmentOptions)
				headAttachmentOptionsModel.Db = core.StringPtr("testString")
				headAttachmentOptionsModel.DocID = core.StringPtr("testString")
				headAttachmentOptionsModel.AttachmentName = core.StringPtr("testString")
				headAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				headAttachmentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				headAttachmentOptionsModel.Rev = core.StringPtr("testString")
				headAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := cloudantService.HeadAttachment(headAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HeadAttachmentOptions model with no property values
				headAttachmentOptionsModelNew := new(cloudantv1.HeadAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = cloudantService.HeadAttachment(headAttachmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAttachment(deleteAttachmentOptions *DeleteAttachmentOptions) - Operation response error`, func() {
		deleteAttachmentPath := "/testString/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAttachmentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAttachment with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteAttachmentOptions model
				deleteAttachmentOptionsModel := new(cloudantv1.DeleteAttachmentOptions)
				deleteAttachmentOptionsModel.Db = core.StringPtr("testString")
				deleteAttachmentOptionsModel.DocID = core.StringPtr("testString")
				deleteAttachmentOptionsModel.AttachmentName = core.StringPtr("testString")
				deleteAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteAttachmentOptionsModel.Rev = core.StringPtr("testString")
				deleteAttachmentOptionsModel.Batch = core.StringPtr("ok")
				deleteAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.DeleteAttachment(deleteAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAttachment(deleteAttachmentOptions *DeleteAttachmentOptions)`, func() {
		deleteAttachmentPath := "/testString/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAttachmentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke DeleteAttachment successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.DeleteAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAttachmentOptions model
				deleteAttachmentOptionsModel := new(cloudantv1.DeleteAttachmentOptions)
				deleteAttachmentOptionsModel.Db = core.StringPtr("testString")
				deleteAttachmentOptionsModel.DocID = core.StringPtr("testString")
				deleteAttachmentOptionsModel.AttachmentName = core.StringPtr("testString")
				deleteAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteAttachmentOptionsModel.Rev = core.StringPtr("testString")
				deleteAttachmentOptionsModel.Batch = core.StringPtr("ok")
				deleteAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.DeleteAttachment(deleteAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteAttachment with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteAttachmentOptions model
				deleteAttachmentOptionsModel := new(cloudantv1.DeleteAttachmentOptions)
				deleteAttachmentOptionsModel.Db = core.StringPtr("testString")
				deleteAttachmentOptionsModel.DocID = core.StringPtr("testString")
				deleteAttachmentOptionsModel.AttachmentName = core.StringPtr("testString")
				deleteAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				deleteAttachmentOptionsModel.Rev = core.StringPtr("testString")
				deleteAttachmentOptionsModel.Batch = core.StringPtr("ok")
				deleteAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.DeleteAttachment(deleteAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteAttachmentOptions model with no property values
				deleteAttachmentOptionsModelNew := new(cloudantv1.DeleteAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.DeleteAttachment(deleteAttachmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAttachment(getAttachmentOptions *GetAttachmentOptions)`, func() {
		getAttachmentPath := "/testString/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAttachmentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Range"]).ToNot(BeNil())
					Expect(req.Header["Range"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetAttachment successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAttachmentOptions model
				getAttachmentOptionsModel := new(cloudantv1.GetAttachmentOptions)
				getAttachmentOptionsModel.Db = core.StringPtr("testString")
				getAttachmentOptionsModel.DocID = core.StringPtr("testString")
				getAttachmentOptionsModel.AttachmentName = core.StringPtr("testString")
				getAttachmentOptionsModel.Accept = core.StringPtr("testString")
				getAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				getAttachmentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getAttachmentOptionsModel.Range = core.StringPtr("testString")
				getAttachmentOptionsModel.Rev = core.StringPtr("testString")
				getAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetAttachment(getAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetAttachment with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetAttachmentOptions model
				getAttachmentOptionsModel := new(cloudantv1.GetAttachmentOptions)
				getAttachmentOptionsModel.Db = core.StringPtr("testString")
				getAttachmentOptionsModel.DocID = core.StringPtr("testString")
				getAttachmentOptionsModel.AttachmentName = core.StringPtr("testString")
				getAttachmentOptionsModel.Accept = core.StringPtr("testString")
				getAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				getAttachmentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getAttachmentOptionsModel.Range = core.StringPtr("testString")
				getAttachmentOptionsModel.Rev = core.StringPtr("testString")
				getAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetAttachment(getAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAttachmentOptions model with no property values
				getAttachmentOptionsModelNew := new(cloudantv1.GetAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetAttachment(getAttachmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutAttachment(putAttachmentOptions *PutAttachmentOptions) - Operation response error`, func() {
		putAttachmentPath := "/testString/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putAttachmentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutAttachment with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PutAttachmentOptions model
				putAttachmentOptionsModel := new(cloudantv1.PutAttachmentOptions)
				putAttachmentOptionsModel.Db = core.StringPtr("testString")
				putAttachmentOptionsModel.DocID = core.StringPtr("testString")
				putAttachmentOptionsModel.AttachmentName = core.StringPtr("testString")
				putAttachmentOptionsModel.Attachment = CreateMockReader("This is a mock file.")
				putAttachmentOptionsModel.ContentType = core.StringPtr("testString")
				putAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				putAttachmentOptionsModel.Rev = core.StringPtr("testString")
				putAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PutAttachment(putAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutAttachment(putAttachmentOptions *PutAttachmentOptions)`, func() {
		putAttachmentPath := "/testString/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putAttachmentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["rev"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke PutAttachment successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PutAttachment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PutAttachmentOptions model
				putAttachmentOptionsModel := new(cloudantv1.PutAttachmentOptions)
				putAttachmentOptionsModel.Db = core.StringPtr("testString")
				putAttachmentOptionsModel.DocID = core.StringPtr("testString")
				putAttachmentOptionsModel.AttachmentName = core.StringPtr("testString")
				putAttachmentOptionsModel.Attachment = CreateMockReader("This is a mock file.")
				putAttachmentOptionsModel.ContentType = core.StringPtr("testString")
				putAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				putAttachmentOptionsModel.Rev = core.StringPtr("testString")
				putAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PutAttachment(putAttachmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PutAttachment with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PutAttachmentOptions model
				putAttachmentOptionsModel := new(cloudantv1.PutAttachmentOptions)
				putAttachmentOptionsModel.Db = core.StringPtr("testString")
				putAttachmentOptionsModel.DocID = core.StringPtr("testString")
				putAttachmentOptionsModel.AttachmentName = core.StringPtr("testString")
				putAttachmentOptionsModel.Attachment = CreateMockReader("This is a mock file.")
				putAttachmentOptionsModel.ContentType = core.StringPtr("testString")
				putAttachmentOptionsModel.IfMatch = core.StringPtr("testString")
				putAttachmentOptionsModel.Rev = core.StringPtr("testString")
				putAttachmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PutAttachment(putAttachmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutAttachmentOptions model with no property values
				putAttachmentOptionsModelNew := new(cloudantv1.PutAttachmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PutAttachment(putAttachmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`DeleteLocalDocument(deleteLocalDocumentOptions *DeleteLocalDocumentOptions) - Operation response error`, func() {
		deleteLocalDocumentPath := "/testString/_local/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLocalDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteLocalDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteLocalDocumentOptions model
				deleteLocalDocumentOptionsModel := new(cloudantv1.DeleteLocalDocumentOptions)
				deleteLocalDocumentOptionsModel.Db = core.StringPtr("testString")
				deleteLocalDocumentOptionsModel.DocID = core.StringPtr("testString")
				deleteLocalDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteLocalDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.DeleteLocalDocument(deleteLocalDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteLocalDocument(deleteLocalDocumentOptions *DeleteLocalDocumentOptions)`, func() {
		deleteLocalDocumentPath := "/testString/_local/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLocalDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke DeleteLocalDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.DeleteLocalDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteLocalDocumentOptions model
				deleteLocalDocumentOptionsModel := new(cloudantv1.DeleteLocalDocumentOptions)
				deleteLocalDocumentOptionsModel.Db = core.StringPtr("testString")
				deleteLocalDocumentOptionsModel.DocID = core.StringPtr("testString")
				deleteLocalDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteLocalDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.DeleteLocalDocument(deleteLocalDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteLocalDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the DeleteLocalDocumentOptions model
				deleteLocalDocumentOptionsModel := new(cloudantv1.DeleteLocalDocumentOptions)
				deleteLocalDocumentOptionsModel.Db = core.StringPtr("testString")
				deleteLocalDocumentOptionsModel.DocID = core.StringPtr("testString")
				deleteLocalDocumentOptionsModel.Batch = core.StringPtr("ok")
				deleteLocalDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.DeleteLocalDocument(deleteLocalDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteLocalDocumentOptions model with no property values
				deleteLocalDocumentOptionsModelNew := new(cloudantv1.DeleteLocalDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.DeleteLocalDocument(deleteLocalDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLocalDocument(getLocalDocumentOptions *GetLocalDocumentOptions) - Operation response error`, func() {
		getLocalDocumentPath := "/testString/_local/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLocalDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for local_seq query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLocalDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetLocalDocumentOptions model
				getLocalDocumentOptionsModel := new(cloudantv1.GetLocalDocumentOptions)
				getLocalDocumentOptionsModel.Db = core.StringPtr("testString")
				getLocalDocumentOptionsModel.DocID = core.StringPtr("testString")
				getLocalDocumentOptionsModel.Accept = core.StringPtr("application/json")
				getLocalDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getLocalDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getLocalDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getLocalDocumentOptionsModel.AttsSince = []string{"testString"}
				getLocalDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getLocalDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetLocalDocument(getLocalDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLocalDocument(getLocalDocumentOptions *GetLocalDocumentOptions)`, func() {
		getLocalDocumentPath := "/testString/_local/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLocalDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for attachments query parameter


					// TODO: Add check for att_encoding_info query parameter


					// TODO: Add check for local_seq query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}`)
				}))
			})
			It(`Invoke GetLocalDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetLocalDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLocalDocumentOptions model
				getLocalDocumentOptionsModel := new(cloudantv1.GetLocalDocumentOptions)
				getLocalDocumentOptionsModel.Db = core.StringPtr("testString")
				getLocalDocumentOptionsModel.DocID = core.StringPtr("testString")
				getLocalDocumentOptionsModel.Accept = core.StringPtr("application/json")
				getLocalDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getLocalDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getLocalDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getLocalDocumentOptionsModel.AttsSince = []string{"testString"}
				getLocalDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getLocalDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetLocalDocument(getLocalDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetLocalDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetLocalDocumentOptions model
				getLocalDocumentOptionsModel := new(cloudantv1.GetLocalDocumentOptions)
				getLocalDocumentOptionsModel.Db = core.StringPtr("testString")
				getLocalDocumentOptionsModel.DocID = core.StringPtr("testString")
				getLocalDocumentOptionsModel.Accept = core.StringPtr("application/json")
				getLocalDocumentOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getLocalDocumentOptionsModel.Attachments = core.BoolPtr(true)
				getLocalDocumentOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				getLocalDocumentOptionsModel.AttsSince = []string{"testString"}
				getLocalDocumentOptionsModel.LocalSeq = core.BoolPtr(true)
				getLocalDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetLocalDocument(getLocalDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLocalDocumentOptions model with no property values
				getLocalDocumentOptionsModelNew := new(cloudantv1.GetLocalDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetLocalDocument(getLocalDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutLocalDocument(putLocalDocumentOptions *PutLocalDocumentOptions) - Operation response error`, func() {
		putLocalDocumentPath := "/testString/_local/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putLocalDocumentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutLocalDocument with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PutLocalDocumentOptions model
				putLocalDocumentOptionsModel := new(cloudantv1.PutLocalDocumentOptions)
				putLocalDocumentOptionsModel.Db = core.StringPtr("testString")
				putLocalDocumentOptionsModel.DocID = core.StringPtr("testString")
				putLocalDocumentOptionsModel.Document = documentModel
				putLocalDocumentOptionsModel.ContentType = core.StringPtr("application/json")
				putLocalDocumentOptionsModel.Batch = core.StringPtr("ok")
				putLocalDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PutLocalDocument(putLocalDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutLocalDocument(putLocalDocumentOptions *PutLocalDocumentOptions)`, func() {
		putLocalDocumentPath := "/testString/_local/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putLocalDocumentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["batch"]).To(Equal([]string{"ok"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "rev": "Rev", "ok": true, "caused_by": "CausedBy", "error": "Error", "reason": "Reason"}`)
				}))
			})
			It(`Invoke PutLocalDocument successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PutLocalDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PutLocalDocumentOptions model
				putLocalDocumentOptionsModel := new(cloudantv1.PutLocalDocumentOptions)
				putLocalDocumentOptionsModel.Db = core.StringPtr("testString")
				putLocalDocumentOptionsModel.DocID = core.StringPtr("testString")
				putLocalDocumentOptionsModel.Document = documentModel
				putLocalDocumentOptionsModel.ContentType = core.StringPtr("application/json")
				putLocalDocumentOptionsModel.Batch = core.StringPtr("ok")
				putLocalDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PutLocalDocument(putLocalDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PutLocalDocument with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				documentModel.Attachments["foo"] = *attachmentModel

				// Construct an instance of the PutLocalDocumentOptions model
				putLocalDocumentOptionsModel := new(cloudantv1.PutLocalDocumentOptions)
				putLocalDocumentOptionsModel.Db = core.StringPtr("testString")
				putLocalDocumentOptionsModel.DocID = core.StringPtr("testString")
				putLocalDocumentOptionsModel.Document = documentModel
				putLocalDocumentOptionsModel.ContentType = core.StringPtr("application/json")
				putLocalDocumentOptionsModel.Batch = core.StringPtr("ok")
				putLocalDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PutLocalDocument(putLocalDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutLocalDocumentOptions model with no property values
				putLocalDocumentOptionsModelNew := new(cloudantv1.PutLocalDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PutLocalDocument(putLocalDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostLocalDocs(postLocalDocsOptions *PostLocalDocsOptions) - Operation response error`, func() {
		postLocalDocsPath := "/testString/_local_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postLocalDocsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostLocalDocs with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostLocalDocsOptions model
				postLocalDocsOptionsModel := new(cloudantv1.PostLocalDocsOptions)
				postLocalDocsOptionsModel.Db = core.StringPtr("testString")
				postLocalDocsOptionsModel.Accept = core.StringPtr("application/json")
				postLocalDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postLocalDocsOptionsModel.Attachments = core.BoolPtr(true)
				postLocalDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postLocalDocsOptionsModel.Descending = core.BoolPtr(true)
				postLocalDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postLocalDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postLocalDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postLocalDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postLocalDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postLocalDocsOptionsModel.Endkey = core.StringPtr("testString")
				postLocalDocsOptionsModel.Key = core.StringPtr("testString")
				postLocalDocsOptionsModel.Keys = []string{"testString"}
				postLocalDocsOptionsModel.Startkey = core.StringPtr("testString")
				postLocalDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostLocalDocs(postLocalDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostLocalDocs(postLocalDocsOptions *PostLocalDocsOptions)`, func() {
		postLocalDocsPath := "/testString/_local_docs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postLocalDocsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_rows": 0, "rows": [{"caused_by": "CausedBy", "error": "Error", "reason": "Reason", "doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "id": "ID", "key": "Key", "value": {"rev": "Rev"}}], "update_seq": "UpdateSeq"}`)
				}))
			})
			It(`Invoke PostLocalDocs successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostLocalDocs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostLocalDocsOptions model
				postLocalDocsOptionsModel := new(cloudantv1.PostLocalDocsOptions)
				postLocalDocsOptionsModel.Db = core.StringPtr("testString")
				postLocalDocsOptionsModel.Accept = core.StringPtr("application/json")
				postLocalDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postLocalDocsOptionsModel.Attachments = core.BoolPtr(true)
				postLocalDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postLocalDocsOptionsModel.Descending = core.BoolPtr(true)
				postLocalDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postLocalDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postLocalDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postLocalDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postLocalDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postLocalDocsOptionsModel.Endkey = core.StringPtr("testString")
				postLocalDocsOptionsModel.Key = core.StringPtr("testString")
				postLocalDocsOptionsModel.Keys = []string{"testString"}
				postLocalDocsOptionsModel.Startkey = core.StringPtr("testString")
				postLocalDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostLocalDocs(postLocalDocsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostLocalDocs with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostLocalDocsOptions model
				postLocalDocsOptionsModel := new(cloudantv1.PostLocalDocsOptions)
				postLocalDocsOptionsModel.Db = core.StringPtr("testString")
				postLocalDocsOptionsModel.Accept = core.StringPtr("application/json")
				postLocalDocsOptionsModel.AttEncodingInfo = core.BoolPtr(true)
				postLocalDocsOptionsModel.Attachments = core.BoolPtr(true)
				postLocalDocsOptionsModel.Conflicts = core.BoolPtr(true)
				postLocalDocsOptionsModel.Descending = core.BoolPtr(true)
				postLocalDocsOptionsModel.IncludeDocs = core.BoolPtr(true)
				postLocalDocsOptionsModel.InclusiveEnd = core.BoolPtr(true)
				postLocalDocsOptionsModel.Limit = core.Int64Ptr(int64(0))
				postLocalDocsOptionsModel.Skip = core.Int64Ptr(int64(0))
				postLocalDocsOptionsModel.UpdateSeq = core.BoolPtr(true)
				postLocalDocsOptionsModel.Endkey = core.StringPtr("testString")
				postLocalDocsOptionsModel.Key = core.StringPtr("testString")
				postLocalDocsOptionsModel.Keys = []string{"testString"}
				postLocalDocsOptionsModel.Startkey = core.StringPtr("testString")
				postLocalDocsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostLocalDocs(postLocalDocsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostLocalDocsOptions model with no property values
				postLocalDocsOptionsModelNew := new(cloudantv1.PostLocalDocsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostLocalDocs(postLocalDocsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostLocalDocsQueries(postLocalDocsQueriesOptions *PostLocalDocsQueriesOptions) - Operation response error`, func() {
		postLocalDocsQueriesPath := "/testString/_local_docs/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postLocalDocsQueriesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostLocalDocsQueries with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostLocalDocsQueriesOptions model
				postLocalDocsQueriesOptionsModel := new(cloudantv1.PostLocalDocsQueriesOptions)
				postLocalDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postLocalDocsQueriesOptionsModel.Accept = core.StringPtr("application/json")
				postLocalDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postLocalDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostLocalDocsQueries(postLocalDocsQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostLocalDocsQueries(postLocalDocsQueriesOptions *PostLocalDocsQueriesOptions)`, func() {
		postLocalDocsQueriesPath := "/testString/_local_docs/queries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postLocalDocsQueriesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"results": [{"total_rows": 0, "rows": [{"caused_by": "CausedBy", "error": "Error", "reason": "Reason", "doc": {"_attachments": {"mapKey": {"content_type": "ContentType", "data": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "digest": "Digest", "encoded_length": 0, "encoding": "Encoding", "follows": false, "length": 0, "revpos": 1, "stub": true}}, "_conflicts": ["Conflicts"], "_deleted": false, "_deleted_conflicts": ["DeletedConflicts"], "_id": "ID", "_local_seq": "LocalSeq", "_rev": "Rev", "_revisions": {"ids": ["Ids"], "start": 1}, "_revs_info": [{"rev": "Rev", "status": "available"}]}, "id": "ID", "key": "Key", "value": {"rev": "Rev"}}], "update_seq": "UpdateSeq"}]}`)
				}))
			})
			It(`Invoke PostLocalDocsQueries successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostLocalDocsQueries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostLocalDocsQueriesOptions model
				postLocalDocsQueriesOptionsModel := new(cloudantv1.PostLocalDocsQueriesOptions)
				postLocalDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postLocalDocsQueriesOptionsModel.Accept = core.StringPtr("application/json")
				postLocalDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postLocalDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostLocalDocsQueries(postLocalDocsQueriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostLocalDocsQueries with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")

				// Construct an instance of the PostLocalDocsQueriesOptions model
				postLocalDocsQueriesOptionsModel := new(cloudantv1.PostLocalDocsQueriesOptions)
				postLocalDocsQueriesOptionsModel.Db = core.StringPtr("testString")
				postLocalDocsQueriesOptionsModel.Accept = core.StringPtr("application/json")
				postLocalDocsQueriesOptionsModel.Queries = []cloudantv1.AllDocsQuery{*allDocsQueryModel}
				postLocalDocsQueriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostLocalDocsQueries(postLocalDocsQueriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostLocalDocsQueriesOptions model with no property values
				postLocalDocsQueriesOptionsModelNew := new(cloudantv1.PostLocalDocsQueriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostLocalDocsQueries(postLocalDocsQueriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`PostEnsureFullCommit(postEnsureFullCommitOptions *PostEnsureFullCommitOptions) - Operation response error`, func() {
		postEnsureFullCommitPath := "/testString/_ensure_full_commit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postEnsureFullCommitPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostEnsureFullCommit with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostEnsureFullCommitOptions model
				postEnsureFullCommitOptionsModel := new(cloudantv1.PostEnsureFullCommitOptions)
				postEnsureFullCommitOptionsModel.Db = core.StringPtr("testString")
				postEnsureFullCommitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostEnsureFullCommit(postEnsureFullCommitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostEnsureFullCommit(postEnsureFullCommitOptions *PostEnsureFullCommitOptions)`, func() {
		postEnsureFullCommitPath := "/testString/_ensure_full_commit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postEnsureFullCommitPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"instance_start_time": "InstanceStartTime", "ok": true}`)
				}))
			})
			It(`Invoke PostEnsureFullCommit successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostEnsureFullCommit(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostEnsureFullCommitOptions model
				postEnsureFullCommitOptionsModel := new(cloudantv1.PostEnsureFullCommitOptions)
				postEnsureFullCommitOptionsModel.Db = core.StringPtr("testString")
				postEnsureFullCommitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostEnsureFullCommit(postEnsureFullCommitOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostEnsureFullCommit with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostEnsureFullCommitOptions model
				postEnsureFullCommitOptionsModel := new(cloudantv1.PostEnsureFullCommitOptions)
				postEnsureFullCommitOptionsModel.Db = core.StringPtr("testString")
				postEnsureFullCommitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostEnsureFullCommit(postEnsureFullCommitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostEnsureFullCommitOptions model with no property values
				postEnsureFullCommitOptionsModelNew := new(cloudantv1.PostEnsureFullCommitOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostEnsureFullCommit(postEnsureFullCommitOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostMissingRevs(postMissingRevsOptions *PostMissingRevsOptions) - Operation response error`, func() {
		postMissingRevsPath := "/testString/_missing_revs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postMissingRevsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostMissingRevs with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostMissingRevsOptions model
				postMissingRevsOptionsModel := new(cloudantv1.PostMissingRevsOptions)
				postMissingRevsOptionsModel.Db = core.StringPtr("testString")
				postMissingRevsOptionsModel.DocumentRevisions = make(map[string][]string)
				postMissingRevsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostMissingRevs(postMissingRevsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostMissingRevs(postMissingRevsOptions *PostMissingRevsOptions)`, func() {
		postMissingRevsPath := "/testString/_missing_revs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postMissingRevsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"missing_revs": {"mapKey": ["Inner"]}}`)
				}))
			})
			It(`Invoke PostMissingRevs successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostMissingRevs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostMissingRevsOptions model
				postMissingRevsOptionsModel := new(cloudantv1.PostMissingRevsOptions)
				postMissingRevsOptionsModel.Db = core.StringPtr("testString")
				postMissingRevsOptionsModel.DocumentRevisions = make(map[string][]string)
				postMissingRevsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostMissingRevs(postMissingRevsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostMissingRevs with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostMissingRevsOptions model
				postMissingRevsOptionsModel := new(cloudantv1.PostMissingRevsOptions)
				postMissingRevsOptionsModel.Db = core.StringPtr("testString")
				postMissingRevsOptionsModel.DocumentRevisions = make(map[string][]string)
				postMissingRevsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostMissingRevs(postMissingRevsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostMissingRevsOptions model with no property values
				postMissingRevsOptionsModelNew := new(cloudantv1.PostMissingRevsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostMissingRevs(postMissingRevsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostRevsDiff(postRevsDiffOptions *PostRevsDiffOptions) - Operation response error`, func() {
		postRevsDiffPath := "/testString/_revs_diff"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postRevsDiffPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostRevsDiff with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostRevsDiffOptions model
				postRevsDiffOptionsModel := new(cloudantv1.PostRevsDiffOptions)
				postRevsDiffOptionsModel.Db = core.StringPtr("testString")
				postRevsDiffOptionsModel.DocumentRevisions = make(map[string][]string)
				postRevsDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.PostRevsDiff(postRevsDiffOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PostRevsDiff(postRevsDiffOptions *PostRevsDiffOptions)`, func() {
		postRevsDiffPath := "/testString/_revs_diff"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postRevsDiffPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": {"missing": ["Missing"], "possible_ancestors": ["PossibleAncestors"]}}`)
				}))
			})
			It(`Invoke PostRevsDiff successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.PostRevsDiff(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostRevsDiffOptions model
				postRevsDiffOptionsModel := new(cloudantv1.PostRevsDiffOptions)
				postRevsDiffOptionsModel.Db = core.StringPtr("testString")
				postRevsDiffOptionsModel.DocumentRevisions = make(map[string][]string)
				postRevsDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.PostRevsDiff(postRevsDiffOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PostRevsDiff with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the PostRevsDiffOptions model
				postRevsDiffOptionsModel := new(cloudantv1.PostRevsDiffOptions)
				postRevsDiffOptionsModel.Db = core.StringPtr("testString")
				postRevsDiffOptionsModel.DocumentRevisions = make(map[string][]string)
				postRevsDiffOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.PostRevsDiff(postRevsDiffOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostRevsDiffOptions model with no property values
				postRevsDiffOptionsModelNew := new(cloudantv1.PostRevsDiffOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.PostRevsDiff(postRevsDiffOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetShardsInformation(getShardsInformationOptions *GetShardsInformationOptions) - Operation response error`, func() {
		getShardsInformationPath := "/testString/_shards"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getShardsInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetShardsInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetShardsInformationOptions model
				getShardsInformationOptionsModel := new(cloudantv1.GetShardsInformationOptions)
				getShardsInformationOptionsModel.Db = core.StringPtr("testString")
				getShardsInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetShardsInformation(getShardsInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetShardsInformation(getShardsInformationOptions *GetShardsInformationOptions)`, func() {
		getShardsInformationPath := "/testString/_shards"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getShardsInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"shards": {"mapKey": ["Inner"]}}`)
				}))
			})
			It(`Invoke GetShardsInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetShardsInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetShardsInformationOptions model
				getShardsInformationOptionsModel := new(cloudantv1.GetShardsInformationOptions)
				getShardsInformationOptionsModel.Db = core.StringPtr("testString")
				getShardsInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetShardsInformation(getShardsInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetShardsInformation with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetShardsInformationOptions model
				getShardsInformationOptionsModel := new(cloudantv1.GetShardsInformationOptions)
				getShardsInformationOptionsModel.Db = core.StringPtr("testString")
				getShardsInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetShardsInformation(getShardsInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetShardsInformationOptions model with no property values
				getShardsInformationOptionsModelNew := new(cloudantv1.GetShardsInformationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetShardsInformation(getShardsInformationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDocumentShardsInfo(getDocumentShardsInfoOptions *GetDocumentShardsInfoOptions) - Operation response error`, func() {
		getDocumentShardsInfoPath := "/testString/_shards/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentShardsInfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDocumentShardsInfo with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDocumentShardsInfoOptions model
				getDocumentShardsInfoOptionsModel := new(cloudantv1.GetDocumentShardsInfoOptions)
				getDocumentShardsInfoOptionsModel.Db = core.StringPtr("testString")
				getDocumentShardsInfoOptionsModel.DocID = core.StringPtr("testString")
				getDocumentShardsInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetDocumentShardsInfo(getDocumentShardsInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocumentShardsInfo(getDocumentShardsInfoOptions *GetDocumentShardsInfoOptions)`, func() {
		getDocumentShardsInfoPath := "/testString/_shards/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDocumentShardsInfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"nodes": ["Nodes"], "range": "Range"}`)
				}))
			})
			It(`Invoke GetDocumentShardsInfo successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetDocumentShardsInfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDocumentShardsInfoOptions model
				getDocumentShardsInfoOptionsModel := new(cloudantv1.GetDocumentShardsInfoOptions)
				getDocumentShardsInfoOptionsModel.Db = core.StringPtr("testString")
				getDocumentShardsInfoOptionsModel.DocID = core.StringPtr("testString")
				getDocumentShardsInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetDocumentShardsInfo(getDocumentShardsInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDocumentShardsInfo with error: Operation validation and request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetDocumentShardsInfoOptions model
				getDocumentShardsInfoOptionsModel := new(cloudantv1.GetDocumentShardsInfoOptions)
				getDocumentShardsInfoOptionsModel.Db = core.StringPtr("testString")
				getDocumentShardsInfoOptionsModel.DocID = core.StringPtr("testString")
				getDocumentShardsInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetDocumentShardsInfo(getDocumentShardsInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentShardsInfoOptions model with no property values
				getDocumentShardsInfoOptionsModelNew := new(cloudantv1.GetDocumentShardsInfoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudantService.GetDocumentShardsInfo(getDocumentShardsInfoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudantService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL: "https://cloudantv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudantService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
					URL: "https://testService/api",
				})
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				})
				err := cloudantService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_URL": "https://cloudantv1/api",
				"CLOUDANT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUDANT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudantService, serviceErr := cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudantService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetActiveTasks(getActiveTasksOptions *GetActiveTasksOptions) - Operation response error`, func() {
		getActiveTasksPath := "/_active_tasks"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getActiveTasksPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetActiveTasks with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetActiveTasksOptions model
				getActiveTasksOptionsModel := new(cloudantv1.GetActiveTasksOptions)
				getActiveTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetActiveTasks(getActiveTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetActiveTasks(getActiveTasksOptions *GetActiveTasksOptions)`, func() {
		getActiveTasksPath := "/_active_tasks"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getActiveTasksPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"changes_done": 0, "database": "Database", "pid": "Pid", "progress": 0, "started_on": 0, "status": "Status", "task": "Task", "total_changes": 0, "type": "Type", "updated_on": 0}]`)
				}))
			})
			It(`Invoke GetActiveTasks successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetActiveTasks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetActiveTasksOptions model
				getActiveTasksOptionsModel := new(cloudantv1.GetActiveTasksOptions)
				getActiveTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetActiveTasks(getActiveTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetActiveTasks with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetActiveTasksOptions model
				getActiveTasksOptionsModel := new(cloudantv1.GetActiveTasksOptions)
				getActiveTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetActiveTasks(getActiveTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUpInformation(getUpInformationOptions *GetUpInformationOptions) - Operation response error`, func() {
		getUpInformationPath := "/_up"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUpInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUpInformation with error: Operation response processing error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetUpInformationOptions model
				getUpInformationOptionsModel := new(cloudantv1.GetUpInformationOptions)
				getUpInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudantService.GetUpInformation(getUpInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUpInformation(getUpInformationOptions *GetUpInformationOptions)`, func() {
		getUpInformationPath := "/_up"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUpInformationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "maintenance_mode"}`)
				}))
			})
			It(`Invoke GetUpInformation successfully`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudantService.GetUpInformation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUpInformationOptions model
				getUpInformationOptionsModel := new(cloudantv1.GetUpInformationOptions)
				getUpInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudantService.GetUpInformation(getUpInformationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetUpInformation with error: Operation request error`, func() {
				cloudantService, serviceErr := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudantService).ToNot(BeNil())

				// Construct an instance of the GetUpInformationOptions model
				getUpInformationOptionsModel := new(cloudantv1.GetUpInformationOptions)
				getUpInformationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudantService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudantService.GetUpInformation(getUpInformationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			cloudantService, _ := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
				URL:           "http://cloudantv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewBulkGetQueryDocument successfully`, func() {
				id := "testString"
				model, err := cloudantService.NewBulkGetQueryDocument(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCorsConfiguration successfully`, func() {
				origins := []string{"testString"}
				model, err := cloudantService.NewCorsConfiguration(origins)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDeleteAttachmentOptions successfully`, func() {
				// Construct an instance of the DeleteAttachmentOptions model
				db := "testString"
				docID := "testString"
				attachmentName := "testString"
				deleteAttachmentOptionsModel := cloudantService.NewDeleteAttachmentOptions(db, docID, attachmentName)
				deleteAttachmentOptionsModel.SetDb("testString")
				deleteAttachmentOptionsModel.SetDocID("testString")
				deleteAttachmentOptionsModel.SetAttachmentName("testString")
				deleteAttachmentOptionsModel.SetIfMatch("testString")
				deleteAttachmentOptionsModel.SetRev("testString")
				deleteAttachmentOptionsModel.SetBatch("ok")
				deleteAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAttachmentOptionsModel).ToNot(BeNil())
				Expect(deleteAttachmentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(deleteAttachmentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAttachmentOptionsModel.AttachmentName).To(Equal(core.StringPtr("testString")))
				Expect(deleteAttachmentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deleteAttachmentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(deleteAttachmentOptionsModel.Batch).To(Equal(core.StringPtr("ok")))
				Expect(deleteAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDatabaseOptions successfully`, func() {
				// Construct an instance of the DeleteDatabaseOptions model
				db := "testString"
				deleteDatabaseOptionsModel := cloudantService.NewDeleteDatabaseOptions(db)
				deleteDatabaseOptionsModel.SetDb("testString")
				deleteDatabaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDatabaseOptionsModel).ToNot(BeNil())
				Expect(deleteDatabaseOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(deleteDatabaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDesignDocumentOptions successfully`, func() {
				// Construct an instance of the DeleteDesignDocumentOptions model
				db := "testString"
				ddoc := "testString"
				deleteDesignDocumentOptionsModel := cloudantService.NewDeleteDesignDocumentOptions(db, ddoc)
				deleteDesignDocumentOptionsModel.SetDb("testString")
				deleteDesignDocumentOptionsModel.SetDdoc("testString")
				deleteDesignDocumentOptionsModel.SetIfMatch("testString")
				deleteDesignDocumentOptionsModel.SetBatch("ok")
				deleteDesignDocumentOptionsModel.SetRev("testString")
				deleteDesignDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDesignDocumentOptionsModel).ToNot(BeNil())
				Expect(deleteDesignDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(deleteDesignDocumentOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(deleteDesignDocumentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deleteDesignDocumentOptionsModel.Batch).To(Equal(core.StringPtr("ok")))
				Expect(deleteDesignDocumentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(deleteDesignDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDocumentOptions successfully`, func() {
				// Construct an instance of the DeleteDocumentOptions model
				db := "testString"
				docID := "testString"
				deleteDocumentOptionsModel := cloudantService.NewDeleteDocumentOptions(db, docID)
				deleteDocumentOptionsModel.SetDb("testString")
				deleteDocumentOptionsModel.SetDocID("testString")
				deleteDocumentOptionsModel.SetIfMatch("testString")
				deleteDocumentOptionsModel.SetBatch("ok")
				deleteDocumentOptionsModel.SetRev("testString")
				deleteDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDocumentOptionsModel).ToNot(BeNil())
				Expect(deleteDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.Batch).To(Equal(core.StringPtr("ok")))
				Expect(deleteDocumentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(deleteDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteIamSessionOptions successfully`, func() {
				// Construct an instance of the DeleteIamSessionOptions model
				deleteIamSessionOptionsModel := cloudantService.NewDeleteIamSessionOptions()
				deleteIamSessionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteIamSessionOptionsModel).ToNot(BeNil())
				Expect(deleteIamSessionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteIndexOptions successfully`, func() {
				// Construct an instance of the DeleteIndexOptions model
				db := "testString"
				ddoc := "testString"
				typeVar := "json"
				index := "testString"
				deleteIndexOptionsModel := cloudantService.NewDeleteIndexOptions(db, ddoc, typeVar, index)
				deleteIndexOptionsModel.SetDb("testString")
				deleteIndexOptionsModel.SetDdoc("testString")
				deleteIndexOptionsModel.SetType("json")
				deleteIndexOptionsModel.SetIndex("testString")
				deleteIndexOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteIndexOptionsModel).ToNot(BeNil())
				Expect(deleteIndexOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(deleteIndexOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(deleteIndexOptionsModel.Type).To(Equal(core.StringPtr("json")))
				Expect(deleteIndexOptionsModel.Index).To(Equal(core.StringPtr("testString")))
				Expect(deleteIndexOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLocalDocumentOptions successfully`, func() {
				// Construct an instance of the DeleteLocalDocumentOptions model
				db := "testString"
				docID := "testString"
				deleteLocalDocumentOptionsModel := cloudantService.NewDeleteLocalDocumentOptions(db, docID)
				deleteLocalDocumentOptionsModel.SetDb("testString")
				deleteLocalDocumentOptionsModel.SetDocID("testString")
				deleteLocalDocumentOptionsModel.SetBatch("ok")
				deleteLocalDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLocalDocumentOptionsModel).ToNot(BeNil())
				Expect(deleteLocalDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(deleteLocalDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLocalDocumentOptionsModel.Batch).To(Equal(core.StringPtr("ok")))
				Expect(deleteLocalDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteReplicationDocumentOptions successfully`, func() {
				// Construct an instance of the DeleteReplicationDocumentOptions model
				docID := "testString"
				deleteReplicationDocumentOptionsModel := cloudantService.NewDeleteReplicationDocumentOptions(docID)
				deleteReplicationDocumentOptionsModel.SetDocID("testString")
				deleteReplicationDocumentOptionsModel.SetIfMatch("testString")
				deleteReplicationDocumentOptionsModel.SetBatch("ok")
				deleteReplicationDocumentOptionsModel.SetRev("testString")
				deleteReplicationDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteReplicationDocumentOptionsModel).ToNot(BeNil())
				Expect(deleteReplicationDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(deleteReplicationDocumentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deleteReplicationDocumentOptionsModel.Batch).To(Equal(core.StringPtr("ok")))
				Expect(deleteReplicationDocumentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(deleteReplicationDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGeoIndexDefinition successfully`, func() {
				index := "testString"
				model, err := cloudantService.NewGeoIndexDefinition(index)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetActiveTasksOptions successfully`, func() {
				// Construct an instance of the GetActiveTasksOptions model
				getActiveTasksOptionsModel := cloudantService.NewGetActiveTasksOptions()
				getActiveTasksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getActiveTasksOptionsModel).ToNot(BeNil())
				Expect(getActiveTasksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAllDbsOptions successfully`, func() {
				// Construct an instance of the GetAllDbsOptions model
				getAllDbsOptionsModel := cloudantService.NewGetAllDbsOptions()
				getAllDbsOptionsModel.SetDescending(true)
				getAllDbsOptionsModel.SetEndkey("testString")
				getAllDbsOptionsModel.SetLimit(int64(0))
				getAllDbsOptionsModel.SetSkip(int64(0))
				getAllDbsOptionsModel.SetStartkey("testString")
				getAllDbsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAllDbsOptionsModel).ToNot(BeNil())
				Expect(getAllDbsOptionsModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(getAllDbsOptionsModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(getAllDbsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getAllDbsOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getAllDbsOptionsModel.Startkey).To(Equal(core.StringPtr("testString")))
				Expect(getAllDbsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAttachmentOptions successfully`, func() {
				// Construct an instance of the GetAttachmentOptions model
				db := "testString"
				docID := "testString"
				attachmentName := "testString"
				getAttachmentOptionsModel := cloudantService.NewGetAttachmentOptions(db, docID, attachmentName)
				getAttachmentOptionsModel.SetDb("testString")
				getAttachmentOptionsModel.SetDocID("testString")
				getAttachmentOptionsModel.SetAttachmentName("testString")
				getAttachmentOptionsModel.SetAccept("testString")
				getAttachmentOptionsModel.SetIfMatch("testString")
				getAttachmentOptionsModel.SetIfNoneMatch("testString")
				getAttachmentOptionsModel.SetRange("testString")
				getAttachmentOptionsModel.SetRev("testString")
				getAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAttachmentOptionsModel).ToNot(BeNil())
				Expect(getAttachmentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.AttachmentName).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.Range).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(getAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCorsInformationOptions successfully`, func() {
				// Construct an instance of the GetCorsInformationOptions model
				getCorsInformationOptionsModel := cloudantService.NewGetCorsInformationOptions()
				getCorsInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCorsInformationOptionsModel).ToNot(BeNil())
				Expect(getCorsInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDatabaseInformationOptions successfully`, func() {
				// Construct an instance of the GetDatabaseInformationOptions model
				db := "testString"
				getDatabaseInformationOptionsModel := cloudantService.NewGetDatabaseInformationOptions(db)
				getDatabaseInformationOptionsModel.SetDb("testString")
				getDatabaseInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDatabaseInformationOptionsModel).ToNot(BeNil())
				Expect(getDatabaseInformationOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getDatabaseInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDbUpdatesOptions successfully`, func() {
				// Construct an instance of the GetDbUpdatesOptions model
				getDbUpdatesOptionsModel := cloudantService.NewGetDbUpdatesOptions()
				getDbUpdatesOptionsModel.SetFeed("continuous")
				getDbUpdatesOptionsModel.SetHeartbeat(int64(0))
				getDbUpdatesOptionsModel.SetTimeout(int64(0))
				getDbUpdatesOptionsModel.SetSince("testString")
				getDbUpdatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDbUpdatesOptionsModel).ToNot(BeNil())
				Expect(getDbUpdatesOptionsModel.Feed).To(Equal(core.StringPtr("continuous")))
				Expect(getDbUpdatesOptionsModel.Heartbeat).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getDbUpdatesOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getDbUpdatesOptionsModel.Since).To(Equal(core.StringPtr("testString")))
				Expect(getDbUpdatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDesignDocumentInformationOptions successfully`, func() {
				// Construct an instance of the GetDesignDocumentInformationOptions model
				db := "testString"
				ddoc := "testString"
				getDesignDocumentInformationOptionsModel := cloudantService.NewGetDesignDocumentInformationOptions(db, ddoc)
				getDesignDocumentInformationOptionsModel.SetDb("testString")
				getDesignDocumentInformationOptionsModel.SetDdoc("testString")
				getDesignDocumentInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDesignDocumentInformationOptionsModel).ToNot(BeNil())
				Expect(getDesignDocumentInformationOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getDesignDocumentInformationOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(getDesignDocumentInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDesignDocumentOptions successfully`, func() {
				// Construct an instance of the GetDesignDocumentOptions model
				db := "testString"
				ddoc := "testString"
				getDesignDocumentOptionsModel := cloudantService.NewGetDesignDocumentOptions(db, ddoc)
				getDesignDocumentOptionsModel.SetDb("testString")
				getDesignDocumentOptionsModel.SetDdoc("testString")
				getDesignDocumentOptionsModel.SetIfNoneMatch("testString")
				getDesignDocumentOptionsModel.SetAttachments(true)
				getDesignDocumentOptionsModel.SetAttEncodingInfo(true)
				getDesignDocumentOptionsModel.SetAttsSince([]string{"testString"})
				getDesignDocumentOptionsModel.SetConflicts(true)
				getDesignDocumentOptionsModel.SetDeletedConflicts(true)
				getDesignDocumentOptionsModel.SetLatest(true)
				getDesignDocumentOptionsModel.SetLocalSeq(true)
				getDesignDocumentOptionsModel.SetMeta(true)
				getDesignDocumentOptionsModel.SetOpenRevs([]string{"testString"})
				getDesignDocumentOptionsModel.SetRev("testString")
				getDesignDocumentOptionsModel.SetRevs(true)
				getDesignDocumentOptionsModel.SetRevsInfo(true)
				getDesignDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDesignDocumentOptionsModel).ToNot(BeNil())
				Expect(getDesignDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getDesignDocumentOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(getDesignDocumentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(getDesignDocumentOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(getDesignDocumentOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(getDesignDocumentOptionsModel.AttsSince).To(Equal([]string{"testString"}))
				Expect(getDesignDocumentOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(getDesignDocumentOptionsModel.DeletedConflicts).To(Equal(core.BoolPtr(true)))
				Expect(getDesignDocumentOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(getDesignDocumentOptionsModel.LocalSeq).To(Equal(core.BoolPtr(true)))
				Expect(getDesignDocumentOptionsModel.Meta).To(Equal(core.BoolPtr(true)))
				Expect(getDesignDocumentOptionsModel.OpenRevs).To(Equal([]string{"testString"}))
				Expect(getDesignDocumentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(getDesignDocumentOptionsModel.Revs).To(Equal(core.BoolPtr(true)))
				Expect(getDesignDocumentOptionsModel.RevsInfo).To(Equal(core.BoolPtr(true)))
				Expect(getDesignDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDocumentOptions successfully`, func() {
				// Construct an instance of the GetDocumentOptions model
				db := "testString"
				docID := "testString"
				getDocumentOptionsModel := cloudantService.NewGetDocumentOptions(db, docID)
				getDocumentOptionsModel.SetDb("testString")
				getDocumentOptionsModel.SetDocID("testString")
				getDocumentOptionsModel.SetIfNoneMatch("testString")
				getDocumentOptionsModel.SetAttachments(true)
				getDocumentOptionsModel.SetAttEncodingInfo(true)
				getDocumentOptionsModel.SetAttsSince([]string{"testString"})
				getDocumentOptionsModel.SetConflicts(true)
				getDocumentOptionsModel.SetDeletedConflicts(true)
				getDocumentOptionsModel.SetLatest(true)
				getDocumentOptionsModel.SetLocalSeq(true)
				getDocumentOptionsModel.SetMeta(true)
				getDocumentOptionsModel.SetOpenRevs([]string{"testString"})
				getDocumentOptionsModel.SetRev("testString")
				getDocumentOptionsModel.SetRevs(true)
				getDocumentOptionsModel.SetRevsInfo(true)
				getDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDocumentOptionsModel).ToNot(BeNil())
				Expect(getDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentOptionsModel.AttsSince).To(Equal([]string{"testString"}))
				Expect(getDocumentOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentOptionsModel.DeletedConflicts).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentOptionsModel.LocalSeq).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentOptionsModel.Meta).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentOptionsModel.OpenRevs).To(Equal([]string{"testString"}))
				Expect(getDocumentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentOptionsModel.Revs).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentOptionsModel.RevsInfo).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDocumentShardsInfoOptions successfully`, func() {
				// Construct an instance of the GetDocumentShardsInfoOptions model
				db := "testString"
				docID := "testString"
				getDocumentShardsInfoOptionsModel := cloudantService.NewGetDocumentShardsInfoOptions(db, docID)
				getDocumentShardsInfoOptionsModel.SetDb("testString")
				getDocumentShardsInfoOptionsModel.SetDocID("testString")
				getDocumentShardsInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDocumentShardsInfoOptionsModel).ToNot(BeNil())
				Expect(getDocumentShardsInfoOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentShardsInfoOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentShardsInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGeoIndexInformationOptions successfully`, func() {
				// Construct an instance of the GetGeoIndexInformationOptions model
				db := "testString"
				ddoc := "testString"
				index := "testString"
				getGeoIndexInformationOptionsModel := cloudantService.NewGetGeoIndexInformationOptions(db, ddoc, index)
				getGeoIndexInformationOptionsModel.SetDb("testString")
				getGeoIndexInformationOptionsModel.SetDdoc("testString")
				getGeoIndexInformationOptionsModel.SetIndex("testString")
				getGeoIndexInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGeoIndexInformationOptionsModel).ToNot(BeNil())
				Expect(getGeoIndexInformationOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getGeoIndexInformationOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(getGeoIndexInformationOptionsModel.Index).To(Equal(core.StringPtr("testString")))
				Expect(getGeoIndexInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGeoOptions successfully`, func() {
				// Construct an instance of the GetGeoOptions model
				db := "testString"
				ddoc := "testString"
				index := "testString"
				getGeoOptionsModel := cloudantService.NewGetGeoOptions(db, ddoc, index)
				getGeoOptionsModel.SetDb("testString")
				getGeoOptionsModel.SetDdoc("testString")
				getGeoOptionsModel.SetIndex("testString")
				getGeoOptionsModel.SetBbox("testString")
				getGeoOptionsModel.SetBookmark("testString")
				getGeoOptionsModel.SetFormat("legacy")
				getGeoOptionsModel.SetG("testString")
				getGeoOptionsModel.SetIncludeDocs(true)
				getGeoOptionsModel.SetLat(float64(-90))
				getGeoOptionsModel.SetLimit(int64(0))
				getGeoOptionsModel.SetLon(float64(-180))
				getGeoOptionsModel.SetNearest(true)
				getGeoOptionsModel.SetRadius(float64(0))
				getGeoOptionsModel.SetRangex(float64(0))
				getGeoOptionsModel.SetRangey(float64(0))
				getGeoOptionsModel.SetRelation("contains")
				getGeoOptionsModel.SetSkip(int64(0))
				getGeoOptionsModel.SetStale("ok")
				getGeoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGeoOptionsModel).ToNot(BeNil())
				Expect(getGeoOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getGeoOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(getGeoOptionsModel.Index).To(Equal(core.StringPtr("testString")))
				Expect(getGeoOptionsModel.Bbox).To(Equal(core.StringPtr("testString")))
				Expect(getGeoOptionsModel.Bookmark).To(Equal(core.StringPtr("testString")))
				Expect(getGeoOptionsModel.Format).To(Equal(core.StringPtr("legacy")))
				Expect(getGeoOptionsModel.G).To(Equal(core.StringPtr("testString")))
				Expect(getGeoOptionsModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(getGeoOptionsModel.Lat).To(Equal(core.Float64Ptr(float64(-90))))
				Expect(getGeoOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getGeoOptionsModel.Lon).To(Equal(core.Float64Ptr(float64(-180))))
				Expect(getGeoOptionsModel.Nearest).To(Equal(core.BoolPtr(true)))
				Expect(getGeoOptionsModel.Radius).To(Equal(core.Float64Ptr(float64(0))))
				Expect(getGeoOptionsModel.Rangex).To(Equal(core.Float64Ptr(float64(0))))
				Expect(getGeoOptionsModel.Rangey).To(Equal(core.Float64Ptr(float64(0))))
				Expect(getGeoOptionsModel.Relation).To(Equal(core.StringPtr("contains")))
				Expect(getGeoOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getGeoOptionsModel.Stale).To(Equal(core.StringPtr("ok")))
				Expect(getGeoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetIamSessionInformationOptions successfully`, func() {
				// Construct an instance of the GetIamSessionInformationOptions model
				getIamSessionInformationOptionsModel := cloudantService.NewGetIamSessionInformationOptions()
				getIamSessionInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getIamSessionInformationOptionsModel).ToNot(BeNil())
				Expect(getIamSessionInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetIndexesInformationOptions successfully`, func() {
				// Construct an instance of the GetIndexesInformationOptions model
				db := "testString"
				getIndexesInformationOptionsModel := cloudantService.NewGetIndexesInformationOptions(db)
				getIndexesInformationOptionsModel.SetDb("testString")
				getIndexesInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getIndexesInformationOptionsModel).ToNot(BeNil())
				Expect(getIndexesInformationOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getIndexesInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLocalDocumentOptions successfully`, func() {
				// Construct an instance of the GetLocalDocumentOptions model
				db := "testString"
				docID := "testString"
				getLocalDocumentOptionsModel := cloudantService.NewGetLocalDocumentOptions(db, docID)
				getLocalDocumentOptionsModel.SetDb("testString")
				getLocalDocumentOptionsModel.SetDocID("testString")
				getLocalDocumentOptionsModel.SetAccept("application/json")
				getLocalDocumentOptionsModel.SetIfNoneMatch("testString")
				getLocalDocumentOptionsModel.SetAttachments(true)
				getLocalDocumentOptionsModel.SetAttEncodingInfo(true)
				getLocalDocumentOptionsModel.SetAttsSince([]string{"testString"})
				getLocalDocumentOptionsModel.SetLocalSeq(true)
				getLocalDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLocalDocumentOptionsModel).ToNot(BeNil())
				Expect(getLocalDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getLocalDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(getLocalDocumentOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(getLocalDocumentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(getLocalDocumentOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(getLocalDocumentOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(getLocalDocumentOptionsModel.AttsSince).To(Equal([]string{"testString"}))
				Expect(getLocalDocumentOptionsModel.LocalSeq).To(Equal(core.BoolPtr(true)))
				Expect(getLocalDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMembershipInformationOptions successfully`, func() {
				// Construct an instance of the GetMembershipInformationOptions model
				getMembershipInformationOptionsModel := cloudantService.NewGetMembershipInformationOptions()
				getMembershipInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMembershipInformationOptionsModel).ToNot(BeNil())
				Expect(getMembershipInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPartitionInformationOptions successfully`, func() {
				// Construct an instance of the GetPartitionInformationOptions model
				db := "testString"
				partitionKey := "testString"
				getPartitionInformationOptionsModel := cloudantService.NewGetPartitionInformationOptions(db, partitionKey)
				getPartitionInformationOptionsModel.SetDb("testString")
				getPartitionInformationOptionsModel.SetPartitionKey("testString")
				getPartitionInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPartitionInformationOptionsModel).ToNot(BeNil())
				Expect(getPartitionInformationOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getPartitionInformationOptionsModel.PartitionKey).To(Equal(core.StringPtr("testString")))
				Expect(getPartitionInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReplicationDocumentOptions successfully`, func() {
				// Construct an instance of the GetReplicationDocumentOptions model
				docID := "testString"
				getReplicationDocumentOptionsModel := cloudantService.NewGetReplicationDocumentOptions(docID)
				getReplicationDocumentOptionsModel.SetDocID("testString")
				getReplicationDocumentOptionsModel.SetIfNoneMatch("testString")
				getReplicationDocumentOptionsModel.SetAttachments(true)
				getReplicationDocumentOptionsModel.SetAttEncodingInfo(true)
				getReplicationDocumentOptionsModel.SetAttsSince([]string{"testString"})
				getReplicationDocumentOptionsModel.SetConflicts(true)
				getReplicationDocumentOptionsModel.SetDeletedConflicts(true)
				getReplicationDocumentOptionsModel.SetLatest(true)
				getReplicationDocumentOptionsModel.SetLocalSeq(true)
				getReplicationDocumentOptionsModel.SetMeta(true)
				getReplicationDocumentOptionsModel.SetOpenRevs([]string{"testString"})
				getReplicationDocumentOptionsModel.SetRev("testString")
				getReplicationDocumentOptionsModel.SetRevs(true)
				getReplicationDocumentOptionsModel.SetRevsInfo(true)
				getReplicationDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReplicationDocumentOptionsModel).ToNot(BeNil())
				Expect(getReplicationDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(getReplicationDocumentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(getReplicationDocumentOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(getReplicationDocumentOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(getReplicationDocumentOptionsModel.AttsSince).To(Equal([]string{"testString"}))
				Expect(getReplicationDocumentOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(getReplicationDocumentOptionsModel.DeletedConflicts).To(Equal(core.BoolPtr(true)))
				Expect(getReplicationDocumentOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(getReplicationDocumentOptionsModel.LocalSeq).To(Equal(core.BoolPtr(true)))
				Expect(getReplicationDocumentOptionsModel.Meta).To(Equal(core.BoolPtr(true)))
				Expect(getReplicationDocumentOptionsModel.OpenRevs).To(Equal([]string{"testString"}))
				Expect(getReplicationDocumentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(getReplicationDocumentOptionsModel.Revs).To(Equal(core.BoolPtr(true)))
				Expect(getReplicationDocumentOptionsModel.RevsInfo).To(Equal(core.BoolPtr(true)))
				Expect(getReplicationDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSchedulerDocsOptions successfully`, func() {
				// Construct an instance of the GetSchedulerDocsOptions model
				getSchedulerDocsOptionsModel := cloudantService.NewGetSchedulerDocsOptions()
				getSchedulerDocsOptionsModel.SetLimit(int64(0))
				getSchedulerDocsOptionsModel.SetSkip(int64(0))
				getSchedulerDocsOptionsModel.SetStates([]string{"initializing"})
				getSchedulerDocsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSchedulerDocsOptionsModel).ToNot(BeNil())
				Expect(getSchedulerDocsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getSchedulerDocsOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getSchedulerDocsOptionsModel.States).To(Equal([]string{"initializing"}))
				Expect(getSchedulerDocsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSchedulerDocumentOptions successfully`, func() {
				// Construct an instance of the GetSchedulerDocumentOptions model
				docID := "testString"
				getSchedulerDocumentOptionsModel := cloudantService.NewGetSchedulerDocumentOptions(docID)
				getSchedulerDocumentOptionsModel.SetDocID("testString")
				getSchedulerDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSchedulerDocumentOptionsModel).ToNot(BeNil())
				Expect(getSchedulerDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(getSchedulerDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSchedulerJobOptions successfully`, func() {
				// Construct an instance of the GetSchedulerJobOptions model
				jobID := "testString"
				getSchedulerJobOptionsModel := cloudantService.NewGetSchedulerJobOptions(jobID)
				getSchedulerJobOptionsModel.SetJobID("testString")
				getSchedulerJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSchedulerJobOptionsModel).ToNot(BeNil())
				Expect(getSchedulerJobOptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(getSchedulerJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSchedulerJobsOptions successfully`, func() {
				// Construct an instance of the GetSchedulerJobsOptions model
				getSchedulerJobsOptionsModel := cloudantService.NewGetSchedulerJobsOptions()
				getSchedulerJobsOptionsModel.SetLimit(int64(0))
				getSchedulerJobsOptionsModel.SetSkip(int64(0))
				getSchedulerJobsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSchedulerJobsOptionsModel).ToNot(BeNil())
				Expect(getSchedulerJobsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getSchedulerJobsOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(getSchedulerJobsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSearchInfoOptions successfully`, func() {
				// Construct an instance of the GetSearchInfoOptions model
				db := "testString"
				ddoc := "testString"
				index := "testString"
				getSearchInfoOptionsModel := cloudantService.NewGetSearchInfoOptions(db, ddoc, index)
				getSearchInfoOptionsModel.SetDb("testString")
				getSearchInfoOptionsModel.SetDdoc("testString")
				getSearchInfoOptionsModel.SetIndex("testString")
				getSearchInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSearchInfoOptionsModel).ToNot(BeNil())
				Expect(getSearchInfoOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getSearchInfoOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(getSearchInfoOptionsModel.Index).To(Equal(core.StringPtr("testString")))
				Expect(getSearchInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSecurityOptions successfully`, func() {
				// Construct an instance of the GetSecurityOptions model
				db := "testString"
				getSecurityOptionsModel := cloudantService.NewGetSecurityOptions(db)
				getSecurityOptionsModel.SetDb("testString")
				getSecurityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSecurityOptionsModel).ToNot(BeNil())
				Expect(getSecurityOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getSecurityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServerInformationOptions successfully`, func() {
				// Construct an instance of the GetServerInformationOptions model
				getServerInformationOptionsModel := cloudantService.NewGetServerInformationOptions()
				getServerInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServerInformationOptionsModel).ToNot(BeNil())
				Expect(getServerInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSessionInformationOptions successfully`, func() {
				// Construct an instance of the GetSessionInformationOptions model
				getSessionInformationOptionsModel := cloudantService.NewGetSessionInformationOptions()
				getSessionInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSessionInformationOptionsModel).ToNot(BeNil())
				Expect(getSessionInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetShardsInformationOptions successfully`, func() {
				// Construct an instance of the GetShardsInformationOptions model
				db := "testString"
				getShardsInformationOptionsModel := cloudantService.NewGetShardsInformationOptions(db)
				getShardsInformationOptionsModel.SetDb("testString")
				getShardsInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getShardsInformationOptionsModel).ToNot(BeNil())
				Expect(getShardsInformationOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(getShardsInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUpInformationOptions successfully`, func() {
				// Construct an instance of the GetUpInformationOptions model
				getUpInformationOptionsModel := cloudantService.NewGetUpInformationOptions()
				getUpInformationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUpInformationOptionsModel).ToNot(BeNil())
				Expect(getUpInformationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUuidsOptions successfully`, func() {
				// Construct an instance of the GetUuidsOptions model
				getUuidsOptionsModel := cloudantService.NewGetUuidsOptions()
				getUuidsOptionsModel.SetCount(int64(1))
				getUuidsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUuidsOptionsModel).ToNot(BeNil())
				Expect(getUuidsOptionsModel.Count).To(Equal(core.Int64Ptr(int64(1))))
				Expect(getUuidsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHeadAttachmentOptions successfully`, func() {
				// Construct an instance of the HeadAttachmentOptions model
				db := "testString"
				docID := "testString"
				attachmentName := "testString"
				headAttachmentOptionsModel := cloudantService.NewHeadAttachmentOptions(db, docID, attachmentName)
				headAttachmentOptionsModel.SetDb("testString")
				headAttachmentOptionsModel.SetDocID("testString")
				headAttachmentOptionsModel.SetAttachmentName("testString")
				headAttachmentOptionsModel.SetIfMatch("testString")
				headAttachmentOptionsModel.SetIfNoneMatch("testString")
				headAttachmentOptionsModel.SetRev("testString")
				headAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(headAttachmentOptionsModel).ToNot(BeNil())
				Expect(headAttachmentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(headAttachmentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(headAttachmentOptionsModel.AttachmentName).To(Equal(core.StringPtr("testString")))
				Expect(headAttachmentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(headAttachmentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(headAttachmentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(headAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHeadDatabaseOptions successfully`, func() {
				// Construct an instance of the HeadDatabaseOptions model
				db := "testString"
				headDatabaseOptionsModel := cloudantService.NewHeadDatabaseOptions(db)
				headDatabaseOptionsModel.SetDb("testString")
				headDatabaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(headDatabaseOptionsModel).ToNot(BeNil())
				Expect(headDatabaseOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(headDatabaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHeadDesignDocumentOptions successfully`, func() {
				// Construct an instance of the HeadDesignDocumentOptions model
				db := "testString"
				ddoc := "testString"
				headDesignDocumentOptionsModel := cloudantService.NewHeadDesignDocumentOptions(db, ddoc)
				headDesignDocumentOptionsModel.SetDb("testString")
				headDesignDocumentOptionsModel.SetDdoc("testString")
				headDesignDocumentOptionsModel.SetIfNoneMatch("testString")
				headDesignDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(headDesignDocumentOptionsModel).ToNot(BeNil())
				Expect(headDesignDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(headDesignDocumentOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(headDesignDocumentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(headDesignDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHeadDocumentOptions successfully`, func() {
				// Construct an instance of the HeadDocumentOptions model
				db := "testString"
				docID := "testString"
				headDocumentOptionsModel := cloudantService.NewHeadDocumentOptions(db, docID)
				headDocumentOptionsModel.SetDb("testString")
				headDocumentOptionsModel.SetDocID("testString")
				headDocumentOptionsModel.SetIfNoneMatch("testString")
				headDocumentOptionsModel.SetLatest(true)
				headDocumentOptionsModel.SetRev("testString")
				headDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(headDocumentOptionsModel).ToNot(BeNil())
				Expect(headDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(headDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(headDocumentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(headDocumentOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(headDocumentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(headDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHeadReplicationDocumentOptions successfully`, func() {
				// Construct an instance of the HeadReplicationDocumentOptions model
				docID := "testString"
				headReplicationDocumentOptionsModel := cloudantService.NewHeadReplicationDocumentOptions(docID)
				headReplicationDocumentOptionsModel.SetDocID("testString")
				headReplicationDocumentOptionsModel.SetIfNoneMatch("testString")
				headReplicationDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(headReplicationDocumentOptionsModel).ToNot(BeNil())
				Expect(headReplicationDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(headReplicationDocumentOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(headReplicationDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHeadSchedulerJobOptions successfully`, func() {
				// Construct an instance of the HeadSchedulerJobOptions model
				jobID := "testString"
				headSchedulerJobOptionsModel := cloudantService.NewHeadSchedulerJobOptions(jobID)
				headSchedulerJobOptionsModel.SetJobID("testString")
				headSchedulerJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(headSchedulerJobOptionsModel).ToNot(BeNil())
				Expect(headSchedulerJobOptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(headSchedulerJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostAllDocsOptions successfully`, func() {
				// Construct an instance of the PostAllDocsOptions model
				db := "testString"
				postAllDocsOptionsModel := cloudantService.NewPostAllDocsOptions(db)
				postAllDocsOptionsModel.SetDb("testString")
				postAllDocsOptionsModel.SetAttEncodingInfo(true)
				postAllDocsOptionsModel.SetAttachments(true)
				postAllDocsOptionsModel.SetConflicts(true)
				postAllDocsOptionsModel.SetDescending(true)
				postAllDocsOptionsModel.SetIncludeDocs(true)
				postAllDocsOptionsModel.SetInclusiveEnd(true)
				postAllDocsOptionsModel.SetLimit(int64(0))
				postAllDocsOptionsModel.SetSkip(int64(0))
				postAllDocsOptionsModel.SetUpdateSeq(true)
				postAllDocsOptionsModel.SetEndkey("testString")
				postAllDocsOptionsModel.SetKey("testString")
				postAllDocsOptionsModel.SetKeys([]string{"testString"})
				postAllDocsOptionsModel.SetStartkey("testString")
				postAllDocsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postAllDocsOptionsModel).ToNot(BeNil())
				Expect(postAllDocsOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postAllDocsOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(postAllDocsOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(postAllDocsOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(postAllDocsOptionsModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(postAllDocsOptionsModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(postAllDocsOptionsModel.InclusiveEnd).To(Equal(core.BoolPtr(true)))
				Expect(postAllDocsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postAllDocsOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postAllDocsOptionsModel.UpdateSeq).To(Equal(core.BoolPtr(true)))
				Expect(postAllDocsOptionsModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(postAllDocsOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(postAllDocsOptionsModel.Keys).To(Equal([]string{"testString"}))
				Expect(postAllDocsOptionsModel.Startkey).To(Equal(core.StringPtr("testString")))
				Expect(postAllDocsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostAllDocsQueriesOptions successfully`, func() {
				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				Expect(allDocsQueryModel).ToNot(BeNil())
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")
				Expect(allDocsQueryModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.InclusiveEnd).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(allDocsQueryModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(allDocsQueryModel.UpdateSeq).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(allDocsQueryModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(allDocsQueryModel.Keys).To(Equal([]string{"testString"}))
				Expect(allDocsQueryModel.Startkey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PostAllDocsQueriesOptions model
				db := "testString"
				postAllDocsQueriesOptionsModel := cloudantService.NewPostAllDocsQueriesOptions(db)
				postAllDocsQueriesOptionsModel.SetDb("testString")
				postAllDocsQueriesOptionsModel.SetQueries([]cloudantv1.AllDocsQuery{*allDocsQueryModel})
				postAllDocsQueriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postAllDocsQueriesOptionsModel).ToNot(BeNil())
				Expect(postAllDocsQueriesOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postAllDocsQueriesOptionsModel.Queries).To(Equal([]cloudantv1.AllDocsQuery{*allDocsQueryModel}))
				Expect(postAllDocsQueriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostApiKeysOptions successfully`, func() {
				// Construct an instance of the PostApiKeysOptions model
				postApiKeysOptionsModel := cloudantService.NewPostApiKeysOptions()
				postApiKeysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postApiKeysOptionsModel).ToNot(BeNil())
				Expect(postApiKeysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostBulkDocsOptions successfully`, func() {
				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				Expect(attachmentModel).ToNot(BeNil())
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)
				Expect(attachmentModel.ContentType).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Data).To(Equal(CreateMockByteArray("This is a mock byte array value.")))
				Expect(attachmentModel.Digest).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.EncodedLength).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Encoding).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Follows).To(Equal(core.BoolPtr(true)))
				Expect(attachmentModel.Length).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Revpos).To(Equal(core.Int64Ptr(int64(1))))
				Expect(attachmentModel.Stub).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				Expect(revisionsModel).ToNot(BeNil())
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))
				Expect(revisionsModel.Ids).To(Equal([]string{"testString"}))
				Expect(revisionsModel.Start).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				Expect(documentRevisionStatusModel).ToNot(BeNil())
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")
				Expect(documentRevisionStatusModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentRevisionStatusModel.Status).To(Equal(core.StringPtr("available")))

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				Expect(documentModel).ToNot(BeNil())
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.Attachments["foo"] = *attachmentModel
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(documentModel.Conflicts).To(Equal([]string{"testString"}))
				Expect(documentModel.Deleted).To(Equal(core.BoolPtr(true)))
				Expect(documentModel.DeletedConflicts).To(Equal([]string{"testString"}))
				Expect(documentModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.LocalSeq).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Revisions).To(Equal(revisionsModel))
				Expect(documentModel.RevsInfo).To(Equal([]cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}))
				Expect(documentModel.GetProperties()).ToNot(BeEmpty())
				Expect(documentModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Attachments["foo"]).To(Equal(*attachmentModel))

				// Construct an instance of the BulkDocs model
				bulkDocsModel := new(cloudantv1.BulkDocs)
				Expect(bulkDocsModel).ToNot(BeNil())
				bulkDocsModel.Docs = []cloudantv1.Document{*documentModel}
				bulkDocsModel.NewEdits = core.BoolPtr(true)
				Expect(bulkDocsModel.Docs).To(Equal([]cloudantv1.Document{*documentModel}))
				Expect(bulkDocsModel.NewEdits).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the PostBulkDocsOptions model
				db := "testString"
				postBulkDocsOptionsModel := cloudantService.NewPostBulkDocsOptions(db)
				postBulkDocsOptionsModel.SetDb("testString")
				postBulkDocsOptionsModel.SetBulkDocs(bulkDocsModel)
				postBulkDocsOptionsModel.SetBody(CreateMockReader("This is a mock file."))
				postBulkDocsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postBulkDocsOptionsModel).ToNot(BeNil())
				Expect(postBulkDocsOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postBulkDocsOptionsModel.BulkDocs).To(Equal(bulkDocsModel))
				Expect(postBulkDocsOptionsModel.Body).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(postBulkDocsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostBulkGetOptions successfully`, func() {
				// Construct an instance of the BulkGetQueryDocument model
				bulkGetQueryDocumentModel := new(cloudantv1.BulkGetQueryDocument)
				Expect(bulkGetQueryDocumentModel).ToNot(BeNil())
				bulkGetQueryDocumentModel.AttsSince = []string{"testString"}
				bulkGetQueryDocumentModel.ID = core.StringPtr("foo")
				bulkGetQueryDocumentModel.OpenRevs = []string{"testString"}
				bulkGetQueryDocumentModel.Rev = core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")
				Expect(bulkGetQueryDocumentModel.AttsSince).To(Equal([]string{"testString"}))
				Expect(bulkGetQueryDocumentModel.ID).To(Equal(core.StringPtr("foo")))
				Expect(bulkGetQueryDocumentModel.OpenRevs).To(Equal([]string{"testString"}))
				Expect(bulkGetQueryDocumentModel.Rev).To(Equal(core.StringPtr("4-753875d51501a6b1883a9d62b4d33f91")))

				// Construct an instance of the PostBulkGetOptions model
				db := "testString"
				postBulkGetOptionsModel := cloudantService.NewPostBulkGetOptions(db)
				postBulkGetOptionsModel.SetDb("testString")
				postBulkGetOptionsModel.SetDocs([]cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel})
				postBulkGetOptionsModel.SetAttachments(true)
				postBulkGetOptionsModel.SetAttEncodingInfo(true)
				postBulkGetOptionsModel.SetLatest(true)
				postBulkGetOptionsModel.SetRevs(true)
				postBulkGetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postBulkGetOptionsModel).ToNot(BeNil())
				Expect(postBulkGetOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postBulkGetOptionsModel.Docs).To(Equal([]cloudantv1.BulkGetQueryDocument{*bulkGetQueryDocumentModel}))
				Expect(postBulkGetOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(postBulkGetOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(postBulkGetOptionsModel.Latest).To(Equal(core.BoolPtr(true)))
				Expect(postBulkGetOptionsModel.Revs).To(Equal(core.BoolPtr(true)))
				Expect(postBulkGetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostChangesOptions successfully`, func() {
				// Construct an instance of the PostChangesOptions model
				db := "testString"
				postChangesOptionsModel := cloudantService.NewPostChangesOptions(db)
				postChangesOptionsModel.SetDb("testString")
				postChangesOptionsModel.SetDocIds([]string{"testString"})
				postChangesOptionsModel.SetFields([]string{"testString"})
				postChangesOptionsModel.SetSelector(make(map[string]interface{}))
				postChangesOptionsModel.SetLastEventID("testString")
				postChangesOptionsModel.SetAttEncodingInfo(true)
				postChangesOptionsModel.SetAttachments(true)
				postChangesOptionsModel.SetConflicts(true)
				postChangesOptionsModel.SetDescending(true)
				postChangesOptionsModel.SetFeed("continuous")
				postChangesOptionsModel.SetFilter("testString")
				postChangesOptionsModel.SetHeartbeat(int64(0))
				postChangesOptionsModel.SetIncludeDocs(true)
				postChangesOptionsModel.SetLimit(int64(0))
				postChangesOptionsModel.SetSeqInterval(int64(1))
				postChangesOptionsModel.SetSince("testString")
				postChangesOptionsModel.SetStyle("testString")
				postChangesOptionsModel.SetTimeout(int64(0))
				postChangesOptionsModel.SetView("testString")
				postChangesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postChangesOptionsModel).ToNot(BeNil())
				Expect(postChangesOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postChangesOptionsModel.DocIds).To(Equal([]string{"testString"}))
				Expect(postChangesOptionsModel.Fields).To(Equal([]string{"testString"}))
				Expect(postChangesOptionsModel.Selector).To(Equal(make(map[string]interface{})))
				Expect(postChangesOptionsModel.LastEventID).To(Equal(core.StringPtr("testString")))
				Expect(postChangesOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(postChangesOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(postChangesOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(postChangesOptionsModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(postChangesOptionsModel.Feed).To(Equal(core.StringPtr("continuous")))
				Expect(postChangesOptionsModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(postChangesOptionsModel.Heartbeat).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postChangesOptionsModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(postChangesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postChangesOptionsModel.SeqInterval).To(Equal(core.Int64Ptr(int64(1))))
				Expect(postChangesOptionsModel.Since).To(Equal(core.StringPtr("testString")))
				Expect(postChangesOptionsModel.Style).To(Equal(core.StringPtr("testString")))
				Expect(postChangesOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postChangesOptionsModel.View).To(Equal(core.StringPtr("testString")))
				Expect(postChangesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostDbsInfoOptions successfully`, func() {
				// Construct an instance of the PostDbsInfoOptions model
				postDbsInfoOptionsModel := cloudantService.NewPostDbsInfoOptions()
				postDbsInfoOptionsModel.SetKeys([]string{"testString"})
				postDbsInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postDbsInfoOptionsModel).ToNot(BeNil())
				Expect(postDbsInfoOptionsModel.Keys).To(Equal([]string{"testString"}))
				Expect(postDbsInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostDesignDocsOptions successfully`, func() {
				// Construct an instance of the PostDesignDocsOptions model
				db := "testString"
				postDesignDocsOptionsModel := cloudantService.NewPostDesignDocsOptions(db)
				postDesignDocsOptionsModel.SetDb("testString")
				postDesignDocsOptionsModel.SetAccept("application/json")
				postDesignDocsOptionsModel.SetAttEncodingInfo(true)
				postDesignDocsOptionsModel.SetAttachments(true)
				postDesignDocsOptionsModel.SetConflicts(true)
				postDesignDocsOptionsModel.SetDescending(true)
				postDesignDocsOptionsModel.SetIncludeDocs(true)
				postDesignDocsOptionsModel.SetInclusiveEnd(true)
				postDesignDocsOptionsModel.SetLimit(int64(0))
				postDesignDocsOptionsModel.SetSkip(int64(0))
				postDesignDocsOptionsModel.SetUpdateSeq(true)
				postDesignDocsOptionsModel.SetEndkey("testString")
				postDesignDocsOptionsModel.SetKey("testString")
				postDesignDocsOptionsModel.SetKeys([]string{"testString"})
				postDesignDocsOptionsModel.SetStartkey("testString")
				postDesignDocsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postDesignDocsOptionsModel).ToNot(BeNil())
				Expect(postDesignDocsOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postDesignDocsOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(postDesignDocsOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(postDesignDocsOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(postDesignDocsOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(postDesignDocsOptionsModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(postDesignDocsOptionsModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(postDesignDocsOptionsModel.InclusiveEnd).To(Equal(core.BoolPtr(true)))
				Expect(postDesignDocsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postDesignDocsOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postDesignDocsOptionsModel.UpdateSeq).To(Equal(core.BoolPtr(true)))
				Expect(postDesignDocsOptionsModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(postDesignDocsOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(postDesignDocsOptionsModel.Keys).To(Equal([]string{"testString"}))
				Expect(postDesignDocsOptionsModel.Startkey).To(Equal(core.StringPtr("testString")))
				Expect(postDesignDocsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostDesignDocsQueriesOptions successfully`, func() {
				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				Expect(allDocsQueryModel).ToNot(BeNil())
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")
				Expect(allDocsQueryModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.InclusiveEnd).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(allDocsQueryModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(allDocsQueryModel.UpdateSeq).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(allDocsQueryModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(allDocsQueryModel.Keys).To(Equal([]string{"testString"}))
				Expect(allDocsQueryModel.Startkey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PostDesignDocsQueriesOptions model
				db := "testString"
				postDesignDocsQueriesOptionsModel := cloudantService.NewPostDesignDocsQueriesOptions(db)
				postDesignDocsQueriesOptionsModel.SetDb("testString")
				postDesignDocsQueriesOptionsModel.SetAccept("application/json")
				postDesignDocsQueriesOptionsModel.SetQueries([]cloudantv1.AllDocsQuery{*allDocsQueryModel})
				postDesignDocsQueriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postDesignDocsQueriesOptionsModel).ToNot(BeNil())
				Expect(postDesignDocsQueriesOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postDesignDocsQueriesOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(postDesignDocsQueriesOptionsModel.Queries).To(Equal([]cloudantv1.AllDocsQuery{*allDocsQueryModel}))
				Expect(postDesignDocsQueriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostDocumentOptions successfully`, func() {
				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				Expect(attachmentModel).ToNot(BeNil())
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)
				Expect(attachmentModel.ContentType).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Data).To(Equal(CreateMockByteArray("This is a mock byte array value.")))
				Expect(attachmentModel.Digest).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.EncodedLength).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Encoding).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Follows).To(Equal(core.BoolPtr(true)))
				Expect(attachmentModel.Length).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Revpos).To(Equal(core.Int64Ptr(int64(1))))
				Expect(attachmentModel.Stub).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				Expect(revisionsModel).ToNot(BeNil())
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))
				Expect(revisionsModel.Ids).To(Equal([]string{"testString"}))
				Expect(revisionsModel.Start).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				Expect(documentRevisionStatusModel).ToNot(BeNil())
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")
				Expect(documentRevisionStatusModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentRevisionStatusModel.Status).To(Equal(core.StringPtr("available")))

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				Expect(documentModel).ToNot(BeNil())
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.Attachments["foo"] = *attachmentModel
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(documentModel.Conflicts).To(Equal([]string{"testString"}))
				Expect(documentModel.Deleted).To(Equal(core.BoolPtr(true)))
				Expect(documentModel.DeletedConflicts).To(Equal([]string{"testString"}))
				Expect(documentModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.LocalSeq).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Revisions).To(Equal(revisionsModel))
				Expect(documentModel.RevsInfo).To(Equal([]cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}))
				Expect(documentModel.GetProperties()).ToNot(BeEmpty())
				Expect(documentModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Attachments["foo"]).To(Equal(*attachmentModel))

				// Construct an instance of the PostDocumentOptions model
				db := "testString"
				postDocumentOptionsModel := cloudantService.NewPostDocumentOptions(db)
				postDocumentOptionsModel.SetDb("testString")
				postDocumentOptionsModel.SetDocument(documentModel)
				postDocumentOptionsModel.SetBody(CreateMockReader("This is a mock file."))
				postDocumentOptionsModel.SetContentType("application/json")
				postDocumentOptionsModel.SetBatch("ok")
				postDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postDocumentOptionsModel).ToNot(BeNil())
				Expect(postDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postDocumentOptionsModel.Document).To(Equal(documentModel))
				Expect(postDocumentOptionsModel.Body).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(postDocumentOptionsModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(postDocumentOptionsModel.Batch).To(Equal(core.StringPtr("ok")))
				Expect(postDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostEnsureFullCommitOptions successfully`, func() {
				// Construct an instance of the PostEnsureFullCommitOptions model
				db := "testString"
				postEnsureFullCommitOptionsModel := cloudantService.NewPostEnsureFullCommitOptions(db)
				postEnsureFullCommitOptionsModel.SetDb("testString")
				postEnsureFullCommitOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postEnsureFullCommitOptionsModel).ToNot(BeNil())
				Expect(postEnsureFullCommitOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postEnsureFullCommitOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostExplainOptions successfully`, func() {
				// Construct an instance of the PostExplainOptions model
				db := "testString"
				postExplainOptionsModel := cloudantService.NewPostExplainOptions(db)
				postExplainOptionsModel.SetDb("testString")
				postExplainOptionsModel.SetBookmark("testString")
				postExplainOptionsModel.SetConflicts(true)
				postExplainOptionsModel.SetExecutionStats(true)
				postExplainOptionsModel.SetFields([]string{"testString"})
				postExplainOptionsModel.SetLimit(int64(0))
				postExplainOptionsModel.SetSelector(make(map[string]interface{}))
				postExplainOptionsModel.SetSkip(int64(0))
				postExplainOptionsModel.SetSort([]map[string]string{make(map[string]string)})
				postExplainOptionsModel.SetStable(true)
				postExplainOptionsModel.SetUpdate("false")
				postExplainOptionsModel.SetUseIndex([]string{"testString"})
				postExplainOptionsModel.SetR(int64(1))
				postExplainOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postExplainOptionsModel).ToNot(BeNil())
				Expect(postExplainOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postExplainOptionsModel.Bookmark).To(Equal(core.StringPtr("testString")))
				Expect(postExplainOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(postExplainOptionsModel.ExecutionStats).To(Equal(core.BoolPtr(true)))
				Expect(postExplainOptionsModel.Fields).To(Equal([]string{"testString"}))
				Expect(postExplainOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postExplainOptionsModel.Selector).To(Equal(make(map[string]interface{})))
				Expect(postExplainOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postExplainOptionsModel.Sort).To(Equal([]map[string]string{make(map[string]string)}))
				Expect(postExplainOptionsModel.Stable).To(Equal(core.BoolPtr(true)))
				Expect(postExplainOptionsModel.Update).To(Equal(core.StringPtr("false")))
				Expect(postExplainOptionsModel.UseIndex).To(Equal([]string{"testString"}))
				Expect(postExplainOptionsModel.R).To(Equal(core.Int64Ptr(int64(1))))
				Expect(postExplainOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostFindOptions successfully`, func() {
				// Construct an instance of the PostFindOptions model
				db := "testString"
				postFindOptionsModel := cloudantService.NewPostFindOptions(db)
				postFindOptionsModel.SetDb("testString")
				postFindOptionsModel.SetBookmark("testString")
				postFindOptionsModel.SetConflicts(true)
				postFindOptionsModel.SetExecutionStats(true)
				postFindOptionsModel.SetFields([]string{"testString"})
				postFindOptionsModel.SetLimit(int64(0))
				postFindOptionsModel.SetSelector(make(map[string]interface{}))
				postFindOptionsModel.SetSkip(int64(0))
				postFindOptionsModel.SetSort([]map[string]string{make(map[string]string)})
				postFindOptionsModel.SetStable(true)
				postFindOptionsModel.SetUpdate("false")
				postFindOptionsModel.SetUseIndex([]string{"testString"})
				postFindOptionsModel.SetR(int64(1))
				postFindOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postFindOptionsModel).ToNot(BeNil())
				Expect(postFindOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postFindOptionsModel.Bookmark).To(Equal(core.StringPtr("testString")))
				Expect(postFindOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(postFindOptionsModel.ExecutionStats).To(Equal(core.BoolPtr(true)))
				Expect(postFindOptionsModel.Fields).To(Equal([]string{"testString"}))
				Expect(postFindOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postFindOptionsModel.Selector).To(Equal(make(map[string]interface{})))
				Expect(postFindOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postFindOptionsModel.Sort).To(Equal([]map[string]string{make(map[string]string)}))
				Expect(postFindOptionsModel.Stable).To(Equal(core.BoolPtr(true)))
				Expect(postFindOptionsModel.Update).To(Equal(core.StringPtr("false")))
				Expect(postFindOptionsModel.UseIndex).To(Equal([]string{"testString"}))
				Expect(postFindOptionsModel.R).To(Equal(core.Int64Ptr(int64(1))))
				Expect(postFindOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostGeoCleanupOptions successfully`, func() {
				// Construct an instance of the PostGeoCleanupOptions model
				db := "testString"
				postGeoCleanupOptionsModel := cloudantService.NewPostGeoCleanupOptions(db)
				postGeoCleanupOptionsModel.SetDb("testString")
				postGeoCleanupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postGeoCleanupOptionsModel).ToNot(BeNil())
				Expect(postGeoCleanupOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postGeoCleanupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostIamSessionOptions successfully`, func() {
				// Construct an instance of the PostIamSessionOptions model
				postIamSessionOptionsModel := cloudantService.NewPostIamSessionOptions()
				postIamSessionOptionsModel.SetAccessToken("testString")
				postIamSessionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postIamSessionOptionsModel).ToNot(BeNil())
				Expect(postIamSessionOptionsModel.AccessToken).To(Equal(core.StringPtr("testString")))
				Expect(postIamSessionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostIndexOptions successfully`, func() {
				// Construct an instance of the Analyzer model
				analyzerModel := new(cloudantv1.Analyzer)
				Expect(analyzerModel).ToNot(BeNil())
				analyzerModel.Name = core.StringPtr("classic")
				analyzerModel.Stopwords = []string{"testString"}
				Expect(analyzerModel.Name).To(Equal(core.StringPtr("classic")))
				Expect(analyzerModel.Stopwords).To(Equal([]string{"testString"}))

				// Construct an instance of the IndexTextOperatorDefaultField model
				indexTextOperatorDefaultFieldModel := new(cloudantv1.IndexTextOperatorDefaultField)
				Expect(indexTextOperatorDefaultFieldModel).ToNot(BeNil())
				indexTextOperatorDefaultFieldModel.Analyzer = analyzerModel
				indexTextOperatorDefaultFieldModel.Enabled = core.BoolPtr(true)
				Expect(indexTextOperatorDefaultFieldModel.Analyzer).To(Equal(analyzerModel))
				Expect(indexTextOperatorDefaultFieldModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the IndexField model
				indexFieldModel := new(cloudantv1.IndexField)
				Expect(indexFieldModel).ToNot(BeNil())
				indexFieldModel.Name = core.StringPtr("testString")
				indexFieldModel.Type = core.StringPtr("boolean")
				indexFieldModel.SetProperty("foo", core.StringPtr("asc"))
				Expect(indexFieldModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(indexFieldModel.Type).To(Equal(core.StringPtr("boolean")))
				Expect(indexFieldModel.GetProperties()).ToNot(BeEmpty())
				Expect(indexFieldModel.GetProperty("foo")).To(Equal(core.StringPtr("asc")))

				// Construct an instance of the IndexDefinition model
				indexDefinitionModel := new(cloudantv1.IndexDefinition)
				Expect(indexDefinitionModel).ToNot(BeNil())
				indexDefinitionModel.DefaultAnalyzer = analyzerModel
				indexDefinitionModel.DefaultField = indexTextOperatorDefaultFieldModel
				indexDefinitionModel.Fields = []cloudantv1.IndexField{*indexFieldModel}
				indexDefinitionModel.IndexArrayLengths = core.BoolPtr(true)
				Expect(indexDefinitionModel.DefaultAnalyzer).To(Equal(analyzerModel))
				Expect(indexDefinitionModel.DefaultField).To(Equal(indexTextOperatorDefaultFieldModel))
				Expect(indexDefinitionModel.Fields).To(Equal([]cloudantv1.IndexField{*indexFieldModel}))
				Expect(indexDefinitionModel.IndexArrayLengths).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the PostIndexOptions model
				db := "testString"
				postIndexOptionsModel := cloudantService.NewPostIndexOptions(db)
				postIndexOptionsModel.SetDb("testString")
				postIndexOptionsModel.SetDdoc("testString")
				postIndexOptionsModel.SetDef(indexDefinitionModel)
				postIndexOptionsModel.SetIndex(indexDefinitionModel)
				postIndexOptionsModel.SetName("testString")
				postIndexOptionsModel.SetPartialFilterSelector(make(map[string]interface{}))
				postIndexOptionsModel.SetPartitioned(true)
				postIndexOptionsModel.SetType("json")
				postIndexOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postIndexOptionsModel).ToNot(BeNil())
				Expect(postIndexOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postIndexOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(postIndexOptionsModel.Def).To(Equal(indexDefinitionModel))
				Expect(postIndexOptionsModel.Index).To(Equal(indexDefinitionModel))
				Expect(postIndexOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(postIndexOptionsModel.PartialFilterSelector).To(Equal(make(map[string]interface{})))
				Expect(postIndexOptionsModel.Partitioned).To(Equal(core.BoolPtr(true)))
				Expect(postIndexOptionsModel.Type).To(Equal(core.StringPtr("json")))
				Expect(postIndexOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostLocalDocsOptions successfully`, func() {
				// Construct an instance of the PostLocalDocsOptions model
				db := "testString"
				postLocalDocsOptionsModel := cloudantService.NewPostLocalDocsOptions(db)
				postLocalDocsOptionsModel.SetDb("testString")
				postLocalDocsOptionsModel.SetAccept("application/json")
				postLocalDocsOptionsModel.SetAttEncodingInfo(true)
				postLocalDocsOptionsModel.SetAttachments(true)
				postLocalDocsOptionsModel.SetConflicts(true)
				postLocalDocsOptionsModel.SetDescending(true)
				postLocalDocsOptionsModel.SetIncludeDocs(true)
				postLocalDocsOptionsModel.SetInclusiveEnd(true)
				postLocalDocsOptionsModel.SetLimit(int64(0))
				postLocalDocsOptionsModel.SetSkip(int64(0))
				postLocalDocsOptionsModel.SetUpdateSeq(true)
				postLocalDocsOptionsModel.SetEndkey("testString")
				postLocalDocsOptionsModel.SetKey("testString")
				postLocalDocsOptionsModel.SetKeys([]string{"testString"})
				postLocalDocsOptionsModel.SetStartkey("testString")
				postLocalDocsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postLocalDocsOptionsModel).ToNot(BeNil())
				Expect(postLocalDocsOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postLocalDocsOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(postLocalDocsOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(postLocalDocsOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(postLocalDocsOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(postLocalDocsOptionsModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(postLocalDocsOptionsModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(postLocalDocsOptionsModel.InclusiveEnd).To(Equal(core.BoolPtr(true)))
				Expect(postLocalDocsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postLocalDocsOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postLocalDocsOptionsModel.UpdateSeq).To(Equal(core.BoolPtr(true)))
				Expect(postLocalDocsOptionsModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(postLocalDocsOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(postLocalDocsOptionsModel.Keys).To(Equal([]string{"testString"}))
				Expect(postLocalDocsOptionsModel.Startkey).To(Equal(core.StringPtr("testString")))
				Expect(postLocalDocsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostLocalDocsQueriesOptions successfully`, func() {
				// Construct an instance of the AllDocsQuery model
				allDocsQueryModel := new(cloudantv1.AllDocsQuery)
				Expect(allDocsQueryModel).ToNot(BeNil())
				allDocsQueryModel.AttEncodingInfo = core.BoolPtr(true)
				allDocsQueryModel.Attachments = core.BoolPtr(true)
				allDocsQueryModel.Conflicts = core.BoolPtr(true)
				allDocsQueryModel.Descending = core.BoolPtr(true)
				allDocsQueryModel.IncludeDocs = core.BoolPtr(true)
				allDocsQueryModel.InclusiveEnd = core.BoolPtr(true)
				allDocsQueryModel.Limit = core.Int64Ptr(int64(0))
				allDocsQueryModel.Skip = core.Int64Ptr(int64(0))
				allDocsQueryModel.UpdateSeq = core.BoolPtr(true)
				allDocsQueryModel.Endkey = core.StringPtr("testString")
				allDocsQueryModel.Key = core.StringPtr("testString")
				allDocsQueryModel.Keys = []string{"testString"}
				allDocsQueryModel.Startkey = core.StringPtr("testString")
				Expect(allDocsQueryModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.InclusiveEnd).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(allDocsQueryModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(allDocsQueryModel.UpdateSeq).To(Equal(core.BoolPtr(true)))
				Expect(allDocsQueryModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(allDocsQueryModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(allDocsQueryModel.Keys).To(Equal([]string{"testString"}))
				Expect(allDocsQueryModel.Startkey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PostLocalDocsQueriesOptions model
				db := "testString"
				postLocalDocsQueriesOptionsModel := cloudantService.NewPostLocalDocsQueriesOptions(db)
				postLocalDocsQueriesOptionsModel.SetDb("testString")
				postLocalDocsQueriesOptionsModel.SetAccept("application/json")
				postLocalDocsQueriesOptionsModel.SetQueries([]cloudantv1.AllDocsQuery{*allDocsQueryModel})
				postLocalDocsQueriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postLocalDocsQueriesOptionsModel).ToNot(BeNil())
				Expect(postLocalDocsQueriesOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postLocalDocsQueriesOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(postLocalDocsQueriesOptionsModel.Queries).To(Equal([]cloudantv1.AllDocsQuery{*allDocsQueryModel}))
				Expect(postLocalDocsQueriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostMissingRevsOptions successfully`, func() {
				// Construct an instance of the PostMissingRevsOptions model
				db := "testString"
				postMissingRevsOptionsModel := cloudantService.NewPostMissingRevsOptions(db)
				postMissingRevsOptionsModel.SetDb("testString")
				postMissingRevsOptionsModel.SetDocumentRevisions(make(map[string][]string))
				postMissingRevsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postMissingRevsOptionsModel).ToNot(BeNil())
				Expect(postMissingRevsOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postMissingRevsOptionsModel.DocumentRevisions).To(Equal(make(map[string][]string)))
				Expect(postMissingRevsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostPartitionAllDocsOptions successfully`, func() {
				// Construct an instance of the PostPartitionAllDocsOptions model
				db := "testString"
				partitionKey := "testString"
				postPartitionAllDocsOptionsModel := cloudantService.NewPostPartitionAllDocsOptions(db, partitionKey)
				postPartitionAllDocsOptionsModel.SetDb("testString")
				postPartitionAllDocsOptionsModel.SetPartitionKey("testString")
				postPartitionAllDocsOptionsModel.SetAttEncodingInfo(true)
				postPartitionAllDocsOptionsModel.SetAttachments(true)
				postPartitionAllDocsOptionsModel.SetConflicts(true)
				postPartitionAllDocsOptionsModel.SetDescending(true)
				postPartitionAllDocsOptionsModel.SetIncludeDocs(true)
				postPartitionAllDocsOptionsModel.SetInclusiveEnd(true)
				postPartitionAllDocsOptionsModel.SetLimit(int64(0))
				postPartitionAllDocsOptionsModel.SetSkip(int64(0))
				postPartitionAllDocsOptionsModel.SetUpdateSeq(true)
				postPartitionAllDocsOptionsModel.SetEndkey("testString")
				postPartitionAllDocsOptionsModel.SetKey("testString")
				postPartitionAllDocsOptionsModel.SetKeys([]string{"testString"})
				postPartitionAllDocsOptionsModel.SetStartkey("testString")
				postPartitionAllDocsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postPartitionAllDocsOptionsModel).ToNot(BeNil())
				Expect(postPartitionAllDocsOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionAllDocsOptionsModel.PartitionKey).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionAllDocsOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionAllDocsOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionAllDocsOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionAllDocsOptionsModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionAllDocsOptionsModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionAllDocsOptionsModel.InclusiveEnd).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionAllDocsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postPartitionAllDocsOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postPartitionAllDocsOptionsModel.UpdateSeq).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionAllDocsOptionsModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionAllDocsOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionAllDocsOptionsModel.Keys).To(Equal([]string{"testString"}))
				Expect(postPartitionAllDocsOptionsModel.Startkey).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionAllDocsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostPartitionFindOptions successfully`, func() {
				// Construct an instance of the PostPartitionFindOptions model
				db := "testString"
				partitionKey := "testString"
				postPartitionFindOptionsModel := cloudantService.NewPostPartitionFindOptions(db, partitionKey)
				postPartitionFindOptionsModel.SetDb("testString")
				postPartitionFindOptionsModel.SetPartitionKey("testString")
				postPartitionFindOptionsModel.SetBookmark("testString")
				postPartitionFindOptionsModel.SetConflicts(true)
				postPartitionFindOptionsModel.SetExecutionStats(true)
				postPartitionFindOptionsModel.SetFields([]string{"testString"})
				postPartitionFindOptionsModel.SetLimit(int64(0))
				postPartitionFindOptionsModel.SetSelector(make(map[string]interface{}))
				postPartitionFindOptionsModel.SetSkip(int64(0))
				postPartitionFindOptionsModel.SetSort([]map[string]string{make(map[string]string)})
				postPartitionFindOptionsModel.SetStable(true)
				postPartitionFindOptionsModel.SetUpdate("false")
				postPartitionFindOptionsModel.SetUseIndex([]string{"testString"})
				postPartitionFindOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postPartitionFindOptionsModel).ToNot(BeNil())
				Expect(postPartitionFindOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionFindOptionsModel.PartitionKey).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionFindOptionsModel.Bookmark).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionFindOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionFindOptionsModel.ExecutionStats).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionFindOptionsModel.Fields).To(Equal([]string{"testString"}))
				Expect(postPartitionFindOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postPartitionFindOptionsModel.Selector).To(Equal(make(map[string]interface{})))
				Expect(postPartitionFindOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postPartitionFindOptionsModel.Sort).To(Equal([]map[string]string{make(map[string]string)}))
				Expect(postPartitionFindOptionsModel.Stable).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionFindOptionsModel.Update).To(Equal(core.StringPtr("false")))
				Expect(postPartitionFindOptionsModel.UseIndex).To(Equal([]string{"testString"}))
				Expect(postPartitionFindOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostPartitionSearchOptions successfully`, func() {
				// Construct an instance of the PostPartitionSearchOptions model
				db := "testString"
				partitionKey := "testString"
				ddoc := "testString"
				index := "testString"
				postPartitionSearchOptionsModel := cloudantService.NewPostPartitionSearchOptions(db, partitionKey, ddoc, index)
				postPartitionSearchOptionsModel.SetDb("testString")
				postPartitionSearchOptionsModel.SetPartitionKey("testString")
				postPartitionSearchOptionsModel.SetDdoc("testString")
				postPartitionSearchOptionsModel.SetIndex("testString")
				postPartitionSearchOptionsModel.SetBookmark("testString")
				postPartitionSearchOptionsModel.SetHighlightFields([]string{"testString"})
				postPartitionSearchOptionsModel.SetHighlightNumber(int64(1))
				postPartitionSearchOptionsModel.SetHighlightPostTag("testString")
				postPartitionSearchOptionsModel.SetHighlightPreTag("testString")
				postPartitionSearchOptionsModel.SetHighlightSize(int64(1))
				postPartitionSearchOptionsModel.SetIncludeDocs(true)
				postPartitionSearchOptionsModel.SetIncludeFields([]string{"testString"})
				postPartitionSearchOptionsModel.SetLimit(int64(3))
				postPartitionSearchOptionsModel.SetQuery("testString")
				postPartitionSearchOptionsModel.SetSort([]string{"testString"})
				postPartitionSearchOptionsModel.SetStale("ok")
				postPartitionSearchOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postPartitionSearchOptionsModel).ToNot(BeNil())
				Expect(postPartitionSearchOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionSearchOptionsModel.PartitionKey).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionSearchOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionSearchOptionsModel.Index).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionSearchOptionsModel.Bookmark).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionSearchOptionsModel.HighlightFields).To(Equal([]string{"testString"}))
				Expect(postPartitionSearchOptionsModel.HighlightNumber).To(Equal(core.Int64Ptr(int64(1))))
				Expect(postPartitionSearchOptionsModel.HighlightPostTag).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionSearchOptionsModel.HighlightPreTag).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionSearchOptionsModel.HighlightSize).To(Equal(core.Int64Ptr(int64(1))))
				Expect(postPartitionSearchOptionsModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionSearchOptionsModel.IncludeFields).To(Equal([]string{"testString"}))
				Expect(postPartitionSearchOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(3))))
				Expect(postPartitionSearchOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionSearchOptionsModel.Sort).To(Equal([]string{"testString"}))
				Expect(postPartitionSearchOptionsModel.Stale).To(Equal(core.StringPtr("ok")))
				Expect(postPartitionSearchOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostPartitionViewOptions successfully`, func() {
				// Construct an instance of the PostPartitionViewOptions model
				db := "testString"
				partitionKey := "testString"
				ddoc := "testString"
				view := "testString"
				postPartitionViewOptionsModel := cloudantService.NewPostPartitionViewOptions(db, partitionKey, ddoc, view)
				postPartitionViewOptionsModel.SetDb("testString")
				postPartitionViewOptionsModel.SetPartitionKey("testString")
				postPartitionViewOptionsModel.SetDdoc("testString")
				postPartitionViewOptionsModel.SetView("testString")
				postPartitionViewOptionsModel.SetAttEncodingInfo(true)
				postPartitionViewOptionsModel.SetAttachments(true)
				postPartitionViewOptionsModel.SetConflicts(true)
				postPartitionViewOptionsModel.SetDescending(true)
				postPartitionViewOptionsModel.SetIncludeDocs(true)
				postPartitionViewOptionsModel.SetInclusiveEnd(true)
				postPartitionViewOptionsModel.SetLimit(int64(0))
				postPartitionViewOptionsModel.SetSkip(int64(0))
				postPartitionViewOptionsModel.SetUpdateSeq(true)
				postPartitionViewOptionsModel.SetEndkey(core.StringPtr("testString"))
				postPartitionViewOptionsModel.SetEndkeyDocid("testString")
				postPartitionViewOptionsModel.SetGroup(true)
				postPartitionViewOptionsModel.SetGroupLevel(int64(1))
				postPartitionViewOptionsModel.SetKey(core.StringPtr("testString"))
				postPartitionViewOptionsModel.SetKeys([]interface{}{"testString"})
				postPartitionViewOptionsModel.SetReduce(true)
				postPartitionViewOptionsModel.SetStable(true)
				postPartitionViewOptionsModel.SetStartkey(core.StringPtr("testString"))
				postPartitionViewOptionsModel.SetStartkeyDocid("testString")
				postPartitionViewOptionsModel.SetUpdate("true")
				postPartitionViewOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postPartitionViewOptionsModel).ToNot(BeNil())
				Expect(postPartitionViewOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionViewOptionsModel.PartitionKey).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionViewOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionViewOptionsModel.View).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionViewOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionViewOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionViewOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionViewOptionsModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionViewOptionsModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionViewOptionsModel.InclusiveEnd).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionViewOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postPartitionViewOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postPartitionViewOptionsModel.UpdateSeq).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionViewOptionsModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionViewOptionsModel.EndkeyDocid).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionViewOptionsModel.Group).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionViewOptionsModel.GroupLevel).To(Equal(core.Int64Ptr(int64(1))))
				Expect(postPartitionViewOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionViewOptionsModel.Keys).To(Equal([]interface{}{"testString"}))
				Expect(postPartitionViewOptionsModel.Reduce).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionViewOptionsModel.Stable).To(Equal(core.BoolPtr(true)))
				Expect(postPartitionViewOptionsModel.Startkey).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionViewOptionsModel.StartkeyDocid).To(Equal(core.StringPtr("testString")))
				Expect(postPartitionViewOptionsModel.Update).To(Equal(core.StringPtr("true")))
				Expect(postPartitionViewOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostReplicateOptions successfully`, func() {
				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				Expect(attachmentModel).ToNot(BeNil())
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)
				Expect(attachmentModel.ContentType).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Data).To(Equal(CreateMockByteArray("This is a mock byte array value.")))
				Expect(attachmentModel.Digest).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.EncodedLength).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Encoding).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Follows).To(Equal(core.BoolPtr(true)))
				Expect(attachmentModel.Length).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Revpos).To(Equal(core.Int64Ptr(int64(1))))
				Expect(attachmentModel.Stub).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				Expect(revisionsModel).ToNot(BeNil())
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))
				Expect(revisionsModel.Ids).To(Equal([]string{"testString"}))
				Expect(revisionsModel.Start).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				Expect(documentRevisionStatusModel).ToNot(BeNil())
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")
				Expect(documentRevisionStatusModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentRevisionStatusModel.Status).To(Equal(core.StringPtr("available")))

				// Construct an instance of the ReplicationCreateTargetParameters model
				replicationCreateTargetParametersModel := new(cloudantv1.ReplicationCreateTargetParameters)
				Expect(replicationCreateTargetParametersModel).ToNot(BeNil())
				replicationCreateTargetParametersModel.N = core.Int64Ptr(int64(1))
				replicationCreateTargetParametersModel.Partitioned = core.BoolPtr(true)
				replicationCreateTargetParametersModel.Q = core.Int64Ptr(int64(1))
				Expect(replicationCreateTargetParametersModel.N).To(Equal(core.Int64Ptr(int64(1))))
				Expect(replicationCreateTargetParametersModel.Partitioned).To(Equal(core.BoolPtr(true)))
				Expect(replicationCreateTargetParametersModel.Q).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the ReplicationDatabaseAuthIam model
				replicationDatabaseAuthIamModel := new(cloudantv1.ReplicationDatabaseAuthIam)
				Expect(replicationDatabaseAuthIamModel).ToNot(BeNil())
				replicationDatabaseAuthIamModel.ApiKey = core.StringPtr("testString")
				Expect(replicationDatabaseAuthIamModel.ApiKey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplicationDatabaseAuth model
				replicationDatabaseAuthModel := new(cloudantv1.ReplicationDatabaseAuth)
				Expect(replicationDatabaseAuthModel).ToNot(BeNil())
				replicationDatabaseAuthModel.Iam = replicationDatabaseAuthIamModel
				Expect(replicationDatabaseAuthModel.Iam).To(Equal(replicationDatabaseAuthIamModel))

				// Construct an instance of the ReplicationDatabase model
				replicationDatabaseModel := new(cloudantv1.ReplicationDatabase)
				Expect(replicationDatabaseModel).ToNot(BeNil())
				replicationDatabaseModel.Auth = replicationDatabaseAuthModel
				replicationDatabaseModel.HeadersVar = make(map[string]string)
				replicationDatabaseModel.URL = core.StringPtr("testString")
				Expect(replicationDatabaseModel.Auth).To(Equal(replicationDatabaseAuthModel))
				Expect(replicationDatabaseModel.HeadersVar).To(Equal(make(map[string]string)))
				Expect(replicationDatabaseModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UserContext model
				userContextModel := new(cloudantv1.UserContext)
				Expect(userContextModel).ToNot(BeNil())
				userContextModel.Db = core.StringPtr("testString")
				userContextModel.Name = core.StringPtr("testString")
				userContextModel.Roles = []string{"_reader"}
				Expect(userContextModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(userContextModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(userContextModel.Roles).To(Equal([]string{"_reader"}))

				// Construct an instance of the ReplicationDocument model
				replicationDocumentModel := new(cloudantv1.ReplicationDocument)
				Expect(replicationDocumentModel).ToNot(BeNil())
				replicationDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				replicationDocumentModel.Conflicts = []string{"testString"}
				replicationDocumentModel.Deleted = core.BoolPtr(true)
				replicationDocumentModel.DeletedConflicts = []string{"testString"}
				replicationDocumentModel.ID = core.StringPtr("testString")
				replicationDocumentModel.LocalSeq = core.StringPtr("testString")
				replicationDocumentModel.Rev = core.StringPtr("testString")
				replicationDocumentModel.Revisions = revisionsModel
				replicationDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				replicationDocumentModel.Cancel = core.BoolPtr(true)
				replicationDocumentModel.CheckpointInterval = core.Int64Ptr(int64(0))
				replicationDocumentModel.ConnectionTimeout = core.Int64Ptr(int64(0))
				replicationDocumentModel.Continuous = core.BoolPtr(true)
				replicationDocumentModel.CreateTarget = core.BoolPtr(true)
				replicationDocumentModel.CreateTargetParams = replicationCreateTargetParametersModel
				replicationDocumentModel.DocIds = []string{"testString"}
				replicationDocumentModel.Filter = core.StringPtr("testString")
				replicationDocumentModel.HTTPConnections = core.Int64Ptr(int64(1))
				replicationDocumentModel.QueryParams = make(map[string]string)
				replicationDocumentModel.RetriesPerRequest = core.Int64Ptr(int64(0))
				replicationDocumentModel.Selector = make(map[string]interface{})
				replicationDocumentModel.SinceSeq = core.StringPtr("testString")
				replicationDocumentModel.SocketOptions = core.StringPtr("testString")
				replicationDocumentModel.Source = replicationDatabaseModel
				replicationDocumentModel.SourceProxy = core.StringPtr("testString")
				replicationDocumentModel.Target = replicationDatabaseModel
				replicationDocumentModel.TargetProxy = core.StringPtr("testString")
				replicationDocumentModel.UseCheckpoints = core.BoolPtr(true)
				replicationDocumentModel.UserCtx = userContextModel
				replicationDocumentModel.WorkerBatchSize = core.Int64Ptr(int64(1))
				replicationDocumentModel.WorkerProcesses = core.Int64Ptr(int64(1))
				replicationDocumentModel.Attachments["foo"] = *attachmentModel
				replicationDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				Expect(replicationDocumentModel.Conflicts).To(Equal([]string{"testString"}))
				Expect(replicationDocumentModel.Deleted).To(Equal(core.BoolPtr(true)))
				Expect(replicationDocumentModel.DeletedConflicts).To(Equal([]string{"testString"}))
				Expect(replicationDocumentModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.LocalSeq).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.Revisions).To(Equal(revisionsModel))
				Expect(replicationDocumentModel.RevsInfo).To(Equal([]cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}))
				Expect(replicationDocumentModel.Cancel).To(Equal(core.BoolPtr(true)))
				Expect(replicationDocumentModel.CheckpointInterval).To(Equal(core.Int64Ptr(int64(0))))
				Expect(replicationDocumentModel.ConnectionTimeout).To(Equal(core.Int64Ptr(int64(0))))
				Expect(replicationDocumentModel.Continuous).To(Equal(core.BoolPtr(true)))
				Expect(replicationDocumentModel.CreateTarget).To(Equal(core.BoolPtr(true)))
				Expect(replicationDocumentModel.CreateTargetParams).To(Equal(replicationCreateTargetParametersModel))
				Expect(replicationDocumentModel.DocIds).To(Equal([]string{"testString"}))
				Expect(replicationDocumentModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.HTTPConnections).To(Equal(core.Int64Ptr(int64(1))))
				Expect(replicationDocumentModel.QueryParams).To(Equal(make(map[string]string)))
				Expect(replicationDocumentModel.RetriesPerRequest).To(Equal(core.Int64Ptr(int64(0))))
				Expect(replicationDocumentModel.Selector).To(Equal(make(map[string]interface{})))
				Expect(replicationDocumentModel.SinceSeq).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.SocketOptions).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.Source).To(Equal(replicationDatabaseModel))
				Expect(replicationDocumentModel.SourceProxy).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.Target).To(Equal(replicationDatabaseModel))
				Expect(replicationDocumentModel.TargetProxy).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.UseCheckpoints).To(Equal(core.BoolPtr(true)))
				Expect(replicationDocumentModel.UserCtx).To(Equal(userContextModel))
				Expect(replicationDocumentModel.WorkerBatchSize).To(Equal(core.Int64Ptr(int64(1))))
				Expect(replicationDocumentModel.WorkerProcesses).To(Equal(core.Int64Ptr(int64(1))))
				Expect(replicationDocumentModel.GetProperties()).ToNot(BeEmpty())
				Expect(replicationDocumentModel.GetProperty("foo")).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(replicationDocumentModel.Attachments["foo"]).To(Equal(*attachmentModel))

				// Construct an instance of the PostReplicateOptions model
				postReplicateOptionsModel := cloudantService.NewPostReplicateOptions()
				postReplicateOptionsModel.SetReplicationDocument(replicationDocumentModel)
				postReplicateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postReplicateOptionsModel).ToNot(BeNil())
				Expect(postReplicateOptionsModel.ReplicationDocument).To(Equal(replicationDocumentModel))
				Expect(postReplicateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostRevsDiffOptions successfully`, func() {
				// Construct an instance of the PostRevsDiffOptions model
				db := "testString"
				postRevsDiffOptionsModel := cloudantService.NewPostRevsDiffOptions(db)
				postRevsDiffOptionsModel.SetDb("testString")
				postRevsDiffOptionsModel.SetDocumentRevisions(make(map[string][]string))
				postRevsDiffOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postRevsDiffOptionsModel).ToNot(BeNil())
				Expect(postRevsDiffOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postRevsDiffOptionsModel.DocumentRevisions).To(Equal(make(map[string][]string)))
				Expect(postRevsDiffOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostSearchAnalyzeOptions successfully`, func() {
				// Construct an instance of the PostSearchAnalyzeOptions model
				postSearchAnalyzeOptionsModel := cloudantService.NewPostSearchAnalyzeOptions()
				postSearchAnalyzeOptionsModel.SetAnalyzer("arabic")
				postSearchAnalyzeOptionsModel.SetText("testString")
				postSearchAnalyzeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postSearchAnalyzeOptionsModel).ToNot(BeNil())
				Expect(postSearchAnalyzeOptionsModel.Analyzer).To(Equal(core.StringPtr("arabic")))
				Expect(postSearchAnalyzeOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(postSearchAnalyzeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostSearchOptions successfully`, func() {
				// Construct an instance of the PostSearchOptions model
				db := "testString"
				ddoc := "testString"
				index := "testString"
				postSearchOptionsModel := cloudantService.NewPostSearchOptions(db, ddoc, index)
				postSearchOptionsModel.SetDb("testString")
				postSearchOptionsModel.SetDdoc("testString")
				postSearchOptionsModel.SetIndex("testString")
				postSearchOptionsModel.SetBookmark("testString")
				postSearchOptionsModel.SetHighlightFields([]string{"testString"})
				postSearchOptionsModel.SetHighlightNumber(int64(1))
				postSearchOptionsModel.SetHighlightPostTag("testString")
				postSearchOptionsModel.SetHighlightPreTag("testString")
				postSearchOptionsModel.SetHighlightSize(int64(1))
				postSearchOptionsModel.SetIncludeDocs(true)
				postSearchOptionsModel.SetIncludeFields([]string{"testString"})
				postSearchOptionsModel.SetLimit(int64(3))
				postSearchOptionsModel.SetQuery("testString")
				postSearchOptionsModel.SetSort([]string{"testString"})
				postSearchOptionsModel.SetStale("ok")
				postSearchOptionsModel.SetCounts([]string{"testString"})
				postSearchOptionsModel.SetDrilldown([][]string{[]string{"testString"}})
				postSearchOptionsModel.SetGroupField("testString")
				postSearchOptionsModel.SetGroupLimit(int64(1))
				postSearchOptionsModel.SetGroupSort([]string{"testString"})
				postSearchOptionsModel.SetRanges(make(map[string]map[string]map[string]string))
				postSearchOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postSearchOptionsModel).ToNot(BeNil())
				Expect(postSearchOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postSearchOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(postSearchOptionsModel.Index).To(Equal(core.StringPtr("testString")))
				Expect(postSearchOptionsModel.Bookmark).To(Equal(core.StringPtr("testString")))
				Expect(postSearchOptionsModel.HighlightFields).To(Equal([]string{"testString"}))
				Expect(postSearchOptionsModel.HighlightNumber).To(Equal(core.Int64Ptr(int64(1))))
				Expect(postSearchOptionsModel.HighlightPostTag).To(Equal(core.StringPtr("testString")))
				Expect(postSearchOptionsModel.HighlightPreTag).To(Equal(core.StringPtr("testString")))
				Expect(postSearchOptionsModel.HighlightSize).To(Equal(core.Int64Ptr(int64(1))))
				Expect(postSearchOptionsModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(postSearchOptionsModel.IncludeFields).To(Equal([]string{"testString"}))
				Expect(postSearchOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(3))))
				Expect(postSearchOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(postSearchOptionsModel.Sort).To(Equal([]string{"testString"}))
				Expect(postSearchOptionsModel.Stale).To(Equal(core.StringPtr("ok")))
				Expect(postSearchOptionsModel.Counts).To(Equal([]string{"testString"}))
				Expect(postSearchOptionsModel.Drilldown).To(Equal([][]string{[]string{"testString"}}))
				Expect(postSearchOptionsModel.GroupField).To(Equal(core.StringPtr("testString")))
				Expect(postSearchOptionsModel.GroupLimit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(postSearchOptionsModel.GroupSort).To(Equal([]string{"testString"}))
				Expect(postSearchOptionsModel.Ranges).To(Equal(make(map[string]map[string]map[string]string)))
				Expect(postSearchOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostViewOptions successfully`, func() {
				// Construct an instance of the PostViewOptions model
				db := "testString"
				ddoc := "testString"
				view := "testString"
				postViewOptionsModel := cloudantService.NewPostViewOptions(db, ddoc, view)
				postViewOptionsModel.SetDb("testString")
				postViewOptionsModel.SetDdoc("testString")
				postViewOptionsModel.SetView("testString")
				postViewOptionsModel.SetAttEncodingInfo(true)
				postViewOptionsModel.SetAttachments(true)
				postViewOptionsModel.SetConflicts(true)
				postViewOptionsModel.SetDescending(true)
				postViewOptionsModel.SetIncludeDocs(true)
				postViewOptionsModel.SetInclusiveEnd(true)
				postViewOptionsModel.SetLimit(int64(0))
				postViewOptionsModel.SetSkip(int64(0))
				postViewOptionsModel.SetUpdateSeq(true)
				postViewOptionsModel.SetEndkey(core.StringPtr("testString"))
				postViewOptionsModel.SetEndkeyDocid("testString")
				postViewOptionsModel.SetGroup(true)
				postViewOptionsModel.SetGroupLevel(int64(1))
				postViewOptionsModel.SetKey(core.StringPtr("testString"))
				postViewOptionsModel.SetKeys([]interface{}{"testString"})
				postViewOptionsModel.SetReduce(true)
				postViewOptionsModel.SetStable(true)
				postViewOptionsModel.SetStartkey(core.StringPtr("testString"))
				postViewOptionsModel.SetStartkeyDocid("testString")
				postViewOptionsModel.SetUpdate("true")
				postViewOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postViewOptionsModel).ToNot(BeNil())
				Expect(postViewOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postViewOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(postViewOptionsModel.View).To(Equal(core.StringPtr("testString")))
				Expect(postViewOptionsModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(postViewOptionsModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(postViewOptionsModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(postViewOptionsModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(postViewOptionsModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(postViewOptionsModel.InclusiveEnd).To(Equal(core.BoolPtr(true)))
				Expect(postViewOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postViewOptionsModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(postViewOptionsModel.UpdateSeq).To(Equal(core.BoolPtr(true)))
				Expect(postViewOptionsModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(postViewOptionsModel.EndkeyDocid).To(Equal(core.StringPtr("testString")))
				Expect(postViewOptionsModel.Group).To(Equal(core.BoolPtr(true)))
				Expect(postViewOptionsModel.GroupLevel).To(Equal(core.Int64Ptr(int64(1))))
				Expect(postViewOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(postViewOptionsModel.Keys).To(Equal([]interface{}{"testString"}))
				Expect(postViewOptionsModel.Reduce).To(Equal(core.BoolPtr(true)))
				Expect(postViewOptionsModel.Stable).To(Equal(core.BoolPtr(true)))
				Expect(postViewOptionsModel.Startkey).To(Equal(core.StringPtr("testString")))
				Expect(postViewOptionsModel.StartkeyDocid).To(Equal(core.StringPtr("testString")))
				Expect(postViewOptionsModel.Update).To(Equal(core.StringPtr("true")))
				Expect(postViewOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostViewQueriesOptions successfully`, func() {
				// Construct an instance of the ViewQuery model
				viewQueryModel := new(cloudantv1.ViewQuery)
				Expect(viewQueryModel).ToNot(BeNil())
				viewQueryModel.AttEncodingInfo = core.BoolPtr(true)
				viewQueryModel.Attachments = core.BoolPtr(true)
				viewQueryModel.Conflicts = core.BoolPtr(true)
				viewQueryModel.Descending = core.BoolPtr(true)
				viewQueryModel.IncludeDocs = core.BoolPtr(true)
				viewQueryModel.InclusiveEnd = core.BoolPtr(true)
				viewQueryModel.Limit = core.Int64Ptr(int64(0))
				viewQueryModel.Skip = core.Int64Ptr(int64(0))
				viewQueryModel.UpdateSeq = core.BoolPtr(true)
				viewQueryModel.Endkey = core.StringPtr("testString")
				viewQueryModel.EndkeyDocid = core.StringPtr("testString")
				viewQueryModel.Group = core.BoolPtr(true)
				viewQueryModel.GroupLevel = core.Int64Ptr(int64(1))
				viewQueryModel.Key = core.StringPtr("testString")
				viewQueryModel.Keys = []interface{}{"testString"}
				viewQueryModel.Reduce = core.BoolPtr(true)
				viewQueryModel.Stable = core.BoolPtr(true)
				viewQueryModel.Startkey = core.StringPtr("testString")
				viewQueryModel.StartkeyDocid = core.StringPtr("testString")
				viewQueryModel.Update = core.StringPtr("true")
				Expect(viewQueryModel.AttEncodingInfo).To(Equal(core.BoolPtr(true)))
				Expect(viewQueryModel.Attachments).To(Equal(core.BoolPtr(true)))
				Expect(viewQueryModel.Conflicts).To(Equal(core.BoolPtr(true)))
				Expect(viewQueryModel.Descending).To(Equal(core.BoolPtr(true)))
				Expect(viewQueryModel.IncludeDocs).To(Equal(core.BoolPtr(true)))
				Expect(viewQueryModel.InclusiveEnd).To(Equal(core.BoolPtr(true)))
				Expect(viewQueryModel.Limit).To(Equal(core.Int64Ptr(int64(0))))
				Expect(viewQueryModel.Skip).To(Equal(core.Int64Ptr(int64(0))))
				Expect(viewQueryModel.UpdateSeq).To(Equal(core.BoolPtr(true)))
				Expect(viewQueryModel.Endkey).To(Equal(core.StringPtr("testString")))
				Expect(viewQueryModel.EndkeyDocid).To(Equal(core.StringPtr("testString")))
				Expect(viewQueryModel.Group).To(Equal(core.BoolPtr(true)))
				Expect(viewQueryModel.GroupLevel).To(Equal(core.Int64Ptr(int64(1))))
				Expect(viewQueryModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(viewQueryModel.Keys).To(Equal([]interface{}{"testString"}))
				Expect(viewQueryModel.Reduce).To(Equal(core.BoolPtr(true)))
				Expect(viewQueryModel.Stable).To(Equal(core.BoolPtr(true)))
				Expect(viewQueryModel.Startkey).To(Equal(core.StringPtr("testString")))
				Expect(viewQueryModel.StartkeyDocid).To(Equal(core.StringPtr("testString")))
				Expect(viewQueryModel.Update).To(Equal(core.StringPtr("true")))

				// Construct an instance of the PostViewQueriesOptions model
				db := "testString"
				ddoc := "testString"
				view := "testString"
				postViewQueriesOptionsModel := cloudantService.NewPostViewQueriesOptions(db, ddoc, view)
				postViewQueriesOptionsModel.SetDb("testString")
				postViewQueriesOptionsModel.SetDdoc("testString")
				postViewQueriesOptionsModel.SetView("testString")
				postViewQueriesOptionsModel.SetQueries([]cloudantv1.ViewQuery{*viewQueryModel})
				postViewQueriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postViewQueriesOptionsModel).ToNot(BeNil())
				Expect(postViewQueriesOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(postViewQueriesOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(postViewQueriesOptionsModel.View).To(Equal(core.StringPtr("testString")))
				Expect(postViewQueriesOptionsModel.Queries).To(Equal([]cloudantv1.ViewQuery{*viewQueryModel}))
				Expect(postViewQueriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutAttachmentOptions successfully`, func() {
				// Construct an instance of the PutAttachmentOptions model
				db := "testString"
				docID := "testString"
				attachmentName := "testString"
				attachment := CreateMockReader("This is a mock file.")
				contentType := "testString"
				putAttachmentOptionsModel := cloudantService.NewPutAttachmentOptions(db, docID, attachmentName, attachment, contentType)
				putAttachmentOptionsModel.SetDb("testString")
				putAttachmentOptionsModel.SetDocID("testString")
				putAttachmentOptionsModel.SetAttachmentName("testString")
				putAttachmentOptionsModel.SetAttachment(CreateMockReader("This is a mock file."))
				putAttachmentOptionsModel.SetContentType("testString")
				putAttachmentOptionsModel.SetIfMatch("testString")
				putAttachmentOptionsModel.SetRev("testString")
				putAttachmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putAttachmentOptionsModel).ToNot(BeNil())
				Expect(putAttachmentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(putAttachmentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(putAttachmentOptionsModel.AttachmentName).To(Equal(core.StringPtr("testString")))
				Expect(putAttachmentOptionsModel.Attachment).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(putAttachmentOptionsModel.ContentType).To(Equal(core.StringPtr("testString")))
				Expect(putAttachmentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(putAttachmentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(putAttachmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutCloudantSecurityConfigurationOptions successfully`, func() {
				// Construct an instance of the SecurityObject model
				securityObjectModel := new(cloudantv1.SecurityObject)
				Expect(securityObjectModel).ToNot(BeNil())
				securityObjectModel.Names = []string{"testString"}
				securityObjectModel.Roles = []string{"testString"}
				Expect(securityObjectModel.Names).To(Equal([]string{"testString"}))
				Expect(securityObjectModel.Roles).To(Equal([]string{"testString"}))

				// Construct an instance of the PutCloudantSecurityConfigurationOptions model
				db := "testString"
				putCloudantSecurityConfigurationOptionsModel := cloudantService.NewPutCloudantSecurityConfigurationOptions(db)
				putCloudantSecurityConfigurationOptionsModel.SetDb("testString")
				putCloudantSecurityConfigurationOptionsModel.SetAdmins(securityObjectModel)
				putCloudantSecurityConfigurationOptionsModel.SetMembers(securityObjectModel)
				putCloudantSecurityConfigurationOptionsModel.SetCloudant(make(map[string][]string))
				putCloudantSecurityConfigurationOptionsModel.SetCouchdbAuthOnly(true)
				putCloudantSecurityConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putCloudantSecurityConfigurationOptionsModel).ToNot(BeNil())
				Expect(putCloudantSecurityConfigurationOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(putCloudantSecurityConfigurationOptionsModel.Admins).To(Equal(securityObjectModel))
				Expect(putCloudantSecurityConfigurationOptionsModel.Members).To(Equal(securityObjectModel))
				Expect(putCloudantSecurityConfigurationOptionsModel.Cloudant).To(Equal(make(map[string][]string)))
				Expect(putCloudantSecurityConfigurationOptionsModel.CouchdbAuthOnly).To(Equal(core.BoolPtr(true)))
				Expect(putCloudantSecurityConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutCorsConfigurationOptions successfully`, func() {
				// Construct an instance of the PutCorsConfigurationOptions model
				putCorsConfigurationOptionsModel := cloudantService.NewPutCorsConfigurationOptions()
				putCorsConfigurationOptionsModel.SetAllowCredentials(true)
				putCorsConfigurationOptionsModel.SetEnableCors(true)
				putCorsConfigurationOptionsModel.SetOrigins([]string{"testString"})
				putCorsConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putCorsConfigurationOptionsModel).ToNot(BeNil())
				Expect(putCorsConfigurationOptionsModel.AllowCredentials).To(Equal(core.BoolPtr(true)))
				Expect(putCorsConfigurationOptionsModel.EnableCors).To(Equal(core.BoolPtr(true)))
				Expect(putCorsConfigurationOptionsModel.Origins).To(Equal([]string{"testString"}))
				Expect(putCorsConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutDatabaseOptions successfully`, func() {
				// Construct an instance of the PutDatabaseOptions model
				db := "testString"
				putDatabaseOptionsModel := cloudantService.NewPutDatabaseOptions(db)
				putDatabaseOptionsModel.SetDb("testString")
				putDatabaseOptionsModel.SetPartitioned(true)
				putDatabaseOptionsModel.SetQ(int64(1))
				putDatabaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putDatabaseOptionsModel).ToNot(BeNil())
				Expect(putDatabaseOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(putDatabaseOptionsModel.Partitioned).To(Equal(core.BoolPtr(true)))
				Expect(putDatabaseOptionsModel.Q).To(Equal(core.Int64Ptr(int64(1))))
				Expect(putDatabaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutDesignDocumentOptions successfully`, func() {
				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				Expect(attachmentModel).ToNot(BeNil())
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)
				Expect(attachmentModel.ContentType).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Data).To(Equal(CreateMockByteArray("This is a mock byte array value.")))
				Expect(attachmentModel.Digest).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.EncodedLength).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Encoding).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Follows).To(Equal(core.BoolPtr(true)))
				Expect(attachmentModel.Length).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Revpos).To(Equal(core.Int64Ptr(int64(1))))
				Expect(attachmentModel.Stub).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				Expect(revisionsModel).ToNot(BeNil())
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))
				Expect(revisionsModel.Ids).To(Equal([]string{"testString"}))
				Expect(revisionsModel.Start).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				Expect(documentRevisionStatusModel).ToNot(BeNil())
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")
				Expect(documentRevisionStatusModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentRevisionStatusModel.Status).To(Equal(core.StringPtr("available")))

				// Construct an instance of the Analyzer model
				analyzerModel := new(cloudantv1.Analyzer)
				Expect(analyzerModel).ToNot(BeNil())
				analyzerModel.Name = core.StringPtr("classic")
				analyzerModel.Stopwords = []string{"testString"}
				Expect(analyzerModel.Name).To(Equal(core.StringPtr("classic")))
				Expect(analyzerModel.Stopwords).To(Equal([]string{"testString"}))

				// Construct an instance of the AnalyzerConfiguration model
				analyzerConfigurationModel := new(cloudantv1.AnalyzerConfiguration)
				Expect(analyzerConfigurationModel).ToNot(BeNil())
				analyzerConfigurationModel.Name = core.StringPtr("classic")
				analyzerConfigurationModel.Stopwords = []string{"testString"}
				analyzerConfigurationModel.Fields = make(map[string]cloudantv1.Analyzer)
				analyzerConfigurationModel.Fields["foo"] = *analyzerModel
				Expect(analyzerConfigurationModel.Name).To(Equal(core.StringPtr("classic")))
				Expect(analyzerConfigurationModel.Stopwords).To(Equal([]string{"testString"}))
				Expect(analyzerConfigurationModel.Fields["foo"]).To(Equal(*analyzerModel))

				// Construct an instance of the SearchIndexDefinition model
				searchIndexDefinitionModel := new(cloudantv1.SearchIndexDefinition)
				Expect(searchIndexDefinitionModel).ToNot(BeNil())
				searchIndexDefinitionModel.Analyzer = analyzerConfigurationModel
				searchIndexDefinitionModel.Index = core.StringPtr("testString")
				Expect(searchIndexDefinitionModel.Analyzer).To(Equal(analyzerConfigurationModel))
				Expect(searchIndexDefinitionModel.Index).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DesignDocumentOptions model
				designDocumentOptionsModel := new(cloudantv1.DesignDocumentOptions)
				Expect(designDocumentOptionsModel).ToNot(BeNil())
				designDocumentOptionsModel.Partitioned = core.BoolPtr(true)
				Expect(designDocumentOptionsModel.Partitioned).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the DesignDocumentViewsMapReduce model
				designDocumentViewsMapReduceModel := new(cloudantv1.DesignDocumentViewsMapReduce)
				Expect(designDocumentViewsMapReduceModel).ToNot(BeNil())
				designDocumentViewsMapReduceModel.Map = core.StringPtr("testString")
				designDocumentViewsMapReduceModel.Reduce = core.StringPtr("testString")
				Expect(designDocumentViewsMapReduceModel.Map).To(Equal(core.StringPtr("testString")))
				Expect(designDocumentViewsMapReduceModel.Reduce).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the GeoIndexDefinition model
				geoIndexDefinitionModel := new(cloudantv1.GeoIndexDefinition)
				Expect(geoIndexDefinitionModel).ToNot(BeNil())
				geoIndexDefinitionModel.Index = core.StringPtr("testString")
				Expect(geoIndexDefinitionModel.Index).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DesignDocument model
				designDocumentModel := new(cloudantv1.DesignDocument)
				Expect(designDocumentModel).ToNot(BeNil())
				designDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				designDocumentModel.Conflicts = []string{"testString"}
				designDocumentModel.Deleted = core.BoolPtr(true)
				designDocumentModel.DeletedConflicts = []string{"testString"}
				designDocumentModel.ID = core.StringPtr("testString")
				designDocumentModel.LocalSeq = core.StringPtr("testString")
				designDocumentModel.Rev = core.StringPtr("testString")
				designDocumentModel.Revisions = revisionsModel
				designDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				designDocumentModel.Autoupdate = core.BoolPtr(true)
				designDocumentModel.Filters = make(map[string]string)
				designDocumentModel.Indexes = make(map[string]cloudantv1.SearchIndexDefinition)
				designDocumentModel.Language = core.StringPtr("testString")
				designDocumentModel.Options = designDocumentOptionsModel
				designDocumentModel.Updates = make(map[string]string)
				designDocumentModel.ValidateDocUpdate = make(map[string]string)
				designDocumentModel.Views = make(map[string]cloudantv1.DesignDocumentViewsMapReduce)
				designDocumentModel.StIndexes = make(map[string]cloudantv1.GeoIndexDefinition)
				designDocumentModel.Attachments["foo"] = *attachmentModel
				designDocumentModel.Indexes["foo"] = *searchIndexDefinitionModel
				designDocumentModel.Views["foo"] = *designDocumentViewsMapReduceModel
				designDocumentModel.StIndexes["foo"] = *geoIndexDefinitionModel
				designDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				Expect(designDocumentModel.Conflicts).To(Equal([]string{"testString"}))
				Expect(designDocumentModel.Deleted).To(Equal(core.BoolPtr(true)))
				Expect(designDocumentModel.DeletedConflicts).To(Equal([]string{"testString"}))
				Expect(designDocumentModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(designDocumentModel.LocalSeq).To(Equal(core.StringPtr("testString")))
				Expect(designDocumentModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(designDocumentModel.Revisions).To(Equal(revisionsModel))
				Expect(designDocumentModel.RevsInfo).To(Equal([]cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}))
				Expect(designDocumentModel.Autoupdate).To(Equal(core.BoolPtr(true)))
				Expect(designDocumentModel.Filters).To(Equal(make(map[string]string)))
				Expect(designDocumentModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(designDocumentModel.Options).To(Equal(designDocumentOptionsModel))
				Expect(designDocumentModel.Updates).To(Equal(make(map[string]string)))
				Expect(designDocumentModel.ValidateDocUpdate).To(Equal(make(map[string]string)))
				Expect(designDocumentModel.GetProperties()).ToNot(BeEmpty())
				Expect(designDocumentModel.GetProperty("foo")).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(designDocumentModel.Attachments["foo"]).To(Equal(*attachmentModel))
				Expect(designDocumentModel.Indexes["foo"]).To(Equal(*searchIndexDefinitionModel))
				Expect(designDocumentModel.Views["foo"]).To(Equal(*designDocumentViewsMapReduceModel))
				Expect(designDocumentModel.StIndexes["foo"]).To(Equal(*geoIndexDefinitionModel))

				// Construct an instance of the PutDesignDocumentOptions model
				db := "testString"
				ddoc := "testString"
				putDesignDocumentOptionsModel := cloudantService.NewPutDesignDocumentOptions(db, ddoc)
				putDesignDocumentOptionsModel.SetDb("testString")
				putDesignDocumentOptionsModel.SetDdoc("testString")
				putDesignDocumentOptionsModel.SetDesignDocument(designDocumentModel)
				putDesignDocumentOptionsModel.SetIfMatch("testString")
				putDesignDocumentOptionsModel.SetBatch("ok")
				putDesignDocumentOptionsModel.SetNewEdits(true)
				putDesignDocumentOptionsModel.SetRev("testString")
				putDesignDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putDesignDocumentOptionsModel).ToNot(BeNil())
				Expect(putDesignDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(putDesignDocumentOptionsModel.Ddoc).To(Equal(core.StringPtr("testString")))
				Expect(putDesignDocumentOptionsModel.DesignDocument).To(Equal(designDocumentModel))
				Expect(putDesignDocumentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(putDesignDocumentOptionsModel.Batch).To(Equal(core.StringPtr("ok")))
				Expect(putDesignDocumentOptionsModel.NewEdits).To(Equal(core.BoolPtr(true)))
				Expect(putDesignDocumentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(putDesignDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutDocumentOptions successfully`, func() {
				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				Expect(attachmentModel).ToNot(BeNil())
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)
				Expect(attachmentModel.ContentType).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Data).To(Equal(CreateMockByteArray("This is a mock byte array value.")))
				Expect(attachmentModel.Digest).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.EncodedLength).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Encoding).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Follows).To(Equal(core.BoolPtr(true)))
				Expect(attachmentModel.Length).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Revpos).To(Equal(core.Int64Ptr(int64(1))))
				Expect(attachmentModel.Stub).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				Expect(revisionsModel).ToNot(BeNil())
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))
				Expect(revisionsModel.Ids).To(Equal([]string{"testString"}))
				Expect(revisionsModel.Start).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				Expect(documentRevisionStatusModel).ToNot(BeNil())
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")
				Expect(documentRevisionStatusModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentRevisionStatusModel.Status).To(Equal(core.StringPtr("available")))

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				Expect(documentModel).ToNot(BeNil())
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.Attachments["foo"] = *attachmentModel
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(documentModel.Conflicts).To(Equal([]string{"testString"}))
				Expect(documentModel.Deleted).To(Equal(core.BoolPtr(true)))
				Expect(documentModel.DeletedConflicts).To(Equal([]string{"testString"}))
				Expect(documentModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.LocalSeq).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Revisions).To(Equal(revisionsModel))
				Expect(documentModel.RevsInfo).To(Equal([]cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}))
				Expect(documentModel.GetProperties()).ToNot(BeEmpty())
				Expect(documentModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Attachments["foo"]).To(Equal(*attachmentModel))

				// Construct an instance of the PutDocumentOptions model
				db := "testString"
				docID := "testString"
				putDocumentOptionsModel := cloudantService.NewPutDocumentOptions(db, docID)
				putDocumentOptionsModel.SetDb("testString")
				putDocumentOptionsModel.SetDocID("testString")
				putDocumentOptionsModel.SetDocument(documentModel)
				putDocumentOptionsModel.SetBody(CreateMockReader("This is a mock file."))
				putDocumentOptionsModel.SetContentType("application/json")
				putDocumentOptionsModel.SetIfMatch("testString")
				putDocumentOptionsModel.SetBatch("ok")
				putDocumentOptionsModel.SetNewEdits(true)
				putDocumentOptionsModel.SetRev("testString")
				putDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putDocumentOptionsModel).ToNot(BeNil())
				Expect(putDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(putDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(putDocumentOptionsModel.Document).To(Equal(documentModel))
				Expect(putDocumentOptionsModel.Body).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(putDocumentOptionsModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(putDocumentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(putDocumentOptionsModel.Batch).To(Equal(core.StringPtr("ok")))
				Expect(putDocumentOptionsModel.NewEdits).To(Equal(core.BoolPtr(true)))
				Expect(putDocumentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(putDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutLocalDocumentOptions successfully`, func() {
				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				Expect(attachmentModel).ToNot(BeNil())
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)
				Expect(attachmentModel.ContentType).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Data).To(Equal(CreateMockByteArray("This is a mock byte array value.")))
				Expect(attachmentModel.Digest).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.EncodedLength).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Encoding).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Follows).To(Equal(core.BoolPtr(true)))
				Expect(attachmentModel.Length).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Revpos).To(Equal(core.Int64Ptr(int64(1))))
				Expect(attachmentModel.Stub).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				Expect(revisionsModel).ToNot(BeNil())
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))
				Expect(revisionsModel.Ids).To(Equal([]string{"testString"}))
				Expect(revisionsModel.Start).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				Expect(documentRevisionStatusModel).ToNot(BeNil())
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")
				Expect(documentRevisionStatusModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentRevisionStatusModel.Status).To(Equal(core.StringPtr("available")))

				// Construct an instance of the Document model
				documentModel := new(cloudantv1.Document)
				Expect(documentModel).ToNot(BeNil())
				documentModel.Attachments = make(map[string]cloudantv1.Attachment)
				documentModel.Conflicts = []string{"testString"}
				documentModel.Deleted = core.BoolPtr(true)
				documentModel.DeletedConflicts = []string{"testString"}
				documentModel.ID = core.StringPtr("testString")
				documentModel.LocalSeq = core.StringPtr("testString")
				documentModel.Rev = core.StringPtr("testString")
				documentModel.Revisions = revisionsModel
				documentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				documentModel.Attachments["foo"] = *attachmentModel
				documentModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(documentModel.Conflicts).To(Equal([]string{"testString"}))
				Expect(documentModel.Deleted).To(Equal(core.BoolPtr(true)))
				Expect(documentModel.DeletedConflicts).To(Equal([]string{"testString"}))
				Expect(documentModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.LocalSeq).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Revisions).To(Equal(revisionsModel))
				Expect(documentModel.RevsInfo).To(Equal([]cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}))
				Expect(documentModel.GetProperties()).ToNot(BeEmpty())
				Expect(documentModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))
				Expect(documentModel.Attachments["foo"]).To(Equal(*attachmentModel))

				// Construct an instance of the PutLocalDocumentOptions model
				db := "testString"
				docID := "testString"
				putLocalDocumentOptionsModel := cloudantService.NewPutLocalDocumentOptions(db, docID)
				putLocalDocumentOptionsModel.SetDb("testString")
				putLocalDocumentOptionsModel.SetDocID("testString")
				putLocalDocumentOptionsModel.SetDocument(documentModel)
				putLocalDocumentOptionsModel.SetBody(CreateMockReader("This is a mock file."))
				putLocalDocumentOptionsModel.SetContentType("application/json")
				putLocalDocumentOptionsModel.SetBatch("ok")
				putLocalDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putLocalDocumentOptionsModel).ToNot(BeNil())
				Expect(putLocalDocumentOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(putLocalDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(putLocalDocumentOptionsModel.Document).To(Equal(documentModel))
				Expect(putLocalDocumentOptionsModel.Body).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(putLocalDocumentOptionsModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(putLocalDocumentOptionsModel.Batch).To(Equal(core.StringPtr("ok")))
				Expect(putLocalDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutReplicationDocumentOptions successfully`, func() {
				// Construct an instance of the Attachment model
				attachmentModel := new(cloudantv1.Attachment)
				Expect(attachmentModel).ToNot(BeNil())
				attachmentModel.ContentType = core.StringPtr("testString")
				attachmentModel.Data = CreateMockByteArray("This is a mock byte array value.")
				attachmentModel.Digest = core.StringPtr("testString")
				attachmentModel.EncodedLength = core.Int64Ptr(int64(0))
				attachmentModel.Encoding = core.StringPtr("testString")
				attachmentModel.Follows = core.BoolPtr(true)
				attachmentModel.Length = core.Int64Ptr(int64(0))
				attachmentModel.Revpos = core.Int64Ptr(int64(1))
				attachmentModel.Stub = core.BoolPtr(true)
				Expect(attachmentModel.ContentType).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Data).To(Equal(CreateMockByteArray("This is a mock byte array value.")))
				Expect(attachmentModel.Digest).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.EncodedLength).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Encoding).To(Equal(core.StringPtr("testString")))
				Expect(attachmentModel.Follows).To(Equal(core.BoolPtr(true)))
				Expect(attachmentModel.Length).To(Equal(core.Int64Ptr(int64(0))))
				Expect(attachmentModel.Revpos).To(Equal(core.Int64Ptr(int64(1))))
				Expect(attachmentModel.Stub).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Revisions model
				revisionsModel := new(cloudantv1.Revisions)
				Expect(revisionsModel).ToNot(BeNil())
				revisionsModel.Ids = []string{"testString"}
				revisionsModel.Start = core.Int64Ptr(int64(1))
				Expect(revisionsModel.Ids).To(Equal([]string{"testString"}))
				Expect(revisionsModel.Start).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the DocumentRevisionStatus model
				documentRevisionStatusModel := new(cloudantv1.DocumentRevisionStatus)
				Expect(documentRevisionStatusModel).ToNot(BeNil())
				documentRevisionStatusModel.Rev = core.StringPtr("testString")
				documentRevisionStatusModel.Status = core.StringPtr("available")
				Expect(documentRevisionStatusModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(documentRevisionStatusModel.Status).To(Equal(core.StringPtr("available")))

				// Construct an instance of the ReplicationCreateTargetParameters model
				replicationCreateTargetParametersModel := new(cloudantv1.ReplicationCreateTargetParameters)
				Expect(replicationCreateTargetParametersModel).ToNot(BeNil())
				replicationCreateTargetParametersModel.N = core.Int64Ptr(int64(1))
				replicationCreateTargetParametersModel.Partitioned = core.BoolPtr(true)
				replicationCreateTargetParametersModel.Q = core.Int64Ptr(int64(1))
				Expect(replicationCreateTargetParametersModel.N).To(Equal(core.Int64Ptr(int64(1))))
				Expect(replicationCreateTargetParametersModel.Partitioned).To(Equal(core.BoolPtr(true)))
				Expect(replicationCreateTargetParametersModel.Q).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the ReplicationDatabaseAuthIam model
				replicationDatabaseAuthIamModel := new(cloudantv1.ReplicationDatabaseAuthIam)
				Expect(replicationDatabaseAuthIamModel).ToNot(BeNil())
				replicationDatabaseAuthIamModel.ApiKey = core.StringPtr("testString")
				Expect(replicationDatabaseAuthIamModel.ApiKey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplicationDatabaseAuth model
				replicationDatabaseAuthModel := new(cloudantv1.ReplicationDatabaseAuth)
				Expect(replicationDatabaseAuthModel).ToNot(BeNil())
				replicationDatabaseAuthModel.Iam = replicationDatabaseAuthIamModel
				Expect(replicationDatabaseAuthModel.Iam).To(Equal(replicationDatabaseAuthIamModel))

				// Construct an instance of the ReplicationDatabase model
				replicationDatabaseModel := new(cloudantv1.ReplicationDatabase)
				Expect(replicationDatabaseModel).ToNot(BeNil())
				replicationDatabaseModel.Auth = replicationDatabaseAuthModel
				replicationDatabaseModel.HeadersVar = make(map[string]string)
				replicationDatabaseModel.URL = core.StringPtr("testString")
				Expect(replicationDatabaseModel.Auth).To(Equal(replicationDatabaseAuthModel))
				Expect(replicationDatabaseModel.HeadersVar).To(Equal(make(map[string]string)))
				Expect(replicationDatabaseModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UserContext model
				userContextModel := new(cloudantv1.UserContext)
				Expect(userContextModel).ToNot(BeNil())
				userContextModel.Db = core.StringPtr("testString")
				userContextModel.Name = core.StringPtr("testString")
				userContextModel.Roles = []string{"_reader"}
				Expect(userContextModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(userContextModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(userContextModel.Roles).To(Equal([]string{"_reader"}))

				// Construct an instance of the ReplicationDocument model
				replicationDocumentModel := new(cloudantv1.ReplicationDocument)
				Expect(replicationDocumentModel).ToNot(BeNil())
				replicationDocumentModel.Attachments = make(map[string]cloudantv1.Attachment)
				replicationDocumentModel.Conflicts = []string{"testString"}
				replicationDocumentModel.Deleted = core.BoolPtr(true)
				replicationDocumentModel.DeletedConflicts = []string{"testString"}
				replicationDocumentModel.ID = core.StringPtr("testString")
				replicationDocumentModel.LocalSeq = core.StringPtr("testString")
				replicationDocumentModel.Rev = core.StringPtr("testString")
				replicationDocumentModel.Revisions = revisionsModel
				replicationDocumentModel.RevsInfo = []cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}
				replicationDocumentModel.Cancel = core.BoolPtr(true)
				replicationDocumentModel.CheckpointInterval = core.Int64Ptr(int64(0))
				replicationDocumentModel.ConnectionTimeout = core.Int64Ptr(int64(0))
				replicationDocumentModel.Continuous = core.BoolPtr(true)
				replicationDocumentModel.CreateTarget = core.BoolPtr(true)
				replicationDocumentModel.CreateTargetParams = replicationCreateTargetParametersModel
				replicationDocumentModel.DocIds = []string{"testString"}
				replicationDocumentModel.Filter = core.StringPtr("testString")
				replicationDocumentModel.HTTPConnections = core.Int64Ptr(int64(1))
				replicationDocumentModel.QueryParams = make(map[string]string)
				replicationDocumentModel.RetriesPerRequest = core.Int64Ptr(int64(0))
				replicationDocumentModel.Selector = make(map[string]interface{})
				replicationDocumentModel.SinceSeq = core.StringPtr("testString")
				replicationDocumentModel.SocketOptions = core.StringPtr("testString")
				replicationDocumentModel.Source = replicationDatabaseModel
				replicationDocumentModel.SourceProxy = core.StringPtr("testString")
				replicationDocumentModel.Target = replicationDatabaseModel
				replicationDocumentModel.TargetProxy = core.StringPtr("testString")
				replicationDocumentModel.UseCheckpoints = core.BoolPtr(true)
				replicationDocumentModel.UserCtx = userContextModel
				replicationDocumentModel.WorkerBatchSize = core.Int64Ptr(int64(1))
				replicationDocumentModel.WorkerProcesses = core.Int64Ptr(int64(1))
				replicationDocumentModel.Attachments["foo"] = *attachmentModel
				replicationDocumentModel.SetProperty("foo", map[string]interface{}{"anyKey": "anyValue"})
				Expect(replicationDocumentModel.Conflicts).To(Equal([]string{"testString"}))
				Expect(replicationDocumentModel.Deleted).To(Equal(core.BoolPtr(true)))
				Expect(replicationDocumentModel.DeletedConflicts).To(Equal([]string{"testString"}))
				Expect(replicationDocumentModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.LocalSeq).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.Revisions).To(Equal(revisionsModel))
				Expect(replicationDocumentModel.RevsInfo).To(Equal([]cloudantv1.DocumentRevisionStatus{*documentRevisionStatusModel}))
				Expect(replicationDocumentModel.Cancel).To(Equal(core.BoolPtr(true)))
				Expect(replicationDocumentModel.CheckpointInterval).To(Equal(core.Int64Ptr(int64(0))))
				Expect(replicationDocumentModel.ConnectionTimeout).To(Equal(core.Int64Ptr(int64(0))))
				Expect(replicationDocumentModel.Continuous).To(Equal(core.BoolPtr(true)))
				Expect(replicationDocumentModel.CreateTarget).To(Equal(core.BoolPtr(true)))
				Expect(replicationDocumentModel.CreateTargetParams).To(Equal(replicationCreateTargetParametersModel))
				Expect(replicationDocumentModel.DocIds).To(Equal([]string{"testString"}))
				Expect(replicationDocumentModel.Filter).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.HTTPConnections).To(Equal(core.Int64Ptr(int64(1))))
				Expect(replicationDocumentModel.QueryParams).To(Equal(make(map[string]string)))
				Expect(replicationDocumentModel.RetriesPerRequest).To(Equal(core.Int64Ptr(int64(0))))
				Expect(replicationDocumentModel.Selector).To(Equal(make(map[string]interface{})))
				Expect(replicationDocumentModel.SinceSeq).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.SocketOptions).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.Source).To(Equal(replicationDatabaseModel))
				Expect(replicationDocumentModel.SourceProxy).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.Target).To(Equal(replicationDatabaseModel))
				Expect(replicationDocumentModel.TargetProxy).To(Equal(core.StringPtr("testString")))
				Expect(replicationDocumentModel.UseCheckpoints).To(Equal(core.BoolPtr(true)))
				Expect(replicationDocumentModel.UserCtx).To(Equal(userContextModel))
				Expect(replicationDocumentModel.WorkerBatchSize).To(Equal(core.Int64Ptr(int64(1))))
				Expect(replicationDocumentModel.WorkerProcesses).To(Equal(core.Int64Ptr(int64(1))))
				Expect(replicationDocumentModel.GetProperties()).ToNot(BeEmpty())
				Expect(replicationDocumentModel.GetProperty("foo")).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(replicationDocumentModel.Attachments["foo"]).To(Equal(*attachmentModel))

				// Construct an instance of the PutReplicationDocumentOptions model
				docID := "testString"
				putReplicationDocumentOptionsModel := cloudantService.NewPutReplicationDocumentOptions(docID)
				putReplicationDocumentOptionsModel.SetDocID("testString")
				putReplicationDocumentOptionsModel.SetReplicationDocument(replicationDocumentModel)
				putReplicationDocumentOptionsModel.SetIfMatch("testString")
				putReplicationDocumentOptionsModel.SetBatch("ok")
				putReplicationDocumentOptionsModel.SetNewEdits(true)
				putReplicationDocumentOptionsModel.SetRev("testString")
				putReplicationDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putReplicationDocumentOptionsModel).ToNot(BeNil())
				Expect(putReplicationDocumentOptionsModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(putReplicationDocumentOptionsModel.ReplicationDocument).To(Equal(replicationDocumentModel))
				Expect(putReplicationDocumentOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(putReplicationDocumentOptionsModel.Batch).To(Equal(core.StringPtr("ok")))
				Expect(putReplicationDocumentOptionsModel.NewEdits).To(Equal(core.BoolPtr(true)))
				Expect(putReplicationDocumentOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(putReplicationDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutSecurityOptions successfully`, func() {
				// Construct an instance of the SecurityObject model
				securityObjectModel := new(cloudantv1.SecurityObject)
				Expect(securityObjectModel).ToNot(BeNil())
				securityObjectModel.Names = []string{"testString"}
				securityObjectModel.Roles = []string{"testString"}
				Expect(securityObjectModel.Names).To(Equal([]string{"testString"}))
				Expect(securityObjectModel.Roles).To(Equal([]string{"testString"}))

				// Construct an instance of the PutSecurityOptions model
				db := "testString"
				putSecurityOptionsModel := cloudantService.NewPutSecurityOptions(db)
				putSecurityOptionsModel.SetDb("testString")
				putSecurityOptionsModel.SetAdmins(securityObjectModel)
				putSecurityOptionsModel.SetMembers(securityObjectModel)
				putSecurityOptionsModel.SetCloudant(make(map[string][]string))
				putSecurityOptionsModel.SetCouchdbAuthOnly(true)
				putSecurityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putSecurityOptionsModel).ToNot(BeNil())
				Expect(putSecurityOptionsModel.Db).To(Equal(core.StringPtr("testString")))
				Expect(putSecurityOptionsModel.Admins).To(Equal(securityObjectModel))
				Expect(putSecurityOptionsModel.Members).To(Equal(securityObjectModel))
				Expect(putSecurityOptionsModel.Cloudant).To(Equal(make(map[string][]string)))
				Expect(putSecurityOptionsModel.CouchdbAuthOnly).To(Equal(core.BoolPtr(true)))
				Expect(putSecurityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSearchIndexDefinition successfully`, func() {
				index := "testString"
				model, err := cloudantService.NewSearchIndexDefinition(index)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
