package lfwfapi

import (
	"context"
	"log"
	"testing"
)

func TestLFWorkflow(t *testing.T) {

	wfClient := NewClient("http://WORKFLOWURLBASEHERE")

	getAllWorkflows, err := wfClient.GetAllWorkflows(context.Background())
	if err != nil {
		t.Errorf("Error encountered in Get All Workflows %s", err)
	} else {
		log.Printf("All workflows %+v", getAllWorkflows)
	}

	getWorkflowParameters, err := wfClient.GetWorkflowParameters(context.Background(), "Workflow Name Here")
	if err != nil {
		t.Errorf("Error encountered in Get Workflow Parameters %s", err)
	} else {
		log.Printf("All Workflow Parameters %+v", getWorkflowParameters)
	}

	runWorkflow, err := wfClient.StartWorkflow(context.Background(), "Workflow Name Here", nil)
	if err != nil {
		t.Errorf("Error encounter in Run Workflow %s", err)
	} else {
		log.Printf("%s", runWorkflow)
	}

}
