package repository

import (
	"calendar/internals/models"
	"calendar/internals/validate"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/chrisyxlee/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostgresUserRepository_CreateUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	userUUID := uuid.NewString()
	login := "myLogin"
	passwordHash := "passHash"
	createdAt := time.Now()

	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	expected := pgxpoolmock.NewRow(userUUID, login, passwordHash, createdAt, false)
	mockPool.EXPECT().
		QueryRow(gomock.Any(), gomock.Any(), gomock.Eq(login), gomock.Eq(passwordHash)).
		Return(expected)

	repo := NewPostgresUserRepository(mockPool)
	user := models.User{
		Login:        login,
		PasswordHash: passwordHash,
	}

	created := models.User{
		Id:           userUUID,
		Login:        login,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
		IsDeleted:    false,
	}

	res, err := repo.CreateUser(user)

	// then
	assert.NoError(t, err)
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, created, res)
}

func TestPostgresUserRepository_GetUserByLogin_get_exist(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	userUUID := uuid.NewString()
	login := "me"
	passwordHash := "hashHere"
	createdAt := time.Now()

	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	columns := []string{"id", "login", "password_hash", "created_at", "is_deleted"}
	pgxRows := pgxpoolmock.NewRows(columns).AddRow(userUUID, login, passwordHash, createdAt, false).ToPgxRows()

	mockPool.EXPECT().
		Query(gomock.Any(), gomock.Any(), gomock.Eq(login)).
		Return(pgxRows, nil)

	repo := NewPostgresUserRepository(mockPool)

	expected := models.User{
		Id:           userUUID,
		Login:        login,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
		IsDeleted:    false,
	}

	res, err := repo.GetUserByLogin(login)

	// then
	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestPostgresUserRepository_GetUserByLogin_absent_in_DB(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "me"

	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Eq(login)).Return(nil, pgx.ErrNoRows)

	repo := NewPostgresUserRepository(mockPool)
	_, err := repo.GetUserByLogin(login)
	assert.Error(t, err)
}

func TestPostgresUserRepository_GetUserByLogin_empty_response_from_DB(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "me"

	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	columns := []string{"id", "login", "password_hash", "created_at", "is_deleted"}
	pgxRows := pgxpoolmock.NewRows(columns).ToPgxRows()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Eq(login)).Return(pgxRows, nil)

	repo := NewPostgresUserRepository(mockPool)

	_, err := repo.GetUserByLogin(login)
	assert.Error(t, err)
}
