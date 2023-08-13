package infra

import (
	"canaanadvisors-test/config"
	"context"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

func NewTemporalClient(ctx context.Context, logger *zap.Logger, tc *config.TempoConfig) (client.Client, error) {
	cl, err := client.Dial(client.Options{
		HostPort: tc.HostPort,
		Namespace: tc.Namespace.Name,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	namespace, err := cl.WorkflowService().DescribeNamespace(ctx, &workflowservice.DescribeNamespaceRequest{
		Namespace: tc.Namespace.Name,
	})
	if namespace != nil && err == nil {
		return cl, nil
	}
	_, err = cl.WorkflowService().RegisterNamespace(ctx, &workflowservice.RegisterNamespaceRequest{
		Namespace:                        tc.Namespace.Name,
		WorkflowExecutionRetentionPeriod: &tc.Namespace.WorkflowExecutionRetentionPeriod,
	})
	return cl, err
}
