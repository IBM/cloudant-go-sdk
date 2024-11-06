/**
 * © Copyright IBM Corporation 2024. All Rights Reserved.
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
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// ErrorResponse is an error augmentation RoundTripper
// that in case of http error response parses the return body
// and converts it to a closer match to the standard error including adding a trace ID
// and appending the CouchDB/Cloudant error reason to the message
type ErrorResponse struct {
	next http.RoundTripper
}

// NewErrorResponse creates a new ErrorResponse middleware
func NewErrorResponse(rt http.RoundTripper) http.RoundTripper {
	return &ErrorResponse{next: rt}
}

// RoundTrip implements RoundTripper interface
func (er *ErrorResponse) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	defer func() {
		if err == nil && resp.StatusCode >= http.StatusBadRequest {
			ct := resp.Header.Get("Content-Type")
			if strings.HasPrefix(ct, "application/json") {
				err = transformError(resp)
				if err != nil {
					resp = nil
				}
			}
		}
	}()

	return er.next.RoundTrip(req)
}

// drainBody reads body into memory and returns two identical ReadClosers
// one of which is safe to read
func drainBody(body io.ReadCloser) (orig, save io.ReadCloser, err error) {
	if body == nil || body == http.NoBody {
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(body); err != nil {
		return nil, body, err
	}
	if err = body.Close(); err != nil {
		return nil, body, err
	}
	r := bytes.NewReader(buf.Bytes())
	return io.NopCloser(&buf), io.NopCloser(r), nil
}

// transformError reads the response's body, parses it as json, augments it,
// encodes back and sets as a new response
func transformError(resp *http.Response) error {
	var err error
	var save io.ReadCloser
	savecl := resp.ContentLength

	save, resp.Body, err = drainBody(resp.Body)
	if err != nil {
		return err
	}

	respError := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&respError)
	if err != nil {
		// since resp is not json we just return it as it is
		resp.Body = save
		return nil
	}

	if _, ok := respError["trace"]; ok {
		resp.Body = save
		return nil
	}

	respErrorWasAugmented := false
	if _, ok := respError["errors"]; !ok {
		err := make(map[string]string)
		if m, ok := respError["error"]; ok {
			err["code"] = m.(string)
			err["message"] = m.(string)
			if m, ok := respError["reason"]; ok && m != "" {
				err["message"] += ": " + m.(string)
			}
			respError["errors"] = []map[string]string{err}
			respErrorWasAugmented = true
		}
	}

	trace := resp.Header.Get("X-Request-Id")
	if trace == "" {
		trace = resp.Header.Get("X-Couch-Request-Id")
	}
	if _, ok := respError["errors"]; ok && trace != "" {
		respError["trace"] = trace
		respErrorWasAugmented = true
	}

	if respErrorWasAugmented {
		var newErrorJson bytes.Buffer
		if err := json.NewEncoder(&newErrorJson).Encode(respError); err == nil {
			save = io.NopCloser(&newErrorJson)
			if resp.Header.Get("Transfer-Encoding") != "chunked" {
				savecl = int64(len(newErrorJson.Bytes()))
			}
		}
	}

	resp.Body = save
	resp.ContentLength = savecl

	return nil
}
