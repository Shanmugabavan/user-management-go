package users

import (
	"time"
	"user-management/api/controller"
	"user-management/bootstrap"
	"user-management/repository"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUserRouter(env *bootstrap.Env, timeout time.Duration, connectionPool *pgxpool.Pool, router *chi.Mux) {
	ur := repository.NewUserRepository(connectionPool)
	uc := &controller.UserController{
		UserRepository: ur,
		Env:            env,
	}

	router.Post("/users", uc.CreateUser)

}
