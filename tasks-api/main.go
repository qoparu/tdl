package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/streadway/amqp"
)

type Task struct {
    Text string `json:"text"`
}

func main() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    ch, err := conn.Channel()
    if err != nil {
        log.Fatal(err)
    }
    defer ch.Close()
    _, err = ch.QueueDeclare("aruzhan_tasks", true, false, false, false, nil)
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
        text := r.URL.Query().Get("msg")
        if text == "" {
            http.Error(w, "missing msg", http.StatusBadRequest)
            return
        }
        task := Task{Text: text}
        body, _ := json.Marshal(task)
        err = ch.Publish("", "aruzhan_tasks", false, false, amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
        if err != nil {
            http.Error(w, "failed to publish", 500)
            return
        }
        log.Printf("Published: %+v", task)
        w.WriteHeader(http.StatusAccepted)
    })

    log.Println("tasks-api listening on :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
