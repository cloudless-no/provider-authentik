package directory

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_group", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.Kind = "Group"
		r.References["users"] = config.Reference{
			TerraformName: "authentik_user",
		}
	})
	p.AddResourceConfigurator("authentik_user", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.Kind = "User"
		r.References["groups"] = config.Reference{
			TerraformName: "authentik_group",
		}
	})
	p.AddResourceConfigurator("authentik_token", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.Kind = "Token"
		r.References["user"] = config.Reference{
			TerraformName: "authentik_user",
		}
	})
	p.AddResourceConfigurator("authentik_source_kerberos", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.Kind = "SourceKerberos"
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["enrollment_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_source_ldap", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.Kind = "SourceLdap"
	})
	p.AddResourceConfigurator("authentik_source_oauth", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.Kind = "SourceOauth"
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
		r.ShortGroup = "directory"
		r.Kind = "SourcePlex"
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["enrollment_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_source_saml", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.Kind = "SourceSaml"
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
