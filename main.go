package main

import (
    "log"
    "net/http"

    "github.com/swer/myapp/handler"
)

func main() {
    http.HandleFunc("/", handler.Hello)
    log.Println("Запуск сервера на http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
 
