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

package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"sync"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authenticator Unit Tests", func() {
	It("Create new Authenticator from environment", func() {
		pwd, err := os.Getwd()
		Expect(err).To(BeNil())
		credentialFilePath := path.Join(pwd, "/testdata/my-credentials.env")
		os.Setenv("IBM_CREDENTIALS_FILE", credentialFilePath)

		auth, err := GetAuthenticatorFromEnvironment("service1")
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))

		auth, err = GetAuthenticatorFromEnvironment("service2")
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(core.AUTHTYPE_IAM))
	})

	It("Create new Authenticator programmatically", func() {
		url, username, password := "http://couch", "user", "pass"
		auth, err := NewCouchDbSessionAuthenticator(url, username, password)
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))

		auth, err = NewCouchDbSessionAuthenticatorFromMap(map[string]string{
			"URL":      url,
			"USERNAME": username,
			"PASSWORD": password,
		})
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))
	})

	It("Test Authenticator instantiation failures", func() {
		errortests := []struct {
			url, user, password string
		}{
			{"", "user", "password"},
			{"localhost", "user", "password"},
			{"https://localhost", "", "password"},
			{"https://localhost", "{invalid-user}", "password"},
			{"https://localhost", "user", ""},
			{"https://localhost", "user", "{invalid-password}"},
		}

		for _, tt := range errortests {
			_, err := NewCouchDbSessionAuthenticator(tt.url, tt.user, tt.password)
			Expect(err).To(HaveOccurred())

			_, err = NewCouchDbSessionAuthenticatorFromMap(map[string]string{
				"URL":      tt.url,
				"USERNAME": tt.user,
				"PASSWORD": tt.password,
			})
			Expect(err).To(HaveOccurred())
		}
	})

	Context("with standard test server", func() {
		var (
			server  *httptest.Server
			request *http.Request
			auth    *CouchDbSessionAuthenticator
			err     error
		)

		BeforeEach(func() {
			callNumber := 0
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				callNumber++
				cookie := &http.Cookie{
					Name:  "AuthSession",
					Value: fmt.Sprintf("fakefake-%d", callNumber),
					Expires: time.Now().
						Add(24*time.Hour - time.Duration(callNumber)*time.Minute),
				}
				http.SetCookie(w, cookie)
				w.WriteHeader(http.StatusOK)
			}))

			builder, err := core.NewRequestBuilder(core.GET).
				ResolveRequestURL("https://localhost", "/db", nil)
			Expect(err).To(BeNil())

			request, err = builder.Build()
			Expect(err).To(BeNil())
			Expect(request).ToNot(BeNil())

			auth, err = NewCouchDbSessionAuthenticator(server.URL, "user", "pass")
			Expect(err).To(BeNil())
			Expect(auth).ToNot(BeNil())
			Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))
		})

		AfterEach(func() {
			server.Close()
		})

		It("Test setting custom http client on Authenticator", func() {
			// set http.Client with custom timeout
			auth.Client = &http.Client{Timeout: time.Second}

			// call to Authenticate sets authenticator's client
			err = auth.Authenticate(request)
			Expect(err).To(BeNil())

			Expect(auth.Client.Timeout).To(Equal(time.Second))
		})

		It("Test authentication re-request after expiration", func(done Done) {
			err = auth.Authenticate(request)
			Expect(err).To(BeNil())

			cookie, err := auth.getCookie()
			Expect(err).To(BeNil())
			Expect(cookie).ToNot(BeNil())

			Expect(cookie.Name).To(Equal("AuthSession"))
			Expect(cookie.Value).To(Equal("fakefake-1"))
			Expect(cookie.Expires).ToNot(BeNil())

			// Fetch cookie again and verify we got old cookie from cache
			oldRefresh := auth.session.refreshTime
			cookie, err = auth.getCookie()
			Expect(err).To(BeNil())
			Expect(cookie.Value).To(Equal("fakefake-1"))
			Expect(auth.session.refreshTime).To(Equal(oldRefresh))

			// Force expiration and verify we got a new cookie
			auth.session.expires = time.Now().Add(-time.Minute)

			// Run getCookie in three parallel threads, to verify
			// that request mutex works and we are querying /_session only once
			var wg sync.WaitGroup

			for i := 1; i <= 3; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					cookie, err = auth.getCookie()
					Expect(err).To(BeNil())
					Expect(cookie.Value).To(Equal("fakefake-2"))
					Expect(auth.session.refreshTime).ToNot(Equal(oldRefresh))
				}()
			}
			wg.Wait()
			close(done)
		})

		It("Test authentication refresh", func(done Done) {
			err = auth.Authenticate(request)
			Expect(err).To(BeNil())

			cookie, err := auth.getCookie()
			Expect(err).To(BeNil())
			Expect(cookie).ToNot(BeNil())
			Expect(cookie.Value).To(Equal("fakefake-1"))

			// Move time into the refresh window
			auth.session.expires = time.Now().Add(30 * time.Minute)
			auth.session.refreshTime = time.Now().Add(-10 * time.Minute)

			// Run getCookie in three parallel threads, to verify
			// that we are still serving stale cached cookie
			// and request mutex works and we are querying /_session only once
			var wg sync.WaitGroup

			for i := 1; i <= 3; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					c, err := auth.getCookie()
					Expect(err).To(BeNil())
					Expect(c).To(Equal(auth.session.cookie))
					Expect(auth.session.cookie.Value).To(Equal("fakefake-1"))
				}()
			}
			wg.Wait()

			// give refresher goroutine time to finish
			time.Sleep(1 * time.Second)

			Expect(cookie).ToNot(Equal(auth.session.cookie))
			Expect(auth.session.cookie.Value).To(Equal("fakefake-2"))

			close(done)
		}, 3.0)
	})

	It("Test authentication failures", func() {
		authError := []byte(`
			{
				"error":"unauthorized",
				"reason":"You are not authorized to access this db."
			}
		`)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(authError)
		}))
		defer server.Close()

		builder, err := core.NewRequestBuilder(core.GET).
			ResolveRequestURL("https://localhost", "/db", nil)
		Expect(err).To(BeNil())

		request, err := builder.Build()
		Expect(err).To(BeNil())
		Expect(request).ToNot(BeNil())

		auth, err := NewCouchDbSessionAuthenticator(server.URL, "user", "pass")
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))

		err = auth.Authenticate(request)
		Expect(err).To(HaveOccurred())
		Expect(string(authError)).To(Equal(err.Error()))
	})
})
