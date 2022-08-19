package http

import (
	"bytes"
	"calendar/internals/aggregate"
	"calendar/internals/hashing"
	"calendar/internals/jwt"
	"calendar/internals/mocks"
	"calendar/internals/models"
	"calendar/internals/validate"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "myLogin"
	password := "hashHashHash"
	hash, err := hashing.HashPassword(password)
	require.NoError(t, err, "test pass should be hashed")
	createdAt := time.Now()

	payload := models.UserRequest{
		Login:    login,
		Password: password,
	}

	exist := models.User{
		Id:           uuid.NewString(),
		Login:        login,
		PasswordHash: hash,
		CreatedAt:    createdAt,
		IsDeleted:    false,
	}

	jwtToken, err := jwt.GenerateToken(login)
	require.NoError(t, err, "token should be generated")
	token := models.Token{
		Token: jwtToken,
	}

	userUsecase := mocks.NewMockUserUsecase(ctrl)
	userUsecase.EXPECT().GetUserByLogin(gomock.Eq(login)).Return(exist, nil)

	tokenUsecase := mocks.NewMockTokenUsecase(ctrl)
	tokenUsecase.EXPECT().CreateToken(gomock.Any()).Return(token, nil)

	calendar := aggregate.Calendar{
		UserCase:  userUsecase,
		TokenCase: tokenUsecase,
	}

	handler := loginHandler(&calendar)

	p, err := json.Marshal(payload)
	require.NoError(t, err, "payload should be parsed to JSON")

	req := httptest.NewRequest("POST", "http://localhost/login", bytes.NewBuffer(p))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	var response models.Token
	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.NoError(t, validate.Struct(response))
	assert.Equal(t, token, response)
}
