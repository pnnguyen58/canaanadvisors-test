package app

import (
	"canaanadvisors-test/config"
	"canaanadvisors-test/core/repositories"
	"canaanadvisors-test/core/workflows"
	"canaanadvisors-test/proto/order"
	"context"
	"errors"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

type Order interface {
	CreateOrchestration(context.Context, *order.OrderCreateRequest) (*order.OrderCreateResponse, error)
}

type orderApp struct {
	logger *zap.Logger
	temporalClient client.Client
	tempoWorkflow *config.Workflow
	repo repositories.Order
}

func NewOrder(logger *zap.Logger, cl client.Client, tcf *config.TempoConfig) Order {
	return &orderApp{
		logger: logger,
		temporalClient: cl,
		tempoWorkflow: tcf.Workflows["canaanadvisors-test-order"],
	}
}

// CreateOrchestration create new order use case
func (oa *orderApp) CreateOrchestration(ctx context.Context, req *order.OrderCreateRequest) (*order.OrderCreateResponse, error) {
	res, err := ExecuteWorkflow[*order.OrderCreateRequest, *order.OrderCreateResponse](
		ctx, oa.logger, oa.temporalClient, oa.tempoWorkflow, workflows.CreateOrderWorkflow, req)
	if err != nil {
		oa.logger.Error(err.Error())
		return nil, errors.New("create order failed")
	}
	if res == nil {
		return new(order.OrderCreateResponse), nil
	}
	return res, nil
}
