// ledger/usecase/get_folio_balance.go
package ledger_usecase

import (
	"context"

	ledger_domain "hotel_system2/internal/ledger/domain"
	ledger_ports "hotel_system2/internal/ledger/ports"
	shared_domain "hotel_system2/internal/shared/domain"
)

type GetFolioBalance struct {
	accountRepo     ledger_ports.AccountRepository
	transactionRepo ledger_ports.TransactionRepository
}

func NewGetFolioBalance(
	accountRepo ledger_ports.AccountRepository,
	transactionRepo ledger_ports.TransactionRepository,
) *GetFolioBalance {
	return &GetFolioBalance{accountRepo: accountRepo, transactionRepo: transactionRepo}
}

func (uc *GetFolioBalance) Execute(ctx context.Context, reservationID string) (shared_domain.Money, error) {
	receivables, err := uc.accountRepo.FindByName(ctx, GuestReceivablesAccountName)
	if err != nil {
		return shared_domain.Money{}, err
	}

	txs, err := uc.transactionRepo.FindByReference(ctx, "reservation", reservationID)
	if err != nil {
		return shared_domain.Money{}, err
	}
	paymentTxs, err := uc.transactionRepo.FindByReference(ctx, "payment", reservationID)
	if err != nil {
		return shared_domain.Money{}, err
	}
	txs = append(txs, paymentTxs...)

	var balanceMinor int64
	for _, tx := range txs {
		if tx.Status() == ledger_domain.StatusReversed {
			continue
		}
		for _, entry := range tx.Entries() {
			if entry.AccountID() != receivables.ID() {
				continue
			}
			switch entry.Type() {
			case ledger_domain.EntryDebit:
				balanceMinor += entry.Amount().AmountMinor
			case ledger_domain.EntryCredit:
				balanceMinor -= entry.Amount().AmountMinor
			}
		}
	}

	return shared_domain.Money{AmountMinor: balanceMinor, Currency: shared_domain.DefaultCurrency}, nil
}