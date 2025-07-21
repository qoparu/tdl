package task

import (
	"testing"
)

func TestInMemoryStore_CreateAndGet(t *testing.T) {
	store := NewInMemoryStore()
	created, err := store.Create(Task{Text: "hello"})
	if err != nil {
		t.Fatalf("Create() error: %v", err)
	}
	if created.ID == 0 {
		t.Errorf("expected ID > 0, got %d", created.ID)
	}
	got, err := store.Get(created.ID)
	if err != nil {
		t.Fatalf("Get() error: %v", err)
	}
	if got.Text != "hello" {
		t.Errorf("expected Text 'hello', got '%s'", got.Text)
	}
}

func TestInMemoryStore_List(t *testing.T) {
	store := NewInMemoryStore()
	_, _ = store.Create(Task{Text: "a"})
	_, _ = store.Create(Task{Text: "b"})
	list, err := store.List()
	if err != nil {
		t.Fatalf("List() error: %v", err)
	}
	if len(list) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(list))
	}
}

func TestInMemoryStore_Update(t *testing.T) {
	store := NewInMemoryStore()
	task, _ := store.Create(Task{Text: "x"})
	updated, err := store.Update(task.ID, Task{Text: "y", Done: true})
	if err != nil {
		t.Fatalf("Update() error: %v", err)
	}
	if updated.Text != "y" || !updated.Done {
		t.Errorf("Update() failed, got: %+v", updated)
	}
	// Проверка что реально изменилось в хранилище
	got, _ := store.Get(task.ID)
	if got.Text != "y" || !got.Done {
		t.Errorf("task in store not updated, got: %+v", got)
	}
}

func TestInMemoryStore_Delete(t *testing.T) {
	store := NewInMemoryStore()
	task, _ := store.Create(Task{Text: "todelete"})
	err := store.Delete(task.ID)
	if err != nil {
		t.Fatalf("Delete() error: %v", err)
	}
	_, err = store.Get(task.ID)
	if err == nil {
		t.Errorf("expected error for deleted task, got nil")
	}
}

func TestInMemoryStore_GetNotFound(t *testing.T) {
	store := NewInMemoryStore()
	_, err := store.Get(999)
	if err == nil {
		t.Errorf("expected error for missing task, got nil")
	}
}

func TestInMemoryStore_UpdateNotFound(t *testing.T) {
	store := NewInMemoryStore()
	_, err := store.Update(123, Task{Text: "none"})
	if err == nil {
		t.Errorf("expected error for update missing, got nil")
	}
}

func TestInMemoryStore_DeleteNotFound(t *testing.T) {
	store := NewInMemoryStore()
	err := store.Delete(123)
	if err == nil {
		t.Errorf("expected error for delete missing, got nil")
	}
}