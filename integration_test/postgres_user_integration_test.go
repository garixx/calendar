package integration_test

import (
	"calendar/internals/config"
	"calendar/internals/database"
	"calendar/internals/hashing"
	"calendar/internals/models"
	"calendar/internals/user/repository"
	"calendar/internals/user/usecase"
	"calendar/pkg/validate"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("error loading env variables:%s", err)
	}
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	cfg, err := config.GetPostgresConfig("../configs/%s/config.yml")
	require.NoError(t, err, "config init failed")

	client, err := database.NewClient(context.Background(), cfg)
	require.NoError(t, err, "db connect failed")
	defer client.Close()

	repo := repository.NewPostgresUserRepository(client)

	// given
	//login := fmt.Sprintf("Test%s", randstr.String(5))
	login := "user"
	password := "password" //fmt.Sprintf("Test%s", randstr.Hex(10))
	hash, err := hashing.HashPassword(password)
	require.NoError(t, err, "should be hashed")
	createdAt := time.Now()

	payload := models.User{
		Login:        login,
		PasswordHash: hash,
	}

	res, err := repo.CreateUser(payload)

	// then
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, login, res.Login)
	assert.Equal(t, hash, res.PasswordHash)
	assert.True(t, res.CreatedAt.After(createdAt))
	assert.Equal(t, false, res.IsDeleted)
	assert.NoError(t, err)
}

func TestCreateUserCase(t *testing.T) {
	t.Parallel()

	cfg, err := config.GetPostgresConfig("../configs/%s/config.yml")
	require.NoError(t, err, "config init failed")

	client, err := database.NewClient(context.Background(), cfg)
	require.NoError(t, err, "db connect failed")
	defer client.Close()

	repo := repository.NewPostgresUserRepository(client)
	userCase := usecase.NewUserUsecase(repo)

	// given
	login := fmt.Sprintf("Test%s", randstr.String(7))
	password := "Test"
	createdAt := time.Now()

	payload := models.UserRequest{
		Login:    login,
		Password: password,
	}

	res, err := userCase.CreateUser(payload)

	// then
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, login, res.Login)
	assert.Equal(t, password+"hashing", res.PasswordHash)
	assert.True(t, res.CreatedAt.After(createdAt))
	assert.Equal(t, false, res.IsDeleted)
	assert.NoError(t, err)
}

func TestGetExistUser(t *testing.T) {
	t.Parallel()

	cfg, err := config.GetPostgresConfig("../configs/%s/config.yml")
	require.NoError(t, err, "config init failed")

	client, err := database.NewClient(context.Background(), cfg)
	require.NoError(t, err, "db connect failed")
	defer client.Close()

	repo := repository.NewPostgresUserRepository(client)

	// given
	login := "user"
	passwordHash, err := hashing.HashPassword("password")
	passwordHash2, err2 := hashing.HashPassword("password")
	log.Println(passwordHash2, err2)
	require.NoError(t, err, "hash failed")

	payload := models.User{
		Login:        login,
		PasswordHash: passwordHash,
	}

	res, err := repo.GetUser(payload)

	// then
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, login, res.Login)
	assert.Equal(t, passwordHash, res.PasswordHash)
	assert.NotNil(t, res.CreatedAt)
	assert.Equal(t, false, res.IsDeleted)
	assert.NoError(t, err)
}

//func TestGetExistUser2(t *testing.T) {
//	t.Parallel()
//	os.Setenv("environment", "qa")
//	os.Setenv("POSTGRES_PASSWORD", "postgres")
//	cfg, err := config.GetPostgresConfig("../configs/%s/config.yml")
//	require.NoError(t, err, "config init failed")
//
//	client, err := database.NewClient(context.Background(), cfg)
//	require.NoError(t, err, "db connect failed")
//	defer client.Close()
//
//	repo := repository.NewPostgresUserRepository(client)
//
//	// given
//	login := "me"
//	passwordHash := "hashHere"
//
//	payload := models.User2{
//		Login:        login,
//		PasswordHash: passwordHash,
//	}
//
//	res, err := repo.GetUser2(payload)
//
//	validate := validator.New()
//	assert.NoError(t, validate.Struct(res))
//	// then
//	assert.Equal(t, login, res.Login)
//	assert.Equal(t, passwordHash, res.PasswordHash)
//	assert.NotNil(t, res.CreatedAt)
//	assert.Equal(t, false, res.IsDeleted)
//	assert.NoError(t, err)
//}
