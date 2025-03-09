package handlers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"log-analyzer/internal/parser"
	"log-analyzer/internal/stats"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	/*
		This function handles the statistics page and renders the statistics template.
	*/
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		http.Error(w, "File not specified", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("uploads", fileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	lines := string(content)
	stat := stats.NewStatistic()
	for _, line := range parser.SplitLines(lines) {
		ip, _, request, code, err := parser.ApachParseLine(line)
		if err == nil {
			stat.AddStatistics(ip, code, request)
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	tmpl.Execute(w, stat)
}
