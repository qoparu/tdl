package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qoparu/tdl/internal/task"
)

// Фейковый брокер для теста
type fakeBroker struct {
	published [][]byte
}

func (f *fakeBroker) Publish(topic string, payload []byte) error {
	f.published = append(f.published, payload)
	return nil
}

func (f *fakeBroker) Close() error { return nil }

func setupTestServer() (*ApiServer, *fakeBroker) {
	store := task.NewInMemoryStore()
	broker := &fakeBroker{}
	return &ApiServer{
		store:  store,
		broker: broker,
		topic:  "test-topic",
	}, broker
}

func TestCreateTask_PublishesMQTT(t *testing.T) {
	api, broker := setupTestServer()

	rr := httptest.NewRecorder()
	body, _ := json.Marshal(task.Task{Text: "Test"})
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
	api.handleCreateTask(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rr.Code)
	}

	if len(broker.published) != 1 {
		t.Fatalf("expected 1 MQTT publish, got %d", len(broker.published))
	}

	var event struct {
		Type string    `json:"type"`
		Task task.Task `json:"task"`
	}
	json.Unmarshal(broker.published[0], &event)
	if event.Type != "created" {
		t.Errorf("expected event type 'created', got %s", event.Type)
	}
	if event.Task.Text != "Test" {
		t.Errorf("expected task text 'Test', got %s", event.Task.Text)
	}
}
