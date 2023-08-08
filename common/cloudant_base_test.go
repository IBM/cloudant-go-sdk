/**
 * © Copyright IBM Corporation 2021, 2023. All Rights Reserved.
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

package common

import (
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/IBM/cloudant-go-sdk/auth"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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

		sdkHeaders := GetSdkHeaders("cloudant", "V1", "GetDocument")
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

		sdkHeaders := GetSdkHeaders("cloudant", "V1", "GetDocumentAsStream")
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

		sdkHeaders := GetSdkHeaders("cloudant", "V1", "GetDocument")
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

		sdkHeaders := GetSdkHeaders("cloudant", "V1", "GetDocument")
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
			sdkHeaders := GetSdkHeaders("cloudant", "V1", "GetDatabase")
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
})
