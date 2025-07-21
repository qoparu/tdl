package task

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	db *pgxpool.Pool
}

func NewPostgresStore(db *pgxpool.Pool) *PostgresStore {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) List() ([]Task, error) {
	rows, err := s.db.Query(context.Background(), "SELECT id, text, done, clock FROM tasks ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Text, &t.Done, &t.Clock); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (s *PostgresStore) Create(t Task) (Task, error) {
	var id int
	err := s.db.QueryRow(context.Background(), "INSERT INTO tasks (text, done, clock) VALUES ($1, $2, $3) RETURNING id", t.Text, t.Done, t.Clock).Scan(&id)
	if err != nil {
		return Task{}, err
	}
	t.ID = id
	return t, nil
}

func (s *PostgresStore) Get(id int) (Task, error) {
	var t Task
	err := s.db.QueryRow(context.Background(), "SELECT id, text, done, clock FROM tasks WHERE id = $1", id).Scan(&t.ID, &t.Text, &t.Done, &t.Clock)
	return t, err
}

func (s *PostgresStore) Update(id int, t Task) (Task, error) {
	_, err := s.db.Exec(context.Background(), "UPDATE tasks SET text = $1, done = $2, clock = $3 WHERE id = $4", t.Text, t.Done, t.Clock, id)
	if err != nil {
		return Task{}, err
	}
	t.ID = id
	return t, nil
}

func (s *PostgresStore) Delete(id int) error {
	_, err := s.db.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	return err
}