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

func ContactConfig() repository.ContactRepository {
	config := config.NewEnvironment()
	db, _ := container.SetupDatabase(config)
	return repository.NewContactRepository(db)
}

func TestContactRepository_GetAllByID_Success_Case1(t *testing.T) {
	config := config.NewEnvironment()
	db, _ := container.SetupDatabase(config)
	contactRepo := repository.NewContactRepository(db)

	contacts := new([]model.Contact)
	UserID := "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9"
	err := contactRepo.GetAllByID(UserID, contacts)
	assert.Nil(t, err)

	jsonData, _ := json.Marshal(contacts)

	log.Println("contacts : ", string(jsonData))

}

func TestContactRepository_Create_Success_Case1(t *testing.T) {
	contactRepo := ContactConfig()

	contact := new(model.Contact)
	contact.UserID = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff9"

	err := contactRepo.Create(contact)
	assert.Nil(t, err)
}

// error case : user id not registered in table users
func TestContactRepository_Create_Error_Case1(t *testing.T) {
	contactRepo := ContactConfig()

	contact := new(model.Contact)
	contact.UserID = "bfccdfc7-d2b5-4c4c-b9fb-8153cc856ff"

	err := contactRepo.Create(contact)
	assert.NotNil(t, err)
}
