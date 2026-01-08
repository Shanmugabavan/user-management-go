package user

import (
	"encoding/json"
	"net/http"
	"user-management/api/controller"
	"user-management/api/controller/user/create"
	"user-management/bootstrap"
	"user-management/domain"

	"github.com/go-chi/chi/v5"
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

func (u *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userEntities, err2 := u.GetAll(r.Context())

	usersDtoResponse := make([]domain.User, 0, len(userEntities))

	for _, u := range userEntities {
		usersDtoResponse = append(usersDtoResponse, domain.User{
			UserId:    u.UserId,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
			Age:       int(u.Age),
			Status:    u.Status,
		})
	}
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(usersDtoResponse)
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	userID, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	userEntity, err2 := u.GetById(r.Context(), userID)

	userResponse := domain.User{
		UserId:    userEntity.UserId,
		FirstName: userEntity.FirstName,
		LastName:  userEntity.LastName,
		Email:     userEntity.Email,
		Phone:     userEntity.Phone,
		Age:       userEntity.Age,
		Status:    userEntity.Status,
	}
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(userResponse)
}
