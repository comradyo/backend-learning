package main

import (
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/zishang520/engine.io/v2/types"
	"github.com/zishang520/engine.io/v2/utils"
	"github.com/zishang520/socket.io/v2/socket"
)

func main() {
	r := gin.Default()

	opts := socket.DefaultServerOptions()
	opts.SetTransports(types.NewSet("polling", "websocket", "webtransport")) // added webtransport
	opts.SetAllowEIO3(true)
	opts.SetCors(&types.Cors{
		Origin:      "*",
		Credentials: false,
	})

	server := socket.NewServer(nil, opts)

	r.GET("/socket.io/*any", gin.WrapH(server.ServeHandler(opts)))
	r.POST("/*any", gin.WrapH(server.ServeHandler(opts)))

	SetupSocketHandler(server)

	r.Run()
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
