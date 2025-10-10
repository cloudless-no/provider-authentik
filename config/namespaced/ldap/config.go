package ldap

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "ldap"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_ldap", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Provider"
		r.References["bind_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["unbind_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_source_ldap", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Source"

		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_source_ldap",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_source_ldap",
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_ldap", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourcePropertyMapping"
	})
}
