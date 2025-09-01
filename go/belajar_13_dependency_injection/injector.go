//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"

	"belajar-12-resful-api/app"
	"belajar-12-resful-api/controller"
	"belajar-12-resful-api/middleware"
	"belajar-12-resful-api/repository"
	"belajar-12-resful-api/service"

	"github.com/go-playground/validator"
	"github.com/google/wire"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build( 
		app.NewDb, 
		validator.New, 
		categorySet, 
		app.NewRouter, 
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}