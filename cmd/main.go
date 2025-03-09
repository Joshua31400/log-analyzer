package main

import (
	"log-analyzer/internal/analyzer"
	server "log-analyzer/internal/web/server"
)

func main() {
	server.StartServer()
	analyzer.Run()
}
