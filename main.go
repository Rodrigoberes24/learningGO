package main

import (
	
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rodrigoberes/rest-api/controller"
	"github.com/rodrigoberes/rest-api/db"
	"github.com/rodrigoberes/rest-api/helper"
	"github.com/rodrigoberes/rest-api/model"
	"github.com/rodrigoberes/rest-api/repository"
	"github.com/rodrigoberes/rest-api/router"
	"github.com/rodrigoberes/rest-api/service"

	
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Print("Started Server!")
	// Database
	db := db.DBconnection()
	validate := validator.New()


	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("clips").AutoMigrate(&model.Clips{})

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)
	clipsRepository := repository.NewClipsREpositoryImpl(db)
	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	clipsService := service.NewClipsServiceImpl(clipsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)
	clipsController := controller.NewClipsController(clipsService)

	// Router
	routes := router.NewRouter(tagsController, clipsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
