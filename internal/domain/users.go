package domain

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Tasks     []Task    `json:"tasks, omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Fetch() ([]User, error)
	GetByID(id uint) (User, error)
	GetByUsername(username string) (User, error)
	Store(user *User) error
}

type UserService interface {
	GetAllUsers() ([]User, error)
	CreateUser(user *User) error
	Login(username, password string) (string, error)
}
