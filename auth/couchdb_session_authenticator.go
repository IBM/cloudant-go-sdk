/**
 * © Copyright IBM Corporation 2020, 2023. All Rights Reserved.
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
	"bytes"
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/IBM/go-sdk-core/v5/core"
)

const (
	AUTHTYPE_COUCHDB_SESSION = "COUCHDB_SESSION"
)

// CouchDbSessionAuthenticator uses username and password to obtain
// CouchDB authentication cookie, and adds the cookie to requests.
type CouchDbSessionAuthenticator struct {
	// [Required] The username and password used to access CouchDB session end-point
	Username, Password string

	// HTTP client used to to obtain CouchDB authentication cookie.
	client *http.Client

	// CouchDB URL inherited from the service request.
	URL string

	// Client's headers inherited from the service request.
	header http.Header

	// Context inherited from from the service request.
	ctx context.Context

	// A flag that indicates whether verification of the server's SSL certificate should be disabled
	DisableSSLVerification bool

	// A session instance that stores and manages the authentication cookie.
	session *session

	// A buffer chanel to hold on refreshed session.
	refresh chan *session

	// Authenticator mutex used in getCookie() to make it thread-safe to use from concurrent goroutines.
	mu sync.Mutex
}

// NewCouchDbSessionAuthenticator constructs a new NewCouchDbSessionAuthenticator instance.
func NewCouchDbSessionAuthenticator(username, password string) (*CouchDbSessionAuthenticator, error) {
	authenticator := &CouchDbSessionAuthenticator{
		Username: username,
		Password: password,
		refresh:  make(chan *session, 1),
	}
	if err := authenticator.Validate(); err != nil {
		return nil, err
	}
	client := core.DefaultHTTPClient()
	authenticator.SetClient(client)
	return authenticator, nil
}

// NewCouchDbSessionAuthenticatorFromMap constructs a new NewCouchDbSessionAuthenticator instance from a map.
func NewCouchDbSessionAuthenticatorFromMap(props map[string]string) (*CouchDbSessionAuthenticator, error) {
	if props == nil {
		return nil, fmt.Errorf(core.ERRORMSG_PROPS_MAP_NIL)
	}
	username := props[core.PROPNAME_USERNAME]
	password := props[core.PROPNAME_PASSWORD]
	return NewCouchDbSessionAuthenticator(username, password)
}

// AuthenticationType returns the authentication type for this authenticator.
func (a *CouchDbSessionAuthenticator) AuthenticationType() string {
	return AUTHTYPE_COUCHDB_SESSION
}

// Validate the authenticator's configuration.
// Ensures the service url, username and password are valid and not nil.
func (a *CouchDbSessionAuthenticator) Validate() error {
	if a.Username == "" {
		return fmt.Errorf(core.ERRORMSG_PROP_MISSING, "Username")
	}

	if a.Password == "" {
		return fmt.Errorf(core.ERRORMSG_PROP_MISSING, "Password")
	}

	if core.HasBadFirstOrLastChar(a.Username) {
		return fmt.Errorf(core.ERRORMSG_PROP_INVALID, "Username")
	}

	if core.HasBadFirstOrLastChar(a.Password) {
		return fmt.Errorf(core.ERRORMSG_PROP_INVALID, "Password")
	}

	return nil
}

// Authenticate adds session authentication cookie to a request.
func (a *CouchDbSessionAuthenticator) Authenticate(request *http.Request) error {
	a.URL = request.URL.Scheme + "://" + request.URL.Host
	a.header = request.Header
	a.ctx = request.Context()

	cookie, err := a.refreshCookie()
	if err != nil {
		return err
	}

	if a.client.Jar == nil && a.session != nil {
		request.AddCookie(cookie)
	}

	return err
}

// SetClient sets the http client for the authenticator.
func (a *CouchDbSessionAuthenticator) SetClient(client *http.Client) {
	a.client = client
}

// refreshCookie checks if an AuthSession cookie needs to be refreshed.
// A new cookie will be fetched and returned from the session end-point
// when needed.
func (a *CouchDbSessionAuthenticator) refreshCookie() (*http.Cookie, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.flushRefreshChannel()

	if a.session == nil || !a.session.isValid() {
		newSession, err := a.requestSession()
		if err != nil {
			return nil, err
		}
		a.session = newSession
	} else if a.session.needsRefresh() {
		// start a background process to refresh the session.
		// the refreshed session will be passed to a buffered channel
		// and updated in a next request at flushRefreshChannel() call.
		go func() {
			// we are intentionally not returning errors to the parent process
			// to avoid raisng error to a client with still valid session.
			session, err := a.requestSession()
			if err != nil {
				return
			}
			a.refresh <- session
		}()
	}

	return a.session.getCookie(), nil
}

// flushRefreshChannel drains authenticator's refresh channel
// and updates session var with instance from the channel.
// This is none-blocking no-op call when channel's empty.
func (a *CouchDbSessionAuthenticator) flushRefreshChannel() {
	select {
	case session := <-a.refresh:
		a.session = session
	default:
	}
}

// requestSession fetches new AuthSession cookie from the server.
func (a *CouchDbSessionAuthenticator) requestSession() (*session, error) {
	builder, err := core.NewRequestBuilder(core.POST).
		ResolveRequestURL(a.URL, "/_session", nil)
	if err != nil {
		return nil, err
	}

	builder.AddHeader(core.CONTENT_TYPE, "application/x-www-form-urlencoded").
		AddFormData("name", "", "", a.Username).
		AddFormData("password", "", "", a.Password).
		WithContext(a.ctx)

	// set all the unique headers from original request's client
	for key, value := range a.header {
		if _, ok := builder.Header[key]; !ok {
			builder.Header[key] = value
		}
	}

	req, err := builder.Build()
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(a.Username, a.Password)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		buff := new(bytes.Buffer)
		_, _ = buff.ReadFrom(resp.Body)

		detailedResponse := &core.DetailedResponse{
			StatusCode: resp.StatusCode,
			Headers:    resp.Header,
			RawResult:  buff.Bytes(),
		}
		err := fmt.Errorf(buff.String())
		return nil, core.NewAuthenticationError(detailedResponse, err)
	}

	var session *session
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "AuthSession" {
			session, err = newSession(cookie)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	if session == nil {
		return nil, fmt.Errorf("Missing AuthSession cookie in the response")
	}

	return session, nil
}
