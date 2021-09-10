package lfwfapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Workflow struct {
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

type StartWFResponse struct {
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

func (c *client) Send(r *http.Request, v interface{}) error {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")

	res, err := c.HttpClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("Error encountered on %v. Got status code %v", r.URL, res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(v)
}

func (c *client) GetAllWorkflows(ctx context.Context) ([]Workflow, error) {
	r, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v/workflow/api/workflow", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	w := []Workflow{}

	err = c.Send(r, &w)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (c *client) StartWorkflow(ctx context.Context, wfname string, p []Parameter) (StartWFResponse, error) {
	s := StartWFResponse{}

	b, err := json.Marshal(startWorkflowRequest{p})
	if err != nil {
		return s, err
	}

	r, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%v/Workflow/api/instances/%v", c.baseURL, url.PathEscape(wfname)), bytes.NewBuffer(b))
	if err != nil {
		return s, err
	}

	err = c.Send(r, &s)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (c *client) GetWorkflowParameters(ctx context.Context, wfname string) ([]workflowParameters, error) {
	r, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%v/Workflow/api/workflow/parameters/%v", c.baseURL, url.PathEscape(wfname)), nil)
	if err != nil {
		return nil, err
	}

	w := []workflowParameters{}

	err = c.Send(r, &w)
	if err != nil {
		return nil, err
	}

	return w, nil
}
