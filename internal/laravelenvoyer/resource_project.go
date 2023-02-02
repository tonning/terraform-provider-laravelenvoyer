package laravelenvoyer

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	apiClient "github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client"
	"github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client/models"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Description: "Laravel Envoyer project.",

		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Project ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Project name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"git_provider": {
				Description:  "Git Provider",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"bitbucket", "github", "gitlab", "gitlab-self"}, true),
			},
			"repository": {
				Description: "Git repository",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description:  "Project Type",
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "laravel-5",
				ValidateFunc: validation.StringInSlice([]string{"laravel-5", "laravel-4", "other"}, true),
			},
			"version": {
				Description: "Version",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"user_id": {
				Description: "User ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"token": {
				Description: "Token",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
			"monitor": {
				Description: "Provider",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"retain_deployments": {
				Description: "Number of deployments to retain",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     5,
			},
			"branch": {
				Description: "Git branch.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "master",
			},
			"push_to_deploy": {
				Description: "Pushing to git will cause the project to deploy.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"composer": {
				Description: "Install composer dependencies.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"composer_dev": {
				Description: "Install composer dev dependencies.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"composer_quiet": {
				Description: "Install composer dependencies quietly.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "[ENVOYER:resourceProjectCreate] Start")
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	opts := &models.CreateProjectRequest{
		Name:              d.Get("name").(string),
		Provider:          d.Get("git_provider").(string),
		Repository:        d.Get("repository").(string),
		Type:              d.Get("type").(string),
		RetainDeployments: d.Get("retain_deployments").(int),
		Monitor:           d.Get("monitor").(string),
		Composer:          d.Get("composer").(bool),
		ComposerDev:       d.Get("composer_dev").(bool),
		ComposerQuiet:     d.Get("composer_quiet").(bool),
	}

	tflog.Debug(ctx, fmt.Sprintf("[ENVOYER:resourceProjectCreate] Start 2: Opts: %#v", opts))
	project, err := client.CreateProject(opts)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(project.Project.Id))

	resourceProjectRead(ctx, d, meta)

	return diags
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	projectId := d.Id()

	server, err := client.GetProject(projectId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(server.Id))
	d.Set("name", server.Name)
	d.Set("git_provider", server.Provider)
	d.Set("plain_repository", server.PlainRepository)
	d.Set("repository", server.Repository)
	d.Set("type", server.Type)
	d.Set("version", server.Version)
	d.Set("token", server.Token)
	d.Set("monitor", server.Monitor)
	d.Set("retain_deployments", server.RetainDeployments)
	d.Set("install_dependencies", server.InstallDependencies)
	d.Set("install_dev_dependencies", server.InstallDevDependencies)
	d.Set("quiet_composer", server.QuietComposer)
	d.Set("user_id", server.UserId)

	log.Printf("[INFO] [ENVOYER:resourceProjectRead] End")

	return diag.Diagnostics{}
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	log.Printf("[INFO] [ENVOYER:resourceProjectUpdate] Start")
	client := meta.(*apiClient.Client)
	projectId := d.Id()

	if d.HasChanges("name", "retain_deployments", "monitor", "composer", "composer_dev", "composer_quiet") {
		projectUpdates := models.UpdateProjectRequest{
			Name:              d.Get("name").(string),
			RetainDeployments: d.Get("retain_deployments").(int),
			Monitor:           d.Get("monitor").(string),
			Composer:          d.Get("composer").(bool),
			ComposerDev:       d.Get("composer_dev").(bool),
			ComposerQuiet:     d.Get("composer_quiet").(bool),
		}

		log.Printf("[INFO] [ENVOYER:resourceProjectUpdate] project updates: %#v", projectUpdates)

		_, err := client.UpdateProject(projectId, projectUpdates)
		if err != nil {
			return err
		}
	}

	if d.HasChanges("git_provider", "repository", "branch", "push_to_deploy") {
		projectSourceUpdates := models.UpdateProjectSourceRequest{
			Provider:     d.Get("git_provider").(string),
			Repository:   d.Get("repository").(string),
			Branch:       d.Get("branch").(string),
			PushToDeploy: d.Get("push_to_deploy").(bool),
		}

		log.Printf("[INFO] [ENVOYER:resourceProjectUpdate] project source updates: %#v", projectSourceUpdates)

		_, err := client.UpdateProjectSource(projectId, projectSourceUpdates)
		if err != nil {
			return err
		}
	}

	return resourceProjectRead(ctx, d, meta)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*apiClient.Client)

	projectId := d.Id()

	err := c.DeleteProject(projectId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diag.Diagnostics{}
}
