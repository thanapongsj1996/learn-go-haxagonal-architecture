package service

type CustomerResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}
