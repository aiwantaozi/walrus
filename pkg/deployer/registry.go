package deployer

import (
	"context"
	"fmt"

	"github.com/seal-io/walrus/pkg/deployer/terraform"
	"github.com/seal-io/walrus/pkg/deployer/types"
)

var dpCreators map[types.Type]types.Creator

func init() {
	// Register deployer creators as below.
	dpCreators = map[types.Type]types.Creator{
		terraform.DeployerType: terraform.NewDeployer,
	}
}

// Get returns types.Deployer with the given types.CreateOptions.
func Get(ctx context.Context, opts types.CreateOptions) (types.Deployer, error) {
	f, exist := dpCreators[opts.Type]
	if !exist {
		return nil, fmt.Errorf("unknown deployer: %s", opts.Type)
	}

	dp, err := f(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("error connecting %s deployer: %w", opts.Type, err)
	}

	return dp, nil
}
