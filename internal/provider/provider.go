// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure NetcalcProvider satisfies various provider interfaces.
var _ provider.Provider = &NetcalcProvider{}

// NetcalcProvider defines the provider implementation.
type NetcalcProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SubnetCalculatorProviderModel describes the provider data model.
type SubnetCalculatorProviderModel struct{}

func (p *NetcalcProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "netcalc"
	resp.Version = p.version
}

func (p *NetcalcProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *NetcalcProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data SubnetCalculatorProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *NetcalcProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSubnetsResource,
	}
}

func (p *NetcalcProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NetcalcProvider{
			version: version,
		}
	}
}
