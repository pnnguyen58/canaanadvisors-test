package app

import (
	"canaanadvisors-test/config"
	"canaanadvisors-test/core/workflows"
	"canaanadvisors-test/proto/notification"
	"context"
	"errors"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

type Notification interface {
	SendOrchestration(context.Context, *notification.NotificationSendRequest) (
		*notification.NotificationSendResponse, error)
}

type notificationApp struct {
	logger *zap.Logger
	temporalClient client.Client
	tempoWorkflow *config.Workflow
}

func NewNotification(logger *zap.Logger, cl client.Client, tcf *config.TempoConfig) Notification {
	return &notificationApp{
		logger: logger,
		temporalClient: cl,
		tempoWorkflow: tcf.Workflows["canaanadvisors-test-notification"],
	}
}
// SendOrchestration send notification use case
func (na *notificationApp) SendOrchestration(ctx context.Context, req *notification.NotificationSendRequest) (
	*notification.NotificationSendResponse, error) {
	res, err := ExecuteWorkflow[*notification.NotificationSendRequest, *notification.NotificationSendResponse](
		ctx, na.logger, na.temporalClient, na.tempoWorkflow, workflows.SendNotificationWorkflow, req)
	if err != nil {
		na.logger.Error(err.Error())
		return nil, errors.New("send notification failed")
	}
	if res == nil {
		return new(notification.NotificationSendResponse), nil
	}
	return res, nil
}