package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{ID: 1, Name: "Thanapong Test", Username: "thanapong"},
		{ID: 2, Name: "Praew Test", Username: "praew"},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.ID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}
