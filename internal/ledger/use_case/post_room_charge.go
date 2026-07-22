package ledger_usecase

import (
	"context"

	"hotel_system2/internal/ledger/domain"
	ledger_ports "hotel_system2/internal/ledger/ports"
	shared_domain "hotel_system2/internal/shared/domain"

	"github.com/google/uuid"
)

const (
	GuestReceivablesAccountName = "Guest Receivables"
	RoomRevenueAccountName      = "Room Revenue"
)

var transactionDescriptionRoomCharge = "Room charge for reservation"

type PostRoomCharge struct {
	accountRepo     ledger_ports.AccountRepository
	transactionRepo ledger_ports.TransactionRepository
}

func NewPostRoomCharge(
	accountRepo ledger_ports.AccountRepository,
	transactionRepo ledger_ports.TransactionRepository,
) *PostRoomCharge {
	return &PostRoomCharge{accountRepo: accountRepo, transactionRepo: transactionRepo}
}

func (uc *PostRoomCharge) Execute(
	ctx context.Context,
	reservationID string,
	amount shared_domain.Money,
) (*domain.Transaction, error) {

	receivables, err := uc.accountRepo.FindByName(ctx, GuestReceivablesAccountName)
	if err != nil {
		return nil, err
	}
	revenue, err := uc.accountRepo.FindByName(ctx, RoomRevenueAccountName)
	if err != nil {
		return nil, err
	}

	debit, err := domain.NewEntry(receivables.ID(), domain.EntryDebit, amount)
	if err != nil {
		return nil, err
	}
	credit, err := domain.NewEntry(revenue.ID(), domain.EntryCredit, amount)
	if err != nil {
		return nil, err
	}

	tx, err := domain.NewTransaction(
		uuid.New().String(),
		string(shared_domain.ReferenceTypeReservation),
		reservationID,
		transactionDescriptionRoomCharge,
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