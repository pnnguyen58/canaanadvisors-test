package app

import (
	"canaanadvisors-test/config"
	"canaanadvisors-test/core/workflows"
	"canaanadvisors-test/proto/user"
	"context"
	"errors"
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

func NewUser(logger *zap.Logger, cl client.Client, tcf *config.TempoConfig) User {
	return &userApp{
		logger: logger,
		temporalClient: cl,
		tempoWorkflow: tcf.Workflows["canaanadvisors-test-user"],
	}
}

// LoginOrchestration login use case
func (ua *userApp) LoginOrchestration(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	res, err := ExecuteWorkflow[*user.LoginRequest, *user.LoginResponse](
		ctx, ua.logger, ua.temporalClient, ua.tempoWorkflow, workflows.LoginWorkflow, req)
	if err != nil {
		ua.logger.Error(err.Error())
		return nil, errors.New("login failed")
	}
	if res == nil {
		return new(user.LoginResponse), nil
	}
	return res, nil
}

// LogoutOrchestration logout use case
func (ua *userApp) LogoutOrchestration(ctx context.Context, req *user.LogoutRequest) (*user.LogoutResponse, error) {
	res, err := ExecuteWorkflow[*user.LogoutRequest, *user.LogoutResponse](
		ctx, ua.logger, ua.temporalClient, ua.tempoWorkflow, workflows.LogoutWorkflow, req)
	if err != nil {
		ua.logger.Error(err.Error())
		return nil, errors.New("logout failed")
	}
	if res == nil {
		return new(user.LogoutResponse), nil
	}
	return res, nil
}
