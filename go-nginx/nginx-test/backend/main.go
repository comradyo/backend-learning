package main

import (
    "fmt"
    "net/http"
)

func route(w http.ResponseWriter, r *http.Request) {
    fmt.Println("GOT REQUEST")
    fmt.Fprintf(w, "Hello, world! 😉 \n")
}

func main() {
    //http.HandleFunc("/api/hello-world", route)
    fmt.Printf("Hello, World!\n")
    http.HandleFunc("/", route)
    http.ListenAndServe(":8888", nil)
}
