package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file", file)
	http.ListenAndServe(":8080", nil)
}

// Download data.zip
func file(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fileName := r.URL.Query().Get("name")
		filePath := "./files/" + fileName

		// Set response headers for zip file
		w.Header().Set("Content-Type", "applicaiton/zip")
		w.Header().Set("Content-Disposition", "attachment; filename='"+fileName+"'")

		http.ServeFile(w, r, filePath)
	}
}
