package usecase

import (
	"calendar/internals/mocks"
	"calendar/internals/models"
	"calendar/pkg/validate"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "myLogin"
	password := "pass"
	hash := "passHash"
	createdAt := time.Now()

	payload := models.UserRequest{
		Login:    login,
		Password: password,
	}

	user := models.User{Login: "myLogin", PasswordHash: "passhash"}

	created := models.User{
		Id:           uuid.NewString(),
		Login:        login,
		PasswordHash: hash,
		CreatedAt:    createdAt,
		IsDeleted:    false,
	}

	repo := mocks.NewMockUserRepository(ctrl)
	repo.EXPECT().CreateUser(gomock.Eq(user)).Return(created, nil)

	usecase := NewUserUsecase(repo)
	res, err := usecase.CreateUser(payload)

	// then
	require.NoError(t, err)
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, res, created)
}
