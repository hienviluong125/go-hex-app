package productmodule

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Status      bool   `json:"status"`
}

type UpdateProduct struct {
	Id          int     `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Quantity    *int    `json:"quantity"`
}

type ProductRepository interface {
	FindAll() ([]Product, error)
	Save(product Product) error
	FindById(id int) (*Product, error)
	UpdateById(id int, updateProductParams UpdateProduct) error
	DestroyById(id int) error
}
