package repository

import (
	"calendar/internals/config"
	"calendar/internals/database"
	"calendar/internals/models"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateUser2(t *testing.T) {
	t.Parallel()

	cfg, err := config.GetPostgresConfig("../../../configs/%s/config.yml")
	require.NoError(t, err, "config init failed")

	client, err := database.NewClient(context.Background(), cfg)
	require.NoError(t, err, "db connect failed")
	defer client.Close()

	repo := NewPostgresUserRepository(client)

	// given
	login := "myLogin1234567893"
	passwordHash := "passHash2345678"
	createdAt := time.Now()

	payload := models.User{
		Login:        login,
		PasswordHash: passwordHash,
	}

	res, err := repo.CreateUser(payload)

	// then
	assert.NotNil(t, res)
	_, errr := uuid.Parse(res.Id.String())
	assert.NoError(t, errr)
	assert.Equal(t, login, res.Login)
	assert.Equal(t, passwordHash, res.PasswordHash)
	assert.True(t, res.CreatedAt.After(createdAt))
	assert.Equal(t, false, res.IsDeleted)
	assert.NoError(t, err)
}

func TestGetExistUser(t *testing.T) {
	t.Parallel()

	cfg, err := config.GetPostgresConfig("../../../configs/%s/config.yml")
	require.NoError(t, err, "config init failed")

	client, err := database.NewClient(context.Background(), cfg)
	require.NoError(t, err, "db connect failed")
	defer client.Close()

	repo := NewPostgresUserRepository(client)

	// given
	login := "me"
	passwordHash := "hashHere"

	payload := models.User{
		Login:        login,
		PasswordHash: passwordHash,
	}

	res, err := repo.GetUser(payload)

	// then
	assert.NotNil(t, res)
	_, errr := uuid.Parse(res.Id.String())
	assert.NoError(t, errr)
	assert.Equal(t, login, res.Login)
	assert.Equal(t, passwordHash, res.PasswordHash)
	assert.NotNil(t, res.CreatedAt)
	assert.Equal(t, false, res.IsDeleted)
	assert.NoError(t, err)
}
