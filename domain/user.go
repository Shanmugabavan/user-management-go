package domain

import (
	"context"

	"github.com/google/uuid"
)

const (
	tableUser = "users"
)

type User struct {
	UserId    uuid.UUID  `json:"userId"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Age       int        `json:"age"`
	Status    UserStatus `json:"status"`
}

type CreateUserRequest struct {
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Age       int        `json:"age"`
	Status    UserStatus `json:"status"`
}

type UserStatus int

const (
	UserStatusActive UserStatus = iota
	UserStatusInactive
)

type UserRepository interface {
	Create(ctx context.Context, user *User) (User, error)
}
