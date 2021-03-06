package workflows

/*
	Rather than creating a separate workflows_test package, this suite is part of
	the workflows package itself in order to provide regression tests for the ProvisionNodeTasks
	vm task templates without exporting these variables for testing. Please be careful
	as this suite will have access to all unexported functions and variables in the workflows
	package. You have been warned

	- The ghost in the air ducts
*/

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/nu7hatch/gouuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rackhd/rackhd-cpi/config"
	"github.com/rackhd/rackhd-cpi/helpers"
)

var _ = Describe("DeprovisionNodeWorkflow", func() {
	Describe("DeprovisionNodeWorkflow", func() {
		It("have the expected structure", func() {
			vendoredWorkflow := deprovisionNodeWorkflow{}
			err := json.Unmarshal(deprovisionNodeWorkflowTemplate, &vendoredWorkflow)
			Expect(err).ToNot(HaveOccurred())

			deprovisionNodeWorkflowFile, err := os.Open("../templates/deprovision_node_workflow.json")
			Expect(err).ToNot(HaveOccurred())
			defer deprovisionNodeWorkflowFile.Close()

			b, err := ioutil.ReadAll(deprovisionNodeWorkflowFile)
			Expect(err).ToNot(HaveOccurred())

			expectedWorkflow := deprovisionNodeWorkflow{}
			err = json.Unmarshal(b, &expectedWorkflow)
			Expect(err).ToNot(HaveOccurred())

			Expect(vendoredWorkflow).To(Equal(expectedWorkflow))
		})

		It("marshalls into the expected JSON document", func() {
			vendoredWorkflow := deprovisionNodeWorkflow{}
			err := json.Unmarshal(deprovisionNodeWorkflowTemplate, &vendoredWorkflow)
			Expect(err).ToNot(HaveOccurred())

			vendoredWorkflowJSON, err := json.Marshal(vendoredWorkflow)
			Expect(err).ToNot(HaveOccurred())

			deprovisionNodeWorkflowFile, err := os.Open("../templates/deprovision_node_workflow.json")
			Expect(err).ToNot(HaveOccurred())
			defer deprovisionNodeWorkflowFile.Close()

			expectedWorkflowJSON, err := ioutil.ReadAll(deprovisionNodeWorkflowFile)
			Expect(err).ToNot(HaveOccurred())

			Expect(vendoredWorkflowJSON).To(MatchJSON(expectedWorkflowJSON))
		})
	})

	Describe("generating the set of deprovision workflow tasks and workflow", func() {
		It("publishes the tasks and workflow", func() {
			u, err := uuid.NewV4()
			Expect(err).ToNot(HaveOccurred())
			uID := u.String()

			apiServer, err := helpers.GetRackHDHost()
			Expect(err).ToNot(HaveOccurred())

			c := config.Cpi{ApiServer: apiServer, RequestID: uID}

			workflowName, err := PublishDeprovisionNodeWorkflow(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(workflowName).To(ContainSubstring(uID))
		})
	})
})
