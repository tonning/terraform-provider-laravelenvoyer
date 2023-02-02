package laravelenvoyer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	apiClient "github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client"
	"github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client/models"
	"log"
	"net/http"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Description: "Laravel Envoyer Environment.",

		CreateContext: resourceEnvironmentCreate,
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentUpdate,
		DeleteContext: resourceEnvironmentDelete,

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
				Description: "Contents of .env",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
			},
			"servers": {
				Description: "Servers",
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectId := d.Get("project_id").(string)
	d.SetId(projectId)

	return resourceEnvironmentUpdate(ctx, d, meta)
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	log.Printf("[INFO] [ENVOYER:resourceEnvironmentRead] start")
	client := meta.(*apiClient.Client)
	projectId := d.Get("project_id").(string)
	opts := apiClient.EnvironmentGetRequest{
		Key: d.Get("key").(string),
	}
	environment, err, response := client.GetEnvironment(projectId, opts)

	if err != nil {
		if response != nil && response.StatusCode == http.StatusNotFound {
			d.SetId("")

			return diag.Diagnostics{}
		}
		resourceEnvironmentDelete(ctx, d, meta)

		return resourceEnvironmentUpdate(ctx, d, meta)
	}

	d.SetId(projectId)
	d.Set("environment", environment.Environment)

	log.Printf("[INFO] [ENVOYER:resourceEnvironmentRead] End")

	return diag.Diagnostics{}
}

func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	log.Printf("[INFO] [ENVOYER:resourceEnvironmentUpdate] Start")
	client := meta.(*apiClient.Client)
	projectId := d.Get("project_id").(string)

	updates := models.EnvironmentUpdateRequest{
		Key:      d.Get("key").(string),
		Contents: d.Get("environment").(string),
		Servers:  d.Get("servers").([]interface{}),
	}

	log.Printf("[INFO] [ENVOYER:resourceEnvironmentUpdate] environment updates: %#v", updates)

	err := client.UpdateEnvironment(projectId, updates)
	if err != nil {
		return err
	}

	d.SetId(projectId)
	//d.Set("servers", d.Get("servers").([]interface{}))
	//d.Set("key", d.Get("key").(string))

	return diag.Diagnostics{}

	//return resourceEnvironmentRead(ctx, d, meta)
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	log.Printf("[INFO] [ENVOYER:resourceEnvironmentDelete] Start")
	c := meta.(*apiClient.Client)

	projectId := d.Get("project_id").(string)

	opts := models.EnvironmentDeleteRequest{
		Key: d.Get("key").(string),
	}

	err, response := c.DeleteEnvironment(projectId, opts)
	if err != nil && response.StatusCode != http.StatusNotFound {
		return err
	}

	d.SetId("")

	return err
}
