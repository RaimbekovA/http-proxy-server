package main

import (
    "log"
    "net/http"
    "http-proxy-server/handlers"
)

func main() {
    http.HandleFunc("/proxy", handlers.ProxyHandler)
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
