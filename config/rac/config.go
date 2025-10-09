package rac

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "rac"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_rac_endpoint", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Endpoint"
		r.References["protocol_provider"] = config.Reference{
			TerraformName: "authentik_provider_rac",
		}
	})
}
