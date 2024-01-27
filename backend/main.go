package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/contact", contactHandler)
    http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `<h1>Welcome to the Go Web Server</h1>
    <p><a href="/about">About Us</a></p>
    <p><a href="/contact">Contact Us</a></p>`)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    userAgent := r.Header.Get("User-Agent")
    fmt.Fprintf(w, "Your User-Agent is: %s", userAgent)
}

func setupCORS(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    setupCORS(&w)
    switch r.Method {
    case "GET":
        // GET リクエストの処理
        json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Go server"})
    case "POST":
        // POST リクエストの処理
        var data struct {
            Message string `json:"message"`
        }
        if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        fmt.Fprintf(w, "Received POST message: %s", data.Message)
    }
}
