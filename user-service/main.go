package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strings"
)

var users = map[string]bool{
    "alice": true,
    "bob":   true,
    "meiirzhan": true,
}

func main() {
    http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
        parts := strings.Split(r.URL.Path, "/")
        if len(parts) < 3 {
            http.Error(w, "not found", 404)
            return
        }
        userID := parts[2]
        _, exists := users[userID]
        json.NewEncoder(w).Encode(map[string]bool{"exists": exists})
    })
    log.Println("user-service listening on :8082")
    log.Fatal(http.ListenAndServe(":8082", nil))
}
