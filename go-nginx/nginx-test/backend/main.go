package main

import (
	"fmt"
	"net/http"
	"os"
)

var port = os.Getenv("BACKEND_PORT")

func route(w http.ResponseWriter, r *http.Request) {
    fmt.Println("GOT REQUEST")
    fmt.Fprintf(w, "Hello, world! 😉 \n")
}

func main() {
    if len(port) == 0 {
        fmt.Printf("[Error] port is empty")
        return
    }
    //http.HandleFunc("/api/hello-world", route)
    fmt.Printf("Hello, World!\n")
    http.HandleFunc("/", route)
    http.ListenAndServe(":"+port, nil)
}
