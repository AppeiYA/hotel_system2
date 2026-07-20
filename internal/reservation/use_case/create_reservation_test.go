package reservation_usecase

import (
	"context"
	// guest_domain "hotel_system2/internal/guest/domain"
	reservation_domain "hotel_system2/internal/reservation/domain"
	"github.com/stretchr/testify/require"
	test_mocks "hotel_system2/internal/tests/mocks"
	"testing"
)

func TestCreateReservation_Success(t *testing.T) {

	repo := &test_mocks.MockReservationRepository{
		CreateFn: func(ctx context.Context, r *reservation_domain.Reservation) error {
			return nil
		},
	}

	uc := NewCreateReservation(
		&test_mocks.MockTransactionManager{},
		repo,
		&test_mocks.MockRoomRepository{},
		&test_mocks.MockGuestRepository{},
		&test_mocks.MockPaymentLookupPort{},
	)

	request := CreateReservationInput{}

	_, err := uc.Execute(context.Background(), request)

	require.NoError(t, err)
}
