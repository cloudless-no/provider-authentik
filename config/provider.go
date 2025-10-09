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
	"github.com/unbounded-tech/provider-authentik/config/customization"
	"github.com/unbounded-tech/provider-authentik/config/directory"
	"github.com/unbounded-tech/provider-authentik/config/enterprise"
	"github.com/unbounded-tech/provider-authentik/config/events"
	"github.com/unbounded-tech/provider-authentik/config/rbac"
	"github.com/unbounded-tech/provider-authentik/config/system"
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
		customization.Configure,
		directory.Configure,
		enterprise.Configure,
		events.Configure,
		// flows_stages.Configure,
		rbac.Configure,
		system.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
