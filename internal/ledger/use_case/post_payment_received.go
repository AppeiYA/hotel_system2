// ledger/usecase/post_payment_received.go
package ledger_usecase

import (
	"context"

	"hotel_system2/internal/ledger/domain"
	ledger_ports "hotel_system2/internal/ledger/ports"
	shared_domain "hotel_system2/internal/shared/domain"

	"github.com/google/uuid"
)

const CashAccountName = "Cash"
var transactionDescriptionPayment = "Payment received for reservation"

type PostPaymentReceived struct {
	accountRepo     ledger_ports.AccountRepository
	transactionRepo ledger_ports.TransactionRepository
}

func NewPostPaymentReceived(
	accountRepo ledger_ports.AccountRepository,
	transactionRepo ledger_ports.TransactionRepository,
) *PostPaymentReceived {
	return &PostPaymentReceived{accountRepo: accountRepo, transactionRepo: transactionRepo}
}

func (uc *PostPaymentReceived) Execute(
	ctx context.Context,
	reservationID string,
	amount shared_domain.Money,
) (*domain.Transaction, error) {

	cash, err := uc.accountRepo.FindByName(ctx, CashAccountName)
	if err != nil {
		return nil, err
	}
	receivables, err := uc.accountRepo.FindByName(ctx, GuestReceivablesAccountName)
	if err != nil {
		return nil, err
	}

	debit, err := domain.NewEntry(cash.ID(), domain.EntryDebit, amount)
	if err != nil {
		return nil, err
	}
	credit, err := domain.NewEntry(receivables.ID(), domain.EntryCredit, amount)
	if err != nil {
		return nil, err
	}

	tx, err := domain.NewTransaction(
		uuid.New().String(),
		string(shared_domain.ReferenceTypePayment),
		reservationID,
		transactionDescriptionPayment,
		[]domain.Entry{debit, credit},
	)
	if err != nil {
		return nil, err
	}

	if err := uc.transactionRepo.Create(ctx, tx); err != nil {
		return nil, err
	}
	return tx, nil
}