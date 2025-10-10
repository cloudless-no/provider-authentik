package oauth2

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "oauth2"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_oauth2", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Provider"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_source_oauth",
		}
		r.References["signing_key"] = config.Reference{
			TerraformName: "authentik_certificate_key_pair",
		}
		r.References["jwt_federation_providers"] = config.Reference{
			TerraformName: "authentik_provider_oauth2",
		}
	})
}
