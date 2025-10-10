package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	applicationsCluster "github.com/unbounded-tech/provider-authentik/config/cluster/applications"
	blueprintsCluster "github.com/unbounded-tech/provider-authentik/config/cluster/blueprints"
	dockerCluster "github.com/unbounded-tech/provider-authentik/config/cluster/docker"
	enterpriseCluster "github.com/unbounded-tech/provider-authentik/config/cluster/enterprise"
	flowsCluster "github.com/unbounded-tech/provider-authentik/config/cluster/flows"
	googleworkspaceCluster "github.com/unbounded-tech/provider-authentik/config/cluster/googleworkspace"
	k8sCluster "github.com/unbounded-tech/provider-authentik/config/cluster/k8s"
	kerberosCluster "github.com/unbounded-tech/provider-authentik/config/cluster/kerberos"
	ldapCluster "github.com/unbounded-tech/provider-authentik/config/cluster/ldap"
	microsoftentraCluster "github.com/unbounded-tech/provider-authentik/config/cluster/microsoftentra"
	notificationsCluster "github.com/unbounded-tech/provider-authentik/config/cluster/notifications"
	oauthCluster "github.com/unbounded-tech/provider-authentik/config/cluster/oauth"
	oauth2Cluster "github.com/unbounded-tech/provider-authentik/config/cluster/oauth2"
	outpostsCluster "github.com/unbounded-tech/provider-authentik/config/cluster/outposts"
	policiesCluster "github.com/unbounded-tech/provider-authentik/config/cluster/policies"
	proxyCluster "github.com/unbounded-tech/provider-authentik/config/cluster/proxy"
	racCluster "github.com/unbounded-tech/provider-authentik/config/cluster/rac"
	radiusCluster "github.com/unbounded-tech/provider-authentik/config/cluster/radius"
	rbacCluster "github.com/unbounded-tech/provider-authentik/config/cluster/rbac"
	samlCluster "github.com/unbounded-tech/provider-authentik/config/cluster/saml"
	scopeCluster "github.com/unbounded-tech/provider-authentik/config/cluster/scope"
	ssfCluster "github.com/unbounded-tech/provider-authentik/config/cluster/ssf"
	stagesCluster "github.com/unbounded-tech/provider-authentik/config/cluster/stages"
	systemCluster "github.com/unbounded-tech/provider-authentik/config/cluster/system"
	usersCluster "github.com/unbounded-tech/provider-authentik/config/cluster/users"

	applicationsNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/applications"
	blueprintsNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/blueprints"
	dockerNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/docker"
	enterpriseNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/enterprise"
	flowsNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/flows"
	googleworkspaceNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/googleworkspace"
	k8sNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/k8s"
	kerberosNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/kerberos"
	ldapNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/ldap"
	microsoftentraNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/microsoftentra"
	notificationsNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/notifications"
	oauthNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/oauth"
	oauth2Namespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/oauth2"
	outpostsNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/outposts"
	policiesNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/policies"
	proxyNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/proxy"
	racNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/rac"
	radiusNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/radius"
	rbacNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/rbac"
	samlNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/saml"
	scopeNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/scope"
	ssfNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/ssf"
	stagesNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/stages"
	systemNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/system"
	usersNamespaced "github.com/unbounded-tech/provider-authentik/config/namespaced/users"
)

const (
	resourcePrefix = "authentik"
	modulePath     = "github.com/unbounded-tech/provider-authentik"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("authentik.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		applicationsCluster.Configure,
		blueprintsCluster.Configure,
		dockerCluster.Configure,
		enterpriseCluster.Configure,
		flowsCluster.Configure,
		googleworkspaceCluster.Configure,
		k8sCluster.Configure,
		kerberosCluster.Configure,
		ldapCluster.Configure,
		microsoftentraCluster.Configure,
		notificationsCluster.Configure,
		oauthCluster.Configure,
		oauth2Cluster.Configure,
		outpostsCluster.Configure,
		policiesCluster.Configure,
		proxyCluster.Configure,
		racCluster.Configure,
		radiusCluster.Configure,
		rbacCluster.Configure,
		samlCluster.Configure,
		scopeCluster.Configure,
		ssfCluster.Configure,
		stagesCluster.Configure,
		systemCluster.Configure,
		usersCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("authentik.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		applicationsNamespaced.Configure,
		blueprintsNamespaced.Configure,
		dockerNamespaced.Configure,
		enterpriseNamespaced.Configure,
		flowsNamespaced.Configure,
		googleworkspaceNamespaced.Configure,
		k8sNamespaced.Configure,
		kerberosNamespaced.Configure,
		ldapNamespaced.Configure,
		microsoftentraNamespaced.Configure,
		notificationsNamespaced.Configure,
		oauthNamespaced.Configure,
		oauth2Namespaced.Configure,
		outpostsNamespaced.Configure,
		policiesNamespaced.Configure,
		proxyNamespaced.Configure,
		racNamespaced.Configure,
		radiusNamespaced.Configure,
		rbacNamespaced.Configure,
		samlNamespaced.Configure,
		scopeNamespaced.Configure,
		ssfNamespaced.Configure,
		stagesNamespaced.Configure,
		systemNamespaced.Configure,
		usersNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
