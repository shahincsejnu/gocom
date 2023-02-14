package order

import (
	"context"

	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type OrderRepository interface {
	GetList(ctx context.Context, userID string) ([]sqlcdb.Order, error)
}

type Usecase struct {
	OrderRepo OrderRepository
}

func (uc *Usecase) GetOrdersList(ctx context.Context, userID string) ([]sqlcdb.Order, error) {
	return uc.OrderRepo.GetList(ctx, userID)
}
