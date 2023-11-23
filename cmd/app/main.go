package main

import (
	"go-chat/container"
	"go-chat/internal/model"
	"go-chat/pkg/config"
	"go-chat/protocols/http"

	"log"
	"os"

	"gorm.io/gorm"
)

func main() {

	if err := container.Container.Invoke(func(server *http.HttpServer, db *gorm.DB, env *config.Environment) {

		// Migrate database
		if env.MYSQL_DB_MIGRATE {
			if err := db.AutoMigrate(&model.User{}, &model.Contact{}); err != nil {
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
