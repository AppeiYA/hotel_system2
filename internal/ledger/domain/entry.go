package domain

import (
	shared_domain "hotel_system2/internal/shared/domain"
)

type Entry struct {
	accountID string
	entryType EntryType
	amount    shared_domain.Money
}

func NewEntry(accountID string, entryType EntryType, amount shared_domain.Money) (Entry, error) {
	if accountID == "" {
		return Entry{}, ErrAccountIDRequired
	}
	if entryType != EntryDebit && entryType != EntryCredit {
		return Entry{}, ErrInvalidEntryType
	}
	if amount.AmountMinor <= 0 {
		return Entry{}, ErrNonPositiveAmount
	}
	return Entry{accountID: accountID, entryType: entryType, amount: amount}, nil
}

func (e Entry) AccountID() string             { return e.accountID }
func (e Entry) Type() EntryType                { return e.entryType }
func (e Entry) Amount() shared_domain.Money     { return e.amount }