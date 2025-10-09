package applications

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "applications"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "GoogleWorkspace"
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_google_workspace",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_google_workspace",
		}
	})

	p.AddResourceConfigurator("authentik_provider_ldap", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "LDAP"
		r.References["bind_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["unbind_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "MicrosoftEntra"
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_microsoft_entra",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_microsoft_entra",
		}
	})

	p.AddResourceConfigurator("authentik_provider_oauth2", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Oauth2"
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

	p.AddResourceConfigurator("authentik_provider_proxy", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Proxy"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_provider_rac", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "RAC"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_rac",
		}
	})

	p.AddResourceConfigurator("authentik_provider_radius", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Radius"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_radius",
		}
	})

	p.AddResourceConfigurator("authentik_provider_saml", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SAML"
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
			TerraformName: "authentik_property_mapping_saml",
		}
	})

	p.AddResourceConfigurator("authentik_provider_scim", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SCIM"
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_scim",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_scim",
		}
	})

	p.AddResourceConfigurator("authentik_provider_ssf", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SSF"
		r.References["jwt_federation_providers"] = config.Reference{
			TerraformName: "authentik_provider_oauth2",
		}
	})

	p.AddResourceConfigurator("authentik_rac_endpoint", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "RACEndpoint"
		r.References["protocol_provider"] = config.Reference{
			TerraformName: "authentik_provider_rac",
		}
	})
}
