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

package common

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`Headers Unit Tests`, func() {
	It("Successfully load SystemInfo", func() {
		sysinfo := GetSystemInfo()
		Expect(sysinfo).ToNot(BeNil())
		Expect(strings.Contains(sysinfo, "lang=")).To(BeTrue())
		Expect(strings.Contains(sysinfo, "arch=")).To(BeTrue())
		Expect(strings.Contains(sysinfo, "os=")).To(BeTrue())
		Expect(strings.Contains(sysinfo, "go.version=")).To(BeTrue())
	})

	It("Check SDK User Agent header", func() {
		var headers = GetSdkHeaders("myService", "v123", "myOperation")
		Expect(headers).ToNot(BeNil())

		actUserAgentHeader, foundIt := headers[headerNameUserAgent]
		Expect(foundIt).To(BeTrue())
		expUserAgentHeader := sdkName + "/" + Version + " " + GetSystemInfo()
		Expect(actUserAgentHeader).To(Equal(expUserAgentHeader))
	})

	It("Check SDK Analytics header", func() {
		var headers = GetSdkHeaders("myService", "v123", "myOperation")
		Expect(headers).ToNot(BeNil())
		analyticsHeader, foundIt := headers[headerNameSdkAnalytics]
		Expect(foundIt).To(BeTrue())
		Expect(analyticsHeader).To(Equal("service_name=myService;service_version=v123;operation_id=myOperation"))
	})
})
