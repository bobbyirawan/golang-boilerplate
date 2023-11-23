package container

import (
	"go-chat/internal/controller"
	"go-chat/internal/repository"
	"go-chat/internal/service"
	"go-chat/pkg/config"
	"go-chat/pkg/utils"
	"go-chat/protocols/http"
	"go-chat/protocols/socket"

	"log"

	"go.uber.org/dig"
)

var Container = dig.New()

func init() {

	// UTILS
	if err := utils.Register(Container); err != nil {
		log.Fatalf("Failed to provide utils: %v", err)
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

	if err := socket.Register(Container); err != nil {
		log.Fatalf("Failed to provide Socket: %v", err)
	}

	// Mendefinisikan objek Environment dalam kontainer
	if err := Container.Provide(config.NewEnvironment); err != nil {
		log.Fatalf("Failed to provide Environment: %v", err)
	}

}
