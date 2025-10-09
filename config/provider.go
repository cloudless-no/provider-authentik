/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/unbounded-tech/provider-authentik/config/applications"
	"github.com/unbounded-tech/provider-authentik/config/blueprints"
	"github.com/unbounded-tech/provider-authentik/config/docker"
	"github.com/unbounded-tech/provider-authentik/config/enterprise"
	"github.com/unbounded-tech/provider-authentik/config/events"
	"github.com/unbounded-tech/provider-authentik/config/flows"
	"github.com/unbounded-tech/provider-authentik/config/k8s"
	"github.com/unbounded-tech/provider-authentik/config/outposts"
	"github.com/unbounded-tech/provider-authentik/config/policies"
	"github.com/unbounded-tech/provider-authentik/config/propertymapping"
	"github.com/unbounded-tech/provider-authentik/config/providers"
	"github.com/unbounded-tech/provider-authentik/config/rac"
	"github.com/unbounded-tech/provider-authentik/config/rbac"
	"github.com/unbounded-tech/provider-authentik/config/sources"
	"github.com/unbounded-tech/provider-authentik/config/stages"
	"github.com/unbounded-tech/provider-authentik/config/system"
	"github.com/unbounded-tech/provider-authentik/config/users"
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
		applications.Configure,
		blueprints.Configure,
		docker.Configure,
		enterprise.Configure,
		events.Configure,
		flows.Configure,
		k8s.Configure,
		outposts.Configure,
		policies.Configure,
		propertymapping.Configure,
		providers.Configure,
		rac.Configure,
		rbac.Configure,
		sources.Configure,
		stages.Configure,
		users.Configure,
		system.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
