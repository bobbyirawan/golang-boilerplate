package main

import (
	"go-chat/config"
	"go-chat/container"
	"go-chat/http"
	"go-chat/internal/model"
	"log"
	"os"

	"gorm.io/gorm"
)

func main() {

	if err := container.Container.Invoke(func(server *http.HttpServer, db *gorm.DB, env *config.Environment) {

		// Migrate database
		if env.DBMigrate {
			if err := db.AutoMigrate(&model.User{}); err != nil {
				log.Fatalf("Failed to auto migrate database: %v", err)
				os.Exit(1)
			}
		}

		// Memulai server HTTP
		server.Start()
	}); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}
