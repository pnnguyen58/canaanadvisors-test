package workflows

import (
	"canaanadvisors-test/core/activities"
	"canaanadvisors-test/proto/order"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/multierr"
	"time"
)

// CreateOrderWorkflow workflows definition
func CreateOrderWorkflow(ctx workflow.Context, flowInput *order.OrderCreateRequest) (
	*order.OrderCreateResponse, error) {
	// Workflow has to check input valid or not
	//inputErr := flowInput.CheckValid()
	//if inputErr != nil {
	//	return nil,
	//		temporal.NewNonRetryableApplicationError("Invalid flow input", common.ErrInvalidInput, inputErr, nil)
	//}

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// This is how you log
	// workflows.GetLogger(ctx).Info("jobInput.Inputs", flowInput.Inputs)

	result := &order.OrderCreateResponse{}
	err := workflow.ExecuteActivity(ctx, activities.CreateOrder, flowInput).Get(ctx, result)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, activities.CreateOrderCompensation, flowInput).
				Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()
	workflow.GetLogger(ctx).Info("Workflow completed.")

	return result, err
}