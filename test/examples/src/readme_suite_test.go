/**
 * © Copyright IBM Corporation 2022, 2023. All Rights Reserved.
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

package main

import (
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

func TestReadme(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "README Suite")
}

var _ = BeforeSuite(func() {
	wiremockUrl := os.Getenv("WIREMOCK_URL")

	// reset wiremock
	_, err := http.Post(wiremockUrl+"/__admin/scenarios/reset", "application/json", nil)
	Expect(err).ShouldNot(HaveOccurred())

	// set authentication
	os.Setenv("CLOUDANT_AUTH_TYPE", "basic")
	os.Setenv("CLOUDANT_USERNAME", "admin")
	os.Setenv("CLOUDANT_PASSWORD", "pass")
	os.Setenv("CLOUDANT_URL", wiremockUrl)
})

var _ = Describe(`Readme integration tests`, func() {
	var runFile, outputFile string

	// AfterEach is an actual assertion step and each It sections acts as a setup
	// this is recommended ginkgo v1 approach to the table tests
	AfterEach(func() {
		cmd := exec.Command("go", "run", runFile)
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).ShouldNot(HaveOccurred())

		// give it a solid timeout for possible CI resource constrains
		session.Wait(5 * time.Second)
		Eventually(session).Should(gexec.Exit())

		expect, err := os.ReadFile(outputFile)
		Expect(err).ShouldNot(HaveOccurred())

		Eventually(session).Should(gbytes.Say(string(expect)))
	})

	It(`Creates db and doc for the first time`, func() {
		runFile = "create_db_and_doc/create_db_and_doc.go"
		outputFile = "../output/create_db_and_doc.txt"
	})

	It(`Gets document from orders database`, func() {
		runFile = "get_info_from_existing_database/get_info_from_existing_database.go"
		outputFile = "../output/get_info_from_existing_database.txt"
	})

	It(`Updates doc for the first time`, func() {
		runFile = "update_doc/update_doc.go"
		outputFile = "../output/update_doc.txt"
	})

	It(`Updates doc for the second time`, func() {
		runFile = "update_doc/update_doc.go"
		outputFile = "../output/update_doc2.txt"
	})

	It(`Deletes existing doc`, func() {
		runFile = "delete_doc/delete_doc.go"
		outputFile = "../output/delete_doc.txt"
	})

	It(`Deletes non-existing doc`, func() {
		runFile = "delete_doc/delete_doc.go"
		outputFile = "../output/delete_doc2.txt"
	})
})
