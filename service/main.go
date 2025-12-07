package main

import (
    "fmt"
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    http.Handle("/metrics", promhttp.Handler())
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Service Running")
    })
    http.ListenAndServe(":8080", nil)
}
