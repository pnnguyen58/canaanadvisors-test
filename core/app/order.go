package app

import (
	"canaanadvisors-test/config"
	"canaanadvisors-test/core/workflows"
	"canaanadvisors-test/proto/order"
	"context"
	"fmt"
	"github.com/google/uuid"
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
}

func NewOrderApp(ctx context.Context, logger *zap.Logger, cl client.Client, tcf *config.TempoConfig) Order {
	return &orderApp{
		logger: logger,
		temporalClient: cl,
		tempoWorkflow: tcf.Workflows["canaanadvisors-test-order"],
	}
}

// CreateOrchestration create new example use case
func (oa *orderApp) CreateOrchestration(ctx context.Context, req *order.OrderCreateRequest) (*order.OrderCreateResponse, error) {
	// Get task config
	taskQueueName := oa.tempoWorkflow.TaskQueueName
	taskQueueID := uuid.New().String()
	taskTimeout := oa.tempoWorkflow.TaskTimeout

	// Get workflow config
	attributes := oa.tempoWorkflow.SearchAttributes
	executionTimeout := oa.tempoWorkflow.ExecutionTimeout
	runTimeout := oa.tempoWorkflow.RunTimeout

	workflowOptions := client.StartWorkflowOptions{
		ID:               taskQueueName + "_" + taskQueueID,
		TaskQueue:        taskQueueName,
		SearchAttributes: attributes,
		WorkflowExecutionTimeout: executionTimeout,
		WorkflowRunTimeout: runTimeout,
	}

	we, err := oa.temporalClient.ExecuteWorkflow(ctx, workflowOptions, workflows.CreateOrderWorkflow, req)
	if err != nil {
		oa.logger.Error("execute workflow failed")
		return nil, err
	}

	ctxWithTimeout, cancelHandler := context.WithTimeout(context.Background(), taskTimeout)
	defer cancelHandler()

	res := &order.OrderCreateResponse{}
	err = we.Get(ctxWithTimeout, &res)
	if err != nil {
		return nil, err
	}
	oa.logger.Info(fmt.Sprintf("execute workflow ID: %v successfully", we.GetID()))
	return res, nil
}
