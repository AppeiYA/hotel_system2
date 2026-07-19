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
	Email     domain.Email `db:"email"`
	Phone     string `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
}

func (g *guestRow) toDomain() *domain.Guest {
	return &domain.Guest{
		ID:        g.ID.String(),
		FirstName: g.FirstName,
		LastName:  g.LastName,
		Email:     g.Email,
		Phone:     g.Phone,
		CreatedAt: g.CreatedAt,
	}
}