package repositories

import (
	"context"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"canaanadvisors-test/core/models"
)

type Order interface {
	Insert(context.Context, models.Order) (models.Order, error)
}

type orderRepository struct {
	logger *zap.Logger
	db *gorm.DB
}

func NewOrderRepository(logger *zap.Logger, db *gorm.DB) Order {
	return &orderRepository{
		logger: logger,
		db: db,
	}
}

func (or *orderRepository) Insert(ctx context.Context, order models.Order) (models.Order, error) {
	err := or.db.Clauses(order.InsertClause()...).
		Model(models.Order{}).
		WithContext(ctx).
		Create(&order).Error
	return order, err
}

var OrderRepositorySet = wire.NewSet(
	wire.Bind(new(Order), new(*orderRepository)),
)