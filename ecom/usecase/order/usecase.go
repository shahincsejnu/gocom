package order

import (
	"context"

	"github.com/shahincsejnu/gocom/ecom/domain/order"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type OrderRepository interface {
	GetList(ctx context.Context, userID string) ([]sqlcdb.Order, error)
	Create(ctx context.Context, opts *order.CreationOptions) (string, error)
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
