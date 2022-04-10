package config

import (
	"net/http"
	"tidy/dbase"
	"tidy/pkg/handler"
	"tidy/pkg/repository"
	"tidy/pkg/service"
)

func Config(db dbase.Database) *http.ServeMux {

	router := http.NewServeMux()
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	handler.Register(router)

	return router
}
