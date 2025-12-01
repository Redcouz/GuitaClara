package transaction

import (
	"errors"
	"time"

	"guitaclara/internal/domain/account"
	"guitaclara/internal/domain/shared"
)

type ID string
type Category string
type Type string

const (
	TypeIncome  Type = "income"
	TypeExpense Type = "expense"
)

type Transaction struct {
	ID          ID
	AccountID   account.ID
	Amount      shared.Money
	Category    Category
	Description string
	Date        time.Time
	Type        Type
	CreatedAt   time.Time
}

func NewTransaction(
	id ID,
	accountID account.ID,
	amount shared.Money,
	category Category,
	desc string,
	txType Type,
) (*Transaction, error) {
	if amount.Amount <= 0 {
		return nil, errors.New("transaction amount must be positive")
	}

	return &Transaction{
		ID:          id,
		AccountID:   accountID,
		Amount:      amount,
		Category:    category,
		Description: desc,
		Date:        time.Now(),
		Type:        txType,
		CreatedAt:   time.Now(),
	}, nil
}
