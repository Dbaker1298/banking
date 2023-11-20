package service

import "github.com/Dbaker1298/banking/domain"

// Primary port
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

// Business logic
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// This connects the Primary port (interface) to the Secondary port
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// This connects the Primary port (interface) to the Secondary port
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
