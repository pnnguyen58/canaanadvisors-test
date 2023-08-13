package app

import (
	"canaanadvisors-test/config"
	"canaanadvisors-test/core/workflows"
	"canaanadvisors-test/proto/management"
	"context"
	"errors"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

type Management interface {
	GetMenuOrchestration(context.Context, *management.MenuGetRequest) (*management.MenuGetResponse, error)
}


type managementApp struct {
	logger *zap.Logger
	temporalClient client.Client
	tempoWorkflow *config.Workflow
}

func NewManagement(logger *zap.Logger, cl client.Client, tcf *config.TempoConfig) Management {
	return &managementApp{
		logger: logger,
		temporalClient: cl,
		tempoWorkflow: tcf.Workflows["canaanadvisors-test-management"],
	}
}
// GetMenuOrchestration get restaurants' menu use case
func (ma *managementApp) GetMenuOrchestration(ctx context.Context, req *management.MenuGetRequest) (
	*management.MenuGetResponse, error) {
	res, err := ExecuteWorkflow[*management.MenuGetRequest, *management.MenuGetResponse](
		ctx, ma.logger, ma.temporalClient, ma.tempoWorkflow, workflows.GetMenuWorkflow, req)
	if err != nil {
		ma.logger.Error(err.Error())
		return &management.MenuGetResponse{
			Error: "get menu failed",
		}, errors.New("get menu failed")
	}
	if res == nil {
		return new(management.MenuGetResponse), nil
	}
	return res, nil
}