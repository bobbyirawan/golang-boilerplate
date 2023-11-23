package dto

type (
	CreateContactReq struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)

type (
	GetAllContactByIDReq struct {
		UserID string `param:"user_id"`
	}

	GetAllContactByIDRes struct {
		Username    string `json:"username"`
		RecipientID string `json:"recipient_id"`
		Email       string `json:"email"`
	}
)
