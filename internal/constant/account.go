package constant

import "github.com/ffauzann/CAI-account/proto/gen"

type AccountCategory string

const (
	AccountCategoryDebit  AccountCategory = "DEBIT"
	AccountCategoryCredit AccountCategory = "CREDIT"
	AccountCategoryLoan   AccountCategory = "LOAN"
)

var CategoryGenToInternalMap = map[gen.AccountCategory]AccountCategory{
	gen.AccountCategory_AC_DEBIT:  AccountCategoryDebit,
	gen.AccountCategory_AC_CREDIT: AccountCategoryCredit,
	gen.AccountCategory_AC_LOAN:   AccountCategoryLoan,
}

var CategoryInternalToGenMap = map[AccountCategory]gen.AccountCategory{
	AccountCategoryDebit:  gen.AccountCategory_AC_DEBIT,
	AccountCategoryCredit: gen.AccountCategory_AC_CREDIT,
	AccountCategoryLoan:   gen.AccountCategory_AC_LOAN,
}
