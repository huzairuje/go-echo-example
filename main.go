package main

import (
	"github.com/huzairuje/chatat_backend_engineer/database"
	_ "github.com/huzairuje/chatat_backend_engineer/docs/products"
	productsDomain "github.com/huzairuje/chatat_backend_engineer/product/routes"
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
// @description you can access the code on github https://github.com/huzairuje/go-echo-example/
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

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
	//initiate routes by domain (this domain just products domain)
	productsDomain.ProductRoutes(router, db)
	//initiate swagger routes
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	//start web (echo) service
	err := router.Start(loadCfg.ServerPort)
	if err != nil {
		logrus.Info(err)
		panic(err)
	}
}
