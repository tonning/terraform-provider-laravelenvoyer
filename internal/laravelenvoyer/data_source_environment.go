package laravelenvoyer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	apiClient "github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client"
	"github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client/models"
	"log"
)

func dataSourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Description: "Get a project's environment (.env).",

		ReadContext: dataSourceEnvironmentRead,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Description: "Project ID.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"key": {
				Description: "Key to unlock environment",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
			},
			"environment": {
				Description: "Project environment",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
			"reset_if_no_key": {
				Type:        schema.TypeBool,
				Description: "This will reset the environment if no key has been set for the project.",
				Optional:    true,
				Default:     true,
			},
		},
	}
}

func dataSourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	log.Printf("[INFO] [ENVOYER:dataSourceEnvironmentRead] start")
	client := meta.(*apiClient.Client)
	projectId := d.Get("project_id").(string)
	opts := apiClient.EnvironmentGetRequest{
		Key: d.Get("key").(string),
	}
	environment, err, response := client.GetEnvironment(projectId, opts)

	if err != nil && response.StatusCode != 422 {
		return diag.FromErr(err)
	}

	if response.StatusCode == 422 && d.Get("reset_if_no_key").(bool) == true {
		client.DeleteEnvironment(projectId, models.EnvironmentDeleteRequest{
			Key: d.Get("key").(string),
		})
	}

	environment, err, _ = client.GetEnvironment(projectId, opts)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(projectId)
	d.Set("environment", environment.Environment)
	log.Printf("[INFO] [ENVOYER:dataSourceEnvironmentRead] Environment: %s", environment)

	return diag.Diagnostics{}
}
