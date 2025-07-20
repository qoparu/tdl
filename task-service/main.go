package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
)

type Task struct {
    Text   string `json:"text"`
    UserID string `json:"user_id"`
}

func main() {
    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "method not allowed", 405)
            return
        }
        var t Task
        if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
            http.Error(w, "bad request", 400)
            return
        }

        // RPC-запрос в user-service
        resp, err := http.Get(fmt.Sprintf("http://localhost:8082/user/%s", t.UserID))
        if err != nil {
            http.Error(w, "user-service unavailable", 500)
            return
        }
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        var check map[string]bool
        json.Unmarshal(body, &check)
        if !check["exists"] {
            http.Error(w, "user not found", 403)
            return
        }

        log.Printf("Task created for user %s: %s", t.UserID, t.Text)
        w.WriteHeader(http.StatusCreated)
    })

    log.Println("task-service listening on :8083")
    log.Fatal(http.ListenAndServe(":8083", nil))
}
