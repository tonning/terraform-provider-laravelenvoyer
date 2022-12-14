package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	"net/http"
	"strings"
)

func (c *Client) GetProject(projectId string) (*Project, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/projects/%s", c.HostURL, projectId), nil)
	log.Printf("[INFO] [LARAVELENVOYER:GetProject] ProjectId: %s, ProjectId: %s", projectId)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	project := ProjectResponse{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] [LARAVELENVOYER:GetProject] Project: %#v, Body: %#v", &project, body)

	return &project.Project, nil
}

func (c *Client) CreateProject(request *CreateProjectRequest) (*ProjectResponse, error) {
	log.Printf("[INFO] [LARAVELENVOYER:CreateProject]")
	rb, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/projects", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	project := ProjectResponse{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] [LARAVELENVOYER:CreateProject] Body: %#v", string(body))
	log.Printf("[INFO] [LARAVELENVOYER:CreateProject] Project: %#v", project)

	return &project, nil
}

func (c *Client) UpdateProject(projectId string, projectUpdates UpdateProjectRequest) (*Project, diag.Diagnostics) {
	rb, err := json.Marshal(projectUpdates)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/projects/%s", c.HostURL, projectId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	project := ProjectResponse{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	return &project.Project, nil
}

func (c *Client) DeleteProject(projectId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/projects/%s", c.HostURL, projectId), nil)
	if err != nil {
		return err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}
