package domain

import "time"

type Transaction struct {
	id            string
	referenceType string
	referenceID   string
	description   string
	status        TransactionStatus
	entries       []Entry
	createdAt     time.Time
}

func NewTransaction(
	id, referenceType, referenceID, description string,
	entries []Entry,
) (*Transaction, error) {
	if len(entries) < 2 {
		return nil, ErrTooFewEntries
	}

	var debitTotal, creditTotal int64
	currency := entries[0].Amount().Currency
	for _, e := range entries {
		if e.Amount().Currency != currency {
			return nil, ErrInvalidCurrency
		}
		switch e.Type() {
		case EntryDebit:
			debitTotal += e.Amount().AmountMinor
		case EntryCredit:
			creditTotal += e.Amount().AmountMinor
		}
	}
	if debitTotal != creditTotal {
		return nil, ErrUnbalancedTransaction
	}

	return &Transaction{
		id:            id,
		referenceType: referenceType,
		referenceID:   referenceID,
		description:   description,
		status:        StatusPosted, // posted immediately; this system has no draft workflow
		entries:       entries,
		createdAt:     time.Now(),
	}, nil
}

func ReconstituteTransaction(
	id, referenceType, referenceID, description string,
	status TransactionStatus,
	entries []Entry,
	createdAt time.Time,
) *Transaction {
	return &Transaction{
		id: id, referenceType: referenceType, referenceID: referenceID,
		description: description, status: status, entries: entries, createdAt: createdAt,
	}
}

func (t *Transaction) ID() string                    { return t.id }
func (t *Transaction) ReferenceType() string          { return t.referenceType }
func (t *Transaction) ReferenceID() string             { return t.referenceID }
func (t *Transaction) Description() string             { return t.description }
func (t *Transaction) Status() TransactionStatus        { return t.status }
func (t *Transaction) Entries() []Entry                 { return t.entries }
func (t *Transaction) CreatedAt() time.Time              { return t.createdAt }

func (t *Transaction) MarkReversed() error {
	if t.status == StatusReversed {
		return ErrAlreadyReversed
	}
	if t.status != StatusPosted {
		return ErrCannotReversePending
	}
	t.status = StatusReversed
	return nil
}
func (t *Transaction) ReversingEntries() ([]Entry, error) {
	reversed := make([]Entry, 0, len(t.entries))
	for _, e := range t.entries {
		flippedType := EntryCredit
		if e.Type() == EntryCredit {
			flippedType = EntryDebit
		}
		entry, err := NewEntry(e.AccountID(), flippedType, e.Amount())
		if err != nil {
			return nil, err
		}
		reversed = append(reversed, entry)
	}
	return reversed, nil
}
