package main

import (
	"canaanadvisors-test/core/app"
	"canaanadvisors-test/proto/management"
	"context"
	"go.uber.org/zap"
)

type ManagementHandler interface {
	GetMenu(context.Context, *management.MenuGetRequest) (*management.MenuGetResponse, error)
}

func NewManagementHandler(logger *zap.Logger, app app.Management) ManagementHandler {
	return &ManagementController{logger: logger, app: app}
}

type ManagementController struct {
	management.UnimplementedManagementServiceServer
	logger *zap.Logger
	app app.Management
}

func (mc *ManagementController) GetMenu(ctx context.Context, req *management.MenuGetRequest) (
	*management.MenuGetResponse, error) {
	return mc.app.GetMenuOrchestration(ctx, req)
}