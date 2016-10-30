package main

import (
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/file/{name}", file)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Download a file based on URL parameters
func file(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		fileName := vars["name"]
		filePath := "./files/" + fileName

		// Check if file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.Error(w, "File Does Not Exist", http.StatusBadRequest)
			return
		}

		// Set response headers for zip file
		w.Header().Set("Content-Type", "applicaiton/zip")
		w.Header().Set("Content-Disposition", "attachment; filename='"+fileName+"'")

		http.ServeFile(w, r, filePath)
	}
}
