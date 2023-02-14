package order

import (
	"context"

	"github.com/shahincsejnu/gocom/ecom/domain/order"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type OrderRepository interface {
	GetList(ctx context.Context, userID string) ([]sqlcdb.Order, error)
	GetOne(ctx context.Context, orderID string) (*sqlcdb.Order, error)
	Create(ctx context.Context, opts *order.CreationOptions) (string, error)
	UpdateOne(ctx context.Context, opts *order.UpdateOptions, orderID string) error
}

type Usecase struct {
	OrderRepo OrderRepository
}

func (uc *Usecase) GetOrdersList(ctx context.Context, userID string) ([]sqlcdb.Order, error) {
	return uc.OrderRepo.GetList(ctx, userID)
}

func (uc *Usecase) CreateOrder(ctx context.Context, opts *order.CreationOptions) (string, error) {
	return uc.OrderRepo.Create(ctx, opts)
}

func (uc *Usecase) GetOneOrder(ctx context.Context, orderID string) (*sqlcdb.Order, error) {
	return uc.OrderRepo.GetOne(ctx, orderID)
}

func (uc *Usecase) UpdateOneOrder(ctx context.Context, opts *order.UpdateOptions, orderID string) error {
	return uc.OrderRepo.UpdateOne(ctx, opts, orderID)
}
