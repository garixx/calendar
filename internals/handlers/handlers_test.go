package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const invalidToken = "Invalid token\n"

func TestRegistrationHandlerSuccess(t *testing.T) {
	user := RegistrationRequest{
		Login:    "me",
		Password: "xxx",
		Username: "meuser",
	}

	payload, _ := json.Marshal(user)
	assert.NotNil(t, payload)

	request, err := http.NewRequest("POST", "/register", bytes.NewReader(payload))
	assert.ErrorIs(t, err, nil)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(RegistrationHandler)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestRegistrationHandlerInvalidPayloads(t *testing.T) {
	tests := []struct {
		name     string
		args     RegistrationRequest
		expected int
	}{
		{"invalid login", RegistrationRequest{Login: "&dsa", Password: "", Username: ""}, http.StatusBadRequest},
		{"invalid pass", RegistrationRequest{Login: "aa", Password: " ", Username: ""}, http.StatusBadRequest},
		{"invalid username", RegistrationRequest{Login: "bb", Password: "a12", Username: "a*^"}, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, _ := json.Marshal(tt)
			assert.NotNil(t, payload)

			request, err := http.NewRequest("POST", "/register", bytes.NewReader(payload))
			assert.ErrorIs(t, err, nil)

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(RegistrationHandler)

			handler.ServeHTTP(recorder, request)

			assert.Equal(t, tt.expected, recorder.Code)
		})
	}
}

func TestHealthCheckHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/health", nil)
	assert.Nil(t, err)
	request.Header.Set("token", "valid")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, `{"alive": true}`, recorder.Body.String())
}

func TestHealthCheckHandlerWithoutToken(t *testing.T) {
	request, err := http.NewRequest("GET", "/health", nil)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusForbidden, recorder.Code)
	assert.Equal(t, invalidToken, recorder.Body.String())
}

func TestHealthCheckHandlerWithInvalidToken(t *testing.T) {
	request, err := http.NewRequest("GET", "/health", nil)
	assert.Nil(t, err)
	request.Header.Set("token", "invalid")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusForbidden, recorder.Code)
	assert.Equal(t, invalidToken, recorder.Body.String())
}
