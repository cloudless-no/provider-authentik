package system

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "system"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_brand", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Brand"
		r.References["flow_authentication"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["flow_device_code"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["flow_invalidation"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["flow_recovery"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["flow_unenrollment"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["flow_user_settings"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["web_certificate"] = config.Reference{
			TerraformName: "authentik_certificate_key_pair",
		}
		r.References["client_certificates"] = config.Reference{
			TerraformName: "authentik_certificate_key_pair",
		}
	})

	p.AddResourceConfigurator("authentik_certificate_key_pair", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "CertificateKeyPair"
	})

	p.AddResourceConfigurator("authentik_system_settings", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SystemSettings"
	})
}
