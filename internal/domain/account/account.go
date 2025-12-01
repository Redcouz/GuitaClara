package account

import (
	"errors"
	"time"

	"guitaclara/internal/domain/household"
	"guitaclara/internal/domain/shared"
)

var (
	ErrInvalidInitialBalance = errors.New("initial balance cannot be nil")
)

type ID string
type Type string

const (
	TypeChecking   Type = "checking"
	TypeSavings    Type = "savings"
	TypeCreditCard Type = "credit_card"
	TypeCash       Type = "cash"
)

// Account es la entidad raíz del agregado Account.
type Account struct {
	ID          ID
	HouseholdID household.ID
	Name        string
	Balance     shared.Money
	Type        Type
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewAccount crea una nueva cuenta. Validaciones de negocio van aquí.
func NewAccount(id ID, householdID household.ID, name string, initialBalance shared.Money, accType Type) (*Account, error) {
	if name == "" {
		return nil, errors.New("account name is required")
	}

	return &Account{
		ID:          id,
		HouseholdID: householdID,
		Name:        name,
		Balance:     initialBalance,
		Type:        accType,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// Credit aumenta el balance (Ingreso).
func (a *Account) Credit(amount shared.Money) error {
	newBalance, err := a.Balance.Add(amount)
	if err != nil {
		return err
	}
	a.Balance = newBalance
	a.UpdatedAt = time.Now()
	return nil
}

// Debit disminuye el balance (Gasto).
func (a *Account) Debit(amount shared.Money) error {
	newBalance, err := a.Balance.Subtract(amount)
	if err != nil {
		return err
	}
	a.Balance = newBalance
	a.UpdatedAt = time.Now()
	return nil
}
