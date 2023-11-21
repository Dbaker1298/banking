package domain

import (
	"strconv"

	"github.com/Dbaker1298/banking/app/errs"
	"github.com/Dbaker1298/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}
