package kerberos

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "kerberos"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_source_kerberos", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Source"
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["enrollment_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_kerberos", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourcePropertyMapping"
	})
}
