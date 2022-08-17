package http

import (
	"calendar/internals/aggregate"
	"calendar/internals/models"
	"calendar/pkg/validate"
	"encoding/json"
	"net/http"
	"time"
)

func registerHandler(useCases *aggregate.Calendar) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.SignUpPayload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil || validate.Struct(payload) != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := useCases.UserCase.CreateUser(payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := models.SignUpResponse{
			Id:        user.Id,
			Login:     user.Login,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		}

		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(res)
	}
}
