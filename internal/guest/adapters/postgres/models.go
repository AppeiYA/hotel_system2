package guest_postgres

import (
	"hotel_system2/internal/guest/domain"
	"time"

	"github.com/google/uuid"
	shared_domain "hotel_system2/internal/shared/domain"
)

type guestRow struct {
	ID        uuid.UUID    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     shared_domain.Email `db:"email"`
	Phone     string `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
}

func (g guestRow) toDomain() *domain.Guest {
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
		Email:     guest.Email(),
		Phone:     guest.Phone(),
		CreatedAt: guest.CreatedAt(),
	}
}