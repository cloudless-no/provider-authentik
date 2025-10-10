package stages

import "github.com/crossplane/upjet/v2/pkg/config"

const ShortGroup = "stages"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_stage_authenticator_duo", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "AuthenticatorDuo"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_stage_authenticator_email", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "AuthenticatorEmail"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_stage_authenticator_endpoint_gdtc", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "AuthenticatorEndpointGdtc"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_stage_authenticator_sms", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "AuthenticatorSms"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_stage_authenticator_static", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "AuthenticatorStatic"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_stage_authenticator_totp", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "AuthenticatorTotp"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_stage_authenticator_validate", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "AuthenticatorValidate"
	})

	p.AddResourceConfigurator("authentik_stage_authenticator_webauthn", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "AuthenticatorWebauthn"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	// Other stages
	p.AddResourceConfigurator("authentik_stage_captcha", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Captcha"
	})

	p.AddResourceConfigurator("authentik_stage_consent", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Consent"
	})

	p.AddResourceConfigurator("authentik_stage_deny", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Deny"
	})

	p.AddResourceConfigurator("authentik_stage_dummy", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Dummy"
	})

	p.AddResourceConfigurator("authentik_stage_email", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Email"
	})

	p.AddResourceConfigurator("authentik_stage_identification", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Identification"
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
		r.ShortGroup = ShortGroup
		r.Kind = "Invitation"
	})

	p.AddResourceConfigurator("authentik_stage_mutual_tls", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "MutualTls"
	})

	p.AddResourceConfigurator("authentik_stage_password", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Password"
		r.References["configure_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_stage_prompt", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Prompt"
		r.References["fields"] = config.Reference{
			TerraformName: "authentik_stage_prompt_field",
		}
		// validation_policies can be multiple policy types - runtime resolution
	})

	p.AddResourceConfigurator("authentik_stage_prompt_field", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "PromptField"
	})

	p.AddResourceConfigurator("authentik_stage_redirect", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Redirect"
		r.References["target_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})

	p.AddResourceConfigurator("authentik_stage_source", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Source"
		// source can be multiple source types - runtime resolution
	})

	p.AddResourceConfigurator("authentik_stage_user_delete", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "UserDelete"
	})

	p.AddResourceConfigurator("authentik_stage_user_login", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "UserLogin"
	})

	p.AddResourceConfigurator("authentik_stage_user_logout", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "UserLogout"
	})

	p.AddResourceConfigurator("authentik_stage_user_write", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "UserWrite"
		r.References["create_users_group"] = config.Reference{
			TerraformName: "authentik_group",
		}
	})
}
