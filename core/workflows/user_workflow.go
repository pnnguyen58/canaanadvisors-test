package workflows

import (
	"canaanadvisors-test/core/activities"
	"canaanadvisors-test/proto/user"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/multierr"
	"time"
)

// LoginWorkflow workflows definition
func LoginWorkflow(ctx workflow.Context, flowInput *user.LoginRequest) (*user.LoginResponse, error) {

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// This is how you log
	// workflows.GetLogger(ctx).Info("jobInput.Inputs", flowInput.Inputs)

	result := &user.LoginResponse{}
	err := workflow.ExecuteActivity(ctx, activities.Login, flowInput).Get(ctx, result)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, activities.LoginCompensation, flowInput).
				Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()
	workflow.GetLogger(ctx).Info("Workflow completed.")

	return result, err
}


// LogoutWorkflow workflows definition
func LogoutWorkflow(ctx workflow.Context, flowInput *user.LogoutRequest) (*user.LogoutResponse, error) {

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// This is how you log
	// workflows.GetLogger(ctx).Info("jobInput.Inputs", flowInput.Inputs)

	result := &user.LogoutResponse{}
	err := workflow.ExecuteActivity(ctx, activities.Logout, flowInput).Get(ctx, result)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, activities.LogoutCompensation, flowInput).
				Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()
	workflow.GetLogger(ctx).Info("Workflow completed.")

	return result, err
}