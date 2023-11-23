package service

import (
	"testing"
)

func TestChatService_HandleConnection(t *testing.T) {
	// Create a mock WebSocket server
	// upgrader := &websocket.Upgrader{}
	// server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	conn, _ := upgrader.Upgrade(w, r, nil)
	// 	defer conn.Close()

	// 	// Handle connection using the ChatService
	// 	chatService := NewChatService(upgrader, NewUserRegistry())
	// 	chatService.HandleConnection(w, r)
	// }))
	// defer server.Close()

	// // Create a WebSocket connection to the mock server
	// u := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	// conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	// if err != nil {
	// 	t.Fatalf("Error establishing WebSocket connection: %v", err)
	// }
	// defer conn.Close()
}
