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
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
)

const (
	AUTHTYPE_COUCHDB_SESSION = "COUCHDB_SESSION"
)

var requestSessionMutex sync.Mutex

// CouchDbSessionAuthenticator uses username and password to obtain
// CouchDB authentication cookie, and adds the cookie to requests.
type CouchDbSessionAuthenticator struct {
	// [Required] The username and password used to access CouchDB session end-point
	Username, Password string

	// [Optional] The http.Client object used to to obtain CouchDB authentication cookie.
	// If not specified by the user, a suitable default Client will be constructed.
	Client *http.Client

	// CouchDB URL inherited from the service config.
	url string

	// Client's headers inherited from the service request.
	header http.Header

	// Context inherited from from the service request.
	ctx context.Context

	// A flag that indicates whether verification of the server's SSL certificate should be disabled; INherired from the service config
	disableSSLVerification bool

	// A session instance that stores and manages the authentication cookie.
	session *session
}

// NewCouchDbSessionAuthenticator constructs a new NewCouchDbSessionAuthenticator instance.
func NewCouchDbSessionAuthenticator(username, password string) (*CouchDbSessionAuthenticator, error) {
	authenticator := &CouchDbSessionAuthenticator{
		Username: username,
		Password: password,
	}
	if err := authenticator.Validate(); err != nil {
		return nil, err
	}
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

// GetAuthenticatorFromEnvironment instantiates an Authenticator
// using service properties retrieved from external config sources.
func GetAuthenticatorFromEnvironment(credentialKey string) (core.Authenticator, error) {
	props, err := core.GetServiceProperties(credentialKey)
	if err != nil {
		return nil, err
	}
	authType, ok := props[core.PROPNAME_AUTH_TYPE]
	if ok && strings.EqualFold(authType, AUTHTYPE_COUCHDB_SESSION) {
		authenticator, err := NewCouchDbSessionAuthenticatorFromMap(props)
		if url, ok := props[core.PROPNAME_SVC_URL]; ok && url != "" {
			authenticator.url = url
		}
		if disableSSLVerification, ok := props[core.PROPNAME_SVC_DISABLE_SSL]; ok && disableSSLVerification != "" {
			boolValue, err := strconv.ParseBool(disableSSLVerification)
			if err == nil && boolValue {
				authenticator.disableSSLVerification = true
			}
		}
		return authenticator, err
	}

	return core.GetAuthenticatorFromEnvironment(credentialKey)
}

// AuthenticationType returns the authentication type for this authenticator.
func (a CouchDbSessionAuthenticator) AuthenticationType() string {
	return AUTHTYPE_COUCHDB_SESSION
}

// Validate the authenticator's configuration.
// Ensures the service url, username and password are valid and not nil.
func (a CouchDbSessionAuthenticator) Validate() error {
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

	a.url = request.URL.Scheme + "://" + request.URL.Host
	a.header = request.Header
	a.ctx = request.Context()

	cookie, err := a.getCookie()
	if err != nil {
		return err
	}

	request.AddCookie(cookie)
	return nil
}

// getCookie returns an AuthSession cookie to be used in a request.
// A new cookie will be fetched from the session end-point when needed.
func (a *CouchDbSessionAuthenticator) getCookie() (*http.Cookie, error) {
	if a.session == nil || !a.session.isValid() {
		err := a.syncRequestSession()
		if err != nil {
			return nil, err
		}
	} else if a.session.needsRefresh() {
		ch := make(chan error)
		go func() {
			ch <- a.requestSession()
		}()
		select {
		case err := <-ch:
			if err != nil {
				return nil, err
			}
		default:
		}
	}

	return a.session.getCookie(), nil
}

// syncRequestSession synchronously checks if the current
// Session cookie in cache is valid. If cookie is not valid
// or does not exist, it'll fetch it from session end-point.
func (a *CouchDbSessionAuthenticator) syncRequestSession() error {
	requestSessionMutex.Lock()
	defer requestSessionMutex.Unlock()

	if a.session != nil && a.session.isValid() {
		return nil
	}

	err := a.requestSession()
	return err
}

// requestSession fetches new AuthSession cookie from the server.
func (a *CouchDbSessionAuthenticator) requestSession() error {
	builder, err := core.NewRequestBuilder(core.POST).
		ResolveRequestURL(a.url, "/_session", nil)
	if err != nil {
		return err
	}

	builder.AddHeader(core.CONTENT_TYPE, core.DEFAULT_CONTENT_TYPE).
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
		return err
	}

	req.SetBasicAuth(a.Username, a.Password)

	if a.Client == nil {
		a.Client = &http.Client{
			Timeout: time.Second * 30,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: a.disableSSLVerification},
			},
		}
	}

	resp, err := a.Client.Do(req)
	if err != nil {
		return err
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
		return core.NewAuthenticationError(detailedResponse, err)
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "AuthSession" {
			a.session, err = newSession(cookie)
			if err != nil {
				return err
			}
			break
		}
	}

	if a.session == nil {
		return fmt.Errorf("Missing AuthSession coookie in the response")
	}

	return nil
}
