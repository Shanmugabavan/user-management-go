package user

import (
	"encoding/json"
	"net/http"
	"user-management/api/controller"
	"user-management/api/controller/user/create"
	"user-management/bootstrap"
	"user-management/domain"

	"github.com/google/uuid"
)

type UserController struct {
	domain.UserRepository
	Env *bootstrap.Env
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest create.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		http.Error(w, controller.JsonError(err.Error()), http.StatusBadRequest)
		return
	}

	user := domain.User{
		UserId:    uuid.New(),
		FirstName: createUserRequest.FirstName,
		LastName:  createUserRequest.LastName,
		Email:     createUserRequest.Email,
		Phone:     createUserRequest.Phone,
		Age:       createUserRequest.Age,
		Status:    createUserRequest.Status,
	}

	createdUser, err2 := u.Create(r.Context(), &user)

	createUserResponse := create.CreateUserResponse{
		UserID: createdUser.UserID,
		Email:  createdUser.Email,
		Status: createdUser.Status,
	}
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(createUserResponse)
}
