package scim

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "scim"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_scim", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Provider"
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_provider_scim",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_provider_scim",
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_scim", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderPropertyMapping"
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_scim", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourcePropertyMapping"
	})
}
