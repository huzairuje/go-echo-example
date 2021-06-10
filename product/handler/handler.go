package handler

import (
	"context"
	"errors"
	"github.com/huzairuje/chatat_backend_engineer/product/repository"
	"github.com/huzairuje/chatat_backend_engineer/product/request"
	"github.com/huzairuje/chatat_backend_engineer/product/service"
	"github.com/huzairuje/chatat_backend_engineer/response"
	"github.com/huzairuje/chatat_backend_engineer/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type productHandler struct {
	context           context.Context
	productRepository repository.Repository
	db                *gorm.DB
}

func NewProductHandler(db *gorm.DB) *productHandler {
	repo := service.NewProductService(db)
	var ctx context.Context
	return &productHandler{
		context:           ctx,
		productRepository: repo,
		db:                db,
	}
}

// Create godoc
// @Summary Create Product
// @Description Post Create Product
// @Tags products
// @Accept */*
// @Produce  json
// @Success 200 {object} entity.Product
// @Router /products [post]
func (p productHandler) Create(ctx echo.Context) error {
	var req request.CreateProductRequest
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, util.BadRequest, nil, err.Error())
	}
	if err := ctx.Validate(req); err != nil {
		return response.ValidationError(ctx, util.ValidationError, nil, err.Error())
	}
	data, err := p.productRepository.Store(p.context, req)
	if err != nil {
		return response.InternalServerError(ctx, util.SomethingWentWrong, nil, err.Error())
	}
	return response.SingleData(ctx, util.OK, data, nil)
}

// List Product godoc
// @Summary List Product
// @Description get List Product
// @Tags products
// @Accept */*
// @Produce  json
// @Success 200 {object} entity.Product
// @Produce  json
// @Router /products [get]
func (p productHandler) List(ctx echo.Context) error {
	data, err := p.productRepository.List()
	if err != nil {
		return response.InternalServerError(ctx, util.SomethingWentWrong, nil, err.Error())
	}
	if len(data) == 0 {
		return response.ListData(ctx,  util.OK, map[string]interface{}{}, nil)
	}
	return response.ListData(ctx, util.OK, data, nil)
}

// Detail Product godoc
// @Summary Detail Product
// @Description get Detail Product
// @Tags products
// @Accept  json
// @Success 200 {object} entity.Product
// @Produce  json
// @Router /products/{id} [get]
func (p productHandler) Detail(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := p.productRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(ctx, util.NotFound, nil, err.Error())
		}
		return response.InternalServerError(ctx, util.SomethingWentWrong, nil, err.Error())
	}
	return response.SingleData(ctx, util.OK, data, nil)
}

// Update Product godoc
// @Summary Update Product
// @Description get Update Product
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Product
// @Router /products/{id} [put]
func (p productHandler) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	var req request.UpdateProductRequest
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, util.BadRequest, nil, err.Error())
	}
	if err := ctx.Validate(req); err != nil {
		return response.ValidationError(ctx, util.ValidationError, nil, err.Error())
	}
	data, err := p.productRepository.Update(p.context, id, req)
	if err != nil {
		return response.InternalServerError(ctx, util.SomethingWentWrong, nil, err.Error())
	}
	return response.SingleData(ctx, util.OK, data, nil)
}

// Delete Product godoc
// @Summary Delete Product
// @Description get Delete Product
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Product
// @Router /products/{id} [delete]
func (p productHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err := p.productRepository.Destroy(p.context, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(ctx, util.NotFound, nil, err.Error())
		}
		return response.InternalServerError(ctx, util.SomethingWentWrong, nil, err.Error())
	}
	return response.SingleData(ctx, util.Success, nil, nil)
}
