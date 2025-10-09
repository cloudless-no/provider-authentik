package flows

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// Flow resources
	p.AddResourceConfigurator("authentik_flow", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "Flow"
		// No references - flow is standalone
	})
	p.AddResourceConfigurator("authentik_flow_stage_binding", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "FlowStageBinding"
		r.References["target"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		// stage can be any stage type - removed placeholder reference for runtime resolution
	})

	// Authenticator stages
	p.AddResourceConfigurator("authentik_stage_authenticator_duo", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageAuthenticatorDuo"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_email", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageAuthenticatorEmail"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_endpoint_gdtc", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageAuthenticatorEndpointGdtc"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_sms", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageAuthenticatorSms"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_static", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageAuthenticatorStatic"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_totp", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageAuthenticatorTotp"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_validate", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageAuthenticatorValidate"
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_webauthn", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageAuthenticatorWebauthn"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	// Other stages
	p.AddResourceConfigurator("authentik_stage_captcha", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageCaptcha"
	})
	p.AddResourceConfigurator("authentik_stage_consent", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageConsent"
	})
	p.AddResourceConfigurator("authentik_stage_deny", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageDeny"
	})
	p.AddResourceConfigurator("authentik_stage_dummy", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageDummy"
	})
	p.AddResourceConfigurator("authentik_stage_email", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageEmail"
	})
	p.AddResourceConfigurator("authentik_stage_identification", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageIdentification"
		r.References["enrollment_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["passwordless_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["recovery_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["password_stage"] = config.Reference{
			TerraformName: "authentik_stage_password",
		}
		r.References["captcha_stage"] = config.Reference{
			TerraformName: "authentik_stage_captcha",
		}
		// sources can be multiple type types - runtime resolution
	})
	p.AddResourceConfigurator("authentik_stage_invitation", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageInvitation"
	})
	p.AddResourceConfigurator("authentik_stage_mutual_tls", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageMutualTls"
	})
	p.AddResourceConfigurator("authentik_stage_password", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StagePassword"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_stage_prompt", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StagePrompt"
		r.References["fields"] = config.Reference{
			TerraformName: "authentik_stage_prompt_field",
		}
		// validation_policies can be multiple policy types - runtime resolution
	})
	p.AddResourceConfigurator("authentik_stage_prompt_field", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StagePromptField"
	})
	p.AddResourceConfigurator("authentik_stage_redirect", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageRedirect"
		r.References["target_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_stage_source", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageSource"
		// source can be multiple source types - runtime resolution
	})
	p.AddResourceConfigurator("authentik_stage_user_delete", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageUserDelete"
	})
	p.AddResourceConfigurator("authentik_stage_user_login", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageUserLogin"
	})
	p.AddResourceConfigurator("authentik_stage_user_logout", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageUserLogout"
	})
	p.AddResourceConfigurator("authentik_stage_user_write", func(r *config.Resource) {
		r.ShortGroup = "flows-stages"
		r.Kind = "StageUserWrite"
		r.References["create_users_group"] = config.Reference{
			TerraformName: "authentik_group",
		}
	})
}
