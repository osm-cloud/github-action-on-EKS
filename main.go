package main

import (
    "fmt"
    "net/http"
)

func health(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "hellow version 2 {status:OK}")
}

func dummy(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "BLUE")
}

func main() {
    http.HandleFunc("/health", health)
    http.HandleFunc("/v1/dummy", dummy)
    http.ListenAndServe(":80", nil)
}