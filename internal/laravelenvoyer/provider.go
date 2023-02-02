package laravelenvoyer

import (
	"context"
	"github.com/hashicorp/terraform-provider-laravelenvoyer/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support Markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example, you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"token": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("LARAVELENVOYER_TOKEN", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"laravelenvoyer_environment": dataSourceEnvironment(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"laravelenvoyer_project":     resourceProject(),
				"laravelenvoyer_server":      resourceServer(),
				"laravelenvoyer_hook":        resourceHook(),
				"laravelenvoyer_environment": resourceEnvironment(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, resourceData *schema.ResourceData) (any, diag.Diagnostics) {
		token := resourceData.Get("token").(string)

		// Warning or errors can be collected in a slice type
		var diags diag.Diagnostics

		envoyerApiClient, err := client.NewClient(nil, &token)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Laravel Envoyer envoyerApiClient",
				Detail:   "Unable to auth user for authenticated Laravel Envoyer envoyerApiClient",
			})
			return nil, diags
		}

		return envoyerApiClient, nil
	}
}
