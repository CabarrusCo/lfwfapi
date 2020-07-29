package lfwfapi

import (
	"log"
	"testing"
)

func TestLFWorkflow(t *testing.T) {

	wfClient := NewClient("http://WORKFLOWURLBASEHERE")

	getAllWorkflows, err := wfClient.GetAllWorkflows()
	if err != nil {
		t.Errorf("Error encountered in Get All Workflows %s", err)
	} else {
		log.Printf("All workflows %+v", getAllWorkflows)
	}

	getWorkflowParameters, err := wfClient.GetWorkflowParameters("PayIT API Alert Blacklist")
	if err != nil {
		t.Errorf("Error encountered in Get Workflow Parameters %s", err)
	} else {
		log.Printf("All Workflow Parameters %+v", getWorkflowParameters)
	}

	runWorkflow, err := wfClient.StartWorkflow("Send Email From Web Service", nil)
	if err != nil {
		t.Errorf("Error encounter in Run Workflow %s", err)
	} else {
		log.Printf("%s", runWorkflow)
	}

}
