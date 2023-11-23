package repository

import (
	"encoding/json"
	"go-chat/container"
	"go-chat/internal/model"
	"go-chat/internal/repository"
	"go-chat/pkg/config"
	"go-chat/pkg/utils"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ChatConfig() repository.ChatRepository {
	config := config.NewEnvironment()
	db, _ := container.SetupDatabase(config)
	return repository.NewChatRepository(db)
}

// test create normally
func TestChatRepository_Create_Success_Case1(t *testing.T) {
	config := config.NewEnvironment()
	db, _ := container.SetupDatabase(config)
	chatRepo := repository.NewChatRepository(db)
	chat := new(model.Chat)

	chat.ID = utils.GenerateUUID()
	chat.Name = "test default chat type"

	err := chatRepo.Create(chat)
	assert.Nil(t, err)

	jsonData, _ := json.Marshal(chat)

	log.Println("chat : ", string(jsonData))

}

func TestChatRepository_GetListMessageUser_Success_Case1(t *testing.T) {
	chatRepo := ChatConfig()

	listChatUser := new([]model.ListDetailChatUser)
	ChatID := "72b84e9c-f65e-4835-aa58-2a1f2185fff4"

	err := chatRepo.ListChatUser(ChatID, listChatUser)
	assert.Nil(t, err)

	jsonData, _ := json.Marshal(listChatUser)
	log.Println("list Chat User : ", string(jsonData))
}
