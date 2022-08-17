package repository

import (
	"calendar/internals/models"
	"calendar/pkg/validate"
	"github.com/chrisyxlee/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
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

func TestGetUser(t *testing.T) {
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
		Query(gomock.Any(), gomock.Any(), gomock.Eq(login), gomock.Eq(passwordHash)).
		Return(pgxRows, nil)

	repo := NewPostgresUserRepository(mockPool)
	user := models.User{
		Login:        login,
		PasswordHash: passwordHash,
	}

	expected := models.User{
		Id:           userUUID,
		Login:        login,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
		IsDeleted:    false,
	}

	res, err := repo.GetUser(user)

	// then
	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

//func TestUserRepositoryMock(t *testing.T) {
//	t.Parallel()
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	// given
//	userUUID := uuid.NewString()
//	login := "myLogin"
//	passwordHash := "passHash"
//	createdAt := time.Now()
//
//	user := models.User{
//		Login:        login,
//		PasswordHash: passwordHash,
//	}
//
//	created := models.User{
//		Id:           userUUID,
//		Login:        login,
//		PasswordHash: passwordHash,
//		CreatedAt:    createdAt,
//		IsDeleted:    false,
//	}
//
//	repo := mocks.NewMockUserRepository(ctrl)
//	repo.
//		EXPECT().
//		CreateUser(gomock.Eq(user)).
//		Return(created, nil)
//
//	res, err := repo.CreateUser(user)
//
//	// then
//	assert.NotNil(t, res)
//	assert.NoError(t, err)
//	assert.Equal(t, created, res)
//}
