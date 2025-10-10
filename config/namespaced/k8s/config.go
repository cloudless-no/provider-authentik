package k8s

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "k8s"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_service_connection_kubernetes", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ServiceConnection"
	})
}
