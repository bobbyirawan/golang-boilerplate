package repository

import (
	"go-chat/internal/model"

	"gorm.io/gorm"
)

type (
	ChatRepository interface {
		Create(chat *model.Chat) error
		ListChatUser(ChatID string, listChatUser *[]model.ListDetailChatUser) error
	}

	chatRepository struct {
		mysql *gorm.DB
	}
)

func NewChatRepository(mysql *gorm.DB) ChatRepository {
	return &chatRepository{mysql: mysql}
}

func (impl *chatRepository) Create(chat *model.Chat) error {
	return impl.mysql.Create(chat).Error
}

func (impl *chatRepository) ListChatUser(ChatID string, listChatUser *[]model.ListDetailChatUser) error {
	return impl.mysql.Table("message_recipients").
		Select("chats.id AS chat_id, chats.chat_type, messages.sender_id, message_recipients.recipient_id, messages.content AS last_message, messages.status AS status_pesan_pengirim, message_recipients.status AS status_pesan_penerima, messages.created_at AS message_created, message_recipients.created_at AS message_recipient_created").
		Joins("JOIN messages ON messages.id = message_recipients.message_id").
		Joins("JOIN chats ON chats.id = messages.chat_id").
		Where("chats.id = ?", ChatID).
		Order("CASE WHEN messages.updated_at IS NOT NULL THEN messages.updated_at ELSE messages.created_at END ASC").
		Scan(&listChatUser).Error
}
