package rbac

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "rbac"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_rbac_initial_permissions", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "InitialPermissions"
		r.References["role"] = config.Reference{
			TerraformName: "authentik_rbac_role",
		}
		// permissions list can reference multiple types, runtime resolution
	})

	p.AddResourceConfigurator("authentik_rbac_permission_role", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "PermissionRole"
		r.References["role"] = config.Reference{
			TerraformName: "authentik_rbac_role",
		}
	})

	p.AddResourceConfigurator("authentik_rbac_permission_user", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "PermissionUser"
		r.References["user"] = config.Reference{
			TerraformName: "authentik_user",
		}
	})

	p.AddResourceConfigurator("authentik_rbac_role", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Role"
	})
}
