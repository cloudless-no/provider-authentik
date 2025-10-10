package oauth

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "oauth"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_source_oauth", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Source"
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["enrollment_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_source_oauth",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_source_oauth",
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_oauth", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourcePropertyMapping"
	})
}
