package http

import (
	"calendar/internals/aggregate"
	"calendar/internals/jwt"
	"calendar/internals/models"
	"context"
	"net/http"
)

func AuthMiddleware(app *aggregate.Calendar) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			if token == "" {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			t, err := app.TokenCase.GetToken(models.Token{Token: token})
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			user, err := jwt.ParseToken(t.Token)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), "username", user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// Another approach
//type authMiddleware struct {
//	app aggregate.Calendar
//}
//
//func NewAuthMiddleware(app aggregate.Calendar) authMiddleware {
//	return authMiddleware{app: app}
//}
//
//func (a *authMiddleware) Authenticate(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//		h.ServeHTTP(w, r)
//	})
//}
