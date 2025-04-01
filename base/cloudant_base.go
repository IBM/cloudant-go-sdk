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
	"fmt"
	"net/http"
	"net/http/cookiejar"
	neturl "net/url"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/cloudant-go-sdk/auth"
	"github.com/IBM/cloudant-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
	"golang.org/x/net/publicsuffix"
)

type BaseService struct {
	serviceUrlPathSegmentsSize int
	*core.BaseService
}

type validationRule struct {
	pathSegmentIndex   int
	errorParameterName string
	operationIds       []string
}

var docIdRule = validationRule{
	pathSegmentIndex:   1,
	errorParameterName: "Document ID",
	operationIds: []string{
		"DeleteDocument",
		"GetDocument",
		"GetDocumentAsMixed",
		"GetDocumentAsRelated",
		"GetDocumentAsStream",
		"HeadDocument",
		"PutDocument",
		"DeleteAttachment",
		"GetAttachment",
		"HeadAttachment",
		"PutAttachment",
	},
}
var attIdRule = validationRule{
	pathSegmentIndex:   2,
	errorParameterName: "Attachment name",
	operationIds: []string{
		"DeleteAttachment",
		"GetAttachment",
		"HeadAttachment",
		"PutAttachment",
	},
}

var validationRules = [...]*validationRule{&docIdRule, &attIdRule}
var rulesByOperation = make(map[string][]*validationRule)

func init() {
	for _, rule := range validationRules {
		// Build a map of operation name to a list of validations
		for _, operationId := range rule.operationIds {
			ruleExists := false
			rules, ok := rulesByOperation[operationId]
			if ok {
				// There are already some rules for this operationId
				// Check if the current rule already exists
				for _, existingRule := range rules {
					if existingRule == rule {
						ruleExists = true
						break
					}
				}
			}
			if !ruleExists {
				// The rule didn't exist, append it
				rulesByOperation[operationId] = append(rules, rule)
			}
		}
	}
}

func NewBaseService(opts *core.ServiceOptions) (*BaseService, error) {
	baseService, err := core.NewBaseService(opts)
	if err != nil {
		return &BaseService{}, err
	}

	// Set a default value for the User-Agent http header.
	baseService.SetUserAgent(buildUserAgent())

	service := &BaseService{0, baseService}
	// Set a default HTTP client
	client := core.DefaultHTTPClient()
	client.Timeout = 6 * time.Minute
	service.SetHTTPClient(client)

	return service, nil
}

func (c *BaseService) Clone() *BaseService {
	baseService := c.BaseService.Clone()
	return &BaseService{c.serviceUrlPathSegmentsSize, baseService}
}

func (c *BaseService) Request(req *http.Request, result interface{}) (detailedResponse *core.DetailedResponse, err error) {
	// Extract the operation ID from the request headers.
	var operationId string
	//nolint
	header := req.Header["X-IBMCloud-SDK-Analytics"][0]
	for _, element := range strings.Split(header, ";") {
		if strings.HasPrefix(element, "operation_id") {
			operationId = strings.Split(element, "=")[1]
			break
		}
	}
	if operationId != "" {
		if rulesToApply, ok := rulesByOperation[operationId]; ok {
			requestUrlPathSegments := strings.Split(strings.Trim(req.URL.EscapedPath(), "/"), "/")
			// In the no-path case the result is a slice with an empty string
			// use a nil slice instead in those cases
			if len(requestUrlPathSegments) == 1 && requestUrlPathSegments[0] == "" {
				requestUrlPathSegments = []string{}
			}
			// Check each validation rule that applies to the operation.
			for _, rule := range rulesToApply {
				// Allow for any path segments that might exist in e.g. the URL of a proxy by offseting from the service URL index.
				pathSegmentIndex := c.serviceUrlPathSegmentsSize + rule.pathSegmentIndex
				if len(requestUrlPathSegments) > pathSegmentIndex {
					segmentToValidate := requestUrlPathSegments[pathSegmentIndex]
					if strings.HasPrefix(segmentToValidate, "_") {
						segmentToValidateMessage, err := neturl.PathUnescape(segmentToValidate)
						if err != nil {
							// If we couldn't unescape for some reason, just error with the escaped form
							segmentToValidateMessage = segmentToValidate
						}
						err = fmt.Errorf("%v %v starts with the invalid _ character", rule.errorParameterName, segmentToValidateMessage)
						return nil, core.SDKErrorf(err, "", "invalid-parameter", common.GetComponentInfo())
					}
				}
			}
		}
	}
	return c.BaseService.Request(req, result)
}

