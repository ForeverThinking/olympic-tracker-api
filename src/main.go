package main

import (
    "fmt"
    "net/http"
    "os"
    "log"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    test := os.Getenv("TEST")

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, test)
    })

    http.ListenAndServe(":8080", nil)
}
