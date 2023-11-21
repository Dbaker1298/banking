package service

import (
	"github.com/Dbaker1298/banking/app/errs"
	"github.com/Dbaker1298/banking/domain"
	"github.com/Dbaker1298/banking/dto"
)

// Primary port
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// Business logic
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// This connects the Primary port (interface) to the Secondary port
func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return s.repo.FindAll(status)
}

// This connects the Primary port (interface) to the Secondary port
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
