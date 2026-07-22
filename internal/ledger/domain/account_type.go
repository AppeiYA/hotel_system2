package domain

type AccountType string

const (
	AccountAsset     AccountType = "ASSET"
	AccountLiability AccountType = "LIABILITY"
	AccountRevenue   AccountType = "REVENUE"
	AccountExpense   AccountType = "EXPENSE"
)

func (t AccountType) IsValid() bool {
	switch t {
	case AccountAsset, AccountLiability, AccountRevenue, AccountExpense:
		return true
	}
	return false
}