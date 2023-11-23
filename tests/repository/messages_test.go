package repository

import (
	"go-chat/container"
	"go-chat/internal/model"
	"go-chat/internal/repository"
	"go-chat/pkg/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MessageConfig() repository.MessageRepository {
	config := config.NewEnvironment()
	db, _ := container.SetupDatabase(config)
	return repository.NewMessageRepository(db)
}

func TestCreateMessageSuccessCase1(t *testing.T) {
	messageRepo := MessageConfig()

	message := new(model.Message)
	message.ChatID = "2b44606e-f4c4-4e8e-b463-85b4b6c3bda4"
	message.SenderID = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9"
	message.Content = "hello"
	err := messageRepo.Create(message)
	assert.Nil(t, err)
}
