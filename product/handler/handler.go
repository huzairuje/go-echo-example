package handler

import (
	"context"
	"database/sql"
	"github.com/huzairuje/chatat_backend_engineer/product/repository"
	"github.com/huzairuje/chatat_backend_engineer/product/request"
	"github.com/huzairuje/chatat_backend_engineer/product/service"
	"github.com/huzairuje/chatat_backend_engineer/response"
	"github.com/huzairuje/chatat_backend_engineer/util"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productRepository repository.Repository
	db                *sql.DB
}

func NewProductHandler(db *sql.DB) *productHandler {
	repo := service.NewProductService(db)
	return &productHandler{
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
	var ctxBuiltIn context.Context
	var req request.CreateProductRequest
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, util.BadRequest, nil, err.Error())
	}
	if err := ctx.Validate(req); err != nil {
		return response.ValidationError(ctx, util.ValidationError, nil, err)
	}
	data, err := p.productRepository.Store(ctxBuiltIn, req)
	if err != nil {
		return response.InternalServerError(ctx, util.SomethingWentWrong, nil, err)
	}
	return response.DataWithoutMeta(ctx, data)
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
		return response.InternalServerError(ctx, util.SomethingWentWrong, nil, err)
	}
	return response.DataWithoutMeta(ctx, data)
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
		return response.InternalServerError(ctx, util.SomethingWentWrong, nil, err)
	}
	return response.DataWithoutMeta(ctx, data)
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
	var ctxBuiltIn context.Context
	var req request.UpdateProductRequest
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, util.BadRequest, nil, err.Error())
	}
	if err := ctx.Validate(req); err != nil {
		return response.ValidationError(ctx, util.ValidationError, nil, err)
	}
	data, err := p.productRepository.Update(ctxBuiltIn, id, req)
	if err != nil {
		return response.InternalServerError(ctx, util.SomethingWentWrong, nil, err)
	}
	return response.DataWithoutMeta(ctx, data)
}

// Delete Product godoc
// @Summary Delete Product
// @Description get Delete Product
// @Tags products
// @Accept  json
// @Produce  json
// @Router /products/{id} [delete]
func (p productHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	var ctxBuiltIn context.Context
	err := p.productRepository.Destroy(ctxBuiltIn, id)
	if err != nil {
		return response.InternalServerError(ctx, util.SomethingWentWrong, nil, err)
	}
	return response.SingleData(ctx, util.Success, nil, nil)
}
