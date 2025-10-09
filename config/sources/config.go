package sources

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "sources"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_source_kerberos", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Kerberos"
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["enrollment_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_source_ldap", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Ldap"
	})

	p.AddResourceConfigurator("authentik_source_oauth", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Oauth"
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

	p.AddResourceConfigurator("authentik_source_plex", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Plex"
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["enrollment_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_source_saml", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Saml"
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
}
