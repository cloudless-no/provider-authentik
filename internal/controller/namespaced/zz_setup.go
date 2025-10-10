// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	application "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/applications/application"
	entitlement "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/applications/entitlement"
	blueprint "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/blueprints/blueprint"
	serviceconnection "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/docker/serviceconnection"
	license "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/enterprise/license"
	flow "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/flows/flow"
	flowstagebinding "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/flows/flowstagebinding"
	provider "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/googleworkspace/provider"
	providerpropertymapping "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/googleworkspace/providerpropertymapping"
	serviceconnectionk8s "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/k8s/serviceconnection"
	source "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/kerberos/source"
	providerldap "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/ldap/provider"
	sourceldap "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/ldap/source"
	sourcepropertymapping "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/ldap/sourcepropertymapping"
	providermicrosoftentra "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/microsoftentra/provider"
	providerpropertymappingmicrosoftentra "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/microsoftentra/providerpropertymapping"
	propertymapping "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/notifications/propertymapping"
	rule "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/notifications/rule"
	transport "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/notifications/transport"
	sourceoauth "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/oauth/source"
	sourcepropertymappingoauth "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/oauth/sourcepropertymapping"
	provideroauth2 "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/oauth2/provider"
	outpost "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/outposts/outpost"
	binding "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/policies/binding"
	dummy "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/policies/dummy"
	eventmatcher "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/policies/eventmatcher"
	expiry "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/policies/expiry"
	expression "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/policies/expression"
	geoip "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/policies/geoip"
	password "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/policies/password"
	reputation "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/policies/reputation"
	uniquepassword "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/policies/uniquepassword"
	mappingproviderscim "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/property/mappingproviderscim"
	mappingsourceplex "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/property/mappingsourceplex"
	mappingsourcescim "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/property/mappingsourcescim"
	scim "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/provider/scim"
	providerconfig "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/providerconfig"
	providerproxy "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/proxy/provider"
	endpoint "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/rac/endpoint"
	providerrac "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/rac/provider"
	providerpropertymappingrac "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/rac/providerpropertymapping"
	providerradius "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/radius/provider"
	providerpropertymappingradius "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/radius/providerpropertymapping"
	initialpermissions "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/rbac/initialpermissions"
	permissionrole "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/rbac/permissionrole"
	permissionuser "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/rbac/permissionuser"
	role "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/rbac/role"
	providersaml "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/saml/provider"
	providerpropertymappingsaml "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/saml/providerpropertymapping"
	sourcesaml "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/saml/source"
	sourcepropertymappingsaml "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/saml/sourcepropertymapping"
	providerpropertymappingscope "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/scope/providerpropertymapping"
	plex "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/source/plex"
	providerssf "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/ssf/provider"
	authenticatorduo "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/authenticatorduo"
	authenticatoremail "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/authenticatoremail"
	authenticatorendpointgdtc "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/authenticatorendpointgdtc"
	authenticatorsms "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/authenticatorsms"
	authenticatorstatic "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/authenticatorstatic"
	authenticatortotp "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/authenticatortotp"
	authenticatorvalidate "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/authenticatorvalidate"
	authenticatorwebauthn "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/authenticatorwebauthn"
	captcha "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/captcha"
	consent "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/consent"
	deny "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/deny"
	dummystages "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/dummy"
	email "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/email"
	identification "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/identification"
	invitation "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/invitation"
	mutualtls "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/mutualtls"
	passwordstages "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/password"
	prompt "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/prompt"
	promptfield "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/promptfield"
	redirect "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/redirect"
	sourcestages "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/source"
	userdelete "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/userdelete"
	userlogin "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/userlogin"
	userlogout "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/userlogout"
	userwrite "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/stages/userwrite"
	brand "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/system/brand"
	certificatekeypair "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/system/certificatekeypair"
	systemsettings "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/system/systemsettings"
	group "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/users/group"
	token "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/users/token"
	user "github.com/unbounded-tech/provider-authentik/internal/controller/namespaced/users/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		entitlement.Setup,
		blueprint.Setup,
		serviceconnection.Setup,
		license.Setup,
		flow.Setup,
		flowstagebinding.Setup,
		provider.Setup,
		providerpropertymapping.Setup,
		serviceconnectionk8s.Setup,
		source.Setup,
		providerldap.Setup,
		sourceldap.Setup,
		sourcepropertymapping.Setup,
		providermicrosoftentra.Setup,
		providerpropertymappingmicrosoftentra.Setup,
		propertymapping.Setup,
		rule.Setup,
		transport.Setup,
		sourceoauth.Setup,
		sourcepropertymappingoauth.Setup,
		provideroauth2.Setup,
		outpost.Setup,
		binding.Setup,
		dummy.Setup,
		eventmatcher.Setup,
		expiry.Setup,
		expression.Setup,
		geoip.Setup,
		password.Setup,
		reputation.Setup,
		uniquepassword.Setup,
		mappingproviderscim.Setup,
		mappingsourceplex.Setup,
		mappingsourcescim.Setup,
		scim.Setup,
		providerconfig.Setup,
		providerproxy.Setup,
		endpoint.Setup,
		providerrac.Setup,
		providerpropertymappingrac.Setup,
		providerradius.Setup,
		providerpropertymappingradius.Setup,
		initialpermissions.Setup,
		permissionrole.Setup,
		permissionuser.Setup,
		role.Setup,
		providersaml.Setup,
		providerpropertymappingsaml.Setup,
		sourcesaml.Setup,
		sourcepropertymappingsaml.Setup,
		providerpropertymappingscope.Setup,
		plex.Setup,
		providerssf.Setup,
		authenticatorduo.Setup,
		authenticatoremail.Setup,
		authenticatorendpointgdtc.Setup,
		authenticatorsms.Setup,
		authenticatorstatic.Setup,
		authenticatortotp.Setup,
		authenticatorvalidate.Setup,
		authenticatorwebauthn.Setup,
		captcha.Setup,
		consent.Setup,
		deny.Setup,
		dummystages.Setup,
		email.Setup,
		identification.Setup,
		invitation.Setup,
		mutualtls.Setup,
		passwordstages.Setup,
		prompt.Setup,
		promptfield.Setup,
		redirect.Setup,
		sourcestages.Setup,
		userdelete.Setup,
		userlogin.Setup,
		userlogout.Setup,
		userwrite.Setup,
		brand.Setup,
		certificatekeypair.Setup,
		systemsettings.Setup,
		group.Setup,
		token.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.SetupGated,
		entitlement.SetupGated,
		blueprint.SetupGated,
		serviceconnection.SetupGated,
		license.SetupGated,
		flow.SetupGated,
		flowstagebinding.SetupGated,
		provider.SetupGated,
		providerpropertymapping.SetupGated,
		serviceconnectionk8s.SetupGated,
		source.SetupGated,
		providerldap.SetupGated,
		sourceldap.SetupGated,
		sourcepropertymapping.SetupGated,
		providermicrosoftentra.SetupGated,
		providerpropertymappingmicrosoftentra.SetupGated,
		propertymapping.SetupGated,
		rule.SetupGated,
		transport.SetupGated,
		sourceoauth.SetupGated,
		sourcepropertymappingoauth.SetupGated,
		provideroauth2.SetupGated,
		outpost.SetupGated,
		binding.SetupGated,
		dummy.SetupGated,
		eventmatcher.SetupGated,
		expiry.SetupGated,
		expression.SetupGated,
		geoip.SetupGated,
		password.SetupGated,
		reputation.SetupGated,
		uniquepassword.SetupGated,
		mappingproviderscim.SetupGated,
		mappingsourceplex.SetupGated,
		mappingsourcescim.SetupGated,
		scim.SetupGated,
		providerconfig.SetupGated,
		providerproxy.SetupGated,
		endpoint.SetupGated,
		providerrac.SetupGated,
		providerpropertymappingrac.SetupGated,
		providerradius.SetupGated,
		providerpropertymappingradius.SetupGated,
		initialpermissions.SetupGated,
		permissionrole.SetupGated,
		permissionuser.SetupGated,
		role.SetupGated,
		providersaml.SetupGated,
		providerpropertymappingsaml.SetupGated,
		sourcesaml.SetupGated,
		sourcepropertymappingsaml.SetupGated,
		providerpropertymappingscope.SetupGated,
		plex.SetupGated,
		providerssf.SetupGated,
		authenticatorduo.SetupGated,
		authenticatoremail.SetupGated,
		authenticatorendpointgdtc.SetupGated,
		authenticatorsms.SetupGated,
		authenticatorstatic.SetupGated,
		authenticatortotp.SetupGated,
		authenticatorvalidate.SetupGated,
		authenticatorwebauthn.SetupGated,
		captcha.SetupGated,
		consent.SetupGated,
		deny.SetupGated,
		dummystages.SetupGated,
		email.SetupGated,
		identification.SetupGated,
		invitation.SetupGated,
		mutualtls.SetupGated,
		passwordstages.SetupGated,
		prompt.SetupGated,
		promptfield.SetupGated,
		redirect.SetupGated,
		sourcestages.SetupGated,
		userdelete.SetupGated,
		userlogin.SetupGated,
		userlogout.SetupGated,
		userwrite.SetupGated,
		brand.SetupGated,
		certificatekeypair.SetupGated,
		systemsettings.SetupGated,
		group.SetupGated,
		token.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
