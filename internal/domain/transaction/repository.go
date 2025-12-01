package transaction

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, tx *Transaction) error
	FindByAccount(ctx context.Context, accountID string) ([]*Transaction, error)
}
