package reservation_usecase

import (
	"context"

	reservation_ports "hotel_system2/internal/reservation/ports"
	room_ports "hotel_system2/internal/room/ports"
	"hotel_system2/internal/shared/db"
    reservation_domain "hotel_system2/internal/reservation/domain"
)

type CheckOut struct {
	txManager       *db.TransactionManager
	reservationRepo reservation_ports.ReservationRepository
	roomRepo        room_ports.RoomRepository
    ledger          reservation_ports.LedgerPort
}

func NewCheckOut(
	txManager *db.TransactionManager,
	reservationRepo reservation_ports.ReservationRepository,
	roomRepo room_ports.RoomRepository,
    ledger reservation_ports.LedgerPort,
) *CheckOut {
	return &CheckOut{
		txManager:       txManager,
		reservationRepo: reservationRepo,
		roomRepo:        roomRepo,
        ledger:          ledger,
	}
}

func (uc *CheckOut) Execute(ctx context.Context, reservationID string) error {
	balance, err := uc.ledger.GetOutstandingBalance(ctx, reservationID)
	if err != nil {
		return err
	}
	if balance.AmountMinor > 0 {
		return reservation_domain.ErrOutstandingBalance
	}

	return uc.txManager.WithinTransaction(ctx, func(ctx context.Context) error {
		reservation, err := uc.reservationRepo.FindByID(ctx, reservationID)
		if err != nil {
			return err
		}
		if err := reservation.CheckOut(); err != nil {
			return err
		}

		room, err := uc.roomRepo.FindByID(ctx, reservation.RoomID())
		if err != nil {
			return err
		}
		if err := room.MarkForCleaning(); err != nil {
			return err
		}

		if err := uc.reservationRepo.Update(ctx, reservation); err != nil {
			return err
		}
		return uc.roomRepo.Update(ctx, room)
	})
}