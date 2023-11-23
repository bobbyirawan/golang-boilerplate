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

func MessageRecipientConfig() repository.MessageRecipientRepository {
	config := config.NewEnvironment()
	db, _ := container.SetupDatabase(config)
	return repository.NewMessageRecipientRepository(db)
}

func TestCreateMessageRecipientSuccessCase1(t *testing.T) {
	MRrepo := MessageRecipientConfig()

	mr := new(model.MessageRecipient)
	mr.MessageID = "746095e5-bf2b-4fe0-b1bd-a55b14888e2f"
	mr.RecipientID = "a8a4960d-1bfe-495a-b53a-1f1325a45188"

	err := MRrepo.Create(mr)
	assert.Nil(t, err)
}

func TestGetAllMessageRecipient_Success_Case1(t *testing.T) {
	MRrepo := MessageRecipientConfig()

	listChat := new([]model.ListChat)
	err := MRrepo.GetAllMessages("bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9", listChat)
	assert.Nil(t, err)

	jsonData, _ := json.Marshal(listChat)

	log.Println("list chat : ", string(jsonData))
}
