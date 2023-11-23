package repository

import (
	"go-chat/internal/model"

	"gorm.io/gorm"
)

type (
	ContactRepository interface {
		Create(contact *model.Contact) error
		GetAllByID(user_id string, contact *[]model.Contact) error
	}

	contactRepository struct {
		db *gorm.DB
	}
)

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{
		db: db,
	}
}

func (impl *contactRepository) Create(contact *model.Contact) error {
	return impl.db.Create(contact).Error
}

func (impl *contactRepository) GetAllByID(id string, contact *[]model.Contact) error {
	return impl.db.Find(contact, "user_id = ?", id).Error
}
