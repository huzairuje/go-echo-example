package repository

import (
	"context"
	"github.com/huzairuje/chatat_backend_engineer/product/entity"
	"github.com/huzairuje/chatat_backend_engineer/product/request"
)

type Repository interface {
	Store(ctx context.Context, req request.CreateProductRequest) (*entity.Product, error)
	List() ([]*entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(ctx context.Context, id string, req request.UpdateProductRequest) (*entity.Product, error)
	Destroy(ctx context.Context, id string) error
}
