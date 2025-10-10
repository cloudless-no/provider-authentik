package users

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "users"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_group", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Group"
		r.References["users"] = config.Reference{
			TerraformName: "authentik_user",
		}
	})

	p.AddResourceConfigurator("authentik_user", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "User"
		r.References["groups"] = config.Reference{
			TerraformName: "authentik_group",
		}
	})

	p.AddResourceConfigurator("authentik_token", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Token"
		r.References["user"] = config.Reference{
			TerraformName: "authentik_user",
		}
	})
}
