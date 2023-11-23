package dto

type Message struct {
	ChatID      *string `json:"chat_id"`      // room chat
	SenderID    string  `json:"sender_id"`    // sender email
	RecipientID string  `json:"recipient_id"` // recipient email
	Content     string  `json:"content"`
}

type ChatReq struct {
	UserID string `query:"user_id"`
}

type (
	GetListChatUserReq struct {
		ChatID string `query:"chat_id"`
	}

	GetListChatUserRes struct {
		Messages interface{} `json:"messages"`
	}
)

type (
	GetListChatReq struct {
		UserID string `param:"user_id"`
	}

	GetListChatRes struct {
		Messages interface{} `json:"messages"`
	}
)
