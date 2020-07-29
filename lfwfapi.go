package lfwfapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type workflow struct {
	BpViewURL           string `json:"bpViewUrl"`
	Description         string `json:"description"`
	Flags               int    `json:"flags"`
	FullScreenBPViewURL string `json:"fullScreenBPViewUrl"`
	HasActiveInstances  bool   `json:"hasActiveInstances"`
	HasActiveRules      bool   `json:"hasActiveRules"`
	LastUpdated         string `json:"lastUpdated"`
	Name                string `json:"name"`
	Version             int    `json:"version"`
	WorkflowID          int    `json:"workflowId"`
	WorkflowViewURL     string `json:"workflowViewUrl"`
}

type Parameter struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type startWorkflowRequest struct {
	ParameterCollection []Parameter `json:"ParameterCollection"`
}

type startWorkflowResponse struct {
	Fault      fault  `json:"fault"`
	InstanceID string `json:"instanceId"`
}

type fault struct {
	Status     int    `json:"Status"`
	Detail     string `json:"Detail"`
	DetailCode int    `json:"DetailCode"`
}

type workflowParameters struct {
	DefaultValues []string `json:"defaultValues"`
	IsMultivalued bool     `json:"isMultivalued"`
	SourceNames   []string `json:"sourceNames"`
	Style         string   `json:"style"`
	TokenName     string   `json:"tokenName"`
	TokenTags     string   `json:"tokenTags"`
}

type client struct {
	baseURL    string
	HttpClient *http.Client
}

func NewClient(url string) *client {
	return &client{
		baseURL:    url,
		HttpClient: &http.Client{},
	}
}

func setReqInfo(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-Type", "application/json")
}

func (c client) GetAllWorkflows() ([]workflow, error) {
	url := fmt.Sprintf("%s/Workflow/api/workflow", c.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	setReqInfo(req)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Expected Status Code 200, Got Status Code %v instead", res.StatusCode)
	}

	workflows := []workflow{}
	err = json.NewDecoder(res.Body).Decode(&workflows)
	if err != nil {
		return nil, err
	}
	return workflows, nil
}

func (c client) StartWorkflow(workflowName string, p []Parameter) (string, error) {
	workflowNameEncoded := url.PathEscape(workflowName)
	url := fmt.Sprintf("%s/Workflow/api/instances/%s", c.baseURL, workflowNameEncoded)

	body, err := json.Marshal(startWorkflowRequest{p})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	setReqInfo(req)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("Expected Status Code 200, Got Status Code %v instead", res.StatusCode)
	}

	result := startWorkflowResponse{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if result.Fault.Status != 0 {
		return "", fmt.Errorf("Problem Starting Workflow: %s", result.Fault.Detail)
	}
	return fmt.Sprintf("Workflow Started Successfully with Instance Id: %s", result.InstanceID), nil
}

func (c client) GetWorkflowParameters(workflowName string) ([]workflowParameters, error) {
	workflowNameEncoded := url.PathEscape(workflowName)
	url := fmt.Sprintf("%s/Workflow/api/workflow/parameters/%s", c.baseURL, workflowNameEncoded)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	setReqInfo(req)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Expected Status Code 200, Got Status Code %v instead", res.StatusCode)
	}

	wfParams := []workflowParameters{}
	err = json.NewDecoder(res.Body).Decode(&wfParams)
	if err != nil {
		return nil, err
	}
	return wfParams, nil
}
