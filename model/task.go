package model

type Task struct{}

func NewTask() *Task {
	task := &Task{}
	return task
}

func (t *Task) SaveTask() {}
