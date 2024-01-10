package resourcecomponents

import (
	"context"

	"go.uber.org/multierr"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	optypes "github.com/seal-io/walrus/pkg/operator/types"
)

type StateResult struct {
	Error         bool
	Transitioning bool
}

func (r *StateResult) Merge(s StateResult) {
	r.merge(s.Error, s.Transitioning)
}

func (r *StateResult) merge(isError, isTransitioning bool) {
	switch {
	case isError:
		r.Error = true
		r.Transitioning = false
	case !r.Error && isTransitioning:
		r.Transitioning = true
	}
}

// State gets status of the given model.ResourceComponent list with the given operator.Operator,
// and represents is ready if both `Error` and `Transitioning` of StateResult are false.
//
// The given model.ResourceComponent item must be instance shape and not data mode.
//
// The given model.ResourceComponent item must specify the following fields:
// Shape, Mode, Status, ID, DeployerType, Type and Name.
func State(
	ctx context.Context,
	op optypes.Operator,
	modelClient model.ClientSet,
	candidates []*model.ResourceComponent,
) (StateResult, error) {
	var sr StateResult

	if op == nil {
		return sr, nil
	}

	// Merge the errors to return them all at once,
	// instead of returning the first error.
	var berr error

	for i := range candidates {
		// Give up the loop if the context is canceled.
		if multierr.AppendInto(&berr, ctx.Err()) {
			break
		}

		// Skip the resource if it is not instance shape or data mode.
		if candidates[i].Shape != types.ResourceComponentShapeInstance ||
			candidates[i].Mode == types.ResourceComponentModeData {
			continue
		}

		// Get status of the resource component.
		st, err := op.GetStatus(ctx, candidates[i])
		if err != nil {
			berr = multierr.Append(berr, err)
		} else {
			sr.merge(st.Error, st.Transitioning)
		}

		if candidates[i].Status.Equal(*st) {
			// Do not update if the status is same as previous.
			continue
		}

		err = modelClient.ResourceComponents().UpdateOne(candidates[i]).
			SetStatus(*st).
			Exec(ctx)
		if err != nil {
			if model.IsNotFound(err) {
				// Resource component has been deleted by other thread processing.
				continue
			}
			berr = multierr.Append(berr, err)
		}
	}

	return sr, berr
}
