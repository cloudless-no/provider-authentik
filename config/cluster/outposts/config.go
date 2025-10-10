package outposts

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "outposts"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_outpost", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Outpost"
		r.References["protocol_providers"] = config.Reference{
			TerraformName: "authentik_provider_proxy",
		}
		r.References["service_connection"] = config.Reference{
			TerraformName: "authentik_service_connection_kubernetes",
		}
	})
}
