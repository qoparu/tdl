package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"

    mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Config struct {
    MQTTBroker   string
    MQTTClientID string
    MQTTTopic    string
    HTTPAddress  string
}

func loadConfig(path string) (*Config, error) {

    return &Config{
        MQTTBroker:   "tcp://test.mosquitto.org:1883",
        MQTTClientID: "go-mqtt-sample",
        MQTTTopic:    "test/topic",
        HTTPAddress:  ":8080",
    }, nil
}

func main() {
    cfgPath := flag.String("config", "config.yaml", "path to config")
    flag.Parse()

    cfg, err := loadConfig(*cfgPath)
    if err != nil {
        log.Fatal(err)
    }

    // Set up MQTT client options
    opts := mqtt.NewClientOptions().AddBroker(cfg.MQTTBroker).SetClientID(cfg.MQTTClientID)
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }
    defer client.Disconnect(250)

    // Simple HTTP handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, world! MQTT is connected: %v", client.IsConnected())
    })

    log.Printf("listening on %s", cfg.HTTPAddress)
    if err := http.ListenAndServe(cfg.HTTPAddress, nil); err != nil {
        log.Fatal(err)
    }
}
