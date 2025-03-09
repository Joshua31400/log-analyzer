package web

import (
	"fmt"
	"log"
	"log-analyzer/internal/web/handlers"
	"net/http"
)

func StartServer() {
	/*
		This function starts the server and listens on port 8080.
	*/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/stats", handlers.StatsHandler)

	fmt.Println("Server is running on http://localhost:8080\nPress Ctrl+C to stop the server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
