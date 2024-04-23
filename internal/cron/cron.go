package cron

import (
	"fmt"
	"os"
	"time"

	"github.com/enrique/cron-bridge/internal/blockchain/theta"
	"github.com/enrique/cron-bridge/internal/model"
	"github.com/enrique/cron-bridge/internal/services"
)

var (
	path = "/ws"
)

func ProcessOrder(order model.Order) {
	wsUrl := os.Getenv("WEBSOCKET_URL")
	wsClient := services.NewWebSocketClient("ws", wsUrl, path, order.CustomerWalletAddress)

	deadline := time.Now().Add(1 * time.Hour)

	for time.Now().Before(deadline) {
		if theta.CheckWalletChange(order.CustomerWalletAddress) {
			err := wsClient.SendMessage("checkWallet", order)

			if err != nil {
				fmt.Println("Failed to send JSON via WebSocket:", err)
				return
			}
		}
		time.Sleep(30 * time.Second)
	}

	wsClient.SendMessage("checkWallet", map[string]string{"error": "Check wallet timeout"})
}

// todo fix when check wallet timeout