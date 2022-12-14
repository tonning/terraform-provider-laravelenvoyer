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

type Project struct {
	Id                     int           `json:"id"`
	UserId                 int           `json:"user_id"`
	Version                int           `json:"version"`
	Name                   string        `json:"name"`
	Provider               string        `json:"provider"`
	PlainRepository        string        `json:"plain_repository"`
	Repository             string        `json:"repository"`
	Type                   string        `json:"type"`
	Branch                 string        `json:"branch"`
	PushToDeploy           bool          `json:"push_to_deploy"`
	WebhookId              interface{}   `json:"webhook_id"`
	Status                 interface{}   `json:"status"`
	ShouldDeployAgain      int           `json:"should_deploy_again"`
	DeploymentStartedAt    interface{}   `json:"deployment_started_at"`
	DeploymentFinishedAt   time.Time     `json:"deployment_finished_at"`
	LastDeploymentStatus   string        `json:"last_deployment_status"`
	DailyDeploys           int           `json:"daily_deploys"`
	WeeklyDeploys          int           `json:"weekly_deploys"`
	LastDeploymentTook     int           `json:"last_deployment_took"`
	RetainDeployments      int           `json:"retain_deployments"`
	EnvironmentServers     []int         `json:"environment_servers"`
	Folders                []interface{} `json:"folders"`
	Monitor                string        `json:"monitor"`
	NewYorkStatus          string        `json:"new_york_status"`
	LondonStatus           string        `json:"london_status"`
	SingaporeStatus        string        `json:"singapore_status"`
	Token                  string        `json:"token"`
	CreatedAt              time.Time     `json:"created_at"`
	UpdatedAt              time.Time     `json:"updated_at"`
	InstallDevDependencies bool          `json:"install_dev_dependencies"`
	InstallDependencies    bool          `json:"install_dependencies"`
	QuietComposer          bool          `json:"quiet_composer"`
	Servers                []struct {
	} `json:"servers"`
	HasEnvironment          bool   `json:"has_environment"`
	HasMonitoringError      bool   `json:"has_monitoring_error"`
	HasMissingHeartbeats    bool   `json:"has_missing_heartbeats"`
	LastDeployedBranch      string `json:"last_deployed_branch"`
	LastDeploymentId        int    `json:"last_deployment_id"`
	LastDeploymentAuthor    string `json:"last_deployment_author"`
	LastDeploymentAvatar    string `json:"last_deployment_avatar"`
	LastDeploymentHash      string `json:"last_deployment_hash"`
	LastDeploymentTimestamp string `json:"last_deployment_timestamp"`
}

type ProjectResponse struct {
	Project Project `json:"project"`
}

type CreateProjectRequest struct {
	Name              string `json:"name"`
	Provider          string `json:"provider"`
	Repository        string `json:"repository"`
	Type              string `json:"type"`
	RetainDeployments int    `json:"retain_deployments"`
	Monitor           string `json:"monitor"`
	Composer          bool   `json:"composer"`
	ComposerDev       bool   `json:"composer_dev"`
	ComposerQuiet     bool   `json:"composer_quiet"`
}

type UpdateProjectRequest struct {
	Name              string `json:"name"`
	RetainDeployments int    `json:"retain_deployments"`
	Monitor           string `json:"monitor"`
	Composer          bool   `json:"composer"`
	ComposerDev       bool   `json:"composer_dev"`
	ComposerQuiet     bool   `json:"composer_quiet"`
}

type UpdateProjectTypeRequeset struct {
	Provider     string `json:"provider"`
	Repository   string `json:"repository"`
	Branch       string `json:"branch"`
	PushToDeploy bool   `json:"push_to_deploy"`
}
