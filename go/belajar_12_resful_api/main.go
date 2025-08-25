package main

import (
	"belajar-12-resful-api/app"
	"belajar-12-resful-api/controller"
	"belajar-12-resful-api/helper"
	"belajar-12-resful-api/middleware"
	"belajar-12-resful-api/repository"
	"belajar-12-resful-api/service"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	categoryController := controller.NewCategoryController(categoryService)
	
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler : middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}