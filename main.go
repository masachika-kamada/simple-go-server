package main

import (
    "fmt"
    "net/http"
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

func contactHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        // POSTリクエストの場合の処理
        r.ParseForm()
        message := r.FormValue("message")
        fmt.Fprintf(w, "Received message: %s", message)
    } else {
        // GETリクエストの場合の処理（フォームを表示）
        fmt.Fprintf(w, `<h1>Contact Us</h1>
        <form action="/contact" method="post">
            <label for="message">Message:</label>
            <input type="text" id="message" name="message">
            <input type="submit" value="Send">
        </form>`)
    }
}
