package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Config struct {
	MQTTBroker   string
	MQTTClientID string
	MQTTTopic    string
	HTTPAddress  string
}

type Task struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var (
	tasks []Task
	mu    sync.Mutex
)

func loadConfig(path string) (*Config, error) {
	_ = path
	return &Config{
		MQTTBroker:   "tcp://test.mosquitto.org:1883",
		MQTTClientID: "go-mqtt-sample",
		MQTTTopic:    "aruzhan/tasks",
		HTTPAddress:  ":8080",
	}, nil
}

// MQTT publisher handler (без изменений)
func publishHandler(client mqtt.Client, cfg *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("msg")
		if text == "" {
			text = fmt.Sprintf("Hello MQTT! Time: %s", time.Now())
		}
		if client != nil {
			token := client.Publish(cfg.MQTTTopic, 0, false, text)
			token.Wait()
		}
		fmt.Fprintf(w, "Published message: %s\n", text)
		log.Printf("Published message: %s", text)
	}
}

// REST API для задач
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	case "POST":
		var t Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		mu.Lock()
		tasks = append(tasks, t)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", 405)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	idxStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	idx, err := strconv.Atoi(idxStr)
	if err != nil || idx < 0 {
		http.Error(w, "Invalid index", 400)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if idx >= len(tasks) {
		http.Error(w, "Not found", 404)
		return
	}
	switch r.Method {
	case "PATCH":
		var patch struct{ Done bool }
		if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		tasks[idx].Done = patch.Done
		w.WriteHeader(204)
	case "DELETE":
		tasks = append(tasks[:idx], tasks[idx+1:]...)
		w.WriteHeader(204)
	default:
		http.Error(w, "Method not allowed", 405)
	}
}

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()

	cfg, err := loadConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message on topic %s: %s\n", msg.Topic(), msg.Payload())
	}

	opts := mqtt.NewClientOptions().
		AddBroker(cfg.MQTTBroker).
		SetClientID(cfg.MQTTClientID).
		SetDefaultPublishHandler(messagePubHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	defer client.Disconnect(250)

	if token := client.Subscribe(cfg.MQTTTopic, 1, messagePubHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	log.Printf("Subscribed to topic: %s", cfg.MQTTTopic)

	// 1. Сервируем фронтенд (index.html и статику) по адресу /
	// Проверь путь: если frontend рядом с папкой cmd, путь "../frontend"
	http.Handle("/", http.FileServer(http.Dir("../frontend")))

	// 2. Обработчики API
    http.Handle("/", http.FileServer(http.Dir("../../frontend")))
	http.HandleFunc("/publish", publishHandler(client, cfg))
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)

	log.Printf("listening on %s", cfg.HTTPAddress)
	if err := http.ListenAndServe(cfg.HTTPAddress, nil); err != nil {
		log.Fatal(err)
	}
}
