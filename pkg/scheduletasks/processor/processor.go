package processor

import (
	"context"
	"fmt"
	"time"

	"github.com/seal-io/utils/pools/gopool"
	"k8s.io/klog/v2"

	"github.com/seal-io/walrus/pkg/scheduletasks/processors/catalog"
	"github.com/seal-io/walrus/pkg/scheduletasks/processors/connector"
	"github.com/seal-io/walrus/pkg/scheduletasks/processors/resource"
	"github.com/seal-io/walrus/pkg/scheduletasks/processors/resourcecomponents"
	"github.com/seal-io/walrus/pkg/scheduletasks/processors/telemetry"
)

// Processor defines the interface to hold the job executing main logic.
type Processor interface {
	// Name define the task name.
	Name() string

	// Process executes the processor main logic.
	Process(ctx context.Context, id string) error
}

var (
	processors = make(map[string]Processor)
	registers  = []Processor{
		new(catalog.Processor),
		new(connector.Processor),
		new(resource.Processor),
		new(resourcecomponents.Processor),
		new(telemetry.Processor),
	}
)

func init() {
	for i := range registers {
		name := registers[i].Name()
		processors[name] = registers[i]
	}
}

func Existed(name string) error {
	_, ok := processors[name]
	if !ok {
		return fmt.Errorf("processor %s not found", name)
	}
	return nil
}

func Process(ctx context.Context, name, id string) (err error) {
	logger := klog.Background().WithName("processor").WithName(name).WithValues("task", id)

	wg := gopool.GroupWithContextIn(ctx)
	wg.Go(func(ctx context.Context) error {
		logger.V(5).Infof("processing, time %s", time.Now().String())

		p, ok := processors[name]
		if !ok {
			return fmt.Errorf("processor %s not found", name)
		}

		err = p.Process(ctx, id)
		if err != nil {
			logger.V(5).Errorf(err, "failed to process, time %s", time.Now().String())
			return err
		}

		logger.V(5).Infof("processed, time %s", time.Now().String())
		return nil
	})

	return wg.Wait()
}
