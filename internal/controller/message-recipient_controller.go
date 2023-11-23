package controller

import (
	"go-chat/internal/service"

	"github.com/labstack/echo/v4"
)

type MessageRecipientController struct {
	service service.Holder
}

func (impl *MessageRecipientController) Routes(g *echo.Group) {
}

func NewMessageRecipientController(service service.Holder) *MessageRecipientController {
	return &MessageRecipientController{
		service: service,
	}
}
