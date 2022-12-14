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

func (c *Client) GetHook(projectId string, hookId string) (*models.Hook, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/projects/%s/hook/%s", c.HostURL, projectId, hookId), nil)
	log.Printf("[INFO] [LARAVELENVOYER:GetHook] ProjectId: %s, HookId: %s", projectId, hookId)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	hook := models.HookResponse{}
	err = json.Unmarshal(body, &hook)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] [LARAVELENVOYER:GetHook] Hook: %#v, Body: %#v", &hook, body)

	return &hook.Hook, nil
}

func (c *Client) CreateHook(projectId string, request *models.HookCreateRequest) (*models.HookResponse, error) {
	log.Printf("[INFO] [LARAVELENVOYER:CreateHook]")
	rb, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/projects/%s/hook", c.HostURL, projectId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	hook := models.HookResponse{}
	err = json.Unmarshal(body, &hook)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] [LARAVELENVOYER:CreateHook] Body: %#v", string(body))
	log.Printf("[INFO] [LARAVELENVOYER:CreateHook] Hook: %#v", hook)

	return &hook, nil
}

func (c *Client) UpdateHook(projectId string, hookId string, hookUpdates models.HookUpdateRequest) (*models.Hook, diag.Diagnostics) {
	rb, err := json.Marshal(hookUpdates)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/projects/%s/hook/%s", c.HostURL, projectId, hookId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	hook := models.HookResponse{}
	err = json.Unmarshal(body, &hook)
	if err != nil {
		return nil, diag.Errorf("Whoops: %s", err)
	}

	return &hook.Hook, nil
}

func (c *Client) DeleteHook(projectId string, hookId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/projects/%s/hook/%s", c.HostURL, projectId, hookId), nil)
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
