package main

import (
	"canaanadvisors-test/core/app"
	"canaanadvisors-test/proto/order"
	"context"
	"go.uber.org/zap"
)

type OrderHandler interface {
	CreateOrder(context.Context, *order.OrderCreateRequest) (*order.OrderCreateResponse, error)
}

func NewOrderHandler(logger *zap.Logger, app app.Order) OrderHandler {
	return &OrderController{logger: logger, app: app}
}

type OrderController struct {
	order.UnimplementedOrderServiceServer
	logger *zap.Logger
	app app.Order
}

func (ec *OrderController) CreateOrder(ctx context.Context, req *order.OrderCreateRequest) (*order.OrderCreateResponse, error) {
	return ec.app.CreateOrchestration(ctx, req)
}
