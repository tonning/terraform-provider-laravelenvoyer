package models

import "time"

type Hook struct {
	UserId    int       `json:"user_id"`
	ActionId  int       `json:"action_id"`
	Timing    string    `json:"timing"`
	Name      string    `json:"name"`
	RunAs     string    `json:"run_as"`
	Script    string    `json:"script"`
	Sequence  int       `json:"sequence"`
	ProjectId int       `json:"project_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Id        int       `json:"id"`
}

type HookResponse struct {
	Hook Hook `json:"hook"`
}

type HookCreateRequest struct {
	Name     string        `json:"name"`
	Script   string        `json:"script"`
	RunAs    string        `json:"runAs"`
	ActionId int           `json:"actionId"`
	Timing   string        `json:"timing"`
	Servers  []interface{} `json:"servers"`
}

type HookUpdateRequest struct {
	Servers []interface{} `json:"servers"`
}
