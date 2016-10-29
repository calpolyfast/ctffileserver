package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/data.zip", data)
    http.ListenAndServe(":8080", nil)
}

func data(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        http.ServeFile(w, r, "./files/data.zip")
    }
}
