package main

import (
	"voyager2/src/server"
	"voyager2/src/server/container"
)

func main() {
	server.StartHttpServer(container.IntializeContainer())
}
