package main

import (
	"SleepWalker/Golang-Chat/internal/router"
	"SleepWalker/Golang-Chat/internal/ws"
)

func main() {
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(wsHandler)
	router.Start("0.0.0.0:8080")
}
