package domain

type Account struct {
	id       string
	name     string
	accType  AccountType
}

func NewAccount(id, name string, accType AccountType) (*Account, error) {
	if name == "" {
		return nil, ErrAccountNameRequired
	}
	if !accType.IsValid() {
		return nil, ErrInvalidAccountType
	}
	return &Account{id: id, name: name, accType: accType}, nil
}

func ReconstituteAccount(id, name string, accType AccountType) *Account {
	return &Account{id: id, name: name, accType: accType}
}

func (a *Account) ID() string {
	return a.id
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) Type() AccountType {
	return a.accType
}