package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gmwallet-connect-go/cmd/dapp-server/utils"
	"gmwallet-connect-go/internal/api/ws_handler"
	"gmwallet-connect-go/internal/ws"
	"log"
	"net/http"
	"time"
)

// websocket server
func main() {
	e := echo.New()

	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("5M"))
	e.Validator = utils.NewCustomValidator()

	hub := ws.NewHub()
	wsHandler := ws_handler.NewHandler(hub)
	go hub.Run()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Wallet Connect API")
	})

	g := e.Group("/api")
	g.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Minute,
	}))
	// create room - use DApp or BApp
	g.GET("/active", wsHandler.ActiveRoom)
	// join room - use Gm wallet
	e.GET("/join/:hash", wsHandler.JoinRoom)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Error running API: %v", err)
	}
}
