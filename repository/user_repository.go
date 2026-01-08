package repository

import (
	"context"
	"user-management/domain"
	"user-management/internal/db"
	"user-management/internal/validator"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	connectionPool *pgxpool.Pool
	queries        *db.Queries
}

func NewUserRepository(pool *pgxpool.Pool) domain.UserRepository {
	return &userRepository{
		connectionPool: pool,
		queries:        db.New(pool),
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) (db.CreateUserRow, error) {
	valError := validator.Validate.Struct(user)
	if valError != nil {
		return db.CreateUserRow{}, valError
	}

	createdd, err := ur.queries.CreateUser(c, db.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Age:       int32(user.Age),
		Status:    int32(user.Status),
	})

	created := db.CreateUserRow{
		UserID: createdd.UserID,
		Email:  createdd.Email,
		Status: createdd.Status,
	}

	return created, err
}

func (ur *userRepository) GetAll(c context.Context) ([]domain.User, error) {
	dbUsers, err := ur.queries.GetAllUsers(c)

	if err != nil {
		return nil, err
	}

	users := make([]domain.User, 0, len(dbUsers))

	for _, u := range dbUsers {
		users = append(users, domain.User{
			UserId:    toUUIDfromPgUUID(u.UserID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
			Age:       int(u.Age),
			Status:    domain.UserStatus(u.Status),
		})
	}

	return users, nil
}

func (ur *userRepository) GetById(c context.Context, id uuid.UUID) (domain.User, error) {
	dbUser, err := ur.queries.GetUser(c, toPgUUID(id))

	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		UserId:    toUUIDfromPgUUID(dbUser.UserID),
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		Phone:     dbUser.Phone,
		Age:       int(dbUser.Age),
		Status:    domain.UserStatus(dbUser.Status),
	}

	return user, nil
}

func toPgUUID(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
}

func toUUIDfromPgUUID(id pgtype.UUID) uuid.UUID {
	if !id.Valid {
		return uuid.Nil
	}
	return id.Bytes
}
