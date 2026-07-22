package ledger_postgres

const (
	FindAccountByID = `
		SELECT id, name, type, currency, created_at
		FROM ledger_account WHERE id = $1`

	FindAccountByName = `
		SELECT id, name, type, currency, created_at
		FROM ledger_account WHERE name = $1`

	InsertTransaction = `
		INSERT INTO ledger_transaction (id, reference_type, reference_id, description, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, reference_type, reference_id, description, status, created_at`

	InsertEntry = `
		INSERT INTO ledger_entry (id, transaction_id, account_id, entry_type, amount)
		VALUES ($1, $2, $3, $4, $5)`

	FindTransactionByID = `
		SELECT id, reference_type, reference_id, description, status, created_at
		FROM ledger_transaction WHERE id = $1`

	FindEntriesByTransactionID = `
		SELECT id, transaction_id, account_id, entry_type, amount
		FROM ledger_entry WHERE transaction_id = $1`

	FindTransactionsByReference = `
		SELECT id, reference_type, reference_id, description, status, created_at
		FROM ledger_transaction
		WHERE reference_type = $1 AND reference_id = $2
		ORDER BY created_at ASC`

	FindEntriesByTransactionIDs = `
		SELECT id, transaction_id, account_id, entry_type, amount
		FROM ledger_entry WHERE transaction_id = ANY($1)`

	UpdateTransactionStatus = `
		UPDATE ledger_transaction SET status = $1 WHERE id = $2`
)