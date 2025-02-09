package main

import (
	"log"
	"regexp"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zishang520/engine.io/v2/types"
	"github.com/zishang520/engine.io/v2/utils"
	"github.com/zishang520/socket.io/v2/socket"
)

func main() {
	r := echo.New()

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	opts := socket.DefaultServerOptions()
	opts.SetTransports(types.NewSet("polling", "websocket", "webtransport")) // added webtransport
	opts.SetAllowEIO3(false)

	server := socket.NewServer(nil, opts)

	r.Any("/socket.io/", echo.WrapHandler(server.ServeHandler(opts)))

	SetupSocketHandler(server)

	err := r.Start(":8080")
	if err != nil {
		log.Fatalf("router Start: %s\n", err)
	}
}

func SetupSocketHandler(server *socket.Server) {
	namespace := server.Of(regexp.MustCompile(`/jackpot`), nil)

	namespace.On("connection", func(clients ...interface{}) {
		println("HEEY")
		utils.Log().Success("On Connect")
		client := clients[0].(*socket.Socket)

		client.On("message", func(msg ...any) {
			log.Println("Received message:", msg)
			client.Emit("message back", msg...)
		})
	})
}
