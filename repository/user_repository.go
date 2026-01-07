package repository

import (
	"context"
	"user-management/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	connectionPool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) domain.UserRepository {
	return &userRepository{
		connectionPool: pool,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) (domain.User, error) {
	const query = `
        INSERT INTO users (user_id, first_name, last_name, email, phone, age, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING user_id, email, status
    `

	row := ur.connectionPool.QueryRow(
		c,
		query,
		user.UserId,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Age,
		user.Status,
	)

	var created domain.User
	err := row.Scan(
		&created.UserId,
		&created.Email,
		&created.Status,
	)
	if err != nil {
		return domain.User{}, err
	}

	return created, nil
}
