package services

import (
	models "github.com/huavanthong/microservice-golang/product-api-v3/models/product"
	"github.com/huavanthong/microservice-golang/product-api-v3/payload"
)

type ProductService interface {
	CreateProduct(pr *payload.RequestCreateProduct) (*models.Product, error)
	FindAllProducts(page int, limit int, currency string) (interface{}, error)
	FindProductByID(id string, currency string) (*models.Product, error)
	FindProductByName(name string, currency string) ([]*models.Product, error)
	FindProductByCategory(category string, currency string) ([]*models.Product, error)
	UpdateProduct(id string, pr *payload.RequestUpdateProduct) (*models.Product, error)
	DeleteProduct(id string) error
}
