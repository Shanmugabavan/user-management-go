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

	createdd, _ := ur.queries.CreateUser(c, db.CreateUserParams{
		UserID:    toPgUUID(user.UserId),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Age:       int32(user.Age),
		Status:    int32(user.Status),
	})

	//const query = `
	//    INSERT INTO users (user_id, first_name, last_name, email, phone, age, status)
	//    VALUES ($1, $2, $3, $4, $5, $6, $7)
	//    RETURNING user_id, email, status
	//`

	//row := ur.connectionPool.QueryRow(
	//	c,
	//	query,
	//	user.UserId,
	//	user.FirstName,
	//	user.LastName,
	//	user.Email,
	//	user.Phone,
	//	user.Age,
	//	user.Status,
	//)

	//var created domain.User
	//err := row.Scan(
	//	&created.UserId,
	//	&created.Email,
	//	&created.Status,
	//)
	//if err != nil {
	//	return domain.User{}, err
	//}

	var created = db.CreateUserRow{
		UserID: createdd.UserID,
		Email:  createdd.Email,
		Status: createdd.Status,
	}

	return created, nil
}

func toPgUUID(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
}
