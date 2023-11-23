package controller

import (
	"go-chat/internal/dto"
	"go-chat/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.GET("/", func(c echo.Context) error {
		token := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
		return c.JSON(http.StatusOK, map[string]string{
			"csrfToken": token,
		})
	})
	e.POST("/signup", impl.SignUp)
	e.POST("/login", impl.LogIn)
	e.GET("/logout", impl.Logout)
}

func (impl *AuthController) SignUp(ctx echo.Context) error {
	auth := &dto.SignUp{}
	if err := ctx.Bind(auth); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request body")
	}

	if err := impl.service.AuthService.SignUp(auth); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}

func (impl *AuthController) LogIn(ctx echo.Context) error {
	login := &dto.LoginReq{}
	if err := ctx.Bind(login); err != nil {
		return err
	}

	res, err := impl.service.AuthService.LogIn(login)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = res.AccessToken
	cookie.Path = "/"
	cookie.Domain = "localhost"
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.MaxAge = 3600
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, res)
}

func (impl *AuthController) Logout(ctx echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Path = ""
	cookie.Domain = ""
	cookie.MaxAge = -1
	cookie.Secure = false
	cookie.HttpOnly = true
	ctx.SetCookie(cookie)

	return ctx.String(http.StatusOK, "logout success")
}
