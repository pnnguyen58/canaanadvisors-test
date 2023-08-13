package app

import (
	"canaanadvisors-test/config"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

func ExecuteWorkflow[Req any, Res any](ctx context.Context, logger *zap.Logger, cl client.Client, cfg *config.Workflow,
	workflow interface{}, req Req, args ...any) (Res, error) {
	var res Res
	if cl == nil {
		return res, errors.New("not found client")
	}
	if cfg == nil {
		return res, errors.New("empty config")
	}
	if workflow == nil {
		return res, errors.New("invalid workflow")
	}
	// Get task config
	taskQueueName := cfg.TaskQueueName
	taskQueueID := uuid.New().String()
	taskTimeout := cfg.TaskTimeout

	// Get workflow config
	attributes := cfg.SearchAttributes
	executionTimeout := cfg.ExecutionTimeout
	runTimeout := cfg.RunTimeout

	workflowOptions := client.StartWorkflowOptions{
		ID:               taskQueueName + "_" + taskQueueID,
		TaskQueue:        taskQueueName,
		SearchAttributes: attributes,
		WorkflowExecutionTimeout: executionTimeout,
		WorkflowRunTimeout: runTimeout,
	}

	we, err := cl.ExecuteWorkflow(ctx, workflowOptions, workflow, req, args)
	if err != nil {
		logger.Error("execute workflow failed")
		return res, err
	}

	ctxWithTimeout, cancelHandler := context.WithTimeout(context.Background(), taskTimeout)
	defer cancelHandler()
	err = we.Get(ctxWithTimeout, &res)
	if err != nil {
		return res, err
	}
	logger.Info(fmt.Sprintf("execute workflow ID: %v successfully", we.GetID()))
	return res, nil
}