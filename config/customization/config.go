package customization

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// Policy resources
	p.AddResourceConfigurator("authentik_policy_binding", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PolicyBinding"
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
		r.ShortGroup = "customization"
		r.Kind = "PolicyDummy"
	})
	p.AddResourceConfigurator("authentik_policy_event_matcher", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PolicyEventMatcher"
	})
	p.AddResourceConfigurator("authentik_policy_expiry", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PolicyExpiry"
	})
	p.AddResourceConfigurator("authentik_policy_expression", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PolicyExpression"
	})
	p.AddResourceConfigurator("authentik_policy_geoip", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PolicyGeoip"
	})
	p.AddResourceConfigurator("authentik_policy_password", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PolicyPassword"
	})
	p.AddResourceConfigurator("authentik_policy_reputation", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PolicyReputation"
	})
	p.AddResourceConfigurator("authentik_policy_unique_password", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PolicyUniquePassword"
	})

	// Property mapping resources
	p.AddResourceConfigurator("authentik_property_mapping_google_workspace", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingGoogleWorkspace"
	})
	p.AddResourceConfigurator("authentik_property_mapping_ldap", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingLdap"
	})
	p.AddResourceConfigurator("authentik_property_mapping_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingMicrosoftEntra"
	})
	p.AddResourceConfigurator("authentik_property_mapping_notification", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingNotification"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingProviderGoogleWorkspace"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingProviderMicrosoftEntra"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_rac", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingProviderRac"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_radius", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingProviderRadius"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_saml", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingProviderSaml"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_scim", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingProviderScim"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_scope", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingProviderScope"
	})
	p.AddResourceConfigurator("authentik_property_mapping_rac", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingRac"
	})
	p.AddResourceConfigurator("authentik_property_mapping_radius", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingRadius"
	})
	p.AddResourceConfigurator("authentik_property_mapping_saml", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingSaml"
	})
	p.AddResourceConfigurator("authentik_property_mapping_scim", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingScim"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_kerberos", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingSourceKerberos"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_ldap", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingSourceLdap"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_oauth", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingSourceOauth"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_plex", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingSourcePlex"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_saml", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingSourceSaml"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_scim", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "PropertyMappingSourceScim"
	})
	p.AddResourceConfigurator("authentik_scope_mapping", func(r *config.Resource) {
		r.ShortGroup = "customization"
		r.Kind = "ScopeMapping"
	})
}
