package guest_postgres

const (
	Create = `
		INSERT INTO guest (
			first_name,
			last_name,
			email,
			phone
		)
		VALUES ($1, $2, $3, $4)
		RETURNING *;
	`

	FindOrCreate = `
		INSERT INTO guest (
			first_name,
			last_name,
			email,
			phone
		)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (email)
		DO UPDATE SET
			first_name = EXCLUDED.first_name,
			last_name  = EXCLUDED.last_name,
			phone      = EXCLUDED.phone
		RETURNING *;
	`

	ExistsByEmail = `
		SELECT EXISTS(
			SELECT 1
			FROM guest
			WHERE email = $1
		);
	`

	FindByEmail = `
		SELECT *
		FROM guest
		WHERE email = $1;
	`
	FindGuestByID = `
		SELECT *
		FROM guest
		WHERE id = $1;
	`
)