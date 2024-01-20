package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)  // ルートURLへのリクエストを処理するハンドラを設定
    http.ListenAndServe(":8080", nil)  // サーバーを8080ポートで起動
}

// handler はHTTPリクエストを処理する関数
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Go Web Server!")  // レスポンスとしてメッセージを送信
}
