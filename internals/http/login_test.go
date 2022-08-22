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
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginHandler_success(t *testing.T) {
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

func Test_loginHandlerNegative_brokenJSON(t *testing.T) {
	handler := loginHandler(&aggregate.Calendar{})

	req := httptest.NewRequest("POST", "http://localhost/login", bytes.NewBuffer([]byte(`{"login": "xx"`)))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	var response models.UserResponse
	err := json.NewDecoder(resp.Body).Decode(&response)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_loginHandlerNegative_login_too_long(t *testing.T) {
	login := "TooLongNameXxx"
	password := "hashHashHash"

	payload := models.UserRequest{
		Login:    login,
		Password: password,
	}

	handler := loginHandler(&aggregate.Calendar{})

	p, err := json.Marshal(payload)
	require.NoError(t, err, "payload should be parsed to JSON")

	req := httptest.NewRequest("POST", "http://localhost/login", bytes.NewBuffer(p))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	var response models.Token
	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_loginHandlerNegative_user_not_created(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "myLogin"
	password := "hashHashHash"

	payload := models.UserRequest{
		Login:    login,
		Password: password,
	}

	userUsecase := mocks.NewMockUserUsecase(ctrl)
	userUsecase.EXPECT().GetUserByLogin(gomock.Eq(login)).Return(models.User{}, pgx.ErrNoRows)

	calendar := aggregate.Calendar{
		UserCase: userUsecase,
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

	assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_loginHandlerNegative_DB_access_error(t *testing.T) {
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

	userUsecase := mocks.NewMockUserUsecase(ctrl)
	userUsecase.EXPECT().GetUserByLogin(gomock.Eq(login)).Return(exist, nil)

	tokenUsecase := mocks.NewMockTokenUsecase(ctrl)
	tokenUsecase.EXPECT().CreateToken(gomock.Any()).Return(models.Token{}, errors.New("mock error"))

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

	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func Test_loginHandlerNegative_wrong_token(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "myLogin"
	password := "hashHashHash"
	hash, err := hashing.HashPassword("anotherHash")
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

	userUsecase := mocks.NewMockUserUsecase(ctrl)
	userUsecase.EXPECT().GetUserByLogin(gomock.Eq(login)).Return(exist, nil)

	calendar := aggregate.Calendar{
		UserCase: userUsecase,
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

	assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
