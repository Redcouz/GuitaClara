package finance

import (
	"context"
	"errors"
	"fmt"

	"guitaclara/internal/domain/account"
	"guitaclara/internal/domain/shared"
	"guitaclara/internal/domain/transaction"
)

// UnitOfWork define una interfaz para manejar transacciones de base de datos atómicas.
// Esto se implementará en infraestructura.
type UnitOfWork interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}

type Service struct {
	accountRepo     account.Repository
	transactionRepo transaction.Repository
	uow             UnitOfWork
}

func NewService(ar account.Repository, tr transaction.Repository, uow UnitOfWork) *Service {
	return &Service{
		accountRepo:     ar,
		transactionRepo: tr,
		uow:             uow,
	}
}

type CreateTransactionDTO struct {
	ID          string
	AccountID   string
	Amount      int64 // En centavos
	Currency    string
	Category    string
	Description string
	Type        string // "income" o "expense"
}

// RecordTransaction orquesta la creación de la transacción y la actualización del balance.
func (s *Service) RecordTransaction(ctx context.Context, dto CreateTransactionDTO) error {
	// 1. Validar y crear Value Objects
	money := shared.NewMoney(dto.Amount, dto.Currency)
	accID := account.ID(dto.AccountID)
	txID := transaction.ID(dto.ID)

	// Ejecutamos dentro de una transacción de base de datos (Unit of Work)
	return s.uow.Do(ctx, func(txCtx context.Context) error {
		// 2. Obtener la cuenta (bloqueo pesimista puede ser necesario en DB real)
		acc, err := s.accountRepo.FindByID(txCtx, accID)
		if err != nil {
			return fmt.Errorf("failed to fetch account: %w", err)
		}
		if acc == nil {
			return errors.New("account not found")
		}

		// 3. Aplicar lógica de dominio
		// Validamos si la moneda coincide
		if acc.Balance.Currency != money.Currency {
			return fmt.Errorf("currency mismatch: account is %s, tx is %s", acc.Balance.Currency, money.Currency)
		}

		txType := transaction.Type(dto.Type)

		// Actualizar balance en memoria (la entidad protege su invariante)
		if txType == transaction.TypeExpense {
			if err := acc.Debit(money); err != nil {
				return err
			}
		} else if txType == transaction.TypeIncome {
			if err := acc.Credit(money); err != nil {
				return err
			}
		} else {
			return errors.New("invalid transaction type")
		}

		// 4. Crear entidad Transacción
		newTx, err := transaction.NewTransaction(
			txID,
			accID,
			money,
			transaction.Category(dto.Category),
			dto.Description,
			txType,
		)
		if err != nil {
			return err
		}

		// 5. Persistir cambios
		if err := s.transactionRepo.Save(txCtx, newTx); err != nil {
			return err
		}

		// Guardamos el nuevo balance
		if err := s.accountRepo.Save(txCtx, acc); err != nil {
			return err
		}

		return nil
	})
}
