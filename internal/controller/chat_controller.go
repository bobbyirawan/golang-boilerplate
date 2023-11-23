package controller

import (
	"go-chat/internal/dto"
	"go-chat/internal/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChatController struct {
	service service.Holder
}

func NewChatController(service service.Holder) *ChatController {
	return &ChatController{
		service: service,
	}
}

func (impl *ChatController) Routes(e *echo.Group) {
	e.GET("/person", impl.Chat)
	e.GET("/:user_id", impl.GetListChat)
	e.GET("/user", impl.GetListChatUser)
}

func (impl *ChatController) Chat(ctx echo.Context) error {
	req := new(dto.ChatReq)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "username is required")
	}

	return impl.service.ChatService.HandleSocketConnection(ctx, req)
}

func (impl *ChatController) GetListChatUser(ctx echo.Context) error {
	var (
		req = new(dto.GetListChatUserReq)
		res = new(dto.GetListChatUserRes)
	)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "username is required")
	}

	log.Println(*req)

	if err := impl.service.ChatService.GetListChatUser(req, res); err != nil {
		return nil
	}

	return ctx.JSON(http.StatusOK, res)
}

func (impl *ChatController) GetListChat(ctx echo.Context) error {
	var (
		req = new(dto.GetListChatReq)
		res = new(dto.GetListChatRes)
	)

	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, "invalid request")
	}

	if err := impl.service.MessageRecipientService.GetAllMessages(req, res); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}
