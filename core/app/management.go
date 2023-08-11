package app

import (
	"canaanadvisors-test/proto/management"
	"context"
)

type Management interface {
	GetMenuOrchestration(context.Context, *management.MenuGetRequest) (*management.MenuGetResponse, error)
}
