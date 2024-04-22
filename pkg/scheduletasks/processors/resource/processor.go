package resource

import (
	"context"

	"k8s.io/klog/v2"

	"github.com/seal-io/walrus/pkg/system"
)

const Name = "resource-relationship-check"

type Processor struct{}

func (in *Processor) Name() string {
	return Name
}

func (in *Processor) Process(ctx context.Context, id string) error {
	_ = klog.Background().WithName("processor").WithName(in.Name()).WithValues("task", id)
	_ = system.LoopbackKubeClient.Get()

	// TODO: your logic here.
	return nil
}
