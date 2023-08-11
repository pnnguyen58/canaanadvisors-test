package workflows

import (
	"canaanadvisors-test/proto/notification"
	"go.temporal.io/sdk/workflow"
)

func SendNotificationWorkflow(ctx workflow.Context, flowInput *notification.NotificationSendRequest) (
	*notification.NotificationSendResponse, error) {
	return nil, nil
}
