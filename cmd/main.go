package main

import (
	"log-analyzer/internal/analyzer"
	"log-analyzer/internal/web"
)

func main() {
	web.StartServer()
	analyzer.Run()
}
