package docker

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "docker"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_service_connection_docker", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ServiceConnection"
		r.References["tls_authentication"] = config.Reference{
			TerraformName: "authentik_certificate_key_pair",
		}
		r.References["tls_verification"] = config.Reference{
			TerraformName: "authentik_certificate_key_pair",
		}
	})
}
