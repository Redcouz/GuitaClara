package account

import (
	"context"
)

// Repository define el contrato para la persistencia de cuentas.
// Se implementará en infrastructure/persistence.
type Repository interface {
	Save(ctx context.Context, account *Account) error
	FindByID(ctx context.Context, id ID) (*Account, error)
	FindByHousehold(ctx context.Context, householdID string) ([]*Account, error)
	UpdateBalance(ctx context.Context, id ID, newBalance int64) error
}
