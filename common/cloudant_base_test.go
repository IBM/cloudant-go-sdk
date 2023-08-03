/**
 * © Copyright IBM Corporation 2021, 2022. All Rights Reserved.
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
	"net/url"
	"os"
	"path"
	"runtime"

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

	It("Create new Authenticator from environment", func() {
		pwd, err := os.Getwd()
		Expect(err).To(BeNil())
		credentialFilePath := path.Join(pwd, "/testdata/my-credentials.env")
		os.Setenv("IBM_CREDENTIALS_FILE", credentialFilePath)

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
		cloudant.BaseService.SetHTTPClient(client)

		Expect(cloudant).ToNot(BeNil())
		Expect(err).To(BeNil())
		Expect(cloudant.BaseService.Client.Jar).ToNot(BeNil())
		Expect(cloudant.BaseService.Client.Jar.Cookies(urlObj)[0]).Should(Equal(cookie))
	})
})
