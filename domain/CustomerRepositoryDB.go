package domain

import (
	"database/sql"

	"github.com/Dbaker1298/banking/app/errs"
	"github.com/Dbaker1298/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

// From the Interface within the domain package, customer.go:
func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		err = d.client.Select(&customers, findAllSql)
		// rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?"
		err = d.client.Select(&customers, findAllSql, status)
		// rows, err = d.client.Query(findAllSql, status)
	}
	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customer table" + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
	// }

	return customers, nil
}

// From the Interface within the domain package, customer.go:
func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer table:" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

// Creating connection to the database
func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{dbClient}
}
