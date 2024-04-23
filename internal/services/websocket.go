package services

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/enrique/cron-bridge/internal/model"
	"github.com/gorilla/websocket"
)

type Order = model.Order

type WebSocketClient struct {
	Conn *websocket.Conn
	Url  url.URL
}

func NewWebSocketClient(scheme, host, path, roomId string) *WebSocketClient {
	u := url.URL{Scheme: scheme, Host: host, Path: path}
	q := u.Query()
	q.Set("roomId", roomId)
	u.RawQuery = q.Encode()

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("error dialing websocket server: %v", err)
		return nil
	}
	return &WebSocketClient{
		Conn: conn,
		Url:  u,
	}
}

func (c *WebSocketClient) SendMessage(action string, data interface{}) error {
	message := map[string]interface{}{
		"action":       action,
		"orderDetails": data,
	}

	if err := c.Conn.WriteJSON(message); err != nil {
		log.Println("error writing JSON to WebSocket:", err)
	}

	return nil
}

func (client *WebSocketClient) WriteJSON(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = client.Conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		return err
	}
	return nil
}

// // ReceiveOrder listens for an order message, decodes it, and returns it
// func (c *WebSocketClient) ReceiveOrder() (Order, error) {
// 	var order Order
// 	_, message, err := c.ReadMessage()
// 	if err != nil {
// 		log.Println("error reading WebSocket message:", err)
// 		return order, err
// 	}

// 	err = json.Unmarshal(message, &order)
// 	if err != nil {
// 		log.Println("error unmarshaling order:", err)
// 		return order, err
// 	}

// 	return order, nil
// }
