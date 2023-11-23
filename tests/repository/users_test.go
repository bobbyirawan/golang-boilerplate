package repository

import (
	"encoding/json"
	"go-chat/container"
	"go-chat/internal/model"
	"go-chat/internal/repository"
	"go-chat/pkg/config"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	config := config.NewEnvironment()
	db, _ := container.SetupDatabase(config)
	useRepo := repository.NewUserRepository(db)

	user := new(model.User)
	user.ID = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9"

	err := useRepo.GetUserByID(user)
	assert.Nil(t, err)

	jsonData, _ := json.Marshal(user)

	log.Println("user : ", string(jsonData))

}
