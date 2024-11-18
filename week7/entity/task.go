package entity

import (
	"time"
)

type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "done"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID       TaskID     `json:"id"` // Tag: json에서 id를 key로 둠.
	Title    string     `json:"title"`
	Status   TaskStatus `json:"status"`
	Created  time.Time  `json:"created"`
	Modified time.Time  `json:"modified" db:"modified"`
}

type Tasks []*Task
