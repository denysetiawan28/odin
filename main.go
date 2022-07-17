package main

import (
	"odin/src/server"
	"odin/src/server/container"
)

func main() {
	server.StartHttpServer(container.IntializeContainer())
}
