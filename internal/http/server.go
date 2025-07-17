package api

import (
	"encoding/json"
	"net/http"

	"github.com/example/tdl/internal/mq"
	"github.com/example/tdl/internal/task"
)

type Server struct {
	Store  task.Store
	Broker mq.Broker
	Topic  string
}

func (s *Server) routes() {
	http.HandleFunc("/tasks", s.handleTasks)
	http.HandleFunc("/tasks/", s.handleTask)
}

func (s *Server) Serve(addr string) error {
	s.routes()
	return http.ListenAndServe(addr, nil)
}

type event struct {
	Type string    `json:"type"`
	Task task.Task `json:"task"`
}

func (s *Server) publish(e event) {
	b, _ := json.Marshal(e)
	_ = s.Broker.Publish(s.Topic, b)
}

func (s *Server) handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks := s.Store.List()
		respond(w, tasks, http.StatusOK)
	case http.MethodPost:
		var t task.Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		t = s.Store.Create(t)
		respond(w, t, http.StatusCreated)
		s.publish(event{Type: "created", Task: t})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/tasks/"):]
	switch r.Method {
	case http.MethodPut:
		var t task.Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updated, ok := s.Store.Update(id, t)
		if !ok {
			http.NotFound(w, r)
			return
		}
		respond(w, updated, http.StatusOK)
		s.publish(event{Type: "updated", Task: updated})
	case http.MethodDelete:
		if !s.Store.Delete(id) {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		s.publish(event{Type: "deleted", Task: task.Task{ID: id}})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func respond(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
