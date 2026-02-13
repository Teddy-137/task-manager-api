package service

import (
	"errors"

	"github.com/teddy-137/task_manager_api/internal/domain"
	"github.com/teddy-137/task_manager_api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
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
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	return s.userRepo.Store(user)
}

func (s *userService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.ID)
}

func (s *userService) UpdateUser(id uint, user *domain.User) error {
	if user.Username == "" {
		return errors.New("invalid username")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	return s.userRepo.Update(id, user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
