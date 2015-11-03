package main

import (
    "log"
    "net/http"
)

func main() {
    rrr := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", rrr))
}
