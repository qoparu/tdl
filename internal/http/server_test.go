package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/example/tdl/internal/task"
)

type fakeBroker struct{ msgs [][]byte }

func (f *fakeBroker) Publish(topic string, payload []byte) error {
	f.msgs = append(f.msgs, payload)
	return nil
}

func newServer() (*Server, *fakeBroker) {
	store := task.NewInMemoryStore()
	broker := &fakeBroker{}
	srv := &Server{Store: store, Broker: broker, Topic: "tasks"}
	srv.routes()
	return srv, broker
}

func TestCreateTask(t *testing.T) {
	srv, broker := newServer()

	body := bytes.NewBufferString(`{"text":"write tests"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	w := httptest.NewRecorder()
	srv.handleTasks(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201 got %d", w.Code)
	}
	if len(broker.msgs) != 1 {
		t.Fatalf("expected 1 message published")
	}
}
