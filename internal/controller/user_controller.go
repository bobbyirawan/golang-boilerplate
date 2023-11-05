package controller

import (
	"go-boilerplate/internal/service"
	"go-boilerplate/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service service.Holder
}

func NewUserController(service service.Holder) *UserController {
	return &UserController{
		service: service,
	}
}

func (impl *UserController) Routes(e *echo.Group) {
	e.GET("/user/:id", impl.GetHandler)
	e.POST("/user", impl.CreateHandler)
	e.PUT("/user/:id", impl.UpdateHandler)
	e.DELETE("/user/:id", impl.DeleteHandler)
}

func (impl *UserController) GetHandler(ctx echo.Context) error {
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

func (impl *UserController) CreateHandler(ctx echo.Context) error {
	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request body")
	}

	if err := impl.service.UserService.CreateUser(user); err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create user")
	}

	return ctx.NoContent(http.StatusCreated)
}

func (impl *UserController) UpdateHandler(ctx echo.Context) error {
	userID := ctx.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := impl.service.UserService.GetUser(uint(id))
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	updatedUser := new(model.User)
	if err := ctx.Bind(updatedUser); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request body")
	}

	user.Username = updatedUser.Username
	user.Email = updatedUser.Email

	if err := impl.service.UserService.UpdateUser(user); err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to update user")
	}

	return ctx.NoContent(http.StatusOK)
}

func (impl *UserController) DeleteHandler(ctx echo.Context) error {
	userID := ctx.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid user ID")
	}

	if err := impl.service.UserService.DeleteUser(uint(id)); err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	return ctx.NoContent(http.StatusNoContent)
}
