/**
 * © Copyright IBM Corporation 2020, 2022. All Rights Reserved.
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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`Headers Unit Tests`, func() {
	It("Check SDK Analytics header", func() {
		var headers = GetSdkHeaders("myService", "v123", "myOperation")
		Expect(headers).ToNot(BeNil())
		analyticsHeader, foundIt := headers[headerNameSdkAnalytics]
		Expect(foundIt).To(BeTrue())
		Expect(analyticsHeader).To(Equal("service_name=myService;service_version=v123;operation_id=myOperation"))
	})

	It("GetComponentInfo", func() {
		var problemComponent = GetComponentInfo()
		Expect(problemComponent).ToNot(BeNil())
		Expect(problemComponent.Name).To(Equal("github.com/IBM/cloudant-go-sdk"))
		Expect(problemComponent.Version).ToNot(BeNil())
	})
})
