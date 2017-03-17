package main

import (
    "log"
    "net/http"

    "./handlers"

    "github.com/arschles/go-in-5-minutes/episode1/storage"
)

func main() {
    db := storage.NewInMemoryDB()
    mux := http.NewServeMux()
    mux.Handle("/get", handlers.GetKey(db))
    mux.Handle("/set", handlers.PutKey(db))
    log.Printf("serving on port 8080")
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
