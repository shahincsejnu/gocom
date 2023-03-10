package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/shahincsejnu/gocom/ecom/domain/order"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type Repository struct {
	DB *sqlcdb.Queries
}

func NewRepository(db *sqlcdb.Queries) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetOne(ctx context.Context, orderID string) (*sqlcdb.Order, error) {
	ord, err := r.DB.GetOrderById(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return &ord, nil
}

func (r *Repository) Create(ctx context.Context, opts *order.CreationOptions) (string, error) {
	arg := sqlcdb.CreateOrderParams{
		ID:        uuid.NewString(),
		ProductID: opts.ProductID,
		UserID:    opts.UserID,
		AddressID: opts.AddressID,
		Quantity:  int32(opts.Quantity),
	}
	err := r.DB.CreateOrder(ctx, arg)
	if err != nil {
		return "", err
	}

	return arg.ID, nil
}

func (r *Repository) UpdateOne(ctx context.Context, opts *order.UpdateOptions, orderID string) error {
	arg := sqlcdb.UpdateOrderParams{
		ID:       orderID,
		Quantity: int32(opts.Quantity),
	}
	return r.DB.UpdateOrder(ctx, arg)
}

func (r *Repository) DeleteOne(ctx context.Context, orderID string) error {
	return r.DB.DeleteOrder(ctx, orderID)
}
