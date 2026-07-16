package reservation_postgres

import (
	"context"
	"hotel_system2/internal/shared/db"
)

type Repository struct {
	db *db.DB
}

func NewRepository(database *db.DB) *Repository {
	return &Repository{db: database}
}

func (r *Repository) executor(ctx context.Context) db.Executor {
	return db.GetExecutor(ctx, r.db)
}