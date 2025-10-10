package flows

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "flows"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// Flow resources
	p.AddResourceConfigurator("authentik_flow", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Flow"
		// No references - flow is standalone
	})
	p.AddResourceConfigurator("authentik_flow_stage_binding", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "FlowStageBinding"
		r.References["target"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		// stage can be any stage type - removed placeholder reference for runtime resolution
	})
}
