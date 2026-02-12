package service

import (
	"errors"
	"github.com/teddy-137/task_manager_api/internal/domain"
)

type taskService struct {
	taskRepo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) domain.TaskService {
	return &taskService{taskRepo: repo}
}

func (s *taskService) GetAllTasks() ([]domain.Task, error) {
	return s.taskRepo.Fetch()
}

func (s *taskService) CreateTask(task *domain.Task) error {
	if task.Title == "" {
		return errors.New("invalid task title.")
	}
	return s.taskRepo.Store(task)
}

func (s *taskService) GetTaskByID(id uint) (domain.Task, error) {
	return s.taskRepo.GetByID(id)
}
