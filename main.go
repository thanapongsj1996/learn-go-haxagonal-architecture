package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"hexagonal-architecture/repository"
)

func main() {
	db, err := sqlx.Open("postgres", "postgresql://localhost:5432/todos-app?sslmode=disable")
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)
	_ = customerRepository

	customers, err := customerRepository.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("customers : ", customers)

	customer, err := customerRepository.GetById(7)
	if err != nil {
		panic(err)
	}
	fmt.Println("customer 7 : ", customer)
}
