package requests

import (
	"time"

	"github.com/Utsavch189/api_mysql/internal/utils"
	"github.com/google/uuid"
)

type User struct {
	Id        string    `json:"id"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}

type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Constructor for user
func NewUser(username, password string) (*User, error) {
	hashedPassword, err := utils.HashString(password)
	if err != nil {
		return &User{}, err
	}
	uid := uuid.New()
	return &User{
		Id:        uid.String(),
		UserName:  username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		IsActive:  true,
	}, nil
}
