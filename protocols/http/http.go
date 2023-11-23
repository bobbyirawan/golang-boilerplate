package http

import (
	"fmt"

	"go-chat/internal/controller"
	"go-chat/pkg/config"

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

// func printHeadersMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// Mencetak semua header yang masuk
// 		fmt.Println("client token:", c.Request().Header.Get("X-CSRF-TOKEN"))

// 		// Melanjutkan ke handler berikutnya
// 		return next(c)
// 	}
// }

func (s *HttpServer) Start() {
	e := echo.New()
	e.HideBanner = false
	e.Debug = true

	// Middleware Echo
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
	}))

	// Register routes
	s.controller.UserController.Routes(e.Group("/user"))
	s.controller.AuthController.Routes(e.Group("/auth"))
	s.controller.ContactController.Routes(e.Group("/contact"))
	s.controller.MessageRecipientController.Routes(e.Group("/getallmessages"))

	s.controller.ChatController.Routes(e.Group("/chats"))

	e.Start(fmt.Sprintf("%s:%s", s.env.HOST, s.env.PORT))
}
