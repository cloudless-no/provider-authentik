package enterprise

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_enterprise_license", func(r *config.Resource) {
		r.ShortGroup = "enterprise"
		r.Kind = "EnterpriseLicense"
	})
}
