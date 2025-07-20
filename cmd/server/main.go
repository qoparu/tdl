package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Config struct {
	MQTTBroker   string
	MQTTClientID string
	MQTTTopic    string
	HTTPAddress  string
}

func loadConfig(path string) (*Config, error) {
	_ = path

	return &Config{
		MQTTBroker:   "tcp://test.mosquitto.org:1883",
		MQTTClientID: "go-mqtt-sample",
		MQTTTopic:    "aruzhan/tasks",
		HTTPAddress:  ":8080",
	}, nil
}

// Эта функция остаётся без изменений
func publishHandler(client mqtt.Client, cfg *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("msg")
		if text == "" {
			text = fmt.Sprintf("Hello MQTT! Time: %s", time.Now())
		}
		// Для тестов, где client может быть nil, нужна проверка
		if client != nil {
			token := client.Publish(cfg.MQTTTopic, 0, false, text)
			token.Wait()
		}
		fmt.Fprintf(w, "Published message: %s\n", text)
		log.Printf("Published message: %s", text)
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

	http.HandleFunc("/publish", publishHandler(client, cfg))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world! MQTT is connected: %v", client.IsConnected())
	})

	log.Printf("listening on %s", cfg.HTTPAddress)
	if err := http.ListenAndServe(cfg.HTTPAddress, nil); err != nil {
		log.Fatal(err)
	}
}