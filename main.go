package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/victormagalhaess/go-api-template/docs"

	"github.com/victormagalhaess/go-api-template/pkg/api"
)

// @title go_api_template
// @version 1.0
// @description This is the backend application for the go_api_template.
// @termsOfService http://swagger.io/terms/

// @contact.name Victor Magalhães
// @contact.email hello@victordias.dev

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	wait := time.Second * 10

	application := api.Application{}
	application.InitializeRouter()
	application.InitializeSwagger()

	server := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      application.Router,
	}

	go func() {
		log.Println("Starting server on port " + port)
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
}