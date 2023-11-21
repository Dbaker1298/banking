package domain

import "github.com/Dbaker1298/banking/app/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// Secondary port, the Repository Interface; Adapter = CustomerRepositoryDB
type CustomerRepository interface {
	// status == 1 status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
