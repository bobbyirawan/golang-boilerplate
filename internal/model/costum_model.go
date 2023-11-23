package model

import "time"

type ListChat struct {
	ChatID      string `gorm:"chat_id"`
	Username    string `gorm:"username"`
	SenderID    string `gorm:"sender_id"`
	RecipientID string `gorm:"recipient_id"`
	LastMessage string `gorm:"last_message"`
	// StatusContentSend string `gorm:"status_content_send"`
	// StatusSendRecieve string `gorm:"status_send_recieve"`
	CreatedAt string `gorm:"created_at"`
}

type ListDetailChatUser struct {
	ChatID                  string    `gorm:"chat_id"`
	ChatType                string    `gorm:"chat_type"`
	SenderID                string    `gorm:"sender_id"`
	RecipientID             string    `gorm:"recipient_id"`
	LastMessage             string    `gorm:"last_message"`
	StatusPesanPengirim     string    `gorm:"status_pesan_pengirim"`
	StatusPesanPenerima     string    `gorm:"status_pesan_penerima"`
	MessageCreated          time.Time `gorm:"message_created"`
	MessageRecipientCreated time.Time `gorm:"message_recipient_created"`
}
