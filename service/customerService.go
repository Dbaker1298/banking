package service

import (
	"github.com/Dbaker1298/banking/app/errs"
	"github.com/Dbaker1298/banking/domain"
)

// Primary port
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// Business logic
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// This connects the Primary port (interface) to the Secondary port
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

// This connects the Primary port (interface) to the Secondary port
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
