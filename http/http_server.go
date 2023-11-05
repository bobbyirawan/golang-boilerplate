package http

import (
	"go-boilerplate/internal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	controller *controller.Holder
}

func NewHttpServer(controller *controller.Holder) *HttpServer {
	return &HttpServer{
		controller: controller,
	}
}

func (s *HttpServer) Start() {
	e := echo.New()

	// Middleware Echo
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register routes
	s.controller.UserController.Routes(e.Group("/user"))

	e.Start(":8080")
}
