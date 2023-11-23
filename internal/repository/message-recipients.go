package repository

import (
	"go-chat/internal/model"

	"gorm.io/gorm"
)

type (
	MessageRecipientRepository interface {
		Create(mr *model.MessageRecipient) error
		GetAllMessages(UserID string, listChat *[]model.ListChat) error
	}

	messageRecipientDependency struct {
		mysql *gorm.DB
	}
)

func NewMessageRecipientRepository(mysql *gorm.DB) MessageRecipientRepository {
	return &messageRecipientDependency{
		mysql: mysql,
	}
}

func (impl *messageRecipientDependency) Create(mr *model.MessageRecipient) error {
	return impl.mysql.Create(mr).Error
}

// SELECT m.chat_id, m.id, m.sender_id, mr.recipient_id, m.content, m.status as pengirim_status, mr.status as penerima_status
// FROM message_recipients mr
// JOIN messages m ON (mr.message_id = m.id);

func (impl *messageRecipientDependency) GetAllMessages(UserID string, listChat *[]model.ListChat) error {
	// query := fmt.Sprintf("c.id as 'chat_id',CASE WHEN (SELECT mrr.recipient_id FROM message_recipients mrr JOIN messages mm ON (mrr.message_id = mm.id) WHERE mrr.id = MAX(mr.id)) = '%s' THEN ( SELECT username FROM users WHERE id = ( SELECT mm.sender_id FROM message_recipients mrr JOIN messages mm ON (mrr.message_id = mm.id) WHERE mrr.id = MAX(mr.id))) ELSE (SELECT username FROM users WHERE id = (SELECT mrr.recipient_id FROM message_recipients mrr JOIN messages mm ON (mrr.message_id = mm.id) WHERE mrr.id = MAX(mr.id))) END AS username, (SELECT mm.sender_id  FROM message_recipients mrr JOIN messages mm ON (mrr.message_id = mm.id) WHERE mrr.id = MAX(mr.id)) as sender_id, (SELECT mrr.recipient_id FROM message_recipients mrr JOIN messages mm ON (mrr.message_id = mm.id) WHERE mrr.id = MAX(mr.id)) as recipient_id, (SELECT mm.content FROM message_recipients mrr JOIN messages mm ON (mrr.message_id = mm.id) WHERE mrr.id = MAX(mr.id)) as last_message,(SELECT mm.created_at  FROM message_recipients mrr JOIN messages mm ON (mrr.message_id = mm.id) WHERE mrr.id = MAX(mr.id)) as content_send", UserID)
	// from := fmt.Sprintf("SELECT m.chat_id, MAX(m.id) AS max_message_id FROM message_recipients mr JOIN messages m ON mr.message_id = m.id WHERE (m.sender_id = '%s' OR mr.recipient_id = '%s' ) GROUP BY m.chat_id", UserID, UserID)

	subQueryFrom := impl.mysql.Table("message_recipients AS mr").
		Select("DISTINCT m.chat_id AS id").
		Joins("JOIN messages m ON mr.message_id = m.id").
		Where("(m.sender_id = ? OR mr.recipient_id = ?)", UserID, UserID)
	subQuery1 := impl.mysql.Table("users").
		Select("username").
		Where("id =  m2.sender_id")
	subQuery2 := impl.mysql.Table("users").
		Select("username").
		Where("id =  mr2.recipient_id")
	subQuery3 := impl.mysql.Table("messages AS m3").
		Select("id").
		Where("m3.chat_id = chat.id").
		Order("m3.created_at DESC").
		Limit(1)

	return impl.mysql.Table("(?) AS chat", subQueryFrom).
		Select("chat.id AS chat_id, m2.sender_id AS sender_id, m2.content AS last_message, mr2.recipient_id  AS recipient_id, m2.created_at, CASE WHEN mr2.recipient_id = ? THEN (?) ELSE (?) END AS username", UserID, subQuery1, subQuery2).
		Joins("JOIN messages m2 ON m2.id = (?)", subQuery3).
		Joins("JOIN message_recipients mr2 ON mr2.message_id = m2.id").
		Order("m2.created_at DESC").
		Scan(listChat).Error
}
