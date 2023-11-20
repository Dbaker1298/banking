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
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
