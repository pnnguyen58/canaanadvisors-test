package main

import (
	"canaanadvisors-test/core/app"
	"canaanadvisors-test/proto/notification"
	"context"
	"go.uber.org/zap"
)

type NotificationHandler interface {
	SendNotification(context.Context, *notification.NotificationSendRequest) (*notification.NotificationSendResponse, error)
}

func NewNotificationHandler(logger *zap.Logger, app app.Notification) NotificationHandler {
	return &NotificationController{logger: logger, app: app}
}

type NotificationController struct {
	notification.UnimplementedNotificationServiceServer
	logger *zap.Logger
	app app.Notification
}

func (nc *NotificationController) SendNotification(ctx context.Context, req *notification.NotificationSendRequest) (
	*notification.NotificationSendResponse, error) {
	return nc.app.SendOrchestration(ctx, req)
}