package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	/*
		This function handles the file upload and redirects to the home page.
	*/
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, handler, err := r.FormFile("logfile")
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	savePath := filepath.Join("uploads", handler.Filename)
	out, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	out.Write(content)

	http.Redirect(w, r, "/?file="+handler.Filename, http.StatusSeeOther)
}
