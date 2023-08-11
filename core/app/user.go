package app

import (
	"canaanadvisors-test/config"
	"canaanadvisors-test/core/workflows"
	"canaanadvisors-test/proto/user"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

type User interface {
	LoginOrchestration(context.Context, *user.LoginRequest) (*user.LoginResponse, error)
	LogoutOrchestration(context.Context, *user.LogoutRequest) (*user.LogoutResponse, error)
}

type userApp struct {
	logger *zap.Logger
	temporalClient client.Client
	tempoWorkflow *config.Workflow
}

func NewUserApp(ctx context.Context, logger *zap.Logger, cl client.Client, tcf *config.TempoConfig) User {
	return &userApp{
		logger: logger,
		temporalClient: cl,
		tempoWorkflow: tcf.Workflows["canaanadvisors-test-user"],
	}
}

// LoginOrchestration login use case
func (aa *userApp) LoginOrchestration(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	// Get task config
	taskQueueName := aa.tempoWorkflow.TaskQueueName
	taskQueueID := uuid.New().String()
	taskTimeout := aa.tempoWorkflow.TaskTimeout

	// Get workflow config
	attributes := aa.tempoWorkflow.SearchAttributes
	executionTimeout := aa.tempoWorkflow.ExecutionTimeout
	runTimeout := aa.tempoWorkflow.RunTimeout

	workflowOptions := client.StartWorkflowOptions{
		ID:               taskQueueName + "_" + taskQueueID,
		TaskQueue:        taskQueueName,
		SearchAttributes: attributes,
		WorkflowExecutionTimeout: executionTimeout,
		WorkflowRunTimeout: runTimeout,
	}

	we, err := aa.temporalClient.ExecuteWorkflow(ctx, workflowOptions, workflows.LoginWorkflow, req)
	if err != nil {
		aa.logger.Error("execute workflow failed")
		return nil, err
	}

	ctxWithTimeout, cancelHandler := context.WithTimeout(context.Background(), taskTimeout)
	defer cancelHandler()

	res := &user.LoginResponse{}
	err = we.Get(ctxWithTimeout, &res)
	if err != nil {
		return nil, err
	}
	aa.logger.Info(fmt.Sprintf("execute workflow ID: %v successfully", we.GetID()))
	return res, nil
}

// LogoutOrchestration logout use case
func (aa *userApp) LogoutOrchestration(ctx context.Context, req *user.LogoutRequest) (*user.LogoutResponse, error) {
	// Get task config
	taskQueueName := aa.tempoWorkflow.TaskQueueName
	taskQueueID := uuid.New().String()
	taskTimeout := aa.tempoWorkflow.TaskTimeout

	// Get workflow config
	attributes := aa.tempoWorkflow.SearchAttributes
	executionTimeout := aa.tempoWorkflow.ExecutionTimeout
	runTimeout := aa.tempoWorkflow.RunTimeout

	workflowOptions := client.StartWorkflowOptions{
		ID:               taskQueueName + "_" + taskQueueID,
		TaskQueue:        taskQueueName,
		SearchAttributes: attributes,
		WorkflowExecutionTimeout: executionTimeout,
		WorkflowRunTimeout: runTimeout,
	}

	we, err := aa.temporalClient.ExecuteWorkflow(ctx, workflowOptions, workflows.LogoutWorkflow, req)
	if err != nil {
		aa.logger.Error("execute workflow failed")
		return nil, err
	}

	ctxWithTimeout, cancelHandler := context.WithTimeout(context.Background(), taskTimeout)
	defer cancelHandler()

	res := &user.LogoutResponse{}
	err = we.Get(ctxWithTimeout, &res)
	if err != nil {
		return nil, err
	}
	aa.logger.Info(fmt.Sprintf("execute workflow ID: %v successfully", we.GetID()))
	return res, nil
}
