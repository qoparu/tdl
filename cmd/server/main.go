package main

import (
	"flag"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	
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
		log.Fatal(err)
	}
}
 6 changes: 6 additions & 0 deletions6  
config.yaml
Viewed
Original file line number	Diff line number	Diff line change
@@ -0,0 +1,6 @@
http:
  address: ":8080"
mqtt:
  broker: "tcp://localhost:1883"
  client_id: "tdl-server"
  topic: "tasks"
