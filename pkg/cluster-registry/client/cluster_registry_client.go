// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/client/clusters"
	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/client/config_items"
	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/client/infrastructure_accounts"
	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/client/node_pool_config_items"
	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/client/node_pools"
)

// Default cluster registry HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new cluster registry HTTP client.
func NewHTTPClient(formats strfmt.Registry) *ClusterRegistry {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cluster registry HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *ClusterRegistry {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cluster registry client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *ClusterRegistry {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(ClusterRegistry)
	cli.Transport = transport
	cli.Clusters = clusters.New(transport, formats)
	cli.ConfigItems = config_items.New(transport, formats)
	cli.InfrastructureAccounts = infrastructure_accounts.New(transport, formats)
	cli.NodePoolConfigItems = node_pool_config_items.New(transport, formats)
	cli.NodePools = node_pools.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// ClusterRegistry is a client for cluster registry
type ClusterRegistry struct {
	Clusters clusters.ClientService

	ConfigItems config_items.ClientService

	InfrastructureAccounts infrastructure_accounts.ClientService

	NodePoolConfigItems node_pool_config_items.ClientService

	NodePools node_pools.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *ClusterRegistry) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Clusters.SetTransport(transport)
	c.ConfigItems.SetTransport(transport)
	c.InfrastructureAccounts.SetTransport(transport)
	c.NodePoolConfigItems.SetTransport(transport)
	c.NodePools.SetTransport(transport)
}
