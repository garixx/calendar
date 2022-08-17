package http

import (
	"bytes"
	"calendar/internals/aggregate"
	"calendar/internals/mocks"
	"calendar/internals/models"
	"calendar/pkg/validate"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	login := "myLogin"
	password := "pass"
	hash := "passHash"
	createdAt := time.Now()

	payload := models.SignUpPayload{
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

	var response models.SignUpResponse
	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.NoError(t, validate.Struct(response))
	assert.Equal(t, payload.Login, response.Login)
	assert.Equal(t, created.Id, response.Id)
}
