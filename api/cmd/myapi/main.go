package main

import (
	"api/api"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func handler1(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from server 1")
}

func handler2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from server 2")
}

func main() {
    // Start Gin server
    port := os.Getenv("API_PORT")

    // Verify if port range or format is correct
    _, err := strconv.ParseUint(port, 10, 16)
    if err != nil {
	    log.Printf("Conversion error or invalid port number: %q", err)
    }

    api.Router(port)
    // First server
    mux1 := http.NewServeMux()
    mux1.HandleFunc("/", handler1)

    // Second server
    mux2 := http.NewServeMux()
    mux2.HandleFunc("/", handler2)

    go func() {
        fmt.Println("Server 1 running on :8081")
        err := http.ListenAndServe(":8081", mux1)
        if err != nil {
            panic(err)
        }
    }()

    /*
    fmt.Println("Server 2 running on :9090")
    err := http.ListenAndServe(":9090", mux2)
    if err != nil {
        panic(err)
    }
    */
}

