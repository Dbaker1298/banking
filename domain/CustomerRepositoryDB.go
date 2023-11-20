package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/Dbaker1298/banking/app/errs"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

// From the Interface within the domain package, customer.go:
func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Printf("Error while querying customer table: %v", err)
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Printf("Error while scanning customer table: %v", err)
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)

	}
	return customers, nil
}

// From the Interface within the domain package, customer.go:
func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"

	row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Printf("Error while scanning customer table: %v", err)
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "david:abcd1234@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	client.SetConnMaxLifetime(time.Minute * 3)

	return CustomerRepositoryDB{client}
}
