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

func (c *Client) GetServer(projectId string, serverId string) (*models.Server, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/projects/%s/servers/%s", c.HostURL, projectId, serverId), nil)
	log.Printf("[INFO] [LARAVELENVOYER:GetServer] ProjectId: %s, ServerId: %s", projectId, serverId)
	if err != nil {
		return nil, err
	}

	body, err, _ := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	server := models.ServerResponse{}
	err = json.Unmarshal(body, &server)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] [LARAVELENVOYER:GetServer] Server: %#v, Body: %#v", &server, body)

	return &server.Server, nil
}

func (c *Client) CreateServer(projectId string, request *models.CreateServerRequest) (*models.ServerResponse, error) {
	log.Printf("[INFO] [LARAVELENVOYER:CreateServer]")
	rb, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/projects/%s/servers", c.HostURL, projectId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err, _ := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	server := models.ServerResponse{}
	err = json.Unmarshal(body, &server)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] [LARAVELENVOYER:CreateServer] Body: %#v", string(body))
	log.Printf("[INFO] [LARAVELENVOYER:CreateServer] Server: %#v", server)

	return &server, nil
}

func (c *Client) UpdateServer(projectId string, serverId string, serverUpdates models.ServerUpdateRequest) (*models.Server, diag.Diagnostics) {
	rb, err := json.Marshal(serverUpdates)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/projects/%s/servers/%s", c.HostURL, projectId, serverId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	body, err, _ := c.doRequest(req)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	server := models.ServerResponse{}
	err = json.Unmarshal(body, &server)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	return &server.Server, nil
}

func (c *Client) DeleteServer(projectId string, serverId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/projects/%s/servers/%s", c.HostURL, projectId, serverId), nil)
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
