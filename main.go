package main

import (
	"gitlab.com/parsaaes/another-url-shortener/server"
)

func main() {
	server.SetupFront()
	server.StartServer()
}
