package service

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/huzairuje/chatat_backend_engineer/product/entity"
	"github.com/huzairuje/chatat_backend_engineer/product/repository"
	"github.com/huzairuje/chatat_backend_engineer/product/request"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"time"
)

type ProductService struct {
	*gorm.DB
}

func NewProductService(db *gorm.DB) repository.Repository {
	return ProductService{db}
}

func (p ProductService) Store(ctx context.Context, req request.CreateProductRequest) (*entity.Product, error) {
	createdAtNow := time.Now()
	var product entity.Product
	product.Title = req.Title
	product.Description = req.Description
	product.Rating = req.Rating
	product.Image = req.Image
	product.CreatedAt = createdAtNow
	err := p.DB.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p ProductService) List() ([]*entity.Product, error) {
	var listProducts []*entity.Product
	err := p.DB.Find(&listProducts).Error
	if err != nil {
		return nil, err
	}
	return listProducts, nil
}

func (p ProductService) FindByID(id string) (*entity.Product, error) {
	uuidFromString, err := uuid.FromString(id)
	if err != nil {
		logrus.Info("err ", uuidFromString)
		return nil, err
	}
	var product entity.Product
	err = p.DB.First(&product, "id = ?", uuidFromString).Error
	if err != nil {
		logrus.Info("err ", product)
		return nil, err
	}
	return &product, nil
}

func (p ProductService) Update(ctx context.Context, id string, req request.UpdateProductRequest) (*entity.Product, error) {
	var product entity.Product
	uuidFromString, _ := uuid.FromString(id)
	product.Id = uuidFromString
	updatedAtNow := time.Now()
	err := p.DB.Model(&product).UpdateColumns(entity.Product{
		Title: req.Title,
		Description: req.Description,
		Rating: req.Rating,
		Image: req.Image,
		UpdatedAt: updatedAtNow,
	}).Error

	if err != nil {
		log.Println("error when commit transaction: ", err.Error())
		return nil, err
	}
	logrus.Println("update product done!")
	return &product, nil
}

func (p ProductService) Destroy(ctx context.Context, id string) error {
	uuidFromString, _ := uuid.FromString(id)
	var product entity.Product
	product.Id = uuidFromString

	isExisting, err := p.FindByID(id)
	if isExisting == nil {
		return err
	}
	deleteAtNow := time.Now()
	err = p.DB.Model(&product).UpdateColumns(entity.Product{
		DeletedAt: deleteAtNow,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
