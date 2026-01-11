package repository

import (
	"context"
	"testing"
	"user-management/domain"
	"user-management/internal/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

type mockDB struct {
	DeleteFunc  func(ctx context.Context, userID pgtype.UUID) (pgtype.UUID, error)
	GetUserFunc func(ctx context.Context, userID pgtype.UUID) (db.User, error)
}

func (m *mockDB) Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, nil
}
func (m *mockDB) Query(context.Context, string, ...interface{}) (pgx.Rows, error) { return nil, nil }

func (m *mockDB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	if sql == db.DeleteUser {
		return &mockRow{id: args[0].(pgtype.UUID)}
	}

	if sql == db.GetUser {
		return &mockRow{
			user: domain.User{
				UserId: toUUIDFromPgUUID(args[0].(pgtype.UUID)),
			},
		}
	}
	return &mockRow{}
}

type mockRow struct {
	id   pgtype.UUID
	user domain.User
}

func (r *mockRow) Scan(dest ...interface{}) error {
	if len(dest) == 1 {
		// Handle Delete (returns just UUID)
		if d, ok := dest[0].(*pgtype.UUID); ok {
			*d = r.id
			return nil
		}
	}

	if len(dest) > 1 {
		if d, ok := dest[0].(*pgtype.UUID); ok {
			*d = toPgUUID(r.user.UserId)
		}
		if d, ok := dest[1].(*string); ok {
			*d = r.user.FirstName
		}
		if d, ok := dest[2].(*string); ok {
			*d = r.user.LastName
		}
		return nil
	}
	return pgx.ErrNoRows
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	testID := uuid.New()

	mDB := &mockDB{}
	queries := db.New(mDB)
	repo := &UserRepository{
		connectionPool: nil,
		queries:        queries,
	}

	deletedID, err := repo.Delete(ctx, testID)

	assert.NoError(t, err)
	assert.Equal(t, testID, deletedID)
}

func TestGetById(t *testing.T) {
	ctx := context.Background()
	testID := uuid.New()

	mDB := &mockDB{}

	queries := db.New(mDB)

	repo := UserRepository{
		connectionPool: nil,
		queries:        queries,
	}

	user, err := repo.GetById(ctx, testID)
	if err != nil {
		return
	}

	assert.NoError(t, err)
	assert.Equal(t, testID, user.UserId)

}

func TestUpdateById(t *testing.T) {
	ctx := context.Background()
	testID := uuid.New()

	mDB := &mockDB{}

	queries := db.New(mDB)

	repo := UserRepository{
		connectionPool: nil,
		queries:        queries,
	}

	user := domain.User{
		UserId:    testID,
		FirstName: "Shanmu",
	}

	updatedUserRow, err := repo.Update(ctx, testID, &user)
	if err != nil {
		return
	}

	assert.NoError(t, err)
	assert.NotNil(t, updatedUserRow)
}
