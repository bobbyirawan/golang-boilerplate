package socket

import (
	"go-chat/protocols/socket"
	"testing"

	"github.com/gorilla/websocket"
)

func ConnectionRegistryConfig() (*socket.SocketConnectionRegistry, *websocket.Conn) {
	registry := socket.NewSocketConnectionRegistry()
	conn := &websocket.Conn{}
	return registry, conn
}

func TestSocketConnectionRegistry_RegisterAndUnregisterConnection(t *testing.T) {
	ur, conn := ConnectionRegistryConfig()

	client := new(socket.Client)

	userID := "user123"

	client.Email = "user123"
	client.Conn = conn

	// Register connection
	ur.RegisterClient(client)

	// Check if the connection is registered
	if _, ok := ur.Connections[userID]; !ok {
		t.Errorf("Expected connection to be registered for user %s, but it's not", userID)
	}

	// Unregister connection
	ur.UnregisterClient(client)

	// Check if the connection is unregistered
	if _, ok := ur.Connections[userID]; ok {
		t.Errorf("Expected connection to be unregistered for user %s, but it's still registered", userID)
	}
}
