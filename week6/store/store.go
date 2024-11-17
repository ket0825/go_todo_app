// `Store` package is responsible for storing and managing tasks. It provides a simple in-memory store for tasks. It has a `TaskStore` struct that holds a map of tasks. It also provides methods to add, get, and list all tasks.

package store

import (
	"errors"
	"fmt"

	"github.com/ket0825/go_todo_app/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) Get(id entity.TaskID) (*entity.Task, error) {
	if ts, ok := ts.Tasks[id]; ok {
		return ts, nil
	}
	return nil, ErrNotFound
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t // why i-1?
		fmt.Printf("i: %d, t: %v\n", i, t)
	}
	return tasks
}
