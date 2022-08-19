package http

import (
	"bytes"
	"calendar/internals/aggregate"
	"calendar/internals/mocks"
	"calendar/internals/models"
	"calendar/internals/validate"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterHandler(t *testing.T) {
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

	created := models.User{
		Id:           uuid.NewString(),
		Login:        login,
		PasswordHash: hash,
		CreatedAt:    createdAt,
		IsDeleted:    false,
	}

	ucase := mocks.NewMockUserUsecase(ctrl)
	ucases := aggregate.Calendar{
		UserCase: ucase,
	}
	handler := registerHandler(&ucases)

	ucase.EXPECT().CreateUser(gomock.Eq(payload)).Return(created, nil)

	p, err := json.Marshal(payload)
	require.NoError(t, err, "payload should be parsed to JSON")

	req := httptest.NewRequest("POST", "http://localhost/register", bytes.NewBuffer(p))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	var response models.UserResponse
	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.NoError(t, validate.Struct(response))
	assert.Equal(t, payload.Login, response.Login)
	assert.Equal(t, created.Id, response.Id)
}

func TestRegisterHandler2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "tooLongLoginText"
	password := "pass"
	//hash := "passHash"
	//createdAt := time.Now()

	payload := models.UserRequest{
		Login:    login,
		Password: password,
	}

	//created := models.User{
	//	Id:           uuid.NewString(),
	//	Login:        login,
	//	PasswordHash: hash,
	//	CreatedAt:    createdAt,
	//	IsDeleted:    false,
	//}

	ucase := mocks.NewMockUserUsecase(ctrl)
	ucases := aggregate.Calendar{
		UserCase: ucase,
	}
	handler := registerHandler(&ucases)

	//ucase.EXPECT().CreateUser(gomock.Eq(payload)).Return(created, nil)

	p, err := json.Marshal(payload)
	require.NoError(t, err, "payload should be parsed to JSON")

	req := httptest.NewRequest("POST", "http://localhost/register", bytes.NewBuffer(p))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	var response models.UserResponse
	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	//assert.NoError(t, validate.Struct(response))
	//assert.Equal(t, payload.Login, response.Login)
	//assert.Equal(t, created.Id, response.Id)
}

func TestRegisterHandler3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	//login := "tooLongLoginText"
	//password := "pass"
	//hash := "passHash"
	//createdAt := time.Now()

	//payload := models.UserRequest{
	//	Login:    login,
	//	Password: password,
	//}

	//created := models.User{
	//	Id:           uuid.NewString(),
	//	Login:        login,
	//	PasswordHash: hash,
	//	CreatedAt:    createdAt,
	//	IsDeleted:    false,
	//}

	ucase := mocks.NewMockUserUsecase(ctrl)
	ucases := aggregate.Calendar{
		UserCase: ucase,
	}
	handler := registerHandler(&ucases)

	//ucase.EXPECT().CreateUser(gomock.Eq(payload)).Return(created, nil)

	//p, err := json.Marshal(payload)
	//require.NoError(t, err, "payload should be parsed to JSON")

	req := httptest.NewRequest("POST", "http://localhost/register", bytes.NewBuffer([]byte(`{"field": "1"`)))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	var response models.UserResponse
	err := json.NewDecoder(resp.Body).Decode(&response)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	//assert.NoError(t, validate.Struct(response))
	//assert.Equal(t, payload.Login, response.Login)
	//assert.Equal(t, created.Id, response.Id)
}

func TestRegisterHandler4(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "myLogin"
	password := "pass"
	//hash := "passHash"
	//createdAt := time.Now()

	payload := models.UserRequest{
		Login:    login,
		Password: password,
	}

	//created := models.User{
	//	Id:           uuid.NewString(),
	//	Login:        login,
	//	PasswordHash: hash,
	//	CreatedAt:    createdAt,
	//	IsDeleted:    false,
	//}

	ucase := mocks.NewMockUserUsecase(ctrl)
	ucases := aggregate.Calendar{
		UserCase: ucase,
	}
	handler := registerHandler(&ucases)

	ucase.EXPECT().CreateUser(gomock.Eq(payload)).Return(models.User{}, errors.New("mock error"))

	p, err := json.Marshal(payload)
	require.NoError(t, err, "payload should be parsed to JSON")

	req := httptest.NewRequest("POST", "http://localhost/register", bytes.NewBuffer(p))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	var response models.UserResponse
	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	//assert.NoError(t, validate.Struct(response))
	//assert.Equal(t, payload.Login, response.Login)
	//assert.Equal(t, created.Id, response.Id)
}
