package payment_postgres

const (
	Create = `
		INSERT INTO payment (
			reservation_id,
			reference,
			amount,
			status,
			method
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *;
	`

	FindByID = `
		SELECT *
		FROM payment
		WHERE id = $1;
	`

	FindByReference = `
		SELECT *
		FROM payment
		WHERE reference = $1;
	`

	Update = `
		UPDATE payment
		SET
			status = $1,
			method = $2
		WHERE id = $3
		RETURNING *;
	`

	FindByReservationID = `
    SELECT *
    FROM payment
    WHERE reservation_id = $1
    LIMIT 1;
	`
)