package mq

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Broker interface {
	Publish(topic string, payload []byte) error
	Close() error
}

type MQTTBroker struct {
	client mqtt.Client
}

func NewMQTTBroker(broker, clientID string) (*MQTTBroker, error) {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &MQTTBroker{client: client}, nil
}

func (b *MQTTBroker) Publish(topic string, payload []byte) error {
	token := b.client.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}

func (b *MQTTBroker) Close() error {
	b.client.Disconnect(250)
	return nil
}
