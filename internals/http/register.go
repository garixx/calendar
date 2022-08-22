package http

import (
	"calendar/internals/aggregate"
	"calendar/internals/models"
	"calendar/internals/validate"
	"encoding/json"
	"net/http"
	"time"
)

func registerHandler(calendar *aggregate.Calendar) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.UserRequest

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := calendar.UserCase.CreateUser(payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := models.UserResponse{
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
