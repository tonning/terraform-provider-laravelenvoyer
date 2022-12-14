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

func resourceHook() *schema.Resource {
	return &schema.Resource{
		Description: "Laravel Envoyer Hook.",

		CreateContext: resourceHookCreate,
		ReadContext:   resourceHookRead,
		UpdateContext: resourceHookUpdate,
		DeleteContext: resourceHookDelete,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Description: "Project ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Hook name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"script": {
				Description: "Script",
				Type:        schema.TypeString,
				Required:    true,
			},
			"run_as": {
				Description: "Run as",
				Type:        schema.TypeString,
				Required:    true,
			},
			"action_id": {
				Description:  "Action ID",
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      2,
				ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4}),
			},
			"timing": {
				Description:  "Timing",
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "after",
				ValidateFunc: validation.StringInSlice([]string{"before", "after"}, true),
			},
			"servers": {
				Description: "Servers",
				Type:        schema.TypeMap,
				Required:    true,
			},
		},
	}
}

func resourceHookCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "[ENVOYER:resourceHookCreate] Start")
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	projectId := strconv.Itoa(d.Get("project_id").(int))

	opts := &models.HookCreateRequest{
		Name:     d.Get("name").(string),
		Script:   d.Get("script").(string),
		RunAs:    d.Get("run_as").(string),
		ActionId: d.Get("action_id").(int),
		Timing:   d.Get("timing").(string),
		Servers:  d.Get("servers").([]int),
	}

	tflog.Debug(ctx, fmt.Sprintf("[ENVOYER:resourceHookCreate] Start 2: Opts: %#v", opts))
	hook, err := client.CreateHook(projectId, opts)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(hook.Hook.Id))

	resourceHookRead(ctx, d, meta)

	return diags
}

func resourceHookRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	hookId := d.Id()
	projectId := strconv.Itoa(d.Get("project_id").(int))

	hook, err := client.GetHook(projectId, hookId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(hook.Id))
	d.Set("user_id", hook.UserId)
	d.Set("action_id", hook.ActionId)
	d.Set("timing", hook.Timing)
	d.Set("name", hook.Name)
	d.Set("run_as", hook.RunAs)
	d.Set("script", hook.Script)
	d.Set("sequence", hook.Sequence)
	d.Set("project_id", hook.ProjectId)
	d.Set("updated_at", hook.UpdatedAt)
	d.Set("created_at", hook.CreatedAt)

	log.Printf("[INFO] [ENVOYER:resourceHookRead] End")

	return diag.Diagnostics{}
}

func resourceHookUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	log.Printf("[INFO] [ENVOYER:resourceHookUpdate] Start")
	client := meta.(*apiClient.Client)
	hookId := d.Id()
	projectId := strconv.Itoa(d.Get("project_id").(int))

	hookUpdates := models.HookUpdateRequest{
		Servers: d.Get("servers").([]int),
	}

	log.Printf("[INFO] [ENVOYER:resourceHookUpdate] project updates: %#v", hookUpdates)

	_, err := client.UpdateHook(projectId, hookId, hookUpdates)
	if err != nil {
		return err
	}

	return resourceHookRead(ctx, d, meta)
}

func resourceHookDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*apiClient.Client)

	hookId := d.Id()
	projectId := strconv.Itoa(d.Get("project_id").(int))

	err := c.DeleteHook(projectId, hookId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diag.Diagnostics{}
}
