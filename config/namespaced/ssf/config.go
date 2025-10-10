package ssf

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "ssf"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("authentik_provider_ssf", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Provider"
		r.References["jwt_federation_providers"] = config.Reference{
			TerraformName: "authentik_provider_oauth2",
		}
	})
}
