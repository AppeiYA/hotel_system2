package room_postgres

import (
	"context"
	"hotel_system2/internal/shared/db"
)

const (
	CREATE_ROOM = `
		INSERT INTO room (room_number, room_type, rate, status)
		VALUES ($1, $2, $3, $4)
		RETURNING *;
	`
	GET_ROOM_BY_ID = `
		SELECT * FROM room
		WHERE id = $1;
	`

	GET_ROOM_BY_NUMBER = `
		SELECT * FROM room
		WHERE room_number = $1;
	`

	LIST_ROOMS = `
		SELECT * FROM room
		ORDER BY created_at DESC;
	`

	UPDATE_ROOM = `
		UPDATE room
		SET room_number = $1,
			room_type = $2,
			rate = $3,
			status = $4
		WHERE id = $5
		RETURNING *;
	`

	DELETE_ROOM = `
		DELETE FROM room
		WHERE id = $1;
	`

	UPDATE_ROOM_STATUS = `
		UPDATE room
		SET status = $1
		WHERE id = $2
		RETURNING *;
	`

	GET_ROOM_BY_ID_FOR_UPDATE = `
		SELECT * FROM room
		WHERE id = $1
		FOR UPDATE;
	`
	FIND_AVAILABLE_ROOM = `
	SELECT r.* FROM room r
	WHERE r.room_type = $1
	  AND r.status = 'available'
	  AND r.id NOT IN (
	      SELECT room_id FROM reservation
	      WHERE status NOT IN ('cancelled', 'checked_out')
	        AND (check_in, check_out) OVERLAPS ($2, $3)
	  )
	LIMIT 1;
	`
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