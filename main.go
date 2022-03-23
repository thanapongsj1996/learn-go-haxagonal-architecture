package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"hexagonal-architecture/repository"
	"hexagonal-architecture/service"
)

func main() {
	db, err := sqlx.Open("postgres", "postgresql://localhost:5432/todos-app?sslmode=disable")
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)

	customers, err := customerService.GetCustomers()
	if err != nil {
		panic(err)
	}
	fmt.Println("customers : ", customers)

	customer, err := customerService.GetCustomer(7)
	if err != nil {
		panic(err)
	}
	fmt.Println("customer : ", customer)
}
