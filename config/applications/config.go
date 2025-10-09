package applications

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_application", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "Application"
		r.References["protocol_provider"] = config.Reference{
			TerraformName: "authentik_provider_oauth2",
		}
	})
	p.AddResourceConfigurator("authentik_application_entitlement", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ApplicationEntitlement"
		r.References["application"] = config.Reference{
			TerraformName: "authentik_application",
		}
	})
	p.AddResourceConfigurator("authentik_outpost", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "Outpost"
		r.References["protocol_providers"] = config.Reference{
			TerraformName: "authentik_provider_proxy",
		}
		r.References["service_connection"] = config.Reference{
			TerraformName: "authentik_service_connection_kubernetes",
		}
	})
	p.AddResourceConfigurator("authentik_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderGoogleWorkspace"
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_google_workspace",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_google_workspace",
		}
	})
	p.AddResourceConfigurator("authentik_provider_ldap", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderLdap"
		r.References["bind_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["unbind_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderMicrosoftEntra"
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_microsoft_entra",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_microsoft_entra",
		}
	})
	p.AddResourceConfigurator("authentik_provider_oauth2", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderOauth2"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_source_oauth",
		}
		r.References["signing_key"] = config.Reference{
			TerraformName: "authentik_certificate_key_pair",
		}
		r.References["jwt_federation_providers"] = config.Reference{
			TerraformName: "authentik_provider_oauth2",
		}
	})
	p.AddResourceConfigurator("authentik_provider_proxy", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderProxy"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_provider_rac", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderRac"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_rac",
		}
	})
	p.AddResourceConfigurator("authentik_provider_radius", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderRadius"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_radius",
		}
	})
	p.AddResourceConfigurator("authentik_provider_saml", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderSaml"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_saml",
		}
	})
	p.AddResourceConfigurator("authentik_provider_scim", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderScim"
		r.References["property_mappings"] = config.Reference{
			TerraformName: "authentik_property_mapping_scim",
		}
		r.References["property_mappings_group"] = config.Reference{
			TerraformName: "authentik_property_mapping_scim",
		}
	})
	p.AddResourceConfigurator("authentik_provider_ssf", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderSsf"
		r.References["jwt_federation_providers"] = config.Reference{
			TerraformName: "authentik_provider_oauth2",
		}
	})
	p.AddResourceConfigurator("authentik_rac_endpoint", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "RacEndpoint"
		r.References["protocol_provider"] = config.Reference{
			TerraformName: "authentik_provider_rac",
		}
	})
	p.AddResourceConfigurator("authentik_service_connection_docker", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ServiceConnectionDocker"
		r.References["tls_authentication"] = config.Reference{
			TerraformName: "authentik_certificate_key_pair",
		}
		r.References["tls_verification"] = config.Reference{
			TerraformName: "authentik_certificate_key_pair",
		}
	})
	p.AddResourceConfigurator("authentik_service_connection_kubernetes", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ServiceConnectionKubernetes"
	})
}
