package socket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	*websocket.Conn
	Email  string
	UserID string
}

type SocketConnectionRegistry struct {
	Connections map[string]*Client
	Mu          sync.Mutex
}

func NewSocketConnectionRegistry() *SocketConnectionRegistry {
	return &SocketConnectionRegistry{
		Connections: make(map[string]*Client),
	}
}

func (cm *SocketConnectionRegistry) RegisterClient(client *Client) {
	cm.Mu.Lock()
	defer cm.Mu.Unlock()
	cm.Connections[client.UserID] = client
	log.Println("ADD NEW CONNECTION: ", client.Email)
}

// UnregisterClient unregisters a client
func (cm *SocketConnectionRegistry) UnregisterClient(client *Client) {
	cm.Mu.Lock()
	defer cm.Mu.Unlock()
	delete(cm.Connections, client.UserID)
	log.Println("REMOVE CONNECTION : ", client.Email)
}
