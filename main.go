package main

import (
	"github.com/huzairuje/chatat_backend_engineer/database"
	_ "github.com/huzairuje/chatat_backend_engineer/docs/products"
	"github.com/huzairuje/chatat_backend_engineer/product/routes"
	"github.com/huzairuje/chatat_backend_engineer/response"
	"github.com/huzairuje/chatat_backend_engineer/util"
	"github.com/huzairuje/chatat_backend_engineer/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Products API
// @version 1.0
// @description This is a products server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http
func main() {
	//load config from config.yml
	loadCfg := util.LoadConfig()
	//initiate database connection
	db := database.CreateDB(loadCfg)
	//initiate echo App instance
	router := echo.New()
	//initiate validator
	router.Validator = validator.NewValidator()
	//custom response for not matching routes
	echo.NotFoundHandler = func(c echo.Context) error {
		// render 404 custom response
		return response.NotFound(c, "Not Matching of Any Routes", nil, "Not Found")
	}
	//initiate routes by domain
	routes.ProductRoutes(router, db)
	//initiate swagger routes
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	//start web (echo) service
	err := router.Start(loadCfg.ServerPort)
	if err != nil {
		logrus.Info(err)
		panic(err)
	}
}
