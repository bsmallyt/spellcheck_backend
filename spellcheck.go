package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            fmt.Fprintf(w, "Hello")
        }
        //fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    fmt.Println("Server listening on 8090")
    http.ListenAndServe(":8090", nil)
}
