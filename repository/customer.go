package repository

type Customer struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	CreatedAt string `db:"createdAt"`
	UpdatedAt string `db:"updatedAt"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}
