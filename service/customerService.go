package service

import "github.com/Dbaker1298/banking/domain"

// Primary port
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// Business logic
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
