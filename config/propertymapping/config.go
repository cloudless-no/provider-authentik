package propertymapping

import "github.com/crossplane/upjet/pkg/config"

const ShortGroup = "propertymapping"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// Deprecated
	// p.AddResourceConfigurator("authentik_property_mapping_google_workspace", func(r *config.Resource) {
	// 	r.ShortGroup = ShortGroup
	// 	r.Kind = "GoogleWorkspace"
	// })

	// Deprecated
	// p.AddResourceConfigurator("authentik_property_mapping_ldap", func(r *config.Resource) {
	// 	r.ShortGroup = ShortGroup
	// 	r.Kind = "Ldap"
	// })

	// Deprecated
	// p.AddResourceConfigurator("authentik_property_mapping_microsoft_entra", func(r *config.Resource) {
	// 	r.ShortGroup = ShortGroup
	// 	r.Kind = "MicrosoftEntra"
	// })

	p.AddResourceConfigurator("authentik_property_mapping_notification", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Notification"
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderGoogleWorkspace"
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderMicrosoftEntra"
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_rac", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderRac"
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_radius", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderRadius"
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_saml", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderSaml"
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_scim", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderScim"
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_scope", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ProviderScope"
	})

	// Deprecated
	// p.AddResourceConfigurator("authentik_property_mapping_rac", func(r *config.Resource) {
	// 	r.ShortGroup = ShortGroup
	// 	r.Kind = "Rac"
	// })

	// Deprecated
	// p.AddResourceConfigurator("authentik_property_mapping_radius", func(r *config.Resource) {
	// 	r.ShortGroup = ShortGroup
	// 	r.Kind = "Radius"
	// })

	// Deprecated
	// p.AddResourceConfigurator("authentik_property_mapping_saml", func(r *config.Resource) {
	// 	r.ShortGroup = ShortGroup
	// 	r.Kind = "Saml"
	// })

	// Deprecated
	// p.AddResourceConfigurator("authentik_property_mapping_scim", func(r *config.Resource) {
	// 	r.ShortGroup = ShortGroup
	// 	r.Kind = "Scim"
	// })

	p.AddResourceConfigurator("authentik_property_mapping_source_kerberos", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourceKerberos"
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_ldap", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourceLdap"
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_oauth", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourceOauth"
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_plex", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourcePlex"
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_saml", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourceSaml"
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_scim", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "SourceScim"
	})

	// Deprecated
	// p.AddResourceConfigurator("authentik_scope_mapping", func(r *config.Resource) {
	// 	r.ShortGroup = ShortGroup
	// 	r.Kind = "ScopeMapping"
	// })
}
