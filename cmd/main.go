package main

import (
	"structured-logging/internal/server"
	"structured-logging/internal/service"
	"structured-logging/utils/log"
)

func main() {
	logger := log.NewLogger()
	svc := service.NewService()

	server.RunHTTPServer(":8081",svc,logger)
}
