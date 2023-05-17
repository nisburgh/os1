package main

import (
        "fmt"
        "log"
        "net/http"
)

func main() {
        http.HandleFunc("/", handler)
        log.Printf("Server started and listening on 0.0.0.0:8080...")
        log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
        log.Printf("Ping from %s", r.RemoteAddr)
        fmt.Fprintln(w, "Hello Kubernetes Beginners!")
}
