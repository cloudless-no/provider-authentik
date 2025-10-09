package applications

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "applications"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_application", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Application"
		r.References["protocol_provider"] = config.Reference{
			TerraformName: "authentik_provider_oauth2",
		}
	})

	p.AddResourceConfigurator("authentik_application_entitlement", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Entitlement"
		r.References["application"] = config.Reference{
			TerraformName: "authentik_application",
		}
	})
}
