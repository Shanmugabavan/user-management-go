package controller

import (
	"encoding/json"
	"net/http"
	"user-management/bootstrap"
	"user-management/domain"

	"github.com/google/uuid"
)

type UserController struct {
	domain.UserRepository
	Env *bootstrap.Env
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest domain.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		http.Error(w, jsonError(err.Error()), http.StatusBadRequest)
		return
	}

	user := domain.User{
		UserId:    uuid.New(),
		FirstName: createUserRequest.FirstName,
		LastName:  createUserRequest.LastName,
		Email:     createUserRequest.Email,
		Phone:     createUserRequest.Phone,
		Age:       createUserRequest.Age,
	}

	_, err2 := u.Create(r.Context(), &user)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
