package policies

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "policies"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_policy_binding", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Binding"
		// target can be multiple types: applications, flows, stages, etc.
		// Examples show focusing on applications
		r.References["target"] = config.Reference{
			TerraformName: "authentik_application",
		}
		r.References["group"] = config.Reference{
			TerraformName: "authentik_group",
		}
		// policy can be any policy type, but examples use specific ones
		// Leave generic for now
	})

	p.AddResourceConfigurator("authentik_policy_dummy", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Dummy"
	})

	p.AddResourceConfigurator("authentik_policy_event_matcher", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "EventMatcher"
	})

	p.AddResourceConfigurator("authentik_policy_expiry", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Expiry"
	})

	p.AddResourceConfigurator("authentik_policy_expression", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Expression"
	})

	p.AddResourceConfigurator("authentik_policy_geoip", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Geoip"
	})

	p.AddResourceConfigurator("authentik_policy_password", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Password"
	})

	p.AddResourceConfigurator("authentik_policy_reputation", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Reputation"
	})

	p.AddResourceConfigurator("authentik_policy_unique_password", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "UniquePassword"
	})

}
