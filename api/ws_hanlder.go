package main

import (
	"canaanadvisors-test/proto/notification"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow any origin for simplicity
	},
}

type WebSocketHandler interface {
	SendMessage(context.Context, *notification.MessageSendRequest) (*notification.MessageSendResponse, error)
	ReceiveMessage(notification.WebSocketService_ReceiveMessageServer) error
}

type WebSocketController struct {
	notification.UnimplementedWebSocketServiceServer
	connections map[string]*websocket.Conn
	mu          sync.RWMutex
	logger *zap.Logger
}

func NewWebSocketHandler(logger *zap.Logger) WebSocketHandler {
	return &WebSocketController{
		connections: make(map[string]*websocket.Conn),
		logger: logger,
	}
}
func (wsc *WebSocketController) SendMessage(ctx context.Context, request *notification.MessageSendRequest) (
	*notification.MessageSendResponse, error) {
	clientID := request.ClientId
	message := request.Message
	err := wsc.SendMessageToClient(clientID, message)
	if err != nil {
		return nil, err
	}

	response := &notification.MessageSendResponse{
		Status: "send successfully",
	}

	return response, nil
}


func (wsc *WebSocketController) ReceiveMessage(stream notification.WebSocketService_ReceiveMessageServer) error {
	request, err := stream.Recv()
	if err != nil {
		return err
	}
	clientID := request.ClientId

	wsc.mu.RLock()
	conn, ok := wsc.connections[clientID]
	wsc.mu.RUnlock()

	if !ok {
		return fmt.Errorf("client connection not found: %s", clientID)
	}

	// Set up a goroutine to receive WebSocket messages
	go func() {
		for {
			_, rawMsg, e := conn.ReadMessage()
			if e != nil {
				wsc.logger.Error(e.Error())
				return
			}

			// Process rawMsg if needed

			response := &notification.MessageReceiveResponse{
				Message: string(rawMsg),
			}

			e = stream.Send(response)
			if e != nil {
				wsc.logger.Error(e.Error())
				return
			}
		}
	}()
	// Block the current goroutine to keep the stream open
	select {}
}
func (wsc *WebSocketController) SendMessageToClient(clientID, message string) error {
	wsc.mu.RLock()
	defer wsc.mu.RUnlock()

	conn, ok := wsc.connections[clientID]
	if !ok {
		return fmt.Errorf("client connection not found: %s", clientID)
	}

	return conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func (wsc *WebSocketController) AddConnection(clientID string, conn *websocket.Conn) {
	wsc.mu.Lock()
	defer wsc.mu.Unlock()

	wsc.connections[clientID] = conn
}

func (wsc *WebSocketController) RemoveConnection(clientID string) {
	wsc.mu.Lock()
	defer wsc.mu.Unlock()

	delete(wsc.connections, clientID)
}

func (wsc *WebSocketController) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		wsc.logger.Error(err.Error())
		return
	}
	defer conn.Close()

	clientID := r.URL.Query().Get("client_id")
	if clientID == "" {
		wsc.logger.Error(err.Error())
		return
	}

	wsc.AddConnection(clientID, conn)
	defer wsc.RemoveConnection(clientID)

	// ... rest of the handleWebSocket code
}