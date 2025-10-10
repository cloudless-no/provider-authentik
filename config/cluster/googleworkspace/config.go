package googleworkspace

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "googleworkspace"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Provider"
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_provider_google_workspace",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_provider_google_workspace",
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderPropertyMapping"
	})
}
