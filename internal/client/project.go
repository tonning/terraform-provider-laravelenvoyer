package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client/models"
	"log"
	"net/http"
	"strings"
)

func (c *Client) GetProject(projectId string) (*models.Project, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/projects/%s", c.HostURL, projectId), nil)
	log.Printf("[INFO] [LARAVELENVOYER:GetProject] ProjectId: %s, ProjectId: %s", projectId)
	if err != nil {
		return nil, err
	}

	body, err, _ := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	project := models.ProjectResponse{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] [LARAVELENVOYER:GetProject] Project: %#v, Body: %#v", &project, body)

	return &project.Project, nil
}

func (c *Client) CreateProject(request *models.CreateProjectRequest) (*models.ProjectResponse, error) {
	log.Printf("[INFO] [LARAVELENVOYER:CreateProject]")
	rb, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/projects", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err, _ := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	project := models.ProjectResponse{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] [LARAVELENVOYER:CreateProject] Body: %#v", string(body))
	log.Printf("[INFO] [LARAVELENVOYER:CreateProject] Project: %#v", project)

	return &project, nil
}

func (c *Client) UpdateProject(projectId string, projectUpdates models.UpdateProjectRequest) (*models.Project, diag.Diagnostics) {
	rb, err := json.Marshal(projectUpdates)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/projects/%s", c.HostURL, projectId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	body, err, _ := c.doRequest(req)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	project := models.ProjectResponse{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	return &project.Project, nil
}

func (c *Client) UpdateProjectSource(projectId string, projectSourceUpdates models.UpdateProjectSourceRequest) (*models.Project, diag.Diagnostics) {
	rb, err := json.Marshal(projectSourceUpdates)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/projects/%s/source", c.HostURL, projectId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	body, err, _ := c.doRequest(req)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	project := models.ProjectResponse{}
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
	body, err, _ := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}
