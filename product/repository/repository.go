package repository

import (
	"github.com/huzairuje/chatat_backend_engineer/product/entity"
	"github.com/huzairuje/chatat_backend_engineer/product/request"
)

type Repository interface {
	Store(req request.CreateProductRequest) (*entity.Product, error)
	List() ([]*entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(id string, req request.UpdateProductRequest) (*entity.Product, error)
	Destroy(id string) error
}
