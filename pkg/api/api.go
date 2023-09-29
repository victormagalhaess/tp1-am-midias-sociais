package api

import (
	"log"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/victormagalhaess/go-api-template/docs"

	"github.com/gorilla/mux"
	"github.com/victormagalhaess/go-api-template/pkg/api/controllers"
	"github.com/victormagalhaess/go-api-template/pkg/api/middlewares"
)

type Application struct {
	Router *mux.Router
}

func (a *Application) InitializeRouter() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *Application) InitializeSwagger() {
	a.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

func (a *Application) initializeRoutes() {
	a.Router.HandleFunc("/api/v1/healthcheck", controllers.Healthcheck).Methods("GET")
	a.Router.HandleFunc("/api/v1/tracker", controllers.Tracker).Methods("POST")
	a.Router.Use(middlewares.Logger(log.Default()))
}
