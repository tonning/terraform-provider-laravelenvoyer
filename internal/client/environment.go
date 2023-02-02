package client

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client/models"
	"log"
	"net/http"
	"strings"
)

type Environment struct {
	Environment string `json:"environment"`
}

type EnvironmentGetRequest struct {
	Key string `json:"key"`
}

func (c *Client) GetEnvironment(projectId string, request EnvironmentGetRequest) (*Environment, error, *http.Response) {
	rb, err := json.Marshal(request)
	if err != nil {
		return nil, err, nil
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/projects/%s/environment", c.HostURL, projectId), strings.NewReader(string(rb)))
	log.Printf("[INFO] [LARAVELENVOYER:GetEnvironment] ProjectId: %s", projectId)
	if err != nil {
		return nil, err, nil
	}

	body, err, response := c.doRequest(req)
	if err != nil {
		return nil, err, response
	}

	environment := Environment{}
	err = json.Unmarshal(body, &environment)
	log.Printf("[INFO] [LARAVELENVOYER:GetEnvironment] Environment: %s", body)
	if err != nil {
		return nil, err, response
	}

	return &environment, nil, response
}

func (c *Client) UpdateEnvironment(projectId string, updates models.EnvironmentUpdateRequest) diag.Diagnostics {
	rb, err := json.Marshal(updates)
	if err != nil {
		return diag.Errorf("Whoops: %s", err)
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/projects/%s/environment", c.HostURL, projectId), strings.NewReader(string(rb)))
	if err != nil {
		return diag.Errorf("Whoops: %s", err)
	}

	err = c.doRequestEmptyBody(req)
	if err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}

func (c *Client) DeleteEnvironment(projectId string, updates models.EnvironmentDeleteRequest) (diag.Diagnostics, *http.Response) {
	rb, err := json.Marshal(updates)
	if err != nil {
		return diag.Errorf("Whoops: %s", err), nil
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/projects/%s/environments", c.HostURL, projectId), strings.NewReader(string(rb)))
	if err != nil {
		return diag.FromErr(err), nil
	}
	body, err, response := c.doRequest(req)
	if err != nil {
		return diag.FromErr(err), response
	}

	if string(body) != "" {
		return diag.Errorf(string(body)), response
	}

	return diag.Diagnostics{}, response
}
