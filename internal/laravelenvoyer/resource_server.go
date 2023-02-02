package laravelenvoyer

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	apiClient "github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client"
	"github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client/models"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform laravelenvoyer scaffolding.",

		CreateContext: resourceServerCreate,
		ReadContext:   resourceServerRead,
		UpdateContext: resourceServerUpdate,
		DeleteContext: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Description: "Project ID.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"user_id": {
				Description: "User ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Server name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"host": {
				Description: "Host / IP address.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"port": {
				Description: "Port.",
				Type:        schema.TypeInt,
				Required:    false,
				Optional:    true,
				Default:     22,
			},
			"php_version": {
				Description: "PHP version.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"connect_as": {
				Description: "Connect as user.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"receives_code_deployments": {
				Description: "Receives code deployments.",
				Type:        schema.TypeBool,
				Required:    false,
				Optional:    true,
				Default:     true,
			},
			"should_restart_fpm": {
				Description: "Restart PHP after deployment.",
				Type:        schema.TypeBool,
				Required:    false,
				Optional:    true,
				Default:     true,
			},
			"deployment_path": {
				Description: "Path to where project lives on the server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"connection_status": {
				Description: "Connection status.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"public_key": {
				Description: "Public key.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"composer_path": {
				Description: "Path to composer",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "composer",
			},
		},
	}
}

func resourceServerCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//log.Printf("[INFO] [ENVOYER:resourceServerCreate] Start")
	tflog.Debug(ctx, "[ENVOYER:resourceServerCreate] Start")
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	projectId := strconv.Itoa(d.Get("project_id").(int))

	opts := &models.CreateServerRequest{
		Name:                    d.Get("name").(string),
		ConnectAs:               d.Get("connect_as").(string),
		Host:                    d.Get("host").(string),
		Port:                    d.Get("port").(int),
		PhpVersion:              d.Get("php_version").(string),
		ReceivesCodeDeployments: d.Get("receives_code_deployments").(bool),
		DeploymentPath:          d.Get("deployment_path").(string),
		RestartFpm:              d.Get("should_restart_fpm").(bool),
	}

	tflog.Debug(ctx, fmt.Sprintf("[ENVOYER:resourceServerCreate] Start 2: Opts: %#v", opts))
	server, err := client.CreateServer(projectId, opts)
	tflog.Debug(ctx, "[ENVOYER:resourceServerCreate] Start 3")

	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "[ENVOYER:resourceServerCreate] Start 4")
	d.SetId(strconv.Itoa(server.Server.Id))
	tflog.Debug(ctx, "[ENVOYER:resourceServerCreate] Start 5")

	resourceServerRead(ctx, d, meta)
	tflog.Debug(ctx, "[ENVOYER:resourceServerCreate] Start 6")

	return diags
}

func resourceServerRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the laravelenvoyer configure method
	client := meta.(*apiClient.Client)

	projectId := strconv.Itoa(d.Get("project_id").(int))
	serverId := d.Id()

	server, err := client.GetServer(projectId, serverId)
	if err != nil {
		d.SetId("")
		return diag.Diagnostics{}
	}

	d.SetId(strconv.Itoa(server.Id))
	log.Printf("[INFO] [ENVOYER:resourceServerRead] Server: %#v", server)

	d.Set("user_id", server.UserId)
	d.Set("name", server.Name)
	d.Set("host", server.IpAddress)
	d.Set("port", server.Port)
	d.Set("php_version", server.PhpVersion)
	d.Set("connect_as", server.ConnectAs)
	d.Set("receives_code_deployments", server.ReceivesCodeDeployments)
	d.Set("should_restart_fpm", server.ShouldRestartFpm)
	d.Set("deployment_path", server.DeploymentPath)
	d.Set("public_key", server.PublicKey)
	d.Set("connection_status", server.ConnectionStatus)

	log.Printf("[INFO] [ENVOYER:resourceServerRead] End")

	return diag.Diagnostics{}
}

func resourceServerUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	log.Printf("[INFO] [ENVOYER:resourceServerUpdate] Start")
	client := meta.(*apiClient.Client)
	projectId := strconv.Itoa(d.Get("project_id").(int))
	serverId := d.Id()

	serverUpdates := models.ServerUpdateRequest{
		Name:                    d.Get("name").(string),
		Host:                    d.Get("host").(string),
		Port:                    d.Get("port").(int),
		PhpVersion:              d.Get("php_version").(string),
		ConnectAs:               d.Get("connect_as").(string),
		ReceivesCodeDeployments: d.Get("receives_code_deployments").(bool),
		RestartFpm:              d.Get("should_restart_fpm").(bool),
		DeploymentPath:          d.Get("deployment_path").(string),
		ComposerPath:            d.Get("composer_path").(string),
	}

	log.Printf("[INFO] [ENVOYER:resourceServerUpdate] server updates: %#v", serverUpdates)

	_, err := client.UpdateServer(projectId, serverId, serverUpdates)
	if err != nil {
		return err
	}

	return resourceServerRead(ctx, d, meta)
}

func resourceServerDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the laravelenvoyer configure method
	// client := meta.(*apiClient)
	c := meta.(*apiClient.Client)

	projectId := strconv.Itoa(d.Get("project_id").(int))
	serverId := d.Id()

	err := c.DeleteServer(projectId, serverId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diag.Diagnostics{}
}
