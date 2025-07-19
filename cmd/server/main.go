package main

import (
	"flag"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/example/tdl/internal/api"
	"github.com/example/tdl/internal/config"
	api "github.com/example/tdl/internal/http"
	"github.com/example/tdl/internal/mq"
	"github.com/example/tdl/internal/task"
)

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	store := task.NewInMemoryStore()

	opts := mqtt.NewClientOptions().AddBroker(cfg.MQTT.Broker).SetClientID(cfg.MQTT.ClientID)
	broker := mq.NewMQTTBroker(opts)
	if err := broker.Connect(); err != nil {
		log.Fatal(err)
	}

	srv := &api.Server{Store: store, Broker: broker, Topic: cfg.MQTT.Topic}

	log.Printf("listening on %s", cfg.HTTP.Address)
	if err := srv.Serve(cfg.HTTP.Address); err != nil {
