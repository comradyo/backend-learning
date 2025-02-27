package main

import (
    "fmt"
    "net/http"
)

func route(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world! 😉 \n")
}

func main() {
    http.HandleFunc("/api/hello-world", route)
    http.ListenAndServe(":8888", nil)
}
