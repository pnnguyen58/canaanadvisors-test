package main

import (
	"canaanadvisors-test/core/app"
	"canaanadvisors-test/proto/user"
	"context"
	"go.uber.org/zap"
)

type UserHandler interface {
	Login(context.Context, *user.LoginRequest) (*user.LoginResponse, error)
	Logout(context.Context, *user.LogoutRequest) (*user.LogoutResponse, error)
}

func NewUserHandler(logger *zap.Logger, app app.User) UserHandler {
	return &UserController{logger: logger, app: app}
}

type UserController struct {
	user.UnimplementedUserServiceServer
	logger *zap.Logger
	app app.User
}

func (ac *UserController) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	return ac.app.LoginOrchestration(ctx, req)
}

func (ac *UserController) Logout(ctx context.Context, req *user.LogoutRequest) (*user.LogoutResponse, error) {
	return ac.app.LogoutOrchestration(ctx, req)
}
