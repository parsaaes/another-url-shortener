package server

import "gitlab.com/parsaaes/another-url-shortener/config"

func StartServer() {
	config.Init(".")
}
