package web

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	/*
		This fonction is used to start the server and handle the different routes
	*/
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/upload", UploadHandler)
	http.HandleFunc("/stats", StatsHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
