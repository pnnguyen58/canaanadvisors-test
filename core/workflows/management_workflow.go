package workflows

import (
	"canaanadvisors-test/proto/management"
	"go.temporal.io/sdk/workflow"
)

func GetMenuWorkflow(ctx workflow.Context, flowInput *management.MenuGetRequest) (
	*management.MenuGetResponse, error) {
	return nil, nil
}
