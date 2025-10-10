package blueprints

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "blueprints"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_blueprint", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Blueprint"
	})
}
