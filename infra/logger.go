package infra

import (
	"context"
	"go.uber.org/zap"
)

func NewLogger(ctx context.Context) (*zap.Logger, error) {
	return zap.NewDevelopment()
}
