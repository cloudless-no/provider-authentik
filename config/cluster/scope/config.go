package scope

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "scope"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_property_mapping_provider_scope", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderPropertyMapping"
	})
}
