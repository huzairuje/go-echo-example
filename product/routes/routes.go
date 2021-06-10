package routes

import (
	"github.com/huzairuje/chatat_backend_engineer/product/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductRoutes(router *echo.Echo, db *gorm.DB)  {
	productHandler := handler.NewProductHandler(db)
	router.POST("/products", productHandler.Create)
	router.GET("/products/:id", productHandler.Detail)
	router.GET("/products", productHandler.List)
	router.PUT("/products/:id", productHandler.Update)
	router.DELETE("/products/:id", productHandler.Delete)
}
