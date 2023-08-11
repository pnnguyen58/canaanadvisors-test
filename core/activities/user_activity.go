package activities

import (
	"canaanadvisors-test/core/models"
	"canaanadvisors-test/proto/user"
	"context"
)

func Login(ctx context.Context, input *user.LoginRequest) (*user.LoginResponse, error) {
	u, err := models.GetUser(input.Username)
	if err != nil {
		return &user.LoginResponse{
			Error: "incorrect username or password",
		}, err
	}
	return &user.LoginResponse{
		Data: &user.User{
			Id: u.Id,
			Username: u.Username,
			Name: u.Name,
			RoleId: u.RoleId,
		},
	}, nil
}

func LoginCompensation(ctx context.Context, input *user.LoginRequest) (*user.LoginResponse, error) {
	return nil, nil
}

func Logout(ctx context.Context, input *user.LogoutRequest) (*user.LogoutResponse, error) {
	return nil, nil
}

func LogoutCompensation(ctx context.Context, input *user.LogoutRequest) (*user.LogoutResponse, error) {
	return nil, nil
}

