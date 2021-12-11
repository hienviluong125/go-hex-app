package productmodule

import (
	"hienviluong125/go-hex-app/errorhandler"

	"github.com/jmoiron/sqlx"
)

type ProductRepositoryMysql struct {
	dbClient *sqlx.DB
}

func NewProductRepositoryMysql(dbClient *sqlx.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{
		dbClient: dbClient,
	}
}

func (repo *ProductRepositoryMysql) FindAll() ([]Product, error) {
	products := []Product{}
	query := "SELECT id, name, description, quantity, status FROM products WHERE status in (1)"
	err := repo.dbClient.Select(&products, query)

	if err != nil {
		return nil, errorhandler.ErrCannotGetRecord("product", err)
	}

	return products, nil
}

func (repo *ProductRepositoryMysql) FindById(id int) (*Product, error) {
	var product Product
	query := "SELECT id, name, description, quantity, status FROM products WHERE id = ? AND status in (1)"
	err := repo.dbClient.Get(&product, query, id)

	if err != nil {
		return nil, errorhandler.ErrCannotGetRecord("product", err)
	}

	return &product, nil
}

func (repo *ProductRepositoryMysql) Save(product Product) error {
	query := "INSERT INTO products(name, description, quantity) VALUES(?, ?, ?)"

	_, err := repo.dbClient.Exec(query, product.Name, product.Description, product.Quantity)

	if err != nil {
		return errorhandler.ErrInternal(err)
	}

	return nil
}

func (repo *ProductRepositoryMysql) UpdateById(id int, updateProductParams UpdateProduct) error {
	query := "UPDATE products set name=?, description=?, quantity=? WHERE id=?"

	_, err := repo.dbClient.Exec(query, updateProductParams.Name, updateProductParams.Description, updateProductParams.Quantity, id)

	if err != nil {
		return errorhandler.ErrInternal(err)
	}

	return nil
}

func (repo *ProductRepositoryMysql) DestroyById(id int) error {
	query := "UPDATE products set status=0 WHERE id=?"

	_, err := repo.dbClient.Exec(query, id)

	if err != nil {
		return errorhandler.ErrInternal(err)
	}

	return nil
}
