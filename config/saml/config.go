package saml

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "saml"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_saml", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Provider"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_provider_saml",
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_saml", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderPropertyMapping"
	})

	p.AddResourceConfigurator("authentik_source_saml", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Source"
		r.References["pre_authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["enrollment_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_source_saml",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_source_saml",
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_saml", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourcePropertyMapping"
	})
}
