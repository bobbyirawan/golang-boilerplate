package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	PORT string

	SSHUsername  string
	SSHPasssword string
	SSHHostname  string
	SSHPort      string
	SSHTunnel    bool

	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
	DBMigrate  bool
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

	dbMigrate, _ := strconv.ParseBool(os.Getenv("DB_MIGRATE"))
	if err != nil {
		log.Fatalf("Failed to convert DB_MIGRATE value: %v", err)
	}

	return &Environment{
		PORT: os.Getenv("PORT"),

		SSHUsername:  os.Getenv("SSH_USERNAME"),
		SSHPasssword: os.Getenv("SSH_PASSWORD"),
		SSHHostname:  os.Getenv("SSH_HOSTNAME"),
		SSHPort:      os.Getenv("SSH_PORT"),
		SSHTunnel:    ssh,

		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort:     os.Getenv("DB_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBMigrate:  dbMigrate,
	}
}
