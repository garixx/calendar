package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

var allowed = "bcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789."
var notAllowed = " .±"

// RegistrationRequest register request payload
type RegistrationRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
}

// LoginRequest login request payload
type LoginRequest struct {
	Login string `json:"login,omitempty"`
	Token string `json:"token,omitempty"`
}

// logout payload
type LogoutRequest struct {
	Token string `json:"token,omitempty"`
}

func validateToken(r *http.Request) bool {
	t := r.Header.Get("token")
	//TODO: replace with real validation
	return t == "valid"
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if !validateToken(r) {
		http.Error(w, "Invalid token", http.StatusForbidden)
		return
	}
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var auth RegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&auth)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !strings.ContainsAny(auth.Login, allowed) {
		http.Error(w, "Bad name", http.StatusBadRequest)
		return
	}
	if strings.Contains(auth.Password, notAllowed) {
		http.Error(w, "Bad password", http.StatusBadRequest)
		return
	}
	if !strings.ContainsAny(auth.Username, allowed) {
		http.Error(w, "Bad username", http.StatusBadRequest)
		return
	}
	//TODO: add to users DB
	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//var auth LoginRequest
	//
	//// Try to decode the request body into the struct. If there is an error,
	//// respond to the client with the error message and a 400 status code.
	//err := json.NewDecoder(r.Body).Decode(&auth)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//logrus.Info("Received login request: username:" + auth.Username + ", password:" + auth.Password)
	//
	//newToken := base64.StdEncoding.EncodeToString([]byte(auth.Username + auth.Password))
	//tokenUsers["Bearer "+newToken] = auth.Username
	//
	//logrus.Infof("tokens: %v", tokenUsers)
	//fmt.Fprintf(w, newToken)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//var auth LogoutRequest
	//
	//err := json.NewDecoder(r.Body).Decode(&auth)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//logrus.Info("Received logout request")
	//
	//delete(tokenUsers, auth.Token)
	//
	//logrus.Infof("tokens: %v", tokenUsers)
	//w.WriteHeader(http.StatusNoContent)
}
