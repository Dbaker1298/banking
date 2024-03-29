package domain

// Mock Adapter	
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "John", "Chicago", "12345", "2000-01-01", "1"},
		{"1002", "Jane", "Chicago", "12345", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
