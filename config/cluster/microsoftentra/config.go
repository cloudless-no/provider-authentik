package microsoftentra

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "microsoftentra"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Provider"
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_provider_microsoft_entra",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_provider_microsoft_entra",
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderPropertyMapping"
	})
}
