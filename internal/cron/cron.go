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

	fmt.Printf("Cronjob started to wallet: %v", order.CustomerWalletAddress)

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

	fmt.Println("Cronjob timed out")

	err := wsClient.SendMessage("checkWallet", map[string]string{"error": "Check wallet timeout"})

	if err != nil {
		fmt.Println("Failed to send JSON via WebSocket when timeout: ", err)
		return
	}
}
