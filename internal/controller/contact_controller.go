package controller

import (
	"go-chat/internal/dto"
	"go-chat/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContactController struct {
	service service.Holder
}

func (impl *ContactController) Routes(g *echo.Group) {
	g.POST("", impl.Create)
	g.GET("/:user_id", impl.GetAllByID)
}

func NewContactController(service service.Holder) *ContactController {
	return &ContactController{
		service: service,
	}
}

func (impl *ContactController) Create(ctx echo.Context) error {
	req := new(dto.CreateContactReq)

	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request body")
	}

	if err := impl.service.ContactService.Create(req); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}

func (impl *ContactController) GetAllByID(ctx echo.Context) error {
	var (
		req = new(dto.GetAllContactByIDReq)
		res = new([]dto.GetAllContactByIDRes)
	)

	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request body")
	}

	if err := impl.service.ContactService.GetAllByID(req, res); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}
