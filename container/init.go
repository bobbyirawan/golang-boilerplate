package container

import (
	"go-boilerplate/config"
	"go-boilerplate/http"
	"go-boilerplate/internal/controller"
	"go-boilerplate/internal/repository"
	"go-boilerplate/internal/service"
	"log"

	"go.uber.org/dig"
)

var Container = dig.New()

func init() {

	// Mendefinisikan objek Environment dalam kontainer
	if err := Container.Provide(config.NewEnvironment); err != nil {
		log.Fatalf("Failed to provide Environment: %v", err)
	}

	// Database Configuration
	if err := Container.Provide(SetupDatabase); err != nil {
		log.Fatalf("Failed to provide Environment: %v", err)
	}

	if err := repository.Register(Container); err != nil {
		log.Fatalf("Failed to provide Repository: %v", err)
	}

	if err := service.Register(Container); err != nil {
		log.Fatalf("Failed to provide Service: %v", err)
	}

	if err := controller.Register(Container); err != nil {
		log.Fatalf("Failed to provide Controller: %v", err)
	}

	if err := Container.Provide(http.NewHttpServer); err != nil {
		log.Fatalf("Failed to provide Http Server: %v", err)
	}
}
