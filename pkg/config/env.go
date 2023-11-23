package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	HOST         string
	PORT         string
	SECRET_KEY   string
	X_CSRF_TOKEN string

	SSH_USERNAME string
	SSH_PASSWORD string
	SSH_HOSTNAME string
	SSH_PORT     string
	SSH_TUNNEL   bool

	MYSQL_USERNAME   string
	MYSQL_PASSWORD   string
	MYSQL_PORT       string
	MYSQL_HOST       string
	MYSQL_DB_NAME    string
	MYSQL_DB_LOC     string
	MYSQL_DB_MIGRATE bool
}

func NewEnvironment() *Environment {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	ssh, err := strconv.ParseBool(os.Getenv("SSH_TUNNEL"))
	if err != nil {
		log.Fatalf("Failed to convert SSH_TUNNEL value: %v", err)
	}

	dbMigrate, _ := strconv.ParseBool(os.Getenv("MYSQL_DB_MIGRATE"))
	if err != nil {
		log.Fatalf("Failed to convert DB_MIGRATE value: %v", err)
	}

	return &Environment{
		HOST:         os.Getenv("HOST"),
		PORT:         os.Getenv("PORT"),
		SECRET_KEY:   os.Getenv("SECRET_KEY"),
		X_CSRF_TOKEN: os.Getenv("X_CSRF_TOKEN"),

		SSH_USERNAME: os.Getenv("SSH_USERNAME"),
		SSH_PASSWORD: os.Getenv("SSH_PASSWORD"),
		SSH_HOSTNAME: os.Getenv("SSH_HOSTNAME"),
		SSH_PORT:     os.Getenv("SSH_PORT"),
		SSH_TUNNEL:   ssh,

		MYSQL_USERNAME:   os.Getenv("MYSQL_USERNAME"),
		MYSQL_PASSWORD:   os.Getenv("MYSQL_PASSWORD"),
		MYSQL_PORT:       os.Getenv("MYSQL_PORT"),
		MYSQL_HOST:       os.Getenv("MYSQL_HOST"),
		MYSQL_DB_NAME:    os.Getenv("MYSQL_DB_NAME"),
		MYSQL_DB_LOC:     os.Getenv("MYSQL_DB_LOC"),
		MYSQL_DB_MIGRATE: dbMigrate,
	}
}
