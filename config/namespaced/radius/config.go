package radius

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "radius"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_radius", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Provider"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_provider_radius",
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_radius", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderPropertyMapping"
	})
}
