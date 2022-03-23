package service

import (
	"hexagonal-architecture/repository"
	"log"
)

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) customerService {
	return customerService{customerRepo: customerRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responses []CustomerResponse
	for _, customer := range customers {
		response := CustomerResponse{
			ID:       customer.ID,
			Name:     customer.Name,
			Username: customer.Username,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepo.GetById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response := CustomerResponse{
		ID:       customer.ID,
		Name:     customer.Name,
		Username: customer.Username,
	}

	return &response, nil
}
