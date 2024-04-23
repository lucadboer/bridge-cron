package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/enrique/cron-bridge/internal"
	"github.com/enrique/cron-bridge/internal/cron"
	"github.com/enrique/cron-bridge/internal/model"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Order = model.Order

var (
	ordersQueue = make(chan Order)
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file, assuming environment variables are set")
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(internal.RateLimiterMiddleware)

	e.POST("/cron", func(c echo.Context) error {
		var order Order
		if err := c.Bind(&order); err != nil {
			log.Println("Error decoding order:", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid order format")
		}

		ordersQueue <- order

		var response = map[string]string{
			"message": "Order received",
		}

		return c.JSON(http.StatusOK, response)
	})

	port := os.Getenv("PORT")

	go handleOrders()

	println("Cron job and HTTP server started...")
	fmt.Println(ordersQueue)

	e.Logger.Fatal(e.Start(":" + port))
}

func handleOrders() {
	for order := range ordersQueue {
		go cron.ProcessOrder(order)
	}
}
