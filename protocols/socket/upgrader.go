package socket

import (
	"fmt"
	"go-chat/pkg/config"
	"net/http"

	"github.com/gorilla/websocket"
)

func NewWebSocket(config *config.Environment) *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// Pastikan permintaan memiliki header Origin yang ada
			origin := r.Header.Get("Origin")

			// Periksa apakah Origin yang diminta ada dalam daftar asal yang diizinkan
			allowedOrigins := []string{fmt.Sprintf("http://%s:%s", config.HOST, "3001"), "http://localhost:3001"}

			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					return true
				}
			}

			return false
		},
	}
}
