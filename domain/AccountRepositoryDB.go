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

/**
 * transaction = make an entry in the transactions table + update the balance in the accounts table
 */
func (d AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? WHERE account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? WHERE account_id = ?`, t.Amount, t.AccountId)
	}

	// in case of error Rollback the transaction, and changes from both tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// getting the last inserted transaction id from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// getting the updated account balance from the accounts table
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	// building the transaction object to return
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDB) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = ?"

	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
