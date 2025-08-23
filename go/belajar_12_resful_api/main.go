package main

import (
	"belajar-12-resful-api/app"
	"belajar-12-resful-api/controller"
	"belajar-12-resful-api/exeption"
	"belajar-12-resful-api/helper"
	"belajar-12-resful-api/middleware"
	"belajar-12-resful-api/repository"
	"belajar-12-resful-api/service"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	categoryController := controller.NewCategoryController(categoryService)
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exeption.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler : middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}