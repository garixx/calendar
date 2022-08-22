package http

import (
	"calendar/internals/aggregate"
	"calendar/internals/hashing"
	"calendar/internals/jwt"
	"calendar/internals/models"
	"calendar/internals/validate"
	"encoding/json"
	"errors"
	"net/http"
)

var wrongCredentials = errors.New("wrong credentials")

func loginHandler(calendar *aggregate.Calendar) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.UserRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := calendar.UserCase.GetUserByLogin(payload.Login)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !hashing.CheckPasswordHash(payload.Password, user.PasswordHash) {
			http.Error(w, wrongCredentials.Error(), http.StatusUnauthorized)
			return
		}

		var token string
		if token, err = jwt.GenerateToken(payload.Login); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Left for training purpose only.
		response, err := calendar.TokenCase.CreateToken(models.Token{Token: token})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write(res)
	}
}
