package main

import (
	"context"
	"log"
	"os"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qoparu/tdl/internal/config"
	"github.com/qoparu/tdl/internal/mq"
	"github.com/qoparu/tdl/internal/task"
	"github.com/rs/cors"
)

// ApiServer держит store и брокер
type ApiServer struct {
	store  task.Store
	broker mq.Broker
	topic  string
}

func main() {
	// Загружаем конфиг
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключаемся к БД
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	dbpool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()
	_, err = dbpool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			text TEXT NOT NULL,
			done BOOLEAN DEFAULT FALSE
		);
	`)
	if err != nil {
		log.Fatalf("Unable to create tasks table: %v\n", err)
	}

	// Брокер MQTT
	broker, err := mq.NewMQTTBroker(cfg.MQTT.Broker, cfg.MQTT.ClientID)
	if err != nil {
		log.Fatalf("Can't connect to MQTT broker: %v", err)
	}
	defer broker.Close()

	api := &ApiServer{
		store:  task.NewPostgresStore(dbpool),
		broker: broker,
		topic:  cfg.MQTT.Topic,
	}

	r := chi.NewRouter()
	r.Get("/tasks", api.handleGetTasks)
	r.Post("/tasks", api.handleCreateTask)
	r.Put("/tasks/{id}", api.handleUpdateTask)
	r.Delete("/tasks/{id}", api.handleDeleteTask)

	handler := cors.Default().Handler(r)

	log.Println("Starting server on", cfg.HTTP.Address)
	if err := http.ListenAndServe(cfg.HTTP.Address, handler); err != nil {
		log.Fatal(err)
	}
}

// --- HTTP Handlers ---

func (s *ApiServer) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := s.store.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, tasks, http.StatusOK)
}

func (s *ApiServer) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var t task.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	createdTask, err := s.store.Create(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, createdTask, http.StatusCreated)
	s.publishEvent("created", createdTask)
}

func (s *ApiServer) handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var t task.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	updatedTask, err := s.store.Update(id, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, updatedTask, http.StatusOK)
	s.publishEvent("updated", updatedTask)
}

func (s *ApiServer) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = s.store.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	s.publishEvent("deleted", task.Task{ID: id})
}

func writeJSON(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

// --- MQTT publish ---

func (s *ApiServer) publishEvent(eventType string, t task.Task) {
	event := struct {
		Type string    `json:"type"`
		Task task.Task `json:"task"`
	}{
		Type: eventType,
		Task: t,
	}
	payload, _ := json.Marshal(event)
	_ = s.broker.Publish(s.topic, payload)
}
