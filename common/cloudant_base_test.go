/**
 * © Copyright IBM Corporation 2021. All Rights Reserved.
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
	"os"
	"path"

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
	})
})
