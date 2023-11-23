package socket

import (
	"log"

	"github.com/gorilla/websocket"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In

	Upgrader                 *websocket.Upgrader
	SocketConnectionRegistry *SocketConnectionRegistry
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewWebSocket); err != nil {
		log.Fatalf("Failed to provide SetupWebsocket: %v", err)
		return err
	}

	if err := container.Provide(NewSocketConnectionRegistry); err != nil {
		log.Fatalf("Failed to provide NewClientManager: %v", err)
		return err
	}

	return nil
}
