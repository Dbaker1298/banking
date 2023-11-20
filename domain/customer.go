package domain

type Customer struct {	
	Id string 
	Namme string
	City string
	Zipcode string
	DateofBirth string
	Status string
}

// Secondary port, the Repository Interface
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

