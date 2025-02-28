package web

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"log-analyzer/internal/parser"
	"log-analyzer/internal/stats"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	/*
		This function handles the home page and renders the index page.
	*/
	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	tmpl.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	/*
		This function handles the file upload and redirects to the statistics page.
	*/
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the file
	file, handler, err := r.FormFile("logfile")
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file locally
	savePath := filepath.Join("uploads", handler.Filename)
	out, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Read file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	out.Write(content)

	// Redirect to the statistics page
	http.Redirect(w, r, "/stats?file="+handler.Filename, http.StatusSeeOther)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	/*
		This function handles the statistics page and renders the statistics template.
	*/
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		http.Error(w, "File not specified", http.StatusBadRequest)
		return
	}

	// Read file
	filePath := filepath.Join("uploads", fileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Process log data
	lines := string(content)
	stat := stats.NewStatistic()
	for _, line := range parser.SplitLines(lines) {
		ip, _, request, code, err := parser.ApachParseLine(line)
		if err == nil {
			stat.AddStatistics(ip, code, request)
		}
	}

	// Render statistics page
	tmpl := template.Must(template.ParseFiles("templates/stats.gohtml"))
	tmpl.Execute(w, stat)
}
