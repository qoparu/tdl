package mq

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Broker interface {
	Publish(topic string, payload []byte) error
}

type MQTTBroker struct {
	client mqtt.Client
}

func NewMQTTBroker(opts *mqtt.ClientOptions) *MQTTBroker {
	client := mqtt.NewClient(opts)
	return &MQTTBroker{client: client}
}

func (b *MQTTBroker) Connect() error {
	if token := b.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (b *MQTTBroker) Publish(topic string, payload []byte) error {
	token := b.client.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}
