package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// applications
	"authentik_application":             config.IdentifierFromProvider,
	"authentik_application_entitlement": config.IdentifierFromProvider,

	// blueprints
	"authentik_blueprint": config.IdentifierFromProvider,

	// docker
	"authentik_service_connection_docker": config.IdentifierFromProvider,

	// enterprise
	"authentik_enterprise_license": config.IdentifierFromProvider,

	// flows
	"authentik_flow":               config.IdentifierFromProvider,
	"authentik_flow_stage_binding": config.IdentifierFromProvider,

	// google workspace
	"authentik_provider_google_workspace":                  config.IdentifierFromProvider,
	"authentik_property_mapping_provider_google_workspace": config.IdentifierFromProvider,

	// k8s
	"authentik_service_connection_kubernetes": config.IdentifierFromProvider,

	// kerberos
	"authentik_source_kerberos": config.IdentifierFromProvider,

	// ldap
	"authentik_provider_ldap":                config.IdentifierFromProvider,
	"authentik_property_mapping_source_ldap": config.IdentifierFromProvider,
	"authentik_source_ldap":                  config.IdentifierFromProvider,

	// microsoftentra
	"authentik_provider_microsoft_entra":                  config.IdentifierFromProvider,
	"authentik_property_mapping_provider_microsoft_entra": config.IdentifierFromProvider,

	// notifications
	"authentik_property_mapping_notification": config.IdentifierFromProvider,
	"authentik_event_rule":                    config.IdentifierFromProvider,
	"authentik_event_transport":               config.IdentifierFromProvider,

	// oauth
	"authentik_property_mapping_source_oauth": config.IdentifierFromProvider,
	"authentik_source_oauth":                  config.IdentifierFromProvider,

	// oauth2
	"authentik_provider_oauth2": config.IdentifierFromProvider,

	// outposts
	"authentik_outpost": config.IdentifierFromProvider,

	// plex
	"authentik_property_mapping_source_plex": config.IdentifierFromProvider,
	"authentik_source_plex":                  config.IdentifierFromProvider,

	// policies
	"authentik_policy_binding":         config.IdentifierFromProvider,
	"authentik_policy_dummy":           config.IdentifierFromProvider,
	"authentik_policy_event_matcher":   config.IdentifierFromProvider,
	"authentik_policy_expiry":          config.IdentifierFromProvider,
	"authentik_policy_expression":      config.IdentifierFromProvider,
	"authentik_policy_geoip":           config.IdentifierFromProvider,
	"authentik_policy_password":        config.IdentifierFromProvider,
	"authentik_policy_reputation":      config.IdentifierFromProvider,
	"authentik_policy_unique_password": config.IdentifierFromProvider,

	// proxy
	"authentik_provider_proxy": config.IdentifierFromProvider,

	// rac
	"authentik_provider_rac":                  config.IdentifierFromProvider,
	"authentik_rac_endpoint":                  config.IdentifierFromProvider,
	"authentik_property_mapping_provider_rac": config.IdentifierFromProvider,

	// radius
	"authentik_provider_radius":                  config.IdentifierFromProvider,
	"authentik_property_mapping_provider_radius": config.IdentifierFromProvider,

	// rbac
	"authentik_rbac_initial_permissions": config.IdentifierFromProvider,
	"authentik_rbac_permission_role":     config.IdentifierFromProvider,
	"authentik_rbac_permission_user":     config.IdentifierFromProvider,
	"authentik_rbac_role":                config.IdentifierFromProvider,

	// saml
	"authentik_provider_saml":                  config.IdentifierFromProvider,
	"authentik_property_mapping_provider_saml": config.IdentifierFromProvider,
	"authentik_property_mapping_source_saml":   config.IdentifierFromProvider,
	"authentik_source_saml":                    config.IdentifierFromProvider,

	// scim
	"authentik_provider_scim":                  config.IdentifierFromProvider,
	"authentik_property_mapping_provider_scim": config.IdentifierFromProvider,
	"authentik_property_mapping_source_scim":   config.IdentifierFromProvider,

	// scope
	"authentik_property_mapping_provider_scope": config.IdentifierFromProvider,

	// ssf
	"authentik_provider_ssf": config.IdentifierFromProvider,

	// stages
	"authentik_stage_authenticator_duo":           config.IdentifierFromProvider,
	"authentik_stage_authenticator_email":         config.IdentifierFromProvider,
	"authentik_stage_authenticator_endpoint_gdtc": config.IdentifierFromProvider,
	"authentik_stage_authenticator_sms":           config.IdentifierFromProvider,
	"authentik_stage_authenticator_static":        config.IdentifierFromProvider,
	"authentik_stage_authenticator_totp":          config.IdentifierFromProvider,
	"authentik_stage_authenticator_validate":      config.IdentifierFromProvider,
	"authentik_stage_authenticator_webauthn":      config.IdentifierFromProvider,
	"authentik_stage_captcha":                     config.IdentifierFromProvider,
	"authentik_stage_consent":                     config.IdentifierFromProvider,
	"authentik_stage_deny":                        config.IdentifierFromProvider,
	"authentik_stage_dummy":                       config.IdentifierFromProvider,
	"authentik_stage_email":                       config.IdentifierFromProvider,
	"authentik_stage_identification":              config.IdentifierFromProvider,
	"authentik_stage_invitation":                  config.IdentifierFromProvider,
	"authentik_stage_mutual_tls":                  config.IdentifierFromProvider,
	"authentik_stage_password":                    config.IdentifierFromProvider,
	"authentik_stage_prompt":                      config.IdentifierFromProvider,
	"authentik_stage_prompt_field":                config.IdentifierFromProvider,
	"authentik_stage_redirect":                    config.IdentifierFromProvider,
	"authentik_stage_source":                      config.IdentifierFromProvider,
	"authentik_stage_user_delete":                 config.IdentifierFromProvider,
	"authentik_stage_user_login":                  config.IdentifierFromProvider,
	"authentik_stage_user_logout":                 config.IdentifierFromProvider,
	"authentik_stage_user_write":                  config.IdentifierFromProvider,

	// system
	"authentik_brand":                config.IdentifierFromProvider,
	"authentik_certificate_key_pair": config.IdentifierFromProvider,
	"authentik_system_settings":      config.IdentifierFromProvider,

	// users
	"authentik_group": config.IdentifierFromProvider,
	"authentik_token": config.IdentifierFromProvider,
	"authentik_user":  config.IdentifierFromProvider,
}

func idWithStub() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		en, _ := config.IDAsExternalName(tfstate)
		return en, nil
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
