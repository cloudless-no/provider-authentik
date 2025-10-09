package events

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "events"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_event_rule", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Rule"
		r.References["destination_group"] = config.Reference{
			TerraformName: "authentik_group",
		}
		// transports likely references event_transport IDs but may be handled differently
	})
	p.AddResourceConfigurator("authentik_event_transport", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Transport"
	})
}
