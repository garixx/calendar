package usecase

import (
	"calendar/internals/jwt"
	"calendar/internals/mocks"
	"calendar/internals/models"
	"calendar/internals/validate"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgresTokenUsecase_CreateToken(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	jwtToken, err := jwt.GenerateToken("me")
	require.NoError(t, err, "jwt token should be generated")

	token := models.Token{
		Token: jwtToken,
	}

	repo := mocks.NewMockTokenRepository(ctrl)
	repo.EXPECT().CreateToken(gomock.Eq(token)).Return(token, nil)

	usecase := NewTokenUsecase(repo)
	res, err := usecase.CreateToken(token)

	// then
	require.NoError(t, err)
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, res, token)
}

func TestPostgresTokenUsecase_GetToken(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	jwtToken, err := jwt.GenerateToken("me")
	require.NoError(t, err, "jwt token should be generated")

	token := models.Token{
		Token: jwtToken,
	}

	repo := mocks.NewMockTokenRepository(ctrl)
	repo.EXPECT().GetToken(gomock.Eq(token)).Return(token, nil)

	usecase := NewTokenUsecase(repo)
	res, err := usecase.GetToken(token)

	// then
	require.NoError(t, err)
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, res, token)
}
