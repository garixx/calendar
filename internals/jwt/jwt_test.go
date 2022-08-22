package jwt

import (
	"calendar/internals/validate"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
)

func TestTokenCreateAndParse(t *testing.T) {
	s := randstr.String(9)
	token, err := GenerateToken(s)
	require.NoError(t, err)
	assert.NoError(t, validate.GetValidator().Var(token, "jwt"))

	parsed, err := ParseToken(token)
	require.NoError(t, err)
	assert.Equal(t, s, parsed)
}
