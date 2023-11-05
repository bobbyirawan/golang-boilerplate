package controller

import (
	"go-boilerplate/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	service service.Holder
}

func NewAuthController(s service.Holder) *AuthController {
	return &AuthController{
		service: s,
	}
}

func (impl *AuthController) Routes(e *echo.Group) {
	e.GET("/user/:id", impl.GetHandler)
}

func (impl *AuthController) GetHandler(ctx echo.Context) error {
	userID := ctx.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := impl.service.UserService.GetUser(uint(id))
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, user)
}
