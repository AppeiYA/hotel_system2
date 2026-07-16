package reservation_postgres

const (
	CreateReservation = `
		INSERT INTO reservation (
			guest_id,
			room_id,
			check_in,
			check_out,
			total_amount,
			status
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *;
	`

	HasOverlap = `
		SELECT EXISTS (
			SELECT 1
			FROM reservation
			WHERE room_id = $1
			AND status NOT IN ('cancelled', 'checked_out')
			AND check_in < $3
			AND check_out > $2
		);
	`

	GetReservationDetails = `
		SELECT
			r.id,
			r.guest_id,
			r.room_id,
			r.check_in,
			r.check_out,
			r.total_amount,
			r.status,
			r.created_at,

			p.id      AS payment_id,
			p.method  AS payment_method,
			p.status  AS payment_status,
			p.amount  AS payment_amount

		FROM reservation r
		LEFT JOIN payment p
			ON p.reservation_id = r.id
		WHERE r.id = $1;
	`

	ListByEmail = `
		SELECT r.*
		FROM reservation r
		JOIN guest g
			ON g.id = r.guest_id
		WHERE g.email = $1
		ORDER BY r.created_at DESC;
	`

	FindByIDForUpdate = `
		SELECT *
		FROM reservation
		WHERE id = $1
		FOR UPDATE;
	`

	UpdateStatus = `
		UPDATE reservation
		SET status = $1
		WHERE id = $2
		RETURNING *;
	`

	FindExpiredPending = `
		SELECT *
		FROM reservation
		WHERE status = 'pending'
		AND created_at < $1
		FOR UPDATE SKIP LOCKED;
	`

	FindNoShow = `
		SELECT *
		FROM reservation
		WHERE status = 'confirmed'
		AND check_out < $1
		FOR UPDATE SKIP LOCKED;
	`
	FindByID = `
		SELECT *
		FROM reservation
		WHERE id = $1;
	`

	ListReservations = `
		SELECT *
		FROM reservation
		ORDER BY created_at DESC;
	`

	UpdateReservation = `
		UPDATE reservation
		SET
			guest_id = $1,
			room_id = $2,
			check_in = $3,
			check_out = $4,
			total_amount = $5,
			status = $6
		WHERE id = $7
		RETURNING *;
	`
)