/**
 * © Copyright IBM Corporation 2021, 2024. All Rights Reserved.
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

package base

import (
	"encoding/json"
	"errors"
	"io"
	"maps"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/IBM/cloudant-go-sdk/auth"
	"github.com/IBM/cloudant-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var expectedErrType *core.SDKProblem

var _ = Describe(`Cloudant custom base service UT`, func() {
	It("Validates a doc ID", func() {
		cloudant, err := NewBaseService(&core.ServiceOptions{
			URL:           "https://cloudant.example",
			Authenticator: &core.NoAuthAuthenticator{},
		})
		Expect(cloudant).ToNot(BeNil())
		Expect(err).To(BeNil())

		pathParamsMap := map[string]string{
			"db":     "testDatabase",
			"doc_id": "_testDocument",
		}

		builder := core.NewRequestBuilder(core.GET)
		_, err = builder.ResolveRequestURL(cloudant.Options.URL, `/{db}/{doc_id}`, pathParamsMap)
		if err != nil {
			return
		}

		sdkHeaders := common.GetSdkHeaders("cloudant", "V1", "GetDocument")
		for headerName, headerValue := range sdkHeaders {
			builder.AddHeader(headerName, headerValue)
		}

		request, err := builder.Build()
		Expect(request).ToNot(BeNil())
		Expect(err).To(BeNil())

		response, err := cloudant.Request(request, nil)
		Expect(response).To(BeNil())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("_testDocument"))
		Expect(errors.As(err, &expectedErrType)).To(BeTrue())
	})

	It("Validates a doc ID with GetDocumentAsStream", func() {
		cloudant, err := NewBaseService(&core.ServiceOptions{
			URL:           "https://cloudant.example",
			Authenticator: &core.NoAuthAuthenticator{},
		})
		Expect(cloudant).ToNot(BeNil())
		Expect(err).To(BeNil())

		pathParamsMap := map[string]string{
			"db":     "testDatabase",
			"doc_id": "_testDocument",
		}

		builder := core.NewRequestBuilder(core.GET)
		_, err = builder.ResolveRequestURL(cloudant.Options.URL, `/{db}/{doc_id}`, pathParamsMap)
		if err != nil {
			return
		}

		sdkHeaders := common.GetSdkHeaders("cloudant", "V1", "GetDocumentAsStream")
		for headerName, headerValue := range sdkHeaders {
			builder.AddHeader(headerName, headerValue)
		}

		request, err := builder.Build()
		Expect(request).ToNot(BeNil())
		Expect(err).To(BeNil())

		response, err := cloudant.Request(request, nil)
		Expect(response).To(BeNil())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("_testDocument"))
		Expect(errors.As(err, &expectedErrType)).To(BeTrue())
	})

	It("Validates a doc ID at long service path", func() {
		cloudant, err := NewBaseService(&core.ServiceOptions{
			URL:           "https://cloudant.example/some/proxy/path",
			Authenticator: &core.NoAuthAuthenticator{},
		})
		Expect(cloudant).ToNot(BeNil())
		Expect(err).To(BeNil())

		pathParamsMap := map[string]string{
			"db":     "testDatabase",
			"doc_id": "_testDocument",
		}

		builder := core.NewRequestBuilder(core.GET)
		_, err = builder.ResolveRequestURL(cloudant.Options.URL, `/{db}/{doc_id}`, pathParamsMap)
		if err != nil {
			return
		}

		sdkHeaders := common.GetSdkHeaders("cloudant", "V1", "GetDocument")
		for headerName, headerValue := range sdkHeaders {
			builder.AddHeader(headerName, headerValue)
		}

		request, err := builder.Build()
		Expect(request).ToNot(BeNil())
		Expect(err).To(BeNil())

		response, err := cloudant.Request(request, nil)
		Expect(response).To(BeNil())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("_testDocument"))
		Expect(errors.As(err, &expectedErrType)).To(BeTrue())
	})

	It("Validates a doc ID at long service path after change", func() {
		cloudant, err := NewBaseService(&core.ServiceOptions{
			URL:           "https://cloudant.example/",
			Authenticator: &core.NoAuthAuthenticator{},
		})
		Expect(cloudant).ToNot(BeNil())
		Expect(err).To(BeNil())

		_ = cloudant.SetServiceURL("https://cloudant.example/some/proxy/path")

		pathParamsMap := map[string]string{
			"db":     "testDatabase",
			"doc_id": "_testDocument",
		}

		builder := core.NewRequestBuilder(core.GET)
		_, err = builder.ResolveRequestURL(cloudant.Options.URL, `/{db}/{doc_id}`, pathParamsMap)
		if err != nil {
			return
		}

		sdkHeaders := common.GetSdkHeaders("cloudant", "V1", "GetDocument")
		for headerName, headerValue := range sdkHeaders {
			builder.AddHeader(headerName, headerValue)
		}

		request, err := builder.Build()
		Expect(request).ToNot(BeNil())
		Expect(err).To(BeNil())

		response, err := cloudant.Request(request, nil)
		Expect(response).To(BeNil())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("_testDocument"))
		Expect(errors.As(err, &expectedErrType)).To(BeTrue())
	})

	It("Validates URL trailing slash is stripped", func() {
		cloudant, err := NewBaseService(&core.ServiceOptions{
			URL:           "https://cloudant.example/",
			Authenticator: &core.NoAuthAuthenticator{},
		})
		Expect(cloudant).ToNot(BeNil())
		Expect(err).To(BeNil())

		builder := core.NewRequestBuilder(core.GET)
		_, err = builder.ResolveRequestURL(cloudant.Options.URL, "/db", nil)
		Expect(err).To(BeNil())
		Expect(builder.URL.String()).To(Equal("https://cloudant.example/db"))
	})

	It("Validates User-Agent is set", func() {
		cloudant, err := NewBaseService(&core.ServiceOptions{
			URL:           "https://cloudant.example",
			Authenticator: &core.NoAuthAuthenticator{},
		})
		Expect(cloudant).ToNot(BeNil())
		Expect(err).To(BeNil())

		Expect(cloudant.UserAgent).ToNot(BeNil())
		Expect(cloudant.UserAgent).To(MatchRegexp("\\w+\\/[0-9.]+\\s\\(.+\\)"))
		Expect(cloudant.UserAgent).To(HavePrefix("cloudant-go-sdk"))
		Expect(cloudant.UserAgent).To(ContainSubstring(runtime.Version()))
		Expect(cloudant.UserAgent).To(ContainSubstring(runtime.GOOS))
		Expect(cloudant.UserAgent).To(ContainSubstring(runtime.GOARCH))
	})

	Context("with IBM_CREDENTIALS env variable", func() {
		BeforeEach(func() {
			pwd, err := os.Getwd()
			Expect(err).To(BeNil())
			credentialFilePath := path.Join(pwd, "/testdata/my-credentials.env")
			err = os.Setenv("IBM_CREDENTIALS_FILE", credentialFilePath)
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			err := os.Unsetenv("IBM_CREDENTIALS_FILE")
			Expect(err).To(BeNil())
		})

		It("Create new Authenticator from environment", func() {
			authenticator, err := GetAuthenticatorFromEnvironment("service1")
			Expect(err).To(BeNil())
			Expect(authenticator).ToNot(BeNil())
			Expect(authenticator.AuthenticationType()).To(Equal(auth.AUTHTYPE_COUCHDB_SESSION))

			sessionAuth, ok := authenticator.(*auth.CouchDbSessionAuthenticator)
			Expect(ok).To(BeTrue())
			Expect(sessionAuth.URL).ToNot(BeZero())
			Expect(sessionAuth.DisableSSLVerification).To(BeFalse())

			authenticator, err = GetAuthenticatorFromEnvironment("service2")
			Expect(err).To(BeNil())
			Expect(authenticator).ToNot(BeNil())
			Expect(authenticator.AuthenticationType()).To(Equal(auth.AUTHTYPE_COUCHDB_SESSION))

			sessionAuth, ok = authenticator.(*auth.CouchDbSessionAuthenticator)
			Expect(ok).To(BeTrue())
			Expect(sessionAuth.URL).ToNot(BeZero())
			Expect(sessionAuth.DisableSSLVerification).To(BeTrue())

			authenticator, err = GetAuthenticatorFromEnvironment("service3")
			Expect(err).To(BeNil())
			Expect(authenticator).ToNot(BeNil())
			Expect(authenticator.AuthenticationType()).To(Equal(core.AUTHTYPE_IAM))

			authenticator, err = GetAuthenticatorFromEnvironment("service4")
			Expect(err).To(BeNil())
			Expect(authenticator).ToNot(BeNil())
			Expect(authenticator.AuthenticationType()).To(Equal(core.AUTHTYPE_IAM))
		})

		It("Validates cookie jar enabled for all auths", func() {
			couchDbAuth, err := GetAuthenticatorFromEnvironment("service1")
			Expect(err).To(BeNil())
			iamAuth, err := GetAuthenticatorFromEnvironment("service3")
			Expect(err).To(BeNil())
			basicAuth, err := GetAuthenticatorFromEnvironment("service5")
			Expect(err).To(BeNil())
			authList := []core.Authenticator{
				&core.NoAuthAuthenticator{},
				couchDbAuth,
				iamAuth,
				basicAuth,
			}
			for _, item := range authList {
				cloudant, err := NewBaseService(&core.ServiceOptions{
					URL:           "https://cloudant.example",
					Authenticator: item,
				})
				Expect(cloudant).ToNot(BeNil())
				Expect(err).To(BeNil())
				Expect(cloudant.BaseService.Client.Jar).ToNot(BeNil())

				// set custom client
				client := &http.Client{Timeout: time.Second}
				cloudant.SetHTTPClient(client)
				Expect(cloudant.BaseService.Client.Jar).ToNot(BeNil())
				Expect(cloudant.BaseService.Client.Timeout).To(Equal(time.Second))
			}
		})

		It("Validates custom cookie jar", func() {
			client := core.DefaultHTTPClient()

			jar, err := cookiejar.New(nil)
			Expect(err).To(BeNil())
			// create cookie for custom jar
			urlObj, _ := url.Parse("http://localhost:8080/")
			cookie := &http.Cookie{
				Name:  "token",
				Value: "some_token",
			}
			jar.SetCookies(urlObj, []*http.Cookie{cookie})
			client.Jar = jar

			iamAuth, err := GetAuthenticatorFromEnvironment("service3")
			Expect(err).To(BeNil())
			cloudant, err := NewBaseService(&core.ServiceOptions{
				URL:           "https://cloudant.example",
				Authenticator: iamAuth,
			})
			cloudant.SetHTTPClient(client)

			Expect(cloudant).ToNot(BeNil())
			Expect(err).To(BeNil())
			Expect(cloudant.BaseService.Client.Jar).ToNot(BeNil())
			Expect(cloudant.BaseService.Client.Jar.Cookies(urlObj)[0]).Should(Equal(cookie))
		})

		It("Validates cookie jar used for CouchDB Session auths", func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/_session":
					if _, err := r.Cookie("AuthSession"); err != nil {
						http.SetCookie(w, &http.Cookie{
							Name:  "AuthSession",
							Value: "fakefake",
						})
					}
					w.WriteHeader(http.StatusOK)
				case "/db":
					w.WriteHeader(http.StatusOK)
				default:
					w.WriteHeader(http.StatusNotFound)
				}
			}))
			defer server.Close()

			auth, err := GetAuthenticatorFromEnvironment("service1")
			Expect(err).To(BeNil())
			Expect(auth).ToNot(BeNil())

			cloudant, err := NewBaseService(&core.ServiceOptions{
				URL:           server.URL,
				Authenticator: auth,
			})
			Expect(err).To(BeNil())
			Expect(cloudant).ToNot(BeNil())
			Expect(cloudant.BaseService.Client.Jar).ToNot(BeNil())

			builder := core.NewRequestBuilder(core.GET)
			_, err = builder.ResolveRequestURL(server.URL, "/db", nil)
			Expect(err).To(BeNil())
			sdkHeaders := common.GetSdkHeaders("cloudant", "V1", "GetDatabase")
			for headerName, headerValue := range sdkHeaders {
				builder.AddHeader(headerName, headerValue)
			}

			request, err := builder.Build()
			Expect(request).ToNot(BeNil())
			Expect(err).To(BeNil())

			_, err = cloudant.Request(request, nil)
			Expect(err).To(BeNil())

			u, err := url.Parse(server.URL)
			Expect(err).To(BeNil())
			var cookie *http.Cookie
			for _, c := range cloudant.BaseService.Client.Jar.Cookies(u) {
				if c.Name == "AuthSession" {
					cookie = c
					break
				}
			}
			Expect(cookie).ToNot(BeNil())
			Expect(cookie.Value).To(Equal("fakefake"))
		})

		It("Validates SetHTTPClient passes URL to CouchDB Session auth", func() {
			a, err := auth.NewCouchDbSessionAuthenticator("foo", "bar")
			Expect(err).To(BeNil())
			Expect(a.URL).To(Equal(""))

			cloudant, err := NewBaseService(&core.ServiceOptions{
				URL:           "https://localhost:8080",
				Authenticator: a,
			})
			Expect(cloudant).ToNot(BeNil())
			Expect(err).To(BeNil())

			newUrl := "https://cloudant.example:5984"
			err = cloudant.SetServiceURL(newUrl)
			Expect(err).To(BeNil())

			Expect(cloudant.BaseService.GetServiceURL()).To(Equal(newUrl))
			Expect(a.URL).To(Equal(newUrl))
		})
	})

	Context("augmentation error tests", func() {
		var (
			service *BaseService
		)

		BeforeEach(func() {
			var err error
			service, err = NewBaseService(&core.ServiceOptions{
				URL:           "https://~replace-with-cloudant-host~.cloudantnosqldb.appdomain.cloud",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(service).ToNot(BeNil())
			Expect(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			service = nil
		})

		It("Validates that ErrorResponse was added", func() {
			var expect *ErrorResponse
			Expect(service.GetHTTPClient().Transport).To(BeAssignableToTypeOf(expect))
		})

		It("Validates that ErrorResponse added to new http clients", func() {
			before := service.GetHTTPClient().Transport

			client := core.DefaultHTTPClient()
			service.SetHTTPClient(client)

			after := service.GetHTTPClient().Transport

			var expect *ErrorResponse
			Expect(after).To(BeAssignableToTypeOf(expect))
			Expect(after).ShouldNot(BeIdenticalTo(before))
		})

		It("Validates that ErrorResponse added only once", func() {
			client := service.GetHTTPClient()
			before := client.Transport
			service.SetHTTPClient(client)

			after := service.GetHTTPClient().Transport

			var expect *ErrorResponse
			Expect(after).To(BeAssignableToTypeOf(expect))
			Expect(after).Should(BeIdenticalTo(before))
		})

		type testConf struct {
			description string
			method      string
			status      int
			headers     map[string]string
			body        string
			expect      string
			stream      bool
		}

		testSuite := func(cfg testConf) {
			It(cfg.description, func() {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					for key, value := range cfg.headers {
						w.Header().Set(key, value)
					}
					if cfg.stream {
						// should be set automatically, but just in case
						w.Header().Set("Transfer-Encoding", "chunked")
						w.Header().Set("X-Content-Type-Options", "nosniff")
					}
					w.WriteHeader(cfg.status)
					if len(cfg.body) > 0 {
						if cfg.stream {
							flusher, ok := w.(http.Flusher)
							Expect(ok).Should(BeTrue())
							for _, line := range strings.Split(cfg.body, "\n") {
								_, _ = w.Write([]byte(strings.TrimSpace(line) + "\n"))
								flusher.Flush()
							}
						} else {
							_, _ = w.Write([]byte(cfg.body))
						}
					}
				}))
				Expect(server).ToNot(BeNil())
				defer server.Close()

				client := service.GetHTTPClient()
				if cfg.method == "" {
					cfg.method = http.MethodGet
				}
				req, err := http.NewRequest(cfg.method, server.URL+"/testdb/testdoc", nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp, err := client.Do(req)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(resp.StatusCode).To(Equal(cfg.status))

				for key, value := range cfg.headers {
					Expect(resp.Header.Values(key)).To(ContainElement(value))
				}

				// this covers both normal and AsStream responses
				bodyBytes, err := io.ReadAll(resp.Body)
				Expect(err).ShouldNot(HaveOccurred())
				defer resp.Body.Close()

				body := make(map[string]interface{})
				err = json.Unmarshal(bodyBytes, &body)
				// test JSON response
				if err == nil {
					expect := make(map[string]interface{})
					err = json.Unmarshal([]byte(cfg.expect), &expect)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(body).To(Equal(expect))
				} else {
					// vs literal equivalence
					Expect(bodyBytes).To(Equal([]byte(cfg.expect)))
				}
			})
		}

		contentTypeHeader := map[string]string{"content-type": "application/json"}
		requestIdHeader := map[string]string{"x-couch-request-id": "testreqid"}
		defaultHeaders := maps.Clone(contentTypeHeader)
		maps.Copy(defaultHeaders, requestIdHeader)

		errorOnlyBody := `{"error": "test_value"}`
		errorReasonBody := `{"error": "test_value", "reason": "A valid test reason"}`

		for _, cfg := range []testConf{
			{
				description: "Validates augmented error without reason without trace",
				status:      http.StatusTeapot,
				headers:     contentTypeHeader,
				body:        errorOnlyBody,
				expect: `{
					"error": "test_value",
					"errors": [{
						"code": "test_value",
						"message": "test_value"
					}]
				}`,
			},
			{
				description: "Validates augmented error without reason with trace",
				status:      http.StatusTeapot,
				headers:     defaultHeaders,
				body:        errorOnlyBody,
				expect: `{
					"trace": "testreqid",
					"error": "test_value",
					"errors": [{
						"code": "test_value",
						"message": "test_value"
					}]
				}`,
			},
			{
				description: "Validates augmented error with reason without trace",
				status:      http.StatusTeapot,
				headers:     contentTypeHeader,
				body:        errorReasonBody,
				expect: `{
					"error": "test_value",
					"reason": "A valid test reason",
					"errors": [{
						"code": "test_value",
						"message": "test_value: A valid test reason"
					}]
				}`,
			},
			{
				description: "Validates augmented error with reason and trace",
				status:      http.StatusTeapot,
				headers:     defaultHeaders,
				body:        errorReasonBody,
				expect: `{
					"trace": "testreqid",
					"error": "test_value",
					"reason": "A valid test reason",
					"errors": [{
						"code": "test_value",
						"message": "test_value: A valid test reason"
					}]
				}`,
			},
			{
				description: "Validates augmented error with reason and trace as a stream",
				status:      http.StatusTeapot,
				headers:     defaultHeaders,
				stream:      true,
				body: `{
					"error": "test_value",
					"reason": "A valid test reason"
				}`,
				expect: `{
					"trace": "testreqid",
					"error": "test_value",
					"reason": "A valid test reason",
					"errors": [{
						"code": "test_value",
						"message": "test_value: A valid test reason"
					}]
				}`,
			},
			{
				description: "Validates augmented error with reason and trace with json charset",
				status:      http.StatusTeapot,
				headers: map[string]string{
					"x-couch-request-id": "testreqid",
					"content-type":       "application/json; charset=utf-8",
				},
				body: errorReasonBody,
				expect: `{
					"trace": "testreqid",
					"error": "test_value",
					"reason": "A valid test reason",
					"errors": [{
						"code": "test_value",
						"message": "test_value: A valid test reason"
					}]
				}`,
			},
			{
				description: "Validates no augmentation on successfull response",
				status:      http.StatusOK,
				headers:     defaultHeaders,
				body:        `{"_id": "testdoc", "_rev": "1-abc", "foo": "bar"}`,
				expect:      `{"_id": "testdoc", "_rev": "1-abc", "foo": "bar"}`,
			},
			{
				description: "Validates no augmentation on HEAD request",
				status:      http.StatusTeapot,
				method:      http.MethodHead,
				headers:     defaultHeaders,
			},
			{
				description: "Validates no augmentation on empty body",
				status:      http.StatusTeapot,
				headers:     defaultHeaders,
				body:        `{}`,
				expect:      `{}`,
			},
			{
				description: "Validates no augmentation of existing trace",
				status:      http.StatusTooManyRequests,
				headers:     defaultHeaders,
				body: `{
					"trace": "testanotherreqid",
					"error": "too_many_requests",
					"reason": "Buy a bigger plan."
				}`,
				expect: `{
					"trace": "testanotherreqid",
					"error": "too_many_requests",
					"reason": "Buy a bigger plan."
				}`,
			},
			{
				description: "Validates no augmentation of existing 'errors' without trace",
				status:      http.StatusForbidden,
				headers:     contentTypeHeader,
				body: `{
					"error": "forbidden",
					"reason": "You must have _reader to access this resource.",
					"errors": [{
						"code": "forbidden",
						"message": "forbidden: You must have _reader to access this resource."
					}]
				}`,
				expect: `{
					"error": "forbidden",
					"reason": "You must have _reader to access this resource.",
					"errors": [{
						"code": "forbidden",
						"message": "forbidden: You must have _reader to access this resource."
					}]
				}`,
			},
			{
				description: "Validates augmented trace with no augmentation of existing 'errors'",
				status:      http.StatusForbidden,
				headers:     defaultHeaders,
				body: `{
					"error": "forbidden",
					"reason": "You must have _reader to access this resource.",
					"errors": [{
						"code": "forbidden",
						"message": "forbidden: You must have _reader to access this resource."
					}]
				}`,
				expect: `{
					"trace": "testreqid",
					"error": "forbidden",
					"reason": "You must have _reader to access this resource.",
					"errors": [{
						"code": "forbidden",
						"message": "forbidden: You must have _reader to access this resource."
					}]
				}`,
			},
			{
				description: "Validates no augmentation on none json response",
				status:      http.StatusBadRequest,
				headers: map[string]string{
					"x-couch-request-id": "testreqid",
					"content-type":       "text/plain",
				},
				body:   `foo`,
				expect: `foo`,
			},
			{
				description: "FIXME! Validates no augmentation on missing content type",
				status:      http.StatusTeapot,
				headers:     requestIdHeader,
				body:        `000`,
				expect:      `000`,
			},
			{
				description: "Validates no augmentation on missing 'error' in response with trace",
				status:      http.StatusBadRequest,
				headers:     defaultHeaders,
				body: `{
					"foo": "bar",
					"reason": "testing"
				}`,
				expect: `{
					"foo": "bar",
					"reason": "testing"
				}`,
			},
			{
				description: "Validates no augmentation on missing 'error' in response without trace",
				status:      http.StatusBadRequest,
				headers:     contentTypeHeader,
				body: `{
					"foo": "bar",
					"reason": "testing"
				}`,
				expect: `{
					"foo": "bar",
					"reason": "testing"
				}`,
			},
			{
				description: "Validates no augmentation on broken json",
				status:      http.StatusBadRequest,
				headers:     contentTypeHeader,
				body:        `{"err`,
				expect:      `{"err`,
			},
			{
				description: "Validates augmented on empty reason with trace",
				status:      http.StatusTeapot,
				headers:     defaultHeaders,
				body: `{
					"error": "test_value",
					"reason": ""
				}`,
				expect: `{
					"trace": "testreqid",
					"error": "test_value",
					"reason": "",
					"errors": [{
						"code": "test_value",
						"message": "test_value"
					}]
				}`,
			},
		} {
			if len(cfg.headers) > 0 {
				testSuite(cfg)
			}
		}
	})
})
