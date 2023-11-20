package domain

type Customer struct {	
	Id string 
	Name string
	City string
	Zipcode string
	DateOfBirth string
	Status string
}

// Secondary port, the Repository Interface; Adapter = CustomerRepositoryDB
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, error)
}

