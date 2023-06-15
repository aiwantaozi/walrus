package types

import (
	"context"

	"github.com/seal-io/seal/pkg/dao/model"
)

// Type indicates the type of Deployer,
// e.g. Terraform, KubeVela, etc.
type Type = string

// Deployer holds the actions that a deployer must satisfy.
type Deployer interface {
	// Type returns Type.
	Type() Type

	// Apply creates/updates the resources of the given service,
	// also cleans stale resources.
	Apply(context.Context, *model.Service, ApplyOptions) error

	// Destroy cleans all resources of the given service.
	Destroy(context.Context, *model.Service, DestroyOptions) error

	// Rollback service with options.
	Rollback(context.Context, *model.Service, RollbackOptions) error
}

// ApplyOptions holds the options of Deployer's Apply action.
type ApplyOptions struct {
	SkipTLSVerify bool
	// Tags is the service revision tags.
	Tags []string
}

// DestroyOptions holds the options of Deployer's Destroy action.
type DestroyOptions struct {
	SkipTLSVerify bool
}

// RollbackOptions hold the options of Deployer's Rollback action.
type RollbackOptions struct {
	SkipTLSVerify bool
	// CloneFrom is the service revision to clone from.
	CloneFrom *model.ServiceRevision
	// Tags is the service revision tags.
	Tags []string
}