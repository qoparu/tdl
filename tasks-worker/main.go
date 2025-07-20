package main

import (
    "encoding/json"
    "log"

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
    msgs, err := ch.Consume("aruzhan_tasks", "", true, false, false, false, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("tasks-worker started")
    for msg := range msgs {
        var task Task
        if err := json.Unmarshal(msg.Body, &task); err != nil {
            log.Println("bad task json:", err)
            continue
        }
        log.Printf("Got new task: %s", task.Text)
        // Здесь можно добавить логику обработки задачи
    }
}
