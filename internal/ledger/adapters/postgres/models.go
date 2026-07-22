package ledger_postgres

import (
	"time"

	"hotel_system2/internal/ledger/domain"
	shared_domain "hotel_system2/internal/shared/domain"
)

type accountRow struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Type      string    `db:"type"`
	Currency  string    `db:"currency"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *accountRow) toDomain() *domain.Account {
	if r == nil {
		return nil
	}
	return domain.ReconstituteAccount(r.ID, r.Name, domain.AccountType(r.Type))
}

type transactionRow struct {
	ID            string    `db:"id"`
	ReferenceType string    `db:"reference_type"`
	ReferenceID   string    `db:"reference_id"`
	Description   string    `db:"description"`
	Status        string    `db:"status"`
	CreatedAt     time.Time `db:"created_at"`
}

type entryRow struct {
	ID            string `db:"id"`
	TransactionID string `db:"transaction_id"`
	AccountID     string `db:"account_id"`
	EntryType     string `db:"entry_type"`
	Amount        int64  `db:"amount"`
}

func assembleTransaction(txRow transactionRow, entryRows []entryRow) (*domain.Transaction, error) {
	entries := make([]domain.Entry, 0, len(entryRows))
	for _, er := range entryRows {
		amount := shared_domain.Money{AmountMinor: er.Amount, Currency: shared_domain.DefaultCurrency}
		entry, err := domain.NewEntry(er.AccountID, domain.EntryType(er.EntryType), amount)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return domain.ReconstituteTransaction(
		txRow.ID,
		txRow.ReferenceType,
		txRow.ReferenceID,
		txRow.Description,
		domain.TransactionStatus(txRow.Status),
		entries,
		txRow.CreatedAt,
	), nil
}