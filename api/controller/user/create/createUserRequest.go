package create

import "user-management/domain"

type CreateUserRequest struct {
	FirstName string            `json:"firstName"`
	LastName  string            `json:"lastName"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	Age       int               `json:"age"`
	Status    domain.UserStatus `json:"status"`
}
