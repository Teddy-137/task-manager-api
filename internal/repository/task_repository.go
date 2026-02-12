package repository

import (
	"github.com/teddy-137/task_manager_api/internal/domain"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Fetch() ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetByID(id uint) (domain.Task, error) {
	var task domain.Task
	err := r.db.First(&task, id).Error
	return task, err
}

func (r *taskRepository) Store(task *domain.Task) error {
	err := r.db.Create(&task).Error
	return err
}
