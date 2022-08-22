package usecase

import (
	"calendar/internals/hashing"
	"calendar/internals/mocks"
	"calendar/internals/models"
	"calendar/internals/validate"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "myLogin"
	password := "password"
	hash, err := hashing.HashPassword(password)
	require.NoError(t, err, "password hash failed")

	payload := models.UserRequest{
		Login:    login,
		Password: password,
	}

	created := models.User{
		Id:           uuid.NewString(),
		Login:        login,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
		IsDeleted:    false,
	}

	repo := mocks.NewMockUserRepository(ctrl)
	repo.EXPECT().CreateUser(gomock.Any()).Return(created, nil)

	usecase := NewUserUsecase(repo)
	res, err := usecase.CreateUser(payload)

	// then
	require.NoError(t, err)
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, res, created)
}

func TestUserUsecase_GetUserByLogin(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "myLogin"
	hash, err := hashing.HashPassword("password")
	require.NoError(t, err, "password hash failed")

	created := models.User{
		Id:           uuid.NewString(),
		Login:        login,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
		IsDeleted:    false,
	}

	repo := mocks.NewMockUserRepository(ctrl)
	repo.EXPECT().GetUserByLogin(gomock.Eq(login)).Return(created, nil)

	usecase := NewUserUsecase(repo)
	res, err := usecase.GetUserByLogin(login)

	// then
	require.NoError(t, err)
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, res, created)
}
