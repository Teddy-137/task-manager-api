package domain

import "time"

type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskRepository interface {
	Fetch() ([]Task, error)
	GetByID(id uint) (Task, error)
	Store(task *Task) error
}

type TaskService interface {
	GetAllTasks() ([]Task, error)
	CreateTask(task *Task) error
	GetTaskByID(id uint) (Task, error)
}
