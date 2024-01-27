package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", messageHandler) // ルートパスで処理
    http.ListenAndServe(":8080", nil)
}

func setupCORS(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
    setupCORS(&w)
    if r.Method == "POST" {
        var data struct {
            Message string `json:"message"`
        }
        if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        processedMessage := strings.ToUpper(data.Message) // メッセージを大文字に変換
        fmt.Fprintf(w, processedMessage)
    }
}
