package dto

type (
	GetUserByIDReq struct {
		UserID string `param:"user_id"`
	}

	GetUserByIDRes struct {
		Username    string `json:"username"`
		Email       string `json:"email"`
		Image       string `json:"image"`
		Description string `json:"description"`
	}
)
