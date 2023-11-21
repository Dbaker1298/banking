package domain

import (
	"github.com/Dbaker1298/banking/app/errs"
	"github.com/Dbaker1298/banking/dto"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
