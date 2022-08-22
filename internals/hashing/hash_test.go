package hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBcryptingIsEasy(t *testing.T) {
	pass := "mypassword"
	hp, err := HashPassword(pass)
	require.NoError(t, err, "HashPassword failed")

	isValid := CheckPasswordHash(pass, hp)
	assert.True(t, isValid)

	notPass := "notthepass"
	errorGot := CheckPasswordHash(hp, notPass)
	assert.False(t, errorGot)
}

func TestBcryptingIsEasy2(t *testing.T) {
	pass := "password"
	hp := "$2a$14$1MEzFbcJXlcBQ8/26tTwK.yF7A2k3TcjWhetZ3CZlvNxUz5LX/NwG"

	res := CheckPasswordHash(pass, hp)
	assert.True(t, res)
}
