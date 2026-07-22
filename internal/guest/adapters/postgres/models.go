package guest_postgres

import (
	"hotel_system2/internal/guest/domain"
	"time"

	"github.com/google/uuid"
)

type guestRow struct {
	ID        uuid.UUID    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
}

func (g guestRow) toDomain() (*domain.Guest, error) {
	return domain.ReconstituteGuest(
		g.ID.String(),
		g.FirstName,
		g.LastName,
		g.Email,
		g.Phone,
		g.CreatedAt,
	)
}

func guestRowFromDomain(guest *domain.Guest) guestRow {
	return guestRow{
		ID:        uuid.MustParse(guest.ID()),
		FirstName: guest.FirstName(),
		LastName:  guest.LastName(),
		Email:     guest.Email().String(),
		Phone:     guest.Phone(),
		CreatedAt: guest.CreatedAt(),
	}
}