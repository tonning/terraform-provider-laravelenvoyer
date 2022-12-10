package client

import "time"

type Server struct {
	Id                      int         `json:"id"`
	ProjectId               int         `json:"project_id"`
	UserId                  int         `json:"user_id"`
	Name                    string      `json:"name"`
	ConnectAs               string      `json:"connect_as"`
	IpAddress               string      `json:"ip_address"`
	Port                    string      `json:"port"`
	PhpSeven                bool        `json:"php_seven"`
	PhpVersion              string      `json:"php_version"`
	Freebsd                 bool        `json:"freebsd"`
	ReceivesCodeDeployments bool        `json:"receives_code_deployments"`
	ShouldRestartFpm        bool        `json:"should_restart_fpm"`
	DeploymentPath          string      `json:"deployment_path"`
	PhpPath                 string      `json:"php_path"`
	ComposerPath            string      `json:"composer_path"`
	PublicKey               string      `json:"public_key"`
	ConnectionStatus        string      `json:"connection_status"`
	CurrentActivity         interface{} `json:"current_activity"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
}

type CreateServerRequest struct {
	Name                    string `json:"name"`
	ConnectAs               string `json:"connectAs"`
	Host                    string `json:"host"`
	Port                    int    `json:"port"`
	PhpVersion              string `json:"phpVersion"`
	ReceivesCodeDeployments bool   `json:"receivesCodeDeployments"`
	DeploymentPath          string `json:"deploymentPath"`
	RestartFpm              bool   `json:"restartFpm"`
	ComposerPath            string `json:"composerPath"`
}

type ServerResponse struct {
	Server Server `json:"server"`
}

type ServerUpdateRequest struct {
	Name string `json:"name"`
}