func (c *BaseService) SetServiceURL(url string) error {
	err := c.BaseService.SetServiceURL(url)
	if err != nil {
		return err
	}
	// Set CouchDb Session's auth URL to Base service URL
	if c.Options.Authenticator.AuthenticationType() == auth.AUTHTYPE_COUCHDB_SESSION {
		a := c.Options.Authenticator.(*auth.CouchDbSessionAuthenticator)
		a.URL = c.GetServiceURL()
	}
	serviceUrl, err := neturl.ParseRequestURI(c.GetServiceURL())
	if err != nil {
		return nil
	}
	serviceUrlPathSegments := strings.Split(strings.Trim(serviceUrl.EscapedPath(), "/"), "/")
	serviceUrlPathSegmentsSize := len(serviceUrlPathSegments)
	// In the no-path case the result is a slice with an empty string
	// set the size to zero in those cases
	if serviceUrlPathSegmentsSize == 1 && serviceUrlPathSegments[0] == "" {
		c.serviceUrlPathSegmentsSize = 0
	} else {
		c.serviceUrlPathSegmentsSize = serviceUrlPathSegmentsSize
	}
	return nil
}

// SetHTTPClient will set "client" as the http.Client instance to be used
// to invoke individual HTTP requests.
// If automatic retries are currently enabled on "service", then
// "client" will be set as the embedded client instance within
// the retryable client; otherwise "client" will be stored
// directly on "service".
func (c *BaseService) SetHTTPClient(client *http.Client) {
	if client.Transport == nil {
		client.Transport = http.DefaultTransport
	}

	// wrap client's transport into ErrorResponse unless it was already wrapped
	if _, ok := client.Transport.(*ErrorResponse); !ok {
		client.Transport = NewErrorResponse(client.Transport)
	}

	// set cookiejar on if it is missing
	if client.Jar == nil {
		// we can ignore the error, jar.New it is actually always returns nil
		client.Jar, _ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	}
	// Set service's HTTP client as CouchDb Session's client to share cookiejar
	if c.Options.Authenticator.AuthenticationType() == auth.AUTHTYPE_COUCHDB_SESSION {
		a := c.Options.Authenticator.(*auth.CouchDbSessionAuthenticator)
		a.SetClient(client)
	}
	c.BaseService.SetHTTPClient(client)
}

// GetAuthenticatorFromEnvironment instantiates an Authenticator
// using service properties retrieved from external config sources.
func GetAuthenticatorFromEnvironment(credentialKey string) (core.Authenticator, error) {
	props, err := core.GetServiceProperties(credentialKey)
	if err != nil {
		return nil, err
	}
	authType, ok := props[core.PROPNAME_AUTH_TYPE]
	if !ok {
		// this property is not a member of core's constants
		authType, ok = props["AUTHTYPE"]
	}

	if ok && strings.EqualFold(authType, auth.AUTHTYPE_COUCHDB_SESSION) {
		authenticator, err := auth.NewCouchDbSessionAuthenticatorFromMap(props)
		if url, ok := props[core.PROPNAME_SVC_URL]; ok && url != "" {
			authenticator.URL = url
		}
		if disableSSLVerification, ok := props[core.PROPNAME_SVC_DISABLE_SSL]; ok && disableSSLVerification != "" {
			boolValue, err := strconv.ParseBool(disableSSLVerification)
			if err == nil && boolValue {
				authenticator.DisableSSLVerification = true
			}
		}
		return authenticator, err
	}

	return core.GetAuthenticatorFromEnvironment(credentialKey)
}

// buildUserAgent builds the user agent string.
func buildUserAgent() string {
	return fmt.Sprintf("cloudant-go-sdk/%s (%s)", common.Version, getSystemInfo())
}

// getSystemInfo returns the system information.
func getSystemInfo() string {
	return fmt.Sprintf("go.version=%s; os.name=%s os.arch=%s lang=go;",
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)
}

func UnmarshalPrimitiveSpecial(rawInput map[string]json.RawMessage, propertyName string, result *map[string]map[string]int64, containerObjectType string) (err error) {
	if containerObjectType != "*cloudantv1.SearchResult" && containerObjectType != "*cloudantv1.SearchResultProperties" {
		err = fmt.Errorf("UnmarshalPrimitiveSpecial is called with the wrong result type: '%s', it should be '*cloudantv1.SearchResult' or '*cloudantv1.SearchResultProperties'", containerObjectType)
		return
	}
	if propertyName != "counts" && propertyName != "ranges" {
		err = fmt.Errorf("UnmarshalPrimitiveSpecial is called with the wrong propertyName: '%s', it should be 'counts' or 'ranges'", propertyName)
		return
	}

	// Unmarshal rawInput with interim map[string]map[string]float64 type:
	var converted map[string]map[string]float64
	err = core.UnmarshalPrimitive(rawInput, propertyName, &converted)
	if err != nil {
		return
	}
	// Marshal it back as rawMsg:
	rawMsg, err := json.Marshal(converted)
	if err != nil {
		err = fmt.Errorf("error marshalling converted property '%s': %s", propertyName, err.Error())
		return
	}

	// Unmarshal rawMsg with final map[string]map[string]int64 type
	// This should not fail if the given number on the map leaf ends with `.0`:
	err = json.Unmarshal(rawMsg, &result)
	return
}
