package domain

type TransactionStatus string

const (
	StatusPending  TransactionStatus = "pending"
	StatusPosted   TransactionStatus = "posted"
	StatusReversed TransactionStatus = "reversed"
)

func (s TransactionStatus) IsValid() bool {
	switch s {
	case StatusPending, StatusPosted, StatusReversed:
		return true
	}
	return false
}