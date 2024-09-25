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

type errorRoundTripper struct {
	rt http.RoundTripper
}

func NewErrorRoundTripper(rt http.RoundTripper) errorRoundTripper {
	return errorRoundTripper{rt: rt}
}

// drainBody reads body into memory, returns two identical ReadClosers
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

// transformError reads the response body, parses it as json, augments it,
// encodes back into json and sets as a new response
func transformError(resp *http.Response) error {
	var err error
	save := resp.Body
	savecl := resp.ContentLength

	save, resp.Body, err = drainBody(resp.Body)
	if err != nil {
		return err
	}

	respError := make(map[string]interface{})
	// decode
	if err := json.NewDecoder(resp.Body).Decode(&respError); err == nil {
		// augment
		trace := resp.Header.Get("X-Couch-Request-Id")
		if trace != "" {
			respError["trace"] = trace
		}

		err := make(map[string]string)
		if m, ok := respError["error"]; ok {
			err["code"] = m.(string)
			err["message"] = m.(string)
		}
		if m, ok := respError["reason"]; ok {
			err["message"] += ": " + m.(string)
		}
		if len(err) > 0 {
			respError["errors"] = []map[string]string{err}
		}

		// encode back
		var newErrorJson bytes.Buffer
		if err := json.NewEncoder(&newErrorJson).Encode(respError); err == nil {
			save = io.NopCloser(&newErrorJson)
			savecl = int64(len(newErrorJson.Bytes()))
		}
	}

	resp.Body = save
	resp.ContentLength = savecl

	return nil
}

func (ert errorRoundTripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
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

	return ert.rt.RoundTrip(req)
}
