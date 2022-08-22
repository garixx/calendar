package http

import (
	"calendar/internals/aggregate"
	"calendar/internals/jwt"
	"calendar/internals/mocks"
	"calendar/internals/models"
	"github.com/thanhpk/randstr"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v4"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
)

func TestMiddlewareExecution_valid_token(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	login := randstr.String(7)
	jwtToken, err := jwt.GenerateToken(login)
	require.NoError(t, err, "token should be generated")
	token := models.Token{
		Token: jwtToken,
	}

	tokenUsecase := mocks.NewMockTokenUsecase(ctrl)
	tokenUsecase.EXPECT().GetToken(gomock.Eq(token)).Return(token, nil)

	calendar := aggregate.Calendar{
		TokenCase: tokenUsecase,
	}

	router := mux.NewRouter()
	router.HandleFunc("/", ctxValueToResponse)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", jwtToken)
	w := httptest.NewRecorder()

	router.Use(AuthMiddleware(&calendar))
	router.ServeHTTP(w, req)
	resp := w.Result()

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, login, string(body))
}

func TestMiddlewareExecution_no_token_header(t *testing.T) {
	router := mux.NewRouter()
	router.Use(AuthMiddleware(&aggregate.Calendar{}))
	router.HandleFunc("/", ctxValueToResponse)

	req := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestMiddlewareExecution_token_not_found_in_DB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	login := randstr.String(7)
	jwtToken, err := jwt.GenerateToken(login)
	require.NoError(t, err, "token should be generated")
	token := models.Token{
		Token: jwtToken,
	}

	tokenUsecase := mocks.NewMockTokenUsecase(ctrl)
	tokenUsecase.EXPECT().GetToken(gomock.Eq(token)).Return(models.Token{}, pgx.ErrNoRows)

	calendar := aggregate.Calendar{
		TokenCase: tokenUsecase,
	}

	router := mux.NewRouter()
	router.HandleFunc("/", ctxValueToResponse)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", jwtToken)
	w := httptest.NewRecorder()

	router.Use(AuthMiddleware(&calendar))
	router.ServeHTTP(w, req)
	resp := w.Result()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func ctxValueToResponse(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	_, _ = w.Write([]byte(username))
}
