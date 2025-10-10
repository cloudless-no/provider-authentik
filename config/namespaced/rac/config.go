package rac

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "rac"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_rac", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Provider"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_provider_rac",
		}
	})

	p.AddResourceConfigurator("authentik_rac_endpoint", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Endpoint"
		r.References["protocol_provider"] = config.Reference{
			TerraformName: "authentik_provider_rac",
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_rac", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderPropertyMapping"
	})
}
