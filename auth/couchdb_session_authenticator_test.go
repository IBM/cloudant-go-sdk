/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
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

package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"sync"
	"sync/atomic"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type contextKey string

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

		sessionAuth, ok := auth.(*CouchDbSessionAuthenticator)
		Expect(ok).To(BeTrue())
		Expect(sessionAuth.url).ToNot(BeZero())
		Expect(sessionAuth.disableSSLVerification).To(BeFalse())

		auth, err = GetAuthenticatorFromEnvironment("service2")
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))

		sessionAuth, ok = auth.(*CouchDbSessionAuthenticator)
		Expect(ok).To(BeTrue())
		Expect(sessionAuth.url).ToNot(BeZero())
		Expect(sessionAuth.disableSSLVerification).To(BeTrue())

		auth, err = GetAuthenticatorFromEnvironment("service3")
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(core.AUTHTYPE_IAM))
	})

	It("Create new Authenticator programmatically", func() {
		username, password := "user", "pass"
		auth, err := NewCouchDbSessionAuthenticator(username, password)
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))

		auth, err = NewCouchDbSessionAuthenticatorFromMap(map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		})
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))
	})

	It("Test Authenticator instantiation failures", func() {
		errortests := []struct {
			user, password string
		}{
			{"", "password"},
			{"{invalid-user}", "password"},
			{"user", ""},
			{"user", "{invalid-password}"},
		}

		for _, tt := range errortests {
			_, err := NewCouchDbSessionAuthenticator(tt.user, tt.password)
			Expect(err).To(HaveOccurred())

			_, err = NewCouchDbSessionAuthenticatorFromMap(map[string]string{
				"USERNAME": tt.user,
				"PASSWORD": tt.password,
			})
			Expect(err).To(HaveOccurred())
		}
	})

	Context("with standard test server", func() {
		var (
			server     *httptest.Server
			request    *http.Request
			auth       *CouchDbSessionAuthenticator
			err        error
			callNumber int32
		)

		BeforeEach(func() {
			callNumber = 0
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				atomic.AddInt32(&callNumber, 1)
				n := int(atomic.LoadInt32(&callNumber))
				cookie := &http.Cookie{
					Name:  "AuthSession",
					Value: fmt.Sprintf("fakefake-%d", n),
					Expires: time.Now().
						Add(24*time.Hour - time.Duration(n)*time.Minute),
				}
				http.SetCookie(w, cookie)
				w.WriteHeader(http.StatusOK)
			}))

			builder, err := core.NewRequestBuilder(core.GET).
				ResolveRequestURL(server.URL, "/db", nil)
			Expect(err).To(BeNil())

			ctx := context.WithValue(context.Background(), contextKey("key"), "abcdefgh")
			request, err = builder.
				AddHeader("X-Req-Id", "abcdefgh").
				WithContext(ctx).
				Build()
			Expect(err).To(BeNil())
			Expect(request).ToNot(BeNil())

			auth, err = NewCouchDbSessionAuthenticator("user", "pass")
			Expect(err).To(BeNil())
			Expect(auth).ToNot(BeNil())
			Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))
		})

		AfterEach(func() {
			server.Close()
		})

		It("Test setting URL, Headers and Context on Authenticator", func() {
			err = auth.Authenticate(request)
			Expect(err).To(BeNil())

			Expect(auth.url).To(Equal(server.URL))
			Expect(auth.header).To(HaveKeyWithValue("X-Req-Id", []string{"abcdefgh"}))
			Expect(auth.ctx.Value(contextKey("key"))).To(Equal("abcdefgh"))
		})

		It("Test setting custom http client on Authenticator", func() {
			// set http.Client with custom timeout
			auth.Client = &http.Client{Timeout: time.Second}

			err = auth.Authenticate(request)
			Expect(err).To(BeNil())

			Expect(auth.Client.Timeout).To(Equal(time.Second))
		})

		It("Test authentication re-request after expiration", func() {
			testDone := make(chan interface{})
			go func() {

				err = auth.Authenticate(request)
				Expect(err).To(BeNil())

				// Fetch a cookie from the cache to confirm it's present
				// and prepare cache state for the next test
				cookie, err := auth.getCookie()
				Expect(err).To(BeNil())
				Expect(cookie).ToNot(BeNil())
				Expect(cookie.Name).To(Equal("AuthSession"))
				Expect(cookie.Value).To(Equal("fakefake-1"))
				Expect(cookie.Expires).ToNot(BeNil())

				// Fetch cookie again to verify cache is working as expected,
				// we are getting old cookie and refresh process's not triggered
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
						defer GinkgoRecover()
						defer wg.Done()
						cookie, err := auth.getCookie()
						Expect(err).To(BeNil())
						Expect(cookie.Value).To(Equal("fakefake-2"))
						Expect(auth.session.refreshTime).ToNot(Equal(oldRefresh))
					}()
				}
				wg.Wait()
				close(testDone)
			}()
			Eventually(testDone, 1.0).Should(BeClosed())
		})

		It("Test authentication refresh", func() {
			testDone := make(chan interface{})
			go func() {
				err = auth.Authenticate(request)
				Expect(err).To(BeNil())

				// Test code path in getCookie() when needsRefresh() is false
				oldCookie, err := auth.getCookie()
				Expect(err).To(BeNil())
				Expect(oldCookie).ToNot(BeNil())
				Expect(oldCookie.Value).To(Equal("fakefake-1"))

				// Move time into the refresh window
				auth.session.refreshTime = time.Now().Add(-10 * time.Minute)

				// Run getCookie in three parallel threads, to verify
				// that we are still serving stale cached cookie
				// and request mutex works and we are querying /_session only once
				var refresherNumber int32
				var wg sync.WaitGroup

				for i := 1; i <= 3; i++ {
					wg.Add(1)
					go func() {
						defer GinkgoRecover()
						defer wg.Done()
						atomic.AddInt32(&refresherNumber, 1)
						cookie, err := auth.getCookie()
						// make sure that at least first refresh is async
						// and returns an old still-valid cookie
						if int(atomic.LoadInt32(&refresherNumber)) == 1 {
							Expect(err).To(BeNil())
							Expect(cookie).To(Equal(oldCookie))
							Expect(cookie.Value).To(Equal("fakefake-1"))
						}
					}()
				}
				wg.Wait()

				// wait for 1s (default duration) to confirm that eventually
				// we'll get a new cookie.
				Eventually(func() (*http.Cookie, error) {
					return auth.getCookie()
				}).ShouldNot(Equal(oldCookie))

				// wait a bit to confirm that we haven't had hits
				// from some late refresh process
				Consistently(func() int {
					return int(atomic.LoadInt32(&callNumber))
				}, "100ms", "100ms").Should(Equal(2))

				close(testDone)
			}()
			Eventually(testDone, 3.0).Should(BeClosed())
		})
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
			ResolveRequestURL(server.URL, "/db", nil)
		Expect(err).To(BeNil())

		request, err := builder.Build()
		Expect(err).To(BeNil())
		Expect(request).ToNot(BeNil())

		auth, err := NewCouchDbSessionAuthenticator("user", "pass")
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))

		err = auth.Authenticate(request)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal(string(authError)))
	})

	It("Test missing AuthSession cookie in the response", func() {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		builder, err := core.NewRequestBuilder(core.GET).
			ResolveRequestURL(server.URL, "/db", nil)
		Expect(err).To(BeNil())

		request, err := builder.Build()
		Expect(err).To(BeNil())
		Expect(request).ToNot(BeNil())

		auth, err := NewCouchDbSessionAuthenticator("user", "pass")
		Expect(err).To(BeNil())
		Expect(auth).ToNot(BeNil())
		Expect(auth.AuthenticationType()).To(Equal(AUTHTYPE_COUCHDB_SESSION))

		err = auth.Authenticate(request)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(HavePrefix("Missing AuthSession cookie"))
	})

	It("Test requestSession fails when auth URL is invalid", func() {
		auth := &CouchDbSessionAuthenticator{}
		_, err := auth.requestSession()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("service URL is empty"))
	})

	It("Test requestSession fails when server's down", func() {
		auth := &CouchDbSessionAuthenticator{}
		auth.url = "http://localhost"
		_, err := auth.requestSession()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).Should(HaveSuffix("connection refused"))
	})
})
