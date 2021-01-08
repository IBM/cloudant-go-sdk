/**
 * © Copyright IBM Corporation 2020, 2021. All Rights Reserved.
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
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Session Unit Tests", func() {
	It("Create new Session", func() {
		cookie := &http.Cookie{
			Name:    "AuthSession",
			Value:   makeAuthToken(),
			Expires: time.Now().Add(24 * time.Hour),
		}

		s, err := newSession(cookie)

		Expect(err).To(BeNil())
		Expect(s).ToNot(BeNil())
		Expect(s.getCookie()).To(Equal(cookie))
		Expect(s.expires).To(Equal(cookie.Expires))
		Expect(s.refreshTime).To(BeTemporally("<", s.expires))
		Expect(s.isValid()).To(BeTrue())
	})

	It("Create new Session for non-persistent cookie", func() {
		cookie := &http.Cookie{
			Name:  "AuthSession",
			Value: makeAuthToken(),
		}

		s, err := newSession(cookie)

		Expect(err).To(BeNil())
		Expect(s).ToNot(BeNil())
		Expect(s.getCookie()).To(Equal(cookie))
		Expect(s.expires).ToNot(BeZero())
		Expect(s.refreshTime).To(BeTemporally("<", s.expires))
		Expect(s.isValid()).To(BeTrue())
	})

	It("Test calculation of refresh time", func() {
		cookie := &http.Cookie{
			Name:    "AuthSession",
			Value:   makeAuthToken(),
			Expires: time.Now().Add(24 * time.Hour),
		}

		s, err := newSession(cookie)
		Expect(err).To(BeNil())
		Expect(s).ToNot(BeNil())

		// 20% of 24h is 4h48m
		expected := 4*time.Hour + 48*time.Minute
		roundedExpiration := s.expires.Sub(s.refreshTime).Round(time.Minute)
		Expect(roundedExpiration).To(Equal(expected))
	})

	It("Test Session validation", func() {
		cookie := &http.Cookie{
			Name:    "AuthSession",
			Value:   makeAuthToken(),
			Expires: time.Now().Add(time.Hour),
		}

		s, err := newSession(cookie)
		Expect(err).To(BeNil())
		Expect(s).ToNot(BeNil())
		Expect(s.isValid()).To(BeTrue())

		s.expires = time.Now().Add(-time.Minute)
		Expect(s.isValid()).ToNot(BeTrue())

		s.expires = time.Now().Add(time.Hour)
		s.cookie = nil
		Expect(s.isValid()).ToNot(BeTrue())
	})

	It("Test needsRefresh function", func() {
		cookie := &http.Cookie{
			Name:    "AuthSession",
			Value:   makeAuthToken(),
			Expires: time.Now().Add(time.Hour),
		}

		s, err := newSession(cookie)
		Expect(err).To(BeNil())
		Expect(s).ToNot(BeNil())
		Expect(s.needsRefresh()).ToNot(BeTrue())

		// Move time into the refresh window
		s.refreshTime = time.Now().Add(-10 * time.Minute)
		Expect(s.needsRefresh()).To(BeTrue())

		expected := time.Now().Add(time.Minute).Round(time.Minute)
		Expect(s.refreshTime).To(BeTemporally("~", expected, time.Minute))
	})
})

// makeAuthToken is a helper function that generates mockup AuthSession tokens
func makeAuthToken() string {
	user := "ea36876f-f058-4b1a-897e-2468f89ba5d5-bluemix"
	ts := time.Now().Add(24 * time.Hour).Unix()
	token := []byte(fmt.Sprintf("%s:%X:fakefake", user, ts))
	return base64.StdEncoding.EncodeToString(token)
}
