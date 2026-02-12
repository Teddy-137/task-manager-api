package service

import (
	"errors"

	"github.com/teddy-137/task_manager_api/internal/domain"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) GetAllUsers() ([]domain.User, error) {
	return s.userRepo.Fetch()
}

func (s *userService) CreateUser(user *domain.User) error {
	if user.Username == "" {
		return errors.New("invalid username")
	}
	return s.userRepo.Store(user)
}
