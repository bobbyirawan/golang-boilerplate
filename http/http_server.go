package http

import (
	"fmt"
	"go-chat/config"
	"go-chat/internal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	controller controller.Holder
	env        *config.Environment
}

func NewHttpServer(controller controller.Holder, env *config.Environment) *HttpServer {
	return &HttpServer{
		controller: controller,
		env:        env,
	}
}

func (s *HttpServer) Start() {
	e := echo.New()

	// Middleware Echo
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register routes
	s.controller.UserController.Routes(e.Group("/user"))
	s.controller.AuthController.Routes(e.Group("/auth"))

	e.Start(fmt.Sprintf(":%s", s.env.PORT))
}
