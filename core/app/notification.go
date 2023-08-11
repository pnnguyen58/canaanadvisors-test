package app

import (
	"canaanadvisors-test/proto/notification"
	"context"
)

type Notification interface {
	SendOrchestration(context.Context, *notification.NotificationSendRequest,
		) (*notification.NotificationSendResponse, error)
}
