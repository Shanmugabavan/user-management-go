package domain

import (
	"context"
	"user-management/internal/db"

	"github.com/google/uuid"
)

const (
	tableUser = "users"
)

type User struct {
	UserId    uuid.UUID  `json:"userId" validate:"required"`
	FirstName string     `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string     `json:"lastName" validate:"required,min=2,max=50"`
	Email     string     `json:"email" validate:"required,email"`
	Phone     string     `json:"phone" validate:"required,e164"`
	Age       int        `json:"age" validate:"required,gt=0"`
	Status    UserStatus `json:"status" validate:"omitempty,oneof=0 1"`
}

type UserStatus int

const (
	UserStatusActive UserStatus = iota
	UserStatusInactive
)

type UserRepository interface {
	Create(ctx context.Context, user *User) (db.CreateUserRow, error)
	GetAll(c context.Context) ([]User, error)
	GetById(c context.Context, id uuid.UUID) (User, error)
}
