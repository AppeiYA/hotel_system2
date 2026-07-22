package domain

import custom_errors "hotel_system2/internal/shared/errors"

var (
	ErrInvalidAccountType = custom_errors.InternalServerError("invalid account type")
	ErrAccountNameRequired = custom_errors.BadException("account name is required")
	ErrInvalidEntryType   = custom_errors.InternalServerError("invalid ledger entry type")
	ErrNonPositiveAmount  = custom_errors.BadException("ledger entry amount must be positive")
	ErrAccountIDRequired  = custom_errors.BadException("account id is required")
	ErrUnbalancedTransaction = custom_errors.BadException("ledger transaction debits and credits must balance")
	ErrTooFewEntries         = custom_errors.BadException("ledger transaction requires at least two entries")
	ErrAlreadyReversed       = custom_errors.BadException("ledger transaction is already reversed")
	ErrCannotReversePending  = custom_errors.BadException("only posted transactions can be reversed")
	ErrInvalidCurrency = custom_errors.ConflictError("all entries in a transaction must share one currency")
)