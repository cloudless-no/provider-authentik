// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	application "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/applications/application"
	entitlement "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/applications/entitlement"
	blueprint "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/blueprints/blueprint"
	serviceconnection "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/docker/serviceconnection"
	license "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/enterprise/license"
	flow "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/flows/flow"
	flowstagebinding "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/flows/flowstagebinding"
	provider "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/googleworkspace/provider"
	providerpropertymapping "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/googleworkspace/providerpropertymapping"
	serviceconnectionk8s "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/k8s/serviceconnection"
	source "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/kerberos/source"
	providerldap "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/ldap/provider"
	sourceldap "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/ldap/source"
	sourcepropertymapping "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/ldap/sourcepropertymapping"
	providermicrosoftentra "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/microsoftentra/provider"
	providerpropertymappingmicrosoftentra "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/microsoftentra/providerpropertymapping"
	propertymapping "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/notifications/propertymapping"
	rule "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/notifications/rule"
	transport "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/notifications/transport"
	sourceoauth "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/oauth/source"
	sourcepropertymappingoauth "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/oauth/sourcepropertymapping"
	provideroauth2 "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/oauth2/provider"
	outpost "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/outposts/outpost"
	binding "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/policies/binding"
	dummy "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/policies/dummy"
	eventmatcher "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/policies/eventmatcher"
	expiry "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/policies/expiry"
	expression "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/policies/expression"
	geoip "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/policies/geoip"
	password "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/policies/password"
	reputation "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/policies/reputation"
	uniquepassword "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/policies/uniquepassword"
	mappingproviderscim "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/property/mappingproviderscim"
	mappingsourceplex "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/property/mappingsourceplex"
	mappingsourcescim "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/property/mappingsourcescim"
	scim "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/provider/scim"
	providerconfig "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/providerconfig"
	providerproxy "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/proxy/provider"
	endpoint "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/rac/endpoint"
	providerrac "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/rac/provider"
	providerpropertymappingrac "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/rac/providerpropertymapping"
	providerradius "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/radius/provider"
	providerpropertymappingradius "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/radius/providerpropertymapping"
	initialpermissions "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/rbac/initialpermissions"
	permissionrole "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/rbac/permissionrole"
	permissionuser "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/rbac/permissionuser"
	role "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/rbac/role"
	providersaml "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/saml/provider"
	providerpropertymappingsaml "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/saml/providerpropertymapping"
	sourcesaml "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/saml/source"
	sourcepropertymappingsaml "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/saml/sourcepropertymapping"
	providerpropertymappingscope "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/scope/providerpropertymapping"
	plex "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/source/plex"
	providerssf "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/ssf/provider"
	authenticatorduo "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/authenticatorduo"
	authenticatoremail "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/authenticatoremail"
	authenticatorendpointgdtc "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/authenticatorendpointgdtc"
	authenticatorsms "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/authenticatorsms"
	authenticatorstatic "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/authenticatorstatic"
	authenticatortotp "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/authenticatortotp"
	authenticatorvalidate "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/authenticatorvalidate"
	authenticatorwebauthn "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/authenticatorwebauthn"
	captcha "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/captcha"
	consent "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/consent"
	deny "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/deny"
	dummystages "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/dummy"
	email "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/email"
	identification "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/identification"
	invitation "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/invitation"
	mutualtls "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/mutualtls"
	passwordstages "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/password"
	prompt "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/prompt"
	promptfield "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/promptfield"
	redirect "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/redirect"
	sourcestages "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/source"
	userdelete "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/userdelete"
	userlogin "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/userlogin"
	userlogout "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/userlogout"
	userwrite "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/stages/userwrite"
	brand "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/system/brand"
	certificatekeypair "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/system/certificatekeypair"
	systemsettings "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/system/systemsettings"
	group "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/users/group"
	token "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/users/token"
	user "github.com/unbounded-tech/provider-authentik/internal/controller/cluster/users/user"
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
