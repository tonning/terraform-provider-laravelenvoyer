package laravelenvoyer

import (
	"context"
	apiClient "github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	environment, err, _ := client.GetEnvironment(projectId, opts)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(projectId)
	d.Set("environment", environment.Environment)
	log.Printf("[INFO] [ENVOYER:dataSourceEnvironmentRead] Environment: %s", environment)

	return diag.Diagnostics{}
}
