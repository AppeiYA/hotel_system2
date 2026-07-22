package ledger_postgres

import (
	"hotel_system2/internal/shared/adapters"
	"hotel_system2/internal/shared/db"
)

type AccountRepository struct {
	adapters.Repository
}

type TransactionRepository struct {
	adapters.Repository
}

func NewAccountRepository(database *db.DB) *AccountRepository {
	return &AccountRepository{
		Repository: adapters.NewRepository(database),
	}
}

func NewTransactionRepository(database *db.DB) *TransactionRepository {
	return &TransactionRepository{
		Repository: adapters.NewRepository(database),
	}
}